package model

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