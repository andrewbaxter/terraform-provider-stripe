package main

import "github.com/dave/jennifer/jen"

type NodeArray struct {
	Elem Node
}

func (n *NodeArray) GenField() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeList"),
		jen.Id("Elem"): n.Elem.GenElem(),
	}
}

func (n *NodeArray) GenElem() jen.Code {
	return jen.Op("&").Qual(SchemaPkg, "Schema").Values(n.GenField())
}

func (n *NodeArray) ReadApi(apiSource jen.Code, tfDest TfDestVal) []jen.Code {
	tfDestId := "dest"
	return []jen.Code{jen.If(jen.List(jen.Id("outerSource"), jen.Id("outerSourceOk").Op(":=")).
		Add(apiSource).Assert(jen.Index().Any()).
		Op(";").Id("outerSourceOk")).Block(
		jen.Comment("NodeArray ReadApi"),
		jen.Id(tfDestId).Op(":=").Index().Any().Values(),
		jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Id("outerSource")).Block(Flatten([][]jen.Code{
			{jen.Id("outerDest").Op(":=").Id(tfDestId)},
			n.Elem.ReadApi(jen.Id("v"), P(ArrayTfDestVal("outerDest"))),
		})...),
		tfDest.Set(jen.Id(tfDestId)),
		jen.Comment("NodeArray ReadApi END"),
	)}
}

func (n *NodeArray) ValidateSetApi(update bool, tfPath *Usable[jen.Code], tfSource TfSourceVal) jen.Code {
	tfSourceId := "outerSource"
	apiDestId := "dest"
	childPath := Unused(func() jen.Code { return jen.Id("path").Dot("IndexInt").Call(jen.Id("i")) })
	validate := n.Elem.ValidateSetApi(false, childPath, P(AnyTfSourceVal("v")))
	pre := []jen.Code{
		jen.Comment("NodeArray ValidateSetApi"),
	}
	if childPath.WasUsed() {
		pre = append(
			pre,
			jen.Id("path").Op(":=").Add(tfPath.Use()),
		)
	}
	return FakeScope(
		jen.Any(),
		Flatten([][]jen.Code{
			pre,
			{
				jen.Id(tfSourceId).Op(":=").Add(tfSource.Get()).Assert(jen.Index().Any()),
				jen.Id(apiDestId).Op(":=").Make(jen.Index().Any(), jen.Len(jen.Id(tfSourceId))),
				jen.For(jen.List(jen.Id("i"), jen.Id("v")).Op(":=").Range().Id(tfSourceId)).Block(
					jen.Id(apiDestId).Index(jen.Id("i")).Op("=").Append(jen.Id(apiDestId), validate),
				),
				jen.Return(jen.Id(apiDestId)),
				jen.Comment("NodeArray ValidateSetApi END"),
			},
		}),
	)
}
