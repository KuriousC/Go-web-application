package main

import "time"

// Todo defines the structure of the Todo
type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos is the slice of Todo
type Todos []Todo
