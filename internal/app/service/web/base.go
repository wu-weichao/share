package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"regexp"
	"share/configs"
	"share/internal/models"
	"strings"
)

func View(c *gin.Context, tpl string, h gin.H) {
	c.HTML(http.StatusOK, tpl, formatViewData(h))
}

func View404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404", formatViewData(gin.H{
		"errorMessage": "抱歉，你访问的页面不存在",
	}))
}

func View500(c *gin.Context) {
	c.HTML(http.StatusInternalServerError, "500", formatViewData(gin.H{
		"errorMessage": "抱歉，服务器出错了",
	}))
}

func formatViewData(h gin.H) gin.H {
	// common info
	data := gin.H{
		"headerTitle":       configs.Html.Title,
		"headerKeywords":    configs.Html.Keywords,
		"headerDescription": configs.Html.Description,
		"footerCopyright":   configs.Html.Copyright,
		"footerIcp":         configs.Html.Icp,
	}
	data["topics"], _ = getViewTopic()

	// merge data
	for s, i := range h {
		data[s] = i
	}
	fmt.Println(data)
	return data
}

func getViewTopic() ([]*models.Topic, error) {
	var topics []*models.Topic
	topics, err := models.TopicGetSimpleAll(map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return topics, nil
}

type Pagination struct {
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	Total     int    `json:"total"`
	TotalPage int    `json:"total_page"`
	PageNums  []int  `json:"page_nums"`
	Prev      int    `json:"prev"`
	Next      int    `json:"next"`
	Link      string `json:"link"`
}

func formatPagination(p *Pagination) *Pagination {
	// 计算页数
	p.TotalPage = int(math.Ceil(float64(p.Total) / float64(p.PageSize)))
	// 页数列表
	var nums []int
	minPageSize := 7
	minNumLen := 5
	if p.TotalPage <= minPageSize {
		for i := 1; i <= p.TotalPage; i++ {
			nums = append(nums, i)
		}
	} else {
		// 两侧数据
		if (p.Page - 1) <= 3 {
			for i := 1; i <= minNumLen+1; i++ {
				nums = append(nums, i)
			}
		} else if p.TotalPage-p.Page <= 3 {
			for i := p.TotalPage - minNumLen; i <= p.TotalPage; i++ {
				nums = append(nums, i)
			}
		} else {
			for i := 1; i < p.TotalPage; i++ {
				if math.Abs(float64(p.Page-i)) <= 2 {
					nums = append(nums, i)
				}
			}
		}
		// append start and end
		var start []int
		if nums[0] != 1 {
			start = append(start, 1)
			if nums[0]-1 > 1 {
				start = append(start, 0)
			}
		}
		var end []int
		if nums[len(nums)-1] != p.TotalPage {
			if p.TotalPage-nums[len(nums)-1] > 1 {
				end = append(end, 0)
			}
			end = append(end, p.TotalPage)
		}
		// merge nums
		n := make([]int, len(start)+len(nums)+len(end))
		at := copy(n, start)
		at += copy(n[at:], nums)
		copy(n[at:], end)
		nums = n
	}
	p.PageNums = nums
	p.Prev = p.Page - 1
	if p.Prev < 1 {
		p.Prev = 1
	}
	p.Next = p.Page + 1
	if p.Next > p.TotalPage {
		p.Next = p.TotalPage
	}
	// format link
	reg, _ := regexp.Compile(`/page/[:\/\w]+`)
	p.Link = strings.TrimRight(reg.ReplaceAllString(p.Link, ""), "/") + "/page/"
	return p
}
