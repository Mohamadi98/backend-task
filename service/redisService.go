package service

import (
	"backend-task/redis"
	"context"
)

func SetKey(key string, value int) error {
	err := redis.RedisClient.Set(context.Background(), key, value, 0).Err()

	if err != nil {
		return err
	}

	return nil
}

func IncrementKey(key string) error {
	err := redis.RedisClient.Incr(context.Background(), key).Err()

	if err != nil {
		return err
	}

	return nil
}

func GetKeyInt(key string) (uint, error) {
	val, err := redis.RedisClient.Get(context.Background(), key).Int()

	if err != nil {
		return uint(val), err
	}

	return uint(val), nil
}

func DeleteKey(key string) error {
	err := redis.RedisClient.Del(context.Background(), key).Err()

	if err != nil {
		return err
	}

	return nil
}
