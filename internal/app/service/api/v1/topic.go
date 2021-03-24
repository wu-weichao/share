package v1

import (
	"github.com/gin-gonic/gin"
	"share/internal/app/service/api"
	"share/internal/models"
	"strconv"
)

type TopicStoreRequest struct {
	Title  string `form:"title" binding:"required"`
	Url    string `form:"url" binding:"required"`
	Sort   int    `form:"sort" binding:"number"`
	Status int    `form:"status" binding:"number,oneof=-1 1"`
}

type TopicResponse struct {
	ID        uint   `json:"id"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	Sort      int    `json:"sort"`
	Status    int    `json:"status"`
}

func GetTopics(c *gin.Context) {
	p := api.NewPagination(c)
	maps := make(map[string]interface{})
	title := c.Query("title")
	if title != "" {
		maps["title LIKE ?"] = "%" + title + "%"
	}
	total, err := models.TopicGetTotal(maps)
	if err != nil {
		api.ErrorRequest(c, "Get topics failed")
		return
	}
	p.Total = total
	if total == 0 {
		api.SuccessPagination(c, []interface{}{}, p)
		return
	}
	topics, err := models.TopicGetAll(p.Page, p.PageSize, maps)
	if err != nil {
		api.ErrorRequest(c, "Get topics failed")
		return
	}

	var r []*TopicResponse
	for _, topic := range topics {
		r = append(r, &TopicResponse{
			ID:        topic.ID,
			CreatedAt: topic.CreatedAt,
			Title:     topic.Title,
			Url:       topic.Url,
			Sort:      topic.Sort,
			Status:    topic.Status,
		})
	}

	api.SuccessPagination(c, r, p)
}

func GetTopic(c *gin.Context) {
	id := c.Param("id")
	topicId, _ := strconv.Atoi(id)
	topic, err := models.TopicGetById(topicId)
	if err != nil {
		api.ErrorRequest(c, "Topic not exists")
		return
	}

	api.Success(c, TopicResponse{
		ID:        topic.ID,
		CreatedAt: topic.CreatedAt,
		Title:     topic.Title,
		Url:       topic.Url,
		Sort:      topic.Sort,
		Status:    topic.Status,
	})
}

func StoreTopic(c *gin.Context) {
	var form TopicStoreRequest
	var err error
	if err = c.ShouldBind(&form); err != nil {
		api.ErrorRequest(c, err.Error())
		return
	}
	if form.Sort == 0 {
		form.Sort = 9999
	}
	if form.Status == 0 {
		form.Status = -1
	}
	// add topic
	tag, err := models.TopicAdd(map[string]interface{}{
		"title":  form.Title,
		"url":    form.Url,
		"sort":   form.Sort,
		"status": form.Status,
	})
	if err != nil {
		api.ErrorRequest(c, "Topic add failed")
		return
	}

	api.Success(c, TopicResponse{
		ID:        tag.ID,
		CreatedAt: tag.CreatedAt,
		Title:     tag.Title,
		Url:       tag.Url,
		Sort:      tag.Sort,
		Status:    tag.Status,
	})
}

func UpdateTopic(c *gin.Context) {
	var form TopicStoreRequest
	var err error
	if err = c.ShouldBind(&form); err != nil {
		api.ErrorRequest(c, err.Error())
		return
	}
	id := c.Param("id")
	topicId, _ := strconv.Atoi(id)
	topic, err := models.TopicGetById(topicId)
	if err != nil {
		api.ErrorRequest(c, "Topic not exists")
		return
	}
	if form.Status == 0 {
		form.Status = topic.Status
	}
	updateData := map[string]interface{}{
		"title":  form.Title,
		"url":    form.Url,
		"sort":   form.Sort,
		"status": form.Status,
	}
	_, err = models.TopicUpdate(topicId, updateData)
	if err != nil {
		api.ErrorRequest(c, "Topic Update failed")
		return
	}

	api.Success(c, "")
}

func DeleteTopic(c *gin.Context) {
	id := c.Param("id")
	topicId, _ := strconv.Atoi(id)
	err := models.TopicDelete(topicId)
	if err != nil {
		api.ErrorRequest(c, "Topic delete failed")
		return
	}
	api.Success(c, "")
}
