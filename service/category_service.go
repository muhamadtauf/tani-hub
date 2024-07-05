package service

import (
	"tani-hub/data/request"
	"tani-hub/data/response"
)

type CategoryService interface {
	Create(category request.CreateCategoryRequest)
	Update(category request.UpdateCategoryRequest)
	Delete(categoryId int)
	FindById(categoryId int) response.CategoryResponse
	FindAll() []response.CategoryResponse
}
