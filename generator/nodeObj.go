package main

import "github.com/dave/jennifer/jen"

type NodeObjField struct {
	Description string
	ApiKey      string
	TfKey       string
	Required    bool
	Spec        Node
}

type NodeObj struct {
	FakeMapField bool // Workaround terraform limitations
	Fields       []*NodeObjField
}

func (n *NodeObj) genFieldInner(outFields jen.Dict, flattenParentOptional bool) {
	for _, field := range n.Fields {
		if objField, isObj := field.Spec.(*NodeObj); isObj {
			objField.genFieldInner(outFields, flattenParentOptional || !field.Required)
		} else {
			fieldOut := field.Spec.GenField()
			fieldOut[jen.Id("Description")] = jen.Lit(field.Description)
			if !flattenParentOptional && field.Required {
				fieldOut[jen.Id("Required")] = jen.True()
			} else {
				fieldOut[jen.Id("Optional")] = jen.True()
			}
			outFields[jen.Lit(field.TfKey)] = jen.Values(fieldOut)
		}
	}
}

func (n *NodeObj) GenField() jen.Dict {
	panic("ASSERTION") // not possible via terraform spec; should already be flattened
}

func (n *NodeObj) GenElem() jen.Code {
	outFields := jen.Dict{}
	if n.FakeMapField {
		outFields[jen.Lit("key")] = jen.Values(jen.Dict{
			jen.Id("Type"):        jen.Qual(SchemaPkg, "TypeString"),
			jen.Id("Description"): jen.Lit("Key for this field in parent map (synthetic to work around Terraform limitations)"),
			jen.Id("Required"):    jen.True(),
		})
	}
	n.genFieldInner(outFields, false)
	return jen.Op("&").Qual(SchemaPkg, "Resource").Values(jen.Dict{
		jen.Id("Schema"): jen.Map(jen.String()).Op("*").Qual(SchemaPkg, "Schema").Values(outFields),
	})
}

func (n *NodeObj) readApiInner(apiSource jen.Code, tfDest TfDestColl) []jen.Code {
	fields := []jen.Code{}
	for _, field := range n.Fields {
		if objField, isObj := field.Spec.(*NodeObj); isObj {
			// Flatten nested objects
			fields = append(fields, objField.readApiInner(apiSource, tfDest)...)
		} else {
			statements := field.Spec.ReadApi(
				jen.Qual(SharedPkg, "Dig").Index(jen.Any()).Call(apiSource, jen.Lit(field.ApiKey)),
				tfDest.Field(field.TfKey),
			)
			if len(statements) > 1 {
				fields = append(fields, jen.Block(statements...))
			} else {
				fields = append(fields, statements[0])
			}
		}
	}
	return fields
}

func (n *NodeObj) ReadApi(apiSource jen.Code, tfDest TfDestVal) []jen.Code {
	apiSourceId := "outerSource"
	tfDestId := "dest"
	statements := []jen.Code{
		jen.Comment("NodeObj ReadApi"),
		jen.Id(apiSourceId).Op(":=").Add(apiSource),
		jen.Id(tfDestId).Op(":=").Map(jen.String()).Any().Values(),
	}
	if n.FakeMapField {
		statements = append(statements,
			jen.Id(tfDestId).Index(jen.Lit("key")).Op("=").Id("k"),
		)
	}
	statements = append(
		statements,
		jen.Block(Flatten([][]jen.Code{
			{
				jen.Id("outerDest").Op(":=").Id(tfDestId),
			},
			n.readApiInner(jen.Id(apiSourceId), P(MapTfDestColl("outerDest"))),
		})...),
		tfDest.Set(jen.Id(tfDestId)),
	)
	return statements
}

func (n *NodeObj) ValidateSetApi(tfPath *Usable[jen.Code], tfSource jen.Code) jen.Code {
	tfSourceId := "outerSource"
	apiDestId := "dest"
	fields := []jen.Code{}
	childPath := Unused(func() jen.Code { return jen.Id("path") })
	if n.FakeMapField {
		fields = append(
			fields,
			jen.Id(apiDestId).Index(jen.Lit("key")).Op("=").
				Qual(SharedPkg, "Dig").Index(jen.String()).Call(jen.Id(tfSourceId), jen.Lit("key")),
		)
	}
	for _, field := range n.Fields {
		if objField, isObj := field.Spec.(*NodeObj); isObj {
			// Flattened nested objects
			fields = append(
				fields,
				jen.Id(apiDestId).Index(jen.Lit(field.ApiKey)).Op("=").Add(objField.ValidateSetApi(childPath, jen.Id(tfSourceId))),
			)
		} else {
			// Normal fields
			childChildPath := Unused(func() jen.Code { return jen.Id("path") })
			if_ := jen.If(jen.Id("v").Op("!=").Nil()).Block(
				jen.Id(apiDestId).Index(jen.Lit(field.ApiKey)).Op("=").Add(field.Spec.ValidateSetApi(childChildPath, jen.Id("v"))),
			)
			if field.Required {
				if_.Else().Block(
					jen.Id("out").Op("=").Id("append").Call(jen.Id("out"), jen.Qual(DiagPkg, "Diagnostic").Values(jen.Dict{
						jen.Id("Severity"):      jen.Qual(DiagPkg, "Error"),
						jen.Id("Summary"):       jen.Lit("Field is missing but required."),
						jen.Id("AttributePath"): childChildPath.Use(),
					})),
				)
			}
			pre := []jen.Code{}
			if childChildPath.WasUsed() {
				pre = append(
					pre,
					jen.Id("path").Op(":=").Add(childPath.Use()).Dot("IndexString").Call(jen.Lit(field.TfKey)),
				)
			}
			fields = append(fields, jen.Block(Flatten([][]jen.Code{
				pre,
				{
					jen.Id("v").Op(":=").Qual(SharedPkg, "Dig").Index(jen.Any()).Call(jen.Id(tfSourceId), jen.Lit(field.TfKey)),
					if_,
				},
			})...))
		}
	}
	pre := []jen.Code{
		jen.Comment("NodeObj ValidateSetApi"),
	}
	if childPath.WasUsed() {
		pre = append(pre, jen.Id("path").Op(":=").Add(tfPath.Use()))
	}
	pre = append(
		pre,
		jen.Id(tfSourceId).Op(":=").Add(tfSource),
		jen.Id(apiDestId).Op(":=").Map(jen.String()).Any().Values(jen.Dict{}),
	)
	return FakeScope(
		jen.Any(),
		Flatten([][]jen.Code{
			pre,
			fields,
			{
				jen.Return(jen.Id(apiDestId)),
			},
		}),
	)
}
