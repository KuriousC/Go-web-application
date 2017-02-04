package main

import "fmt"

// --------------------------- Global Variable ------------------------------
var currentId int

var todos Todos

// --------------------------------------------------------------------------

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

// RepoFindTodo function will find the to do item with given id
func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

// RepoCreateTodo function will create To do item with given list
// in this case, todo id will be automatically generated
func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

// RepoDestroyTodo function will delete the to do item with given id
func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
