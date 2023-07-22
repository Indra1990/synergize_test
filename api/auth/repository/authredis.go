package repository

import "context"

type AuthRepositoryRedis interface {
	StoreToken(ctx context.Context, keyToken string, token string) (err error)
	RemoveToken(ctx context.Context, keyToken string) (err error)
}
