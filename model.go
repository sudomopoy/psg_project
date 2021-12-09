package main

import "time"

type toDo_model struct {
	id         int
	username   string
	domain     string
	password   string
	created_at time.Time
}
