package port

type RoomRepository interface {
	CreateRoom(roomName string) error
}
