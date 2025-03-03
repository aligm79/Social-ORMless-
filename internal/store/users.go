package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID int64 `json:"id"`
	Username int64 `json:"username"`
	Email int64 `json:"email"`
	Password int64 `json:"-"`
	CreatedAt int64 `json:"created_at"`
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id, created_at
	`

	err := s.db.QueryRowContext(
		ctx, 
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}