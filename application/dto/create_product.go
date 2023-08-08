package dto

type CreateProduct struct {
	Name  string  `json:"name"`
	Stock uint    `json:"stock"`
	Price float64 `json:"price"`
}
