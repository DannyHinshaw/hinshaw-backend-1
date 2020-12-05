package cache

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type MockService struct {
	Client redis.Cmdable
}

// Create a new redis mock for testing.
func NewMockRedis() *MockService {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return &MockService{Client: client}
}

// Get all keys in redis
func (s *MockService) GetAllRedisKeys() []string {
	var allKeys []string
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = s.Client.Scan(cursor, "", 0).Result()
		if err != nil {
			panic(err)
		}

		allKeys = append(allKeys, keys...)
		if cursor == 0 {
			break
		}
	}

	return allKeys
}

// Set a string key in redis.
func (s *MockService) SetKeyStringRedis(key string, val string) error {
	err := s.Client.Set(key, val, 0).Err()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// Get a string key in redis.
func (s *MockService) GetKeyInRedis(key string) string {
	val := s.Client.Get(key).Val()
	return val
}

// Check if a key is in redis.
func (s *MockService) IsKeyInRedis(key string) bool {
	scrapeKey, err := s.Client.Get(key).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		return false
	} else if scrapeKey == "" {
		return false
	}

	return true
}

// Util function to remove a key/val from redis.
func (s *MockService) ExpireKey(key string) {
	s.Client.Set(key, "", time.Millisecond)
}
