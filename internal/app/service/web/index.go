package web

import (
	"github.com/gin-gonic/gin"
	"share/internal/models"
	"strconv"
)

func Homepage(c *gin.Context) {
	var page int
	var pageSize int = 10
	if num := c.Param("num"); num != "" {
		page, _ = strconv.Atoi(num)
	}
	if page <= 0 {
		page = 1
	}
	// get articles
	maps := map[string]interface{}{
		"published = ?": 1,
	}
	total, err := models.ArticleGetTotal(maps)
	if err != nil {
		View500(c)
		return
	}
	if total == 0 {
		View404(c)
		return
	}
	articles, err := models.ArticleGetAll(page, pageSize, maps, "published_at desc")
	if err != nil {
		View500(c)
		return
	}

	//for _, article := range articles {
	//
	//}

	View(c, "home.tmpl", gin.H{
		"articles": articles,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
	//[]map[string]interface{}{
	//	{
	//		"title":     "文章1",
	//		"create_at": "2020-02-26",
	//		"views":     100,
	//		"tags":      "golang",
	//		"intro":     "xxxxxxxxxxxxxxxxxx",
	//	},
	//	{
	//		"title":     "文章2",
	//		"create_at": "2020-02-25",
	//		"views":     342,
	//		"tags":      "php",
	//		"intro":     "yyyyyyyyyyyyyyyyy",
	//	},
	//}
}
