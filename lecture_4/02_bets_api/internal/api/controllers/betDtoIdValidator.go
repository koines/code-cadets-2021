package controllers

// BetDtoIdValidator validates event update requests.
type BetDtoIdValidator interface {
	IdIsValid(id string) bool
}
