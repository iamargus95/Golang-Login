package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Hello World!!")

	database, _ := sql.Open("sqlite3", "./mydata.db")

	statement1, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, email TEXT, password TEXT)")
	statement1.Exec()

	statement2, _ := database.Prepare("INSERT INTO people (email, password) VALUES (? , ?)")
	statement2.Exec("email4", "password4")

	rows, _ := database.Query("SELECT id, email, password FROM people")

	var id int
	var email string
	var password string

	for rows.Next() {
		rows.Scan(&id, &email, &password)
		fmt.Println(strconv.Itoa(id) + ":" + email + " " + password)
	}
}
