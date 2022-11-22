package main

import "github.com/dave/jennifer/jen"

type NodeArray struct {
	Elem Node
}

func (n *NodeArray) GenSchema() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeList"),
		jen.Id("Elem"): jen.Op("&").Qual(SchemaPkg, "Schema").Values(n.Elem.GenSchema()),
	}
}

func (n *NodeArray) Validate(path jen.Code, v jen.Code) *jen.Code {
	validate := n.Elem.Validate(
		jen.Id("path").Dot("IndexInt").Call(jen.Id("i")),
		jen.Id("v"),
	)
	if validate == nil {
		return nil
	}
	var out jen.Code = jen.Block( // compiler workaround...
		jen.Id("path").Op(":=").Add(path),
		jen.Id("outerV").Op(":=").Add(v),
		jen.For(jen.List(jen.Id("i"), jen.Id("v")).Op(":=").Range().Id("outerV").Assert(jen.Index().Any())).Block(*validate),
	)
	return &out
}
