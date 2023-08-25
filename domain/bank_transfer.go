/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package domain

import "time"

// BankTransfer content struct deifinition
type BankTransfer struct {
	Id                   string    `json:"id"`
	AccountOriginId      string    `json:"account_origin_id"`
	AccountDestinationId string    `json:"account_destination_id"`
	Amount               float64   `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

// BankTransfer usecase methods deifinition
type BankTransferUsecase interface {
}

// BankTransfer repository methods deifinition
type BankTransferRepository interface {
	TransferBetweenAccounts(amount float64, accountDestinationId string) error
}
