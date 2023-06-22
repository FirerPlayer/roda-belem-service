package web

import (
	"encoding/json"
	"net/http"

	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
	usecase "github.com/firerplayer/roda-belem-service/internal/usecase/review"
)

type WebReviewHandler struct {
	CreateReviewUseCase                      usecase.CreateReviewUsecase
	FindReviewByIDUseCase                    usecase.FindReviewByIDUseCase
	FindReviewsByPlaceIDUseCase              usecase.FincdReviewsByPlaceIDUseCase
	FindReviewsByUserIDUseCase               usecase.FindReviewsByUserIDUseCase
	UpdateReviewByIDUseCase                  usecase.UpdateReviewByIDUseCase
	DeleteReviewByIDUseCase                  usecase.DeleteReviewByIDUseCase
	AddAccessibilityFeatureByReviewIDUseCase usecase.AddAccessibilityFeaturesByReviewIDUseCase
}

func NewWebReviewHandler(
	createReviewUseCase usecase.CreateReviewUsecase,
	findReviewByIDUseCase usecase.FindReviewByIDUseCase,
	findReviewsByPlaceIDUseCase usecase.FincdReviewsByPlaceIDUseCase,
	findReviewsByUserIDUseCase usecase.FindReviewsByUserIDUseCase,
	updateReviewByIDUseCase usecase.UpdateReviewByIDUseCase,
	deleteReviewByIDUseCase usecase.DeleteReviewByIDUseCase,
	addAccessibilityFeatureByReviewIDUseCase usecase.AddAccessibilityFeaturesByReviewIDUseCase,
) *WebReviewHandler {
	return &WebReviewHandler{
		CreateReviewUseCase:                      createReviewUseCase,
		FindReviewByIDUseCase:                    findReviewByIDUseCase,
		FindReviewsByPlaceIDUseCase:              findReviewsByPlaceIDUseCase,
		FindReviewsByUserIDUseCase:               findReviewsByUserIDUseCase,
		UpdateReviewByIDUseCase:                  updateReviewByIDUseCase,
		DeleteReviewByIDUseCase:                  deleteReviewByIDUseCase,
		AddAccessibilityFeatureByReviewIDUseCase: addAccessibilityFeatureByReviewIDUseCase,
	}
}

func (h *WebReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
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

func (h *WebReviewHandler) FindReviewByID(w http.ResponseWriter, r *http.Request) {
	var input dto.FindReviewByIDInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := h.FindReviewByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *WebReviewHandler) FindReviewsByPlaceID(w http.ResponseWriter, r *http.Request) {
	var input dto.FindReviewsByPlaceIDInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

func (h *WebReviewHandler) FindReviewsByUserID(w http.ResponseWriter, r *http.Request) {
	var input dto.FindReviewsByUserIDInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

func (h *WebReviewHandler) UpdateReviewByID(w http.ResponseWriter, r *http.Request) {
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

func (h *WebReviewHandler) DeleteReviewByID(w http.ResponseWriter, r *http.Request) {
	var input dto.DeleteReviewByIDInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.DeleteReviewByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *WebReviewHandler) AddAccessibilityFeatureByReviewID(w http.ResponseWriter, r *http.Request) {
	var input dto.AddAccessibilityFeaturesByReviewIDInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.AddAccessibilityFeatureByReviewIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
