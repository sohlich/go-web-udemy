package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goadesign/goa/uuid"
)

var tpl = template.Must(template.ParseGlob("./template/*.gohtml"))

func main() {
	log.Printf("Application %s starting.", "Photo blog")

	http.HandleFunc("/login", login)
	http.HandleFunc("/home", secure(home))

	http.ListenAndServe(":8080", nil)

}

func secure(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check scure cookie
		if _, err := r.Cookie("auth"); err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		h(w, r)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("serving: Login")
	switch r.Method {
	case http.MethodGet:
		tpl.ExecuteTemplate(w, "login.gohtml", nil)
	case http.MethodPost:
		auth := &http.Cookie{
			Name:     "auth",
			Value:    uuid.NewV4().String(),
			Expires:  time.Now().Add(72 * time.Hour),
			HttpOnly: true,
		}
		http.SetCookie(w, auth)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

type Picture struct {
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("serving: Home")
	c, _ := r.Cookie("auth")

	if r.Method == http.MethodPost {
		f, fh, e := r.FormFile("photo")
		defer f.Close()
		if e != nil {
			log.Println("cannot read file from form: " + e.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		os.MkdirAll("./storage/"+c.Value, os.ModePerm)
		outF, e := os.OpenFile("./storage/"+c.Value+"/"+fh.Filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		defer outF.Close()
		if e != nil {
			log.Println("cannot open file: " + e.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if _, e = io.Copy(outF, f); e != nil {
			log.Println("cannot write file: " + e.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	data := struct {
		UserName string
		Pictures []Picture
	}{
		c.Value,
		make([]Picture, 0),
	}
	if err := tpl.ExecuteTemplate(w, "home.gohtml", data); err != nil {
		log.Println("err serving home: ", err.Error())
		return
	}
}
