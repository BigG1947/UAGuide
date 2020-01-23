package router

import (
	"UAGuide/models"
	"UAGuide/pagination"
	"html/template"
	"log"
	"net/http"
)

func signUp(writer http.ResponseWriter, request *http.Request) {
	session, err := sessionStore.Get(request, userSessionName)
	if err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
		return
	}

	if checkLoginUser(session) {
		http.Redirect(writer, request, "/cabinet", 302)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/signUp.html"))
	err = tmpl.Execute(writer, nil)
	if err != nil {
		log.Printf("Error in login routes: %s\n", err)
		print500ErrorPage(writer, request, err)
	}
}

func registration(writer http.ResponseWriter, request *http.Request) {
	session, err := sessionStore.Get(request, userSessionName)
	if err != nil {
		log.Printf("Error in getting session: %s\n", err)
		print500ErrorPage(writer, request, err)
		return
	}

	if checkLoginUser(session) {
		http.Redirect(writer, request, "/cabinet", 302)
		return
	}

	var u models.User
	u.Email = request.FormValue("email")
	u.FIO = request.FormValue("fio")
	u.Phone = request.FormValue("phone")
	u.Password = []byte(request.FormValue("password"))
	if err := u.Add(db); err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
		return
	}

	session.Values["user"] = u.Id
	if err := session.Save(request, writer); err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
		return
	}
	http.Redirect(writer, request, "/cabinet", 302)
	return
}

func signIn(writer http.ResponseWriter, request *http.Request) {
	session, err := sessionStore.Get(request, userSessionName)
	if err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
		return
	}

	if checkLoginUser(session) {
		http.Redirect(writer, request, "/cabinet", 302)
		return
	}

	var errorMessage string
	var user models.User

	if request.Method == http.MethodPost {
		user.Email = request.FormValue("email")
		user.Password = []byte(request.FormValue("password"))
		if err := user.Login(db); err == nil {
			session.Values["user"] = user.Id
			if err := session.Save(request, writer); err != nil {
				log.Printf("%s\n", err)
				print500ErrorPage(writer, request, err)
				return
			}
			http.Redirect(writer, request, "/cabinet", 302)
			return
		} else {
			log.Printf("%s\n", err)
			errorMessage = "Не вірний логін чи пароль"
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/signIn.html"))
	err = tmpl.Execute(writer, map[string]interface{}{
		"errorMessage": errorMessage,
		"user":         user,
	})
	if err != nil {
		log.Printf("Error in login routes: %s\n", err)
		print500ErrorPage(writer, request, err)
	}
}

func exit(writer http.ResponseWriter, request *http.Request) {
	session, err := sessionStore.Get(request, userSessionName)
	if err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
		return
	}

	if !checkLoginUser(session) {
		http.Redirect(writer, request, "/signIn", 302)
		return
	}

	session.Options.MaxAge = -1
	if err := session.Save(request, writer); err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
		return
	}
	http.Redirect(writer, request, "/signIn", 302)
	return
}

func cabinet(writer http.ResponseWriter, request *http.Request) {
	session, err := sessionStore.Get(request, userSessionName)
	if err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
		return
	}

	if !checkLoginUser(session) {
		http.Redirect(writer, request, "/signIn", 302)
		return
	}

	var u models.User
	u.Id = session.Values["user"].(int)
	if err := u.Get(db); err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/cabinet/cabinet.html"))
	err = tmpl.Execute(writer, map[string]interface{}{
		"user": u,
	})
	if err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
	}

}

func settingProfile(writer http.ResponseWriter, request *http.Request) {
	session, err := sessionStore.Get(request, userSessionName)
	if err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
		return
	}

	if !checkLoginUser(session) {
		http.Redirect(writer, request, "/signIn", 302)
		return
	}

	var u models.User
	u.Id = session.Values["user"].(int)
	if err := u.Get(db); err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
		return
	}

	var successMessage string

	if request.Method == http.MethodPost {
		if len(request.FormValue("fio")) > 0 {
			u.FIO = request.FormValue("fio")
		}
		if len(request.FormValue("phone")) > 0 {
			u.Phone = request.FormValue("phone")
		}
		if len(request.FormValue("email")) > 0 {
			u.Email = request.FormValue("email")
		}
		if len(request.FormValue("photo")) > 0 {
			u.Photo = request.FormValue("photo")
		}
		if err := u.Update(db); err != nil {
			log.Printf("%s\n", err)
			print500ErrorPage(writer, request, err)
			return
		}

		successMessage = "Дані успішно збереженні."
	}

	tmpl := template.Must(template.ParseFiles("templates/cabinet/settings.html"))
	err = tmpl.Execute(writer, map[string]interface{}{
		"successMessage": successMessage,
		"user":           u,
	})
	if err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
	}
}

func userPlaces(writer http.ResponseWriter, request *http.Request) {
	session, err := sessionStore.Get(request, userSessionName)
	if err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
		return
	}

	if !checkLoginUser(session) {
		http.Redirect(writer, request, "/signIn", 302)
		return
	}

	var u models.User
	u.Id = session.Values["user"].(int)
	if err := u.Get(db); err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
		return
	}

	var placeList models.PlaceList
	_, err = pagination.NewPaginator(db, &placeList, 1, 10, nil)

	tmpl := template.Must(template.ParseFiles("templates/cabinet/places.html"))
	err = tmpl.Execute(writer, map[string]interface{}{
		"user": u,
	})
	if err != nil {
		log.Printf("%s\n", err)
		print500ErrorPage(writer, request, err)
	}
}
