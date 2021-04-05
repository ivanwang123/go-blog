package stores

import (
	"fmt"

	"github.com/ivanwang123/go-blog/models"
	"github.com/jmoiron/sqlx"
)

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id int) (models.Post, error) {
	var post models.Post
	if err := s.Get(&post, `SELECT * FROM posts WHERE id=$1`, id); err != nil {
		return models.Post{}, fmt.Errorf("Error getting post: %w", err)
	}
	return post, nil
}

func (s *PostStore) Posts() ([]models.Post, error) {
	posts := []models.Post{}
	if err := s.Select(&posts, `SELECT * FROM posts`); err != nil {
		return []models.Post{}, fmt.Errorf("Error getting posts: %w", err)
	}
	return posts, nil
}

func (s *PostStore) PaginatedPosts(page int) ([]models.Post, int, error) {
	posts := []models.Post{}
	limit := 10
	offset := (page - 1) * limit
	if err := s.Select(&posts, `SELECT * FROM posts ORDER BY created_at DESC OFFSET $1 ROWS LIMIT $2`, offset, limit); err != nil {
		return []models.Post{}, 0, fmt.Errorf("Error getting posts: %w", err)
	}
	return posts, limit, nil
}

func (s *PostStore) CreatePost(p *models.Post) error {
	if err := s.Get(p, `INSERT INTO posts (title, content, user_id) VALUES ($1, $2, $3) RETURNING *`,
		p.Title, p.Content, p.UserId); err != nil {
		return fmt.Errorf("Error creating post: %w", err)
	}
	return nil
}

func (s *PostStore) UpdatePost(p *models.Post) error {
	if err := s.Get(p, `UPDATE posts SET title=$1, content=$2 WHERE id=$3 RETURNING *`,
		p.Title, p.Content, p.Id); err != nil {
		return fmt.Errorf("Error updating post: %w", err)
	}
	return nil
}

func (s *PostStore) DeletePost(id int) error {
	if _, err := s.Exec(`DELETE FROM posts WHERE id=$1`, id); err != nil {
		return fmt.Errorf("Error deleting post: %w", err)
	}
	return nil
}
