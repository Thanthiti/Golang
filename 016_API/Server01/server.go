package main

import (
	"encoding/json"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message":"Hello world"}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(resp)
}
func greetHandle(w http.ResponseWriter, r *http.Request) {
	
	name := r.URL.Query().Get("ggg")
	if name == ""{
		name = "Thanthiti"
	}
	resp := map[string]string{"message":"Hello my name's " + name +"!"}
	json.NewEncoder(w).Encode(resp)

}

func main() {
	// http.HandleFunc("/", greet)
	http.HandleFunc("/", greetHandle)
	http.ListenAndServe(":8080", nil)
}