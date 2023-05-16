package main

import "fmt"

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
	// db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// repo := repository.NewProductRepositoryMysql(db)
	// createProductUseCase := usecase.NewCreateProductUseCase(repo)
	// listProductsUseCase := usecase.NewListProductsUseCase(repo)

	// // Web
	// productHandlers := web.NewProductHandlers(createProductUseCase, listProductsUseCase)
	// r := chi.NewRouter()
	// r.Post("/products", productHandlers.CreateProduct)
	// r.Get("/products/all", productHandlers.ListProducts)

	// go http.ListenAndServe(":8000", r)
	fmt.Println("Tem que fazer muita coisa ainda")
}
