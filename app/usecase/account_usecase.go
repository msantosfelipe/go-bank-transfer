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

type accountUsecase struct {
	repository domain.AccountRepository
}

func NewAccountUsecase(repository domain.AccountRepository) domain.AccountUsecase {
	return &accountUsecase{
		repository: repository,
	}
}

func (uc *accountUsecase) CreateAccount(request domain.AccountCreatorRequest) (*domain.AccountCreatorResponse, error) {
	count, err := uc.repository.CountAccountByCpf(request.Cpf)
	if err != nil {
		return nil, err
	}

	if count > 0 {
		logrus.Warn(domain.ErrAccountConflict.Error())
		return nil, domain.ErrAccountConflict
	}

	hashedPassword, err := hashPassword(request.Secret)
	if err != nil {
		logrus.Error("error hashing password - ", err)
		return nil, err
	}

	return uc.repository.CreateAccount(request.Name, request.Cpf, hashedPassword)
}

func hashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	return string(hashedPasswordBytes), err
}
