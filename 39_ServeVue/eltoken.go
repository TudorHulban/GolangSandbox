package main

// see https://vuejsdevelopers.com/2019/04/15/api-security-jwt-json-web-tokens/
// eltoken needs to handle additions to cache as single point

import (
	"fmt"
	"time"

	cache "github.com/TudorHulban/cachejwt"
	"github.com/dgrijalva/jwt-go"
)

type JWTCustomClaims struct {
	TokenExpiration int64 `json:"unix"`
}

// JWTClaims - payload to JWT
type JWTClaims struct {
	JWTCustomClaims

	jwt.StandardClaims
}

type JWTToken struct {
	Token string
}

// NewClaims - see https://jwt.io. pExpiration in seconds, please multiply for nano seconds
func newClaims(secondsAfter int64) *JWTClaims {
	return &JWTClaims{
		JWTCustomClaims{
			TokenExpiration: time.Now().UnixNano() + int64(secondsAfter*time.Second.Nanoseconds()),
		},
	}
}

// New - constructor for token string. check with https://jwt.io/
func (t *JWTToken) New(secondsAfter int64, signingMethod int, secret string, cache *cache.Cache, item *cache.Item) (string, error) {
	token := new(jwt.Token)
	claims := newClaims(secondsAfter)

	switch signingMethod {
	case 1:
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	default:
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	}

	stringToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	cache.Add(stringToken, item) //keep expiration in one place only. here
	return stringToken, nil
}

// Validate - validates and returns JWT token
func (t *JWTToken) Validate(s JWTToken) (*jwt.Token, error) {
	return jwt.ParseWithClaims(s.Token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, err := token.Method.(*jwt.SigningMethodHMAC); err == false {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return "", nil
	})
}

func (t *JWTToken) DecodeClaims(s JWTToken) (*JWTCustomClaims, error) {
	tok, err := t.Validate(s)
	if err.Error() != "key is of invalid type" {
		return nil, err
	}

	claims, _ := tok.Claims.(*JWTClaims)
	return &JWTCustomClaims{
		TokenExpiration: claims.TokenExpiration,
	}, nil
}
