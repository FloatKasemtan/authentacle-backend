package redis

import (
	"github.com/floatkasemtan/authentacle-service/init/config"
	"github.com/redis/go-redis/v9"
)

var Instance *redis.Client

func Initialize() {
	Instance = redis.NewClient(&redis.Options{
		Addr:     config.C.RedisHost,
		Password: config.C.RedisPassword,
		DB:       config.C.RedisDb,
	})
}
