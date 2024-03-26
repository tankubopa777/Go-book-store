package jwtauth

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

type (
	AuthFactory interface {
		SignToken() string
	}

	Claims struct {
		UserId string `json:"user_id"`
		RoleCode int `json:"role_code"`
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

func NewRefreshToken(secret string, expiredAt int64, claims *Claims) AuthFactory {
	return &refreshToken{
		authConcrete: &authConcrete{
			Secrect: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer: "tankubopabook.com",
					Subject: "refresh-token",
					Audience: []string{"tankubopabook.com"},
					ExpiresAt: jwtTimeDurationCal(expiredAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt: jwt.NewNumericDate(now()),
				},
			},
		},
}
}

func ReloadToken(secret string, expiredAt int64, claims *Claims) string{
	obj := &refreshToken{
		authConcrete: &authConcrete{
			Secrect: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer: "tankubopabook.com",
					Subject: "refresh-token",
					Audience: []string{"tankubopabook.com"},
					ExpiresAt: jwtTimeRepeatAdapter(expiredAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt: jwt.NewNumericDate(now()),
				},
			},
		},
	}

	return obj.SignToken()
}

func NewApiKey(secret string, ) AuthFactory {
	return &apiKey{
		authConcrete: &authConcrete{
			Secrect: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: &Claims{},
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer: "tankubopabook.com",
					Subject: "api-key",
					Audience: []string{"tankubopabook.com"},
					ExpiresAt: jwtTimeDurationCal(31560000),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt: jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func ParseToken(secret string, tokenString string) (*AuthMapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthMapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("error: token format is invalid")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("error: token is expired")
		} else {
			return nil ,errors.New("error: token is invalid")
		}
	}

	if claims, ok := token.Claims.(*AuthMapClaims); ok {
		return claims, nil
	} else {
	return nil, errors.New("error: token is invalid")
	}
}

// Apikey generator
var apiKeyInstant string
var once sync.Once

func SetApiKey(secret string) {
	once.Do(func () {
		apiKeyInstant = NewApiKey(secret).SignToken()
	})
}

func SetApiKeyInContext(pctx *context.Context) {
	*pctx = metadata.NewOutgoingContext(*pctx, metadata.Pairs("api-key", apiKeyInstant))
}