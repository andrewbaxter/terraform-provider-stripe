package main

import (
	"github.com/andrewbaxter/terraform-provider-stripe/shared"
	"github.com/dave/jennifer/jen"
)

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
	return []jen.Code{jen.If(jen.List(jen.Id("outerSource"), jen.Id("outerSourceOk").Op(":=")).
		Add(apiSource).Assert(jen.Map(jen.String()).Any()).
		Op(";").Id("outerSourceOk")).Block(
		jen.Comment("NodeFakeMap ReadApi"),
		jen.Id(tfDestId).Op(":=").Index().Any().Values(),
		jen.For(jen.List(jen.Id("k"), jen.Id("v")).Op(":=").Range().Id("outerSource")).Block(Flatten([][]jen.Code{
			{
				jen.Id("outerDest").Op(":=").Op("&").Id(tfDestId),
			},
			n.Elem.ReadApi(jen.Id("v"), shared.Pointer(ArrayTfDestVal("outerDest"))),
		})...),
		tfDest.Set(jen.Id(tfDestId)),
		jen.Comment("NodeFakeMap ReadApi END"),
	)}
}

func (n *NodeFakeMap) ValidateSetApi(update bool, tfPath *Usable[jen.Code], tfSource TfSourceVal) jen.Code {
	tfSourceId := "outerSource"
	apiDestId := "dest"
	childPath := Unused(func() jen.Code { return jen.Add(tfPath.Use()).Dot("IndexInt").Call(jen.Id("i")) })
	validate := n.Elem.ValidateSetApi(update, childPath, shared.Pointer(AnyTfSourceVal("v")))
	pre := []jen.Code{
		jen.Comment("NodeFakeMap ValidateSetApi"),
	}
	var index jen.Code
	if childPath.WasUsed() {
		index = jen.Id("i")
	} else {
		index = jen.Id("_")
	}
	return FakeScope(
		Flatten([][]jen.Code{
			pre,
			{
				jen.Id(tfSourceId).Op(":=").Add(tfSource.Get()).Assert(jen.Index().Any()),
				jen.Id(apiDestId).Op(":=").Map(jen.String()).Any().Values(),
				jen.For(jen.List(index, jen.Id("v")).Op(":=").Range().Id(tfSourceId)).Block(
					jen.List(jen.Id("subRes"), jen.Id("subResOk")).Op(":=").Add(validate),
					jen.If(jen.Id("subResOk")).Block(
						jen.Id("key").Op(":=").Qual(SharedPkg, "DigDelete").Index(jen.String()).Call(jen.Id("subRes"), jen.Lit("key")),
						jen.Id(apiDestId).Index(jen.Id("key")).Op("=").Id("subRes"),
					),
				),
				jen.Return(jen.Id(apiDestId), jen.Lit(true)),
				jen.Comment("NodeFakeMap ValidateSetApi END"),
			},
		}),
	)
}

func (n *NodeFakeMap) IsNotDefault(id jen.Code) jen.Code {
	return jen.Len(jen.Add(id).Assert(jen.Index().Any())).Op(">").Lit(0)
}

func (n *NodeFakeMap) CanRead() bool {
	return n.Elem.CanRead()
}
