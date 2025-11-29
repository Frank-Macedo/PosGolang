package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(productRepository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{
		repository: productRepository,
	}
}

// This product was not supposed to be return, whe should return a DTO instead.
// However, we will return it for now to keep the example simple.
func (u *ProductUseCase) GetProduct(id int) (Product, error) {
	return u.repository.GetProduct(id)
}
