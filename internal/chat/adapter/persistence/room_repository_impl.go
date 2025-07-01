package persistence

import (
	"context"
	"time"

	"media.hiway.chat/internal/chat/domain/model"
	"media.hiway.chat/internal/chat/domain/port"

	"github.com/jackc/pgx/v5/pgxpool"
	uuid "github.com/satori/go.uuid"
)

type roomRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewRoomRepository(db *pgxpool.Pool) port.RoomRepository {
	return &roomRepositoryImpl{db: db}
}

func (repository *roomRepositoryImpl) CreateRoom(roomName string) (*model.RoomModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	uuid := uuid.NewV4().String()
	_, err := repository.db.Exec(ctx, `INSERT INTO rooms (id, name) VALUES ($1, $2)`, uuid, roomName)

	room := &model.RoomModel{
		ID:           uuid,
		Name:         roomName,
		Participants: make(map[string]*model.UserModel),
	}
	return room, err
}

func (repository *roomRepositoryImpl) GetAllRooms() (*[]model.RoomModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := repository.db.Query(ctx, `SELECT id, name FROM rooms`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []model.RoomModel
	for rows.Next() {
		var room model.RoomModel
		if err := rows.Scan(&room.ID, &room.Name); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return &rooms, nil
}

func (repository *roomRepositoryImpl) GetRoomByID(roomID string) (*model.RoomModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := repository.db.QueryRow(ctx, `SELECT id, name FROM rooms WHERE id = $1`, roomID)

	var room model.RoomModel

	if err := row.Scan(&room.ID, &room.Name); err != nil {
		return nil, err
	}

	return &room, nil

}

func (repository *roomRepositoryImpl) UpdateRoomName(roomID string, roomName string) (*model.RoomModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "UPDATE rooms SET name = $1 WHERE id = $2 RETURNING id, name"
	var room model.RoomModel

	if err := repository.db.QueryRow(ctx, query, roomName, roomID).Scan(&room.ID, &room.Name); err != nil {
		return nil, err
	}
	return &room, nil
}

func (repository *roomRepositoryImpl) DeleteRoom(roomID string) (*model.RoomModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var room model.RoomModel
	query := "DELETE FROM rooms WHERE id = $1 RETURNING id, name"

	err := repository.db.QueryRow(ctx, query, roomID).Scan(&room.ID, &room.Name)
	if err != nil {
		return nil, err
	}

	return &room, nil
}
