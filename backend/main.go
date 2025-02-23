// hello-world.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	port := ":7000"
	fmt.Println("Server starting at port 7000...")
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
