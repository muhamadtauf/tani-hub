package repository

import (
	"errors"
	"gorm.io/gorm"
	"tani-hub/data/request"
	"tani-hub/helper"
	"tani-hub/model"
	"time"
)

type CategoryRepositoryImpl struct {
	Db *gorm.DB
}

func NewCategoryRepositoryImpl(Db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{Db: Db}
}

func (t CategoryRepositoryImpl) Save(category model.Category) {
	result := t.Db.Create(&category)
	helper.ErrorPanic(result.Error)

}

func (t CategoryRepositoryImpl) Update(category model.Category) {
	var updateCategory = request.UpdateCategoryRequest{
		Id:        category.Id,
		Name:      category.Name,
		UpdatedAt: time.Now(),
	}
	result := t.Db.Model(&category).Updates(updateCategory)
	helper.ErrorPanic(result.Error)
}

func (t CategoryRepositoryImpl) Delete(categoryId int) {
	var category model.Category
	result := t.Db.Where("id = ?", categoryId).Delete(&category)
	helper.ErrorPanic(result.Error)
}

func (t CategoryRepositoryImpl) FindById(categoryId int) (model.Category, error) {
	var category model.Category
	result := t.Db.Find(&category, categoryId)
	if result != nil {
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (t CategoryRepositoryImpl) FindAll() []model.Category {
	var category []model.Category
	results := t.Db.Find(&category)
	helper.ErrorPanic(results.Error)
	return category
}
