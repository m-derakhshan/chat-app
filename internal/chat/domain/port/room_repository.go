package port

import "media.hiway.chat/internal/chat/domain/model"

type RoomRepository interface {
	CreateRoom(roomName string) error

	GetAllRooms() (*[]model.RoomModel, error)
}
