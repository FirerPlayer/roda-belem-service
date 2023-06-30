package main

import (
	"database/sql"
	"fmt"

	"github.com/firerplayer/roda-belem-service/configs"
	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/infra/filters"
	"github.com/firerplayer/roda-belem-service/internal/infra/repository"
	"github.com/go-chi/chi/v5"

	"github.com/firerplayer/roda-belem-service/internal/infra/web"
	"github.com/firerplayer/roda-belem-service/internal/infra/web/middlewares"
	"github.com/firerplayer/roda-belem-service/internal/infra/web/webserver"
	usecasePlaces "github.com/firerplayer/roda-belem-service/internal/usecase/places"
	usecaseReviews "github.com/firerplayer/roda-belem-service/internal/usecase/review"
	usecaseUser "github.com/firerplayer/roda-belem-service/internal/usecase/user"
	"googlemaps.github.io/maps"

	//import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	conn, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Configuração do webserver e rotas
	webServer := webserver.NewWebServer("8080")

	repoPlaceMySql := repository.NewPlaceRepositoryMySQL(conn)
	repoReviewMySql := repository.NewReviewRepositoryMySQL(conn)
	repoUserMySql := repository.NewUserRepositoryMysql(conn)
	million := 1000000
	cuckooFilter, err := filters.NewCuckooFilter(uint(10 * million))
	if err != nil {
		panic(err)
	}
	googleClient, err := maps.NewClient(maps.WithAPIKey(configs.GoogleMapsApiKey))
	if err != nil {
		panic(err)
	}

	setupPlacesHandlers(webServer, repoPlaceMySql, googleClient, cuckooFilter)
	setupReviewHandlers(webServer, repoReviewMySql)
	setupUsersHandlers(webServer, repoUserMySql, configs.JwtSecretKey)
	fmt.Println("Server running on port: ", webServer.WebServerPort)
	webServer.Start()
}

func setupPlacesHandlers(webserver *webserver.WebServer, repoPlaceMySql gateway.PlacesGateway, googleClient *maps.Client, cuckooFilter *filters.CuckooFilter) {
	// Places
	createPlaceUseCase := usecasePlaces.NewCreatePlaceUseCase(repoPlaceMySql, cuckooFilter)
	deletePlaceByIdUseCase := usecasePlaces.NewDeletePlaceByIDUseCase(repoPlaceMySql)
	findPlaceByIdUseCase := usecasePlaces.NewFindPlaceByIdUseCase(repoPlaceMySql)
	findNearbyPlacesUseCase := usecasePlaces.NewFindNearbyPlacesUseCase(repoPlaceMySql, googleClient, cuckooFilter)
	findPlacesByAccessibilityFeatureUseCase := usecasePlaces.NewFindPlacesByAccessibilityFeatureUseCase(repoPlaceMySql)
	updatePlaceByIdUseCase := usecasePlaces.NewUpdatePlaceByIDUseCase(repoPlaceMySql)
	saveFilterUseCase := usecasePlaces.NewSaveFilterUseCase(cuckooFilter)

	placesHandlers := web.NewWebPlacesHandlers(
		*createPlaceUseCase,
		*deletePlaceByIdUseCase,
		*findNearbyPlacesUseCase,
		*findPlaceByIdUseCase,
		*findPlacesByAccessibilityFeatureUseCase,
		*updatePlaceByIdUseCase,
		*saveFilterUseCase,
	)

	webserver.AddPostHandler("/places/create", placesHandlers.CreatePlace)
	webserver.AddDeleteHandler("/places/delete", placesHandlers.DeletePlaceByID)
	webserver.AddGetHandler("/places/find", placesHandlers.FindPlaceByID)
	webserver.AddGetHandler("/places/nearby", placesHandlers.FindNearbyPlaces)
	webserver.AddGetHandler("/places/accessibility", placesHandlers.FindPlacesByAccessibilityFeature)
	webserver.AddPostHandler("/places/update", placesHandlers.UpdatePLaceByID)
	webserver.AddGetHandler("/places/server/save-filter", placesHandlers.SaveFilter)

}

func setupReviewHandlers(webserver *webserver.WebServer, repoReviewMySql gateway.ReviewsGateway) {
	// Reviews
	createReviewUseCase := usecaseReviews.NewCreateReviewUseCase(repoReviewMySql)
	deleteReviewByIdUseCase := usecaseReviews.NewDeleteReviewByIDUseCase(repoReviewMySql)
	addAccessibilityFeaturesByReviewIdUseCase := usecaseReviews.NewAddAccessibilityFeaturesByReviewIdUseCase(repoReviewMySql)
	findReviewByIdUseCase := usecaseReviews.NewFindReviewByIdUseCase(repoReviewMySql)
	findReviewsByPlaceIdUseCase := usecaseReviews.NewFindReviewsByPlaceIDUseCase(repoReviewMySql)
	findReviewByUserIdUsecase := usecaseReviews.NewFindReviewsByUserIDUseCase(repoReviewMySql)
	updateReviewByIdUseCase := usecaseReviews.NewUpdateReviewByIDUseCase(repoReviewMySql)

	reviewsHandlers := web.NewWebReviewHandler(
		*createReviewUseCase,
		*findReviewByIdUseCase,
		*findReviewsByPlaceIdUseCase,
		*findReviewByUserIdUsecase,
		*updateReviewByIdUseCase,
		*deleteReviewByIdUseCase,
		*addAccessibilityFeaturesByReviewIdUseCase,
	)

	webserver.AddGroupRoute(func(r chi.Router) {
		r.Use(middlewares.AuthJWTMiddleware)
		r.Post("/reviews/create", reviewsHandlers.CreateReview)
		r.Delete("/reviews/delete", reviewsHandlers.DeleteReviewByID)
		r.Post("/reviews/update", reviewsHandlers.UpdateReviewByID)
		r.Get("/reviews/add-feature", reviewsHandlers.AddAccessibilityFeatureByReviewID)
	})

	webserver.AddGetHandler("/reviews/find-by-userid", reviewsHandlers.FindReviewsByUserID)
	webserver.AddGetHandler("/reviews/find", reviewsHandlers.FindReviewByID)
	webserver.AddGetHandler("/reviews/find-by-placeid", reviewsHandlers.FindReviewsByPlaceID)
}

func setupUsersHandlers(webserver *webserver.WebServer, repoUserMySql gateway.UsersGateway, jwtSecretKey string) {
	// Users
	createUserUseCase := usecaseUser.NewCreateUserUseCase(repoUserMySql)
	deleteUserByIdUseCase := usecaseUser.NewDeleteUserByIDUseCase(repoUserMySql)
	findUserByIdUseCase := usecaseUser.NewFindUserByIdUseCase(repoUserMySql)
	findUserByEmailUseCase := usecaseUser.NewFindUserByEmailUsecase(repoUserMySql)
	updateUserByIdUseCase := usecaseUser.NewUpdateUserUseCase(repoUserMySql)
	updateUserPointsByUserIdUseCase := usecaseUser.NewUpdateUserPointsByUserIDUseCase(repoUserMySql)
	addFavoriteUseCase := usecaseUser.NewAddFavoritesUseCase(repoUserMySql)
	listAllUsersUseCase := usecaseUser.NewListAllUsersUsecase(repoUserMySql)
	deleteFavoriteByUserIdUseCase := usecaseUser.NewDeleteFavoriteUseCase(repoUserMySql)
	authenticationJwtUseCase := usecaseUser.NewAuthenticationJwtUseCase(repoUserMySql, jwtSecretKey)

	usersHandlers := web.NewWebUserHandler(
		*createUserUseCase,
		*listAllUsersUseCase,
		*findUserByIdUseCase,
		*findUserByEmailUseCase,
		*updateUserByIdUseCase,
		*deleteUserByIdUseCase,
		*updateUserPointsByUserIdUseCase,
		*addFavoriteUseCase,
		*deleteFavoriteByUserIdUseCase,
		*authenticationJwtUseCase,
	)
	webserver.AddGroupRoute(func(r chi.Router) {
		r.Use(middlewares.AuthJWTMiddleware)
		r.Post("/users/update", usersHandlers.UpdateUserByID)
		r.Get("/users/update-points", usersHandlers.UpdateUserPointsByUserId)
		r.Delete("/users/delete-favorite", usersHandlers.DeleteFavoriteByUserIdAndPlaceId)
		r.Get("/users/add-favorite", usersHandlers.AddFavoriteByUserIdAndPlaceId)
	})

	webserver.AddPostHandler("/users/create", usersHandlers.CreateUser)
	webserver.AddDeleteHandler("/users/delete", usersHandlers.DeleteUserByID)
	webserver.AddGetHandler("/users/list", usersHandlers.ListAllUsers)
	webserver.AddGetHandler("/users/find", usersHandlers.FindUserByID)
	webserver.AddGetHandler("/users/find-by-email", usersHandlers.FindUserByEmail)
	webserver.AddPostHandler("/users/login", usersHandlers.LoginUser)
}
