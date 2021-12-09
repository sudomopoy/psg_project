package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/fatih/structs"
)

func main() {
	arg := os.Args
	if len(arg) < 2 {
		fmt.Println(cli_HELLO)
		fmt.Println(cli_HELP)
		return
	}
	db, _ := sql.Open("sqlite3", "F:\\psg\\database\\passw.db")
	tdp_intro(arg[1:], db)
}
func tdp_intro(args []string, db *sql.DB) {
	db.Exec(`create table if not exists passw 
	(id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		username text, 
		domain text, 
		password text, 
		create_time datetime
	)`)
	if len(args) == 4 {
		flag := args[0]
		if contains(cm.ADD_PASSWORD_ITEM, flag) {
			pass := addPassword(db, args[1], args[2], args[3])
			fmt.Println(cli_SUCSESSFUL_ADDED_TODO)
			fmt.Println(pass)
		} else {
			fmt.Println(cli_NOT_FOUND_ERROR)
			return
		}
	} else if len(args) == 2 {
		flag := args[0]
		if contains(cm.GET_PASSWORD_ITEM, flag) {
			toDos := getSinglePassword(db, args[1])
			showTable(toDos)
		} else if contains(cm.HELP_COMMAND, flag) {
			for _, v := range structs.Values(cm) {
				fmt.Println(v)
			}
		}
	} else if len(args) == 1 {
		flag := args[0]
		if contains(cm.TODO_LIST_ITEMS, flag) {
			toDos := getPasswords(db)
			showTable(toDos)
		} else if contains(cm.HELP_COMMAND, flag) {
			for _, v := range structs.Values(cm) {
				fmt.Println(v)
			}
		}
	} else {
		fmt.Println(cli_NOT_FOUND_ERROR)
		return
	}
}
