package main

const cli_HELLO = "Hello :)"
const cli_NOT_FOUND_ERROR = "Connot undrestand command. \n -h for help"
const cli_HELP = "-h for help"
const cli_ENTER_A_VALID_COMMAND = "please enter a valid command."
const cli_SUCSESSFUL_ADDED_TODO = "ToDo Added to List"
const cli_SUCSESSFUL_REMOVED_TODO = "ToDo Removed from List"
const cli_NOTHING_TO_DO = "nothing to do! \nStart with: add 'todoTitle'"

type commands struct {
	ADD_PASSWORD_ITEM      []string
	GET_PASSWORD_ITEM      []string
	DELETE_PASSWORD_ITEM   []string
	DELETE_ALL_TODOS_ITEMS []string
	TODO_LIST_ITEMS        []string
	SURE_COMMAND           []string
	HELP_COMMAND           []string
}

var cm commands = commands{
	ADD_PASSWORD_ITEM:      []string{"add", "-n"},
	GET_PASSWORD_ITEM:      []string{"get", "-g"},
	DELETE_PASSWORD_ITEM:   []string{"del", "-d", "remove"},
	DELETE_ALL_TODOS_ITEMS: []string{"dal", "-c", "rma"},
	TODO_LIST_ITEMS:        []string{"l", "ls", "list"},
	SURE_COMMAND:           []string{"y", "Y", "Yes", "yes"},
	HELP_COMMAND:           []string{"-h", "--help", "?"},
}
