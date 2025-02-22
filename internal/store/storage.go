package store

import (
	"context"
	"database/sql"

	response "github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

type Storage struct {
	Blogs interface {
		Create(context.Context, response.Blog) error
		GetAll(context.Context) ([]response.Blog, error)
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
