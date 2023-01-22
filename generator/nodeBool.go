package main

import "github.com/dave/jennifer/jen"

type NodeBool struct {
}

func (n *NodeBool) GenField() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeBool"),
	}
}

func (n *NodeBool) GenElem() jen.Code {
	return jen.Op("&").Qual(SchemaPkg, "Schema").Values(n.GenField())
}

func (n *NodeBool) ReadApi(apiSource jen.Code, tfDest TfDestVal) []jen.Code {
	return []jen.Code{tfDest.Set(apiSource)}
}

func (n *NodeBool) ValidateSetApi(update bool, tfPath *Usable[jen.Code], tfSource TfSourceVal) jen.Code {
	return jen.List(tfSource.Get(), jen.Lit(true))
}

func (n *NodeBool) IsNotDefault(id jen.Code) jen.Code {
	return jen.Add(id).Assert(jen.Bool())
}

func (n *NodeBool) CanRead() bool {
	return true
}
