# cdc-observer

This project is inspired by the challenges encountered at the onset of my second career role, where I was tasked with developing a Change Data Capture (CDC) data synchronization system. Our test environment was far from ideal: each table in the live database had over 10 or 20 fields, some of which were in JSON type. Creating test data to observe the performance of my data processing logic after some data changes proved to be an arduous task.This project seeks to revolutionize this testing experience, making it more efficient and less stressful.

# getting started

## 1. easy to define a customized database

The demo example posted below can help you build the customized database. 

````go
func TestNewDatabaseAndAddNewTable(t *testing.T) {
	db, err := NewDatabase("elliot_test_database", "127.0.0.1", 3307, "root", "123456")
	if err != nil {
		t.Fatal(err)
	}
	table, err := NewTableBuilder("test_table").AddFieldInt("test_field_int").AddFieldVarchar("test_field_string").Submit()
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

````

# architecture

1. Build the docker container based on the user's customized database definition.
2. Run the docker container via Docker API in golang
3. User can trigger a CDC event by the command this project provided
4. Print the CDC event in std ouput.
   <img width="1415" alt="image" src="https://github.com/user-attachments/assets/e8ec487f-130b-4e39-8941-70c188afd318">


# reference

1. what is CDC: https://www.thoughtworks.com/en-au/insights/blog/architecture/change_data_capture