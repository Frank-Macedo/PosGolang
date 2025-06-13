package tax

import (
	"github.com/stretchr/testify/mock"
)

type TaxRpositoryMock struct {
	mock.Mock
}

func (m *TaxRpositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)

}
