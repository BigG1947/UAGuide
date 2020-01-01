package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Init() *mux.Router {
	r := mux.NewRouter()

	// ServeStaticFiles
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// SiteRoutes
	r.HandleFunc("/", index)
	r.HandleFunc("/signUp", signUp)
	r.HandleFunc("/signIn", signIn)

	// Test Routes
	r.HandleFunc("/api/city", callCityApi)
	return r
}


