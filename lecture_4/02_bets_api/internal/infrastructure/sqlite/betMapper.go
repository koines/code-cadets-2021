package sqlite

import (
	domainmodels "github.com/koines/code-cadets-2021/lecture_4/02_bets_api/internal/api/controllers/models"
	storagemodels "github.com/koines/code-cadets-2021/lecture_4/02_bets_api/internal/infrastructure/sqlite/models"
)

type BetMapper interface {
	MapStorageBetToDomainBet(storageBet storagemodels.Bet) domainmodels.BetDto
}
