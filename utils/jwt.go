package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(Data map[string]interface{}, ExpiredAt time.Duration) (string, error) {

	expiredAt := time.Now().Add(time.Duration(time.Second) * ExpiredAt).Unix()

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	claims := jwt.MapClaims{}
	claims["expiredAt"] = expiredAt
	claims["authorization"] = true

	for i, v := range Data {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}
