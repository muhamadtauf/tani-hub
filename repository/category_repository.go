package repository

import (
	"tani-hub/model"
)

type CategoryRepository interface {
	Save(category model.Category)
	Update(category model.Category)
	Delete(categoryId int)
	FindById(categoryId int) (category model.Category, err error)
	FindAll() []model.Category
}
