package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/saipulmuiz/mnc-test-tahap2/config"
)

func VerifyToken(tokenString string) (*Claims, error) {
	errResponse := errors.New("Token-Invalid")

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTConfig.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errResponse
	}

	claims := token.Claims.(*Claims)

	return claims, nil
}

type Claims struct {
	UserID      string `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID, phoneNumber string) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.JWTConfig.AccessExpiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTConfig.AccessSecret))
}

func GenerateRefreshToken(userID string) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.JWTConfig.RefreshExpiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTConfig.RefreshSecret))
}
