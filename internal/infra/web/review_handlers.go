package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
	usecase "github.com/firerplayer/roda-belem-service/internal/usecase/review"
)

type WebReviewHandlers struct {
	CreateReviewUseCase                       usecase.CreateReviewUsecase
	FindReviewByIDUseCase                     usecase.FindReviewByIDUseCase
	FindReviewsByPlaceIDUseCase               usecase.FincdReviewsByPlaceIDUseCase
	FindReviewsByUserIDUseCase                usecase.FindReviewsByUserIDUseCase
	UpdateReviewByIDUseCase                   usecase.UpdateReviewByIDUseCase
	DeleteReviewByIDUseCase                   usecase.DeleteReviewByIDUseCase
	AddAccessibilityFeaturesByReviewIDUseCase usecase.AddAccessibilityFeaturesByReviewIDUseCase
}

func NewWebReviewHandler(
	createReviewUseCase usecase.CreateReviewUsecase,
	findReviewByIDUseCase usecase.FindReviewByIDUseCase,
	findReviewsByPlaceIDUseCase usecase.FincdReviewsByPlaceIDUseCase,
	findReviewsByUserIDUseCase usecase.FindReviewsByUserIDUseCase,
	updateReviewByIDUseCase usecase.UpdateReviewByIDUseCase,
	deleteReviewByIDUseCase usecase.DeleteReviewByIDUseCase,
	addAccessibilityFeaturesByReviewIDUseCase usecase.AddAccessibilityFeaturesByReviewIDUseCase,
) *WebReviewHandlers {
	return &WebReviewHandlers{
		CreateReviewUseCase:                       createReviewUseCase,
		FindReviewByIDUseCase:                     findReviewByIDUseCase,
		FindReviewsByPlaceIDUseCase:               findReviewsByPlaceIDUseCase,
		FindReviewsByUserIDUseCase:                findReviewsByUserIDUseCase,
		UpdateReviewByIDUseCase:                   updateReviewByIDUseCase,
		DeleteReviewByIDUseCase:                   deleteReviewByIDUseCase,
		AddAccessibilityFeaturesByReviewIDUseCase: addAccessibilityFeaturesByReviewIDUseCase,
	}
}

func (h *WebReviewHandlers) CreateReview(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateReviewInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.CreateReviewUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *WebReviewHandlers) FindReviewByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	input := dto.FindReviewByIDInputDTO{Id: id}
	output, err := h.FindReviewByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *WebReviewHandlers) FindReviewsByPlaceID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("placeId")
	if id == "" {
		http.Error(w, "placeId is required", http.StatusBadRequest)
		return
	}
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	if err != nil {
		limit = 30
	}
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	if err != nil {
		offset = 0
	}
	input := dto.FindReviewsByPlaceIDInputDTO{
		PlaceID: id,
		Limit:   int(limit),
		Offset:  int(offset),
	}
	output, err := h.FindReviewsByPlaceIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *WebReviewHandlers) FindReviewsByUserID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	if err != nil {
		limit = 30
	}
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	if err != nil {
		offset = 0
	}
	input := dto.FindReviewsByUserIDInputDTO{
		UserID: id,
		Limit:  int(limit),
		Offset: int(offset),
	}
	output, err := h.FindReviewsByUserIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *WebReviewHandlers) UpdateReviewByID(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateReviewByIDInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.UpdateReviewByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *WebReviewHandlers) DeleteReviewByID(w http.ResponseWriter, r *http.Request) {
	var input dto.DeleteReviewByIDInputDTO
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	input.ID = id
	err := h.DeleteReviewByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *WebReviewHandlers) AddAccessibilityFeatureByReviewID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	feat := r.URL.Query().Get("feature")
	if feat == "" {
		http.Error(w, "feature is required", http.StatusBadRequest)
		return
	}
	input := dto.AddAccessibilityFeaturesByReviewIDInputDTO{
		ReviewID: id,
		Features: feat,
	}
	err := h.AddAccessibilityFeaturesByReviewIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
