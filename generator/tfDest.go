package main

import "github.com/dave/jennifer/jen"

type TfDestColl interface {
	Field(key string) TfDestVal
	FlatField(key string) TfDestColl
}

type TfDestVal interface {
	Set(value jen.Code) jen.Code
}

// Map
type MapTfDestColl struct {
	Base     string
	TfPrefix string
}

func (d *MapTfDestColl) Field(key string) TfDestVal {
	return &MapTfDestVal{
		Base: d.Base,
		Key:  jen.Lit(d.TfPrefix + key),
	}
}

func (d *MapTfDestColl) FlatField(key string) TfDestColl {
	return &MapTfDestColl{
		Base:     d.Base,
		TfPrefix: d.TfPrefix + key + "_",
	}
}

type MapTfDestVal struct {
	Base string
	Key  jen.Code
}

func (d *MapTfDestVal) Set(val jen.Code) jen.Code {
	return jen.Id(d.Base).Index(d.Key).Op("=").Add(val)
}

// Root
type RootTfDestColl struct {
	Base     string
	TfPrefix string
}

func (d *RootTfDestColl) Field(key string) TfDestVal {
	return &RootTfDestVal{
		Base: d.Base,
		Key:  d.TfPrefix + key,
	}
}

func (d *RootTfDestColl) FlatField(key string) TfDestColl {
	return &RootTfDestColl{
		Base:     d.Base,
		TfPrefix: d.TfPrefix + key + "_",
	}
}

type RootTfDestVal struct {
	Base string
	Key  string
}

func (d *RootTfDestVal) Set(value jen.Code) jen.Code {
	return jen.Id(d.Base).Dot("Set").Call(jen.Lit(d.Key), value)
}

// Array
type ArrayTfDestVal string

func (d *ArrayTfDestVal) Set(val jen.Code) jen.Code {
	return jen.Op("*").Id(string(*d)).Op("=").Append(jen.Op("*").Id(string(*d)), val)
}
