package request

import "time"

type UpdateArticleRequest struct {
	Id        int    `validate:"required"`
	Title     string `validate:"required" json:"title"`
	Content   string `validate:"required" json:"content"`
	IsAtHome  bool   `json:"is_at_home"`
	UpdatedAt time.Time
}
