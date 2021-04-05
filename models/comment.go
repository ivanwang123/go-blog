package models

import "time"

type Comment struct {
	Id        int       `json:"id"`
	Content   string    `json:"content" validate:"required"`
	PostId    int       `json:"postId" db:"post_id" validate:"omitempty,required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
