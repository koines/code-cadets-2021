package bootstrap

import (
	"github.com/koines/code-cadets-2021/lecture_4/02_bets_api/cmd/config"
	"github.com/koines/code-cadets-2021/lecture_4/02_bets_api/internal/api"
	"github.com/koines/code-cadets-2021/lecture_4/02_bets_api/internal/api/controllers"
	"github.com/koines/code-cadets-2021/lecture_4/02_bets_api/internal/api/controllers/validators"
	"github.com/koines/code-cadets-2021/lecture_4/02_bets_api/internal/domain/mappers"
	"github.com/koines/code-cadets-2021/lecture_4/02_bets_api/internal/domain/services"
	"github.com/koines/code-cadets-2021/lecture_4/02_bets_api/internal/infrastructure/sqlite"
)

func newBetDtoIdValidator() *validators.BetDtoIdValidator {
	return validators.NewBetDtoIdValidator()
}

func newBetDtoStatusValidator() *validators.BetDtoStatusValidator {
	return validators.NewBetDtoStatusValidator()
}

func newBetRepository(dbExecutor sqlite.DatabaseExecutor, dbMapper sqlite.BetMapper) *sqlite.BetRepository {
	return sqlite.NewBetRepository(dbExecutor, dbMapper)
}

func newBetDtoService(betRepository services.BetRepository) *services.BetDtoService {
	return services.NewBetDtoService(betRepository)
}

func newController(idValidator controllers.BetDtoIdValidator, statusValidator controllers.BetDtoStatusValidator, betService controllers.BetDtoService) *controllers.Controller {
	return controllers.NewController(idValidator, statusValidator, betService)
}

// Api bootstraps the http server.
func Api(dbExecutor sqlite.DatabaseExecutor) *api.WebServer {
	idValidator := newBetDtoIdValidator()
	statusValidator := newBetDtoStatusValidator()
	dbMapper := mappers.NewBetMapper()
	repository := newBetRepository(dbExecutor, dbMapper)
	betService := newBetDtoService(repository)
	controller := newController(idValidator, statusValidator, betService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
