package room

import (
	"encoding/json"
	"errors"
	"io"

	"media.hiway.chat/internal/chat/domain/model"
	"media.hiway.chat/internal/chat/domain/port"
)

type getRoomRequest struct {
	ID string `json:"room_id"`
}

type getRoomByID struct {
	repository port.RoomRepository
}

func NewGetRoomByID(repository port.RoomRepository) *getRoomByID {
	return &getRoomByID{
		repository: repository,
	}
}

func (g getRoomByID) Execute(body io.ReadCloser) (*model.RoomModel, error) {

	var req getRoomRequest

	if err := json.NewDecoder(body).Decode(&req); err != nil {
		return nil, errors.New("invalid request body")
	}

	if req.ID == "" {
		return nil, errors.New("room ID cannot be empty")
	}

	room, err := g.repository.GetRoomByID(req.ID)
	if err != nil {
		return nil, err
	}

	return room, nil
}
