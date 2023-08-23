/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package usecase

import "github.com/msantosfelipe/go-bank-transfer/domain"

type accountUsecase struct {
}

func NewAccountUsecase() domain.AccountUsecase {
	return &accountUsecase{}
}

func (uc *accountUsecase) CreateAccount(accountRequest domain.AccountCreatorRequest) (*domain.AccountCreatorResponse, error) {
	return nil, nil
}
