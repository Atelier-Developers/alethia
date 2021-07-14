package auth

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

type RedisService struct {
	Auth   AuthInterface
	Client *redis.Client
}

func NewRedisDB(host, port, password string, ctx context.Context) (*RedisService, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	return &RedisService{
		Auth:   NewAuth(redisClient, ctx),
		Client: redisClient,
	}, nil
}
