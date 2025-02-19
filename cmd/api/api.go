package main

import (
	"log"
	"net/http"
	"time"

	middleware "github.com/ariefzainuri96/go-api-blogging/cmd/api/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/v1/blog/", BlogRouter())

	mux.Handle("/v1/auth/", AuthRouter())

	return mux
}

func (app *application) run(mux *http.ServeMux) error {

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      stack(mux),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	log.Printf("Server has started on %s", app.config.addr)

	return srv.ListenAndServe()
}
