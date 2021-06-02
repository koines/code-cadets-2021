package services

import "github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api/controllers/models"

// BetService implements event related functions.
type BetService struct {
	betReceivedPublisher BetReceivedPublisher
}

// NewBetService creates a new instance of BetService.
func NewBetService(betReceivedPublisher BetReceivedPublisher) *BetService {
	return &BetService{
		betReceivedPublisher: betReceivedPublisher,
	}
}

// CreateBet sends event update message to the queues.
func (e BetService) CreateBet(bet models.BetRequest) error {
	return e.betReceivedPublisher.Publish(bet)
}
