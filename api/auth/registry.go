package auth

import (
	"context"
	"net/http"
	"synergize/api/auth/usecase"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

type AuthHttpRouterRegistry struct {
	authService usecase.AuthService
}

func NewAuthHttpRouterRegistry(authService usecase.AuthService) *AuthHttpRouterRegistry {
	return &AuthHttpRouterRegistry{authService: authService}
}

func (a *AuthHttpRouterRegistry) Register(w http.ResponseWriter, r *http.Request) {
	cmd := usecase.RegisterRequest{}
	if decodeErr := render.DecodeJSON(r.Body, &cmd); decodeErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, decodeErr.Error())
		return
	}

	if err := a.authService.Register(cmd); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, "user register success")
}

func (a *AuthHttpRouterRegistry) Login(w http.ResponseWriter, r *http.Request) {
	cmd := usecase.LoginRequest{}
	ctx := context.Background()

	if decodeErr := render.DecodeJSON(r.Body, &cmd); decodeErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, decodeErr.Error())
		return
	}

	token, tokernErr := a.authService.Login(ctx, cmd)
	if tokernErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, tokernErr.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, token)
}

func (a *AuthHttpRouterRegistry) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	_, claims, claimsErr := jwtauth.FromContext(r.Context())

	if claimsErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, claimsErr.Error())
		return
	}

	if err := a.authService.Logout(ctx, claims); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)

	render.JSON(w, r, "success logout")
}
