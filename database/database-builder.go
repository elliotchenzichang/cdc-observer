package database

import (
	"fmt"
)

// DatabaseManager this data structure aims to help the user build the database they want.
// the main functionalities of this structure as following:
// 1. sync the table from the same or the other database.
// 2. user can customized his database, such as adding, modifying and deleting a column
// 3. self serve defining one table with specific code.
// this project don't need to build multiple database, all tables you need can be sync into one unify database.
// todo build the database builder API first
type DatabaseBuilder struct {
	tables map[string]*Table
}

func NewDatabaseBuilder(opt *DatabaseBuilderOptions) *DatabaseBuilder {
	return nil
}

func (dbb *DatabaseBuilder) AddTable(name string, table *Table) error {
	if _, exist := dbb.tables[name]; exist {
		return fmt.Errorf("this table :%s is already existed", name)
	}
	dbb.tables[name] = table
	return nil
}

func (dbb *DatabaseBuilder) DeleteTable(name string) error {
	if _, exist := dbb.tables[name]; exist {
		return fmt.Errorf("can't delete a unextisted table: %s", name)
	}
	delete(dbb.tables, name)
	return nil
}
