package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/goadesign/goa/uuid"
)

const (
	STORAGE = "./storage"
)

var tpl = template.Must(template.ParseGlob("./template/*.gohtml"))

func main() {
	log.Printf("Application %s starting.", "Photo blog")
	http.HandleFunc("/login", login)
	http.HandleFunc("/home", secure(home))
	http.Handle("/image/", http.StripPrefix("/image", http.FileServer(http.Dir(STORAGE))))
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

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("serving: Home")
	c, _ := r.Cookie("auth")

	var errString string

	// Handle POST
	if r.Method == http.MethodPost {
		if err := homePost(r, c.Value); err != nil {
			errString = err.Error()
		}
	}

	// Load pictures and assemble
	// to page.
	pictures := listFiles(c.Value)
	data := struct {
		UserName string
		Pictures []Picture
		Error    string
	}{
		c.Value,
		pictures,
		errString,
	}
	if err := tpl.ExecuteTemplate(w, "home.gohtml", data); err != nil {
		log.Println("err serving home: ", err.Error())
		return
	}
}

func homePost(r *http.Request, userid string) error {
	f, fh, e := r.FormFile("photo")
	defer f.Close()
	if e != nil {
		return e
	}
	return storeFile(userid, fh.Filename, f)
}
