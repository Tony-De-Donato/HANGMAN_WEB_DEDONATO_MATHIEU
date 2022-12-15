package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpl1 := template.Must(template.ParseFiles("acceuil_hangman.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.FormValue("Hangman1") == "1" {
			tpml2 := template.Must(template.ParseFiles("jeu_hangman.html"))
			http.HandleFunc("/jeu", func(w http.ResponseWriter, r *http.Request) {
				tpml2.Execute(w, nil)
			})
		}
		tmpl1.Execute(w, nil)
	})

	http.ListenAndServe(":80", nil)

}

func lancer(nbr int) {
	tpml2 := template.Must(template.ParseFiles("jeu_hangman.html"))
	http.HandleFunc("/jeu", func(w http.ResponseWriter, r *http.Request) {
		tpml2.Execute(w, nil)
	})
}
