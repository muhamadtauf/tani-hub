package repository

import (
	"errors"
	"gorm.io/gorm"
	"tani-hub/data/request"
	"tani-hub/helper"
	"tani-hub/model"
	"time"
)

type ArticleRepositoryImpl struct {
	Db *gorm.DB
}

func NewArticleRepositoryImpl(Db *gorm.DB) ArticleRepository {
	return &ArticleRepositoryImpl{Db: Db}
}

func (t ArticleRepositoryImpl) Save(article model.Article) {
	result := t.Db.Create(&article)
	helper.ErrorPanic(result.Error)
}

func (t ArticleRepositoryImpl) Update(article model.Article) {
	var updateArticle = request.UpdateArticleRequest{
		Id:        article.Id,
		Title:     article.Title,
		Content:   article.Content,
		IsAtHome:  article.IsAtHome,
		UpdatedAt: time.Now(),
	}
	result := t.Db.Model(&article).Updates(updateArticle)
	helper.ErrorPanic(result.Error)
}

func (t ArticleRepositoryImpl) Delete(articleId int) {
	var article model.Article
	result := t.Db.Where("id = ?", articleId).Delete(&article)
	helper.ErrorPanic(result.Error)
}

func (t ArticleRepositoryImpl) FindById(articleId int) (model.Article, error) {
	var article model.Article
	result := t.Db.Find(&article, articleId)
	if result != nil {
		return article, nil
	} else {
		return article, errors.New("article is not found")
	}
}

func (t ArticleRepositoryImpl) FindAll() []model.Article {
	var article []model.Article
	results := t.Db.Find(&article)
	helper.ErrorPanic(results.Error)
	return article
}
