package articleResults

import (
	model "github.com/jrsmth97/gin-rest-api/models"
)

type Service interface {
	ArticlesService() (*[]model.ArticleEntity, string)
}

type service struct {
	repository Repository
}

func NewServiceResults(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ArticlesService() (*[]model.ArticleEntity, string) {

	results, err := s.repository.ArticlesRepository()

	return results, err
}
