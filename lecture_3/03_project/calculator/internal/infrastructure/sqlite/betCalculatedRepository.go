package sqlite

import (
	"context"
	"database/sql"
	domainmodels "github.com/koines/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "github.com/koines/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
	"github.com/pkg/errors"
)

// BetCalcRepository provides methods that operate on bets SQLite database.
type BetCalcRepository struct {
	dbExecutor DatabaseExecutor
	betMapper  BetMapper
}

// NewBetRepository creates and returns a new BetCalcRepository.
func NewBetRepository(dbExecutor DatabaseExecutor, betMapper BetMapper) *BetCalcRepository {
	return &BetCalcRepository{
		dbExecutor: dbExecutor,
		betMapper:  betMapper,
	}
}

// InsertBet inserts the provided bet into the database. An error is returned if the operation
// has failed.
func (r *BetCalcRepository) InsertBet(ctx context.Context, bet domainmodels.Bet) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryInsertBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "bet repository failed to insert a bet with id "+bet.Id)
	}
	return nil
}

func (r *BetCalcRepository) queryInsertBet(ctx context.Context, bet storagemodels.BetCalc) error {
	insertBetSQL := "INSERT INTO calc_bets(id, selection_id, selection_coefficient, payment) VALUES (?, ?, ?, ?)"
	statement, err := r.dbExecutor.PrepareContext(ctx, insertBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.Id, bet.SelectionId, bet.SelectionCoefficient, bet.Payment)
	return err
}

// UpdateBet updates the provided bet in the database. An error is returned if the operation
// has failed.
func (r *BetCalcRepository) UpdateBet(ctx context.Context, bet domainmodels.Bet) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryUpdateBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "bet repository failed to update a bet with id "+bet.Id)
	}
	return nil
}

func (r *BetCalcRepository) queryUpdateBet(ctx context.Context, bet storagemodels.BetCalc) error {
	updateBetSQL := "UPDATE calc_bets SET selection_id=?, selection_coefficient=?, payment=? WHERE id=?"

	statement, err := r.dbExecutor.PrepareContext(ctx, updateBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.SelectionId, bet.SelectionCoefficient, bet.Payment, bet.Id)
	return err
}

// GetBetByID fetches a bet from the database and returns it. The second returned value indicates
// whether the bet exists in DB. If the bet does not exist, an error will not be returned.
func (r *BetCalcRepository) GetBetByID(ctx context.Context, id string) (domainmodels.Bet, bool, error) {
	storageBet, err := r.queryGetBetByID(ctx, id)
	if err == sql.ErrNoRows {
		return domainmodels.Bet{}, false, nil
	}
	if err != nil {
		return domainmodels.Bet{}, false, errors.Wrap(err, "bet repository failed to get a bet with id "+id)
	}

	domainBet := r.betMapper.MapStorageBetToDomainBet(storageBet)
	return domainBet, true, nil
}

func (r *BetCalcRepository) queryGetBetByID(ctx context.Context, id string) (storagemodels.BetCalc, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM calc_bets WHERE id='"+id+"';")
	if err != nil {
		return storagemodels.BetCalc{}, err
	}
	defer row.Close()

	// This will move to the "next" result (which is the only result, because a single bet is fetched).
	row.Next()

	var selectionId string
	var selectionCoefficient string
	var payment int

	err = row.Scan(&id, &selectionId, &selectionCoefficient, &payment)
	if err != nil {
		return storagemodels.BetCalc{}, err
	}

	return storagemodels.BetCalc{
		Id:                   id,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
	}, nil
}

func (r *BetCalcRepository) queryGetBetsBySelectionID(ctx context.Context, id string) ([]storagemodels.BetCalc, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM calc_bets WHERE selection_id='"+id+"';")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var betsCalc []storagemodels.BetCalc

	// This will move to the "next" result (which is the only result, because a single bet is fetched).
	for row.Next() {
		var selectionId string
		var selectionCoefficient string
		var payment int
		var idBet string

		err = row.Scan(&idBet, &selectionId, &selectionCoefficient, &payment)
		if err != nil {
			return nil, err
		}

		var betCalc = storagemodels.BetCalc{
			Id:                   idBet,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
		}

		betsCalc = append(betsCalc, betCalc)
	}

	return betsCalc, nil
}

func (r *BetCalcRepository) GetBetsBySelectionID(ctx context.Context, id string) ([]domainmodels.Bet, bool, error) {
	storageBets, err := r.queryGetBetsBySelectionID(ctx, id)
	if err == sql.ErrNoRows {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, errors.Wrap(err, "bet repository failed to get a bet with id "+id)
	}

	var domainBets []domainmodels.Bet

	for _, storageBet := range storageBets {
		domainBet := r.betMapper.MapStorageBetToDomainBet(storageBet)
		domainBets = append(domainBets, domainBet)
	}

	return domainBets, true, nil
}
