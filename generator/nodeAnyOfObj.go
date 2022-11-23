package main

import (
	"github.com/dave/jennifer/jen"
)

type NodeAnyOfObjs struct {
	Options []*NodeObj
}

func (n *NodeAnyOfObjs) synthObj() *NodeObj {
	fields := []*NodeObjField{}
	for _, o := range n.Options {
		for _, prop := range o.Fields {
			fields = append(fields, &NodeObjField{
				Description: prop.Description,
				ApiKey:      prop.ApiKey,
				Required:    false,
				Spec:        prop.Spec,
			})
		}
	}
	return &NodeObj{
		Fields: fields,
	}
}

func (n *NodeAnyOfObjs) GenField() jen.Dict {
	return n.synthObj().GenField()
}

func (n *NodeAnyOfObjs) GenElem() jen.Code {
	return n.synthObj().GenElem()
}

func (n *NodeAnyOfObjs) ReadApi(apiSource jen.Code, tfDest TfDestVal) []jen.Code {
	return n.synthObj().ReadApi(apiSource, tfDest)
}

func (n *NodeAnyOfObjs) ValidateSetApi(tfPath *Usable[jen.Code], tfSource jen.Code) jen.Code {
	tfSourceId := "outerSource"
	optionStatements := []jen.Code{}
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
			cond.Call(jen.Lit(f.ApiKey), jen.Id(tfSourceId))
		}
		optionStatements = append(optionStatements, jen.If(cond).Block(
			jen.Return(o.ValidateSetApi(Unused(func() jen.Code { return jen.Id("path") }), jen.Id(tfSourceId))),
		))
	}
	return FakeScope(
		jen.Any(),
		Flatten(
			[][]jen.Code{
				{
					jen.Id("path").Op(":=").Add(tfPath.Use()),
					jen.Id(tfSourceId).Op(":=").Add(tfSource).Assert(jen.Map(jen.String()).Any()),
				},
				optionStatements,
				{
					jen.Id("out").Op("=").Id("append").Call(jen.Id("out"), jen.Qual(DiagPkg, "Diagnostic").Values(jen.Dict{
						jen.Id("Severity"): jen.Qual(DiagPkg, "Error"),
						jen.Id("Summary"): jen.Qual("fmt", "Sprintf").
							Call(jen.Lit("This field is a union, but no set of required fields match any union member.")),
						jen.Id("AttributePath"): jen.Id("path"),
					})),
					jen.Return(jen.Nil()),
				},
			},
		),
	)
}
