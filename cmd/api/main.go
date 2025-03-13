package main

import (
	"log"

	"github.com/hmdfrds/social/internal/env"
	"github.com/hmdfrds/social/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using existing environment variables")
	}
	cfg := config{
		addr: env.GetString("ADDR", ":8081"),
	}
	store := store.NewStorage(nil)
	app := &application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
