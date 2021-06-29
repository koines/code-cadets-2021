package rabbitmq

import (
	"encoding/json"
	"github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/infrastructure/rabbitmq/models"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/pkg/errors"

	"github.com/streadway/amqp"
	"log"

	domainmodels "github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/api/controllers/models"
)

const contentTypeTextPlain = "text/plain"

// BetReceivedPublisher handles event update queue publishing.
type BetReceivedPublisher struct {
	exchange  string
	queueName string
	mandatory bool
	immediate bool
	publisher QueuePublisher
}

// NewBetReceivedPublisher create a new instance of betReceivedPublisher.
func NewBetReceivedPublisher(
	exchange string,
	queueName string,
	mandatory bool,
	immediate bool,
	publisher QueuePublisher,
) *BetReceivedPublisher {
	return &BetReceivedPublisher{
		exchange:  exchange,
		queueName: queueName,
		mandatory: mandatory,
		immediate: immediate,
		publisher: publisher,
	}
}

func getRandomUUID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.WithMessage(err, "id generator failed")
	}

	return id.String(), nil
}

// Publish publishes an event update message to the queue.
func (p *BetReceivedPublisher) Publish(betDto domainmodels.BetRequest) error {
	betId, err := getRandomUUID()
	if err != nil {
		return err
	}
	bet := &models.BetDto{
		Id:                   betId,
		CustomerId:           betDto.CustomerId,
		SelectionId:          betDto.SelectionId,
		SelectionCoefficient: betDto.SelectionCoefficient,
		Payment:              betDto.Payment,
	}

	betJson, err := json.Marshal(bet)
	if err != nil {
		return err
	}

	err = p.publisher.Publish(
		p.exchange,
		p.queueName,
		p.mandatory,
		p.immediate,
		amqp.Publishing{
			ContentType: contentTypeTextPlain,
			Body:        betJson,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Sent %s", betJson)
	return nil
}
