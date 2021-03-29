package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"share/internal/app/service/api"
	"share/internal/database"
	"share/internal/models"
	"strconv"
	"time"
)

func GetNewVisitCount(c *gin.Context) {
	count := 0
	if database.Redis != nil {
		uvTotal := database.Redis.PFCount("uv_total").Val()
		allUvTotal := database.Redis.PFCount("uv_total", fmt.Sprintf("uv:%s", time.Now().Format("20060102"))).Val()
		count = int(allUvTotal - uvTotal)
	}
	api.Success(c, count)
}

func GetVisitCount(c *gin.Context) {
	count := 0
	if database.Redis != nil {
		allUvTotal := database.Redis.PFCount("uv_total", fmt.Sprintf("uv:%s", time.Now().Format("20060102"))).Val()
		count = int(allUvTotal)
	}
	api.Success(c, count)
}

func GetViewCount(c *gin.Context) {
	count := 0
	if database.Redis != nil {
		count, _ = strconv.Atoi(database.Redis.Get("pv_total").Val())
	} else {
		count, _ = models.ArticleViewCount()
	}
	api.Success(c, count)
}

func GetArticlyCount(c *gin.Context) {
	count, _ := models.ArticleGetTotal(map[string]interface{}{})
	api.Success(c, count)
}

func GetStatisticsRange(c *gin.Context) {
	startDay := c.DefaultQuery("start", time.Now().AddDate(0, 0, -6).Format("20060102"))
	endDay := c.DefaultQuery("end", time.Now().Format("20060102"))
	if startDay == "" || endDay == "" {
		api.ErrorRequest(c, "day range empty")
		return
	}
	var days []string
	fmt.Println(startDay, endDay)
	var views []int
	for startDay <= endDay {
		t, _ := time.Parse("20060102", startDay)
		days = append(days, t.Format("01-02"))
		var count = 0
		if database.Redis != nil {
			count, _ = strconv.Atoi(database.Redis.Get(fmt.Sprintf("pv:%s", startDay)).Val())
		}
		views = append(views, count)
		startDay = t.AddDate(0, 0, 1).Format("20060102")
	}
	api.Success(c, map[string]interface{}{
		"days":  days,
		"views": views,
	})

}
