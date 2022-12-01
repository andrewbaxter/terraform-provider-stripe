package main

import (
	"github.com/andrewbaxter/terraform-provider-stripe/shared"
	"github.com/dave/jennifer/jen"
)

type NodeMap struct {
	Elem Node
}

func (n *NodeMap) GenField() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeMap"),
		jen.Id("Elem"): n.Elem.GenElem(),
	}
}

func (n *NodeMap) GenElem() jen.Code {
	return jen.Op("&").Qual(SchemaPkg, "Schema").Values(n.GenField())
}

func (n *NodeMap) ReadApi(apiSource jen.Code, tfDest TfDestVal) []jen.Code {
	tfDestId := "dest"
	return []jen.Code{
		jen.Comment("NodeMap ReadApi"),
		jen.Id(tfDestId).Op(":=").Map(jen.String()).Any().Values(),
		// price lists currency_options in spec but not in reality
		jen.If(jen.List(jen.Id("preKv"), jen.Id("preKvOk")).Op(":=").
			Add(apiSource).Assert(jen.Map(jen.String()).Any())).Op(";").Id("preKvOk").Block(
			jen.For(jen.List(jen.Id("k"), jen.Id("v")).Op(":=").Range().Id("preKv")).Block(Flatten([][]jen.Code{
				{
					jen.Id("outerDest").Op(":=").Id(tfDestId),
				},
				n.Elem.ReadApi(jen.Id("v"), shared.Pointer(MapTfDestVal{
					Base: "outerDest",
					Key:  jen.Id("k"),
				})),
			})...),
		),
		tfDest.Set(jen.Id(tfDestId)),
	}
}

func (n *NodeMap) ValidateSetApi(update bool, tfPath *Usable[jen.Code], tfSource TfSourceVal) jen.Code {
	apiDestId := "dest"
	childPath := Unused(func() jen.Code { return jen.Id("path").Dot("IndexString").Call(jen.Id("k")) })
	validate := n.Elem.ValidateSetApi(update, childPath, shared.Pointer(AnyTfSourceVal("v")))
	pre := []jen.Code{
		jen.Comment("NodeMap ValidateSetApi"),
	}
	if childPath.WasUsed() {
		pre = append(
			pre,
			jen.Id("path").Op(":=").Add(tfPath.Use()),
		)
	}
	tfSourceId := "outerSource"
	return FakeScope(
		jen.Any(),
		Flatten([][]jen.Code{
			pre,
			{
				jen.Id(tfSourceId).Op(":=").Add(tfSource.Get()).Assert(jen.Map(jen.String()).Any()),
				jen.Id(apiDestId).Op(":=").Map(jen.String()).Any().Values(),
				jen.For(jen.List(jen.Id("k"), jen.Id("v")).Op(":=").Range().Id(tfSourceId)).Block(
					jen.Id(apiDestId).Index(jen.Id("k")).Op("=").Add(validate),
				),
				jen.Return(jen.Id(apiDestId)),
				jen.Comment("NodeMap ValidateSetApi END"),
			},
		}),
	)
}

func (n *NodeMap) IsNotDefault(id jen.Code) jen.Code {
	return jen.Len(jen.Add(id).Assert(jen.Map(jen.String()).Any())).Op(">").Lit(0)
}

func (n *NodeMap) CanRead() bool {
	return n.Elem.CanRead()
}
