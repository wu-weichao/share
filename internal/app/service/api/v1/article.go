package v1

import (
	"github.com/gin-gonic/gin"
	"share/internal/app/service/api"
	"share/internal/models"
	"strconv"
	"strings"
)

type ArticleRequest struct {
	Tags        string `form:"tags" binding:"required"`
	Title       string `form:"title" binding:"required"`
	Cover       string `form:"cover"`
	Keywords    string `form:"keywords"`
	Description string `form:"description"`
	Content     string `form:"content" binding:"required"`
	Type        int    `form:"type" binding:"number"`
}

type ArticleResponse struct {
	ID          uint   `json:"id"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Type        int    `json:"type"`

	Tags []*ArticleTagResponse `json:"tags"`
}

type ArticleTagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

func GetArticles(c *gin.Context) {
	p := api.NewPagination(c)
	var maps map[string]interface{}
	total, err := models.ArticleGetTotal(maps)
	if err != nil {
		api.ErrorRequest(c, "Get Articles failed")
		return
	}
	p.Total = total
	if total == 0 {
		api.SuccessPagination(c, []interface{}{}, p)
		return
	}
	articles, err := models.ArticleGetAll(p.Page, p.PageSize, maps)
	if err != nil {
		api.ErrorRequest(c, "Get Articles failed")
		return
	}
	// get Article tag ids
	var articleIds []int
	for _, article := range articles {
		articleIds = append(articleIds, int(article.ID))
	}
	articleTags, err := models.ArticleTagGetByArticleIds(articleIds)
	if err != nil {
		api.ErrorRequest(c, "Get Articles failed")
		return
	}
	var r []*ArticleResponse
	for _, article := range articles {
		var resTags []*ArticleTagResponse
		if tags, ok := articleTags[int(article.ID)]; ok {
			for _, tag := range tags {
				resTags = append(resTags, &ArticleTagResponse{
					ID:   tag.ID,
					Name: tag.Name,
					Flag: tag.Flag,
				})
			}
		}
		r = append(r, &ArticleResponse{
			ID:          article.ID,
			CreatedAt:   article.CreatedAt,
			UpdatedAt:   article.UpdatedAt,
			Title:       article.Title,
			Cover:       article.Cover,
			Keywords:    article.Keywords,
			Description: article.Description,
			//Content:     article.Content,
			Type: article.Type,
			Tags: resTags,
		})

	}

	api.SuccessPagination(c, r, p)
}

func GetArticle(c *gin.Context) {
	id := c.Param("id")
	artId, _ := strconv.Atoi(id)
	article, err := models.ArticleGetById(artId)
	if err != nil {
		api.ErrorRequest(c, "Article not exists")
		return
	}
	//tagID
	//article.Tags

	api.Success(c, ArticleResponse{
		ID:          article.ID,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
		Title:       article.Title,
		Cover:       article.Cover,
		Keywords:    article.Keywords,
		Description: article.Description,
		Content:     article.Content,
		Type:        article.Type,
		Tags:        nil,
	})

}

func StoreArticle(c *gin.Context) {
	var form ArticleRequest
	var err error
	if err = c.ShouldBind(&form); err != nil {
		api.ErrorRequest(c, err.Error())
		return
	}
	// check tags
	var tagIds []int
	for _, s := range strings.Split(form.Tags, ",") {
		i, _ := strconv.Atoi(s)
		tagIds = append(tagIds, i)
	}
	tags, err := models.TagGetByIds(tagIds)
	if err != nil || len(tags) == 0 {
		api.ErrorRequest(c, "Tags params failed")
		return
	}
	if form.Keywords == "" {
		var words []string
		for _, tag := range tags {
			words = append(words, tag.Name)
		}
		form.Keywords = strings.Join(words, ",")
	}
	if form.Type == models.ArticleTypeDefault {
		form.Type = models.ArticleTypeMarkdown
	}
	// add article
	article, err := models.ArticleAdd(map[string]interface{}{
		"title":       form.Title,
		"cover":       form.Cover,
		"keywords":    form.Keywords,
		"description": form.Description,
		"content":     form.Content,
		"type":        form.Type,
		"tags":        tagIds,
	})
	if err != nil {
		api.ErrorRequest(c, "Article add failed")
		return
	}
	var articleTags []*ArticleTagResponse
	for _, tag := range tags {
		articleTags = append(articleTags, &ArticleTagResponse{
			ID:   tag.ID,
			Name: tag.Name,
			Flag: tag.Flag,
		})
	}

	api.Success(c, ArticleResponse{
		ID:          article.ID,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
		Title:       article.Title,
		Cover:       article.Cover,
		Keywords:    article.Keywords,
		Description: article.Description,
		Content:     article.Content,
		Type:        article.Type,
		Tags:        articleTags,
	})

}

func UpdateArticle(c *gin.Context) {

}

func DeleteArticle(c *gin.Context) {

}
