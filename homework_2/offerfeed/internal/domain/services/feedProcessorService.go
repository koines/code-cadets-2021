package services

import (
	"context"
	"github.com/koines/code-cadets-2021/homework_2/offerfeed/internal/domain/models"
	"log"
	"reflect"
)

type FeedProcessorService struct {
	feeds []Feed
	queue Queue
}

func NewFeedProcessorService(
	feeds []Feed,
	queue Queue,
) *FeedProcessorService {
	return &FeedProcessorService{
		feeds: feeds,
		queue: queue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	defer log.Printf("shutting down %s", f)

	source := f.queue.GetSource()
	defer close(source)

	var updateChannels []chan models.Odd
	for _, feedIter := range f.feeds {
		updateChannels = append(updateChannels, feedIter.GetUpdates())
	}

	cases := make([]reflect.SelectCase, len(updateChannels))
	for i, ch := range updateChannels {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		}
	}

	for len(cases) > 0 {
		chosen, value, ok := reflect.Select(cases)
		if !ok {
			cases = append(cases[:chosen], cases[chosen+1:]...)
			continue
		}

		input := value.Interface().(models.Odd)
		input.Coefficient *= 2
		source <- input
	}

	return nil
}

func (f *FeedProcessorService) String() string {
	return "feed processor service"
}

type Feed interface {
	GetUpdates() chan models.Odd
}

type Queue interface {
	GetSource() chan models.Odd
}
