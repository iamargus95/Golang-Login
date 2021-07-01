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

	http.HandleFunc("/login", loginpage)
	http.HandleFunc("/register", registerpage)
	http.ListenAndServe(":8000", nil)
}

func loginpage(w http.ResponseWriter, r *http.Request) {

	tpl = template.Must(template.ParseGlob("./login.html"))
	tpl.ExecuteTemplate(w, "login.html", nil)

	formemail := r.FormValue("email")
	formpsw := r.FormValue("psw")
	// onclickregister := r.FormValue("")

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

func registerpage(w http.ResponseWriter, r *http.Request) {

	formemail1 := r.FormValue("email")
	formpsw1 := r.FormValue("psw")

	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Println(err)
	}

	defer database.Close()

	stmt1, err1 := database.Prepare("INSERT INTO users (email, password) VALUES (?, ?)")
	stmt1.Exec(formemail1, formpsw1)
	if err1 != nil {
		log.Fatal(err1)
	}

	var email string

	err1 = stmt1.QueryRow(formemail1).Scan(email)
	if err1 != nil {
		log.Fatal(err1)
	}

	if email == formemail1 {
		fmt.Println("This email is already in use.")
		stmt1.Query("DELETE FROM users WHERE id = (SELECT MAX(id) FROM users")
	} else {
		fmt.Println("Successfully created a new account.")
	}
}
