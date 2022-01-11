package db

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

var redisDB *redis.Client

func init() {
	log.Println("init redis")
	redisDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
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

