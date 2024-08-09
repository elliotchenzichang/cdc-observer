package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Name          string
	dbClient      *gorm.DB
	tables        map[string]*Table
	pendingTables map[string]*Table
}

func NewDatabase(name string, addr string, port int, username string, password string) (*Database, error) {
	db := &Database{
		Name:          name,
		tables:        map[string]*Table{},
		pendingTables: map[string]*Table{},
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, addr, port, name)
	dbClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.dbClient = dbClient
	return db, nil
}

func (db *Database) ExistedTable(name string) bool {
	_, existed := db.tables[name]
	_, existedInPendingTables := db.pendingTables[name]
	return existed || existedInPendingTables
}

func (db *Database) AddTable(name string, table *Table) error {
	if db.ExistedTable(name) {
		return fmt.Errorf("table %s already existed", name)
	}
	db.pendingTables[name] = table
	return nil
}

func (db *Database) DeleteTable(name string) error {
	if !db.ExistedTable(name) {
		return fmt.Errorf("table %s doesn't existed", name)
	}
	delete(db.tables, name)
	delete(db.pendingTables, name)
	return nil
}

// Apply the database schema to real database
func (db *Database) Apply() error {
	return nil
}

// Clean the database env from database
func (db *Database) Clean() error {
	return nil
}
