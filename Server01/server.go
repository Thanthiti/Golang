package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))

	data := struct {
		Title string
		Name  string
	}{
		Title: "My Website",
		Name:  "Thanthiti",
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handle)
	fmt.Println("Serving at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
