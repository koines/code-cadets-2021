package main

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/sethgrid/pester"

	"log"
)

const betsAPIPath = "http://127.0.0.1:8081/bets?status=active"
const eventAPIPath = "http://127.0.0.1:8080/event/update"
const wonOutcome = "won"
const lostOutcome = "lost"

type eventUpdateDto struct {
	Id      string `json:"id"`
	Outcome string `json:"outcome"`
}

type betDto struct {
	Id                   string  `json:"id"`
	CustomerId           string  `json:"customerId"`
	SelectionId          string  `json:"selectionId"`
	SelectionCoefficient float64 `json:"selectionCoefficient"`
	Payment              float64 `json:"payment"`
}

func getActiveBets(httpClient pester.Client) ([]betDto, error) {
	httpResponse, err := httpClient.Get(betsAPIPath)
	if err != nil {
		return nil, err
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var decodedContent []betDto
	error := json.Unmarshal(bodyContent, &decodedContent)
	if error != nil {
		return nil, error
	}

	return decodedContent, nil
}

func getDistinctMatches(bets []betDto) map[string]bool {
	matches := make(map[string]bool)

	for _, bet := range bets {
		matches[bet.SelectionId] = true
	}

	return matches
}

func publish(event eventUpdateDto) error {
	eventUpdateJson, err := json.Marshal(event)
	if err != nil {
		return errors.WithMessage(err, "failed to marshal event update")
	}

	_, error := http.Post(eventAPIPath, "application/json",
		bytes.NewBuffer(eventUpdateJson))
	if error != nil {
		return errors.WithMessage(err, "post event update")
	}

	log.Printf("Sent %s", eventUpdateJson)

	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	httpClient := pester.New()

	activeBets, err := getActiveBets(*httpClient)
	if err != nil {
		log.Fatalf("retrive active bets: %s", err)
	}

	distinctMatches := getDistinctMatches(activeBets)

	for key, _ := range distinctMatches {
		var outcome string
		if rand.Float64() > 0.5 {
			outcome = lostOutcome
		} else {
			outcome = wonOutcome
		}

		eventUpdate := &eventUpdateDto{
			Id:      key,
			Outcome: outcome,
		}

		error := publish(*eventUpdate)
		if error != nil {
			log.Fatalf("publish: %s", error)
		}
	}
}
