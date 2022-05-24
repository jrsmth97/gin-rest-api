package registerAuth

import (
	model "github.com/jrsmth97/gin-rest-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(input *model.UsersEntity) (*model.UsersEntity, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterRepository(input *model.UsersEntity) (*model.UsersEntity, string) {

	var users model.UsersEntity
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &users, <-errorCode
	}

	users.Fullname = input.Fullname
	users.Email = input.Email
	users.Password = input.Password
	users.Active = input.Active

	addNewUser := db.Debug().Create(&users)
	db.Commit()

	if addNewUser.Error != nil {
		errorCode <- "REGISTER_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
