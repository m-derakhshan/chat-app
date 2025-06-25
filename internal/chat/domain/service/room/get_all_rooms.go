package room

import (
	"media.hiway.chat/internal/chat/domain/model"
	"media.hiway.chat/internal/chat/domain/port"
)

type GetAllRooms struct {
	repository port.RoomRepository
}

func NewGetAllRooms(repository port.RoomRepository) *GetAllRooms {
	return &GetAllRooms{
		repository: repository,
	}
}

func (g *GetAllRooms) Execute() (*[]model.RoomModel, error) {

	rooms, err := g.repository.GetAllRooms()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}
