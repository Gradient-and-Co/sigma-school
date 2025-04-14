package redis

import (
	"context"
	"encoding/json"
	"github.com/Gradient-and-Co/sigma-school/internal/adapter/auth/port"
	"github.com/redis/go-redis/v9"
	"time"
)

type SessionStorage struct {
	redisClient *redis.Client
}

func NewSessionStorage(redisClient *redis.Client) *SessionStorage {
	return &SessionStorage{redisClient: redisClient}
}

func (s *SessionStorage) Get(refreshToken string) (port.AuthSession, error) {
	sessionJson, err := s.redisClient.Get(context.Background(), refreshToken).Bytes()
	if err != nil {
		return port.AuthSession{}, err
	}

	var session port.AuthSession
	err = json.Unmarshal(sessionJson, &session)
	if err != nil {
		return port.AuthSession{}, err
	}

	return session, nil
}

func (s *SessionStorage) Put(refreshToken string, session port.AuthSession,
	expireTime time.Duration) error {
	sessionJson, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return s.redisClient.Set(context.Background(), refreshToken, sessionJson, expireTime).Err()
}

func (s *SessionStorage) Delete(refreshToken string) error {
	return s.redisClient.Del(context.Background(), refreshToken).Err()
}
