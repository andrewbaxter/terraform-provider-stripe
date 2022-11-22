package main

import "github.com/dave/jennifer/jen"

type NodeBool struct {
}

func (n *NodeBool) GenSchema() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeBool"),
	}
}

func (n *NodeBool) Validate(path jen.Code, v jen.Code) jen.Code {
	return jen.Block()
}
