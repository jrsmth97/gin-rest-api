package resetAuth

import model "github.com/jrsmth97/gin-rest-api/models"

type Service interface {
	ResetService(input *InputReset) (*model.UsersEntity, string)
}

type service struct {
	repository Repository
}

func NewServiceReset(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResetService(input *InputReset) (*model.UsersEntity, string) {

	users := model.UsersEntity{
		Email:    input.Email,
		Password: input.Password,
		Active:   input.Active,
	}

	resetResult, errResult := s.repository.ResetRepository(&users)

	return resetResult, errResult
}
