package user

import (
	"errors"

	"media.hiway.chat/internal/chat/domain/model"
	"media.hiway.chat/internal/chat/domain/port"
)

type createUser struct {
	repository port.UserRepository
}

func NewCreateUser(repository port.UserRepository) *createUser {
	return &createUser{
		repository: repository,
	}
}

func (c *createUser) Execute(firstName string, lastName string) (*model.UserModel, error) {

	if firstName == "" || lastName == "" {
		return nil, errors.New(("user first name and last name cannot be empty"))
	}

	user, err := c.repository.CreateUser(firstName, lastName)
	if err != nil {
		return nil, err
	}
	return user, nil

}
