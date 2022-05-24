package route

import (
	"github.com/gin-gonic/gin"
	createArticle "github.com/jrsmth97/gin-rest-api/controllers/article-controllers/create"
	articleResults "github.com/jrsmth97/gin-rest-api/controllers/article-controllers/results"
	handlerCreateArticle "github.com/jrsmth97/gin-rest-api/handlers/article-handlers/create"
	handlerArticleResults "github.com/jrsmth97/gin-rest-api/handlers/article-handlers/results"
	middleware "github.com/jrsmth97/gin-rest-api/middlewares"
	"gorm.io/gorm"
)

func InitArticleRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Student
	*/
	createArticleRepository := createArticle.CreateNewRepository(db)
	createArticleService := createArticle.NewServiceCreate(createArticleRepository)
	createArticleHandler := handlerCreateArticle.NewHandlerCreateArticle(createArticleService)

	resultsArticleRepository := articleResults.NewRepositoryResults(db)
	resultsArticleService := articleResults.NewServiceResults(resultsArticleRepository)
	resultsArticleHandler := handlerArticleResults.NewHandlerArticleResults(resultsArticleService)

	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.GET("/articles", resultsArticleHandler.ArticleResultsHandler)
	groupRoute.POST("/article", createArticleHandler.CreateArticleHandler)
}
