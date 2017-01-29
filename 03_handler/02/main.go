package main

import (
	"html/template"
	"net/http"
)

var (
	tpl = template.Must(template.ParseGlob("*.gohtml"))
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleOther))
	http.ListenAndServe(":8080", mux)
}

func handleOther(rw http.ResponseWriter, req *http.Request) {
	data := struct {
		Name    string
		Surname string
	}{
		"Radek",
		"Sohlich",
	}
	tpl.ExecuteTemplate(rw, "Index", data)
}
