/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package usecase

import (
	"github.com/msantosfelipe/go-bank-transfer/domain"
)

type transferUsecase struct {
	repository domain.TransferRepository
}

func NewTransferUsecase(repository domain.TransferRepository) domain.TransferUsecase {
	return &transferUsecase{
		repository: repository,
	}
}

func (uc *transferUsecase) TransferBetweenAccounts(
	originAccountId string, request domain.TransferRequest,
) (*domain.TransferCreatorResponse, error) {
	originUuid, err := generateAccountUuid(originAccountId)
	if err != nil {
		return nil, err
	}

	destinationUuid, err := generateAccountUuid(request.AccountDestinationId)
	if err != nil {
		return nil, err
	}

	return uc.repository.TransferBetweenAccounts(request.Amount, originUuid, destinationUuid)
}

func (uc *transferUsecase) GetAccountOriginTransfers(accountOriginId string) (*domain.TransferList, error) {
	parsedUUID, err := generateAccountUuid(accountOriginId)
	if err != nil {
		return nil, err
	}

	transfers, err := uc.repository.GetAccountOriginTransfers(parsedUUID)
	if err != nil {
		return nil, err
	}

	return &domain.TransferList{
		Transfers: transfers,
	}, nil
}
