package main

import (
	"go-crud/app"
	"go-crud/controller"
	"go-crud/helper"
	"go-crud/repository"
	"go-crud/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	categoryRepository := repository.NewCategoryRepostory(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:2000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
