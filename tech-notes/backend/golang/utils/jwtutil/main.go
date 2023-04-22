package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTUtil interface {
	Generate(userUID string) (string, time.Time, error)
	ParseUserUID(tokenString string) (string, error)
	Refresh(tokenString string) (string, time.Time, error)
}

type baseJWTUtil struct {
	jwtKey []byte
}

func NewJWTUtil(jwtKey []byte) JWTUtil {
	return &baseJWTUtil{jwtKey: jwtKey}
}

type Claims struct {
	UserUID string `json:"user_uid"`
	jwt.RegisteredClaims
}

func (b *baseJWTUtil) Generate(userUID string) (string, time.Time, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		UserUID: userUID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(b.jwtKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}

func (b *baseJWTUtil) ParseUserUID(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return b.jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", errors.New("invalid token signature")
		}

		return "", err
	}
	if !token.Valid {
		return "", err
	}

	if time.Until(claims.ExpiresAt.Time) < 30*time.Second {
		return "", errors.New("token already expired")
	}

	return claims.UserUID, nil
}

func (b *baseJWTUtil) Refresh(tokenString string) (string, time.Time, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return b.jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", time.Time{}, errors.New("invalid token signature")
		}

		return "", time.Time{}, err
	}
	if !token.Valid {
		return "", time.Time{}, err
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(b.jwtKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}

func main() {
	jwtUtil := NewJWTUtil([]byte("secret_key"))

	tokenString, expirationTime, err := jwtUtil.Generate("test")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tokenString)
	fmt.Println(expirationTime)
}
