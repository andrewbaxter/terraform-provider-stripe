package main

import (
	"github.com/dave/jennifer/jen"
)

type TfSourceVal interface {
	Get() jen.Code
	GetOk() jen.Code
	TfFieldName() string
	Ok() jen.Code
	Field(key string) TfSourceVal
}

// Root (create)
type CreateRootTfSourceVal struct {
	Base string
	Key  string
}

func (d *CreateRootTfSourceVal) Get() jen.Code {
	return jen.Id(d.Base).Dot("Get").Call(jen.Lit(d.Key))
}

func (d *CreateRootTfSourceVal) GetOk() jen.Code {
	return jen.Id(d.Base).Dot("GetOk").Call(jen.Lit(d.Key))
}

func (d *CreateRootTfSourceVal) Ok() jen.Code {
	return jen.Id("inResourceData").Call(jen.Lit(d.Key), jen.Id(d.Base))
}

func (d *CreateRootTfSourceVal) TfFieldName() string {
	return d.Key
}

func (d *CreateRootTfSourceVal) Field(key string) TfSourceVal {
	var newKey string
	if d.Key == "" {
		newKey = key
	} else {
		newKey = d.Key + "_" + key
	}
	return &CreateRootTfSourceVal{
		Base: d.Base,
		Key:  newKey,
	}
}

// Root (update)
type UpdateRootTfSourceVal struct {
	Base string
	Key  string
}

func (d *UpdateRootTfSourceVal) Get() jen.Code {
	return jen.Id(d.Base).Dot("Get").Call(jen.Lit(d.Key))
}

func (d *UpdateRootTfSourceVal) GetOk() jen.Code {
	return jen.Id(d.Base).Dot("GetOk").Call(jen.Lit(d.Key))
}

func (d *UpdateRootTfSourceVal) Ok() jen.Code {
	return jen.Id(d.Base).Dot("HasChange").Call(jen.Lit(d.Key))
}

func (d *UpdateRootTfSourceVal) TfFieldName() string {
	return d.Key
}

func (d *UpdateRootTfSourceVal) Field(key string) TfSourceVal {
	var newKey string
	if d.Key == "" {
		newKey = key
	} else {
		newKey = d.Key + "_" + key
	}
	return &UpdateRootTfSourceVal{
		Base: d.Base,
		Key:  newKey,
	}
}

// Any
type AnyTfSourceVal string

func (d *AnyTfSourceVal) Get() jen.Code {
	return jen.Id(string(*d))
}

func (d *AnyTfSourceVal) GetOk() jen.Code {
	return jen.List(jen.Id(string(*d)), jen.True())
}

func (d *AnyTfSourceVal) Ok() jen.Code {
	return jen.True()
}

func (d *AnyTfSourceVal) TfFieldName() string {
	return string(*d)
}

func (d *AnyTfSourceVal) Field(key string) TfSourceVal {
	return &MapTfSourceVal{
		Base: string(*d),
		Key:  key,
	}
}

// Map
type MapTfSourceColl string

func (d *MapTfSourceColl) Field(key string) TfSourceVal {
	return &MapTfSourceVal{
		Base: string(*d),
		Key:  key,
	}
}

type MapTfSourceVal struct {
	Base string
	Key  string
}

func (d *MapTfSourceVal) Get() jen.Code {
	return jen.Id(d.Base).Assert(jen.Map(jen.String()).Any()).Index(jen.Lit(d.Key))
}

func (d *MapTfSourceVal) GetOk() jen.Code {
	return jen.Id(d.Base).Assert(jen.Map(jen.String()).Any()).Index(jen.Lit(d.Key))
}

func (d *MapTfSourceVal) Ok() jen.Code {
	return jen.Qual(SharedPkg, "InMap").Call(jen.Lit(d.Key), jen.Id(d.Base).Assert(jen.Map(jen.String()).Any()))
}

func (d *MapTfSourceVal) TfFieldName() string {
	return d.Key
}

func (d *MapTfSourceVal) Field(key string) TfSourceVal {
	return &MapTfSourceVal{
		Base: d.Base,
		Key:  d.Key + "_" + key,
	}
}
