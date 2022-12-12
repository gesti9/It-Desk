package service

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func DataBase() {
	database, _ := sql.Open("sqlite3", "./users.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, email TEXT, phone TEXT, password TEXT)")
	statement.Exec()

	statement, _ = database.Prepare("INSERT INTO people (firstname, email, phone, password) VALUES (?,?,?,?)")
	// statement.Exec("Sultan", "sultan.astana0@gmail.com", "+77004010700", "Iddqd777")

	rows, _ := database.Query("SELECT id, firstname, email, phone, password  FROM people")
	var id int
	var firstname string
	var email string
	var phone string
	var password string

	for rows.Next() {
		rows.Scan(&id, &firstname, &email, &phone, &password)
		fmt.Println(id, firstname, email, phone, password)
	}
}
