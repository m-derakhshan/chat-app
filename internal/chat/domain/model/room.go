package model

type RoomModel struct {
	ID           string
	Name         string
	Participants map[string]*UserModel
}
