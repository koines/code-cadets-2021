package controllers

// BetDtoStatusValidator validates event update requests.
type BetDtoStatusValidator interface {
	StatusIsValid(id string) bool
}
