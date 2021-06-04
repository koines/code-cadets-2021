package models

// BetCalc represents a DTO for calculated bets.
type BetCalc struct {
	Id                   string
	SelectionId          string
	SelectionCoefficient int
	Payment              int
}
