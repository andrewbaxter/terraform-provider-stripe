package main

import (
	"github.com/andrewbaxter/terraform-provider-stripe/shared"
	"github.com/dave/jennifer/jen"
)

type NodeObjFieldBehavior string

const (
	NBUserRequired         NodeObjFieldBehavior = NodeObjFieldBehavior("required")
	NBUserOptional         NodeObjFieldBehavior = NodeObjFieldBehavior("optional")
	NBUserOptionalComputed NodeObjFieldBehavior = NodeObjFieldBehavior("optionalcomputed")
	NBComputed             NodeObjFieldBehavior = NodeObjFieldBehavior("computed")
)

type NodeObjField struct {
	Description string
	Key         string
	Behavior    NodeObjFieldBehavior
	Default     jen.Code
	// Only read if named sibling field is true
	ConditionalOn string
	Updatable     bool // Only if user behavior
	Readable      bool // Must be true for computed
	Spec          Node
}

type NodeObj struct {
	FakeMapField bool // Workaround terraform limitations
	Fields       []*NodeObjField
}

func (n *NodeObj) genFieldInner(outFields jen.Dict, flattenParentOptional bool, flattenPrefix string) {
	for _, field := range n.Fields {
		if objField, isObj := field.Spec.(*NodeObj); isObj {
			objField.genFieldInner(outFields, flattenParentOptional || field.Behavior != NBUserRequired, flattenPrefix+field.Key+"_")
		} else {
			fieldOut := field.Spec.GenField()
			fieldOut[jen.Id("Description")] = jen.Lit(field.Description)
			behavior := field.Behavior
			if flattenParentOptional && behavior == NBUserRequired {
				behavior = NBUserOptional
			}
			switch behavior {
			case NBUserRequired:
				fieldOut[jen.Id("Required")] = jen.True()
				if !field.Updatable {
					fieldOut[jen.Id("ForceNew")] = jen.True()
				}
			case NBUserOptional:
				fieldOut[jen.Id("Optional")] = jen.True()
				if !field.Updatable {
					fieldOut[jen.Id("ForceNew")] = jen.True()
				}
				if field.Default != nil {
					fieldOut[jen.Id("Default")] = field.Default
				}
			case NBUserOptionalComputed:
				fieldOut[jen.Id("Optional")] = jen.True()
				fieldOut[jen.Id("Computed")] = jen.True()
				if !field.Updatable {
					fieldOut[jen.Id("ForceNew")] = jen.True()
				}
				if field.Default != nil {
					fieldOut[jen.Id("Default")] = field.Default
				}
			case NBComputed:
				fieldOut[jen.Id("Computed")] = jen.True()
			default:
				panic("ASSERTION")
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
		if !field.Readable || !field.Spec.CanRead() {
			continue
		}
		var read []jen.Code
		if objField, isObj := field.Spec.(*NodeObj); isObj {
			// Flatten nested objects
			read = objField.readApiFlat(
				jen.Qual(SharedPkg, "DigAny").Call(apiSource, jen.Lit(field.Key)),
				tfDest.FlatField(field.Key),
			)
		} else {
			statements := field.Spec.ReadApi(
				jen.Qual(SharedPkg, "DigAny").Call(apiSource, jen.Lit(field.Key)),
				tfDest.Field(field.Key),
			)
			if len(statements) > 1 {
				read = []jen.Code{jen.Block(statements...)}
			} else {
				read = []jen.Code{statements[0]}
			}
		}
		if field.ConditionalOn != "" {
			fields = append(
				fields,
				jen.
					If(jen.Qual(SharedPkg, "Dig").
						Index(jen.Bool()).
						Call(apiSource, jen.Lit(field.ConditionalOn))).
					Block(read...),
			)
		} else {
			fields = append(fields, read...)
		}
	}
	return fields
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
		if f.Behavior != NBUserRequired {
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
		// The required nested fields serve as keys that pull in all other fields
		if f.Behavior != NBUserRequired {
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
			cond.Op("||").Add(inner)
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
	// This is a workaround for required fields in optional nested objects
	// i.e. if all missing, object is missing therefore okay
	anyPresent := Unused(func() jen.Code { return jen.Id(anyPresentId) })
	for _, field := range n.Fields {
		if update && !field.Updatable {
			continue
		}
		if objField, isObj := field.Spec.(*NodeObj); isObj {
			// Flattened nested objects
			write := jen.Block(
				jen.List(jen.Id("res"), jen.Id("resOk")).Op(":=").Add(objField.ValidateSetApi(
					update,
					tfPath,
					tfSource.Field(field.Key),
				)),
				jen.If(jen.Id("resOk")).Block(jen.Id(apiDestId).Index(jen.Lit(field.Key)).Op("=").Id("res")),
			)
			if field.ConditionalOn != "" {
				fields = append(
					fields,
					jen.If(tfSource.Field(field.ConditionalOn).Get()).Block(write),
				)
			} else {
				fields = append(
					fields,
					write,
				)
			}
		} else {
			// Normal fields
			childChildPath := Unused(func() jen.Code { return jen.Id("path") })
			tfSourceField := tfSource.Field(field.Key)
			if_ := jen.If(jen.List(jen.Id("v"), jen.Id("vOk")).Op(":=").Add(tfSourceField.GetOk()).
				Op(";").Id("vOk").Op("&&").Add(field.Spec.IsNotDefault(jen.Id("v")))).Block(
				jen.List(jen.Id("res"), jen.Id("resOk")).Op(":=").Add(field.Spec.ValidateSetApi(
					update,
					childChildPath,
					shared.Pointer(AnyTfSourceVal("v")),
				)),
				jen.If(jen.Id("resOk")).Block(jen.Id(apiDestId).Index(jen.Lit(field.Key)).Op("=").Id("res")),
			)
			if field.Behavior == NBUserRequired {
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
					jen.Id("path").Op(":=").Add(tfPath.Use()).Dot("IndexString").Call(jen.Lit(tfSourceField.TfFieldName())),
				)
			}
			if field.ConditionalOn != "" {
				if_ = jen.If(jen.List(jen.Id("v"), jen.Id("vOk")).Op(":=").Add(tfSource.Field(field.ConditionalOn).GetOk()).
					Op(";").Id("vOk").Op("&&").Add(jen.Id("v").Assert(jen.Bool()))).Block(if_)
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
		Flatten([][]jen.Code{
			pre,
			fields,
			{
				jen.Return(jen.Id(apiDestId), jen.Len(jen.Id(apiDestId)).Op(">").Lit(0)),
			},
		}),
	)
}

func (n *NodeObj) IsNotDefault(id jen.Code) jen.Code {
	return jen.True()
}

func (n *NodeObj) CanRead() bool {
	for _, f := range n.Fields {
		if f.Readable && f.Spec.CanRead() {
			return true
		}
	}
	return false
}
