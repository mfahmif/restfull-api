package main

import (
	"net/http"
	"restful-api/app"
	"restful-api/controller"
	"restful-api/helper"
	"restful-api/middleware"
	"restful-api/repository"
	"restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.Newdb()
	validator := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validator)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.New_Auth_Middleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
