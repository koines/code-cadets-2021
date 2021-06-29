package controllers

import "github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api/controllers/models"

// BetValidator validates event update requests.
type BetValidator interface {
	BetIsValid(bet models.BetRequest) bool
}
