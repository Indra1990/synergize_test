package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func ProviderCacheRedis(ctx context.Context) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("cache.redis.host") + ":" + viper.GetString("cache.redis.port"),
		Password: viper.GetString("cache.redis.password"),
		DB:       viper.GetInt("cache.redis.database"),
	})

	res, resErr := redisClient.Ping(ctx).Result()
	if resErr != nil {
		panic(resErr)
	}

	fmt.Println("redis success connect : ", res, "OK")

	return redisClient
}
