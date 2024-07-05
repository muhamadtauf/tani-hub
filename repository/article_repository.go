package repository

import (
	"tani-hub/model"
)

type ArticleRepository interface {
	Save(artcile model.Article)
	Update(artcile model.Article)
	Delete(artcileId int)
	FindById(artcileId int) (artcile model.Article, err error)
	FindAll() []model.Article
}
