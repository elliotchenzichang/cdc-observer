package database

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

var (
	FieldNoExist = errors.New("field no exist")
)

var (
	Int64Type   = reflect.TypeOf(int64(0))
	IntType     = reflect.TypeOf(0)
	StringType  = reflect.TypeOf("")
	BoolType    = reflect.TypeOf(true)
	Uint8Type   = reflect.TypeOf(uint8(0))
	Int8Type    = reflect.TypeOf(int8(0))
	Float64Type = reflect.TypeOf(0.0)
	UInt64Type  = reflect.TypeOf(uint64(0))
)

type StructBuilder struct {
	fields []reflect.StructField
}

type Struct struct {
	typ   reflect.Type
	index map[string]int
}

type Instance struct {
	instance reflect.Value
	index    map[string]int
}

func NewStructBuilder() *StructBuilder {
	return &StructBuilder{}
}

func (s Struct) New() *Instance {
	return &Instance{reflect.New(s.typ).Elem(), s.index}
}

func (b *StructBuilder) AddField(field string, typ reflect.Type, Tag reflect.StructTag) *StructBuilder {
	b.fields = append(b.fields, reflect.StructField{Name: field, Type: typ, Tag: Tag})
	return b
}

func (b *StructBuilder) AddString(name string, tag reflect.StructTag) *StructBuilder {
	return b.AddField(name, StringType, tag)
}

func (b *StructBuilder) AddBool(name string, tag reflect.StructTag) *StructBuilder {
	return b.AddField(name, BoolType, tag)
}

func (b *StructBuilder) AddInt(name string, tag reflect.StructTag) *StructBuilder {
	return b.AddField(name, IntType, tag)
}

func (b *StructBuilder) AddInt64(name string, tag reflect.StructTag) *StructBuilder {
	return b.AddField(name, Int64Type, tag)
}

func (b *StructBuilder) AddUint8(name string, tag reflect.StructTag) *StructBuilder {
	return b.AddField(name, Uint8Type, tag)
}

func (b *StructBuilder) AddInt8(name string, tag reflect.StructTag) *StructBuilder {
	return b.AddField(name, Int8Type, tag)
}

func (b *StructBuilder) AddFloat64(name string, tag reflect.StructTag) *StructBuilder {
	return b.AddField(name, Float64Type, tag)
}

func (b *StructBuilder) AddUint64(name string, tag reflect.StructTag) *StructBuilder {
	return b.AddField(name, UInt64Type, tag)
}

func (b *StructBuilder) Build() *Struct {
	stu := reflect.StructOf(b.fields)
	index := make(map[string]int)
	for i := 0; i < stu.NumField(); i++ {
		index[stu.Field(i).Name] = i
	}
	return &Struct{stu, index}
}

func (in *Instance) Field(name string) (reflect.Value, error) {
	if i, ok := in.index[name]; ok {
		return in.instance.Field(i), nil
	} else {
		return reflect.Value{}, FieldNoExist
	}
}
func (in *Instance) SetString(name, value string) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetString(value)
	}
}

func (in *Instance) SetBool(name string, value bool) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetBool(value)
	}
}

func (in *Instance) SetInt64(name string, value int64) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetInt(value)
	}
}

func (in *Instance) SetUint8(name string, value uint8) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetUint(uint64(value))
	}
}

func (in *Instance) SetInt8(name string, value int8) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetInt(int64(value))
	}
}

func (in *Instance) SetFloat64(name string, value float64) {
	if i, ok := in.index[name]; ok {
		in.instance.Field(i).SetFloat(value)
	}
}
func (in *Instance) Interface() interface{} {
	return in.instance.Interface()
}

func (in *Instance) Addr() interface{} {
	return in.instance.Addr().Interface()
}

func buildTag(columnName string) reflect.StructTag {
	tag := fmt.Sprintf(`db:"%s"`, columnName)
	return reflect.StructTag(tag)
}
