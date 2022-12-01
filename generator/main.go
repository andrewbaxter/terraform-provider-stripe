package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"sort"
	"strings"

	"github.com/andrewbaxter/terraform-provider-stripe/shared"
	"github.com/dave/jennifer/jen"
)

func DeepCopy[T any](x T) T {
	a, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	x = shared.Default[T]()
	err = json.Unmarshal(a, &x)
	if err != nil {
		panic(err)
	}
	return x
}

func Last[T any](x []T) T {
	return x[len(x)-1]
}

func DigOpt[T any](obj any, field string) *T {
	v, found := obj.(map[string]any)[field]
	if !found {
		return nil
	}
	v1, isT := v.(T)
	if !isT {
		panic("dug value not convertible")
	}
	return &v1
}

func IndexOpt[T any](arr []T, i int) T {
	if i >= len(arr) {
		return shared.Default[T]()
	}
	return arr[i]
}

func Bury(obj any, field string, data string) {
	m := obj.(map[string]any)
	old := m[field]
	if data != "" && old != "" {
		m[field] = fmt.Sprintf("%s %s", data, old)
	} else {
		m[field] = data
	}
}

func Map[T any, G any](x []T, m func(v T) G) []G {
	out := make([]G, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = m(x[i])
	}
	return out
}

func MapKVs[K comparable, V any, G any](x map[K]V, m func(k K, v V) G) []G {
	out := []G{}
	for k, v := range x {
		out = append(out, m(k, v))
	}
	return out
}

func Filter[T any](x []T, m func(v T) bool) []T {
	out := []T{}
	for _, v := range x {
		if m(v) {
			out = append(out, v)
		}
	}
	return out
}

func Flatten[T any](x [][]T) []T {
	total := 0
	for _, y := range x {
		total += len(y)
	}
	out := make([]T, 0, total)
	for _, y := range x {
		out = append(out, y...)
	}
	return out
}

type Usable[T any] struct {
	used bool
	f    func() T
}

func (u *Usable[T]) Use() T {
	u.used = true
	return u.f()
}

func (u *Usable[T]) WasUsed() bool {
	return u.used
}

func Unused[T any](f func() T) *Usable[T] {
	return &Usable[T]{
		used: false,
		f:    f,
	}
}

func FromSnake(s string) []string {
	return strings.Split(s, "_")
}

func ToSnake(parts []string) string {
	return strings.Join(parts, "_")
}

func ToCamel(parts []string) string {
	out := strings.Builder{}
	for _, part := range parts {
		out.WriteString(strings.ToUpper(string(part[0])) + part[1:])
	}
	return out.String()
}

func Reunmarshal[T any](bla any) T {
	var out T
	err := json.Unmarshal([]byte(shared.Must(json.Marshal(bla))), &out)
	if err != nil {
		panic(err)
	}
	return out
}

type OpenApiString struct {
	Description string   `json:"description"`
	Enum        []string `json:"enum"`
	MaxLength   *int     `json:"maxLength"`
}

type OpenApiArray struct {
	Description string `json:"description"`
	Items       any    `json:"items"`
}

type OpenApiObj struct {
	Type                 string         `json:"type"`
	Description          string         `json:"description"`
	AdditionalProperties any            `json:"additionalProperties"`
	Properties           map[string]any `json:"properties"`
	Required             []string       `json:"required"`
}

const SchemaPkg = "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
const DiagPkg = "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
const CtyPkg = "github.com/hashicorp/go-cty/cty"
const SharedPkg = "github.com/andrewbaxter/terraform-provider-stripe/shared"
const DiagId = "out"

func CondAddDiagErr(exit bool, path jen.Code, t string, messageArgs ...jen.Code) *jen.Statement {
	messageArgs = append([]jen.Code{
		jen.Lit(t + ": %s"),
	}, messageArgs...)
	messageArgs = append(messageArgs, jen.Id("err"))
	statements := []jen.Code{
		jen.Id(DiagId).Op("=").Id("append").Call(jen.Id(DiagId), jen.Qual(DiagPkg, "Diagnostic").Values(jen.Dict{
			jen.Id("Severity"):      jen.Qual(DiagPkg, "Error"),
			jen.Id("Summary"):       jen.Qual("fmt", "Sprintf").Call(messageArgs...),
			jen.Id("AttributePath"): path,
		})),
	}
	if exit {
		statements = append(statements, jen.Return(jen.Id(DiagId)))
	}
	return jen.If(jen.Id("err").Op("!=").Nil()).Block(statements...)
}

type Node interface {
	// Generate a partial field definition
	GenField() jen.Dict
	// Generate a complete schema for use in a collection (map/list)
	GenElem() jen.Code
	// Read the value from apiSource and place it in tfDest
	ReadApi(apiSource jen.Code, tfDest TfDestVal) []jen.Code
	// Expression that checks the raw (any) value in tfSource and returns the value in the shape the api expects.
	// For objects, this means turning a flattened map into nested maps for nested object specs.
	ValidateSetApi(update bool, tfPath *Usable[jen.Code], tfSource TfSourceVal) jen.Code
	IsNotDefault(id jen.Code) jen.Code
}

var FixupFieldDefault = map[string]map[string]jen.Code{
	"price": {
		"billing_scheme": jen.Lit("per_unit"),
		"tax_behavior":   jen.Lit("unspecified"),
	},
}
var FixupFieldString = map[string]bool{
	"price/tiers/[]/up_to":                  true,
	"price/currency_options/tiers/[]/up_to": true,
}

func BuildNode(
	path []string,
	seenRefs map[string]bool,
	refs map[string]OpenApiObj,
	createSpec any,
	updateSpec any,
	getSpec any,
) *Node {
	if FixupFieldString[strings.Join(path, "/")] {
		return shared.Pointer[Node](&NodeString{})
	}
	deref := func(s any) (any, string) {
		if s == nil {
			return nil, ""
		}
		ref := DigOpt[string](s, "$ref")
		if ref == nil {
			return s, ""
		}
		elemName := Last(strings.Split(*ref, "/"))
		return Reunmarshal[any](refs[elemName]), elemName
	}
	filterAnyOf := func(anyOf []any) []any {
		filtered := []any{}
		var anyPrim any = nil
		str := false
		allowedValues := []string{}
		for _, sub := range anyOf {
			sub, _ := deref(sub)
			type_ := shared.Dig[string](sub, "type")
			switch type_ {
			case "integer", "number", "boolean":
				anyPrim = sub
				allowedValues = append(allowedValues, type_)
			case "string":
				enum := shared.Dig[[]any](sub, "enum")
				if enum != nil && len(enum) == 1 && enum[0].(string) == "" {
					continue
				}
				str = true
				anyPrim = sub
				if len(enum) > 0 {
					for _, e := range enum {
						allowedValues = append(allowedValues, fmt.Sprintf("\"%s\"", e.(string)))
					}
				} else {
					allowedValues = append(allowedValues, type_)
				}
			case "object":
				// nop
			case "array":
				// nop
			default:
				panic("TODO")
			}
			filtered = append(filtered, sub)
		}
		if str {
			return []any{map[string]any{
				"type":        "string",
				"description": fmt.Sprintf("Allowed values: %s", strings.Join(allowedValues, ", ")),
			}}
		} else if anyPrim != nil {
			return []any{anyPrim}
		} else {
			return filtered
		}
	}

	{
		var ref_ = ""
		if s, ref := deref(createSpec); ref != "" {
			if seenRefs[ref] {
				log.Printf("Encountered infinite ref loop at %s", ref)
				return nil
			}
			ref_ = ref
			createSpec = s
		}
		if s, ref := deref(updateSpec); ref != "" {
			if seenRefs[ref] {
				log.Printf("Encountered infinite ref loop at %s", ref)
				return nil
			}
			ref_ = ref
			updateSpec = s
		}
		if s, ref := deref(getSpec); ref != "" {
			if seenRefs[ref] {
				log.Printf("Encountered infinite ref loop at %s", ref)
				return nil
			}
			ref_ = ref
			getSpec = s
		}
		if ref_ != "" {
			seenRefs = DeepCopy(seenRefs)
			seenRefs[ref_] = true
		}
	}
	deanyof := func(s any) any {
		// get rid of weird mismatches where create is obj, get is anyof with 1 obj option
		if s == nil {
			return nil
		}
		anyof := shared.Dig[[]any](s, "anyOf")
		anyof = filterAnyOf(anyof)
		if len(anyof) == 1 {
			return anyof[0]
		}
		return s
	}
	createSpec = deanyof(createSpec)
	updateSpec = deanyof(updateSpec)
	getSpec = deanyof(getSpec)
	{
		ref_ := ""
		if s, ref := deref(createSpec); ref != "" {
			if seenRefs[ref] {
				log.Printf("Encountered infinite ref loop at %s", ref)
				return nil
			}
			ref_ = ref
			createSpec = s
		}
		if s, ref := deref(updateSpec); ref != "" {
			if seenRefs[ref] {
				log.Printf("Encountered infinite ref loop at %s", ref)
				return nil
			}
			ref_ = ref
			updateSpec = s
		}
		if s, ref := deref(getSpec); ref != "" {
			if seenRefs[ref] {
				log.Printf("Encountered infinite ref loop at %s", ref)
				return nil
			}
			ref_ = ref
			getSpec = s
		}
		if ref_ != "" {
			seenRefs = DeepCopy(seenRefs)
			seenRefs[ref_] = true
		}
	}

	desc := ""
	if createSpec != nil {
		desc = desc + shared.Dig[string](createSpec, "description")
	}
	if getSpec != nil {
		desc = desc + shared.Dig[string](getSpec, "description")
	}

	var type_ = ""
	if createSpec != nil {
		type_ = shared.Dig[string](createSpec, "type")
	}
	var getType_ = ""
	if getSpec != nil {
		type_ = shared.Dig[string](getSpec, "type")
		if type_ == "" {
			type_ = getType_
		}
	}
	if type_ != "" && getType_ != "" && type_ != getType_ {
		// bad spec design
		log.Printf("Node with mismatched types %s and %s, skipping", type_, getType_)
		return nil
	}
	if type_ == "" {
		m := func(s any) []any {
			if s == nil {
				return nil
			}
			return shared.Dig[[]any](s, "anyOf")
		}
		createAnyOf := filterAnyOf(m(createSpec))
		updateAnyOf := filterAnyOf(m(updateSpec))
		getAnyOf := filterAnyOf(m(getSpec))
		if createAnyOf != nil || getAnyOf != nil {
			type El struct {
				Create any // not nil
				Update any
				Get    any
			}
			filtered := []El{}
			allObj := true
			if len(createAnyOf) > 0 {
				for i, sub := range createAnyOf {
					filtered = append(filtered, El{
						Create: sub,
						Update: IndexOpt(updateAnyOf, i),
						Get:    IndexOpt(getAnyOf, i),
					})
				}
			} else {
				for i, sub := range getAnyOf {
					filtered = append(filtered, El{
						Create: nil,
						Update: IndexOpt(updateAnyOf, i),
						Get:    sub,
					})
				}
			}
			if len(filtered) < 2 {
				panic("ASSERTION")
			}
			if allObj {
				options := []*NodeObj{}
				for _, o := range filtered {
					node := BuildNode(path, seenRefs, refs, o.Create, o.Update, o.Get)
					if node == nil {
						continue
					}
					objNode := (*node).(*NodeObj)
					if !objNode.checkAnyReq() {
						continue
					}
					options = append(options, objNode)
				}
				if len(options) == 1 {
					return shared.Pointer[Node](options[0])
				}
				if len(options) == 0 {
					log.Printf("AnyOf field with no viable options, skipping")
					return nil
				}
				return shared.Pointer[Node](&NodeAnyOfObjs{
					Options: options,
				})
			}
			panic("ASSERTION")
		}
	}
	switch type_ {
	case "string":
		var spec OpenApiString
		if createSpec != nil {
			spec = Reunmarshal[OpenApiString](createSpec)
		} else {
			spec = Reunmarshal[OpenApiString](getSpec)
		}
		enum := spec.Enum
		return shared.Pointer[Node](&NodeString{
			Enum: enum,
		})
	case "boolean":
		return shared.Pointer[Node](&NodeBool{})
	case "integer":
		return shared.Pointer[Node](&NodeInt{})
	case "number":
		return shared.Pointer[Node](&NodeFloat{})
	case "object":
		m := func(s any) *OpenApiObj {
			if s == nil {
				return nil
			}
			return shared.Pointer(Reunmarshal[OpenApiObj](s))
		}
		createSpec := m(createSpec)
		updateSpec := m(updateSpec)
		getSpec := m(getSpec)
		m1 := func(s *OpenApiObj) map[string]any {
			if s == nil {
				return nil
			}
			return s.Properties
		}
		updateFields := m1(updateSpec)
		getFields := m1(getSpec)
		if (createSpec != nil && len(createSpec.Properties) > 0) || len(getFields) > 0 {
			fields := []*NodeObjField{}
			fieldsByKey := map[string]*NodeObjField{}
			if createSpec != nil {
				createRequired := map[string]bool{}
				for _, paramName := range createSpec.Required {
					createRequired[paramName] = true
				}
				for propName, prop := range createSpec.Properties {
					propNode := BuildNode(append(path, propName), seenRefs, refs, prop, updateFields[propName], getFields[propName])
					if propNode == nil {
						continue
					}
					var behavior NodeObjFieldBehavior
					if createRequired[propName] {
						behavior = NBUserRequired
					} else {
						behavior = NBUserOptional
					}
					field := &NodeObjField{
						Description: shared.Dig[string](prop, "description"),
						Key:         propName,
						Behavior:    behavior,
						Updatable:   shared.InMap(propName, updateFields),
						Readable:    shared.InMap(propName, getFields),
						Spec:        *propNode,
					}
					fields = append(fields, field)
					fieldsByKey[propName] = field
				}
			}

			for propName, prop := range getFields {
				desc := shared.Dig[string](prop, "description")
				if createField, has := fieldsByKey[propName]; has {
					if createField.Description == "" {
						createField.Description = desc
					}
					continue
				}
				propNode := BuildNode(append(path, propName), seenRefs, refs, nil, updateFields[propName], getFields[propName])
				if propNode == nil {
					continue
				}
				field := &NodeObjField{
					Description: desc,
					Key:         propName,
					Behavior:    NBComputed,
					Readable:    true,
					Spec:        *propNode,
				}
				fields = append(fields, field)
			}

			if len(fields) == 0 {
				// Some crazy empty objects in the api...
				log.Printf("Object field with no viable fields, skipping")
				return nil
			}

			sort.Slice(fields, func(i, j int) bool {
				return fields[i].Key < fields[j].Key
			})

			return shared.Pointer[Node](&NodeObj{
				Fields: fields,
			})
		} else if (createSpec != nil && createSpec.AdditionalProperties != nil) || (getSpec != nil && getSpec.AdditionalProperties != nil) {
			m := func(s *OpenApiObj) any {
				if s == nil {
					return nil
				}
				return s.AdditionalProperties
			}
			elem := BuildNode(
				path,
				seenRefs,
				refs,
				m(createSpec),
				m(updateSpec),
				m(getSpec),
			)
			if elem == nil {
				return nil
			}
			if objElem, isObj := (*elem).(*NodeObj); isObj {
				objElem.FakeMapField = true
				return shared.Pointer[Node](&NodeFakeMap{
					Elem: *elem,
				})
			} else {
				return shared.Pointer[Node](&NodeMap{
					Elem: *elem,
				})
			}
		} else {
			log.Printf("Object missing properties or additionalProperties, skipping")
			return nil
		}
	case "array":
		m := func(s any) any {
			if s == nil {
				return nil
			}
			return Reunmarshal[OpenApiArray](s).Items
		}
		elem := BuildNode(append(path, "[]"), seenRefs, refs, m(createSpec), m(updateSpec), m(getSpec))
		if elem == nil {
			return nil
		}
		return shared.Pointer[Node](&NodeArray{
			Elem: *elem,
		})
	default:
		panic(fmt.Sprintf("Unhandled schema param type [[%s]]", type_))
	}
}

func main() {
	// Load + preprocess spec
	type OpenApiMethod struct {
		Description string `json:"description"`
		RequestBody struct {
			Content struct {
				FormUrlencoded struct {
					Schema any `json:"schema"`
				} `json:"application/x-www-form-urlencoded"`
			} `json:"content"`
		} `json:"requestBody"`
		Responses struct {
			R200 struct {
				Content struct {
					Json struct {
						Schema struct {
							Ref string `json:"$ref"`
						} `json:"schema"`
					} `json:"application/json"`
				} `json:"content"`
			} `json:"200"`
		} `json:"responses"`
	}
	var spec struct {
		Components struct {
			Schemas map[string]OpenApiObj `json:"schemas"`
		} `json:"components"`
		Paths map[string]map[string]OpenApiMethod `json:"paths"`
	}
	err := json.Unmarshal(shared.Must(ioutil.ReadFile("stripe_openapi/openapi/spec3.json")), &spec)
	if err != nil {
		panic(err)
	}
	instanceMethods := map[string]map[string]OpenApiMethod{}
	for path, methods := range spec.Paths {
		parts := strings.Split(path, "/")
		if Last(parts)[0] == '{' {
			instanceMethods[strings.Join(parts[0:len(parts)-1], "/")] = methods
		}
	}

	jenResources := jen.Dict{}
	for path, methods := range spec.Paths {
		if strings.Contains(path, "{") {
			//log.Printf("Skipping %s, parameterized = no collection", path)
			continue
		}
		post, hasPost := methods["post"]
		if !hasPost {
			//log.Printf("Skipping %s, no post method", path)
			continue
		}
		objName := Last(strings.Split(post.Responses.R200.Content.Json.Schema.Ref, "/"))
		objSpec := spec.Components.Schemas[objName]

		var update any
		{
			updateRaw, hasUpdate := instanceMethods[path]["post"]
			if hasUpdate {
				update = updateRaw.RequestBody.Content.FormUrlencoded.Schema
			}
		}

		_, hasDelete := instanceMethods[path]["delete"]

		if post.RequestBody.Content.FormUrlencoded.Schema == nil {
			log.Printf("Skipping %s, no post schema", path)
			continue
		}
		fields := (*BuildNode(
			[]string{objName},
			map[string]bool{},
			spec.Components.Schemas,
			post.RequestBody.Content.FormUrlencoded.Schema,
			update,
			Reunmarshal[any](objSpec),
		)).(*NodeObj)

		hasActive := false
		noneUpdatable := true
		idField := ""
		for _, f := range fields.Fields {
			if f.Key == "active" {
				hasActive = true
			}
			if f.Key == "id" {
				idField = f.Key
				f.Behavior = NBComputed
			}
			if f.Updatable {
				noneUpdatable = false
			}
			default_, hasDefault := FixupFieldDefault[objName][f.Key]
			if hasDefault {
				f.Default = default_
			}
		}
		if idField == "" {
			log.Printf("Skipping %s, no id field", path)
			continue
		}
		if noneUpdatable {
			log.Printf("Skipping %s, no updatable fields", path)
			continue
		}
		if !hasDelete && !hasActive {
			log.Printf("Skipping %s, no active and no delete", path)
			continue
		}
		fields.Fields = Filter(fields.Fields, func(v *NodeObjField) bool {
			switch v.Key {
			case "expand", "metadata", "active": // api meta, not resource
				return false
			case "last_finalization_error": // weird fields that I don't want to support
				return false
			default:
				return true
			}
		})

		collNameParts := []string{}
		for _, seg := range strings.Split(objName, ".") {
			collNameParts = append(collNameParts, FromSnake(seg)...)
		}
		collName := ToSnake(collNameParts)
		facilId := "f"
		getClientStmt := jen.Id(facilId).Op(":=").Id("f0").Assert(jen.Op("*").Id("Facilitator"))
		createDiagStmt := jen.Id(DiagId).Op(":=").Qual(DiagPkg, "Diagnostics").Values(jen.Dict{})
		inputName := "d"
		idExpr := jen.Id(inputName).Dot("Id").Call()
		idUrlPathExpr := jen.Qual("fmt", "Sprintf").Call(jen.Lit(path+"/%v"), jen.Id(inputName).Dot("Get").Call(jen.Lit(idField)))
		emptyPathExpr := Unused(func() jen.Code { return jen.Qual(CtyPkg, "Path").Values(jen.Dict{}) })

		topSchemaFields := jen.Dict{}
		fields.genFieldInner(topSchemaFields, false, "")

		jenResFile := jen.NewFile("main")
		funcName := "resources_" + collName
		jenResFile.Func().Id(funcName).Params().Op("*").Qual(SchemaPkg, "Resource").Block(
			jen.Return(jen.Op("&").Qual(SchemaPkg, "Resource").Values(jen.Dict{
				jen.Id("Schema"): jen.Map(jen.String()).Op("*").Qual(SchemaPkg, "Schema").Values(topSchemaFields),

				jen.Id("CreateContext"): jen.Func().Params(
					jen.Id("ctx").Qual("context", "Context"),
					jen.Id(inputName).Op("*").Qual(SchemaPkg, "ResourceData"),
					jen.Id("f0").Any(),
				).Qual(DiagPkg, "Diagnostics").BlockFunc(func(g *jen.Group) {
					paramsId := "params"
					g.Add(getClientStmt)
					g.Add(createDiagStmt)
					g.Add(jen.Id(paramsId).Op(":=").Add(fields.ValidateSetApi(
						false,
						emptyPathExpr,
						&CreateRootTfSourceVal{
							Base: inputName,
							Key:  "",
						})))
					g.Add(jen.If(jen.Len(jen.Id(DiagId)).Op(">").Lit(0)).Block(
						jen.Return(jen.Id(DiagId)),
					))
					g.Add(jen.List(jen.Id("res"), jen.Id("err")).Op(":=").Id(facilId).Dot("Post").Call(jen.Id("ctx"), jen.Lit(path), jen.Id(paramsId)))
					g.Add(CondAddDiagErr(true, emptyPathExpr.Use(), fmt.Sprintf("failed to create new %s", objName)))
					g.Add(jen.Id(inputName).Dot("SetId").Call(jen.Qual(SharedPkg, "Dig").Index(jen.String()).Call(jen.Id("res"), jen.Lit(idField))))
					g.Add(jen.Return(jen.Id(DiagId)))
				}),

				jen.Id("ReadContext"): jen.Func().Params(
					jen.Id("ctx").Qual("context", "Context"),
					jen.Id(inputName).Op("*").Qual(SchemaPkg, "ResourceData"),
					jen.Id("f0").Any(),
				).Qual(DiagPkg, "Diagnostics").BlockFunc(func(g *jen.Group) {
					g.Add(getClientStmt)
					g.Add(createDiagStmt)
					g.Add(jen.List(jen.Id("res"), jen.Id("err")).Op(":=").Id(facilId).Dot("Get").Call(jen.Id("ctx"), idUrlPathExpr))
					g.Add(CondAddDiagErr(true, emptyPathExpr.Use(), fmt.Sprintf("failed to look up %s %%s", objName), idExpr))
					for _, statement := range fields.readApiInner(jen.Id("res"), &RootTfDestColl{Base: inputName, TfPrefix: ""}) {
						g.Add(statement)
					}
					g.Add(jen.Return(jen.Id(DiagId)))
				}),

				jen.Id("UpdateContext"): jen.Func().Params(
					jen.Id("ctx").Qual("context", "Context"),
					jen.Id(inputName).Op("*").Qual(SchemaPkg, "ResourceData"),
					jen.Id("f0").Any(),
				).Qual(DiagPkg, "Diagnostics").BlockFunc(func(g *jen.Group) {
					paramsId := "params"
					g.Add(getClientStmt)
					g.Add(createDiagStmt)
					g.Add(jen.Id(paramsId).Op(":=").Add(fields.ValidateSetApi(
						true,
						emptyPathExpr,
						&UpdateRootTfSourceVal{
							Base: inputName,
							Key:  "",
						},
					)))
					g.Add(jen.If(jen.Len(jen.Id(DiagId)).Op(">").Lit(0)).Block(
						jen.Return(jen.Id(DiagId)),
					))
					g.Add(jen.List(jen.Id("_"), jen.Id("err")).Op(":=").Id(facilId).Dot("Post").Call(
						jen.Id("ctx"),
						idUrlPathExpr,
						jen.Id(paramsId),
					))
					g.Add(CondAddDiagErr(true, emptyPathExpr.Use(), fmt.Sprintf("failed to update %s %%s", objName), idExpr))
					g.Add(jen.Return(jen.Id(DiagId)))
				}),

				jen.Id("DeleteContext"): jen.Func().Params(
					jen.Id("ctx").Qual("context", "Context"),
					jen.Id(inputName).Op("*").Qual(SchemaPkg, "ResourceData"),
					jen.Id("f0").Any(),
				).Qual(DiagPkg, "Diagnostics").BlockFunc(func(g *jen.Group) {
					g.Add(getClientStmt)
					g.Add(createDiagStmt)
					if hasDelete {
						g.Add(jen.Id("err").Op(":=").Id(facilId).Dot("Delete").Call(jen.Id("ctx"), idUrlPathExpr))
						g.Add(CondAddDiagErr(true, emptyPathExpr.Use(), fmt.Sprintf("failed to delete %s %%s", objName), idExpr))
					} else if hasActive {
						g.Add(jen.List(jen.Id("_"), jen.Id("err")).Op(":=").Id(facilId).Dot("Post").Call(
							jen.Id("ctx"),
							idUrlPathExpr,
							jen.Map(jen.String()).Any().Values(jen.Dict{
								jen.Lit("active"): jen.False(),
							}),
						))
						g.Add(CondAddDiagErr(true, emptyPathExpr.Use(), fmt.Sprintf("failed to set active to false on %s %%s", objName), idExpr))
					} else {
						panic("ASSERTION")
					}
					g.Add(jen.Return(jen.Id(DiagId)))
				}),
			})),
		)
		jenResources[jen.Lit("stripe_"+collName)] = jen.Id(funcName).Call()
		err = jenResFile.Save(fmt.Sprintf("generated/resources_%s.go", collName))
		if err != nil {
			panic(err)
		}
	}

	// Wrap and output
	{
		jenFile := jen.NewFile("main")
		jenFile.Func().Id("resources").Params().Map(jen.String()).Op("*").Qual(SchemaPkg, "Resource").Block(
			jen.Return(jen.Map(jen.String()).Op("*").Qual(SchemaPkg, "Resource").Values(jenResources)),
		)
		gitHash, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
		if err != nil {
			panic(err)
		}
		jenFile.Func().Id("gitHash").Params().String().Block(
			jen.Return(jen.Lit(strings.TrimSpace(string(gitHash)))),
		)
		err = jenFile.Save("generated/resources.go")
		if err != nil {
			panic(err)
		}
	}

	log.Printf("Done.")
}

func FakeScope(t jen.Code, body []jen.Code) jen.Code {
	return jen.Func().Params().Add(t).Block(body...).Call()
}
