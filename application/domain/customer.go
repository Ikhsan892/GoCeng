package domain

import (
	"encoding/json"
	"errors"
	"github.com/ikhsan892/goceng/application/value_objects"
	db "github.com/ikhsan892/goceng/sqlc"
)

type Customer struct {
	person  db.User
	address value_objects.Address
}

var (
	ErrInvalidName = errors.New("Invalid Name")
	ErrInvalidId   = errors.New("Invalid ID")
)

func NewCustomer(id uint, name string, address []byte) (Customer, error) {
	if id == 0 {
		return Customer{}, ErrInvalidId
	}

	if name == "" {
		return Customer{}, ErrInvalidName
	}

	var addressVo value_objects.Address

	err := json.Unmarshal(address, &addressVo)
	if err != nil {
		return Customer{}, err
	}

	return Customer{
		person: db.User{
			ID:       int64(id),
			Username: name,
			Address:  address,
		},
		address: addressVo,
	}, nil
}

func (c Customer) GetName() string {
	return c.person.Username
}

func (c Customer) GetId() int64 {
	return c.person.ID
}

func (c Customer) GetAddress() value_objects.Address {
	return c.address
}
