package mock

import (
	"context"
	"github.com/ikhsan892/goceng/application/domain"
	"github.com/ikhsan892/goceng/application/domain/customer"
	db "github.com/ikhsan892/goceng/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
)

type TestCustomerPostgresRepository struct {
	mock.Mock
}

func (c *TestCustomerPostgresRepository) toAggregate(user db.User) (domain.Customer, error) {
	param := c.Called(user)

	return param.Get(0).(domain.Customer), param.Error(1)
}

func (c *TestCustomerPostgresRepository) WithTx(tx pgx.Tx) *customer.CustomerPostgresRepository {
	param := c.Called(tx)

	return param.Get(0).(*customer.CustomerPostgresRepository)
}

func (c *TestCustomerPostgresRepository) GetById(ctx context.Context, id uint) (domain.Customer, error) {
	param := c.Called(ctx, id)

	return param.Get(0).(domain.Customer), param.Error(1)

}

func (c *TestCustomerPostgresRepository) CreateCustomer(ctx context.Context, request db.CreateCustomerParams) error {
	param := c.Called(ctx, request)

	return param.Error(0)
}
