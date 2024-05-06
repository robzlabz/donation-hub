package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on :8180")
	http.ListenAndServe(":8180", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
}
