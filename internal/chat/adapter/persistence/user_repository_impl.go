package persistence

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	uuid "github.com/satori/go.uuid"
	"media.hiway.chat/internal/chat/domain/model"
	"media.hiway.chat/internal/chat/domain/port"
)

type userRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) port.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) CreateUser(firstName string, lastName string) (*model.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	uuid := uuid.NewV4().String()
	query := `INSERT INTO users (id, first_name, last_name) VALUES ($1,$2, $3)`
	if _, err := r.db.Exec(ctx, query, uuid, firstName, lastName); err != nil {
		return nil, err
	}
	user := &model.UserModel{
		ID:        uuid,
		FirstName: firstName,
		LastName:  lastName,
	}
	return user, nil
}

func (r *userRepositoryImpl) GetUserByID(userID string) (*model.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `SELECT id, first_name, last_name FROM users WHERE id = $1`
	row := r.db.QueryRow(ctx, query, userID)
	user := &model.UserModel{}
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryImpl) UpdateUser(userID string, firstName string, lastName string) (*model.UserModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string
	if firstName == "" {
		query = `UPDATE users SET last_name = $1 WHERE id = $2 RETURNING id, first_name, last_name`
	} else if lastName == "" {
		query = `UPDATE users SET first_name = $1 WHERE id = $2 RETURNING id, first_name, last_name`
	} else {
		query = `UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3 RETURNING id, first_name, last_name`
	}

	var user model.UserModel

	if err := r.db.QueryRow(ctx, query, firstName, lastName, userID).Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *userRepositoryImpl) DeleteUser(userID string) (*model.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM users WHERE id = $1 RETURNING id, first_name, last_name`
	var user model.UserModel

	if err := r.db.QueryRow(ctx, query, userID).Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
		return nil, err
	}

	return &user, nil
}
