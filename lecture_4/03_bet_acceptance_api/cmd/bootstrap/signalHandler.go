package bootstrap

import "github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/tasks"

// SignalHandler bootstraps the signal handler.
func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
