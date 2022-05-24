package activationAuth

import model "github.com/jrsmth97/gin-rest-api/models"

type Service interface {
	ActivationService(input *InputActivation) (*model.UsersEntity, string)
}

type service struct {
	repository Repository
}

func NewServiceActivation(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ActivationService(input *InputActivation) (*model.UsersEntity, string) {
	users := model.UsersEntity{
		Email:  input.Email,
		Active: input.Active,
	}

	activationResult, activationError := s.repository.ActivationRepository(&users)

	return activationResult, activationError
}
