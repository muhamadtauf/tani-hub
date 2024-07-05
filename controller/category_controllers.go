package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tani-hub/data/request"
	"tani-hub/data/response"
	"tani-hub/helper"
	"tani-hub/service"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(service service.CategoryService) *CategoryController {
	return &CategoryController{categoryService: service}
}

func (controller *CategoryController) Create(ctx *gin.Context) {
	createCategoryRequest := request.CreateCategoryRequest{}
	err := ctx.ShouldBindJSON(&createCategoryRequest)
	helper.ErrorPanic(err)

	controller.categoryService.Create(createCategoryRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryController) Update(ctx *gin.Context) {
	updateCategoryRequest := request.UpdateCategoryRequest{}
	err := ctx.ShouldBindJSON(&updateCategoryRequest)
	helper.ErrorPanic(err)

	categoryId := ctx.Param("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.ErrorPanic(err)

	updateCategoryRequest.Id = id

	controller.categoryService.Update(updateCategoryRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryController) Delete(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.ErrorPanic(err)
	controller.categoryService.Delete(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *CategoryController) FindById(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.ErrorPanic(err)

	categoryResponse := controller.categoryService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryController) FindAll(ctx *gin.Context) {
	categoryResponse := controller.categoryService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
