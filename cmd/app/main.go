package main

import (
	"fmt"
	"log"

	"github.com/firerplayer/roda-belem-service/internal/domain/entity"
	"github.com/joho/godotenv"
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
	// LoadEnv()
	user, err := entity.NewUser("email", "username", "password1234")
	if err != nil {
		log.Fatalf("Falha ao criar o usuário: %v", err)
	}
	fmt.Println(user.Email)
	fmt.Println(user.Username)
	fmt.Println(user.Password)

}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Falha ao carregar o arquivo .env: %v", err)
	}
	log.Println("Arquivo .env carregado")
}
