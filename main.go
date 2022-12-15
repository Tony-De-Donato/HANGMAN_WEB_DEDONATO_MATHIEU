package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpl1 := template.Must(template.ParseFiles("acceuil_hangman.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.FormValue("Hangman1") == "1" {
			tmpl1 = template.Must(template.ParseFiles("jeu_hangman.html"))
			lancer(1)
		}
		tmpl1.Execute(w, nil)
	})

	http.ListenAndServe(":80", nil)

}
func lancer(nbr int) {

}
