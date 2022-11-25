package main

import (
	"github.com/dave/jennifer/jen"
)

type NodeObjField struct {
	Description string
	Key         string
	Computed    bool
	Required    bool
	Updatable   bool
	Readable    bool
	Spec        Node
}

type NodeObj struct {
	FakeMapField bool // Workaround terraform limitations
	Fields       []*NodeObjField
}

func (n *NodeObj) genFieldInner(outFields jen.Dict, flattenParentOptional bool, flattenPrefix string) {
	for _, field := range n.Fields {
		if objField, isObj := field.Spec.(*NodeObj); isObj {
			objField.genFieldInner(outFields, flattenParentOptional || !field.Required, flattenPrefix+field.Key+"_")
		} else {
			fieldOut := field.Spec.GenField()
			fieldOut[jen.Id("Description")] = jen.Lit(field.Description)
			if field.Computed {
				fieldOut[jen.Id("Computed")] = jen.True()
			} else {
				if !flattenParentOptional && field.Required {
					fieldOut[jen.Id("Required")] = jen.True()
				} else {
					fieldOut[jen.Id("Optional")] = jen.True()
				}
				if !field.Updatable {
					fieldOut[jen.Id("ForceNew")] = jen.True()
				}
			}
			outFields[jen.Lit(flattenPrefix+field.Key)] = jen.Values(fieldOut)
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
	n.genFieldInner(outFields, false, "")
	return jen.Op("&").Qual(SchemaPkg, "Resource").Values(jen.Dict{
		jen.Id("Schema"): jen.Map(jen.String()).Op("*").Qual(SchemaPkg, "Schema").Values(outFields),
	})
}

func (n *NodeObj) readApiInner(apiSource jen.Code, tfDest TfDestColl) []jen.Code {
	fields := []jen.Code{}
	for _, field := range n.Fields {
		if !field.Readable {
			continue
		}
		if objField, isObj := field.Spec.(*NodeObj); isObj {
			if !objField.checkAnyReadable() {
				continue
			}
			// Flatten nested objects
			fields = append(fields, objField.readApiFlat(
				jen.Qual(SharedPkg, "DigAny").Call(apiSource, jen.Lit(field.Key)),
				tfDest.FlatField(field.Key),
			)...)
		} else {
			statements := field.Spec.ReadApi(
				jen.Qual(SharedPkg, "DigAny").Call(apiSource, jen.Lit(field.Key)),
				tfDest.Field(field.Key),
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

func (n *NodeObj) checkAnyReadable() bool {
	for _, f := range n.Fields {
		if !f.Readable {
			continue
		}

		if objField, isObj := f.Spec.(*NodeObj); isObj {
			if objField.checkAnyReadable() {
				return true
			}
		} else {
			return true
		}
	}

	return false
}

func (n *NodeObj) readApiFlat(apiSource jen.Code, tfDest TfDestColl) []jen.Code {
	apiSourceId := "outerSource"
	return []jen.Code{
		jen.If(jen.List(jen.Id(apiSourceId), jen.Id("isMap")).Op(":=").Add(apiSource).Assert(jen.Map(jen.String()).Any()).Op(";").
			Id("isMap")).
			Block(Flatten([][]jen.Code{
				{jen.Comment("NodeObj readApiFlat")},
				n.readApiInner(jen.Id(apiSourceId), tfDest),
				{jen.Comment("NodeObj readApiFlat END")},
			})...),
	}
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
			{jen.Id("outerDest").Op(":=").Id(tfDestId)},
			n.readApiInner(jen.Id(apiSourceId), &MapTfDestColl{Base: "outerDest", TfPrefix: ""}),
		})...),
		tfDest.Set(jen.Id(tfDestId)),
		jen.Comment("NodeObj ReadApi END"),
	)
	return statements
}

func (n *NodeObj) checkAnyReq() bool {
	for _, f := range n.Fields {
		if !f.Required {
			continue
		}

		if objField, isObj := f.Spec.(*NodeObj); isObj {
			if objField.checkAnyReq() {
				return true
			}
		} else {
			return true
		}
	}

	return false
}

func (n *NodeObj) buildAnyTfFieldsPresent(tfSource TfSourceVal) *jen.Statement {
	var cond *jen.Statement = nil
	for _, f := range n.Fields {
		if !f.Required {
			continue
		}
		var inner jen.Code = nil
		if objField, isObj := f.Spec.(*NodeObj); isObj {
			inner1 := objField.buildAnyTfFieldsPresent(tfSource.Field(f.Key))
			if inner1 == nil {
				continue
			}
			inner = inner1
		} else {
			inner = tfSource.Field(f.Key).Ok()
		}
		if cond == nil {
			cond = jen.Add(inner)
		} else {
			cond.Op("&&").Add(inner)
		}
	}
	return cond
}

func (n *NodeObj) ValidateSetApi(update bool, tfPath *Usable[jen.Code], tfSource TfSourceVal) jen.Code {
	apiDestId := "dest"
	fields := []jen.Code{}
	if n.FakeMapField {
		fields = append(
			fields,
			jen.Id(apiDestId).Index(jen.Lit("key")).Op("=").Add(tfSource.Field("key").Get()),
		)
	}
	anyPresentId := "anyPresent"
	anyPresent := Unused(func() jen.Code { return jen.Id(anyPresentId) })
	for _, field := range n.Fields {
		if update && !field.Updatable {
			continue
		}
		if objField, isObj := field.Spec.(*NodeObj); isObj {
			// Flattened nested objects
			fields = append(
				fields,
				jen.Id(apiDestId).Index(jen.Lit(field.Key)).Op("=").Add(objField.ValidateSetApi(
					update,
					tfPath,
					tfSource.Field(field.Key),
				)),
			)
		} else {
			// Normal fields
			childChildPath := Unused(func() jen.Code { return jen.Id("path") })
			if_ := jen.If(jen.List(jen.Id("v"), jen.Id("vOk")).Op(":=").Add(tfSource.Field(field.Key).GetOk()).Op(";").Id("vOk")).Block(
				jen.Id(apiDestId).Index(jen.Lit(field.Key)).Op("=").Add(field.Spec.ValidateSetApi(
					update,
					childChildPath,
					P(AnyTfSourceVal("v")),
				)),
			)
			if field.Required {
				if_.Else().If(anyPresent.Use()).Block(
					jen.Id("out").Op("=").Id("append").Call(jen.Id("out"), jen.Qual(DiagPkg, "Diagnostic").Values(jen.Dict{
						jen.Id("Severity"): jen.Qual(DiagPkg, "Error"),
						jen.Id("Summary"): jen.Qual("fmt", "Sprintf").Call(
							jen.Lit("Field is missing but required. (at %s)"),
							jen.Id("fmtPath").Call(jen.Id("path")),
						),
						jen.Id("AttributePath"): childChildPath.Use(),
					})),
				)
			}
			pre := []jen.Code{}
			if childChildPath.WasUsed() {
				pre = append(
					pre,
					jen.Id("path").Op(":=").Add(tfPath.Use()).Dot("IndexString").Call(jen.Lit(field.Key)),
				)
			}
			fields = append(fields, jen.Block(Flatten([][]jen.Code{
				pre,
				{if_},
			})...))
		}
	}
	pre := []jen.Code{
		jen.Comment("NodeObj ValidateSetApi"),
		jen.Id(apiDestId).Op(":=").Map(jen.String()).Any().Values(jen.Dict{}),
	}
	if anyPresent.WasUsed() {
		cond := n.buildAnyTfFieldsPresent(tfSource)
		if cond == nil {
			cond = jen.False()
		}
		pre = append(pre, jen.Id(anyPresentId).Op(":=").Add(cond))
	}
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
