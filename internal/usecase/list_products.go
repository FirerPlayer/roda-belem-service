package usecase

// import "github.com/firerplayer/hexagonal-arch-go/internal/repository"

// type ListProductsOutputDto struct {
// 	ID    string
// 	Name  string
// 	Price float64
// }

// type ListProductsUseCase struct {
// 	ProductRepository repository.ProductRepository
// }

// func NewListProductsUseCase(productRepository repository.ProductRepository) *ListProductsUseCase {
// 	return &ListProductsUseCase{
// 		ProductRepository: productRepository,
// 	}
// }

// func (uc *ListProductsUseCase) Execute() ([]*ListProductsOutputDto, error) {
// 	products, err := uc.ProductRepository.FindAll()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var output []*ListProductsOutputDto
// 	for _, product := range products {
// 		output = append(output, &ListProductsOutputDto{
// 			ID:    product.ID,
// 			Name:  product.Name,
// 			Price: product.Price,
// 		})
// 	}
// 	return output, nil
// }
