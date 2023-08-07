package domain_test

import (
	"github.com/ikhsan892/goceng/application/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomerDomain(t *testing.T) {
	address := []byte(`{"street": "", "province": "", "city": ""}`)

	tableTest := []struct {
		id       uint
		name     string
		wantErr  bool
		errData  error
		testcase string
	}{
		{
			id:       1,
			name:     "",
			wantErr:  true,
			errData:  domain.ErrInvalidName,
			testcase: "GIVEN name is empty",
		},
		{
			id:       0,
			name:     "",
			wantErr:  true,
			errData:  domain.ErrInvalidId,
			testcase: "GIVEN id is 0",
		},
		{
			id:       1,
			name:     "ikhsan",
			wantErr:  false,
			errData:  domain.ErrInvalidId,
			testcase: "GIVEN name and id is valid and object is not nil",
		},
	}

	for _, s := range tableTest {
		t.Run(s.testcase, func(t *testing.T) {
			data, err := domain.NewCustomer(s.id, s.name, address)
			if s.wantErr {
				assert.Error(t, err, s.errData)
			} else {
				assert.NotNil(t, data)
			}
		})
	}
}
