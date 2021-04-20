package web

import (
	"github.com/gin-gonic/gin"
	"share/internal/models"
	"strconv"
)

type ArticleDetail struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Html        string `json:"html"`
	Type        int    `json:"type"`
	View        int    `json:"view"`
	PublishedAt int    `json:"published_at"`
	Tags        []string
}

func ShowArticle(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		Homepage(c)
		return
	}
	var artId int
	artId, _ = strconv.Atoi(id)
	// get articel detail
	article, err := models.ArticleGetById(artId)
	if err != nil {
		// redirect to homepage
		Homepage(c)
		return
	}
	// add view count
	models.ArticleViewAdd(artId, 1)

	detail := ArticleDetail{
		ID:          article.ID,
		Title:       article.Title,
		Cover:       article.Cover,
		Keywords:    article.Keywords,
		Description: article.Description,
		Html:        article.Html,
		Type:        article.Type,
		View:        article.View + 1,
		PublishedAt: article.PublishedAt,
		Tags:        []string{},
	}
	tags, err := models.ArticleTagGetByArticleId(artId)
	if err != nil {
		View500(c)
		return
	}
	for _, tag := range tags {
		detail.Tags = append(detail.Tags, tag.Name)
	}

	View(c, "article", gin.H{
		"headerTitle":       detail.Title,
		"headerKeywords":    detail.Keywords,
		"headerDescription": detail.Description,
		"article":           detail,
	})
}
