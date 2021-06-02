package bootstrap

import (
	"github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/cmd/config"
	"github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api"
	"github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api/controllers"
	"github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api/controllers/validators"
	"github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/domain/services"
	"github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/infrastructure/rabbitmq"
	"github.com/streadway/amqp"
)

const coefficientBound = 10.0
const paymentLowerBound = 2.0
const paymentUpperBound = 100.0

func newBetValidator() *validators.BetValidator {
	return validators.NewBetValidator(coefficientBound, paymentLowerBound, paymentUpperBound)
}

func newBetPublisher(publisher rabbitmq.QueuePublisher) *rabbitmq.BetReceivedPublisher {
	return rabbitmq.NewBetReceivedPublisher(
		config.Cfg.Rabbit.PublisherExchange,
		config.Cfg.Rabbit.PublisherBetQueueQueue,
		config.Cfg.Rabbit.PublisherMandatory,
		config.Cfg.Rabbit.PublisherImmediate,
		publisher,
	)
}

func newBetService(publisher services.BetReceivedPublisher) *services.BetService {
	return services.NewBetService(publisher)
}

func newController(betValidator controllers.BetValidator, betService controllers.BetService) *controllers.Controller {
	return controllers.NewController(betValidator, betService)
}

// Api bootstraps the http server.
func Api(rabbitMqChannel *amqp.Channel) *api.WebServer {
	betValidator := newBetValidator()
	betPublisher := newBetPublisher(rabbitMqChannel)
	betService := newBetService(betPublisher)
	controller := newController(betValidator, betService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
