package usecase

// import (
// 	"github.com/firerplayer/hexagonal-arch-go/internal/entity"
// 	"github.com/firerplayer/hexagonal-arch-go/internal/repository"
// )

// type CreateProductInputDto struct {
// 	Name  string  `json:"name"`
// 	Price float64 `json:"price"`
// }

// type CreateProductOutputDto struct {
// 	ID    string
// 	Name  string
// 	Price float64
// }

// type CreateProductUseCase struct {
// 	ProductRepository repository.ProductRepository
// }

// func NewCreateProductUseCase(productRepository repository.ProductRepository) *CreateProductUseCase {
// 	return &CreateProductUseCase{
// 		ProductRepository: productRepository,
// 	}
// }

// func (uc *CreateProductUseCase) Execute(input CreateProductInputDto) (*CreateProductOutputDto, error) {
// 	product := entity.NewProduct(input.Name, input.Price)
// 	err := uc.ProductRepository.Create(product)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &CreateProductOutputDto{
// 		ID:    product.ID,
// 		Name:  product.Name,
// 		Price: product.Price,
// 	}, nil
// }
