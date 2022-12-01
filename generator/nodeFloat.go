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

func (n *NodeFloat) ValidateSetApi(update bool, tfPath *Usable[jen.Code], tfSource TfSourceVal) jen.Code {
	return tfSource.Get()
}

func (n *NodeFloat) IsNotDefault(id jen.Code) jen.Code {
	return jen.Add(id).Assert(jen.Float64()).Op("!=").Lit(0.0)
}

func (n *NodeFloat) CanRead() bool {
	return true
}
