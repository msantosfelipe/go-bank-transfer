/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package usecase

import (
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
	login, err := uc.repository.GetLoginByCpf(credentials.Cpf)
	if err != nil {
		if err == domain.ErrNoRowsInResultSet {
			return nil, domain.ErrInvalidLogin
		}
		return nil, err
	}

	if ok := doPasswordsMatch(login.Secret, credentials.Secret); !ok {
		logrus.Error("invalid password")
		return nil, domain.ErrInvalidLogin
	}

	// TODO criar JWT

	return &domain.JwtToken{
		Token: "jwt",
	}, nil
}

func doPasswordsMatch(hashedPassword, requestPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(requestPassword))
	return err == nil
}
