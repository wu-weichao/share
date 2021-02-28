package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"share/internal/app/middleware"
	"share/internal/app/service/api"
	"share/internal/models"
)

func Login(c *gin.Context) {
	var user models.User
	err := models.DB.Where("id = ?", 1).First(&user).Error
	if err != nil {
		fmt.Printf("user models error %+v\n", err)
	}
	// login success
	jwt := middleware.Jwt{}
	token, expire, err := jwt.Create(&user)
	if err != nil {
		fmt.Pintf("login error %+v\n", err)
	}
	c.JSON(http.StatusOK, api.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data: map[string]interface{}{
			"token":  token,
			"expire": expire,
		},
	})
}
