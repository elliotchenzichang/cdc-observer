package database

type Table struct {
	Fields []*Field
}

func (t *Table) AddField(name string, dataType DataType) {
	t.Fields = append(t.Fields, &Field{Name: name, Type: dataType})
}
