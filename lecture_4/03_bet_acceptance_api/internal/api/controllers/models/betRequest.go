package models

type BetRequest struct {
	CustomerId           string  `json:"customer_id"`
	SelectionId          string  `json:"selection_id"`
	SelectionCoefficient float64 `json:"selection_coefficient"`
	Payment              float64 `json:"payment"`
}
