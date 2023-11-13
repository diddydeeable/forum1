package model

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

// InitDB initializes the db variable to be used across the DAL
func InitDB(database *sql.DB) {
	db = database
}

func HashPassword(password string) (string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println(":()") 
	}
	return string(bytes)
}

// function to create a new user in the database
func CreateUser(Username, Email, Password string) error {
	var UserQuery string = "INSERT INTO Users(Username, Email, Password) VALUES(?,?,?)"

	hashedPassword := HashPassword(Password)

	//Prepare the SQL statement
	PrepUserStatement, err := db.Prepare(UserQuery)
	if err != nil {
		return err
	}

	defer PrepUserStatement.Close()

	_, err = PrepUserStatement.Exec(Username, Email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
