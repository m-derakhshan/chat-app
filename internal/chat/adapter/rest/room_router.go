package rest

import (
	"github.com/gorilla/mux"
	"media.hiway.chat/internal/chat/app/room"
)

func RegisterRoomRoutes(router *mux.Router) {

	handler := room.NewRoomHandler()

	router.HandleFunc("/rooms", handler.GetAllRooms).Methods(("GET"))
}
