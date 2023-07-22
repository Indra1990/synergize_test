package usecase

import "context"

type AuthService interface {
	Register(cmd RegisterRequest) (err error)
	Login(ctx context.Context, cmd LoginRequest) (dto LoginResponse, err error)
	Logout(ctx context.Context, claims map[string]interface{}) (err error)
}
