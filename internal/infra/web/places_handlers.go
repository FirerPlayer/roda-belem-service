package web

import (
	"encoding/json"
	"net/http"

	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
	usecase "github.com/firerplayer/roda-belem-service/internal/usecase/places"
)

type WebPlacesHandler struct {
	CreatePLaceUseCase                      usecase.CreatePlaceUseCase
	DeletePlaceByIDUseCase                  usecase.DeletePlaceByIDUseCase
	FindNearbyPlacesUseCase                 usecase.FindNearbyPlacesUseCase
	FindPlaceByIDUseCase                    usecase.FindPlaceByIDUseCase
	FindPlacesByAccessibilityFeatureUseCase usecase.FindPlacesByAccessibilityFeatureUseCase
	UpdatePLaceByIDUseCase                  usecase.UpdatePLaceByIDUseCase
}

func NewWebServiceHandler(
	createPLaceUseCase usecase.CreatePlaceUseCase,
	deletePlaceByIDUseCase usecase.DeletePlaceByIDUseCase,
	findNearbyPlacesUseCase usecase.FindNearbyPlacesUseCase,
	findPlaceByIDUseCase usecase.FindPlaceByIDUseCase,
	findPlacesByAccessibilityFeatureUseCase usecase.FindPlacesByAccessibilityFeatureUseCase,
	updatePLaceByIDUseCase usecase.UpdatePLaceByIDUseCase,
) *WebPlacesHandler {
	return &WebPlacesHandler{
		CreatePLaceUseCase:                      createPLaceUseCase,
		DeletePlaceByIDUseCase:                  deletePlaceByIDUseCase,
		FindNearbyPlacesUseCase:                 findNearbyPlacesUseCase,
		FindPlaceByIDUseCase:                    findPlaceByIDUseCase,
		FindPlacesByAccessibilityFeatureUseCase: findPlacesByAccessibilityFeatureUseCase,
		UpdatePLaceByIDUseCase:                  updatePLaceByIDUseCase,
	}
}
func (h *WebPlacesHandler) CreatePlace(w http.ResponseWriter, r *http.Request) {
	var input dto.CreatePlaceInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.CreatePLaceUseCase.Execute(r.Context(), &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *WebPlacesHandler) DeletePlaceByID(w http.ResponseWriter, r *http.Request) {
	var input dto.DeletePlaceByIDInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.DeletePlaceByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *WebPlacesHandler) FindNearbyPlaces(w http.ResponseWriter, r *http.Request) {
	var input dto.FindNearbyPlacesInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := h.FindNearbyPlacesUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *WebPlacesHandler) FindPlaceByID(w http.ResponseWriter, r *http.Request) {
	var input dto.FindPlaceByIDInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := h.FindPlaceByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *WebPlacesHandler) FindPlacesByAccessibilityFeature(w http.ResponseWriter, r *http.Request) {
	var input dto.FindPlacesByAccessibilityFeatureInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := h.FindPlacesByAccessibilityFeatureUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *WebPlacesHandler) UpdatePLaceByID(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdatePlaceByIDInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.UpdatePLaceByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
