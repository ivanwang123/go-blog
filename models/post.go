package models

import "time"

type Post struct {
	Id        int       `json:"id"`
	Title     string    `json:"title" validate:"required,min=1,max=255"`
	Content   string    `json:"content" validate:"required"`
	UserId    int       `json:"userId" db:"user_id" validate:"omitempty,required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
