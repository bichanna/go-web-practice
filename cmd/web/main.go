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
	
	log.Printf("Starting application on port %s\n", portNumber)

  srv := &http.Server{
    Addr: portNumber,
    Handler: routes(&app),
  }

  if err = srv.ListenAndServe(); err != nil {
    log.Fatal(err)
  }
}