package room

import (
	"errors"

	"media.hiway.chat/internal/chat/domain/model"
	"media.hiway.chat/internal/chat/domain/port"
)

type createRoom struct {
	repository port.RoomRepository
}

func NewCreateRoom(repository port.RoomRepository) *createRoom {
	return &createRoom{
		repository: repository,
	}
}

func (c *createRoom) Execute(roomName string) (*model.RoomModel, error) {

	if roomName == "" {
		return nil, errors.New("room name cannot be empty")
	}

	room, err := c.repository.CreateRoom(roomName)
	if err != nil {
		return nil, err
	}
	return room, nil
}
