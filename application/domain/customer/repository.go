package customer

import (
	"context"
	"github.com/ikhsan892/goceng/application/domain"
	db "github.com/ikhsan892/goceng/sqlc"
)

type CustomerRepository interface {
	GetById(ctx context.Context, id uint) (domain.Customer, error)
	CreateCustomer(ctx context.Context, request db.CreateCustomerParams) error
}
