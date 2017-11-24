package main

import (
	"fmt"
	"net/http"
	"math/rand"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Random number: %d", rand.Intn(100))
	})
	http.ListenAndServe(":8080", nil)
}
