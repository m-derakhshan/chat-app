package model

type RoomModel struct {
	ID           string                `json:"id"`
	Name         string                `json:"name"`
	Participants map[string]*UserModel `json:"participants,omitempty"`
}
