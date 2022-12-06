package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Choix   string
	Success bool
}

func main() {
	tmpl1 := template.Must(template.ParseFiles("acceuil_hangman.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl1.Execute(w, nil)
			return
		}
		details := User{
			Choix:   r.FormValue("choix"),
			Success: true,
		}
		tmpl1.Execute(w, details)
	})
	http.ListenAndServe(":80", nil)
}
