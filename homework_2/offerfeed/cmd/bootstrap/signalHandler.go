package bootstrap

import "github.com/koines/code-cadets-2021/homework_2/offerfeed/internal/tasks"

func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
