package controller

import (
	"forumhub/auth"
	"html/template"
	"log"
	"net/http"

)
func createPost(w http.ResponseWriter, r *http.Request) {
//need to also add if!Authenticated
if r.Method != "Post" {
//check if user is autheticated and has a cookie

	//Redirect to login page
	http.Redirect(w,r,"/login", http.StatusSeeOther)
	return // return to avoid further execution
}

fTitle := r.FormValue("title")
fCategory := r.FormValue("category")
fBody := r.FormValue("body")
fUsername := GetUsername(r)

InsertPost(fUsername, fTitle, fCategory, fBody)


myTemplate, err := template.ParseFiles("./view/static/templates/layout.html", "./view/static/pages/content.html")
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


func GetUsername(){

}