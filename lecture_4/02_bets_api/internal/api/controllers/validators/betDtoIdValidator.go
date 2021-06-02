package validators

// BetDtoIdValidator validates event update requests.
type BetDtoIdValidator struct{}

// NewBetDtoIdValidator creates a new instance of BetDtoIdValidator.
func NewBetDtoIdValidator() *BetDtoIdValidator {
	return &BetDtoIdValidator{}
}

// IdIsValid checks if event update is valid.
// Id is not empty
// Outcome is `lost`or `won`
func (e *BetDtoIdValidator) IdIsValid(id string) bool {
	if id != "" {
		return true
	}

	return false
}
