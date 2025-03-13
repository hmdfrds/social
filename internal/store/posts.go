package store

import (
	"context"
	"database/sql"

	"github.com/go-pg/pg"
)

type Post struct {
	ID        int64    `json:"id"`
	Context   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}
type PostsStore struct {
	db *sql.DB
}

func (s *PostsStore) Create(ctx context.Context, post *Post) error {
	query := `INSERT INTO post (content, title, users_id, tags) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`

	err := s.db.QueryRowContext(ctx, query, post.Context, post.Title, post.UserID, pg.Array(post.Tags)).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	return err
}
