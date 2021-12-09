package main

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func addPassword(db *sql.DB, username string, domain string, _length string) string {
	tx, _ := db.Begin()
	lenInt, _ := strconv.Atoi(_length)
	password := passGen(lenInt)
	current_time := time.Now()
	stmt, _ := tx.Prepare("insert into passw (username,domain,password,create_time) values (?,?,?,?)")
	_, err := stmt.Exec(username, domain, password, current_time)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	return password
}

func getPasswords(db *sql.DB) [][]string {
	rows, err := db.Query("select * from passw")
	if err != nil {
		log.Fatal(err)
	}
	todos := [][]string{}
	for rows.Next() {
		var tempTodo toDo_model
		err =
			rows.Scan(&tempTodo.id, &tempTodo.username, &tempTodo.domain, &tempTodo.password, &tempTodo.created_at)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, []string{strconv.Itoa(tempTodo.id), tempTodo.username, tempTodo.domain, tempTodo.password, tempTodo.created_at.Format("2 Jan 2006 15:04")})
	}

	return todos
}
func getSinglePassword(db *sql.DB, domain string) [][]string {
	rows, err := db.Query("select * from passw")
	if err != nil {
		log.Fatal(err)
	}
	todos := [][]string{}
	for rows.Next() {
		var tempTodo toDo_model
		err =
			rows.Scan(&tempTodo.id, &tempTodo.username, &tempTodo.domain, &tempTodo.password, &tempTodo.created_at)
		if err != nil {
			log.Fatal(err)
		}
		if tempTodo.domain == domain {
			todos = append(todos, []string{strconv.Itoa(tempTodo.id), tempTodo.username, tempTodo.domain, tempTodo.password, tempTodo.created_at.Format("2 Jan 2006 15:04")})
		}
	}

	return todos
}
