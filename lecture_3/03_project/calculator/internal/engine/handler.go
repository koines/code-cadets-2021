package engine

import (
	"context"
	rabbitmqmodels "github.com/koines/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Handler interface {
	HandleBets(ctx context.Context, betsReceived <-chan rabbitmqmodels.Bet)
	HandleEventUpdates(ctx context.Context, betsCalculated <-chan rabbitmqmodels.EventUpdate) <-chan rabbitmqmodels.BetCalculated
}
