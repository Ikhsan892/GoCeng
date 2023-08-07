package customer

import (
	"context"
	"github.com/ikhsan892/goceng/application/domain"
	db "github.com/ikhsan892/goceng/sqlc"
	"github.com/jackc/pgx/v5"
)

type CustomerRepository interface {
	WithTx(tx pgx.Tx) *CustomerPostgresRepository
	GetById(ctx context.Context, id uint) (domain.Customer, error)
	CreateCustomer(ctx context.Context, request db.CreateCustomerParams) error
}
