package main

import "github.com/dave/jennifer/jen"

type NodeFakeMap struct {
	Elem Node
}

func (n *NodeFakeMap) GenField() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeList"),
		jen.Id("Elem"): n.Elem.GenElem(),
	}
}

func (n *NodeFakeMap) GenElem() jen.Code {
	return jen.Op("&").Qual(SchemaPkg, "Schema").Values(n.GenField())
}

func (n *NodeFakeMap) ReadApi(apiSource jen.Code, tfDest TfDestVal) []jen.Code {
	tfDestId := "dest"
	return []jen.Code{
		jen.Comment("NodeFakeMap ReadApi"),
		jen.Id(tfDestId).Op(":=").Index().Any().Values(),
		jen.For(jen.List(jen.Id("k"), jen.Id("v")).Op(":=").Range().Add(apiSource).Assert(jen.Map(jen.String()).Any())).Block(Flatten([][]jen.Code{
			{
				jen.Id("outerDest").Op(":=").Id(tfDestId),
			},
			n.Elem.ReadApi(jen.Id("v"), P(ArrayTfDestVal("outerDest"))),
		})...),
		tfDest.Set(jen.Id(tfDestId)),
	}
}

func (n *NodeFakeMap) ValidateSetApi(tfPath *Usable[jen.Code], tfSource jen.Code) jen.Code {
	tfSourceId := "outerSource"
	apiDestId := "dest"
	childPath := Unused(func() jen.Code { return jen.Id("path").Dot("IndexInt").Call(jen.Id("i")) })
	validate := n.Elem.ValidateSetApi(
		childPath,
		jen.Id("v"),
	)
	pre := []jen.Code{
		jen.Comment("NodeFakeMap ValidateSetApi"),
	}
	if childPath.WasUsed() {
		pre = append(
			pre,
			jen.Id("path").Op(":=").Add(tfPath.Use()),
		)
	}
	var index jen.Code
	if childPath.WasUsed() {
		index = jen.Id("i")
	} else {
		index = jen.Id("_")
	}
	return FakeScope(
		jen.Any(),
		Flatten([][]jen.Code{
			pre,
			{
				jen.Id(tfSourceId).Op(":=").Add(tfSource).Assert(jen.Index().Any()),
				jen.Id(apiDestId).Op(":=").Map(jen.String()).Any().Values(),
				jen.For(jen.List(index, jen.Id("v")).Op(":=").Range().Id(tfSourceId)).Block(
					jen.Id("subRes").Op(":=").Add(validate),
					jen.Id("key").Op(":=").Qual(SharedPkg, "DigDelete").Index(jen.String()).Call(jen.Id("subRes"), jen.Lit("key")),
					jen.Id(apiDestId).Index(jen.Id("key")).Op("=").Id("subRes"),
				),
				jen.Return(jen.Id(apiDestId)),
			},
		}),
	)
}
