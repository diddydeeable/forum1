package model

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDatabase() {
	var err error
	//Open SQLite database
	db, err = sql.Open("sqlite3", "./model/forum.db")
	// fmt.Print("opened successfully")
	if err != nil {
		log.Fatal(err, "failure to open the schema file")
	}
	// Read the schema.sql file
	schema, err := os.ReadFile("./model/schema.sql")
	// fmt.Print("read successfully")
	if err != nil {
		log.Fatal(err,"failure to read the schema when initalizing db")
	}
	// Execute the SQL commands from the schema.sql file
	_, err = db.Exec(string(schema))
	// fmt.Print("executed successfully", err)
	if err != nil {
		log.Fatal(err, "failure to execute the schema when initalizing db")
	}
}
