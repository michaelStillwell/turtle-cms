package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/michaelStillwell/turtle-cms/pkg"
)

const (
	userIDKey = "userID"
)

func GetUserIDKey() string {
	return userIDKey
}

func UserIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		u, ok := ctx.Get("user").(*jwt.Token)
		if !ok {
			return next(ctx)
		}

		claims := u.Claims.(*pkg.JWTClaims)
		if claims.UserID > 0 {
			ctx.Set(GetUserIDKey(), int64(claims.UserID))
		}

		return next(ctx)
	}
}
