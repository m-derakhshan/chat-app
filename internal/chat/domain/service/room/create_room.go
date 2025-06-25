package room

import (
	"errors"

	"media.hiway.chat/internal/chat/domain/port"
)

func NewCreateRoom(repository port.RoomRepository) *CreateRoom {
	return &CreateRoom{
		repository: repository,
	}
}

type CreateRoom struct {
	repository port.RoomRepository
}

func (c *CreateRoom) Execute(roomName string) error {
	if roomName == "" {
		return errors.New("room name cannot be empty")
	}

	if err := c.repository.CreateRoom(roomName); err != nil {
		return err
	}

	return nil
}
