package infrastructures

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func NewPostgreSQL(ctx context.Context, Url string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, Url)
	if err != nil {
		log.Fatal("Unable to connect to database")
		return nil, err
	}
	log.Println("Database successfully connected")
	return conn, nil
}
