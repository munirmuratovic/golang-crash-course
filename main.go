// main.go
package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{"message": "Hello from bare metal Go! Changed again!"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
