package main

import (
	"UAGuide/database"
	"UAGuide/router"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	db := database.Connect()
	r := router.Init(db)
	log.Fatal(http.ListenAndServe(":8088", r))
}
