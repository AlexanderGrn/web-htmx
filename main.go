package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World!")
		io.WriteString(w, "\n"+r.Method)
	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":80", nil))
}
