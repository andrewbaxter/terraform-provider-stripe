package main

import "github.com/dave/jennifer/jen"

type NodeFloat struct {
}

func (n *NodeFloat) GenSchema() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeFloat"),
	}
}

func (n *NodeFloat) Validate(path jen.Code, v jen.Code) *jen.Code {
	return nil
}
