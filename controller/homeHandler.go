package controller

import (
	//"errors"
	//"html/template"

	"html/template"
	"log"
	"net/http"
	//"forumhub/view"
)




func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home Handler hit")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	myTemplate, err := template.ParseFiles("./view/static/templates/layout.html", "./view/static/pages/index.html")
	if err != nil {
		log.Print("error parsing the template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(myTemplate)

	if err := myTemplate.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Print("error executing the template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
