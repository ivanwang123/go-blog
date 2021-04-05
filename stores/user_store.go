package stores

import (
	"fmt"

	"github.com/ivanwang123/go-blog/models"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	*sqlx.DB
}

func (s *UserStore) User(id int) (models.User, error) {
	var user models.User
	if err := s.Get(&user, `SELECT * FROM users WHERE id=$1`, id); err != nil {
		return models.User{}, fmt.Errorf("Error getting user: %w", err)
	}
	return user, nil
}

func (s *UserStore) UserByEmail(email string) (models.User, error) {
	var user models.User
	if err := s.Get(&user, `SELECT * FROM users WHERE email=$1`, email); err != nil {
		return models.User{}, fmt.Errorf("Error getting user: %w", err)
	}
	return user, nil
}

func (s *UserStore) Users() ([]models.User, error) {
	users := []models.User{}
	if err := s.Select(&users, `SELECT * FROM users`); err != nil {
		return []models.User{}, fmt.Errorf("Error getting users: %w", err)
	}
	return users, nil
}

func (s *UserStore) CreateUser(u *models.User) error {
	if err := s.Get(u, `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *`,
		u.Username, u.Email, u.Password); err != nil {
		return fmt.Errorf("Error creating user: %w", err)
	}
	return nil
}

func (s *UserStore) UpdateUser(u *models.User) error {
	if err := s.Get(u, `UPDATE users SET username=$1, email=$2 WHERE id=$3 RETURNING *`,
		u.Username, u.Email, u.Id); err != nil {
		return fmt.Errorf("Error updating user: %w", err)
	}
	return nil
}

func (s *UserStore) DeleteUser(id int) error {
	if _, err := s.Exec(`DELETE FROM users WHERE id=$1`, id); err != nil {
		return fmt.Errorf("Error deleting user: %w", err)
	}
	return nil
}
