package stores

import (
	"fmt"

	"github.com/ivanwang123/go-blog/models"
	"github.com/jmoiron/sqlx"
)

type LikeStore struct {
	*sqlx.DB
}

func (s *LikeStore) LikedByUser(userId, postId int) bool {
	var exists int
	if err := s.Get(&exists, `SELECT COUNT(1) FROM likes WHERE user_id=$1 AND post_id=$2`, userId, postId); err != nil {
		return false
	}

	if exists > 0 {
		return true
	} else {
		return false
	}
}

func (s *LikeStore) LikesByPost(postId int) ([]models.Like, error) {
	likes := []models.Like{}
	if err := s.Select(&likes, `SELECT * FROM likes WHERE post_id=$1`, postId); err != nil {
		return []models.Like{}, fmt.Errorf("Error getting likes: %w", err)
	}
	return likes, nil
}

func (s *LikeStore) ToggleLike(l *models.Like) bool {
	var exists int
	if err := s.Get(&exists, `SELECT COUNT(1) FROM likes WHERE user_id=$1 AND post_id=$2`, l.UserId, l.PostId); err != nil {
		return false
	}

	if exists > 0 {
		if err := s.deleteLike(l.UserId, l.PostId); err != nil {
			return true
		}
		return false
	} else {
		if err := s.createLike(l.UserId, l.PostId); err != nil {
			return false
		}
		return true
	}
}

func (s *LikeStore) createLike(userId, postId int) error {
	if _, err := s.Exec(`INSERT INTO likes (user_id, post_id) VALUES ($1, $2) RETURNING *`,
		userId, postId); err != nil {
		return fmt.Errorf("Error creating like: %w", err)
	}
	return nil
}

func (s *LikeStore) deleteLike(userId, postId int) error {
	if _, err := s.Exec(`DELETE FROM likes WHERE user_id=$1 AND post_id=$2`, userId, postId); err != nil {
		return fmt.Errorf("Error deleting like: %w", err)
	}
	return nil
}
