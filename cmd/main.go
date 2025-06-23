package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"media.hiway.chat/internal/chat/adapter/middleware"
	"media.hiway.chat/internal/chat/adapter/rest"
)

func main() {
	router := mux.NewRouter()
	rest.RegisterRoomRoutes(router)

	router.Use(middleware.LoggingMiddleware)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
