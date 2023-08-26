/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package domain

import (
	"github.com/google/uuid"
)

// Transfer content struct deifinition
type Transfer struct {
	Id                   string  `json:"id"`
	AccountOriginId      string  `json:"account_origin_id"`
	AccountDestinationId string  `json:"account_destination_id"`
	Amount               float64 `json:"amount"`
	CreatedAt            string  `json:"created_at"`
}

// TransferRequest content struct deifinition
type TransferRequest struct {
	AccountDestinationId string  `json:"account_destination_id"`
	Amount               float64 `json:"amount"`
}

type TransferList struct {
	Transfers []Transfer `json:"transfers"`
}

// TransferCreatorResponse content struct deifinition
// 'old' and 'new' values placed for better visualization
type TransferCreatorResponse struct {
	Id                           string  `json:"id"`
	OldAccountOriginBalance      float64 `json:"old_account_origin_balance"`
	NewAccountOriginBalance      float64 `json:"new_account_origin_balance"`
	OldAccountDestinationBalance float64 `json:"old_account_destination_balance"`
	NewAccountDestinationBalance float64 `json:"new_account_destination_balance"`
}

// Transfer usecase methods deifinition
type TransferUsecase interface {
	TransferBetweenAccounts(
		originAccountId string, request TransferRequest,
	) (*TransferCreatorResponse, error)
	GetAccountOriginTransfers(accountOriginId string) (*TransferList, error)
}

// Transfer repository methods deifinition
type TransferRepository interface {
	TransferBetweenAccounts(
		amount float64, accountOriginId, accountDestinationId uuid.UUID,
	) (*TransferCreatorResponse, error)
	GetAccountOriginTransfers(accountOriginId uuid.UUID) ([]Transfer, error)
}
