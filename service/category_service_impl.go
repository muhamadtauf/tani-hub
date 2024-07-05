package service

import (
	"github.com/go-playground/validator/v10"
	"tani-hub/data/request"
	"tani-hub/data/response"
	"tani-hub/helper"
	"tani-hub/model"
	"tani-hub/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		Validate:           validate,
	}
}

func (t CategoryServiceImpl) Create(category request.CreateCategoryRequest) {
	err := t.Validate.Struct(category)
	helper.ErrorPanic(err)
	categoryModel := model.Category{
		Name: category.Name,
	}
	t.CategoryRepository.Save(categoryModel)
}

func (t CategoryServiceImpl) Update(category request.UpdateCategoryRequest) {
	categoryData, err := t.CategoryRepository.FindById(category.Id)
	helper.ErrorPanic(err)
	categoryData.Name = category.Name
	t.CategoryRepository.Update(categoryData)
}

func (t CategoryServiceImpl) Delete(categoryId int) {
	t.CategoryRepository.Delete(categoryId)
}

func (t CategoryServiceImpl) FindById(categoryId int) response.CategoryResponse {
	categoryData, err := t.CategoryRepository.FindById(categoryId)
	helper.ErrorPanic(err)

	categoryResponse := response.CategoryResponse{
		Id:   categoryData.Id,
		Name: categoryData.Name,
	}
	return categoryResponse
}

func (t CategoryServiceImpl) FindAll() []response.CategoryResponse {
	result := t.CategoryRepository.FindAll()

	var categories []response.CategoryResponse
	for _, value := range result {
		category := response.CategoryResponse{
			//Id:        value.Id,
			Name:      value.Name,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		categories = append(categories, category)
	}
	return categories
}
