package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

	type Todo struct {
		ID   int    `json:"id"`
		Task string `json:"task"`
	}

	var todos  = []Todo{}
	var nextID = 1

	func addTodo(w http.ResponseWriter, r *http.Request) {
		var t Todo
		json.NewDecoder(r.Body).Decode(&t)
		fmt.Println("Do Here")
		t.ID = nextID
		nextID++
		todos = append(todos, t)
		json.NewEncoder(w).Encode(t)
		
	}
	
	func listTodos(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Show To do list")
		json.NewEncoder(w).Encode(todos)
	}

	func deleteTodo(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/todo/")
		id, _ := strconv.Atoi(idStr)
		for i, t := range todos {
			if t.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				break
			}
		}
		w.WriteHeader(http.StatusNoContent)
	}

	func main() {
		http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				listTodos(w, r)
			case "POST":
				addTodo(w, r)
			}
		})
		http.HandleFunc("/todo/", deleteTodo)
		http.ListenAndServe(":8080", nil)
	}
