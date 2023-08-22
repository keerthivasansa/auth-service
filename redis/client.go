package redis

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var _client *redis.Client

func GetClient() *redis.Client {
	if _client != nil {
		return _client
	}
	address := os.Getenv("REDIS_ADDR")
	db := os.Getenv("REDIS_DB")
	password := os.Getenv("REDIS_PASS")
	db_index, err := strconv.Atoi(db)
	if err != nil {
		db_index = 0
	}
	_client = redis.NewClient(&redis.Options{
		Addr:     address,
		DB:       db_index,
		Password: password,
	})
	return _client
}
