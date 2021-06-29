package validators

type BetDtoIdValidator struct{}

// NewBetDtoIdValidator creates a new instance of BetDtoIdValidator.
func NewBetDtoIdValidator() *BetDtoIdValidator {
	return &BetDtoIdValidator{}
}

// IdIsValid checks if id is valid.
// Id is not empty
func (e *BetDtoIdValidator) IdIsValid(id string) bool {
	if id != "" {
		return true
	}

	return false
}
