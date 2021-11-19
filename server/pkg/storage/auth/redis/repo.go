package redis

import (
	"context"
	"errors"
	"github.com/Davidmnj91/myrents/pkg/auth/domain"
	"github.com/Davidmnj91/myrents/pkg/types"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisRepository struct {
	client *redis.Client
	ttl    int64
}

func NewRepository(client *redis.Client, ttl int64) domain.Repository {
	return &redisRepository{client, ttl}
}

func (r *redisRepository) GetSession(ctx context.Context, uuid types.UUID) (domain.Session, error) {
	uuidStr := uuid.String()
	sessionStr, err := r.client.Get(ctx, uuidStr).Result()

	if err == redis.Nil {
		return domain.Session{}, errors.New("key does not exists")
	}
	if err != nil {
		return domain.Session{}, err
	}

	session, err := ToDomain([]byte(sessionStr))
	if err != nil {
		return domain.Session{}, err
	}
	return session, nil

}

func (r *redisRepository) CreateSession(ctx context.Context, session domain.Session) error {
	sessionStr, err := ToRedis(session)
	if err != nil {
		return err
	}

	timeout := time.Millisecond * time.Duration(r.ttl)
	err = r.client.Set(ctx, session.UserUUID.String(), sessionStr, timeout).Err()

	return err
}

func (r *redisRepository) RefreshSession(ctx context.Context, session domain.Session) error {
	timeout := time.Millisecond * time.Duration(r.ttl)

	_, err := r.client.Expire(ctx, session.UserUUID.String(), timeout).Result()

	return err
}

func (r *redisRepository) RemoveSession(ctx context.Context, userUUID types.UUID) error {
	return r.client.Del(ctx, userUUID.String()).Err()
}
