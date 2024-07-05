package request

import "time"

type UpdateCategoryRequest struct {
	Id        int    `validate:"required"`
	Name      string `validate:"required" json:"name"`
	UpdatedAt time.Time
}
