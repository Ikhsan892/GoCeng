package mock

import (
	"context"
	"github.com/ikhsan892/goceng/application/domain"
	"github.com/ikhsan892/goceng/application/domain/product"
	db "github.com/ikhsan892/goceng/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
)

type TestProductPostgresRepository struct {
	mock.Mock
}

func (c *TestProductPostgresRepository) toAggregate(product db.Product) (*domain.Product, error) {
	param := c.Called(product)

	return param.Get(0).(*domain.Product), param.Error(1)
}

func (c *TestProductPostgresRepository) WithTx(tx pgx.Tx) *product.ProductPostgresRepository {
	param := c.Called(tx)

	return param.Get(0).(*product.ProductPostgresRepository)
}

func (c *TestProductPostgresRepository) GetByIds(ctx context.Context, ids []int32) ([]domain.Product, error) {
	param := c.Called(ctx, ids)

	return param.Get(0).([]domain.Product), param.Error(1)
}
