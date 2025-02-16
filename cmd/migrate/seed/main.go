package main

import (
	"log"

	"github.com/lucaiatropulus/social/internal/db"
	"github.com/lucaiatropulus/social/internal/env"
	"github.com/lucaiatropulus/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
