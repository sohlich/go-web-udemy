package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/me/", handleMe)
	mux.HandleFunc("/dog", handleDog)
	mux.HandleFunc("/", handleOther)
	http.ListenAndServe(":8080", mux)
}

func handleMe(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "I'm %s.", "Radomir")
}

func handleDog(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "Handling dog.")
}

func handleOther(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "Other")
}
