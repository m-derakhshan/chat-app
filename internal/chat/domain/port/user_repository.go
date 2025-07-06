package port

import "media.hiway.chat/internal/chat/domain/model"

type UserRepository interface {
	CreateUser(firstName string, lastName string) (*model.UserModel, error)
	GetUserByID(userID string) (*model.UserModel, error)
	UpdateUser(userID string, firstName string, lastName string) (*model.UserModel, error)
	DeleteUser(userID string) (*model.UserModel, error)
}