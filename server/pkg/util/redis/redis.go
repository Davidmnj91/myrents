package util

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type RedisConfiguration struct {
	host string
	port int
	pass string
	db   int
}

func NewRedisConfiguration(host string, port int, pass string, db int) *RedisConfiguration {
	return &RedisConfiguration{host, port, pass, db}
}

// ConnectRedis to a database handle from a redis configuration.
func ConnectRedis(configuration *RedisConfiguration) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", configuration.host, configuration.port),
		Password: configuration.pass,
		DB:       configuration.db,
	})

	err := client.Ping(context.TODO()).Err()

	if err != nil {
		return nil, err
	}

	return client, nil
}
