/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package usecase

import (
	"github.com/google/uuid"
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/sirupsen/logrus"
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
	originUuid, destinationUuid, err := generateUuids(originAccountId, request.AccountDestinationId)
	if err != nil {
		return nil, err
	}

	return uc.repository.TransferBetweenAccounts(request.Amount, originUuid, destinationUuid)
}

func generateUuids(originId, destinationId string) (uuid.UUID, uuid.UUID, error) {
	originUuid, err := uuid.Parse(originId)
	if err != nil {
		logrus.Error("invalid account origin id")
		return uuid.UUID{}, uuid.UUID{}, domain.ErrInvalidAccountId
	}

	destinationUuid, err := uuid.Parse(destinationId)
	if err != nil {
		logrus.Error("invalid account destination id")
		return uuid.UUID{}, uuid.UUID{}, domain.ErrInvalidAccountId
	}

	return originUuid, destinationUuid, nil
}
