package handlerArticleResults

import (
	"net/http"

	"github.com/gin-gonic/gin"
	articleResults "github.com/jrsmth97/gin-rest-api/controllers/article-controllers/results"
	util "github.com/jrsmth97/gin-rest-api/utils"
)

type handler struct {
	service articleResults.Service
}

func NewHandlerArticleResults(service articleResults.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ArticleResultsHandler(ctx *gin.Context) {

	results, errResult := h.service.ArticlesService()

	switch errResult {

	case "RESULTS_ARTICLE_NOT_FOUND_404":
		util.APIResponse(ctx, "Article data is not exists", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "Retrieve article data successfully", http.StatusOK, http.MethodPost, results)
	}
}
