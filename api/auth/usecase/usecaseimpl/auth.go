package usecaseimpl

import (
	"context"
	"fmt"
	"synergize/api/auth/repository"
	"synergize/api/auth/usecase"
	"synergize/entity"
	"synergize/utils"

	"github.com/go-chi/jwtauth/v5"
)

type AuthService struct {
	authRepo    repository.AuthRepository
	tokenAuth   *jwtauth.JWTAuth
	authRepoRds repository.AuthRepositoryRedis
}

func NewAuthService(authRepo repository.AuthRepository, tokenAuth *jwtauth.JWTAuth, authRepoRds repository.AuthRepositoryRedis) *AuthService {
	return &AuthService{
		authRepo:    authRepo,
		tokenAuth:   tokenAuth,
		authRepoRds: authRepoRds,
	}
}

func (a *AuthService) Register(cmd usecase.RegisterRequest) (err error) {
	if err = validateRegister(cmd); err != nil {
		return
	}

	pass, passErr := utils.GenerateHashPassword(cmd.Password)
	if passErr != nil {
		err = passErr
		return
	}

	user := entity.User{
		Username:    cmd.Username,
		Email:       cmd.Email,
		PhoneNumber: cmd.PhoneNumber,
		Password:    pass,
	}

	if err = a.authRepo.CreateUser(&user); err != nil {
		return
	}

	return
}

func (a *AuthService) Login(ctx context.Context, cmd usecase.LoginRequest) (dto usecase.LoginResponse, err error) {
	user, userErr := a.authRepo.FindByUserEmail(cmd.Email)
	if userErr != nil {
		err = userErr
		return
	}

	if err = utils.PasswordCheck(cmd.Password, user.Password); err != nil {
		return
	}

	claims := map[string]interface{}{
		"user_id": user.ID,
		"name":    user.Username,
	}

	createToken, createTokenErr := a.createToken(ctx, claims)
	if createTokenErr != nil {
		err = createTokenErr
		return
	}

	dto = usecase.LoginResponse{
		AccessToken: createToken,
	}

	return
}

func (a *AuthService) Logout(ctx context.Context, claims map[string]interface{}) (err error) {
	userId := claims["user_id"]
	rdsKeyUser := fmt.Sprintf("%s_%v", "token", userId)

	if err = a.authRepoRds.RemoveToken(ctx, rdsKeyUser); err != nil {
		return
	}

	return
}

func (a *AuthService) createToken(ctx context.Context, claims map[string]interface{}) (string, error) {
	jwtauth.SetIssuedNow(claims)
	_, token, tokenErr := a.tokenAuth.Encode(claims)
	if tokenErr != nil {
		return "", tokenErr
	}

	userId := claims["user_id"]
	rdsKeyUser := fmt.Sprintf("%s_%v", "token", userId)

	//token store to redis
	if err := a.authRepoRds.StoreToken(ctx, rdsKeyUser, token); err != nil {
		return "", err
	}

	return token, nil
}
