package database

import "testing"

func TestNewDatabaseAndAddNewTable(t *testing.T) {
	db, err := NewDatabase("elliot_test_database", "127.0.0.1", 3307, "elliot_test", "123456")
	if err != nil {
		t.Fatal(err)
	}
	table, err := NewTableBuilder().AddFieldInt("test_field_int").Submit()
	if err != nil {
		t.Fatal(err)
	}
	err = db.AddTable("elliot_test_table", table)
	if err != nil {
		t.Fatal(err)
	}
}
