package main

import (
	"github.com/veloandy/webserver/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/viewxml", handlers.XMLHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
