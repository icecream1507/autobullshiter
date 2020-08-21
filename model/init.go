package model

import (
	"os"

	"github.com/go-redis/redis"
)

// DB 数据库连接
var DB *redis.Client

// Init 初始化数据库连接
func Init() {
	db := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PWD"),
		DB:       0,
	})

	_, err := db.Ping().Result()

	if err != nil {
		panic(err)
	}

	DB = db
}
