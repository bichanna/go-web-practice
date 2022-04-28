package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// renderTemplate parses template and writes it
func renderTemplate(w http.ResponseWriter, temp string) {
	parsedTemp, _ := template.ParseFiles(fmt.Sprintf("./templates/%s.html", temp))
	err := parsedTemp.Execute(w, nil)
	if err != nil {
		log.Println("error parsing template:", err)
	}
}
