package v1

import (
	"github.com/gin-gonic/gin"
	"share/internal/app/middleware"
	"share/internal/app/service/api"
	"share/internal/models"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required" message:"email is required"`
	Password string `form:"password" binding:"required" message:"password is required"`
}

type UserResponse struct {
	ID        uint     `json:"id"`
	CreatedAt int      `json:"created_at"`
	Name      string   `json:"name"`
	Avatar    string   `json:"avatar"`
	Email     string   `json:"email"`
	Phone     string   `json:"phone"`
	Status    int      `json:"status"`
	Roles     []string `json:"roles"`
}

func Login(c *gin.Context) {
	// validate params
	var form LoginRequest
	var err error
	if err = c.ShouldBind(&form); err != nil {
		api.ErrorRequest(c, err.Error())
		return
	}
	// get user
	user, err := models.UserGetByEmail(form.Email)
	if err != nil {
		api.ErrorRequest(c, "Incorrect username or password")
		return
	}
	if models.UserEncodePassword(form.Password) != user.Password {
		api.ErrorRequest(c, "Incorrect username or password")
		return
	}

	// login success
	jwt := middleware.NewJwt("")
	jwt, err = jwt.Create(user)
	if err != nil {
		api.ErrorRequest(c, "Token generation failed")
		return
	}
	api.Success(c, map[string]interface{}{
		"token":  jwt.Token,
		"expire": jwt.Expire,
	})
}

func LoginUserInfo(c *gin.Context) {
	// get login user
	jwtUser, _ := c.Get("user")
	user, err := models.UserGetById(jwtUser.(*middleware.CustomClaims).ID)
	if err != nil {
		api.ErrorRequest(c, "User not exists")
		return
	}
	// get user role
	var roles []string
	if user.Type == models.UserTypeAdmin {
		roles = append(roles, "admin")
	}

	api.Success(c, UserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Name:      user.Name,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Phone:     user.Phone,
		Roles:     roles,
		Status:    user.Status,
	})
}
