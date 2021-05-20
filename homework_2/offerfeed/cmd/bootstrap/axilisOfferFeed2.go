package bootstrap

import (
	"github.com/koines/code-cadets-2021/homework_2/offerfeed/internal/infrastructure/http"
	stdhttp "net/http"
)

func AxilisOfferFeed2() *http.AxilisOfferFeed2 {
	httpClient := &stdhttp.Client{}
	return http.NewAxilisOfferFeed2(httpClient)
}
