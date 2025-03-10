package store

import (
	"context"
	"database/sql"

	response "github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

type Storage struct {
	Blogs interface {
		GetAll(context.Context) ([]response.Blog, error)
		CreateWithDB(context.Context, *response.Blog) error
		GetById(context.Context, int64) (response.Blog, error)
		DeleteById(context.Context, int64) error
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
