package main

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tpl = template.Must(template.ParseGlob("templates/*"))
)

func main() {
	http.Handle("/resources/*", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
