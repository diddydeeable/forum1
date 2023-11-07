package view

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

//This global variable stores all the parsed templates
var Templates map[string]*template.Template

//func LoadTemplates loads all the templates from the directory
func LoadTemplates(templatesDir string) error{
templateFiles, err := filepath.Glob(filepath.Join(templatesDir,"*.html"))
if err != nil {
	return fmt.Errorf("error with layout file globbing: %w", err)
}

//Parse each template file
for _,templateFile := range templateFiles {
	//Base returns the last element of path -> filename will be used as the key for the map
	filename := filepath.Base(templateFile)

	//Parse the template file and add it to the map
	tmpl, err := template.ParseFiles(templateFile)

	if err != nil {
		return fmt.Errorf("error with passing file: %w", err)
	}

	    // Use the filename without the extension as the key in the map.
	key := strings.TrimSuffix(filename, filepath.Ext(filename))
		Templates[key] = tmpl
	}

	return nil

}

func(t Template) Execute(w http.ResponseWriter, data interface{}){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.tmpl.Execute(w, data)
	if err != nil {
		log.Printf("error executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

//We can parse the template even when we don't have access an http.ResponseWriter
//We can parse the template before our web server starts
func Parse(filepath string)(Template, error){
	tmpl, err := template.ParseFiles(filepath)
	if err != nil{
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		tmpl:tmpl,
		}, nil

}