package service

import (
	"context"
	"github.com/ikhsan892/goceng/application/domain/product"
	"github.com/ikhsan892/goceng/application/dto"
	db "github.com/ikhsan892/goceng/sqlc"
)

type ProductService struct {
	ctx     context.Context
	product product.ProductRepository
}

type ProductServiceConfiguration func(ps *ProductService) error

func NewWithProductPostgreRepository(pr *product.ProductPostgresRepository) func(ps *ProductService) error {
	return func(ps *ProductService) error {
		ps.product = pr

		return nil
	}
}

func NewProductService(ctx context.Context, cfg ...ProductServiceConfiguration) (*ProductService, error) {
	pr := &ProductService{ctx: ctx}

	for _, configuration := range cfg {
		if err := configuration(pr); err != nil {
			return nil, err
		}
	}

	return pr, nil
}

func (p ProductService) CreateProduct(createProductDto dto.CreateProduct) error {
	err := p.product.CreateProduct(p.ctx, db.CreateProductParams{
		Name:    createProductDto.Name,
		Column2: createProductDto.Price,
		Column3: int32(createProductDto.Stock),
	})
	if err != nil {
		return err
	}

	return nil

}
