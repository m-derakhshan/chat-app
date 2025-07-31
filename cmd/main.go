// @title           Chat API
// @version         1.0
// @description     The APIs of Chat application.
// @host            localhost:8080
// @BasePath        /

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "media.hiway.chat/docs"
	"media.hiway.chat/internal/chat/adapter/middleware"
	"media.hiway.chat/internal/chat/adapter/persistence"
	"media.hiway.chat/internal/chat/adapter/rest"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err:=persistence.RunMigrations(); err !=nil{
		log.Fatalf("Failed to run migrations: %v", err)
	}
	
	router := mux.NewRouter()

	rest.RegisterRoomRoutes(router)
	rest.RegisterUserRoutes(router)

	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	router.Use(middleware.LoggingMiddleware)

	log.Println("Starting server on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
