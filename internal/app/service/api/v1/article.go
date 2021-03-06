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
	Description string `form:"description" binding:"required"`
	Content     string `form:"content" binding:"required"`
	Type        int    `form:"type" binding:"number"`
	Published   int    `form:"published" binding:"number"`
}

type ArticleResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Type        int    `json:"type"`
	View        int    `json:"view"`
	Published   int    `json:"published"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
	PublishedAt int    `json:"published_at"`

	Tags []*ArticleTagResponse `json:"tags"`
}

type ArticleTagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

func GetArticles(c *gin.Context) {
	p := api.NewPagination(c)
	maps := make(map[string]interface{})
	tags := c.Query("tags")
	if tags != "" {
		var tagIds []int
		for _, s := range strings.Split(tags, ",") {
			i, _ := strconv.Atoi(s)
			tagIds = append(tagIds, i)
		}
		articleIds, err := models.ArticleIdGetByTagIds(tagIds)
		if err != nil {
			api.ErrorRequest(c, "Get Articles failed")
		}
		maps["id IN ?"] = articleIds
	}
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
	sort := c.Query("sort")
	articles, err := models.ArticleGetAll(p.Page, p.PageSize, maps, sort)
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
			Title:       article.Title,
			Cover:       article.Cover,
			Keywords:    article.Keywords,
			Description: article.Description,
			//Content:     article.Content,
			Type:        article.Type,
			View:        article.View,
			Published:   article.Published,
			CreatedAt:   article.CreatedAt,
			UpdatedAt:   article.UpdatedAt,
			PublishedAt: article.PublishedAt,

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
	tags, err := models.ArticleTagGetByArticleId(artId)
	if err != nil {
		api.ErrorRequest(c, "Get article failed")
		return
	}
	//article.Tags
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
		Title:       article.Title,
		Cover:       article.Cover,
		Keywords:    article.Keywords,
		Description: article.Description,
		Content:     article.Content,
		Type:        article.Type,
		Published:   article.Published,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
		PublishedAt: article.PublishedAt,
		Tags:        articleTags,
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
		"published":   form.Published,
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
	var form ArticleRequest
	var err error
	if err = c.ShouldBind(&form); err != nil {
		api.ErrorRequest(c, err.Error())
		return
	}
	// get article
	id := c.Param("id")
	artId, _ := strconv.Atoi(id)
	_, err = models.ArticleGetById(artId)
	if err != nil {
		api.ErrorRequest(c, "Article not exists")
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

	// update article
	article, err := models.ArticleUpdate(artId, map[string]interface{}{
		"title":       form.Title,
		"cover":       form.Cover,
		"keywords":    form.Keywords,
		"description": form.Description,
		"content":     form.Content,
		"type":        form.Type,
		"published":   form.Published,
		"tags":        tagIds,
	})
	if err != nil {
		api.ErrorRequest(c, "Article update failed")
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

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	artId, _ := strconv.Atoi(id)
	err := models.ArticleDelete(artId)
	if err != nil {
		api.ErrorRequest(c, "Article delete failed")
		return
	}
	api.Success(c, "")
}

func PublishArticle(c *gin.Context) {
	id := c.Param("id")
	artId, _ := strconv.Atoi(id)
	if models.ArticlePublish(artId, 1) == false {
		api.ErrorRequest(c, "Article publish failed")
		return
	}
	api.Success(c, nil)
}

func UnpublishArticle(c *gin.Context) {
	id := c.Param("id")
	artId, _ := strconv.Atoi(id)
	if models.ArticlePublish(artId, 0) == false {
		api.ErrorRequest(c, "Article unpublish failed")
		return
	}
	//TODO: remove html

	api.Success(c, nil)
}
