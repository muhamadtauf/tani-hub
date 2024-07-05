package request

type CreateArticleRequest struct {
	Title    string `validate:"required" json:"title"`
	Content  string `validate:"required" json:"content"`
	IsAtHome bool   `json:"is_at_home"`
}
