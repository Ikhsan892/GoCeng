package service

import (
	"context"
	"github.com/ikhsan892/goceng/application/domain"
	"github.com/ikhsan892/goceng/application/domain/customer"
	"github.com/ikhsan892/goceng/application/domain/product"
	"github.com/ikhsan892/goceng/application/dto"
)

type OrderService struct {
	customer customer.CustomerRepository
	product  product.ProductRepository
	context  context.Context
}

type OrderServiceConfiguration func(os *OrderService) error

func WithPostgresRepository(ctx context.Context, customerRepo customer.CustomerRepository, productRepo product.ProductRepository) func(os *OrderService) error {
	return func(os *OrderService) error {
		os.customer = customerRepo
		os.product = productRepo
		return nil
	}
}

func NewOrderService(ctx context.Context, cfgs ...OrderServiceConfiguration) (*OrderService, error) {
	os := &OrderService{}
	os.context = ctx

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func (os OrderService) CreateOrder(request dto.CreateOrder) (domain.Order, error) {
	cust, err := os.customer.GetById(os.context, request.Person_id)
	if err != nil {
		return domain.Order{}, err
	}

	var productIds []int32
	for _, itemRequest := range request.Products {
		productIds = append(productIds, int32(itemRequest.Id))
	}

	products, err := os.product.GetByIds(os.context, productIds)
	if err != nil {
		return domain.Order{}, err
	}

	order := domain.NewOrder(request.TotalAmount, request.PaymentType, request.Products).
		SetCustomer(cust).
		SetItems(products)

	return order.Build()
}
