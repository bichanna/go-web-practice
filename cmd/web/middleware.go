package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

func logToConsole(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("Hit the page:", r.URL.String())
    next.ServeHTTP(w, r)
  })
}

func noSurf(next http.Handler) http.Handler {
  csrfHandler := nosurf.New(next)
  csrfHandler.SetBaseCookie(http.Cookie{
    HttpOnly: true,
    Path: "/",
    Secure: false,
    SameSite: http.SameSiteLaxMode,
  })
  return csrfHandler
}