package user

import (
	"errors"

	"media.hiway.chat/internal/chat/domain/model"
	"media.hiway.chat/internal/chat/domain/port"
)

type deleteUser struct {
	respository port.UserRepository
}

func NewDeleteUser(repository port.UserRepository) *deleteUser {
	return &deleteUser{
		respository: repository,
	}
}

func (c deleteUser) Execute(userID string) (*model.UserModel, error) {

	if userID == "" {
		return nil, errors.New("user id cannot be empty")
	}

	user, err := c.respository.DeleteUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
