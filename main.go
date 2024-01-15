package main

import (
	"fmt"
	"forumhub/controller"
	"forumhub/model"
	"log"
	"net/http"
	//"forumhub/view"
)

func init() {
	model.InitDatabase()
	fmt.Println("Success connected to database")

}

func Server() {
	//create a file server to handle static files
	fs := http.FileServer(http.Dir("/view/static"))

	//initialize a new servemux and register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.Handle("/view/static/", http.StripPrefix("/view/static/", fs))
	mux.HandleFunc("/", pathHandler)

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
		controller.HomeHandler(w, r)
	case "/register":
		controller.RegisterHandler(w, r)
	case "/registerauth":
		controller.RegisterAuthHandler(w, r)
	case "/login":
		controller.LoginHandler(w, r)
	case "/loginauth":
		controller.LoginAuthHandler(w, r)

	default:
		controller.ErrorHandler(w, r, 404)
	}

}

func main() {
	Server()
	http.HandleFunc("/", pathHandler)
}
