package port

import "media.hiway.chat/internal/chat/domain/model"

type RoomRepository interface {
	CreateRoom(roomName string) (*model.RoomModel, error)

	GetAllRooms() (*[]model.RoomModel, error)

	GetRoomByID(roomID string) (*model.RoomModel, error)

	UpdateRoomName(roomID string, roomName string) (*model.RoomModel, error)

	DeleteRoom(roomID string) (*model.RoomModel, error)
}
