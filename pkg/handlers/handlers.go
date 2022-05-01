package handlers

import (
	"net/http"

	"github.com/bichanna/go-web-practice/pkg/config"
	"github.com/bichanna/go-web-practice/pkg/models"
	"github.com/bichanna/go-web-practice/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{a}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(repo *Repository) {
	Repo = repo
}

// Home is the home page handler
func (repo *Repository) Home(w http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(w, "home", &models.TemplateData{})
}

// About is the about page handler
func (repo *Repository) About(w http.ResponseWriter, request *http.Request) {
  stringMap := make(map[string]string)
  stringMap["test"] = "hello again"
	render.RenderTemplate(w, "about", &models.TemplateData{StringMap: stringMap})
}
