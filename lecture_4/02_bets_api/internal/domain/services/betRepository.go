package services

import (
	"context"
	domainmodels "github.com/koines/code-cadets-2021/lecture_4/02_bets_api/internal/api/controllers/models"
)

type BetRepository interface {
	GetBetByID(ctx context.Context, id string) (domainmodels.BetDto, bool, error)
	GetBetsByStatus(ctx context.Context, status string) ([]domainmodels.BetDto, bool, error)
	GetBetsByCustomerId(ctx context.Context, userId string) ([]domainmodels.BetDto, bool, error)
}
