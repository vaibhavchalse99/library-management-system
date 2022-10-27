package users

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/vaibhavchalse99/config"
	"github.com/vaibhavchalse99/db"
)

func createToken(userId uuid.UUID, role db.RoleValue) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	byteSecretKey := []byte(config.SecretHashKey())
	tokenString, err := token.SignedString(byteSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func mapUserData(dbUser db.User) (user User) {
	user.ID = dbUser.ID
	user.Email = dbUser.Email
	user.Name = dbUser.Name
	user.Role = dbUser.Role
	user.Password = ""
	user.CreatedAt = dbUser.CreatedAt
	user.UpdatedAt = dbUser.UpdatedAt
	return
}
