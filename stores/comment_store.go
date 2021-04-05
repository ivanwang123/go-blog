package stores

import (
	"fmt"

	"github.com/ivanwang123/go-blog/models"
	"github.com/jmoiron/sqlx"
)

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) Comment(id int) (models.Comment, error) {
	var comment models.Comment
	if err := s.Get(&comment, `SELECT * FROM comments WHERE id=$1`, id); err != nil {
		return models.Comment{}, fmt.Errorf("Error getting comment: %w", err)
	}
	return comment, nil
}

func (s *CommentStore) CommentsByPost(postId int) ([]models.Comment, error) {
	comments := []models.Comment{}
	if err := s.Select(&comments, `SELECT * FROM comments WHERE post_id=$1`, postId); err != nil {
		return []models.Comment{}, fmt.Errorf("Error getting comments: %w", err)
	}
	return comments, nil
}

func (s *CommentStore) CreateComment(c *models.Comment) error {
	if err := s.Get(c, `INSERT INTO comments (content, post_id) VALUES ($1, $2) RETURNING *`,
		c.Content, c.PostId); err != nil {
		return fmt.Errorf("Error creating comment: %w", err)
	}
	return nil
}

func (s *CommentStore) UpdateComment(c *models.Comment) error {
	if err := s.Get(c, `UPDATE comments SET content=$1 WHERE id=$2 RETURNING *`,
		c.Content, c.Id); err != nil {
		return fmt.Errorf("Error updating comment: %w", err)
	}
	return nil
}

func (s *CommentStore) DeleteComment(id int) error {
	if _, err := s.Exec(`DELETE FROM comments WHERE id=$1`, id); err != nil {
		return fmt.Errorf("Error deleting comment: %w", err)
	}
	return nil
}
