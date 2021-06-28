package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var tpl *template.Template

func main() {
	fmt.Println("Hello World!!")

	// database, _ := sql.Open("sqlite3", "./mydata.db")

	// statement1, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, email TEXT, password TEXT)")
	// statement1.Exec()

	// statement2, _ := database.Prepare("INSERT INTO people (email, password) VALUES (? , ?)")
	// statement2.Exec("email4", "password4") //Placeholder data. Get input from login.html

	// rows, _ := database.Query("SELECT id, email, password FROM people")

	// var (
	// 	id       int
	// 	email    string
	// 	password string
	// )

	http.HandleFunc("/", login)
	http.ListenAndServe(":8000", nil)
}

func init() {
	tpl = template.Must(template.ParseGlob("./login.html"))
}

func login(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.html", nil)
}
