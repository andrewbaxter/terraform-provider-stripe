package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/dave/jennifer/jen"
)

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func Last[T any](x []T) T {
	return x[len(x)-1]
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

func Default[T any]() T {
	var out T
	return out
}

func Dig[T any](bla any, field string) T {
	v, found := bla.(map[string]any)[field]
	if !found {
		return Default[T]()
	}
	v1, isT := v.(T)
	if !isT {
		panic("dug value not convertible")
	}
	return v1
}

func Reunmarshal[T any](bla any) T {
	var out T
	err := json.Unmarshal([]byte(Must(json.Marshal(bla))), &out)
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

var LineRegexp = regexp.MustCompile(`(?m)^.*$`)

func GenSchema(type_ string, spec any) jen.Dict {
	if type_ == "" {
		anyOf := Dig[[]any](spec, "anyOf")
		if anyOf != nil {
			filtered := []any{}
			allObj := true
			for _, sub := range anyOf {
				type_ := Dig[string](sub, "type")
				switch type_ {
				case "integer":
					return GenSchema(type_, sub)
				case "number":
					return GenSchema(type_, sub)
				case "boolean":
					return GenSchema(type_, sub)
				case "string":
					enum := Dig[[]any](sub, "enum")
					if enum != nil {
						if len(enum) == 1 && enum[0].(string) == "" {
							continue
						}
					} else {
						return GenSchema(type_, sub)
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
				return GenSchema(Dig[string](filtered[0], "type"), filtered[0])
			}
			if allObj {
				// Can't handle unions, so do the poor workaround of merging everything
				// into one object with all optional fields... can potentially verify later though
				desc := strings.Builder{}
				desc.WriteString("One of the following:\n\n")
				props := map[string]any{}
				for _, elAny := range filtered {
					el := Reunmarshal[OpenApiObj](elAny)
					desc.WriteString(fmt.Sprintf("* %s", LineRegexp.ReplaceAllString(el.Description, "   $0")))
					desc.WriteString(fmt.Sprintf("   Required:\n\n"))
					for key, prop := range el.Properties {
						desc.WriteString(fmt.Sprintf("   * `%s`\n", key))
						props[key] = prop
					}
				}
				return GenSchema("object", Reunmarshal[any](OpenApiObj{
					Description: desc.String(),
					Properties:  props,
					Required:    nil,
				}))
			}
			panic("TODO")
		}
	}
	switch type_ {
	case "string":
		return GenStringSchema()
	case "boolean":
		return GenBoolSchema()
	case "integer":
		return GenIntSchema()
	case "number":
		return GenFloatSchema()
	case "object":
		return GenObjSchema(spec)
	case "array":
		return GenArraySchema(spec)
	default:
		panic(fmt.Sprintf("Unhandled schema param type [[%s]]", type_))
	}
}

func GenStringSchema() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeString"),
	}
}

func GenBoolSchema() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeBool"),
	}
}

func GenIntSchema() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeInt"),
	}
}

func GenFloatSchema() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeFloat"),
	}
}

func GenArraySchema(spec any) jen.Dict {
	typed := Reunmarshal[OpenApiArray](spec)
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeList"),
		jen.Id("Elem"): jen.Op("&").Qual(SchemaPkg, "Schema").Values(GenSchema(Dig[string](typed.Items, "type"), typed.Items)),
	}
}

func GenObjParamSchemas(updatable map[string]bool, spec any) *jen.Statement {
	typed := Reunmarshal[OpenApiObj](spec)
	required := map[string]bool{}
	for _, paramName := range typed.Required {
		required[paramName] = true
	}
	return jen.Map(jen.String()).Op("*").Qual(SchemaPkg, "Schema").Values(func() jen.Dict {
		fields := jen.Dict{}
		for paramName, param := range typed.Properties {
			typedParam := Reunmarshal[OpenApiObjField](param)
			field := GenSchema(typedParam.Type, param)
			field[jen.Id("Description")] = jen.Lit(typedParam.Description)
			field[jen.Id("Required")] = func() *jen.Statement {
				if required[paramName] {
					return jen.True()
				} else {
					return jen.False()
				}
			}()
			field[jen.Id("ForceNew")] = func() *jen.Statement {
				if updatable[paramName] {
					return jen.False()
				} else {
					return jen.True()
				}
			}()
			fields[jen.Lit(paramName)] = jen.Op("&").Qual(SchemaPkg, "Schema").Values(field)
		}
		return fields
	}())
}

func GenObjSchema(spec any) jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeMap"),
		jen.Id("Elem"): jen.Op("&").Qual(SchemaPkg, "Resource").Values(jen.Dict{
			jen.Id("Schema"): GenObjParamSchemas(nil, spec),
		}),
	}
}

func OpenApiTypeToJen(t string) *jen.Statement {
	switch t {
	case "string":
		return jen.String()
	case "boolean":
		return jen.Bool()
	case "integer":
		return jen.Int64()
	case "object":
		return jen.Map(jen.String()).Any()
	case "array":
		return jen.Index().Any()
	default:
		return jen.Any()
		//panic(fmt.Sprintf("No known go type corresponding to openapi type [[%s]]", t))
	}
}

func Reraise(t string, args ...jen.Code) *jen.Statement {
	args = append([]jen.Code{
		jen.Lit("%s: %w"),
		jen.Lit(t),
	}, args...)
	args = append(args, jen.Id("err"))
	return jen.If(jen.Id("err").Op("!=").Nil()).Block(
		jen.Return(jen.Qual("fmt", "Errorf").Call(args...)),
	)
}

func main() {
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
	err := json.Unmarshal(Must(ioutil.ReadFile("stripe_openapi/openapi/spec3.json")), &spec)
	if err != nil {
		panic(err)
	}
	settingToken := "token"
	jenResources := jen.Dict{}
	stripePkg := "github.com/stripe/stripe-go/v74"
	stripeClientPkg := stripePkg + "/client"
	diagPkg := "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	instanceMethods := map[string]map[string]OpenApiMethod{}
	for path, methods := range spec.Paths {
		parts := strings.Split(path, "/")
		if Last(parts)[0] == '{' {
			instanceMethods[strings.Join(parts[0:len(parts)-1], "/")] = methods
		}
	}
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
		postSchema := Reunmarshal[OpenApiObj](post.RequestBody.Content.FormUrlencoded.Schema)
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
		objName := Last(strings.Split(post.Responses.R200.Content.Json.Schema.Ref, "/"))
		collNameParts := []string{}
		for _, seg := range strings.Split(path, "/")[2:] {
			collNameParts = append(collNameParts, FromSnake(seg)...)
		}
		collName := ToSnake(collNameParts)
		camelCollName := ToCamel(collNameParts)
		clientId := "client"
		getClientStmt := jen.Id(clientId).Op(":=").Id("m").Assert(jen.Op("*").Qual(stripeClientPkg, "API"))
		inputName := "d"
		idStmt := jen.Id(inputName).Dot("Id").Call()
		required := map[string]bool{}
		for _, paramName := range postSchema.Required {
			required[paramName] = true
		}

		jenResources[jen.Lit(collName)] = jen.Op("&").Qual(SchemaPkg, "Resource").Values(jen.Dict{
			jen.Id("Schema"): GenObjParamSchemas(updatable, post.RequestBody.Content.FormUrlencoded.Schema),

			jen.Id("CreateWithoutTimeout"): jen.Func().Params(
				jen.Id("_").Qual("context", "Context"),
				jen.Id(inputName).Op("*").Qual(SchemaPkg, "ResourceData"),
				jen.Id("m").Any(),
			).Qual(diagPkg, "Diagnostics").BlockFunc(func(g *jen.Group) {
				paramsId := "params"
				g.Add(getClientStmt)
				g.Add(jen.Id(paramsId).Op(":=").Qual(stripePkg, fmt.Sprintf("%sParams", camelCollName)).Values(jen.Dict{}))
				for paramName, param := range postSchema.Properties {
					camelParamName := ToCamel(FromSnake(paramName))
					paramType := Dig[string](param, "type")
					if required[paramName] {
						g.Add(
							jen.Id(paramsId).Dot(camelParamName).Op("=").
								Id(inputName).Dot("Get").Call(jen.Lit(paramName)).Assert(OpenApiTypeToJen(paramType)),
						)
					} else {
						g.If(jen.List(
							jen.Id("val"),
							jen.Id("exists"),
						).Op(":=").Id(inputName).Dot("GetOk").Call(jen.Lit(paramName)).Op(";").Id("exists")).Block(
							jen.Id(paramsId).Dot(camelParamName).Op("=").Op("&").Id("val"),
						)
					}
				}
				g.Add(jen.List(jen.Id("res"), jen.Id("err")).Op(":=").Id(clientId).Dot(camelCollName).Dot("New").Call(jen.Id(paramsId)))
				g.Add(Reraise(fmt.Sprintf("failed to create new %s", objName)))
				g.Add(jen.Id(inputName).Dot("SetId").Call(jen.Id("res").Dot("ID")))
				g.Add(jen.Return(jen.Nil()))
			}),

			jen.Id("ReadWithoutTimeout"): jen.Func().Params(
				jen.Id("_").Qual("context", "Context"),
				jen.Id(inputName).Op("*").Qual(SchemaPkg, "ResourceData"),
				jen.Id("m").Any(),
			).Qual(diagPkg, "Diagnostics").BlockFunc(func(g *jen.Group) {
				g.Add(getClientStmt)
				g.Add(jen.List(jen.Id("res"), jen.Id("err")).Op(":=").Id(clientId).Dot(camelCollName).Dot("Get").Call(idStmt))
				g.Add(Reraise(fmt.Sprintf("failed to look up %s %%s", objName), idStmt))
				for paramName := range postSchema.Properties {
					camelParamName := ToCamel(FromSnake(paramName))
					g.Add(jen.Id(inputName).Dot("Set").Call(
						jen.Lit(paramName),
						jen.Id("res").Dot(camelParamName),
					))
				}
				g.Add(jen.Return(jen.Nil()))
			}),

			jen.Id("UpdateWithoutTimeout"): jen.Func().Params(
				jen.Id("_").Qual("context", "Context"),
				jen.Id(inputName).Op("*").Qual(SchemaPkg, "ResourceData"),
				jen.Id("m").Any(),
			).Qual(diagPkg, "Diagnostics").BlockFunc(func(g *jen.Group) {
				paramsId := "params"
				g.Add(getClientStmt)
				g.Add(jen.Id(paramsId).Op(":=").Qual(stripePkg, fmt.Sprintf("%sParams", camelCollName)).Values(jen.Dict{}))
				for paramName, param := range postSchema.Properties {
					if !updatable[paramName] {
						continue
					}
					camelParamName := ToCamel(FromSnake(paramName))
					paramType := Dig[string](param, "type")
					g.If(jen.Id(inputName).Dot("HasChange").Call(jen.Lit(paramName))).Block(
						jen.Id(paramsId).Dot(camelParamName).Op("=").
							Op("&").Id(inputName).Dot("Get").Call(jen.Lit(paramName)).Assert(OpenApiTypeToJen(paramType)),
					)
				}
				g.Add(jen.List(jen.Id("_"), jen.Id("err")).Op(":=").Id(clientId).Dot(camelCollName).Dot("Update").Call(
					idStmt,
					jen.Id(paramsId),
				))
				g.Add(Reraise(fmt.Sprintf("failed to update %s %%s", objName), idStmt))
				g.Add(jen.Return(jen.Nil()))
			}),

			jen.Id("DeleteWithoutTimeout"): jen.Func().Params(
				jen.Id("_").Qual("context", "Context"),
				jen.Id(inputName).Op("*").Qual(SchemaPkg, "ResourceData"),
				jen.Id("m").Any(),
			).Qual(diagPkg, "Diagnostics").BlockFunc(func(g *jen.Group) {
				g.Add(getClientStmt)
				g.Add(jen.Id("err").Op(":=").Id(clientId).Dot(camelCollName).Dot("Delete").Call(idStmt))
				g.Add(Reraise(fmt.Sprintf("failed to delete %s %%s", objName), idStmt))
				g.Add(jen.Return(jen.Nil()))
			}),
		})
	}
	jenFile := jen.NewFile("main")
	pluginPkg := "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	jenFile.Func().Id("main").Params().Block(
		jen.Qual(pluginPkg, "Serve").Call(
			jen.Op("&").Qual(pluginPkg, "ServeOpts").Values(jen.Dict{
				jen.Id("ProviderFunc"): jen.Func().
					Params().
					Op("*").Qual(SchemaPkg, "Provider").
					Block(
						jen.Return(jen.Op("&").Qual(SchemaPkg, "Provider").Values(jen.Dict{
							jen.Id("Schema"): jen.Map(jen.String()).Op("*").Qual(SchemaPkg, "Schema").Values(jen.Dict{
								jen.Lit(settingToken): jen.Values(jen.Dict{
									jen.Id("Type"): jen.Qual(SchemaPkg, "TypeString"),
								}),
							}),
							jen.Id("ConfigureFunc"): jen.Func().Params(
								jen.Id("d").Op("*").Qual(SchemaPkg, "ResourceData"),
							).Parens(jen.List(jen.Any(), jen.Error())).Block(
								jen.Qual(stripePkg, "SetAppInfo").Call(
									jen.Op("&").Qual(stripePkg, "AppInfo").Values(jen.Dict{
										jen.Id("Name"): jen.Lit("github.com/andrewbaxter/terraform-provider-stripe"),
									}),
								),
								jen.Id("client").Op(":=").Op("&").Qual(stripeClientPkg, "API").Values(jen.Dict{}),
								jen.Id("client").Dot("Init").Call(
									jen.Id("d").Dot("Get").Call(jen.Lit(settingToken)).Assert(jen.String()),
									jen.Nil(),
								),
								jen.Return(
									jen.Id("client"),
									jen.Nil(),
								),
							),
							jen.Id("ResourcesMap"): jen.Map(jen.String()).Op("*").Qual(SchemaPkg, "Resource").Values(jenResources),
						})),
					),
			})),
	)
	err = jenFile.Save("main.go")
	if err != nil {
		panic(err)
	}
	log.Printf("Done.")
}
