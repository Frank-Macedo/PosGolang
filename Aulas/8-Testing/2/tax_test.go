package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)

	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)

	println(err.Error())

	// if assert.Error(t, err, "amount must be lower than 0") {
	// 	assert.Contains(t, err.Error(), "must be lower than 0")
	// }

	assert.Equal(t, 0.0, tax)

}

func TestCalculateTaxAndSave(t *testing.T) {

	repository := &TaxRpositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)

}
