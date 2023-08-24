/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package usecase

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
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
	if ok := isValidCpf(request.Cpf); !ok {
		logrus.Error(domain.ErrAccountInvalidCpf.Error())
		return nil, domain.ErrAccountInvalidCpf
	}

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

func (uc *accountUsecase) GetAccounts() (*domain.AccountList, error) {
	accounts, err := uc.repository.GetAccounts()
	if err != nil {
		return nil, err
	}

	// just for debug logging and use the secret and balance values
	for _, i := range accounts {
		level := logrus.GetLevel()
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug(fmt.Sprintf("cpf: %s, secret: %s, balance: %v", i.Cpf, i.Secret, i.Balance))
		logrus.SetLevel(level)
	}

	return &domain.AccountList{Accounts: accounts}, nil
}

func (uc *accountUsecase) GetAccountBalance(accountId string) (*domain.AccountBalance, error) {
	parsedUUID, err := uuid.Parse(accountId)
	if err != nil {
		logrus.Error("error parsing uuid - ", err)
		return nil, err
	}

	balance, err := uc.repository.GetAccountBalance(parsedUUID)
	if err != nil {
		return nil, err
	}

	return &domain.AccountBalance{Balance: balance}, nil
}

func isValidCpf(cpf string) bool {
	validPattern := regexp.MustCompile(`^\d{11}$`)
	return validPattern.MatchString(cpf)
}

func hashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	return string(hashedPasswordBytes), err
}
