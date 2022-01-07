package db

import "github.com/go-redis/redis"

var redisDB *redis.Client

func init() {
	redisDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6699",
		Password: "",
		DB: 0,
	})
}

func GetRedisDB() *redis.Client {
	return redisDB
}

