package main

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type NodeString struct {
	Enum []string
}

func (n *NodeString) GenField() jen.Dict {
	return jen.Dict{
		jen.Id("Type"): jen.Qual(SchemaPkg, "TypeString"),
	}
}

func (n *NodeString) GenElem() jen.Code {
	return jen.Op("&").Qual(SchemaPkg, "Schema").Values(n.GenField())
}

func (n *NodeString) ReadApi(apiSource jen.Code, tfDest TfDestVal) []jen.Code {
	return []jen.Code{tfDest.Set(apiSource)}
}

func (n *NodeString) ValidateSetApi(update bool, tfPath *Usable[jen.Code], tfSource TfSourceVal) jen.Code {
	tfSourceId := "outerSource"
	statements := []jen.Code{
		jen.Id(tfSourceId).Op(":=").Add(tfSource.Get()),
	}
	if len(n.Enum) > 0 {
		statements = append(
			statements,
			jen.If(
				jen.Op("!").Id("inEnum").Call(
					jen.Id(tfSourceId).Assert(jen.String()),
					jen.Index().String().Values(Map(
						n.Enum,
						func(e string) jen.Code { return jen.Lit(e) },
					)...),
				),
			).Block(jen.Id("out").Op("=").Id("append").Call(jen.Id("out"), jen.Qual(DiagPkg, "Diagnostic").Values(jen.Dict{
				jen.Id("Severity"):      jen.Qual(DiagPkg, "Error"),
				jen.Id("Summary"):       jen.Qual("fmt", "Sprintf").Call(jen.Lit(fmt.Sprintf("Field must be one of: %s", strings.Join(n.Enum, ", ")))),
				jen.Id("AttributePath"): tfPath.Use(),
			}))),
		)
	}
	statements = append(statements, jen.Return(jen.Id(tfSourceId)))
	return FakeScope(jen.Any(), statements)
}
