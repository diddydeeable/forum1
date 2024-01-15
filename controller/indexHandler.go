package controller

import (
	"net/http"
	"html/template"
)

// indexHandler doesn't do anything but generate HTML and write it to the ResponseWriter
func IndexHandler(w http.ResponseWriter, r *http.Request){
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))	
	comments, err := model.Threads(); if err == nil {
		_, err := session(w, r)
		templates.ExecuteTemplate(w,"layout", comments)
	}
}