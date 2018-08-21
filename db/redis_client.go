package db

import (
	"github.com/go-redis/redis"
)

// RedisClient return an instance of redis.Client
func RedisClient(addr string, password string, db int) (*redis.Client, error) {
	var client *redis.Client
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return client, err
	}
	return client, nil
}
