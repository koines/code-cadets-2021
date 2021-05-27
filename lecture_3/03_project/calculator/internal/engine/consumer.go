package engine

import (
	"context"
	rabbitmqmodels "github.com/koines/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Consumer interface {
	ConsumeBets(ctx context.Context) (<-chan rabbitmqmodels.Bet, error)
	ConsumeEventUpdates(ctx context.Context) (<-chan rabbitmqmodels.EventUpdate, error)
}
