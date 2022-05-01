package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bichanna/go-web-practice/pkg/config"
	"github.com/bichanna/go-web-practice/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplate(ac *config.AppConfig) {
  app = ac
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.html");
	if err != nil { return myCache, err }

	for _, page := range pages {
    name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil { return myCache, err }

    matches, err := filepath.Glob("./templates/*.layout.html")
    if err != nil { return myCache, err }

    if len(matches) > 0 {
      templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")
      if err != nil {
        return myCache, err
      }
    }
    myCache[name] = templateSet
	}
	
  return myCache, nil
}

// addDefaultData 
func addDefaultData(tempData *models.TemplateData) *models.TemplateData {
  return tempData
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, temp string, tempData *models.TemplateData) {
  var tc map[string]*template.Template
  var err error

  if !app.UseCache {
    tc, _ = CreateTemplateCache()
  } else {
    tc = app.TemplateCache
  }

  t, ok := tc[fmt.Sprintf("%s.page.html", temp)]
  if !ok {
    log.Fatal("Could not get template from template cache") 
  }

  buf := new(bytes.Buffer)
  tempData = addDefaultData(tempData)
  if err = t.Execute(buf, tempData); err != nil {
    log.Println(err)
  }

  if _, err = buf.WriteTo(w); err != nil {
    log.Println("Error writing template to the browser")
  }
}