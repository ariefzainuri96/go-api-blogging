package main

import (
	"fmt"
	"net/http"
)

func postBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("post blog"))
}

func getBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("get blog"))
}

func getBlogById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get blog by id"))
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("delete blog"))
}

func putBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("put blog"))
}

func patchBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("patch blog"))
}

func blogComments(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("blog comments %s", id)))
}

func BlogRouter() http.Handler {
	blogRouter := http.NewServeMux()

	blogRouter.Handle("POST /", http.HandlerFunc(postBlog))
	blogRouter.Handle("GET /", http.HandlerFunc(getBlog))
	blogRouter.Handle("GET /{id}/comments", http.HandlerFunc(blogComments))
	blogRouter.Handle("GET /{id}", http.HandlerFunc(getBlogById))
	blogRouter.Handle("DELETE /{id}", http.HandlerFunc(deleteBlog))
	blogRouter.Handle("PUT /{id}", http.HandlerFunc(putBlog))
	blogRouter.Handle("PATCH /{id}", http.HandlerFunc(patchBlog))

	return http.StripPrefix("/v1/blog", blogRouter)
}
