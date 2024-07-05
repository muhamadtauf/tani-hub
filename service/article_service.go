package service

import (
	"tani-hub/data/request"
	"tani-hub/data/response"
)

type ArticleService interface {
	Create(article request.CreateArticleRequest)
	Update(article request.UpdateArticleRequest)
	Delete(articleId int)
	FindById(articleId int) response.ArticleResponse
	FindAll() []response.ArticleResponse
}
