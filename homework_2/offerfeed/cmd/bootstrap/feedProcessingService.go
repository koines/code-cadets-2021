package bootstrap

import "github.com/koines/code-cadets-2021/homework_2/offerfeed/internal/domain/services"

func FeedProcessingService(feeds []services.Feed, queue services.Queue) *services.FeedProcessorService {
	return services.NewFeedProcessorService(feeds, queue)
}
