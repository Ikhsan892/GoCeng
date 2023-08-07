package domain_test

import (
	"github.com/ikhsan892/goceng/application/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderDomain(t *testing.T) {
	cust, _ := domain.NewCustomer(1, "ikhsan", []byte(`{"street": "", "province": "", "city": ""}`))
	product1, _ := domain.NewProduct(1, "Indomie", 3000, 1)
	product2, _ := domain.NewProduct(2, "Aqua", 2000, 2)
	zeroStock, _ := domain.NewProduct(3, "Wafer", 4000, 0)

	tableTests := []struct {
		amount      float64
		paymentType string
		customer    domain.Customer
		request     []domain.ProductRequest
		products    []domain.Product
		testCase    string
		wantErr     bool
		errData     error
	}{
		{
			amount:      0,
			paymentType: "",
			customer:    domain.Customer{},
			products:    nil,
			request:     []domain.ProductRequest{},
			testCase:    "Payment type is not set",
			wantErr:     true,
			errData:     domain.ErrInvalidPaymentType,
		},
		{
			amount:      0,
			paymentType: "transfer",
			customer:    domain.Customer{},
			request:     []domain.ProductRequest{},
			products:    nil,
			testCase:    "Customer not set",
			wantErr:     true,
			errData:     domain.ErrPersonCannotEmpty,
		},
		{
			amount:      1000,
			paymentType: "transfer",
			customer:    cust,
			products:    []domain.Product{},
			request:     []domain.ProductRequest{},
			testCase:    "Product Empty",
			wantErr:     true,
			errData:     domain.ErrProductCannotEmpty,
		},
		{
			amount:      1000,
			paymentType: "transfer",
			customer:    cust,
			products: []domain.Product{
				*product1,
				*product2,
			},
			request: []domain.ProductRequest{
				{
					Id:       1,
					Quantity: 2,
				},
				{
					Id:       2,
					Quantity: 5,
				},
			},
			testCase: "Amount incorrect",
			wantErr:  true,
			errData:  domain.ErrAmountIsNotCorrect,
		},
		{
			amount:      1000,
			paymentType: "transfer",
			customer:    cust,
			products: []domain.Product{
				*product1,
				*product2,
			},
			request: []domain.ProductRequest{
				{
					Id:       1,
					Quantity: 2,
				},
				{
					Id:       2,
					Quantity: 5,
				},
			},
			testCase: "Amount incorrect 2",
			wantErr:  true,
			errData:  domain.ErrAmountIsNotCorrect,
		},
		{
			amount:      5000,
			paymentType: "transfer",
			customer:    cust,
			products: []domain.Product{
				*product1,
				*product2,
				*zeroStock,
			},
			request: []domain.ProductRequest{
				{
					Id:       1,
					Quantity: 2,
				},
				{
					Id:       2,
					Quantity: 5,
				},
				{
					Id:       3,
					Quantity: 1,
				},
			},
			testCase: "Product has zero",
			wantErr:  true,
			errData:  domain.ErrProductHasZeroStock,
		},
		{
			amount:      5000,
			paymentType: "transfer",
			customer:    cust,
			products: []domain.Product{
				*product1,
				*product2,
			},
			request: []domain.ProductRequest{
				{
					Id:       1,
					Quantity: 1,
				},
				{
					Id:       2,
					Quantity: 1,
				},
			},
			testCase: "Correct all of it",
			wantErr:  false,
			errData:  nil,
		},
		{
			amount:      5000,
			paymentType: "transfer",
			customer:    cust,
			products: []domain.Product{
				*product1,
				*product2,
			},
			request: []domain.ProductRequest{
				{
					Id:       1,
					Quantity: 1,
				},
				{
					Id:       2,
					Quantity: 1,
				},
				{
					Id:       3,
					Quantity: 1,
				},
			},
			testCase: "Product and Request length not same",
			wantErr:  true,
			errData:  domain.ErrProductRequestedNotSameAsRecordResult,
		},
	}

	for _, test := range tableTests {
		t.Run(test.testCase, func(t *testing.T) {
			o := domain.NewOrder(test.amount, test.paymentType, test.request)
			orderDomain, err := o.SetCustomer(test.customer).
				SetItems(test.products).
				Build()

			if test.wantErr {
				assert.Error(t, err, test.errData)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, orderDomain.GetAmount(), test.amount)
				assert.Greater(t, len(orderDomain.GetProducts()), 0)
				assert.NotNil(t, orderDomain.GetCustomer())
				assert.Equal(t, orderDomain.GetPaymentType(), test.paymentType)
			}
		})
	}
}
