package rest

import (
	"github.com/gorilla/mux"
	"media.hiway.chat/internal/chat/controller/room/command"
	"media.hiway.chat/internal/chat/controller/room/query"
)

func RegisterRoomRoutes(router *mux.Router) {
	router.HandleFunc("/rooms", query.GetAllRooms).Methods(("GET"))
	router.HandleFunc("/room", query.GetRoomByID).Methods("GET")
	router.HandleFunc("/room", command.AddRoom).Methods(("POST"))
	router.HandleFunc("/room", command.UpdateRoomName).Methods("PATCH")
	router.HandleFunc("/room", command.DeleteRoom).Methods("DELETE")
}
