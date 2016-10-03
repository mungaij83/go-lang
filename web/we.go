package main

import (
	"fmt"
	"html/template"
	"net/http"
	//"io/ioutil"
)

type Page struct {
	Body  []byte
	Title string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func renderTemplate(w http.ResponseWriter, r *http.Request, page string, pg *Page) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Fprintf(w, "<p>Ops!! an errror occured loading page<p><p>", err, "</p>")
	}
	t.Execute(w, nil)
}
func other(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "index.html", nil)
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view", other)
	http.ListenAndServe(":8080", nil)
}
