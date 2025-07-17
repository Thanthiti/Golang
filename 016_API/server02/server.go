package main

import (
	"encoding/json"
	"fmt"
	"net/http"
 )

type Echo struct {
	Text string `json:"Text"`
}


func echoHandle(w http.ResponseWriter, r *http.Request) {

	var input Echo
	json.NewDecoder(r.Body).Decode(&input)
	if input.Text == ""{
		input.Text = "Non Text!!"
	}
	resp := map[string]string{"Said": input.Text}
	w.Header().Set("Content-Type","Application/json")
	json.NewEncoder(w).Encode(resp)
}

func greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w,"Hello Word!")
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/echo", echoHandle)
	http.ListenAndServe(":8080", nil)
}