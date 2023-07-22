package main

import (
	"context"
	"log"
	"synergize/api"
	"synergize/db"
	utilpkg "synergize/utils"

	"github.com/go-chi/jwtauth/v5"
)

func main() {
	ctx := context.Background()
	conn, dbErr := db.SetupDatabasePgSQLConnection()
	if dbErr != nil {
		log.Fatal("cannot connect to db :", dbErr)

	}

	var tokenAuth *jwtauth.JWTAuth
	tokenAuth = utilpkg.ProvideJWTAuth()

	cacheRds := db.ProviderCacheRedis(ctx)

	server := api.NewServer(conn, tokenAuth, cacheRds)
	serverErr := server.Start(ctx)
	if serverErr != nil {
		log.Fatal("cannot start server :", serverErr)
	}
}
