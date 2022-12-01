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
			behavior := prop.Behavior
			if behavior == NBUserRequired {
				behavior = NBUserOptional
			}
			fields = append(fields, &NodeObjField{
				Description: prop.Description,
				Key:         prop.Key,
				Behavior:    behavior,
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

func (n *NodeAnyOfObjs) ValidateSetApi(update bool, tfPath *Usable[jen.Code], tfSource TfSourceVal) jen.Code {
	optionStatements := []jen.Code{}
	for _, o := range n.Options {
		cond := o.buildAnyTfFieldsPresent(tfSource)
		if cond == nil {
			panic("ASSERTION") // no identifiable fields for variant
		}
		optionStatements = append(optionStatements, jen.If(cond).Block(
			jen.Return(o.ValidateSetApi(
				update,
				Unused(func() jen.Code { return jen.Id("path") }),
				tfSource,
			)),
		))
	}
	return FakeScope(
		jen.Any(),
		Flatten(
			[][]jen.Code{
				{
					jen.Id("path").Op(":=").Add(tfPath.Use()),
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

func (n *NodeAnyOfObjs) IsNotDefault(id jen.Code) jen.Code {
	return jen.True()
}
