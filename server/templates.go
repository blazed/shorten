package server

import "net/http"
import "html/template"

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) bool {
	if t, err := template.New(tmpl).ParseFiles("./templates/" + tmpl); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	} else if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	return true
}
