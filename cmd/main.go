package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
