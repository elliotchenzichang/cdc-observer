package database

import "fmt"

type Database struct {
	Name   string
	tables map[string]*Table
}

func NewDatabase(name string) *Database {
	return &Database{
		Name:   name,
		tables: map[string]*Table{},
	}
}

func (db *Database) ExistedTable(name string) bool {
	_, existed := db.tables[name]
	return existed
}

func (db *Database) AddTable(name string, table *Table) error {
	if db.ExistedTable(name) {
		return fmt.Errorf("table %s already existed", name)
	}
	db.tables[name] = table
	return nil
}
