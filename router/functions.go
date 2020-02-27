package router

import (
	"github.com/gorilla/sessions"
	"html/template"
	"io/ioutil"
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
func uploadProfileImages(file []byte) (string, error) {
	tempFile, err := ioutil.TempFile("upload/profile-photo/", "upload-photo-*.jpg")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	_, err = tempFile.Write(file)
	if err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}

func uploadPlacesImages(file []byte) (string, error) {
	tempFile, err := ioutil.TempFile("upload/places-photo/", "upload-photo-*.jpg")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	_, err = tempFile.Write(file)
	if err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}
