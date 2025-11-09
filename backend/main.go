package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Backend response for %s", r.URL.Path)
	}

	// Register both the exact path and the prefix so requests to /books and
	// /books/... are handled.
	http.HandleFunc("/books", handler)
	http.HandleFunc("/books/", handler)
	fmt.Println("Backend server running on :8081")
	http.ListenAndServe(":8081", nil)
}
