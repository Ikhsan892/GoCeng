package order

import (
	"context"
	db "github.com/ikhsan892/goceng/sqlc"
	"github.com/jackc/pgx/v5"
)

type OrderPostgresRepository interface {
	WithTx(tx pgx.Tx) *db.Queries
	SaveOrder(ctx context.Context)
}
