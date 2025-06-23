package room

import "net/http"

type RoomHandler struct {
}

func NewRoomHandler() *RoomHandler {
	return &RoomHandler{}
}

func (handler *RoomHandler) GetAllRooms(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
