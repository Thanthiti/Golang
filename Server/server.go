package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    tmpl := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Go Server UI</title>
        <link rel="stylesheet" href="Server/static/style.css">
        <script>
            async function fetchData() {
                let response = await fetch('/api/data');
                let data = await response.json();
                alert('Data: ' + JSON.stringify(data)); 
            }
        </script>
    </head>
    <body>
        <h1>Hello from Go!ssddd</h1>
        <button onclick="fetchData()">Fetch Data</button>
    </body>
    </html>
    `
    t, _ := template.New("webpage").Parse(tmpl)
    t.Execute(w, nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message":"Hello from API","status":"success"}`))
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/api/data", apiHandler)

    // Serve static files from the "static" directory
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":8080", nil)
    fmt.Println("Listing on port 8080....!!")
}
