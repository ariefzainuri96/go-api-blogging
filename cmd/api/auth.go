package main

import (
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("login"))
}

func AuthRouter() http.Handler {
	authRouter := http.NewServeMux()

	authRouter.Handle("POST /login", http.HandlerFunc(login))

	return http.StripPrefix("/v1/auth", authRouter)
}
