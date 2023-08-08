package product

import (
	"context"
	"github.com/ikhsan892/goceng/application/domain"
	db "github.com/ikhsan892/goceng/sqlc"
)

type ProductRepository interface {
	GetByIds(ctx context.Context, ids []int32) ([]domain.Product, error)
	CreateProduct(ctx context.Context, arg db.CreateProductParams) error
}
