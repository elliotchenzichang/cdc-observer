package database

// DatabaseManager this data structure aims to help the user build the database they want.
// the main functionalities of this structure as following:
// 1. sync the table from the same or the other database.
// 2. user can customized his database, such as adding, modifying and deleting a column
// 3. self serve defining one table with specific code.
// this project don't need to build multiple database, all tables you need can be sync into one unify database.
type DatabaseBuilder struct {
}

func NewDatabaseBuilder(opt *DatabaseBuilderOptions) *DatabaseBuilder {
	return nil
}
