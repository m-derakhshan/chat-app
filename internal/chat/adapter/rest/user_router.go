package rest

import (
	"github.com/gorilla/mux"
	"media.hiway.chat/internal/chat/controller/user/command"
	"media.hiway.chat/internal/chat/controller/user/query"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/user", query.GetUserByID).Methods("GET")
	router.HandleFunc("/user", command.AddUser).Methods("POST")
	router.HandleFunc("/user", command.UpdateUser).Methods("PATCH")
	router.HandleFunc("/user", command.DeleteUser).Methods("DELETE")
}
