package main

import (
	"database/sql"
	"fmt"

	"github.com/firerplayer/roda-belem-service/configs"
	"github.com/firerplayer/roda-belem-service/internal/domain/gateway"
	"github.com/firerplayer/roda-belem-service/internal/infra/blooms"
	"github.com/firerplayer/roda-belem-service/internal/infra/repository"
	"github.com/firerplayer/roda-belem-service/internal/infra/web"
	"github.com/firerplayer/roda-belem-service/internal/infra/web/webserver"
	usecasePlaces "github.com/firerplayer/roda-belem-service/internal/usecase/places"
	usecaseReviews "github.com/firerplayer/roda-belem-service/internal/usecase/review"
	usecaseUser "github.com/firerplayer/roda-belem-service/internal/usecase/user"
	"googlemaps.github.io/maps"

	//import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// import (
// 	"database/sql"
// 	"net/http"

// 	"github.com/firerplayer/hexagonal-arch-go/internal/infra/repository"
// 	"github.com/firerplayer/hexagonal-arch-go/internal/infra/web"
// 	"github.com/firerplayer/hexagonal-arch-go/internal/usecase"
// 	"github.com/go-chi/chi/v5"

// 	_ "github.com/go-sql-driver/mysql"
// )

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
	bloomFilter := blooms.NewBloomFilter(100000000)
	googleClient, err := maps.NewClient(maps.WithAPIKey(configs.GoogleMapsApiKey))
	if err != nil {
		panic(err)
	}

	setupPlacesHandlers(webServer, repoPlaceMySql, googleClient, bloomFilter)
	setupReviewHandlers(webServer, repoReviewMySql)
	setupUsersHandlers(webServer, repoUserMySql)
	fmt.Println("Server running on port: ", webServer.WebServerPort)
	webServer.Start()
}

func setupPlacesHandlers(webserver *webserver.WebServer, repoPlaceMySql gateway.PlacesGateway, googleClient *maps.Client, bloomFilter *blooms.BloomFilter) {
	// Places
	createPlaceUseCase := usecasePlaces.NewCreatePlaceUseCase(repoPlaceMySql, bloomFilter)
	deletePlaceByIdUseCase := usecasePlaces.NewDeletePlaceByIDUseCase(repoPlaceMySql)
	findPlaceByIdUseCase := usecasePlaces.NewFindPlaceByIdUseCase(repoPlaceMySql, bloomFilter)
	findNearbyPlacesUseCase := usecasePlaces.NewFindNearbyPlacesUseCase(repoPlaceMySql, googleClient, bloomFilter)
	findPlacesByAccessibilityFeatureUseCase := usecasePlaces.NewFindPlacesByAccessibilityFeatureUseCase(repoPlaceMySql)
	updatePlaceByIdUseCase := usecasePlaces.NewUpdatePlaceByIDUseCase(repoPlaceMySql)

	placesHandlers := web.NewWebPlacesHandlers(
		*createPlaceUseCase,
		*deletePlaceByIdUseCase,
		*findNearbyPlacesUseCase,
		*findPlaceByIdUseCase,
		*findPlacesByAccessibilityFeatureUseCase,
		*updatePlaceByIdUseCase,
	)

	webserver.AddHandler("/places/create", placesHandlers.CreatePlace)
	webserver.AddHandler("/places/delete", placesHandlers.DeletePlaceByID)
	webserver.AddHandler("/places/find", placesHandlers.FindPlaceByID)
	webserver.AddHandler("/places/nearby", placesHandlers.FindNearbyPlaces)
	webserver.AddHandler("/places/feature", placesHandlers.FindPlacesByAccessibilityFeature)
	webserver.AddHandler("/places/update", placesHandlers.UpdatePLaceByID)

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
	webserver.AddHandler("/reviews/create", reviewsHandlers.CreateReview)
	webserver.AddHandler("/reviews/delete", reviewsHandlers.DeleteReviewByID)
	webserver.AddHandler("/reviews/find", reviewsHandlers.FindReviewByID)
	webserver.AddHandler("/reviews/find-by-placeid", reviewsHandlers.FindReviewsByPlaceID)
	webserver.AddHandler("/reviews/find-by-userid", reviewsHandlers.FindReviewsByUserID)
	webserver.AddHandler("/reviews/update", reviewsHandlers.UpdateReviewByID)
	webserver.AddHandler("/reviews/add-feature", reviewsHandlers.AddAccessibilityFeatureByReviewID)
}

func setupUsersHandlers(webserver *webserver.WebServer, repoUserMySql gateway.UsersGateway) {
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
	)

	webserver.AddHandler("/users/create", usersHandlers.CreateUser)
	webserver.AddHandler("/users/delete", usersHandlers.DeleteUserByID)
	webserver.AddHandler("/users/list", usersHandlers.ListAllUsers)
	webserver.AddHandler("/users/find", usersHandlers.FindUserByID)
	webserver.AddHandler("/users/find-by-email", usersHandlers.FindUserByEmail)
	webserver.AddHandler("/users/update", usersHandlers.UpdateUserByID)
	webserver.AddHandler("/users/update-points", usersHandlers.UpdateUserPointsByUserId)
	webserver.AddHandler("/users/add-favorite", usersHandlers.AddFavoriteByUserIdAndPlaceId)
	webserver.AddHandler("/users/delete-favorite", usersHandlers.DeleteFavoriteByUserIdAndPlaceId)
}
