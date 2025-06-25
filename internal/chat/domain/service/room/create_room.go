package room

import (
	"encoding/json"
	"errors"
	"io"

	"media.hiway.chat/internal/chat/domain/port"
)

type createRoomRequest struct {
	Name string `json:"room_name"`
}

type createRoom struct {
	repository port.RoomRepository
}

func NewCreateRoom(repository port.RoomRepository) *createRoom {
	return &createRoom{
		repository: repository,
	}
}

func (c *createRoom) Execute(body io.ReadCloser) error {

	var req createRoomRequest

	if err := json.NewDecoder(body).Decode(&req); err != nil {
		return errors.New("invalid request body")
	}

	if req.Name == "" {
		return errors.New("room name cannot be empty")
	}

	if err := c.repository.CreateRoom(req.Name); err != nil {
		return err
	}

	return nil
}
