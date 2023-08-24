/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package usecase

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/msantosfelipe/go-bank-transfer/config"
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type loginUsecase struct {
	repository domain.LoginRepository
}

func NewLoginUsecase(repository domain.LoginRepository) domain.LoginUsecase {
	return &loginUsecase{
		repository: repository,
	}
}

func (uc *loginUsecase) AuthenticateUser(credentials domain.Login) (*domain.JwtToken, error) {
	login, accountId, err := uc.repository.GetLoginAndAccount(credentials.Cpf)
	if err != nil {
		if err.Error() == domain.ErrNoRowsInResultSet.Error() {
			logrus.Error("invalid password")
			return nil, domain.ErrInvalidLogin
		}
		return nil, err
	}

	if ok := doPasswordsMatch(login.Secret, credentials.Secret); !ok {
		logrus.Error("invalid password")
		return nil, domain.ErrInvalidLogin
	}

	jwtToken, err := generateToken(accountId)
	if err != nil {
		logrus.Error("error generatin token - ", err)
		return nil, err
	}

	return &domain.JwtToken{
		Token: jwtToken,
	}, nil
}

func doPasswordsMatch(hashedPassword, requestPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(requestPassword))
	return err == nil
}

func generateToken(accountId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account_origin_id": accountId,
		"exp":               time.Now().Add(time.Minute * time.Duration(config.ENV.JwtTokenExpMinutes)).Unix(),
	})
	return token.SignedString([]byte(config.ENV.JwtTokenSecret))
}
