package request

type CreateCategoryRequest struct {
	Name string `validate:"required" json:"name"`
}
