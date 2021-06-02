package validators

const lostStatus = "lost"
const wonStatus = "won"
const activeStatus = "active"

// BetDtoIdValidator validates event update requests.
type BetDtoStatusValidator struct{}

// NewBetDtoIdValidator creates a new instance of BetDtoIdValidator.
func NewBetDtoStatusValidator() *BetDtoStatusValidator {
	return &BetDtoStatusValidator{}
}

// IdIsValid checks if event update is valid.
// Id is not empty
// Outcome is `lost`or `won`
func (e *BetDtoStatusValidator) StatusIsValid(status string) bool {
	if status == lostStatus || status == wonStatus || status == activeStatus {
		return true
	}

	return false
}
