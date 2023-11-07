package main

import (
	"log"
	"net/http"
	"fmt"
	"forumhub/dal"
	"forumhub/handlers"
	"forumhub/view"
)

func executeTemplate(w http.ResponseWriter, filepath string){
	tmpl, err := view.Parse(filepath)
	if err != nil{
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
tmpl.Execute(w, nil)
}
func init() {
	dal.InitDatabase()
	fmt.Println("Success connected to database")

	executeTemplate(w,"./static")
	
}

func Server() {
	fs := http.FileServer(http.Dir("static"))

	//initialize a new servemux and register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", handlers.HomeHandler)

	//Start a new web server listening on port 7000
	log.Print("Starting server on :7001")
	err := http.ListenAndServe(":7001", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlers.HomeHandler(w, r)
	// case "/register":
	// 	RegisterHandler(w, r)
	// case "/login":
	// 	LoginHandler(w, r)
	default:
		handlers.ErrorHandler(w, r, 404)
	}

}

func main() {
	Server()
	http.HandleFunc("/", pathHandler)
}
