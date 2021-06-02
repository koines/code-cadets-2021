package controllers

import "github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api/controllers/models"

// BetService implements bet related functions.
type BetService interface {
	CreateBet(bet models.BetRequest) error
}
