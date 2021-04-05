package models

type Like struct {
	Id     int `json:"id"`
	UserId int `json:"userId" db:"user_id" validate:"required"`
	PostId int `json:"postId" db:"post_id" validate:"required"`
}
