package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var tpl *template.Template

func main() {

	http.HandleFunc("/", login)
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":8000", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseGlob("./login.html"))
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {

	formemail := r.FormValue("email")
	formpsw := r.FormValue("psw")

	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Println(err)
	}

	defer database.Close()

	stmt, err := database.Prepare("SELECT password FROM users WHERE email=?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var password string
	err = stmt.QueryRow(formemail).Scan(&password)
	if err != nil {
		log.Fatal(err)
	}

	if password == formpsw {
		fmt.Println("YAY")
	} else {
		fmt.Println("KNOPE")
	}

}
