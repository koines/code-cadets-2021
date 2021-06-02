package validators

import (
	"github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api/controllers/models"
)

// BetValidator validates bet requests.
type BetValidator struct {
	coefficientUpperBound float64
	paymentLowerBound     float64
	paymentUpperBound     float64
}

// NewBetValidator creates a new instance of BetValidator.
func NewBetValidator(coefficentBound float64, paymentLowerBound float64, paymentUpperBound float64) *BetValidator {
	return &BetValidator{
		coefficientUpperBound: coefficentBound,
		paymentLowerBound:     paymentLowerBound,
		paymentUpperBound:     paymentUpperBound,
	}
}

// BetIsValid checks if bet is valid.
func (e *BetValidator) BetIsValid(betRequest models.BetRequest) bool {
	if e.IdIsValid(betRequest.CustomerId) && e.IdIsValid(betRequest.SelectionId) && e.CoefficientIsValid(betRequest.SelectionCoefficient) && e.PaymentIsValid(betRequest.Payment) {
		return true
	}

	return false
}

func (e *BetValidator) IdIsValid(id string) bool {
	if id == "" {
		return false
	}

	return true
}

func (e *BetValidator) CoefficientIsValid(coefficient float64) bool {
	if coefficient < 0.0 || coefficient > e.coefficientUpperBound {
		return false
	}

	return true
}

func (e *BetValidator) PaymentIsValid(payment float64) bool {
	if payment < e.paymentLowerBound || payment > e.paymentUpperBound {
		return false
	}

	return true
}
