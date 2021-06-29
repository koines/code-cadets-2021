package main

import (
	"fmt"
	"github.com/koines/code-cadets-2021/homework_2/offerfeed/cmd/bootstrap"
	"github.com/koines/code-cadets-2021/homework_2/offerfeed/internal/domain/services"
	"github.com/koines/code-cadets-2021/homework_2/offerfeed/internal/tasks"
)

func main() {
	signalHandler := bootstrap.SignalHandler()

	feed := bootstrap.AxilisOfferFeed()
	feed2 := bootstrap.AxilisOfferFeed2()
	queue := bootstrap.OrderedQueue()
	processingService := bootstrap.FeedProcessingService([]services.Feed{feed, feed2}, queue)

	// blocking call, start "the application"
	tasks.RunTasks(signalHandler, feed, feed2, queue, processingService)

	fmt.Println("program finished gracefully")
}
