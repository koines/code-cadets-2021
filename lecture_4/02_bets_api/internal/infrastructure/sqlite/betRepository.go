package sqlite

import (
	"context"
	"database/sql"
	domainmodels "github.com/koines/code-cadets-2021/lecture_4/02_bets_api/internal/api/controllers/models"
	storagemodels "github.com/koines/code-cadets-2021/lecture_4/02_bets_api/internal/infrastructure/sqlite/models"
	"github.com/pkg/errors"
)

// BetRepository provides methods that operate on bets SQLite database.
type BetRepository struct {
	dbExecutor DatabaseExecutor
	betMapper  BetMapper
}

// NewBetRepository creates and returns a new BetRepository.
func NewBetRepository(dbExecutor DatabaseExecutor, betMapper BetMapper) *BetRepository {
	return &BetRepository{
		dbExecutor: dbExecutor,
		betMapper:  betMapper,
	}
}

// GetBetByID fetches a bet from the database and returns it. The second returned value indicates
// whether the bet exists in DB. If the bet does not exist, an error will not be returned.
func (r *BetRepository) GetBetByID(ctx context.Context, id string) (domainmodels.BetDto, bool, error) {
	storageBet, err := r.queryGetBetByID(ctx, id)
	if err == sql.ErrNoRows {
		return domainmodels.BetDto{}, false, nil
	}
	if err != nil {
		return domainmodels.BetDto{}, false, errors.Wrap(err, "bet repository failed to get a bet with id "+id)
	}

	domainBet := r.betMapper.MapStorageBetToDomainBet(storageBet)
	return domainBet, true, nil
}

func (r *BetRepository) queryGetBetByID(ctx context.Context, id string) (storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE id='"+id+"';")
	if err != nil {
		return storagemodels.Bet{}, err
	}
	defer row.Close()

	// This will move to the "next" result (which is the only result, because a single bet is fetched).
	row.Next()

	bet := storagemodels.Bet{}
	var payoutSql sql.NullInt64

	err = row.Scan(&bet.Id, &bet.CustomerId, &bet.Status, &bet.SelectionId, &bet.SelectionCoefficient, &bet.Payment, &payoutSql)
	if err != nil {
		return storagemodels.Bet{}, err
	}

	var payout int
	if payoutSql.Valid {
		payout = int(payoutSql.Int64)
	}


	return storagemodels.Bet{
		Id:                   id,
		CustomerId:           bet.CustomerId,
		Status:               bet.Status,
		SelectionId:          bet.SelectionId,
		SelectionCoefficient: bet.SelectionCoefficient,
		Payment:              bet.Payment,
		Payout:               payout,
	}, nil
}

// GetBetsByStatus fetches all bets from the database that have provided status
// and returns them. The second returned value indicates whether the bet exists in DB.
// If the bet does not exist, an error will not be returned.
func (r *BetRepository) GetBetsByStatus(ctx context.Context, status string) ([]domainmodels.BetDto, bool, error) {
	storageBets, err := r.queryGetBetByStatus(ctx, status)
	if err == sql.ErrNoRows {
		return []domainmodels.BetDto{}, false, nil
	}
	if err != nil {
		return []domainmodels.BetDto{}, false, errors.Wrap(err, "bets repository failed to get a bet with status"+status)
	}

	var domainBets []domainmodels.BetDto
	for _, storageBet := range storageBets {
		domainBets = append(domainBets, r.betMapper.MapStorageBetToDomainBet(storageBet))
	}

	return domainBets, true, nil
}

func (r *BetRepository) queryGetBetByStatus(ctx context.Context, status string) ([]storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE status='"+status+"';")
	if err != nil {
		return []storagemodels.Bet{}, errors.Wrap(err, "query")
	}
	defer row.Close()

	var bets []storagemodels.Bet
	var payoutSql sql.NullInt64

	bet := storagemodels.Bet{}

	// A loop over all returned rows.
	for row.Next() {
		err = row.Scan(&bet.Id, &bet.CustomerId, &bet.Status, &bet.SelectionId, &bet.SelectionCoefficient, &bet.Payment, &payoutSql)
		if err != nil {
			return []storagemodels.Bet{}, err
		}

		var payout int
		if payoutSql.Valid {
			payout = int(payoutSql.Int64)
		}


		bets = append(bets, storagemodels.Bet{
			Id:                   bet.Id,
			CustomerId:           bet.CustomerId,
			Status:               bet.Status,
			SelectionId:          bet.SelectionId,
			SelectionCoefficient: bet.SelectionCoefficient,
			Payment:              bet.Payment,
			Payout:               payout,
		})
	}

	return bets, nil
}

// GetBetsByCustomerId fetches all bets from the database that have provided status
// and returns them. The second returned value indicates whether the bet exists in DB.
// If the bet does not exist, an error will not be returned.
func (r *BetRepository) GetBetsByCustomerId(ctx context.Context, customerId string) ([]domainmodels.BetDto, bool, error) {
	storageBets, err := r.queryGetBetByCustomerId(ctx, customerId)
	if err == sql.ErrNoRows {
		return []domainmodels.BetDto{}, false, nil
	}
	if err != nil {
		return []domainmodels.BetDto{}, false, errors.Wrap(err, "bets repository failed to get a bet with selection id "+customerId)
	}

	var domainBets []domainmodels.BetDto
	for _, storageBet := range storageBets {
		domainBets = append(domainBets, r.betMapper.MapStorageBetToDomainBet(storageBet))
	}

	return domainBets, true, nil
}

func (r *BetRepository) queryGetBetByCustomerId(ctx context.Context, customerId string) ([]storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE customer_id='"+customerId+"';")
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer row.Close()

	var bets []storagemodels.Bet

	// A loop over all returned rows.
	for row.Next() {
		var status string
		var id string
		var selectionId string
		var selectionCoefficient int
		var payment int
		var payoutSql sql.NullInt64

		err = row.Scan(&id, &customerId, &status, &selectionId, &selectionCoefficient, &payment, &payoutSql)
		if err != nil {
			return []storagemodels.Bet{}, err
		}

		var payout int

		if payoutSql.Valid {
			payout = int(payoutSql.Int64)
		}

		bets = append(bets, storagemodels.Bet{
			Id:                   id,
			CustomerId:           customerId,
			Status:               status,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
			Payout:               payout,
		})
	}

	return bets, nil
}
