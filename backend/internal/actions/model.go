package actions

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Action struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// GetAll récupère toutes les actions.
func GetAll(ctx context.Context, db *pgx.Conn) ([]Action, error) {
	rows, err := db.Query(ctx, "SELECT id, name, created_at FROM actions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actions []Action
	for rows.Next() {
		var a Action
		if err := rows.Scan(&a.ID, &a.Name, &a.CreatedAt); err != nil {
			return nil, err
		}
		actions = append(actions, a)
	}
	return actions, rows.Err()
}

// Insert ajoute une action en BDD et renvoie la nouvelle action.
func Insert(ctx context.Context, db *pgx.Conn, name string) (Action, error) {
	var a Action
	err := db.QueryRow(ctx,
		"INSERT INTO actions (name) VALUES ($1) RETURNING id, name, created_at",
		name).Scan(&a.ID, &a.Name, &a.CreatedAt)
	return a, err
}
