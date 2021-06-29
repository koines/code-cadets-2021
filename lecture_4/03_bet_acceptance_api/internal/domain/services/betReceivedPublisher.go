package services

import "github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api/controllers/models"

// BetReceivedPublisher handles bet received queue publishing.
type BetReceivedPublisher interface {
	Publish(bet models.BetRequest) error
}
