package room

import (
	"errors"

	"media.hiway.chat/internal/chat/domain/model"
	"media.hiway.chat/internal/chat/domain/port"
)

type getRoomByID struct {
	repository port.RoomRepository
}

func NewGetRoomByID(repository port.RoomRepository) *getRoomByID {
	return &getRoomByID{
		repository: repository,
	}
}


func (g getRoomByID) Execute(roomID string) (*model.RoomModel, error) {

	if roomID == "" {
		return nil, errors.New("room id cannot be empty")
	}

	room, err := g.repository.GetRoomByID(roomID)
	if err != nil {
		return nil, err
	}

	return room, nil
}
