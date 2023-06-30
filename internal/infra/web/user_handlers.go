package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
	usecase "github.com/firerplayer/roda-belem-service/internal/usecase/user"
)

type WebUserHandlers struct {
	CreateUserUseCase                       usecase.CreateUserUsecase
	ListAllUsersUseCase                     usecase.ListAllUsersUsecase
	FindUserByIdUseCase                     usecase.FindUserByIDUseCase
	FindUserByEmailUseCase                  usecase.FindUserByEmailUsecase
	UpdateUserByIDUseCase                   usecase.UpdateUserByIDUseCase
	DeleteUserByIDUseCase                   usecase.DeleteUserByIDUseCase
	UpdateUserPointsByUserIdUseCase         usecase.UpdateUserPointsByUserIdUseCase
	AddFavoriteByUserIdAndPlaceIdUseCase    usecase.AddFavoritesUseCase
	DeleteFavoriteByUserIdAndPlaceIdUseCase usecase.DeleteFavoriteUseCase
	AuthenticateJwtUseCase                  usecase.AuthenticationJwtUseCase
}

func NewWebUserHandler(
	createUserUseCase usecase.CreateUserUsecase,
	listAllUsersUseCase usecase.ListAllUsersUsecase,
	findUserByIdUseCase usecase.FindUserByIDUseCase,
	findUserByEmailUseCase usecase.FindUserByEmailUsecase,
	updateUserByIDUseCase usecase.UpdateUserByIDUseCase,
	deleteUserByIDUseCase usecase.DeleteUserByIDUseCase,
	updateUserPointsByUserIdUseCase usecase.UpdateUserPointsByUserIdUseCase,
	addFavoriteByUserIdAndPlaceIdUseCase usecase.AddFavoritesUseCase,
	deleteFavoriteByUserIdAndPlaceIdUseCase usecase.DeleteFavoriteUseCase,
	authenticateJwtUseCase usecase.AuthenticationJwtUseCase,
) *WebUserHandlers {
	return &WebUserHandlers{
		CreateUserUseCase:                       createUserUseCase,
		ListAllUsersUseCase:                     listAllUsersUseCase,
		FindUserByIdUseCase:                     findUserByIdUseCase,
		FindUserByEmailUseCase:                  findUserByEmailUseCase,
		UpdateUserByIDUseCase:                   updateUserByIDUseCase,
		DeleteUserByIDUseCase:                   deleteUserByIDUseCase,
		UpdateUserPointsByUserIdUseCase:         updateUserPointsByUserIdUseCase,
		AddFavoriteByUserIdAndPlaceIdUseCase:    addFavoriteByUserIdAndPlaceIdUseCase,
		DeleteFavoriteByUserIdAndPlaceIdUseCase: deleteFavoriteByUserIdAndPlaceIdUseCase,
		AuthenticateJwtUseCase:                  authenticateJwtUseCase,
	}
}

func (h *WebUserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.CreateUserUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *WebUserHandlers) ListAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.ListAllUsersUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h *WebUserHandlers) FindUserByID(w http.ResponseWriter, r *http.Request) {
	// err := json.NewDecoder(r.Body).Decode(&input)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	input := dto.FindUserByIDInputDTO{UserId: id}
	user, err := h.FindUserByIdUseCase.Execute(r.Context(), input.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *WebUserHandlers) FindUserByEmail(w http.ResponseWriter, r *http.Request) {
	var input dto.FindUserByEmailInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.FindUserByEmailUseCase.Execute(r.Context(), input.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *WebUserHandlers) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateUserInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.UpdateUserByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *WebUserHandlers) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	input := dto.DeleteUserByIdInputDTO{UserId: id}
	err := h.DeleteUserByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *WebUserHandlers) UpdateUserPointsByUserId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("userId")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	points, err := strconv.ParseInt(r.URL.Query().Get("points"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input := dto.UpdateUserPointsByUserIDInputDTO{UserId: id, Points: int(points)}

	err = h.UpdateUserPointsByUserIdUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *WebUserHandlers) AddFavoriteByUserIdAndPlaceId(w http.ResponseWriter, r *http.Request) {
	var input dto.AddFavoritesInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.AddFavoriteByUserIdAndPlaceIdUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *WebUserHandlers) DeleteFavoriteByUserIdAndPlaceId(w http.ResponseWriter, r *http.Request) {
	var input dto.DeleteFavoriteByUserIdAndPlaceIdInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.DeleteFavoriteByUserIdAndPlaceIdUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *WebUserHandlers) LoginUser(w http.ResponseWriter, r *http.Request) {
	var input dto.AuthenticateJwtUserInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := h.AuthenticateJwtUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+output.Token)
	w.WriteHeader(http.StatusAccepted)

}
