package main

import (
	"context"
	"log"
	"net/http"

	"github.com/breyting/statmylife/internal/actions"
	"github.com/jackc/pgx/v5"
)

func main() {
	connStr := "postgres://sml:smlpass@localhost:5432/statmylife"
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close(context.Background())

	h := &actions.Handler{DB: db}

	http.HandleFunc("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(`{"message":"pong"}`))
	})

	http.HandleFunc("/actions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetActions(w, r)
		case http.MethodPost:
			h.CreateAction(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("API listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
