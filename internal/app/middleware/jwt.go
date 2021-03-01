package middleware

import (
	"net/http"
	"share/internal/app/service/api"

	//"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"share/configs"
	"share/internal/models"
	"time"
)

type Jwt struct {
	Token  string
	Secret string
	Expire int64
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

func NewJwt(token string) *Jwt {
	return &Jwt{
		Token:  token,
		Secret: configs.Jwt.Secret,
	}
}

func (j *Jwt) Create(user *models.User) (*Jwt, error) {
	now := time.Now()
	expire := now.Add(time.Duration(configs.Jwt.Ttl) * time.Second).Unix()
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
	token, err := tokenClaims.SignedString([]byte(j.Secret))
	j.Token = token
	j.Expire = expire
	return j, err
}

func (j *Jwt) Parse() (*CustomClaims, error) {
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(j.Token, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(j.Secret), nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get header Authorization: [Bearer token]
		token := c.GetHeader("Authorization")
		flag := true
		msg := ""
		if token == "" {
			flag = false
			msg = "Failed to authenticate because of bad credentials or an invalid authorization header"
		} else {
			token = token[7:]
			// validate token
			j := NewJwt(token)
			claims, err := j.Parse()
			if err != nil {
				flag = false
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorMalformed:
					msg = "Token format error"
				case jwt.ValidationErrorExpired:
					msg = "Token has expired"
				case jwt.ValidationErrorNotValidYet:
					msg = "Token verification error"
				default:
					msg = "Unable to authenticate with invalid token"
				}
			} else {
				c.Set("user", claims)
			}
		}
		// token validate failed
		if !flag {
			c.JSON(http.StatusUnauthorized, api.Response{
				Code:    http.StatusUnauthorized,
				Message: msg,
				Data:    "",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
