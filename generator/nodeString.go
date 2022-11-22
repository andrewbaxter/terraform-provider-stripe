package main

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type NodeString struct {
	Enum []string
}

func (n *NodeString) GenSchema() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeString"),
	}
}

func (n *NodeString) Validate(path jen.Code, v jen.Code) jen.Code {
	if len(n.Enum) > 0 {
		return jen.If(
			jen.Id("inEnum").Call(
				v,
				jen.Index().String().Values(Map(
					n.Enum,
					func(e string) jen.Code { return jen.Lit(e) },
				)...),
			),
		).Block(jen.Id("out").Op("=").Id("append").Call(jen.Id("out"), jen.Qual(DiagPkg, "Diagnostic").Values(jen.Dict{
			jen.Id("Severity"):      jen.Qual(DiagPkg, "Error"),
			jen.Id("Summary"):       jen.Qual("fmt", "Sprintf").Call(jen.Lit(fmt.Sprintf("Field must be one of: %s", strings.Join(n.Enum, ", ")))),
			jen.Id("AttributePath"): jen.Id("path"),
		})))
	} else {
		return jen.Block()
	}
}
