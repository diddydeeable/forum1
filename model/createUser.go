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

func GetUser(username string)(User, error){
	var result User
	stmt := "SELECT * FROM Users WHERE username = ?"
	
	row := db.QueryRow(stmt, username)
	
	err := row.Scan(&result.ID, &result.Username, &result.Email, &result.PasswordHash)
	//if err, username already taken
	if err != nil {
		return result, err
	}
	return result, nil
}

func HashPassword(password string) (string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt error :", err) 
	}
	return string(bytes)
}

// function to create a new user in the database
func CreateUser(Username, Email, Password string) error {
	GetUser(Username)

	var UserQuery string = "INSERT INTO Users(Username, Email, Password) VALUES(?,?,?)"

	hashedPassword := HashPassword(Password)

	//Prepare the SQL statement
	PrepUserStatement, err := db.Prepare(UserQuery)
	if err != nil  {
		return err
	}

	defer PrepUserStatement.Close()

	_, err = PrepUserStatement.Exec(Username, Email, hashedPassword)
	if err != nil {
		return err
	}
//fmt.Fprint(w,"congratulations, your account has been successfully created")

	return nil
}
