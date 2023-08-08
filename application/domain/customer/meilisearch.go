package customer

import (
	"context"
	"github.com/ikhsan892/goceng/application/domain"
	db "github.com/ikhsan892/goceng/sqlc"
)

type CustomerMeilisearchRepository struct {
}

func (c CustomerMeilisearchRepository) GetById(ctx context.Context, id uint) (domain.Customer, error) {
	return domain.Customer{}, nil
}

func (c CustomerMeilisearchRepository) CreateCustomer(ctx context.Context, request db.CreateCustomerParams) error {
	return nil
}
