package forgotAuth

import (
	model "github.com/jrsmth97/gin-rest-api/models"
)

type Service interface {
	ForgotService(input *InputForgot) (*model.UsersEntity, string)
}

type service struct {
	repository Repository
}

func NewServiceForgot(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ForgotService(input *InputForgot) (*model.UsersEntity, string) {

	users := model.UsersEntity{
		Email: input.Email,
	}

	resultRegister, errRegister := s.repository.ForgotRepository(&users)

	return resultRegister, errRegister
}
