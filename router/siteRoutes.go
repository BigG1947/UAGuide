package router

import (
	"html/template"
	"log"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(writer, nil)
	if err != nil {
		log.Printf("Error in index routes: %s\n", err)
	}
}
