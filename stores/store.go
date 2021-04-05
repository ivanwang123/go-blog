package stores

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	*UserStore
	*PostStore
	*CommentStore
	*LikeStore
}

func NewStore(dbString string) (*Store, error) {
	db, err := sqlx.Connect("pgx", dbString)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to database: %w", err)
	}

	return &Store{
		UserStore:    &UserStore{DB: db},
		PostStore:    &PostStore{DB: db},
		CommentStore: &CommentStore{DB: db},
		LikeStore:    &LikeStore{DB: db},
	}, nil
}
