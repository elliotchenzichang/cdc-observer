package database

import "testing"

func TestNewDatabaseAndAddNewTable(t *testing.T) {
	db, err := NewDatabase("elliot_test_database", "127.0.0.1", 3307, "root", "123456")
	if err != nil {
		t.Fatal(err)
	}
	table, err := NewTableBuilder("test_table_name").AddFieldInt("test_field_int").Submit()
	if err != nil {
		t.Fatal(err)
	}
	err = db.AddTable(table)
	if err != nil {
		t.Fatal(err)
	}
	err = db.Apply()
	if err != nil {
		t.Fatal(err)
	}
}
