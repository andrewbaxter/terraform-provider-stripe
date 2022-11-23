package main

import "github.com/dave/jennifer/jen"

type NodeFloat struct {
}

func (n *NodeFloat) GenField() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeFloat"),
	}
}

func (n *NodeFloat) GenElem() jen.Code {
	return jen.Op("&").Qual(SchemaPkg, "Schema").Values(n.GenField())
}

func (n *NodeFloat) ReadApi(apiSource jen.Code, tfDest TfDestVal) []jen.Code {
	return []jen.Code{tfDest.Set(apiSource)}
}

func (n *NodeFloat) ValidateSetApi(tfPath *Usable[jen.Code], tfSource jen.Code) jen.Code {
	return tfSource
}
