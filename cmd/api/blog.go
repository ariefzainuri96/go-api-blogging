package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

func (app *application) postBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("post blog"))

	app.store.Blogs.Create(context.Background())
}

func (app *application) getBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get blog"))
}

func (app *application) getBlogById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	baseResp := response.BaseResponse{}

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "invalid id"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success"
	blogResp, _ := response.BlogResponse{
		BaseResponse: baseResp,
		Blog: response.Blog{
			ID:    int64(id),
			Title: "test",
			Body:  "test",
		},
	}.MarshalBlogResponse()
	w.WriteHeader(http.StatusOK)
	w.Write(blogResp)
}

func (app *application) deleteBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("delete blog"))
}

func (app *application) putBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("put blog"))
}

func (app *application) patchBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("patch blog"))
}

func (app *application) blogComments(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("blog comments %s", id)))
}

func (app *application) BlogRouter() *http.ServeMux {
	blogRouter := http.NewServeMux()

	blogRouter.HandleFunc("POST /", app.postBlog)
	blogRouter.HandleFunc("GET /", app.getBlog)
	blogRouter.HandleFunc("GET /{id}/comments", app.blogComments)
	blogRouter.HandleFunc("GET /{id}", app.getBlogById)
	blogRouter.HandleFunc("DELETE /{id}", app.deleteBlog)
	blogRouter.HandleFunc("PUT /{id}", app.putBlog)
	blogRouter.HandleFunc("PATCH /{id}", app.patchBlog)

	// Catch-all route for undefined paths
	blogRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	return blogRouter
}
