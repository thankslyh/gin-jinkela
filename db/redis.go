package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"jinkela/setting"
	"log"
	"time"
)

var redisDB *redis.Client

func SetRedisUp() {
	fmt.Println("init redis=", setting.RedisSetting.Host)
	redisDB = redis.NewClient(&redis.Options{
		Addr: setting.RedisSetting.Host,
		Password: "",
		DB: 0,
	})
	if err := redisDB.Ping().Err(); err != nil {
		log.Fatal(err)
		return
	}
	redisDB.Set("test77777", "7777", time.Millisecond*1000*1000)
	log.Println(redisDB.Get("test77777").Val())
}

func GetRedisDB() *redis.Client {
	return redisDB
}

