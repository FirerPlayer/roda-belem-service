package web

func denis() {

}

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/firerplayer/hexagonal-arch-go/internal/usecase"
// )

// type ProductHandlers struct {
// 	CreateProductUsecase *usecase.CreateProductUseCase
// 	ListProductUseCase   *usecase.ListProductsUseCase
// }

// func NewProductHandlers(createProductUsecase *usecase.CreateProductUseCase, listProductsUseCase *usecase.ListProductsUseCase) *ProductHandlers {
// 	return &ProductHandlers{
// 		CreateProductUsecase: createProductUsecase,
// 		ListProductUseCase:   listProductsUseCase,
// 	}
// }

// func (h *ProductHandlers) CreateProduct(w http.ResponseWriter, r *http.Request) {
// 	var input usecase.CreateProductInputDto
// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	output, err := h.CreateProductUsecase.Execute(input)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(output)
// }

// func (h *ProductHandlers) ListProducts(w http.ResponseWriter, r *http.Request) {
// 	output, err := h.ListProductUseCase.Execute()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(output)
// }
