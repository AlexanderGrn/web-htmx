package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Film struct {
	Title    string
	Director string
}

const LOG_FILE = "./log/log.log"

func main() {
	f, err := os.OpenFile(LOG_FILE, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second)
		log.Print("HTMX request received")
		log.Print(r.Header.Get("Hx-Request"))

		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		log.Print(title)
		log.Print(director)

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)
	log.Fatal(http.ListenAndServe(":80", nil))
}
