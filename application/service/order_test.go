package service_test

import (
	"context"
	"github.com/ikhsan892/goceng/application/domain"
	"github.com/ikhsan892/goceng/application/dto"
	"github.com/ikhsan892/goceng/application/service"
	"github.com/ikhsan892/goceng/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type cust struct {
	id      uint
	name    string
	address []byte
}

func TestCreateOrder(t *testing.T) {
	testTable := []struct {
		request  dto.CreateOrder
		products []struct {
			id    uint
			name  string
			price float64
			stock uint
		}
		testcase   string
		customer   *cust
		custErr    error
		productErr error
		wantErr    bool
		errorData  error
	}{
		{
			testcase: "Success create order",
			products: []struct {
				id    uint
				name  string
				price float64
				stock uint
			}{
				{
					id:    1,
					name:  "Indomie",
					price: 3000,
					stock: 200,
				},
			},
			request: dto.CreateOrder{
				Person_id:   1,
				TotalAmount: 6000,
				Products: []domain.ProductRequest{
					{
						Id:       1,
						Quantity: 2,
					},
				},
				PaymentType: "transfer",
			},
			customer: &cust{
				id:      1,
				name:    "Ikhsan",
				address: []byte(`{"street": "", "province": "", "city": ""}`),
			},
			custErr:    nil,
			productErr: nil,
			wantErr:    false,
			errorData:  nil,
		},
		{
			testcase: "Success create order 2",
			products: []struct {
				id    uint
				name  string
				price float64
				stock uint
			}{
				{
					id:    1,
					name:  "Indomie",
					price: 3000,
					stock: 200,
				},
				{
					id:    2,
					name:  "Aqua",
					price: 5000,
					stock: 20,
				},
			},
			request: dto.CreateOrder{
				Person_id:   1,
				TotalAmount: 21000,
				Products: []domain.ProductRequest{
					{
						Id:       1,
						Quantity: 2,
					},
					{
						Id:       2,
						Quantity: 3,
					},
				},
				PaymentType: "transfer",
			},
			customer: &cust{
				id:      1,
				name:    "Ikhsan",
				address: []byte(`{"street": "", "province": "", "city": ""}`),
			},
			custErr:    nil,
			productErr: nil,
			wantErr:    false,
			errorData:  nil,
		},
		{
			testcase: "Create order with less amount",
			products: []struct {
				id    uint
				name  string
				price float64
				stock uint
			}{
				{
					id:    1,
					name:  "Indomie",
					price: 3000,
					stock: 200,
				},
				{
					id:    2,
					name:  "Aqua",
					price: 5000,
					stock: 20,
				},
			},
			request: dto.CreateOrder{
				Person_id:   1,
				TotalAmount: 20000,
				Products: []domain.ProductRequest{
					{
						Id:       1,
						Quantity: 2,
					},
					{
						Id:       2,
						Quantity: 3,
					},
				},
				PaymentType: "transfer",
			},
			customer: &cust{
				id:      1,
				name:    "Ikhsan",
				address: []byte(`{"street": "", "province": "", "city": ""}`),
			},
			custErr:    nil,
			productErr: nil,
			wantErr:    true,
			errorData:  domain.ErrAmountIsNotCorrect,
		},
		{
			testcase: "Create order but one product has empty stock",
			products: []struct {
				id    uint
				name  string
				price float64
				stock uint
			}{
				{
					id:    1,
					name:  "Indomie",
					price: 3000,
					stock: 200,
				},
				{
					id:    2,
					name:  "Aqua",
					price: 5000,
					stock: 0,
				},
			},
			request: dto.CreateOrder{
				Person_id:   1,
				TotalAmount: 20000,
				Products: []domain.ProductRequest{
					{
						Id:       1,
						Quantity: 2,
					},
					{
						Id:       2,
						Quantity: 3,
					},
				},
				PaymentType: "transfer",
			},
			customer: &cust{
				id:      1,
				name:    "Ikhsan",
				address: []byte(`{"street": "", "province": "", "city": ""}`),
			},
			custErr:    nil,
			productErr: nil,
			wantErr:    true,
			errorData:  domain.ErrProductHasZeroStock,
		},
	}

	ctx := context.Background()

	for _, s := range testTable {
		t.Run(s.testcase, func(t *testing.T) {
			customerRepo := new(mock.TestCustomerPostgresRepository)
			productRepo := new(mock.TestProductPostgresRepository)

			var user domain.Customer
			if s.customer != nil {
				user, _ = domain.NewCustomer(s.customer.id, s.customer.name, s.customer.address)
			}

			var products []domain.Product
			var ids []int32
			if len(s.products) > 0 {
				for _, p := range s.products {
					product, _ := domain.NewProduct(p.id, p.name, p.price, p.stock)
					products = append(products, *product)
					ids = append(ids, int32(p.id))
				}
			}

			customerRepo.On("GetById", ctx, uint(1)).Return(user, s.custErr)
			productRepo.On("GetByIds", ctx, ids).Return(products, s.productErr)

			os, _ := service.NewOrderService(ctx,
				service.WithPostgresRepository(
					ctx,
					customerRepo,
					productRepo,
				),
			)

			order, err := os.CreateOrder(s.request)

			if s.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.NotNil(t, order)
				assert.Nil(t, err)
			}

		})
	}

}
