package main

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
}

// Handler function for the endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{Message: "Hello, World!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/hello", helloHandler) // Register endpoint
    http.ListenAndServe(":8080", nil)      // Start server on port 8080
}