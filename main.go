package main

import (
	"go-crud/app"
	"go-crud/controller"
	"go-crud/helper"
	"go-crud/middleware"
	"go-crud/repository"
	"go-crud/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	categoryRepository := repository.NewCategoryRepostory(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:2000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
