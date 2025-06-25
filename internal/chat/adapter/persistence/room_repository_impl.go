package persistence

import (
	"context"
	"time"

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



func (repository *roomRepositoryImpl) CreateRoom(roomName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	uuid := uuid.NewV4().String()
	_, err := repository.db.Exec(ctx, `INSERT INTO rooms (id, name) VALUES ($1, $2)`, uuid, roomName)
	return err
}
