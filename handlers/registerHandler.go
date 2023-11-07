package handlers

import (
	"forumhub/dal"
	"html/template"
	"net/http"

//	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
FirstName string
LastName string
UserName string
Email string


}


func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse form data
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		// Check if the email is already taken (you need to implement this)
		if IsEmailTaken(email) {
			// Handle the case where the email is already taken
			http.Error(w, "Email is already registered", http.StatusConflict)
			return
		}

		// // Generate a UUID for the user
		// userID := uuid.NewV4().String()
		// // Store user information in the database (you need to implement this)
		// if err := dal.CreateUser(userID, username, email, password); err != nil {
		// 	// Handle error
		// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// 	return
		// }

		// Redirect the user to the login page or home page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	// Render the registration form template (you need to create this template)
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		// Handle error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
