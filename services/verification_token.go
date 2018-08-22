package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
)

// VerificationRedisNS is the namespace to store tokens in redis in the
// following format.
// "verification:<id>": <verificationCode>
const VerificationRedisNS = "verification"

// VerificationTTL is how long the token lives
const VerificationTTL = 10 * time.Minute

func getKey(id string) string {
	return fmt.Sprintf(`%s:%s`, VerificationRedisNS, id)
}

// CreateVerificationToken makes a temporary token inside redis to be used to
// verify the email account.
func CreateVerificationToken(redisClient *redis.Client, id string) (string, error) {
	verificationCode := fmt.Sprintf(
		`%d%d%d%d`,
		rand.Intn(9),
		rand.Intn(9),
		rand.Intn(9),
		rand.Intn(9),
	)

	err := redisClient.Set(
		getKey(id),
		verificationCode,
		VerificationTTL,
	).Err()
	if err != nil {
		return verificationCode, err
	}

	return verificationCode, nil
}

// GetVerificationToken returns token from redis.
func GetVerificationToken(redisClient *redis.Client, id string) (string, error) {
	return redisClient.Get(
		getKey(id),
	).Result()
}
