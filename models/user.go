package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required,min=1,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"-" validate:"omitempty,min=8,max=255"`
}
