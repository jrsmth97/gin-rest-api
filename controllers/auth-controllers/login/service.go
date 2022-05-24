package loginAuth

import (
	model "github.com/jrsmth97/gin-rest-api/models"
)

type Service interface {
	LoginService(input *InputLogin) (*model.UsersEntity, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginService(input *InputLogin) (*model.UsersEntity, string) {

	user := model.UsersEntity{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&user)

	return resultLogin, errLogin
}
