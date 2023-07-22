package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-redis/redis/v8"
)

func CheckTokenInRedis(cacheRds *redis.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			_, claims, claimsErr := jwtauth.FromContext(r.Context())
			if claimsErr != nil {
				http.Error(w, claimsErr.Error(), http.StatusUnauthorized)
				return
			}

			userId := fmt.Sprintf("%v", claims["user_id"])
			rdsKeyUser := fmt.Sprintf("%s_%v", "token", userId)

			rdsUsr, rdsUsrErr := cacheRds.Get(ctx, rdsKeyUser).Result()
			if rdsUsrErr == redis.Nil {
				// key does not exist
				http.Error(w, "token expired", http.StatusUnauthorized)
				return
			}

			if jwtauth.TokenFromHeader(r) != rdsUsr {
				http.Error(w, "token mismatch", http.StatusUnauthorized)
				return
			}

			if rdsUsrErr != nil {
				http.Error(w, rdsUsrErr.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hfn)
	}
}
