package customer

import (
	"context"
	"github.com/ikhsan892/goceng/application/domain"
	db "github.com/ikhsan892/goceng/sqlc"
)

type CustomerPostgresRepository struct {
	sql *db.Queries
}

func NewCustomerPostgresRepository(tx *db.Queries) CustomerRepository {
	return &CustomerPostgresRepository{sql: tx}
}

func (c CustomerPostgresRepository) toAggregate(user db.User) (domain.Customer, error) {
	return domain.NewCustomer(uint(user.ID), user.Username, user.Address)
}

func (c CustomerPostgresRepository) GetById(ctx context.Context, id uint) (domain.Customer, error) {
	user, err := c.sql.GetById(ctx, int64(id))
	if err != nil {
		return domain.Customer{}, err
	}

	return c.toAggregate(user)

}

func (c CustomerPostgresRepository) CreateCustomer(ctx context.Context, request db.CreateCustomerParams) error {
	err := c.sql.CreateCustomer(ctx, request)
	if err != nil {
		return err
	}

	return nil
}
