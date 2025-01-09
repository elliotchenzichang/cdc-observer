package database

import "testing"

func TestTable_Basic(t *testing.T) {
	table, err := NewTableBuilder("test_table_name", nil).AddFieldInt("test_field_int").Submit()
	if err != nil {
		t.Fatal(err)
	}
	err = table.Apply()
	if err != nil {
		t.Fatal(err)
	}
}
