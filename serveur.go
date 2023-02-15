package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	templateshtml := template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/siteferme", func(w http.ResponseWriter, r *http.Request) {
		mdp := r.PostFormValue("auth-login")
		if mdp == "lemotdepasseestreslongjetel'avoue" {

			http.Redirect(w, r, "/goodjob", http.StatusFound)

		}

		templateshtml.ExecuteTemplate(w, "index2.html", "")
	})

	http.HandleFunc("/goodjob", func(w http.ResponseWriter, r *http.Request) {

		templateshtml.ExecuteTemplate(w, "index3.html", "")
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
