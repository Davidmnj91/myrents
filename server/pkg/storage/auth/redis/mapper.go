package redis

import (
	"encoding/json"
	"github.com/Davidmnj91/myrents/pkg/auth/domain"
	"github.com/Davidmnj91/myrents/pkg/types"
)

func ToRedis(session domain.Session) ([]byte, error) {
	redisSession := &Session{
		UserUUID: session.UserUUID.String(),
		Username: session.Username,
	}

	return json.Marshal(redisSession)
}

func ToDomain(sessionStr []byte) (domain.Session, error) {
	var redisSession Session

	err := json.Unmarshal(sessionStr, &redisSession)
	if err != nil {
		return domain.Session{}, err
	}

	uuid, err := types.Parse(redisSession.UserUUID)
	if err != nil {
		return domain.Session{}, err
	}

	return domain.Session{
		UserUUID: uuid,
		Username: redisSession.Username,
	}, nil
}
