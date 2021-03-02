package v1

import (
	"github.com/gin-gonic/gin"
	"share/internal/app/service/api"
	"share/internal/models"
	"strconv"
)

type TagStoreRequest struct {
	Name        string `form:"name" binding:"required"`
	Flag        string `form:"flag" binding:"required"`
	Icon        string `form:"icon"`
	Description string `form:"description"`
}

type TagResponse struct {
	ID          uint   `json:"id"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
	Name        string `json:"name"`
	Flag        string `json:"flag"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func GetTags(c *gin.Context) {
	p := api.NewPagination(c)
	var maps map[string]interface{}
	total, err := models.TagGetTotal(maps)
	if err != nil {
		api.ErrorRequest(c, "Get tags failed")
		return
	}
	p.Total = total
	if total == 0 {
		api.SuccessPagination(c, []interface{}{}, p)
		return
	}
	tags, err := models.TagGetAll(p.Page, p.PageSize, maps)
	if err != nil {
		api.ErrorRequest(c, "Get tags failed")
		return
	}

	var r []*TagResponse
	for _, tag := range tags {
		r = append(r, &TagResponse{
			ID:          tag.ID,
			CreatedAt:   tag.CreatedAt,
			Name:        tag.Name,
			Flag:        tag.Flag,
			Icon:        tag.Icon,
			Description: tag.Description,
			Status:      tag.Status,
		})
	}

	api.SuccessPagination(c, r, p)
}

func GetTag(c *gin.Context) {
	id := c.Param("id")
	tagId, _ := strconv.Atoi(id)
	tag, err := models.TagGetById(tagId)
	if err != nil {
		api.ErrorRequest(c, "Tag not exists")
		return
	}

	api.Success(c, TagResponse{
		ID:          tag.ID,
		CreatedAt:   tag.CreatedAt,
		Name:        tag.Name,
		Flag:        tag.Flag,
		Icon:        tag.Icon,
		Description: tag.Description,
		Status:      tag.Status,
	})
}

func StoreTag(c *gin.Context) {
	var form TagStoreRequest
	var err error
	if err = c.ShouldBind(&form); err != nil {
		api.ErrorRequest(c, err.Error())
		return
	}
	// tag exists
	if _, err = models.TagGetByFlag(form.Flag); err == nil {
		api.ErrorRequest(c, "Tag is exists")
		return
	}
	// add tag
	tag, err := models.TagAdd(map[string]interface{}{
		"name":        form.Name,
		"flag":        form.Flag,
		"icon":        form.Icon,
		"description": form.Description,
	})
	if err != nil {
		api.ErrorRequest(c, "Tag add failed")
		return
	}

	api.Success(c, TagResponse{
		ID:          tag.ID,
		CreatedAt:   tag.CreatedAt,
		Name:        tag.Name,
		Flag:        tag.Flag,
		Icon:        tag.Icon,
		Description: tag.Description,
		Status:      tag.Status,
	})
}

func UpdateTag(c *gin.Context) {
	id := c.Param("id")
	tagId, _ := strconv.Atoi(id)

	var err error
	tag, err := models.TagGetById(tagId)
	if err != nil {
		api.ErrorRequest(c, "Tag not exists")
		return
	}
	var form TagStoreRequest
	if err = c.ShouldBind(&form); err != nil {
		api.ErrorRequest(c, err.Error())
		return
	}
	if form.Flag != tag.Flag {
		if _, err = models.TagGetByFlag(form.Flag); err == nil {
			api.ErrorRequest(c, "Tag is exists")
			return
		}
	}

	_, err = models.TagUpdate(tagId, map[string]interface{}{
		"name":        form.Name,
		"flag":        form.Flag,
		"icon":        form.Icon,
		"description": form.Description,
	})
	if err != nil {
		api.ErrorRequest(c, "Tag Update failed")
		return
	}

	api.Success(c, "")
}
