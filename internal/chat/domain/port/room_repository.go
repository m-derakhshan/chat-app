package port

import "media.hiway.chat/internal/chat/domain/model"

type RoomRepository interface {
	GetAllRooms() ([]*model.RoomModel, error)

	GetRoomByID(id string) (*model.RoomModel, error)

	GetParticipants(roomID string) ([]*model.UserModel, error)

	AddUserToRoom(roomID string, user *model.UserModel) error

	RemoveUserFromRoom(roomID string, userID string) error
}
