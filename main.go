package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	mot        string
	tentatives int
	pose       string
	deja       string
}

func main() {
	tmpl1 := template.Must(template.ParseFiles("jeu_hangman.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl1.Execute(w, nil)
			return
		}
		details := Page{
			mot:        "LEMOT",
			tentatives: 10,
			pose:       "img",
			deja:       "a,g,s,t,motdefou,rgb",
		}
		tmpl1.Execute(w, details)
	})
	http.ListenAndServe(":80", nil)
}
