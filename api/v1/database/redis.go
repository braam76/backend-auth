package database

import (
	"os"

	"github.com/gofiber/storage/redis"
	"github.com/joho/godotenv"
)

var Redis *redis.Storage

func InitRedis() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	Redis = redis.New(redis.Config{
		Host:     "127.0.0.1",
		Port:     6380,
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: 0,
	})
}
