package database

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type Table struct {
	name      string
	dbClient  *gorm.DB
	Structure interface{}
	Fields    map[string]*Field
}

func NewTable(name string) *Table {
	return &Table{
		name:   name,
		Fields: map[string]*Field{},
	}
}

func (t *Table) addField(name string, dataType string) {
	t.Fields[name] = &Field{Name: name, Type: dataType}
}

// Apply the table schema to real database
func (t *Table) Apply() error {
	builder := NewStructBuilder()
	for _, definition := range t.Fields {
		switch definition.Type {
		case SMALL_INT:
			builder.AddInt8(CamelString(definition.Name), reflect.StructTag(fmt.Sprintf("gorm: \"column:%s\"", definition.Name)))
		case MEDIUM_INT:
			builder.AddInt32(CamelString(definition.Name), reflect.StructTag(fmt.Sprintf("gorm: \"column:%s\"", definition.Name)))
		case INT:
			builder.AddInt(CamelString(definition.Name), reflect.StructTag(fmt.Sprintf("gorm: \"column:%s\"", definition.Name)))
		case BIG_INT:
			builder.AddInt64(CamelString(definition.Name), reflect.StructTag(fmt.Sprintf("gorm: \"column:%s\"", definition.Name)))
		case VARCHAR, TEXT, BLOB:
			builder.AddString(CamelString(definition.Name), reflect.StructTag(fmt.Sprintf("gorm: \"column:%s\"", definition.Name)))
		}
	}
	instance := builder.Build().New()
	t.Structure = instance.Interface()
	err := t.dbClient.Table(t.name).Migrator().CreateTable(t.Structure)
	return err
}

// Clean the table env from database
func (t *Table) Clean() error {
	return nil
}

func (t *Table) AddRow(r *Row) {

}

func (t *Table) UpdateRow() {

}

func (t *Table) DeleteRow() {

}

func (t *Table) AddRows(rs Rows) {

}

type TableBuilder struct {
	table *Table
}

func NewTableBuilder(name string) *TableBuilder {
	return &TableBuilder{
		table: NewTable(name),
	}
}

func (tb *TableBuilder) AddFieldTinyInt(name string) *TableBuilder {
	tb.table.addField(name, TINY_INT)
	return tb
}

func (tb *TableBuilder) AddFieldSmallInt(name string) *TableBuilder {
	tb.table.addField(name, SMALL_INT)
	return tb
}

func (tb *TableBuilder) AddFieldMediumInt(name string) *TableBuilder {
	tb.table.addField(name, MEDIUM_INT)
	return tb
}

func (tb *TableBuilder) AddFieldInt(name string) *TableBuilder {
	tb.table.addField(name, INT)
	return tb
}

func (tb *TableBuilder) AddFieldBigInt(name string) *TableBuilder {
	tb.table.addField(name, BIG_INT)
	return tb
}

func (tb *TableBuilder) AddFieldDecimal(name string) *TableBuilder {
	tb.table.addField(name, DECIMAL)
	return tb
}

func (tb *TableBuilder) AddFieldFloat(name string) *TableBuilder {
	tb.table.addField(name, FLOAT)
	return tb
}

func (tb *TableBuilder) AddFieldDouble(name string) *TableBuilder {
	tb.table.addField(name, SMALL_INT)
	return tb
}

func (tb *TableBuilder) AddFieldDate(name string) *TableBuilder {
	tb.table.addField(name, DTAE)
	return tb
}

func (tb *TableBuilder) AddFieldTime(name string) *TableBuilder {
	tb.table.addField(name, TIME)
	return tb
}

func (tb *TableBuilder) AddFieldYear(name string) *TableBuilder {
	tb.table.addField(name, YEAR)
	return tb
}

func (tb *TableBuilder) AddFieldDateTime(name string) *TableBuilder {
	tb.table.addField(name, DATETIME)
	return tb
}

func (tb *TableBuilder) AddFieldTimestamp(name string) *TableBuilder {
	tb.table.addField(name, TIMESTAMP)
	return tb
}

func (tb *TableBuilder) AddFieldChar(name string) *TableBuilder {
	tb.table.addField(name, CHAR)
	return tb
}

func (tb *TableBuilder) AddFieldVarchar(name string) *TableBuilder {
	tb.table.addField(name, VARCHAR)
	return tb
}

func (tb *TableBuilder) AddFieldText(name string) *TableBuilder {
	tb.table.addField(name, TEXT)
	return tb
}

func (tb *TableBuilder) AddFieldBlob(name string) *TableBuilder {
	tb.table.addField(name, BLOB)
	return tb
}

func (tb *TableBuilder) AddFieldEnum(name string) *TableBuilder {
	tb.table.addField(name, ENUM)
	return tb
}

func (tb *TableBuilder) AddFieldSet(name string) *TableBuilder {
	tb.table.addField(name, SET)
	return tb
}

func (tb *TableBuilder) AddFieldPoint(name string) *TableBuilder {
	tb.table.addField(name, POINT)
	return tb
}

func (tb *TableBuilder) AddFieldLineString(name string) *TableBuilder {
	tb.table.addField(name, LINESTRING)
	return tb
}

func (tb *TableBuilder) AddFieldJSON(name string) *TableBuilder {
	tb.table.addField(name, JSON)
	return tb
}

func (tb *TableBuilder) Submit() (*Table, error) {
	return tb.table, nil
}
