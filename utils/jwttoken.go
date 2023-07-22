package utils

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

func ProvideJWTAuth() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte(viper.GetString("jwt.jwtKey")), nil)
}
