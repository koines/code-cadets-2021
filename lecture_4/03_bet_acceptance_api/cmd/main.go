package main

import (
	"github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/cmd/bootstrap"
	"github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/cmd/config"
	"github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/tasks"
	"log"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	rabbitMqChannel := bootstrap.RabbitMq()
	signalHandler := bootstrap.SignalHandler()
	api := bootstrap.Api(rabbitMqChannel)

	log.Println("Bootstrap finished. Bet acceptance API is starting")

	tasks.RunTasks(signalHandler, api)

	log.Println("Bet acceptance API finished gracefully")
}
