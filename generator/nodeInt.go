package main

import "github.com/dave/jennifer/jen"

type NodeInt struct {
}

func (n *NodeInt) GenSchema() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeInt"),
	}
}

func (n *NodeInt) Validate(path jen.Code, v jen.Code) jen.Code {
	return jen.Block()
}
