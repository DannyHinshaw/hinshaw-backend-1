package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"os"
	"time"
)

type IService interface {
	SetJWTRedis(key string, val string) error
	GetKeyInRedis(key string) string
	IsKeyInRedis(key string) bool
	GetAllRedisKeys() []string
	ExpireKey(key string)
}

type Service struct {
	Client redis.Cmdable
}

// Handle environment assigned redis url.
func getRedisUrl() string {
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		return "redis:6379"
	}

	return redisUrl
}

// Creates a new redis client.
func NewRedisClient() *Service {
	redisUrl := getRedisUrl()
	println("redisUrl::", redisUrl)
	client := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test ping
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong, err)

	return &Service{Client: client}
}

// Get all keys in redis
func (s *Service) GetAllRedisKeys() []string {
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
func (s *Service) SetJWTRedis(jwt string, userId string) error {
	err := s.Client.Set(jwt, userId, time.Hour).Err()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// Get a string key in redis.
func (s *Service) GetKeyInRedis(key string) string {
	val := s.Client.Get(key).Val()
	return val
}

// Check if a key is in redis.
func (s *Service) IsKeyInRedis(key string) bool {
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
func (s *Service) ExpireKey(key string) {
	s.Client.Set(key, "", time.Millisecond)
}
