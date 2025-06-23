package port

import "media.hiway.chat/internal/chat/domain/model"

type RoomRepository interface {
	CreateRoom(roomName string) error

	GetAllRooms() ([]*model.RoomModel, error)

	GetRoomByID(id string) (*model.RoomModel, error)

	UpdateRoom(id string, roomName string) error

	DeleteRoom(id string) error

	AddUserToRoom(roomID string, user *model.UserModel) error

	GetParticipants(roomID string) ([]*model.UserModel, error)

	RemoveUserFromRoom(roomID string, userID string) error
}
