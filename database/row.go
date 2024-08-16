package database

type Rows []Row

// Row defines the real data in database,
// and the data should be fited to the table schema definition
type Row struct {
	data map[string]interface{}
}

func NewRow(tableName string) *Row {
	return &Row{
		data: map[string]interface{}{},
	}
}

func (r *Row) AddField(name string, value interface{}) *Row {
	r.data[name] = value
	return r
}

func (r *Row) Submit() map[string]interface{} {
	return r.data
}

func (r *Row) Validate() error {
	return nil
}
