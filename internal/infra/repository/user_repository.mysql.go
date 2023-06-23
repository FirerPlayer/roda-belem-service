package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
	db "github.com/firerplayer/roda-belem-service/internal/infra/mysql"
	"github.com/google/uuid"
)

type UserRepositoryMysql struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewUserRepositoryMysql(dbt *sql.DB) *UserRepositoryMysql {
	return &UserRepositoryMysql{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}

func (r *UserRepositoryMysql) CreateUser(ctx context.Context, user *entity.User) error {
	params := db.CreateUserParams{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Points:    int32(user.Points),
		CreatedAt: sql.NullTime{Time: user.CreatedAt},
		UpdatedAt: sql.NullTime{Time: user.UpdatedAt},
	}
	missions, err := json.Marshal(user.Missions)
	if err != nil {
		return err
	}
	params.Missions = missions
	return r.Queries.CreateUser(ctx, params)
}

func (r *UserRepositoryMysql) ListAllUsers(ctx context.Context) ([]*entity.User, error) {
	usersDb, err := r.Queries.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	var users []*entity.User
	for _, userDb := range usersDb {
		var user entity.User
		err := HydrateUser(userDb, &user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *UserRepositoryMysql) FindUserById(ctx context.Context, id string) (*entity.User, error) {
	userDb, err := r.Queries.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	var user entity.User
	err = HydrateUser(userDb, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryMysql) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	userDb, err := r.Queries.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	var user entity.User
	err = HydrateUser(userDb, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryMysql) UpdateUserById(ctx context.Context, id string, user *entity.User) error {
	params := db.UpdateUserByIdParams{
		ID:        id,
		Email:     user.Email,
		Avatar:    sql.NullString{String: string(user.Avatar)},
		Username:  user.Username,
		Password:  user.Password,
		Points:    int32(user.Points),
		UpdatedAt: sql.NullTime{Time: user.UpdatedAt},
	}
	missions, err := json.Marshal(user.Missions)
	if err != nil {
		return err
	}
	params.Missions = missions
	return r.Queries.UpdateUserById(ctx, params)
}

func (r *UserRepositoryMysql) DeleteUserById(ctx context.Context, id string) error {
	return r.Queries.DeleteUserById(ctx, id)
}

func (r *UserRepositoryMysql) UpdateUserPointsByUserId(ctx context.Context, userId string, points int) error {
	params := db.UpdateUserPointsByUserIdParams{
		ID:     userId,
		Points: int32(points),
	}

	return r.Queries.UpdateUserPointsByUserId(ctx, params)
}

func (r *UserRepositoryMysql) AddFavoriteByUserIdAndPlaceId(ctx context.Context, userId string, placeId string) error {
	params := db.CreateFavoriteParams{
		UserID:  sql.NullString{String: userId},
		PlaceID: sql.NullString{String: placeId},
	}

	return r.Queries.CreateFavorite(ctx, params)
}

func (r *UserRepositoryMysql) DeleteFavoriteByUserIdAndPlaceId(ctx context.Context, userId string, placeId string) error {
	params := db.DeleteFavoriteByUserIdAndPlaceIdParams{
		UserID:  sql.NullString{String: userId},
		PlaceID: sql.NullString{String: placeId},
	}
	return r.Queries.DeleteFavoriteByUserIdAndPlaceId(ctx, params)
}

func (r *UserRepositoryMysql) FindFavoritesByUserId(ctx context.Context, userId string) ([]string, error) {
	favoritesDb, err := r.Queries.FindFavoritesByUserId(ctx, sql.NullString{String: userId})
	if err != nil {
		return nil, err
	}
	var favorites []string
	for _, favoriteDb := range favoritesDb {
		favorites = append(favorites, favoriteDb.String)
	}
	return favorites, nil
}

// HydrateUser hydrates a user entity with data from a userDb record.
//
// userDb: The database record to hydrate from.
// user: The user entity to hydrate.
// error: Returns an error if the JSON unmarshal fails.
func HydrateUser(userDb db.User, user *entity.User) error {
	user.ID = uuid.MustParse(userDb.ID)
	user.Email = userDb.Email
	user.Avatar = []byte(userDb.Avatar.String)
	user.Username = userDb.Username
	user.Password = userDb.Password
	user.Points = int(userDb.Points)
	user.CreatedAt = userDb.CreatedAt.Time
	user.UpdatedAt = userDb.UpdatedAt.Time
	err := json.Unmarshal(userDb.Missions, &user.Missions)
	if err != nil {
		return err
	}
	return nil
}
