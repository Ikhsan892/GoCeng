package dto

type CreateCustomerRequest struct {
	Username string
	Password string
	Email    string
	Address  []byte
}
