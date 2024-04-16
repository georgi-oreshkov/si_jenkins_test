package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	queryParams := r.URL.Query()
	name := queryParams.Get("name")

	// If name is not provided, respond with a bad request status
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}

	// Write response
	response := fmt.Sprintf("Hello, %s!", name)
	fmt.Fprint(w, response)
}

func main() {
	// Define HTTP route
	http.HandleFunc("/hello", helloHandler)

	// Start HTTP server on port 8080
	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8070", nil))
}
