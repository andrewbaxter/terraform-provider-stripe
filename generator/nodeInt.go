package main

import "github.com/dave/jennifer/jen"

type NodeInt struct {
}

func (n *NodeInt) GenField() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeInt"),
	}
}

func (n *NodeInt) GenElem() jen.Code {
	return jen.Op("&").Qual(SchemaPkg, "Schema").Values(n.GenField())
}

func (n *NodeInt) ReadApi(apiSource jen.Code, tfDest TfDestVal) []jen.Code {
	return []jen.Code{tfDest.Set(apiSource)}
}

func (n *NodeInt) ValidateSetApi(update bool, tfPath *Usable[jen.Code], tfSource TfSourceVal) jen.Code {
	return tfSource.Get()
}
