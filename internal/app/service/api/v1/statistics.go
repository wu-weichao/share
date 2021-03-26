package v1

import (
	"github.com/gin-gonic/gin"
	"share/internal/app/service/api"
	"share/internal/database"
	"share/internal/models"
)

func GetNewVisitCount(c *gin.Context) {
	count := 0
	if database.Redis != nil {

	}
	api.Success(c, count)
}

func GetVisitCount(c *gin.Context) {
	count := 0
	if database.Redis != nil {

	}
	api.Success(c, count)
}

func GetViewCount(c *gin.Context) {
	count, _ := models.ArticleViewCount()
	api.Success(c, count)
}

func GetArticlyCount(c *gin.Context) {
	count, _ := models.ArticleGetTotal(map[string]interface{}{})
	api.Success(c, count)
}
