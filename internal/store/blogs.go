package store

import (
	"context"
	"database/sql"
)

type BlogsStore struct {
	db *sql.DB
}

func (s *BlogsStore) Create(ctx context.Context) error {
	return nil
}
