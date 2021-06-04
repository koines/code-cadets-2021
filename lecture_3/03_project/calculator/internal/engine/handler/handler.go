package handler

import (
	"context"
	domainmodels "github.com/koines/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	rabbitmqmodels "github.com/koines/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
	"github.com/mattn/go-sqlite3"
	"log"
)

const outcomeWon = "won"
const outcomeLost = "lost"

// Handler handles bets received and bets calculated.
type Handler struct {
	betCalcRepository BetCalcRepository
}

// New creates and returns a new Handler.
func New(betCalcRepository BetCalcRepository) *Handler {
	return &Handler{
		betCalcRepository: betCalcRepository,
	}
}

// HandleBets handles bets received.
func (h *Handler) HandleBets(
	ctx context.Context,
	bets <-chan rabbitmqmodels.Bet,
) {

	go func() {
		for bet := range bets {
			log.Println("Processing bet received, betId:", bet.Id)

			// Calculate the domain bet based on the incoming bet received.
			domainBet := domainmodels.Bet{
				Id:                   bet.Id,
				SelectionId:          bet.SelectionId,
				SelectionCoefficient: bet.SelectionCoefficient,
				Payment:              bet.Payment,
			}

			// Insert the domain bet into the repository.
			err := h.betCalcRepository.InsertBet(ctx, domainBet)
			if err == sqlite3.ErrConstraint {
				log.Println("bet already exists")
				continue
			} else if err != nil {
				log.Println("Failed to insert bet, error: ", err)
				continue
			}
		}
	}()
}

// HandleEventUpdates handles bets calculated.
func (h *Handler) HandleEventUpdates(
	ctx context.Context,
	eventUpdates <-chan rabbitmqmodels.EventUpdate,
) <-chan rabbitmqmodels.BetCalculated {
	resultingBets := make(chan rabbitmqmodels.BetCalculated)

	go func() {
		defer close(resultingBets)

		for eventUpdate := range eventUpdates {
			log.Println("Processing event update, betId:", eventUpdate.Id)

			// Fetch the domain bet.
			domainBets, exists, err := h.betCalcRepository.GetBetsBySelectionID(ctx, eventUpdate.Id)
			if err != nil {
				log.Println("Failed to fetch a bet which should be updated, error: ", err)
				continue
			}
			if !exists {
				log.Println("A bet which should be updated does not exist, betId: ", eventUpdate.Id)
				continue
			}

			// Calculate the resulting bet, which should be published.
			for _, domainBet := range domainBets {
				status := outcomeLost
				payout := 0.0

				if eventUpdate.Outcome == outcomeWon {
					payout = domainBet.Payment * domainBet.SelectionCoefficient
					status = eventUpdate.Outcome
				}

				resultingBet := rabbitmqmodels.BetCalculated{
					Id:     domainBet.Id,
					Status: status,
					Payout: payout,
				}

				select {
				case resultingBets <- resultingBet:
				case <-ctx.Done():
					return
				}

			}
		}
	}()

	return resultingBets
}
