package main


import (
	"net/http"
	"fmt"
	"html/template"
)


func homeHandler(w http.ResponseWriter, r *http.Request)  {

	if r.URL.Path != "/"{
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return 
	}
}