/*
 * MIT License
 *
 * Copyright (c) 2023 Felipe Maia Santos
 *
 */

package usecase

import (
	"regexp"

	"github.com/google/uuid"
	"github.com/msantosfelipe/go-bank-transfer/domain"
	"github.com/sirupsen/logrus"
)

func isValidCpf(cpf string) bool {
	validPattern := regexp.MustCompile(`^\d{11}$`)
	return validPattern.MatchString(cpf)
}

func generateAccountUuid(id string) (uuid.UUID, error) {
	parsedUuid, err := uuid.Parse(id)
	if err != nil {
		logrus.Error("invalid id")
		return uuid.UUID{}, domain.ErrInvalidAccountId
	}

	return parsedUuid, nil
}
