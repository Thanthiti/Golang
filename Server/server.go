package main

import (
    "net/http"
    "html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
    tmpl := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Go Server UI</title>
        <script>
            async function fetchData() {
                let response = await fetch('/api/data');
                let data = await response.json();
                alert('Data: ' + JSON.stringify(data));
            }
        </script>
    </head>
    <body>
        <h1>Hello from Go!</h1>
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
    http.ListenAndServe(":8080", nil)
}
