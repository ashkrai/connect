// hello-world.go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/rs/cors"
)

// Generate a random number
func randomNumberHandler(w http.ResponseWriter, r *http.Request) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) // Generate a random number between 0 and 99

	w.Header().Set("Content-Type", "application/json")
	// Return the random number as a JSON response
	fmt.Fprintf(w, `{"random_number": %d}`, randomNumber)
}

func main() {
	http.HandleFunc("/random-number", randomNumberHandler) // API endpoint

	port := ":7000"
	fmt.Println("Server starting at port 7000...")

	// Enable CORS for all origins
	handler := cors.Default().Handler(http.DefaultServeMux)

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal(err)
	}
}
