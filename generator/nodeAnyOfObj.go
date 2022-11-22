package main

import (
	"github.com/dave/jennifer/jen"
)

type NodeAnyOfObjs struct {
	Options []*NodeObj
}

func (n *NodeAnyOfObjs) GenSchema() jen.Dict {
	fields := []NodeObjField{}
	for _, o := range n.Options {
		for _, prop := range o.Fields {
			fields = append(fields, NodeObjField{
				Description: prop.Description,
				Name:        prop.Name,
				Required:    false,
				Spec:        prop.Spec,
			})
		}
	}
	return (&NodeObj{
		Fields: fields,
	}).GenSchema()
}

func (n *NodeAnyOfObjs) Validate(path jen.Code, v jen.Code) jen.Code {
	statements := []jen.Code{
		jen.Id("path").Op(":=").Add(path),
		jen.Id("outerV").Op(":=").Add(v).Assert(jen.Map(jen.String()).Any()),
		jen.Id("oneOk").Op(":=").False(),
	}
	for _, o := range n.Options {
		var cond *jen.Statement = nil
		for _, f := range o.Fields {
			if !f.Required {
				continue
			}
			if cond == nil {
				cond = jen.Id("inMap")
			} else {
				cond.Op("&&").Id("inMap")
			}
			cond.Call(jen.Lit(f.Name), jen.Id("outerV"))
		}
		statements = append(statements, jen.If(cond).Block(
			jen.Id("oneOk").Op("=").True(),
			o.Validate(jen.Id("path"), jen.Id("outerV")),
		))
	}
	statements = append(statements, jen.If(jen.Op("!").Id("oneOk")).Block(
		jen.Id("out").Op("=").Id("append").Call(jen.Id("out"), jen.Qual(DiagPkg, "Diagnostic").Values(jen.Dict{
			jen.Id("Severity"): jen.Qual(DiagPkg, "Error"),
			jen.Id("Summary"): jen.Qual("fmt", "Sprintf").
				Call(jen.Lit("This field is a union, but no set of required fields match any union member.")),
			jen.Id("AttributePath"): jen.Id("path"),
		})),
	))
	// if not oneOk diag err
	return jen.Block(statements...)
}
