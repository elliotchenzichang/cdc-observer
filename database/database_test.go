package database

import "testing"

func TestNewDatabaseAndAddNewTable(t *testing.T) {
	db := NewDatabase("elliot_test_database")
	table, err := NewTableBuilder().AddFieldInt("test_field_int").Submit()
	if err != nil {
		t.Fatal(err)
	}
	err = db.AddTable("elliot_test_table", table)
	if err != nil {
		t.Fatal(err)
	}
}
