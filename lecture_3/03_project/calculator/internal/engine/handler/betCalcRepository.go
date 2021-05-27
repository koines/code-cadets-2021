package handler

import (
	"context"
	domainmodels "github.com/koines/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
)

type BetCalcRepository interface {
	InsertBet(ctx context.Context, bet domainmodels.Bet) error
	UpdateBet(ctx context.Context, bet domainmodels.Bet) error
	GetBetsBySelectionID(ctx context.Context, id string) ([]domainmodels.Bet, bool, error)
}
