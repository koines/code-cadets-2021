package bootstrap

import (
	"github.com/koines/code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http"
	stdhttp "net/http"
)

func AxilisOfferFeed() *http.AxilisOfferFeed {
	httpClient := &stdhttp.Client{}
	return http.NewAxilisOfferFeed(httpClient)
}
