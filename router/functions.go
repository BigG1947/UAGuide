package router

import (
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

func checkLoginUser(session *sessions.Session) bool {
	if _, ok := session.Values["user"]; !ok || session.IsNew {
		return false
	}
	return true
}
func print500ErrorPage(writer http.ResponseWriter, request *http.Request, err error) {
	tmpl := template.Must(template.ParseFiles("templates/errors/500.html"))
	if err := tmpl.Execute(writer, err); err != nil {
		log.Printf("%s\n", err)
		return
	}
}
func uploadImages() error {
	return nil
}
