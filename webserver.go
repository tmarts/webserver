package main

import (
	"github.com/veloandy/webserver/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/viewxml", handlers.XMLHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
