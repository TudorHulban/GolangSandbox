package main

// see https://vuejsdevelopers.com/2019/04/15/api-security-jwt-json-web-tokens/
// eltoken needs to handle additions to cache as single point

import (
	"fmt"
	"log"
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
func newClaims(pSecondsAfter int64) *JWTClaims {
	var instance = new(JWTClaims)
	instance.TokenExpiration = time.Now().UnixNano() + int64(pSecondsAfter*time.Second.Nanoseconds())
	log.Println("claims expiration: ", time.Unix(0, instance.TokenExpiration))
	return instance
}

// New - constructor for token string. check with https://jwt.io/
func (t *JWTToken) New(pSecondsAfter int64, pSigningMethod int, pSecret string, pCache *cache.Cache, pItem *cache.Item) (string, error) {
	var tok = new(jwt.Token)
	claims := newClaims(pSecondsAfter)

	switch pSigningMethod {
	case 1:
		tok = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	default:
		tok = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	}

	stringToken, err := tok.SignedString([]byte(pSecret))
	if err != nil {
		return "", err
	}

	pCache.Add(stringToken, pItem) //keep expiration in one place only. here
	return stringToken, nil
}

// Validate - validates and returns JWT token
func (t *JWTToken) Validate(pString JWTToken) (*jwt.Token, error) {
	tok, err := jwt.ParseWithClaims(pString.Token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {

		_, err := token.Method.(*jwt.SigningMethodHMAC)
		if err == false {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return "", nil
	})
	return tok, err
}

func (t *JWTToken) DecodeClaims(pString JWTToken) (*JWTCustomClaims, error) {
	var custom = new(JWTCustomClaims)

	tok, err := t.Validate(pString)
	if err.Error() != "key is of invalid type" {
		return nil, err
	}
	claims, _ := tok.Claims.(*JWTClaims)
	custom.TokenExpiration = claims.TokenExpiration
	return custom, nil
}
