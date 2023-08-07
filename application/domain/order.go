package domain

import (
	"errors"
)

// aggregate root
type Order struct {
	id          uint
	customer    Customer
	items       []Product
	request     []ProductRequest
	amount      float64
	paymentType string
}

type ProductRequest struct {
	Id       uint
	Quantity uint
}

var (
	ErrProductCannotEmpty                    = errors.New("Product cannot be empty")
	ErrPersonCannotEmpty                     = errors.New("This order must have a person")
	ErrAmountIsZero                          = errors.New("Amount cannot be less than 1")
	ErrInvalidPaymentType                    = errors.New("Invalid payment type")
	ErrAmountIsNotCorrect                    = errors.New("Amount is not correct")
	ErrProductHasZeroStock                   = errors.New("Product has zero stock ")
	ErrProductRequestedNotSameAsRecordResult = errors.New("Product and Request data length is not match")
)

func NewOrder(amount float64, paymentType string, products []ProductRequest) *Order {
	return &Order{
		amount:      amount,
		paymentType: paymentType,
		request:     products,
	}
}

func (o *Order) SetItems(products []Product) *Order {
	o.items = products

	return o
}

func (o *Order) SetCustomer(customer Customer) *Order {
	o.customer = customer

	return o
}

func (o *Order) SetId(id uint) *Order {
	o.id = id

	return o
}

func (o Order) GetItems() []Product {
	return o.items
}

func (o Order) calculateTotalPrice() float64 {
	var total float64

	for i := 0; i < len(o.items); i++ {
		if o.items[i].id == o.request[i].Id {
			total += o.items[i].GetPrice() * float64(o.request[i].Quantity)
		}
	}

	return total
}

func (o Order) checkProductHasZeroStock() bool {

	for _, item := range o.items {
		if item.GetStock() < 1 {
			return true
		}
	}

	return false

}

func (o Order) GetId() uint {
	return o.id
}

func (o Order) GetCustomer() Customer {
	return o.customer
}
func (o Order) GetProducts() []Product {
	return o.items
}
func (o Order) GetAmount() float64 {
	return o.amount
}
func (o Order) GetPaymentType() string {
	return o.paymentType
}

func (o Order) Build() (Order, error) {
	if o.paymentType == "" {
		return Order{}, ErrInvalidPaymentType
	}

	if o.customer.GetName() == "" {
		return Order{}, ErrPersonCannotEmpty
	}

	if len(o.items) < 1 {
		return Order{}, ErrProductCannotEmpty
	}

	if o.amount < 1 {
		return Order{}, ErrAmountIsZero
	}

	if len(o.request) != len(o.items) {
		return Order{}, ErrProductRequestedNotSameAsRecordResult
	}

	if o.checkProductHasZeroStock() {
		return Order{}, ErrProductHasZeroStock
	}

	if o.amount != o.calculateTotalPrice() {
		return Order{}, ErrAmountIsNotCorrect
	}

	return o, nil
}
