package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io"
	"io/ioutil"

	"github.com/gorilla/mux"
)

// Index is the first landing page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// TodoIndex is the Todo list page
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	/*// We are commenting out the pre-defined todo list
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}*/

	// Set the contenttype to json, and tell client to expect json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// Write status code
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}

	/*if err := json.NewEncoder(w).Encode(json.MarshalIndent(todos, "", "\t")); err != nil {
		panic(err)
	}*/

	/*err := json.MarshalIndent(todos, "", "\t")
	if err != nil {
		fmt.Fprintln("error:", err)
	}*/
}

// TodoShow is the landing page displaying the todo list
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	/*for i := 0; i < 10; i++{
	    if Todos.
	}*/
	fmt.Fprintln(w, "Todo show:", todoId)
}

// TodoCreate cuntion creates handler for it
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	// Opening body for the Request
	// LimitReader is to protect against DOS attack
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	// we Unmarshal the json to store to our todo list
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		// error message if something goes wrong
		// in this case we are returning 422
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
	// if it all goes well then we return status 201
}
