package services

import (
	"context"
	domainmodels "github.com/koines/code-cadets-2021/lecture_4/02_bets_api/internal/api/controllers/models"
	"github.com/pkg/errors"
)

// BetDtoService implements event related functions.
type BetDtoService struct {
	betRepository BetRepository
}

// NewBetDtoService creates a new instance of BetDtoService.
func NewBetDtoService(betRepository BetRepository) *BetDtoService {
	return &BetDtoService{
		betRepository: betRepository,
	}
}

func (e BetDtoService) GetByID(ctx context.Context, id string) (domainmodels.BetDto, error) {
	bet, found, err := e.betRepository.GetBetByID(ctx, id)
	if !found {
		return domainmodels.BetDto{}, errors.WithMessage(err, "Bet with specified id does not exist in table")
	}

	return bet, err
}

func (e BetDtoService) GetAll(ctx context.Context, id string) ([]domainmodels.BetDto, error) {
	bets, found, err := e.betRepository.GetBetsByCustomerId(ctx, id)
	if !found {
		return bets, errors.WithMessage(err, "No such bets in database")
	}

	return bets, err
}

func (e BetDtoService) GetByStatus(ctx context.Context, status string) ([]domainmodels.BetDto, error) {
	bets, found, err := e.betRepository.GetBetsByStatus(ctx, status)
	if !found {
		return bets, errors.WithMessage(err, "No such bets in database")
	}

	return bets, err
}
