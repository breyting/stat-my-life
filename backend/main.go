package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

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

	http.HandleFunc("/actions/", func(w http.ResponseWriter, r *http.Request) {
		// Pour /actions/{id}/events
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		if len(pathParts) == 3 && pathParts[0] == "actions" && pathParts[2] == "events" {
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, "Invalid action ID", http.StatusBadRequest)
				return
			}

			switch r.Method {
			case http.MethodGet:
				h.GetOccurrences(w, r, id)
			case http.MethodPost:
				h.CreateOccurrence(w, r, id)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
			return
		}

		http.NotFound(w, r)
	})

	log.Println("API listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
