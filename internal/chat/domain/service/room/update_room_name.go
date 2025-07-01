package room

import (
	"errors"

	"media.hiway.chat/internal/chat/domain/model"
	"media.hiway.chat/internal/chat/domain/port"
)

type updateRoomName struct {
	respository port.RoomRepository
}

func NewUpdateRoomName(repository port.RoomRepository) *updateRoomName {
	return &updateRoomName{
		respository: repository,
	}
}

func (c updateRoomName) Execute(roomID string, name string) (*model.RoomModel, error) {

	if roomID == "" {
		return nil, errors.New("room id cannot be empty")
	}
	if name == "" {
		return nil, errors.New("room name cannot be empty")
	}

	room, err := c.respository.UpdateRoomName(roomID, name)
	if err != nil {
		return nil, err
	}
	return room, nil

}
