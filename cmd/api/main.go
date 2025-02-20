package main

import (
	"log"

	"github.com/ariefzainuri96/go-api-blogging/internal/store"
)

func main() {
	cfg := config{
		addr: ":8080",
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
