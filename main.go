package main

import (
	"UAGuide/router"
	"log"
	"net/http"
)

func main() {
	r := router.Init()
	log.Fatal(http.ListenAndServe(":8088", r))
}
