package service

import (
	"github.com/go-playground/validator/v10"
	"tani-hub/data/request"
	"tani-hub/data/response"
	"tani-hub/helper"
	"tani-hub/model"
	"tani-hub/repository"
)

type ArticleServiceImpl struct {
	ArticleRepository repository.ArticleRepository
	Validate          *validator.Validate
}

func NewArticleServiceImpl(articleRepository repository.ArticleRepository, validate *validator.Validate) ArticleService {
	return &ArticleServiceImpl{
		ArticleRepository: articleRepository,
		Validate:          validate,
	}
}

func (t ArticleServiceImpl) Create(article request.CreateArticleRequest) {
	err := t.Validate.Struct(article)
	helper.ErrorPanic(err)
	articleModel := model.Article{
		Title:    article.Title,
		Content:  article.Content,
		IsAtHome: article.IsAtHome,
	}
	t.ArticleRepository.Save(articleModel)
}

func (t ArticleServiceImpl) Update(article request.UpdateArticleRequest) {
	articleData, err := t.ArticleRepository.FindById(article.Id)
	helper.ErrorPanic(err)
	articleData.Title = article.Title
	articleData.Content = article.Content
	articleData.IsAtHome = article.IsAtHome
	t.ArticleRepository.Update(articleData)
}

func (t ArticleServiceImpl) Delete(articleId int) {
	t.ArticleRepository.Delete(articleId)
}

func (t ArticleServiceImpl) FindById(articleId int) response.ArticleResponse {
	articleData, err := t.ArticleRepository.FindById(articleId)
	helper.ErrorPanic(err)

	articleResponse := response.ArticleResponse{
		Id:        articleData.Id,
		Title:     articleData.Title,
		Content:   articleData.Content,
		IsAtHome:  articleData.IsAtHome,
		CreatedAt: articleData.CreatedAt,
		UpdatedAt: articleData.UpdatedAt,
	}
	return articleResponse
}

func (t ArticleServiceImpl) FindAll() []response.ArticleResponse {
	result := t.ArticleRepository.FindAll()

	var categories []response.ArticleResponse
	for _, value := range result {
		article := response.ArticleResponse{
			Id:        value.Id,
			Title:     value.Title,
			Content:   value.Content,
			IsAtHome:  value.IsAtHome,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		categories = append(categories, article)
	}
	return categories
}
