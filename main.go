package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var (
	tpl  *template.Template
	tpl1 *template.Template
)

func main() {

	fmt.Println("Starting listener on http://localhost:8080")
	http.HandleFunc("/", website)
	http.HandleFunc("/login", loginpage)
	http.HandleFunc("/register", registerpage)
	http.HandleFunc("/loginh", loginh)
	http.HandleFunc("/registerh", registerh)
	http.ListenAndServe(":8080", nil)
}

func website(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseGlob("./website.html"))
	tpl.ExecuteTemplate(w, "website.html", nil)
}

func loginpage(w http.ResponseWriter, r *http.Request) {

	tpl = template.Must(template.ParseGlob("./login.html"))
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func loginh(w http.ResponseWriter, r *http.Request) {

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
		fmt.Println("Login Successful.")
	} else {
		fmt.Println("Login Failed.")
	}

}

func registerpage(w http.ResponseWriter, r *http.Request) {

	tpl1 = template.Must(template.ParseGlob("./register.html"))
	tpl1.ExecuteTemplate(w, "register.html", nil)
}

func registerh(w http.ResponseWriter, r *http.Request) {

	formemail1 := r.FormValue("email")
	formpsw1 := r.FormValue("psw")

	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Println(err)
	}

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
		fmt.Println("Registeration Successful")
	}

	defer database.Close()

}
