package consumer

import (
	"context"
	rabbitmqmodels "github.com/koines/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type BetConsumer interface {
	Consume(ctx context.Context) (<-chan rabbitmqmodels.Bet, error)
}
