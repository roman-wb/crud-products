package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Product_Const(t *testing.T) {
	assert.Equal(t, "The Name field is required.", ProductValidationNameRequired)
	assert.Equal(t, "The Price must be greater than or equal 0.", ProductValidationPriceGte)
}

func Test_Product_Validate(t *testing.T) {
	testCases := []struct {
		name         string
		product      Product
		wantLen      int
		wantMessages []string
	}{
		{
			name:         "valid with price 0",
			product:      Product{Name: "Name", Price: 0},
			wantLen:      0,
			wantMessages: []string{},
		},
		{
			name:         "valid with price gt 1",
			product:      Product{Name: "Name", Price: 1},
			wantLen:      0,
			wantMessages: []string{},
		},
		{
			name:         "invalid all",
			product:      Product{Name: "", Price: -1},
			wantLen:      2,
			wantMessages: []string{ProductValidationNameRequired, ProductValidationPriceGte},
		},
		{
			name:         "invalid name",
			product:      Product{Name: "", Price: 0},
			wantLen:      1,
			wantMessages: []string{ProductValidationNameRequired},
		},
		{
			name:         "invalid price",
			product:      Product{Name: "Name", Price: -1},
			wantLen:      1,
			wantMessages: []string{ProductValidationPriceGte},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			messages := tc.product.Validate()

			assert.Equal(t, tc.wantLen, len(messages))
			assert.Equal(t, tc.wantMessages, messages)
		})
	}
}

func Test_Product_Fill(t *testing.T) {
	type Params struct {
		Name  string  "json:\"name\""
		Price float64 "json:\"price\""
	}

	testCases := []struct {
		name       string
		wantParams Params
	}{
		{
			name: "Present params",
			wantParams: Params{
				Name:  "Name 1",
				Price: 100.99,
			},
		},
		{
			name:       "Blank params",
			wantParams: Params{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			product := &Product{}
			product.Fill(tc.wantParams)

			assert.Equal(t, tc.wantParams.Name, product.Name)
			assert.Equal(t, tc.wantParams.Price, product.Price)
		})
	}
}
