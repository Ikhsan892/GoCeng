package domain

import "errors"

type Product struct {
	id    uint
	name  string
	price float64
	stock uint
}

var (
	ErrIdIsLessThanOne = errors.New("Id cannot less than one")
)

func NewProduct(id uint, name string, price float64, stock uint) (*Product, error) {
	if id == 0 {
		return nil, ErrIdIsLessThanOne
	}

	return &Product{
		id:    id,
		name:  name,
		price: price,
		stock: stock,
	}, nil
}

func (p Product) GetId() uint {
	return p.id
}

func (p Product) GetName() string {
	return p.name
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p Product) GetStock() uint {
	return p.stock
}
