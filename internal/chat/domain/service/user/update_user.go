package user

import (
	"errors"

	"media.hiway.chat/internal/chat/domain/model"
	"media.hiway.chat/internal/chat/domain/port"
)

type updateUser struct {
	repository port.UserRepository
}

func NewUpdateUser(repository port.UserRepository) *updateUser {
	return &updateUser{
		repository: repository,
	}
}

func (c *updateUser) Execute(userID string, firstName string, lastName string) (*model.UserModel, error) {

	if userID == "" {
		return nil, errors.New(("user ID cannot be empty"))
	}

	if firstName == "" && lastName == "" {
		return nil, errors.New(("both first name and last name cannot be empty"))
	}

	user, err := c.repository.UpdateUser(userID, firstName, lastName)
	if err != nil {
		return nil, err
	}
	return user, nil

}
