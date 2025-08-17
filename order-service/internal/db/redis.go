package db

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedisDB(connectionString string) (*redis.Client, error) {
	options, err := redis.ParseURL(connectionString)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(options)
	ctx := context.Background()

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
