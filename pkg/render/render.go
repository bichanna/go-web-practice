package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, temp string) {
	parsedTemp, _ := template.ParseFiles(fmt.Sprintf("./templates/%s.html", temp))
	err := parsedTemp.Execute(w, nil)
	if err != nil {
		log.Println("error parsing template:", err)
	}
}
