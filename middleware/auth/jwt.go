package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
)

var (
	accessExp  = jwt.NewNumericDate(time.Now().Add(time.Hour * 1))
	refreshExp = jwt.NewNumericDate(time.Now().Add(time.Hour * 120))
	secretKey  = []byte(os.Getenv("JWT_KEY"))
)

type Token struct {
	Token string
	ExpAt time.Time
}

func GenerateAccessToken(userID string, email string) (*Token, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"email":  email,
		"sub":    "ACCESS_TOKEN",
		"exp":    accessExp,
		"iat":    jwt.NewNumericDate(time.Now()),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	at := new(Token)
	at.Token = token
	at.ExpAt = accessExp.Time

	return at, nil
}

func GenerateRefreshToken(email string) (*Token, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"sub":   "REFRESH_TOKEN",
		"exp":   refreshExp,
		"iat":   jwt.NewNumericDate(time.Now()),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	rt := new(Token)
	rt.Token = token
	rt.ExpAt = refreshExp.Time

	return rt, nil
}

func VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Println("jwt.VerifyToken() error: %v", err)
		return nil, err
	}

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, jwt.ErrTokenExpired
	}

	return claims, nil
}
