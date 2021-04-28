package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"share/internal/database"
	"time"
)

func Statistics() gin.HandlerFunc {
	return func(c *gin.Context) {
		if database.Redis != nil {
			// uv
			uvKey := fmt.Sprintf("uv:%s", time.Now().Format("20060102"))
			database.Redis.PFAdd(uvKey, c.ClientIP())
			// request ip
			database.Redis.ZAdd("request_ips", redis.Z{float64(time.Now().Unix()), c.ClientIP()})
			// pv
			pvKey := fmt.Sprintf("pv:%s", time.Now().Format("20060102"))
			database.Redis.Incr(pvKey)
			database.Redis.Incr("pv_total")
		}
		c.Next()
	}
}
