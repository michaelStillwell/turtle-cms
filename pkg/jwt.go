package pkg

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	jwtKey      = "secret"
	cookieName  = "session"
	exipiration = time.Hour * 72
)

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID int
	Tenant string
}

func GetJwtKey() []byte {
	return []byte(jwtKey)
}

func GetCookieName() string {
	return cookieName
}

func GenerateTokenAndCookie(ctx echo.Context, userId int, tenant string) error {
	token, err := generateJWT(userId, tenant)
	if err != nil {
		return err
	}

	setCookie(ctx, GetCookieName(), token, time.Now().Add(exipiration))
	return nil
}

func RemoveCookie(ctx echo.Context) {
	setCookie(ctx, GetCookieName(), "", time.Now().Add(-time.Hour))
}

func generateJWT(userID int, tenant string) (string, error) {
	claims := &JWTClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exipiration)),
		},
		userID,
		tenant,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetJwtKey())
}

func setCookie(ctx echo.Context, name, token string, expiration time.Time) {
	ctx.SetCookie(&http.Cookie{
		Name:     name,
		Value:    token,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true,
	})
}
