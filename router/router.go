package router

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

var sessionStore *sessions.CookieStore
var db *sql.DB

const (
	userSessionName  = "user-session-uaguide"
	adminSessionName = "admin-session-uaguide"
)

func Init(conn *sql.DB) *mux.Router {
	sessionStore = sessions.NewCookieStore([]byte("SECRET-KEY"))
	sessionStore.MaxAge(0)
	db = conn
	r := mux.NewRouter()

	// ServeStaticFiles
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.PathPrefix("/upload/").Handler(http.StripPrefix("/upload/", http.FileServer(http.Dir("upload"))))

	// SiteRoutes
	r.HandleFunc("/", index)

	// UserRoutes
	r.HandleFunc("/signUp", signUp)
	r.HandleFunc("/registration", registration).Methods(http.MethodPost)
	r.HandleFunc("/signIn", signIn)
	r.HandleFunc("/exit", exit)
	r.HandleFunc("/cabinet", cabinet)
	r.HandleFunc("/cabinet/setting", settingProfile).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/cabinet/places", userPlaces)
	r.HandleFunc("/cabinet/places/suggestion", userPlacesSuggestion).Methods(http.MethodGet, http.MethodPost)

	// Test Routes
	r.HandleFunc("/api/city", callCityApi)
	r.HandleFunc("/api/checkPhone", checkPhone)
	r.HandleFunc("/api/checkEmail", checkEmail)
	r.HandleFunc("/api/uploadProfileImages", apiUploadProfileImages).Methods(http.MethodPost)
	r.HandleFunc("/api/uploadPlacesImages", apiUploadPlacesImages).Methods(http.MethodPost)
	return r
}
