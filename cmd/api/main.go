package main

import (
	"fmt"
	"log"
	"social/internal/db"
	"social/internal/env"
	"social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt("MAX_OPEN_CONNS", 32),
			maxIdleConns: env.GetInt("MAX_IDLE_CONNS", 32),
			maxIdleTime: env.GetString("MAX_IDLE_TIME", "10m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	
	if err != nil {
		fmt.Print(err)
	}

	defer db.Close()
	log.Println("database is connected")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store: store,
	}


	mux := app.mount()
	log.Fatal(app.run(mux))
}