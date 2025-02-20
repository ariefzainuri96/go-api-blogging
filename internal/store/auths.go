package store

import (
	"context"
	"database/sql"
)

type AuthsStore struct {
	db *sql.DB
}

func (s *AuthsStore) Login(ctx context.Context) error {
	return nil
}

func (s *AuthsStore) Create(ctx context.Context) error {
	return nil
}
