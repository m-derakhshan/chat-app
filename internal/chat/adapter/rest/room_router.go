package rest

import (
	"github.com/gorilla/mux"
	"media.hiway.chat/internal/chat/controller/room/command"
	"media.hiway.chat/internal/chat/controller/room/query"
)

func RegisterRoomRoutes(router *mux.Router) {
	router.HandleFunc("/rooms", query.GetAllRooms).Methods(("GET"))
	router.HandleFunc("/rooms", command.AddRoom).Methods(("POST"))
}
