package main

import "github.com/dave/jennifer/jen"

type NodeObjField struct {
	Description string
	Name        string
	Required    bool
	Spec        Node
}

type NodeObj struct {
	Fields []NodeObjField
}

func (n *NodeObj) GenSchema() jen.Dict {
	outFields := jen.Dict{}
	for _, field := range n.Fields {
		fieldOut := field.Spec.GenSchema()
		fieldOut[jen.Id("Description")] = jen.Lit(field.Description)
		fieldOut[jen.Id("Required")] = func() *jen.Statement {
			if field.Required {
				return jen.True()
			} else {
				return jen.False()
			}
		}()
		outFields[jen.Lit(field.Name)] = jen.Values(fieldOut)
	}
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeMap"),
		jen.Id("Elem"): jen.Op("&").Qual(SchemaPkg, "Resource").Values(jen.Dict{
			jen.Id("Schema"): jen.Map(jen.String()).Op("*").Qual(SchemaPkg, "Schema").Values(outFields),
		}),
	}
}

func (n *NodeObj) Validate(path jen.Code, v jen.Code) *jen.Code {
	fields := []jen.Code{}
	for _, field := range n.Fields {
		validate := field.Spec.Validate(jen.Id("path"), jen.Id("v"))

		if validate == nil && !field.Required {
			continue
		}

		yesStatements := []jen.Code{}
		if validate != nil {
			yesStatements = append(yesStatements, *validate)
		}
		if_ := jen.If(jen.Id("v").Op("!=").Nil()).Block(yesStatements...)
		if field.Required {
			if_.Else().Block(
				jen.Id("out").Op("=").Id("append").Call(jen.Id("out"), jen.Qual(DiagPkg, "Diagnostic").Values(jen.Dict{
					jen.Id("Severity"):      jen.Qual(DiagPkg, "Error"),
					jen.Id("Summary"):       jen.Lit("Field is missing but required."),
					jen.Id("AttributePath"): jen.Id("path"),
				})),
			)
		}
		fields = append(fields, jen.Block(
			jen.Id("path").Op(":=").Id("path").Dot("IndexString").Call(jen.Lit(field.Name)),
			jen.Id("v").Op(":=").Qual(SharedPkg, "Dig").Index(jen.Any()).Call(jen.Id("outerV"), jen.Lit(field.Name)),
			if_,
		))
	}
	if len(fields) == 0 {
		return nil
	}
	var out jen.Code = jen.Block(append(
		[]jen.Code{
			jen.Id("path").Op(":=").Add(path),
			jen.Id("outerV").Op(":=").Add(v),
		},
		fields...,
	)...)
	return &out
}
