package mappers

import (
	domainmodels "github.com/koines/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "github.com/koines/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
	"math"
	"strconv"
)

// BetMapper maps storage bets to domain bets and vice versa.
type BetMapper struct {
}

// NewBetStorageMapper creates and returns a new BetMapper.
func NewBetMapper() *BetMapper {
	return &BetMapper{}
}

// MapDomainBetToStorageBet maps the given domain bet into storage bet. Floating point values will
// be converted to corresponding integer values of the storage bet by multiplying them with 100.
func (m *BetMapper) MapDomainBetToStorageBet(domainBet domainmodels.Bet) storagemodels.BetCalc {
	return storagemodels.BetCalc{
		Id:                   domainBet.Id,
		SelectionId:          domainBet.SelectionId,
		SelectionCoefficient: strconv.FormatFloat(math.Round(domainBet.SelectionCoefficient*100), 'E', -1, 64),
		Payment:              int(math.Round(domainBet.Payment * 100)),
	}
}

// MapStorageBetToDomainBet maps the given storage bet into domain bet. Floating point values will
// be converted from corresponding integer values of the storage bet by dividing them with 100.
func (m *BetMapper) MapStorageBetToDomainBet(storageBet storagemodels.BetCalc) domainmodels.Bet {
	coefficient, _ := strconv.ParseFloat(storageBet.SelectionCoefficient, 64)
	coefficient /= 100
	return domainmodels.Bet{
		Id:                   storageBet.Id,
		SelectionId:          storageBet.SelectionId,
		SelectionCoefficient: coefficient,
		Payment:              float64(storageBet.Payment) / 100,
	}
}
