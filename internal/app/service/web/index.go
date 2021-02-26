package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"posts": []map[string]interface{}{
			{
				"title":     "文章1",
				"create_at": "2020-02-26",
				"views":     100,
				"tags":      "golang",
				"intro":     "xxxxxxxxxxxxxxxxxx",
			},
			{
				"title":     "文章2",
				"create_at": "2020-02-25",
				"views":     342,
				"tags":      "php",
				"intro":     "yyyyyyyyyyyyyyyyy",
			},
		},
	})
}
