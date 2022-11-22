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

func (n *NodeString) Validate(path jen.Code, v jen.Code) *jen.Code {
	if len(n.Enum) > 0 {
		var out jen.Code = jen.If(
			jen.Id("inEnum").Call(
				jen.Add(v).Assert(jen.String()),
				jen.Index().String().Values(Map(
					n.Enum,
					func(e string) jen.Code { return jen.Lit(e) },
				)...),
			),
		).Block(jen.Id("out").Op("=").Id("append").Call(jen.Id("out"), jen.Qual(DiagPkg, "Diagnostic").Values(jen.Dict{
			jen.Id("Severity"):      jen.Qual(DiagPkg, "Error"),
			jen.Id("Summary"):       jen.Qual("fmt", "Sprintf").Call(jen.Lit(fmt.Sprintf("Field must be one of: %s", strings.Join(n.Enum, ", ")))),
			jen.Id("AttributePath"): path,
		})))
		return &out
	} else {
		return nil
	}
}
