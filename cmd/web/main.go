package main

import (
	"log"
	"net/http"

	"github.com/bichanna/go-web-practice/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}