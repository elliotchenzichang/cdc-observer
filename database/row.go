package database

type Rows []Row

// Row defines the real data in database,
// and the data should be fited to the table schema definition
type Row struct {
	data map[string]interface{}
}

func NewRow() *Row {
	return &Row{
		data: map[string]interface{}{},
	}
}

func (r *Row) AddField(name string, value interface{}) *Row {
	r.data[name] = value
	return r
}

func (r *Row) Validate() error {
	return nil
}

type RowBuilder struct {
	r *Row
}

func NewRowBuilder() *RowBuilder {
	return &RowBuilder{
		r: NewRow(),
	}
}

func (rb *RowBuilder) AddField(name string, value interface{}) *RowBuilder {
	rb.r.AddField(name, value)
	return rb
}

func (rb *RowBuilder) Submit() *Row {
	return rb.r
}
