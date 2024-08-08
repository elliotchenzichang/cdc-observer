package database

type Table struct {
	Fields map[string]*Field
}

func (t *Table) AddField(name string, dataType interface{}) {
	t.Fields[name] = &Field{Name: name, Type: dataType}
}
