package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedis(host string, port uint16, username string, password string, dbname int) (redisClient *redis.Client, err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Username: username,
		Password: password,
		DB:       dbname,
	})

	err = redisClient.Ping(context.Background()).Err()

	return
}
