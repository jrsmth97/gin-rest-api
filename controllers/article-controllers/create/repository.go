package createArticle

import (
	model "github.com/jrsmth97/gin-rest-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateArticleRepository(input *model.ArticleEntity) (*model.ArticleEntity, string)
}

type repository struct {
	db *gorm.DB
}

func CreateNewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateArticleRepository(input *model.ArticleEntity) (*model.ArticleEntity, string) {

	var articles model.ArticleEntity
	db := r.db.Model(&articles)
	errorCode := make(chan string, 1)

	articles.Title = input.Title
	articles.Content = input.Content
	articles.AuthorId = input.AuthorId

	newArticle := db.Debug().Create(&articles)
	db.Commit()

	if newArticle.Error != nil {
		errorCode <- "FAILED_CREATE_ARTICLE_403"
		return &articles, <-errorCode
	}

	errorCode <- "nil"
	return &articles, <-errorCode
}
