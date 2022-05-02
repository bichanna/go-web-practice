package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bichanna/go-web-practice/pkg/config"
	"github.com/bichanna/go-web-practice/pkg/handlers"
	"github.com/bichanna/go-web-practice/pkg/render"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
  // Change to true in production
  app.InProduction = false
  app.UseCache = false

  session = scs.New()
  session.Lifetime = 24 * time.Hour
  session.Cookie.Persist = true
  session.Cookie.SameSite = http.SameSiteLaxMode
  session.Cookie.Secure = app.InProduction
  
  templateCache, err := render.CreateTemplateCache()
  if err != nil {
    log.Fatal("Cannot create template cache", err)
  }
  app.TemplateCache = templateCache

  repo := handlers.NewRepo(&app)
  handlers.NewHandlers(repo)

  render.NewTemplate(&app)
	
	log.Printf("Starting application on port %s\n", portNumber)

  server := &http.Server{
    Addr: portNumber,
    Handler: routes(&app),
  }

  if err = server.ListenAndServe(); err != nil {
    log.Fatal(err)
  }
}