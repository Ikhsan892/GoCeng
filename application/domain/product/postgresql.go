package product

import (
	"context"
	"github.com/ikhsan892/goceng/application/domain"
	db "github.com/ikhsan892/goceng/sqlc"
	"github.com/jackc/pgx/v5"
)

type ProductPostgresRepository struct {
	sql *db.Queries
}

func NewProductPostgresRepository(tx *db.Queries) *ProductPostgresRepository {
	return &ProductPostgresRepository{sql: tx}
}

func (c ProductPostgresRepository) toAggregate(product db.Product) (*domain.Product, error) {
	var price float64
	var stock uint
	err := product.Price.Scan(&price)
	err = product.Stock.Scan(&stock)
	if err != nil {
		return nil, err
	}

	return domain.NewProduct(uint(product.ID), product.Name, price, stock)
}

func (c *ProductPostgresRepository) WithTx(tx pgx.Tx) *ProductPostgresRepository {
	c.sql.WithTx(tx)

	return c
}

func (c ProductPostgresRepository) GetByIds(ctx context.Context, ids []int32) ([]domain.Product, error) {
	products, err := c.sql.GetProductsByIds(ctx, ids)

	var result []domain.Product

	if err != nil {
		return []domain.Product{{}}, err
	}

	for _, product := range products {
		var p *domain.Product
		p, err = c.toAggregate(product)
		result = append(result, *p)
	}

	return result, nil
}
