package main

import (
	"net/http"

	"gorm.io/driver/sqlite"

	"github.com/felipefabricio/golang-rest-api/configs"
	"github.com/felipefabricio/golang-rest-api/internal/entity"
	"github.com/felipefabricio/golang-rest-api/internal/infra/database"
	"github.com/felipefabricio/golang-rest-api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig("configs")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Client{}, &entity.Product{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/products/{id}", productHandler.GetProductById)
	r.Post("/products", productHandler.CreateProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	http.ListenAndServe(":8080", r)
}
