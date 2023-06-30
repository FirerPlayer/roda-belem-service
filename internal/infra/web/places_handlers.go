package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
	usecase "github.com/firerplayer/roda-belem-service/internal/usecase/places"
)

type WebPlacesHandlers struct {
	CreatePLaceUseCase                      usecase.CreatePlaceUseCase
	DeletePlaceByIDUseCase                  usecase.DeletePlaceByIDUseCase
	FindNearbyPlacesUseCase                 usecase.FindNearbyPlacesUseCase
	FindPlaceByIDUseCase                    usecase.FindPlaceByIDUseCase
	FindPlacesByAccessibilityFeatureUseCase usecase.FindPlacesByAccessibilityFeatureUseCase
	UpdatePLaceByIDUseCase                  usecase.UpdatePLaceByIDUseCase
	SaveFilterUseCase                       usecase.SaveFilterUseCase
}

func NewWebPlacesHandlers(
	createPLaceUseCase usecase.CreatePlaceUseCase,
	deletePlaceByIDUseCase usecase.DeletePlaceByIDUseCase,
	findNearbyPlacesUseCase usecase.FindNearbyPlacesUseCase,
	findPlaceByIDUseCase usecase.FindPlaceByIDUseCase,
	findPlacesByAccessibilityFeatureUseCase usecase.FindPlacesByAccessibilityFeatureUseCase,
	updatePLaceByIDUseCase usecase.UpdatePLaceByIDUseCase,
	saveFilterUseCase usecase.SaveFilterUseCase,
) *WebPlacesHandlers {
	return &WebPlacesHandlers{
		CreatePLaceUseCase:                      createPLaceUseCase,
		DeletePlaceByIDUseCase:                  deletePlaceByIDUseCase,
		FindNearbyPlacesUseCase:                 findNearbyPlacesUseCase,
		FindPlaceByIDUseCase:                    findPlaceByIDUseCase,
		FindPlacesByAccessibilityFeatureUseCase: findPlacesByAccessibilityFeatureUseCase,
		UpdatePLaceByIDUseCase:                  updatePLaceByIDUseCase,
		SaveFilterUseCase:                       saveFilterUseCase,
	}
}
func (h *WebPlacesHandlers) CreatePlace(w http.ResponseWriter, r *http.Request) {
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

func (h *WebPlacesHandlers) DeletePlaceByID(w http.ResponseWriter, r *http.Request) {
	var input dto.DeletePlaceByIDInputDTO
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	input.ID = id
	err := h.DeletePlaceByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *WebPlacesHandlers) FindNearbyPlaces(w http.ResponseWriter, r *http.Request) {
	var input dto.FindNearbyPlacesInputDTO
	var err error
	lat := r.URL.Query().Get("lat")
	lng := r.URL.Query().Get("lng")
	radius := r.URL.Query().Get("radius")
	isFromGoogle := r.URL.Query().Get("isFromGoogle")
	if lat == "" || lng == "" || radius == "" || isFromGoogle == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	input.Lat, err = strconv.ParseFloat(lat, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input.Lng, err = strconv.ParseFloat(lng, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input.Radius, err = strconv.ParseFloat(radius, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input.IsFromGoogle, err = strconv.ParseBool(isFromGoogle)
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

func (h *WebPlacesHandlers) FindPlaceByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	input := dto.FindPlaceByIDInputDTO{ID: id}
	output, err := h.FindPlaceByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *WebPlacesHandlers) FindPlacesByAccessibilityFeature(w http.ResponseWriter, r *http.Request) {
	var input dto.FindPlacesByAccessibilityFeatureInputDTO
	feat := r.URL.Query().Get("feature")
	if feat == "" {
		http.Error(w, "feature is required", http.StatusBadRequest)
		return
	}
	input.AccessibilityFeature = feat
	output, err := h.FindPlacesByAccessibilityFeatureUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *WebPlacesHandlers) UpdatePLaceByID(w http.ResponseWriter, r *http.Request) {
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

func (h *WebPlacesHandlers) SaveFilter(w http.ResponseWriter, r *http.Request) {
	err := h.SaveFilterUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
