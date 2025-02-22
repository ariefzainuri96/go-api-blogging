package store

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	response "github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
	utils "github.com/ariefzainuri96/go-api-blogging/internal/utils"
)

type BlogsStore struct {
	db *sql.DB
}

func (s *BlogsStore) Create(ctx context.Context, body response.Blog) error {
	dir, err := os.Getwd() // Gets the working directory
	if err != nil {
		return err
	}

	path := filepath.Join(dir, "/internal/db/blogs.json")

	var blogs []response.Blog
	err = utils.LoadJsonData(path, &blogs)

	if err != nil {
		return err
	}

	blogs = append(blogs, body)

	if err := utils.SaveToJson(path, blogs); err != nil {
		return err
	}

	return nil
}

func (s *BlogsStore) GetAll(ctx context.Context) ([]response.Blog, error) {
	var blogs []response.Blog

	dir, err := os.Getwd() // Gets the working directory
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, "/internal/db/blogs.json")

	err = utils.LoadJsonData(path, &blogs)

	if err != nil {
		return nil, err
	}

	return blogs, nil
}
