package createArticle

import (
	model "github.com/jrsmth97/gin-rest-api/models"
)

type Service interface {
	CreateArticleService(input *CreateArticleInput) (*model.ArticleEntity, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateArticleService(input *CreateArticleInput) (*model.ArticleEntity, string) {
	articles := model.ArticleEntity{
		Title:    input.Title,
		Content:  input.Content,
		AuthorId: input.AuthorId,
	}

	createArticleResult, errCreate := s.repository.CreateArticleRepository(&articles)

	return createArticleResult, errCreate
}
