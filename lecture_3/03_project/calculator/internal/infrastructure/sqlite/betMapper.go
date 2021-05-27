package sqlite

import (
	domainmodels "github.com/koines/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "github.com/koines/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
)

type BetMapper interface {
	MapDomainBetToStorageBet(domainBet domainmodels.Bet) storagemodels.BetCalc
	MapStorageBetToDomainBet(storageBet storagemodels.BetCalc) domainmodels.Bet
}
