package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Blogs interface {
		Create(context.Context) error
	}
	Auths interface {
		Login(context.Context) error
		Create(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Blogs: &BlogsStore{db},
		Auths: &AuthsStore{db},
	}
}
