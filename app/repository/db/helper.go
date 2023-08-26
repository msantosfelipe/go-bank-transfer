/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package db

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/msantosfelipe/go-bank-transfer/app/repository/db/queries"
	"github.com/sirupsen/logrus"
)

func microMoneytoMoney(microMoney int64) float64 {
	return float64(microMoney) / 1000000.0
}

func moneyToMicroMoney(floatValue float64) int64 {
	return int64(floatValue * 1000000)
}

func getBalances(accounts []queries.GetAccountsBalancesRow, originId, destinationId uuid.UUID) (int64, int64, error) {
	var originBalance, destinationBalance int64
	var err error
	for _, i := range accounts {
		switch i.ID {
		case originId:
			originBalance = i.Balance
		case destinationId:
			destinationBalance = i.Balance
		default:
			logrus.Error("account ids do not match")
			err = errors.New("account ids do not match")
			return 0, 0, err
		}
	}

	return originBalance, destinationBalance, err
}

func transferBalances(originBalance, destinationBalance, transferAmount int64) (int64, int64) {
	originBalance -= transferAmount
	destinationBalance += transferAmount
	return originBalance, destinationBalance
}

func formatDate(date time.Time) string {
	return date.Format("02/01/2006 15:04:05")
}
