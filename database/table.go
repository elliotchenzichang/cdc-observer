package database

type DataType uint8

type Table struct {
	Fields []*Field
}

func (t *Table) AddField(name string, dataType DataType) {
	t.Fields = append(t.Fields, &Field{Name: name, Type: dataType})
}

type Field struct {
	Name string
	Type DataType
}
