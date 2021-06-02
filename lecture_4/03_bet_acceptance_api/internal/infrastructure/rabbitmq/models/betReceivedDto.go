package models

type BetDto struct {
	Id                   string  `json:"id"`
	CustomerId           string  `json:"customer_id"`
	SelectionId          string  `json:"selection_id"`
	SelectionCoefficient float64 `json:"selection_coefficient"`
	Payment              float64 `json:"payment"`
}
