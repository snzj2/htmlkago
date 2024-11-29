package main

import (
	"html/template"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("styles"))
	mux.Handle("/styles/", http.StripPrefix("/styles/", fs))
	fs1 := http.FileServer(http.Dir("image"))
	mux.Handle("/image/", http.StripPrefix("/image/", fs1))

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
