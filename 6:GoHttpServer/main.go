package main

import (
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to golang server Testing")
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Greeting %s!", name)
}
func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/greet", handleGreet)
	fmt.Println("Server is running in port 8080")
	http.ListenAndServe(":8080", nil)
}
