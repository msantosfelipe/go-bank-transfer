/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package usecase

import (
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/msantosfelipe/go-bank-transfer/infrastructure/crypto"
	"github.com/msantosfelipe/go-bank-transfer/infrastructure/jwt"
	"github.com/sirupsen/logrus"
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

	if ok := crypto.DoPasswordsMatch(login.Secret, credentials.Secret); !ok {
		logrus.Error("invalid password")
		return nil, domain.ErrInvalidLogin
	}

	jwtToken, err := jwt.GenerateToken(accountId)
	if err != nil {
		logrus.Error("error generatin token - ", err)
		return nil, err
	}

	return &domain.JwtToken{
		Token: jwtToken,
	}, nil
}
