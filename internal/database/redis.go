package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"share/configs"
	"sync"
)

var once sync.Once

var Redis *redis.Client

func init() {
	once.Do(func() {
		r := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", configs.Redis.Host, configs.Redis.Port),
			Password: configs.Redis.Password,
			DB:       configs.Redis.Database,
		})
		_, err := r.Ping().Result()
		if err != nil {
			fmt.Printf("redis connection error: %+v\n", err)
			return
		}
		Redis = r
	})
}
