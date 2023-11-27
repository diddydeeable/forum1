package controller

import (
	"fmt"
	"forumhub/model"
	"html/template"
	"log"
	"net/http"
	"unicode"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Register Handler hit")
	if r.URL.Path != "/register" {
		http.NotFound(w, r)
		return
	}

	myTemplate, err := template.ParseFiles("./view/static/templates/layout.html", "./view/static/pages/register.html")
	if err != nil {
		log.Print("error parsing the template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	myTemplate.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Print("error executing the template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// registerAuthHandler creates new user in database
func RegisterAuthHandler(w http.ResponseWriter, r *http.Request) {
	/*
		1. check email criteria
		2. check username criteria
		3. check password criteria
		4. check if username already exists in database
		5. create bcrypt hash from password
		6. insert username and password hash into database
	*/

	fmt.Println("********** registerAuthHandler running ***********")

	r.ParseForm()

	username := r.FormValue("username")
	// checking if username is alphanumeric only
	var nameAlphaNumeric = true
	for _, char := range username {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			nameAlphaNumeric = false
		}
	}
	// check length of username and password
	var nameLength bool

	if len(username) >= 5 && len(username) <= 15 {
		nameLength = true
	}
	if !nameLength || !nameAlphaNumeric {
		myTemplate, err := template.ParseFiles("./view/static/templates/layout.html", "./view/static/pages/register.html")
		if err != nil {
			log.Print("you can't use that username, check the criteria")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//check this -> execute layout.html or register.html?
		myTemplate.ExecuteTemplate(w, "register.html", "Please check username and password criteria")

	}
	password := r.FormValue("password")
	fmt.Println("password: ", password, "password length: ", len(password))
	// variables that must pass for password creation
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength bool
	for _, char := range password {
		switch {
		//func IsNumber(r rune) bool
		case unicode.IsLower(char):
			pswdLowercase = true
		case unicode.IsUpper(char):
			pswdUppercase = true
		case unicode.IsNumber(char):
			pswdNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true
		}
	}
	if len(password) > 6 && len(password) < 21 {
		pswdLength = true
	}
	fmt.Println("pswdLowercase: ", pswdLowercase, "pswdUppercase: ", pswdUppercase, "pswdNumber: ", pswdNumber, "pswdSpecial: ", pswdSpecial, "pswdLength", pswdLength)

	// if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength {
	// 	// myTemplate, err := template.ParseFiles("./view/static/templates/layout.html", "./view/static/pages/register.html")
	// 	// if err != nil {
	// 	// 	log.Print("error parsing the template")
	// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	// 	return
	// 	// }
	// 	// data := model.PageData{
	// 	// 	ErrorMessage: "Please check username and password criteria",
	// 	// }
	// 	//check this -> execute layout.html or register.html?

	// }

	http.Redirect(w, r, "/", http.StatusSeeOther)
	//start cookie

	email := r.FormValue("email")

	model.HashPassword(password)
	model.CreateUser(username, email, password)

}
