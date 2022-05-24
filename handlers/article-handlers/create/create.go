package createArticleHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	createArticle "github.com/jrsmth97/gin-rest-api/controllers/article-controllers/create"
	util "github.com/jrsmth97/gin-rest-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handler struct {
	service createArticle.Service
}

func NewHandlerCreateArticle(service createArticle.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateArticleHandler(ctx *gin.Context) {

	var input createArticle.CreateArticleInput
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Title",
				Message: "title is required on body",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	userIdSession, _ := ctx.Get("user_id")
	input.AuthorId = userIdSession.(string)

	_, errCreateArticle := h.service.CreateArticleService(&input)

	if errCreateArticle == "CREATE_ARTICLE_FAILED_403" {
		util.APIResponse(ctx, "Create new article failed", http.StatusForbidden, http.MethodPost, nil)
		return
	}

	util.APIResponse(ctx, "Create new article successfully", http.StatusCreated, http.MethodPost, nil)
}
