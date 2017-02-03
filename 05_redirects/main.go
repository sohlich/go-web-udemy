package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"text/template"
	"time"
)

var (
	tpl = template.Must(template.ParseGlob("./templates/*"))
)

func main() {
	log.Printf("Application %s starting.", "Redirects")

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/home", homeHandler)
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		errMsg := r.FormValue("error")
		tpl.ExecuteTemplate(w, "Login", errMsg)
	case http.MethodPost:
		l := r.FormValue("login")
		p := r.FormValue("password")
		if l == "radek" && p == "pass" {
			authCookie := &http.Cookie{
				Name:     "auth",
				Value:    fmt.Sprintf("login:%s", l),
				HttpOnly: true,
				Expires:  time.Now().Add(48 * time.Hour),
			}

			http.SetCookie(w, authCookie)

			http.Redirect(w, r, "/home", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/login?error="+url.QueryEscape("bad login or password"), http.StatusSeeOther)
		}
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("auth")

	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	tmpl := "Home1"
	switch r.Method {
	case http.MethodGet:
		if r.FormValue("tmpl") != "" {
			tmpl = "Home2"
		}
	case http.MethodPost:
		c.MaxAge = -1 // Remove the cookie
		http.SetCookie(w, c)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	data := struct {
		Login    string
		Template string
	}{
		c.Value,
		tmpl,
	}

	tpl.ExecuteTemplate(w, "Home", data)
}
