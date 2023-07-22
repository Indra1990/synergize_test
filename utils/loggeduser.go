package utils

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-chi/jwtauth/v5"
)

func GetUserIdByToken(ctx context.Context) (userid uint, err error) {
	_, claims, claimsErr := jwtauth.FromContext(ctx)
	if claimsErr != nil {
		err = claimsErr
		return
	}

	userIdClaim := claims["user_id"]
	userClaimId := fmt.Sprintf("%v", userIdClaim)
	userId, userIdErr := strconv.ParseUint(userClaimId, 10, 64)
	if userIdErr != nil {
		err = userIdErr
		return
	}

	userid = uint(userId)
	return
}
