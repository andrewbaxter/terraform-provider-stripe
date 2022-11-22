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
		outFields[jen.Lit(field.Name)] = jen.Op("&").Qual(SchemaPkg, "Schema").Values(fieldOut)
	}
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeMap"),
		jen.Id("Elem"): jen.Op("&").Qual(SchemaPkg, "Resource").Values(jen.Dict{
			jen.Id("Schema"): jen.Map(jen.String()).Op("*").Qual(SchemaPkg, "Schema").Values(outFields),
		}),
	}
}

func (n *NodeObj) Validate(path jen.Code, v jen.Code) jen.Code {
	fields := []jen.Code{}
	for _, field := range n.Fields {
		fields = append(fields, field.Spec.Validate(
			jen.Id("path").Dot("IndexString").Call(jen.Lit(field.Name)),
			jen.Qual(SharedPkg, "Dig").Index(jen.Any()).Call(jen.Id("outerV")),
		))
	}
	return jen.Block(append(
		[]jen.Code{
			jen.Id("path").Op(":=").Add(path),
			jen.Id("outerV").Op(":=").Add(v),
		},
		fields...,
	)...)
}
