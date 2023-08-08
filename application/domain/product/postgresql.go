package product

import (
	"context"
	"github.com/ikhsan892/goceng/application/domain"
	db "github.com/ikhsan892/goceng/sqlc"
)

type ProductPostgresRepository struct {
	sql *db.Queries
}

func NewProductPostgresRepository(tx *db.Queries) *ProductPostgresRepository {
	return &ProductPostgresRepository{sql: tx}
}

func (c ProductPostgresRepository) CreateProduct(ctx context.Context, arg db.CreateProductParams) error {
	return c.sql.CreateProduct(ctx, arg)
}

func (c ProductPostgresRepository) GetByIds(ctx context.Context, ids []int32) ([]domain.Product, error) {
	products, err := c.sql.GetProductsByIds(ctx, ids)

	var result []domain.Product

	if err != nil {
		return []domain.Product{{}}, err
	}

	for _, product := range products {
		var p *domain.Product
		p, err = domain.NewProduct(uint(product.ID), product.Name, float64(product.Price), uint(product.Stock))
		result = append(result, *p)
	}

	return result, nil
}
