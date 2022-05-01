package main

import (
	"net/http"

	"github.com/bichanna/go-web-practice/pkg/config"
	"github.com/bichanna/go-web-practice/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(appConfig *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

  // middlewares
  mux.Use(middleware.Recoverer)
  mux.Use(logToConsole)
  mux.Use(noSurf)

  // routes
  mux.Get("/", handlers.Repo.Home)
  mux.Get("/about", handlers.Repo.About)

  return mux
}