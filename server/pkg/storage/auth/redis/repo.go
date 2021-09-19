package redis

import (
	"context"
	"errors"
	"github.com/Davidmnj91/myrents/pkg/domain/auth"
	"github.com/Davidmnj91/myrents/pkg/domain/types"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisRepository struct {
	client *redis.Client
	ttl    int64
}

func NewRepository(client *redis.Client, ttl int64) auth.Repository {
	return &redisRepository{client, ttl}
}

func (r *redisRepository) GetSession(ctx context.Context, uuid domain.UUID) (auth.Session, error) {
	sessionStr, err := r.client.Get(ctx, uuid.String()).Result()

	if err == redis.Nil {
		return auth.Session{}, errors.New("key does not exists")
	}
	if err != nil {
		return auth.Session{}, err
	}

	session, err := ToDomain([]byte(sessionStr))
	if err != nil {
		return auth.Session{}, err
	}
	return session, nil

}

func (r *redisRepository) CreateSession(ctx context.Context, session auth.Session) error {
	sessionStr, err := ToRedis(session)
	if err != nil {
		return err
	}

	_, err = r.client.Set(ctx, session.UserUUID.String(), sessionStr, time.Duration(r.ttl)).Result()

	if err != nil {
		return err
	}
	return nil
}

func (r *redisRepository) RefreshSession(ctx context.Context, session auth.Session) error {
	_, err := r.client.Expire(ctx, session.UserUUID.String(), time.Duration(r.ttl)).Result()

	if err != nil {
		return err
	}
	return nil
}

func (r *redisRepository) RemoveSession(ctx context.Context, userUUID domain.UUID) error {
	_, err := r.client.Del(ctx, userUUID.String()).Result()

	if err != nil {
		return err
	}
	return nil
}
