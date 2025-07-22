package user

import (
	"errors"

	"media.hiway.chat/internal/chat/domain/model"
	"media.hiway.chat/internal/chat/domain/port"
)

type getUserByID struct {
	repository port.UserRepository
}

func NewGetUserByID(repository port.UserRepository) *getUserByID {
	return &getUserByID{
		repository: repository,
	}
}

func (g getUserByID) Execute(userID string) (*model.UserModel, error) {

	if userID == "" {
		return nil, errors.New("user id cannot be empty")
	}

	user, err := g.repository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
