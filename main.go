package main

import (
	"belajar_belajar_golang_restfull_api/app"
	"belajar_belajar_golang_restfull_api/exception"
	"belajar_belajar_golang_restfull_api/handler"
	"belajar_belajar_golang_restfull_api/helper"
	"belajar_belajar_golang_restfull_api/middleware"
	"belajar_belajar_golang_restfull_api/repository"
	"belajar_belajar_golang_restfull_api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	validate := validator.New()
	db := app.NewDB()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryHandler.FindAll)
	router.GET("/api/categories/:categoryId", categoryHandler.FindById)

	router.POST("/api/categories", categoryHandler.Create)
	router.PUT("/api/categories/:categoryId/update", categoryHandler.Update)
	router.DELETE("/api/categories/:categoryId/delete", categoryHandler.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	router.PanicHandler = exception.ErrorHandler

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
