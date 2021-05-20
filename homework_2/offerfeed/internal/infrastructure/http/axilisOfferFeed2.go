package http

import (
	"context"
	"github.com/koines/code-cadets-2021/homework_2/offerfeed/internal/domain/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const axilisFeedURL2 = "http://18.193.121.232/axilis-feed-2"

type AxilisOfferFeed2 struct {
	updates    chan models.Odd
	httpClient *http.Client
}

func NewAxilisOfferFeed2(
	httpClient *http.Client,
) *AxilisOfferFeed2 {
	return &AxilisOfferFeed2{
		updates:    make(chan models.Odd),
		httpClient: httpClient,
	}
}

func (a *AxilisOfferFeed2) Start(ctx context.Context) error {
	defer close(a.updates)
	defer log.Printf("shutting down %s", a)

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-time.After(time.Second):
			response, err := a.httpClient.Get(axilisFeedURL2)
			if err != nil {
				log.Println("axilis offer feed2, http get", err)
				continue
			}
			a.processResponse(ctx, response)
		}
	}
}

func (a *AxilisOfferFeed2) GetUpdates() chan models.Odd {
	return a.updates
}

func (a *AxilisOfferFeed2) String() string {
	return "axilis offer feed2"
}

func (a *AxilisOfferFeed2) processResponse(ctx context.Context, response *http.Response) {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("axilis offer feed2, ioutil readAll", err)
		return
	}

	bodyDecoded := strings.Split(string(body), "\n")

	for _, row := range bodyDecoded {
		items := strings.Split(row, ",")

		coeff, err := strconv.ParseFloat(items[3], 64)
		if err != nil {
			log.Println("axilis offer feed2, parse float")
			return
		}

		odd := models.Odd{
			Id:          items[0],
			Name:        items[1],
			Match:       items[2],
			Coefficient: coeff,
			Timestamp:   time.Time{},
		}

		select {
		case <-ctx.Done():
			return
		case a.updates <- odd:
			// do nothing
		}
	}

}
