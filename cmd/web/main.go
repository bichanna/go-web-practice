package main

import (
	"log"
	"net/http"

	"github.com/bichanna/go-web-practice/pkg/config"
	"github.com/bichanna/go-web-practice/pkg/handlers"
	"github.com/bichanna/go-web-practice/pkg/render"
)

const portNumber = ":8080"

func main() {
  var app config.AppConfig
  
  templateCache, err := render.CreateTemplateCache()
  if err != nil {
    log.Fatal("Cannot create template cache", err)
  }
  app.TemplateCache = templateCache
  // Change to true
  app.UseCache = false

  repo := handlers.NewRepo(&app)
  handlers.NewHandlers(repo)

  render.NewTemplate(&app)
	
  http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}