package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Mot        string
	Tentatives int
	Pose       string
	Deja       string
}

func main() {
	tmpl1 := template.Must(template.ParseFiles("jeu_hangman.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		details := Page{
			Mot:        "LEMOT",
			Tentatives: 10,
			Pose:       "img",
			Deja:       "a",
		}
		tmpl1.Execute(w, details)
	})
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.ListenAndServe(":80", nil)
	
}
