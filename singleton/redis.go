package singleton

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/go-redis/redis"
	"sync"
)

var redisInstance *redis.Client
var redisOnce sync.Once
//https://godoc.org/github.com/go-redis/redis#pkg-examples
func RedisSingleton() *redis.Client {
	redisOnce.Do(func() {
		config, _ := config.NewConfig("ini", "./conf/redis.conf")
		redisHost := config.String("host")
		redisPort := config.String("port")
		redisPassword := config.String("password")
		redisInstance = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s",redisHost,redisPort),
			Password:  redisPassword,
			DB:       4,
		})
	})
	return redisInstance
}
