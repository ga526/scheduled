package singleton

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

var (
	redisHost = "10.200.124.189"
	redisPort = "6382"
	redisPassword = "bjt123torx"
)

var redisInstance *redis.Client
var redisOnce sync.Once
//https://godoc.org/github.com/go-redis/redis#pkg-examples
func RedisSingleton() *redis.Client {
	redisOnce.Do(func() {
		redisInstance = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s",redisHost,redisPort),
			Password:  redisPassword,
			DB:       4,
		})
	})
	return redisInstance
}
