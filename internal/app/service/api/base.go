package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func SuccessPagination(c *gin.Context, data interface{}, p *Pagination) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data: map[string]interface{}{
			"list":       data,
			"pagination": p,
		},
	})
}

func Error(c *gin.Context, code int, r *Response) {
	c.JSON(code, r)
}

func ErrorRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    "",
	})
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}

func NewPagination(c *gin.Context) *Pagination {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return &Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}

// TODO: 优化报错信息
//func ParseErrorMessage(err error, obj interface{}) (msg string) {
//
//	s := reflect.TypeOf(obj).Elem()
//	objName := s.Name()
//	fmt.Println(s.Name())
//	for  i := 0; i < s.NumField(); i++ {
//		fmt.Println(s.Field(i).Name)
//		st := s.Field(i).Tag
//		key := st.Get("form")
//		message := st.Get("message")
//		if key == "" || message == "" {
//			continue
//		}
//		fmt.Println(objName + "." + s.Field(i).Name)
//		fmt.Println(s.Field(i).Tag.Get("form"), s.Field(i).Tag.Get("message"))
//	}
//
//	return
//}
