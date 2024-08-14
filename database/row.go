package database

type Rows []Row

// Row defines the real data in database,
// and the data should be fited to the table schema definition
type Row struct {
	table string
}

func NewRow(tableName string) *Row {
	return &Row{table: tableName}
}
