package controllers

import (
	"context"
	domainmodels "github.com/koines/code-cadets-2021/lecture_4/02_bets_api/internal/api/controllers/models"
)

// BetDtoService implements event related functions.
type BetDtoService interface {
	GetByID(ctx context.Context, betDtoId string) (domainmodels.BetDto, error)
	GetAll(ctx context.Context, betDtoUserId string) ([]domainmodels.BetDto, error)
	GetByStatus(ctx context.Context, betDtoStatus string) ([]domainmodels.BetDto, error)
}
