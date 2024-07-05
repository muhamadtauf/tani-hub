package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"tani-hub/config"
	"tani-hub/controller"
	"tani-hub/helper"
	"tani-hub/initializer"
	"tani-hub/model"
	"tani-hub/repository"
	"tani-hub/router"
	"tani-hub/service"
	"time"
)

func main() {

	//Database
	initializer.ConnectToDb()
	db := config.DatabaseConnection()
	validate := validator.New()

	//db.Table("tags").AutoMigrate(&model.Tags{})
	//db.Table("users").AutoMigrate(&model.User{})
	//
	//db.Table("products").AutoMigrate(&model.Product{})
	//db.Table("categories").AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Category{}, &model.Product{}, &model.Article{})

	tagRepository := repository.NewTagsRepositoryImpl(db)
	tagService := service.NewTagServiceImpl(tagRepository, validate)
	tagController := controller.NewTagController(tagService)

	categoryRepository := repository.NewCategoryRepositoryImpl(db)
	categoryService := service.NewCategoryServiceImpl(categoryRepository, validate)
	categoryController := controller.NewCategoryController(categoryService)

	articleRepository := repository.NewArticleRepositoryImpl(db)
	articleService := service.NewArticleServiceImpl(articleRepository, validate)
	articleController := controller.NewArticleController(articleService)

	//router
	routes := router.NewRouter(tagController, categoryController, articleController)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
