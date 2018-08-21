package db

import (
	"github.com/go-redis/redis"
)

// RedisClient return an instance of redis.Client
func RedisClient(addr string, password string, db int) (error, *redis.Client) {
	var client *redis.Client
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return err, client
	}
	return nil, client
}
