package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

//logHitPath prints to the console which path has been hit.
func logHitPath(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("Hit the page:", r.URL.String())
    next.ServeHTTP(w, r)
  })
}

// noSurf adds CSRF protection to all post requests.
func noSurf(next http.Handler) http.Handler {
  csrfHandler := nosurf.New(next)
  csrfHandler.SetBaseCookie(http.Cookie{
    HttpOnly: true,
    Path: "/",
    Secure: app.InProduction,
    SameSite: http.SameSiteLaxMode,
  })
  return csrfHandler
}

// sessionLoad loads and saves the session on every request.
func sessionLoad(next http.Handler) http.Handler {
  return session.LoadAndSave(next)
}