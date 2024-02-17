package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/renanmav/GoExpert-API/configs"
	"github.com/renanmav/GoExpert-API/internal/entity"
	"github.com/renanmav/GoExpert-API/internal/infra/database"
	"github.com/renanmav/GoExpert-API/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	config := configs.LoadConfig(".")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&entity.User{}, &entity.Product{})
	if err != nil {
		panic(err)
	}

	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/products", productHandler.CreateProduct)
	router.Get("/products/{id}", productHandler.GetProduct)

	err = http.ListenAndServe(":"+config.WebServerPort, router)
	if err != nil {
		return
	}
}
