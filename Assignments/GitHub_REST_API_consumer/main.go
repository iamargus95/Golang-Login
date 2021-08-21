package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/google/go-github/v38/github"
)

var (
	tpl          *template.Template
	formUsername string
)

func main() {
	fmt.Println("Starting listener on http://localhost:8080")
	http.HandleFunc("/", userName)
	http.ListenAndServe(":8080", nil)
}

func userName(w http.ResponseWriter, r *http.Request) {

	tpl = template.Must(template.ParseGlob("./index.html"))
	tpl.ExecuteTemplate(w, "index.html", nil)

	formUsername = r.FormValue("username")

	fmt.Println(formUsername)

	client := github.NewClient(nil)
}
