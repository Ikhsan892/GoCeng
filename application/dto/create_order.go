package dto

import "github.com/ikhsan892/goceng/application/domain"

type CreateOrder struct {
	Person_id   uint                    `json:"person_id"`
	TotalAmount float64                 `json:"total_amount"`
	Products    []domain.ProductRequest `json:"products"`
	PaymentType string                  `json:"payment_type"`
}
