package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Connect() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := RedisClient.Ping(context.Background()).Result()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected To Redis Successfuly!: %v\n", pong)
}
