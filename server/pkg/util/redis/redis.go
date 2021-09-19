package util

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisConfiguration struct {
	address string
	pass    string
	db      int
}

func NewRedisConfiguration(address string, pass string, db int) *RedisConfiguration {
	return &RedisConfiguration{address, pass, db}
}

// ConnectRedis to a database handle from a redis configuration.
func ConnectRedis(configuration *RedisConfiguration) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     configuration.address,
		Password: configuration.pass,
		DB:       configuration.db,
	})

	err := client.Set(context.TODO(), "key", "value", 0).Err()

	if err != nil {
		return nil, err
	}

	return client, nil
}
