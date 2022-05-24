package articleResults

import (
	model "github.com/jrsmth97/gin-rest-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	ArticlesRepository() (*[]model.ArticleEntity, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryResults(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ArticlesRepository() (*[]model.ArticleEntity, string) {

	var articles []model.ArticleEntity
	db := r.db.Model(&articles)
	errorCode := make(chan string, 1)

	results := db.Debug().Select("*").Find(&articles)

	if results.Error != nil {
		errorCode <- "RESULTS_STUDENT_NOT_FOUND_404"
		return &articles, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &articles, <-errorCode
}
