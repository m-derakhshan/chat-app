package room

import (
	"errors"

	"media.hiway.chat/internal/chat/domain/model"
	"media.hiway.chat/internal/chat/domain/port"
)

type deleteRoom struct {
	respository port.RoomRepository
}

func NewDeleteRoom(repository port.RoomRepository) *deleteRoom {
	return &deleteRoom{
		respository: repository,
	}
}

func (c deleteRoom) Execute(roomID string) (*model.RoomModel, error) {

	if roomID == "" {
		return nil, errors.New("room id cannot be empty")
	}

	room, err := c.respository.DeleteRoom(roomID)
	if err != nil {
		return nil, err
	}

	return room, nil
}
