package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

)

// Model
	type Todo struct {
		ID   int    `json:"id"`
		Task string `json:"task"`
	}
	//Controller 
	type TodoController struct{
		todos []Todo
		nextID int
	}

	// Constructor FN
	func NewTodoController() *TodoController{
		return &TodoController{
			todos: []Todo{},
			nextID: 1,
		}
	}
	
	// Method Post
	func (c *TodoController) addTodo(w http.ResponseWriter, r *http.Request) {
		var t Todo

		if err := json.NewDecoder(r.Body).Decode(&t);err != nil{
			http.Error(w,`{"Error" : "Invalid Json"}`,http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(t.Task)== ""{
			http.Error(w,`{"Error" : "Task is required"}`,http.StatusBadRequest)
			return
		}

		t.ID = c.nextID 	
		c.nextID++		
		c.todos = append(c.todos, t)
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(t)
	}
	
	// Method Get
	func (c *TodoController) listTodos(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(c.todos)
	}

	// Method Put
	func (c *TodoController) UpdateTodos(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idstr := strings.TrimPrefix(r.URL.Path,"/todo/edit/")
		id,err  := strconv.Atoi(idstr)

		if err != nil {
			http.Error(w,`{"error" : "invalid ID (PUT)"}`,http.StatusBadRequest)
			return
		}
		for _,data := range c.todos{
			if id == data.ID{
				json.NewEncoder(w).Encode(data)
				break
			}
		}
	}
	
	// Method Delete
	func (c *TodoController) deleteTodo(w http.ResponseWriter, r *http.Request) {

		idStr := strings.TrimPrefix(r.URL.Path, "/todo/")
		id, err := strconv.Atoi(idStr)

		if err != nil || id <= 0{
			http.Error(w,`{"error" : "invalid ID"}`,http.StatusBadRequest)
		}
		
		found := false
		for i, data := range c.todos {
			if data.ID == id {
				c.todos = append(c.todos[:i], c.todos[i+1:]...)
				found = true
				break
			}
		}
		if !found{
			http.Error(w,`{"error" : "Todo not found"}`,http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}

	func main() {
		ctrl := NewTodoController()
		fmt.Println(ctrl.nextID)
		http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				ctrl.listTodos(w, r)
			case http.MethodPost:
				ctrl.addTodo(w, r)
			}
		})
		http.HandleFunc("/todo/", ctrl.deleteTodo)
		http.HandleFunc("/todo/edit/", ctrl.UpdateTodos)
		http.ListenAndServe(":8080", nil)
	}
