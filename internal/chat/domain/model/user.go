package model

type UserModel struct {
	ID        string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
}
