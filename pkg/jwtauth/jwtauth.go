package jwtauth

import (
	"math"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	AuthFactory interface {
		SignToken() string
	}

	Claims struct {
		Id string `json:"id"`
		RoleCode string `json:"role_code"`
	}

	AuthMapClaims struct {
		*Claims
		jwt.RegisteredClaims
	}

	authConcrete struct {
		Secrect []byte
		Claims *AuthMapClaims `json:"claims"`
	}

	accessToken struct {*authConcrete}
	refreshToken struct {*authConcrete}
	apiKey struct {*authConcrete}
)

func (a *authConcrete) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.Claims)
	ss, _ := token.SignedString(a.Secrect)
	return ss
}

func now() time.Time{
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc)
}

func jwtTimeDurationCal(t int64) *jwt.NumericDate {
    return jwt.NewNumericDate(now().Add(time.Duration(t * int64(math.Pow10(9)))))
}

func jwtTimeRepeatAdapter(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Unix(t, 0))
}

func NewAccessToken(secret string, expiredAt int64, claims *Claims) AuthFactory {
	return &accessToken{
		authConcrete: &authConcrete{
			Secrect: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer: "tankubopabook.com",
					Subject: "access-token",
					Audience: []string{"tankubopabook.com"},
					ExpiresAt: jwtTimeDurationCal(expiredAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt: jwt.NewNumericDate(now()),
				},
			},
		},
}
}