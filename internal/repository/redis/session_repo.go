package repository

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type SessionRepo struct {
	cache *redis.Client
	ttl   time.Duration
}

func NewSessionRepo(r *redis.Client, ttl time.Duration) *SessionRepo {
	return &SessionRepo{
		cache: r,
		ttl:   ttl,
	}
}

func (repo *SessionRepo) key(chatID int64) string {
	return fmt.Sprintf("session:%d", chatID)
}

//func (repo *SessionRepo) Save(session *model.Session) error {
//	data, err := json.Marshal(session)
//	if err != nil {
//		return err
//	}
//	return repo.cache.Set(context.Background(), repo.key(session.ChatID), data, repo.ttl)
//}
//
//func (repo *SessionRepo) Load(chatID int64) (*model.Session, error) {
//	raw, err := repo.cache.Get(repo.key(chatID))
//	if err != nil {
//		return nil, err
//	}
//	var s model.Session
//	if err := json.Unmarshal(raw, &s); err != nil {
//		return nil, err
//	}
//	return &s, nil
//}
//
//func (repo *SessionRepo) Delete(chatID int64) error {
//	return repo.cache.Del(repo.key(chatID))
//}
