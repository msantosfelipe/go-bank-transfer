package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/msantosfelipe/go-bank-transfer/config"
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/sirupsen/logrus"
)

const (
	accountOriginId = "account_origin_id"
	tokenExpiration = "exp"
)

func GenerateToken(accountId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		accountOriginId: accountId,
		tokenExpiration: time.Now().Add(time.Minute * time.Duration(config.ENV.JwtTokenExpMinutes)).Unix(),
	})

	tokenSigned, err := token.SignedString([]byte(config.ENV.JwtTokenSecret))
	if err != nil {
		return "", err
	}

	return "Bearer " + tokenSigned, nil
}

func IsTokenValid(context *gin.Context) bool {
	_, err := extractToken(context)
	return err == nil
}

func ExtractAccountOriginId(context *gin.Context) (string, error) {
	token, err := extractToken(context)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id := fmt.Sprintf("%s", claims[accountOriginId])
		return id, nil
	}

	return "", domain.ErrUserNotAuthorized
}

func extractToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := context.Query("token")

	if tokenString == "" {
		bearerToken := context.Request.Header.Get("Authorization")
		if len(strings.Split(bearerToken, " ")) == 2 {
			tokenString = strings.Split(bearerToken, " ")[1]
		}
	}

	if tokenString == "" {
		logrus.Error("invalid token")
		return nil, domain.ErrUserNotAuthorized
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.ENV.JwtTokenSecret), nil
	})
	if err != nil {
		logrus.Error("invalid token - ", err)
		return nil, err
	}

	return token, err
}
