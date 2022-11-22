package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"github.com/andrewbaxter/terraform-provider-stripe/shared"
	"github.com/dave/jennifer/jen"
)

func Last[T any](x []T) T {
	return x[len(x)-1]
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

func P[T any](x T) *T {
	return &x
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

type OpenApiPrim struct {
	Description string `json:"description"`
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

type OpenApiObjField struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

type OpenApiObj struct {
	Description string         `json:"description"`
	Properties  map[string]any `json:"properties"`
	Required    []string       `json:"required"`
}

const SchemaPkg = "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
const DiagPkg = "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
const CtyPkg = "github.com/hashicorp/go-cty/cty"
const SharedPkg = "github.com/andrewbaxter/terraform-provider-stripe/shared"

func AddDiagErr(exit bool, path jen.Code, t string, messageArgs ...jen.Code) *jen.Statement {
	messageArgs = append([]jen.Code{
		jen.Lit(t + ": %s"),
	}, messageArgs...)
	messageArgs = append(messageArgs, jen.Id("err"))
	statements := []jen.Code{
		jen.Id("out").Op("=").Id("append").Call(jen.Id("out"), jen.Qual(DiagPkg, "Diagnostic").Values(jen.Dict{
			jen.Id("Severity"):      jen.Qual(DiagPkg, "Error"),
			jen.Id("Summary"):       jen.Qual("fmt", "Sprintf").Call(messageArgs...),
			jen.Id("AttributePath"): path,
		})),
	}
	if exit {
		statements = append(statements, jen.Return(jen.Id("out")))
	}
	return jen.If(jen.Id("err").Op("!=").Nil()).Block(statements...)
}

type Node interface {
	GenSchema() jen.Dict
	Validate(path jen.Code, v jen.Code) *jen.Code
}

func BuildNode(spec any) Node {
	type_ := shared.Dig[string](spec, "type")
	if type_ == "" {
		anyOf := shared.Dig[[]any](spec, "anyOf")
		if anyOf != nil {
			filtered := []any{}
			allObj := true
			for _, sub := range anyOf {
				type_ := shared.Dig[string](sub, "type")
				switch type_ {
				case "integer", "number", "boolean":
					return BuildNode(sub)
				case "string":
					enum := shared.Dig[[]any](sub, "enum")
					if enum != nil {
						if len(enum) == 1 && enum[0].(string) == "" {
							continue
						}
					} else {
						return BuildNode(sub)
					}
					allObj = false
				case "object":
					// nop
				case "array":
					// nop
				default:
					panic("TODO")
				}
				filtered = append(filtered, sub)
			}
			if len(filtered) == 1 {
				return BuildNode(filtered[0])
			}
			if allObj {
				return &NodeAnyOfObjs{
					Options: Map(filtered, func(o any) *NodeObj {
						return BuildNode(o).(*NodeObj)
					}),
				}
			}
			panic("TODO")
		}
	}
	switch type_ {
	case "string":
		spec := Reunmarshal[OpenApiString](spec)
		enum := spec.Enum
		return &NodeString{
			Enum: enum,
		}
	case "boolean":
		return &NodeBool{}
	case "integer":
		return &NodeInt{}
	case "number":
		return &NodeFloat{}
	case "object":
		spec := Reunmarshal[OpenApiObj](spec)
		required := map[string]bool{}
		for _, paramName := range spec.Required {
			required[paramName] = true
		}
		return &NodeObj{
			Fields: MapKVs(spec.Properties, func(propName string, prop any) NodeObjField {
				return NodeObjField{
					Description: shared.Dig[string](prop, "description"),
					Name:        propName,
					Required:    required[propName],
					Spec:        BuildNode(prop),
				}
			}),
		}
	case "array":
		spec := Reunmarshal[OpenApiArray](spec)
		return &NodeArray{
			Elem: BuildNode(spec.Items),
		}
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
			log.Printf("Skipping %s, parameterized = no collection", path)
			continue
		}
		post, hasPost := methods["post"]
		if !hasPost {
			log.Printf("Skipping %s, no post method", path)
			continue
		}
		objName := Last(strings.Split(post.Responses.R200.Content.Json.Schema.Ref, "/"))
		idField := ""
		for propName := range spec.Components.Schemas[objName].Properties {
			if propName == "id" {
				idField = propName
				break
			}
		}
		if idField == "" {
			log.Printf("Skipping %s, no id field", path)
			continue
		}

		updatable := map[string]bool{}
		{
			update, hasUpdate := instanceMethods[path]["post"]
			if hasUpdate {
				updateSchema := Reunmarshal[OpenApiObj](update.RequestBody.Content.FormUrlencoded.Schema)
				for paramName := range updateSchema.Properties {
					updatable[paramName] = true
				}
			}
		}
		_, hasDelete := instanceMethods[path]["delete"]

		if post.RequestBody.Content.FormUrlencoded.Schema == nil {
			log.Printf("Skipping %s, no post schema", path)
			continue
		}
		fields := BuildNode(post.RequestBody.Content.FormUrlencoded.Schema).(*NodeObj)

		hasActive := false
		for _, f := range fields.Fields {
			if f.Name == "active" {
				hasActive = true
				break
			}
		}

		if !hasDelete && !hasActive {
			log.Printf("Skipping %s, no active and no delete", path)
			continue
		}

		fields.Fields = Filter(fields.Fields, func(v NodeObjField) bool {
			switch v.Name {
			case "expand", "metadata", "active":
				return false
			default:
				return true
			}
		})

		collNameParts := []string{}
		for _, seg := range strings.Split(path, "/")[2:] {
			collNameParts = append(collNameParts, FromSnake(seg)...)
		}
		collName := ToSnake(collNameParts)
		facilId := "f"
		getClientStmt := jen.Id(facilId).Op(":=").Id("f0").Assert(jen.Op("*").Id("Facilitator"))
		createDiagStmt := jen.Id("out").Op(":=").Qual(DiagPkg, "Diagnostics").Values(jen.Dict{})
		inputName := "d"
		idExpr := jen.Id(inputName).Dot("Id").Call()
		idUrlPathExpr := jen.Qual("fmt", "Sprintf").Call(jen.Lit(path+"/%v"), jen.Id(inputName).Dot("Get").Call(jen.Lit(idField)))
		emptyPathExpr := jen.Qual(CtyPkg, "Path").Values(jen.Dict{})

		topSchemaFields := jen.Dict{}
		for _, field := range fields.Fields {
			fieldOut := field.Spec.GenSchema()
			fieldOut[jen.Id("Description")] = jen.Lit(field.Description)
			if field.Required {
				fieldOut[jen.Id("Required")] = jen.True()
			} else {
				fieldOut[jen.Id("Optional")] = jen.True()
			}
			if !updatable[field.Name] {
				fieldOut[jen.Id("ForceNew")] = jen.True()
			}
			topSchemaFields[jen.Lit(field.Name)] = jen.Values(fieldOut)
		}

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
					g.Add(jen.Id(paramsId).Op(":=").Map(jen.String()).Any().Values(jen.Dict{}))
					for _, field := range fields.Fields {
						if field.Required {
							statements := []jen.Code{
								jen.Id("v").Op(":=").Id(inputName).Dot("Get").Call(jen.Lit(field.Name)),
							}
							validate := field.Spec.Validate(emptyPathExpr, jen.Id("v"))
							if validate != nil {
								statements = append(statements, *validate)
							}
							statements = append(statements, jen.Id(paramsId).Index(jen.Lit(field.Name)).Op("=").Id("v"))
							g.Block(statements...)
						} else {
							statements := []jen.Code{}
							validate := field.Spec.Validate(emptyPathExpr, jen.Id("v"))
							if validate != nil {
								statements = append(statements, *validate)
							}
							statements = append(statements, jen.Id(paramsId).Index(jen.Lit(field.Name)).Op("=").Id("v"))
							g.If(jen.List(
								jen.Id("v"),
								jen.Id("exists"),
							).Op(":=").Id(inputName).Dot("GetOk").Call(jen.Lit(field.Name)).Op(";").Id("exists")).Block(statements...)
						}
					}
					g.Add(jen.List(jen.Id("res"), jen.Id("err")).Op(":=").Id(facilId).Dot("Post").Call(jen.Id("ctx"), jen.Lit(path), jen.Id(paramsId)))
					g.Add(AddDiagErr(true, emptyPathExpr, fmt.Sprintf("failed to create new %s", objName)))
					g.Add(jen.Id(inputName).Dot("SetId").Call(jen.Qual(SharedPkg, "Dig").Index(jen.String()).Call(jen.Id("res"), jen.Lit(idField))))
					g.Add(jen.Return(jen.Id("out")))
				}),

				jen.Id("ReadContext"): jen.Func().Params(
					jen.Id("ctx").Qual("context", "Context"),
					jen.Id(inputName).Op("*").Qual(SchemaPkg, "ResourceData"),
					jen.Id("f0").Any(),
				).Qual(DiagPkg, "Diagnostics").BlockFunc(func(g *jen.Group) {
					g.Add(getClientStmt)
					g.Add(createDiagStmt)
					g.Add(jen.List(jen.Id("res"), jen.Id("err")).Op(":=").Id(facilId).Dot("Get").Call(jen.Id("ctx"), idUrlPathExpr))
					g.Add(AddDiagErr(true, emptyPathExpr, fmt.Sprintf("failed to look up %s %%s", objName), idExpr))
					for _, field := range fields.Fields {
						g.Add(jen.Id(inputName).Dot("Set").Call(
							jen.Lit(field.Name),
							jen.Qual(SharedPkg, "Dig").Index(jen.Any()).Call(jen.Id("res"), jen.Lit(field.Name)),
						))
					}
					g.Add(jen.Return(jen.Id("out")))
				}),

				jen.Id("UpdateContext"): jen.Func().Params(
					jen.Id("ctx").Qual("context", "Context"),
					jen.Id(inputName).Op("*").Qual(SchemaPkg, "ResourceData"),
					jen.Id("f0").Any(),
				).Qual(DiagPkg, "Diagnostics").BlockFunc(func(g *jen.Group) {
					paramsId := "params"
					g.Add(getClientStmt)
					g.Add(createDiagStmt)
					g.Add(jen.Id(paramsId).Op(":=").Map(jen.String()).Any().Values(jen.Dict{}))
					for _, field := range fields.Fields {
						if !updatable[field.Name] {
							continue
						}
						statements := []jen.Code{
							jen.Id("v").Op(":=").Id(inputName).Dot("Get").Call(jen.Lit(field.Name)),
						}
						validate := field.Spec.Validate(emptyPathExpr, jen.Id("v"))
						if validate != nil {
							statements = append(statements, *validate)
						}
						statements = append(statements, jen.Id(paramsId).Index(jen.Lit(field.Name)).Op("=").Id("v"))
						g.If(jen.Id(inputName).Dot("HasChange").Call(jen.Lit(field.Name))).Block(statements...)
					}
					g.Add(jen.List(jen.Id("_"), jen.Id("err")).Op(":=").Id(facilId).Dot("Post").Call(
						jen.Id("ctx"),
						idUrlPathExpr,
						jen.Id(paramsId),
					))
					g.Add(AddDiagErr(true, emptyPathExpr, fmt.Sprintf("failed to update %s %%s", objName), idExpr))
					g.Add(jen.Return(jen.Id("out")))
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
						g.Add(AddDiagErr(true, emptyPathExpr, fmt.Sprintf("failed to delete %s %%s", objName), idExpr))
					} else if hasActive {
						g.Add(jen.List(jen.Id("_"), jen.Id("err")).Op(":=").Id(facilId).Dot("Post").Call(
							jen.Id("ctx"),
							idUrlPathExpr,
							jen.Map(jen.String()).Any().Values(jen.Dict{
								jen.Lit("active"): jen.False(),
							}),
						))
						g.Add(AddDiagErr(true, emptyPathExpr, fmt.Sprintf("failed to set active to false on %s %%s", objName), idExpr))
					} else {
						panic("ASSERTION")
					}
					g.Add(jen.Return(jen.Id("out")))
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
