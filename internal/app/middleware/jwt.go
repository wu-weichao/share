package middleware

import (
	//"fmt"
	//"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"share/configs"
	"share/internal/models"
	"time"
)

type Jwt struct {
	Secret string
}

type CustomClaims struct {
	jwt.StandardClaims

	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Type   int    `json:"type"`
}

func (j *Jwt) Create(user *models.User) (token string, expire int64, err error) {
	j.Secret = configs.Jwt.Secret

	now := time.Now()
	expire = now.Add(time.Duration(configs.Jwt.Ttl) * time.Second).Unix()
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire,
			IssuedAt:  now.Unix(),
			Issuer:    "share",
			NotBefore: now.Unix() - 1000,
		},
		ID:     user.ID,
		Name:   user.Name,
		Avatar: user.Avatar,
		Email:  user.Email,
		Phone:  user.Phone,
		Type:   user.Type,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString([]byte(j.Secret))
	return
}

//func JwtCheck() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// get header Authorization: [Bearer token]
//		token := c.GetHeader("Authorization")
//		flag := true
//		if token == "" {
//			flag = false
//		} else {
//			token = token[7:]
//			// validate token
//			_, err := jwt.Parse()
//		}
//
//
//		if token == "" {
//			c.Abort()
//		}
//		token = token[7:]
//		fmt.Printf("token %+v\n", token)
//
//		c.Next()
//	}
//}
