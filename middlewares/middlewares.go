package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/context"
	"github.com/vaibhavchalse99/api"
	"github.com/vaibhavchalse99/config"
	"github.com/vaibhavchalse99/users"
)

func getAuthTokenFromHeader(r *http.Request) (token string, err error) {
	bearerToken := r.Header.Get("Authorization")
	arr := strings.Split(bearerToken, " ")
	if len(arr) == 2 {
		token = arr[1]
		return
	}
	return token, ErrTokenNotPresent
}

func verifyJwt(bearerToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(config.SecretHashKey()), nil
	})
	return token, err
}

func IsLoggedIn(next http.HandlerFunc, service users.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tokenValue, err := getAuthTokenFromHeader(r)
		fmt.Println(tokenValue)
		if err == ErrTokenNotPresent {
			api.Error(rw, http.StatusNotFound, api.Response{Message: ErrTokenNotPresent.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		token, err := verifyJwt(tokenValue)
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		if token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok {
				userId := claims["user_id"].(string)
				response, err := service.GetById(r.Context(), userId)
				if err == users.ErrUserNotExist {
					api.Error(rw, http.StatusNotFound, api.Response{Message: users.ErrUserNotExist.Error()})
					return
				}
				if err != nil {
					api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
					return
				}
				context.Set(r, "user", response.User)
				next(rw, r)
			}
		} else {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: ErrTokenNotPresent.Error()})
		}
	})
}
