package room

type CreateRoomRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"` // Room name must be between 1 and 100 characters
}
