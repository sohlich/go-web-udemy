package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	tpl = template.Must(template.ParseGlob("templates/*"))
)

func main() {
	http.HandleFunc("/special/data.dat", specialFileHandler)
	http.HandleFunc("/special/", specialHandler)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func specialHandler(w http.ResponseWriter, r *http.Request) {
	// Serve specific file or directory
	http.ServeFile(w, r, "./special/golden.jpg") // Serves file
}

func specialFileHandler(w http.ResponseWriter, r *http.Request) {
	// Serve content needs all data for the content
	http.ServeContent(w, r, "data", time.Now(), strings.NewReader("Hello"))
}
