package redis

import (
	"encoding/json"
	"github.com/Davidmnj91/myrents/pkg/domain/auth"
	domain "github.com/Davidmnj91/myrents/pkg/domain/types"
)

func ToRedis(session auth.Session) ([]byte, error) {
	redisSession := &Session{
		UserUUID: session.UserUUID.String(),
		Username: session.Username,
	}

	return json.Marshal(redisSession)
}

func ToDomain(sessionStr []byte) (auth.Session, error) {
	var redisSession Session

	err := json.Unmarshal(sessionStr, &redisSession)
	if err != nil {
		return auth.Session{}, err
	}

	uuid, err := domain.Parse(redisSession.UserUUID)
	if err != nil {
		return auth.Session{}, err
	}

	return auth.Session{
		UserUUID: uuid,
		Username: redisSession.Username,
	}, nil
}
