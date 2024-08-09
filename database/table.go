package database

type Table struct {
	Fields map[string]*Field
}

func NewTable() *Table {
	return &Table{
		Fields: map[string]*Field{},
	}
}

func (t *Table) AddField(name string, dataType interface{}) {
	t.Fields[name] = &Field{Name: name, Type: dataType}
}

// Apply the table schema to real database
func (t *Table) Apply() error {
	return nil
}

// Clean the table env from database
func (t *Table) Clean() error {
	return nil
}

type TableBuilder struct {
	table *Table
}

func NewTableBuilder() *TableBuilder {
	return &TableBuilder{
		table: NewTable(),
	}
}

func (tb *TableBuilder) AddFieldTinyInt(name string) *TableBuilder {
	tb.table.AddField(name, new(TinyInt))
	return tb
}

func (tb *TableBuilder) AddFieldSmallInt(name string) *TableBuilder {
	tb.table.AddField(name, new(SmallInt))
	return tb
}

func (tb *TableBuilder) AddFieldMediumInt(name string) *TableBuilder {
	tb.table.AddField(name, new(MediumInt))
	return tb
}

func (tb *TableBuilder) AddFieldInt(name string) *TableBuilder {
	tb.table.AddField(name, new(Int))
	return tb
}

func (tb *TableBuilder) AddFieldBigInt(name string) *TableBuilder {
	tb.table.AddField(name, new(BigInt))
	return tb
}

func (tb *TableBuilder) AddFieldDecimal(name string) *TableBuilder {
	tb.table.AddField(name, new(Decimal))
	return tb
}

func (tb *TableBuilder) AddFieldFloat(name string) *TableBuilder {
	tb.table.AddField(name, new(Float))
	return tb
}

func (tb *TableBuilder) AddFieldDouble(name string) *TableBuilder {
	tb.table.AddField(name, new(SmallInt))
	return tb
}

func (tb *TableBuilder) AddFieldDate(name string) *TableBuilder {
	tb.table.AddField(name, new(Date))
	return tb
}

func (tb *TableBuilder) AddFieldTime(name string) *TableBuilder {
	tb.table.AddField(name, new(Time))
	return tb
}

func (tb *TableBuilder) AddFieldYear(name string) *TableBuilder {
	tb.table.AddField(name, new(Year))
	return tb
}

func (tb *TableBuilder) AddFieldDateTime(name string) *TableBuilder {
	tb.table.AddField(name, new(DateTime))
	return tb
}

func (tb *TableBuilder) AddFieldTimestamp(name string) *TableBuilder {
	tb.table.AddField(name, new(Timestamp))
	return tb
}

func (tb *TableBuilder) AddFieldChar(name string) *TableBuilder {
	tb.table.AddField(name, new(Char))
	return tb
}

func (tb *TableBuilder) AddFieldVarchar(name string) *TableBuilder {
	tb.table.AddField(name, new(Varchar))
	return tb
}

func (tb *TableBuilder) AddFieldText(name string) *TableBuilder {
	tb.table.AddField(name, new(Text))
	return tb
}

func (tb *TableBuilder) AddFieldBlob(name string) *TableBuilder {
	tb.table.AddField(name, new(Blob))
	return tb
}

func (tb *TableBuilder) AddFieldEnum(name string) *TableBuilder {
	tb.table.AddField(name, new(Enum))
	return tb
}

func (tb *TableBuilder) AddFieldSet(name string) *TableBuilder {
	tb.table.AddField(name, new(Set))
	return tb
}

func (tb *TableBuilder) AddFieldPoint(name string) *TableBuilder {
	tb.table.AddField(name, new(Point))
	return tb
}

func (tb *TableBuilder) AddFieldLineString(name string) *TableBuilder {
	tb.table.AddField(name, new(LineString))
	return tb
}

func (tb *TableBuilder) AddFieldJSON(name string) *TableBuilder {
	tb.table.AddField(name, new(Json))
	return tb
}

func (tb *TableBuilder) Submit() (*Table, error) {
	return tb.table, nil
}
