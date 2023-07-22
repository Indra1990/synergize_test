package repositoryredis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type RepositoryRedis struct {
	cacheRds *redis.Client
}

func NewRepositoryRedis(cacheRds *redis.Client) *RepositoryRedis {
	return &RepositoryRedis{
		cacheRds: cacheRds,
	}
}

func (r *RepositoryRedis) StoreToken(ctx context.Context, keyToken string, token string) (err error) {

	rdsCmd := r.cacheRds.Set(ctx, keyToken, token, viper.GetDuration("redis.expired")*time.Minute)

	if rdsCmd.Err() != nil {
		err = rdsCmd.Err()
	}
	return
}

func (r *RepositoryRedis) RemoveToken(ctx context.Context, key string) (err error) {

	rdsDelKey := r.cacheRds.Del(ctx, key).Err()

	if rdsDelKey == redis.Nil {
		err = rdsDelKey
		return
	}

	if rdsDelKey != nil {
		err = rdsDelKey
		return
	}

	return
}
