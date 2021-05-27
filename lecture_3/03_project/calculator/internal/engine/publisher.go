package engine

import (
	"context"
	rabbitmqmodels "github.com/koines/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Publisher interface {
	PublishBetCalculated(ctx context.Context, bets <-chan rabbitmqmodels.BetCalculated)
}
