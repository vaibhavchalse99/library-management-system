package users

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/vaibhavchalse99/config"
)

func createToken(userId uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	byteSecretKey := []byte(config.SecretHashKey())
	tokenString, err := token.SignedString(byteSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
