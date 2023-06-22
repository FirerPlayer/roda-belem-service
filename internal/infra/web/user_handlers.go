package web

import (
	"encoding/json"
	"net/http"

	"github.com/firerplayer/roda-belem-service/internal/usecase/dto"
	usecase "github.com/firerplayer/roda-belem-service/internal/usecase/user"
)

/*
type UsersGateway interface {
	CreateUser(ctx context.Context, user *entity.User) error
	ListAllUsers(ctx context.Context) ([]*entity.User, error)
	FindUserById(ctx context.Context, id string) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUserById(ctx context.Context, id string, user *entity.User) error
	DeleteUserById(ctx context.Context, id string) error
	UpdateUserPointsByUserId(ctx context.Context, userId string, points int) error
	AddFavoriteByUserIdAndPlaceId(ctx context.Context, userId string, placeId string) error
	DeleteFavoriteByUserIdAndPlaceId(ctx context.Context, userId string, placeId string) error
	// FindFavoritesByUserId returns a slice of strings containing the favorites of a user given their user ID.
	//
	// ctx is the context of the request.
	// userId is the ID of the user whose favorites are being searched.
	// It returns a slice of strings representing the favorites of the user and an error if any occurred.
	//
	FindFavoritesByUserId(ctx context.Context, userId string) ([]string, error)
}
*/

type WebUserHandler struct {
	CreateUserUseCase                       usecase.CreateUserUsecase
	ListAllUsersUseCase                     usecase.ListAllUsersUsecase
	FindUserByIdUseCase                     usecase.FindUserByIDUseCase
	FindUserByEmailUseCase                  usecase.FindUserByEmailUsecase
	UpdateUserByIDUseCase                   usecase.UpdateUserByIDUseCase
	DeleteUserByIDUseCase                   usecase.DeleteUserByIDUseCase
	UpdateUserPointsByUserIdUseCase         usecase.UpdateUserPointsByUserIdUseCase
	AddFavoriteByUserIdAndPlaceIdUseCase    usecase.AddFavoritesUseCase
	DeleteFavoriteByUserIdAndPlaceIdUseCase usecase.DeleteFavoriteUseCase
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
) *WebUserHandler {
	return &WebUserHandler{
		CreateUserUseCase:                       createUserUseCase,
		ListAllUsersUseCase:                     listAllUsersUseCase,
		FindUserByIdUseCase:                     findUserByIdUseCase,
		FindUserByEmailUseCase:                  findUserByEmailUseCase,
		UpdateUserByIDUseCase:                   updateUserByIDUseCase,
		DeleteUserByIDUseCase:                   deleteUserByIDUseCase,
		UpdateUserPointsByUserIdUseCase:         updateUserPointsByUserIdUseCase,
		AddFavoriteByUserIdAndPlaceIdUseCase:    addFavoriteByUserIdAndPlaceIdUseCase,
		DeleteFavoriteByUserIdAndPlaceIdUseCase: deleteFavoriteByUserIdAndPlaceIdUseCase,
	}
}

func (h *WebUserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
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

func (h *WebUserHandler) ListAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.ListAllUsersUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h *WebUserHandler) FindUserByID(w http.ResponseWriter, r *http.Request) {
	var input dto.FindUserByIDInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.FindUserByIdUseCase.Execute(r.Context(), input.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *WebUserHandler) FindUserByEmail(w http.ResponseWriter, r *http.Request) {
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

func (h *WebUserHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
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

func (h *WebUserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	var input dto.DeleteUserByIdInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.DeleteUserByIDUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *WebUserHandler) UpdateUserPointsByUserId(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateUserPointsByUserIDInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.UpdateUserPointsByUserIdUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *WebUserHandler) AddFavoriteByUserIdAndPlaceId(w http.ResponseWriter, r *http.Request) {
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

func (h *WebUserHandler) DeleteFavoriteByUserIdAndPlaceId(w http.ResponseWriter, r *http.Request) {
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
