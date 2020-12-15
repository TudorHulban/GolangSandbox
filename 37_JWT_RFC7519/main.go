package main

// see https://vuejsdevelopers.com/2019/04/15/api-security-jwt-json-web-tokens/

import (
	"fmt"

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

// NewClaims - see https://jwt.io
func NewClaims(expirationSecs int64) *JWTClaims {
	instance := new(JWTClaims)
	instance.TokenExpiration = expirationSecs

	return instance
}

func NewToken(claims *JWTClaims, signingMethod int, secret string) (string, error) {
	var tok = new(jwt.Token)

	switch signingMethod {
	case 1:
		tok = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	default:
		tok = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	}

	stringToken, err := tok.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return stringToken, nil
}

// ValidateToken - validates and returns JWT token
func ValidateToken(stringToken string) (*jwt.Token, error) {
	tok, err := jwt.ParseWithClaims(stringToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, err := token.Method.(*jwt.SigningMethodHMAC); err == false {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return "", nil
	})
	return tok, err
}

func DecodeClaims(stringToken string) (*JWTCustomClaims, error) {
	var custom = new(JWTCustomClaims)

	tok, err := ValidateToken(stringToken)
	if err.Error() != "key is of invalid type" {
		return nil, err
	}

	claims, _ := tok.Claims.(*JWTClaims)
	custom.TokenExpiration = claims.TokenExpiration
	return custom, nil
}

func main() {
	c := NewClaims(60)

	t, err := NewToken(c, 1, "xxx")
	fmt.Println(t, err)

	x, _ := DecodeClaims(t)
	fmt.Println(x.TokenExpiration)
}
