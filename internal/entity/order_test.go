package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfGetsAnErrorIfIdIsBlank(t *testing.T) {
	order := Order{}

	assert.Error(t, order.Validate(), "id is required")
}

func TestIfGetsAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{
		ID: "any_id",
	}

	assert.Error(t, order.Validate(), "Price is muster be greater than 0")
}

func TestIfGetsAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{
		ID:    "any_id",
		Price: 10,
	}

	assert.Error(t, order.Validate(), "Tax is muster be greater than 0")
}

func TestIfCalculateFinalPriceIfPriceIsInvalid(t *testing.T) {
	order := Order{
		ID:    "any_id",
		Price: -10,
	}

	assert.Error(t, order.CalculateFinalPrice(), "Price is muster be greater than 0")
}

func TestIfCalculateFinalPriceIfTaxIsInvalid(t *testing.T) {
	order := Order{
		ID:    "any_id",
		Price: 10,
		Tax:   -5,
	}

	assert.Error(t, order.CalculateFinalPrice(), "Tax is muster be greater than 0")
}

func TestFinalPrice(t *testing.T) {
	order := Order{
		ID:    "any_id",
		Price: 10,
		Tax:   5,
	}

	assert.NoError(t, order.Validate())
	assert.Equal(t, "any_id", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 5.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 15.0, order.FinalPrice)
}
