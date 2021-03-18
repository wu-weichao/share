package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func View(c *gin.Context, tpl string, h gin.H) {
	c.HTML(http.StatusOK, tpl, h)
}

func View404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.tmpl", gin.H{
		"message": "抱歉，你访问的页面不存在",
	})
}

func View500(c *gin.Context) {
	c.HTML(http.StatusInternalServerError, "500.tmpl", gin.H{
		"message": "抱歉，服务器出错了",
	})
}
