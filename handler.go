package main


import (
	"net/http"
	"fmt"
	"html/template"
)




var tmpl = template.Must(template.template.ParseGlob("templates/*.html"))

func homeHandler(w http.ResponseWriter, r *http.Request)  {

	if r.URL.Path != "/"{
		http.Error(w, "404 Method Not Found", http.StatusNotFound)
		return 
	}

	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return 
	}

	tmpl.ExecuteTemplate(w, "home.html", []Artist)
}