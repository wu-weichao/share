package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"share/configs"
)

func View(c *gin.Context, tpl string, h gin.H) {
	c.HTML(http.StatusOK, tpl, formatViewData(h))
}

func View404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.tmpl", formatViewData(gin.H{
		"errorMessage": "抱歉，你访问的页面不存在",
	}))
}

func View500(c *gin.Context) {
	c.HTML(http.StatusInternalServerError, "500.tmpl", formatViewData(gin.H{
		"errorMessage": "抱歉，服务器出错了",
	}))
}

func formatViewData(h gin.H) gin.H {
	// common info
	data := gin.H{
		"headTitle":       configs.Html.Title,
		"headKeywords":    configs.Html.Keywords,
		"headDescription": configs.Html.Description,
	}
	// merge data
	for s, i := range h {
		data[s] = i
	}
	return data
}
