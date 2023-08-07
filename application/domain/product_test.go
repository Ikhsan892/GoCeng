package domain_test

import (
	"github.com/ikhsan892/goceng/application/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct(t *testing.T) {
	p, _ := domain.NewProduct(1, "Indomie", 2000, 10)

	assert.NotNil(t, p)
}

func TestProductErrorIdIsLessThanOne(t *testing.T) {
	_, err := domain.NewProduct(0, "Indomie", 2000, 10)

	assert.Error(t, err, domain.ErrIdIsLessThanOne)
}
