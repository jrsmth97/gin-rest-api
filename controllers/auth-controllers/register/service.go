package registerAuth

import (
	model "github.com/jrsmth97/gin-rest-api/models"
)

type Service interface {
	RegisterService(input *InputRegister) (*model.UsersEntity, string)
}

type service struct {
	repository Repository
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input *InputRegister) (*model.UsersEntity, string) {

	users := model.UsersEntity{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
		Active:   true,
	}

	resultRegister, errRegister := s.repository.RegisterRepository(&users)

	return resultRegister, errRegister
}
