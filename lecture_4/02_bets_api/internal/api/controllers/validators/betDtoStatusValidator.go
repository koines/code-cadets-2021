package validators

const lostStatus = "lost"
const wonStatus = "won"
const activeStatus = "active"

type BetDtoStatusValidator struct{}

// NewBetDtoStatusValidator creates a new instance of BetDtoStatusValidator.
func NewBetDtoStatusValidator() *BetDtoStatusValidator {
	return &BetDtoStatusValidator{}
}

// StatusIsValid checks if status is valid.
// Status is `lost`, `won` or `active`
func (e *BetDtoStatusValidator) StatusIsValid(status string) bool {
	if status == lostStatus || status == wonStatus || status == activeStatus {
		return true
	}

	return false
}
