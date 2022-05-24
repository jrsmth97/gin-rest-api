package createArticle

type CreateArticleInput struct {
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	AuthorId string `json:"author_id" validate:"required"`
}
