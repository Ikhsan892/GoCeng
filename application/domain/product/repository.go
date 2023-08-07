package product

import (
	"context"
	"github.com/ikhsan892/goceng/application/domain"
	"github.com/jackc/pgx/v5"
)

type ProductRepository interface {
	WithTx(tx pgx.Tx) *ProductPostgresRepository
	GetByIds(ctx context.Context, ids []int32) ([]domain.Product, error)
}
