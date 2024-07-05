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

type ArticleController struct {
	articleService service.ArticleService
}

func NewArticleController(service service.ArticleService) *ArticleController {
	return &ArticleController{articleService: service}
}

func (controller *ArticleController) Create(ctx *gin.Context) {
	createArticleRequest := request.CreateArticleRequest{}
	err := ctx.ShouldBindJSON(&createArticleRequest)
	helper.ErrorPanic(err)

	controller.articleService.Create(createArticleRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ArticleController) Update(ctx *gin.Context) {
	updateArticleRequest := request.UpdateArticleRequest{}
	err := ctx.ShouldBindJSON(&updateArticleRequest)
	helper.ErrorPanic(err)

	articleId := ctx.Param("articleId")
	id, err := strconv.Atoi(articleId)
	helper.ErrorPanic(err)

	updateArticleRequest.Id = id

	controller.articleService.Update(updateArticleRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ArticleController) Delete(ctx *gin.Context) {
	articleId := ctx.Param("articleId")
	id, err := strconv.Atoi(articleId)
	helper.ErrorPanic(err)
	controller.articleService.Delete(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *ArticleController) FindById(ctx *gin.Context) {
	articleId := ctx.Param("articleId")
	id, err := strconv.Atoi(articleId)
	helper.ErrorPanic(err)

	articleResponse := controller.articleService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   articleResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ArticleController) FindAll(ctx *gin.Context) {
	articleResponse := controller.articleService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   articleResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
