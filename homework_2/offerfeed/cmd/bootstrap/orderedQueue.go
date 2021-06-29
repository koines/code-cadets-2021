package bootstrap

import "github.com/koines/code-cadets-2021/homework_2/offerfeed/internal/infrastructure/queue"

func OrderedQueue() *queue.OrderedQueue {
	return queue.NewOrderedQueue()
}
