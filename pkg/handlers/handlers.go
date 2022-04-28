package handlers

import (
	"net/http"

	"github.com/bichanna/go-web-practice/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(w, "home")
}

// About is the about page handler
func About(w http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(w, "about")
}