package resendAuth

import (
	model "github.com/jrsmth97/gin-rest-api/models"
)

type Service interface {
	ResendService(input *InputResend) (*model.UsersEntity, string)
}

type service struct {
	repository Repository
}

func NewServiceResend(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResendService(input *InputResend) (*model.UsersEntity, string) {

	users := model.UsersEntity{
		Email: input.Email,
	}

	resultRegister, errRegister := s.repository.ResendRepository(&users)

	return resultRegister, errRegister
}
