package web

import (
	"github.com/gin-gonic/gin"
	"share/internal/models"
	"strconv"
)

type Article struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	View        int    `json:"view"`
	PublishedAt int    `json:"published_at"`

	Tags []string
}

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
	var total int
	var articles []*models.Article
	var articleList []*Article
	var err error
	if total, err = models.ArticleGetTotal(maps); err != nil {
		View500(c)
		return
	}
	if total > 0 {
		if articles, err = models.ArticleGetAll(page, pageSize, maps, "published_at desc"); err != nil {
			View500(c)
			return
		}
		if len(articles) > 0 {
			var articleIds []int
			for _, article := range articles {
				articleIds = append(articleIds, int(article.ID))
			}
			articleTags, err := models.ArticleTagGetByArticleIds(articleIds)
			if err != nil {
				View500(c)
				return
			}
			for _, article := range articles {
				item := &Article{
					ID:          article.ID,
					Title:       article.Title,
					Cover:       article.Cover,
					Description: article.Description,
					View:        article.View,
					PublishedAt: article.PublishedAt,
				}
				if tags, ok := articleTags[int(article.ID)]; ok {
					for _, tag := range tags {
						item.Tags = append(item.Tags, tag.Name)
					}
				}
				articleList = append(articleList, item)
			}
		}
	}
	View(c, "home.tmpl", gin.H{
		"articles": articleList,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}
