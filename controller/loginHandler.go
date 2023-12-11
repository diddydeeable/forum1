package controller

import (
	"forumhub/model"
	"html/template"
	"log"
	"net/http"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"forumhub/auth"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// check for cookie session
	log.Println("Login Handler hit")
	if r.URL.Path != "/login" {
		http.NotFound(w, r)
		return
	}

	myTemplate, err := template.ParseFiles("./view/static/templates/layout.html", "./view/static/pages/login.html")
	if err != nil {
		log.Print("error parsing the template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = myTemplate.ExecuteTemplate(w, "layout", nil)

	if err != nil {
		log.Print("error executing the template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func LoginAuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	person, err := model.GetUser(username)
	if err != nil {
		log.Print("error checking user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if CheckPasswordHash(password, person.PasswordHash) {
		sessionUuid, err := uuid.NewV4()
		if err != nil {
			log.Print("error generating UUID")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    sessionUuid.String(),
			HttpOnly: true,
			Path: "/",
			//10h 
			MaxAge:36000,
		}
		http.SetCookie(w, &cookie)
		auth.SetCookieToMap(cookie.Value,person)
		log.Print("cookie created")
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
  