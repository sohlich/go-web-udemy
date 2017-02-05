package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"path/filepath"

	"github.com/goadesign/goa/uuid"
)

var tpl = template.Must(template.ParseGlob("./template/*.gohtml"))

func main() {
	log.Printf("Application %s starting.", "Photo blog")

	http.HandleFunc("/login", login)
	http.HandleFunc("/home", secure(home))
	http.Handle("/image/", http.StripPrefix("/image", http.FileServer(http.Dir("./storage"))))
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
	Name string
	URL  string
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("serving: Home")
	c, _ := r.Cookie("auth")

	var errString string
	if r.Method == http.MethodPost {
		if err := homePost(r, c.Value); err != nil {
			errString = err.Error()
		}
	}

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

	os.MkdirAll("./storage/"+userid, os.ModePerm)
	outF, e := os.OpenFile("./storage/"+userid+"/"+fh.Filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer outF.Close()
	if e != nil {
		return e
	}
	if _, e = io.Copy(outF, f); e != nil {
		return e
	}
	return nil
}

func listFiles(path string) []Picture {
	files, err := ioutil.ReadDir("./storage/" + path)
	if err != nil {
		return make([]Picture, 0)
	}

	paths := make([]Picture, len(files))
	for idx, file := range files {
		paths[idx].Name = file.Name()
		paths[idx].URL = filepath.Join("/image", path, file.Name())
	}
	return paths
}
