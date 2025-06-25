package rest

import (
	"github.com/gorilla/mux"
	"media.hiway.chat/internal/chat/app/room/command"
	"media.hiway.chat/internal/chat/app/room/query"
)

func RegisterRoomRoutes(router *mux.Router) {
	router.HandleFunc("/rooms", query.GetAllRooms).Methods(("GET"))
	router.HandleFunc("/rooms", command.AddRome).Methods(("POST"))
}
