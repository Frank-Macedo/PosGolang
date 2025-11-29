package usecase

import (
	"github.com/Frank-Macedo/20-cleanArch/internal/entity"
)

type GetOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *GetOrdersUseCase) Execute() ([]OrderOutputDTO, error) {

	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		output = append(output, dto)
	}

	return output, nil
}
