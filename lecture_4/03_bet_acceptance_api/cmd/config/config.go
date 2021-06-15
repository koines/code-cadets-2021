package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Cfg is the single instance of configuration that gets automatically populated from the
// environment variables once the  module loads.
var Cfg Config

// Config contains all the configuration needed for service to work.
type Config struct {
	Rabbit   rabbitConfig `split_words:"true"`
	Api      apiConfig    `split_words:"true"`
	ConstVar paramConfig  `split_words:"true"`
}

type apiConfig struct {
	ReadWriteTimeoutMs int `split_words:"true" default:"10000"`
	Port               int `split_words:"true" default:"8082"`
}

type paramConfig struct {
	CoefficientBound  float64 `split_words:"true" default:"10.0"`
	PaymentLowerBound float64 `split_words:"true" default:"2.0"`
	PaymentUpperBound float64 `split_words:"true" default:"100.0"`
}

type rabbitConfig struct {
	PublisherBetQueueQueue string `split_words:"true"  default:"bets-received"`
	PublisherExchange      string `split_words:"true" default:""`
	PublisherMandatory     bool   `split_words:"true" default:"false"`
	PublisherImmediate     bool   `split_words:"true" default:"false"`
	ConnectionQueue        string `split_words:"true" default:"amqp://guest:guest@localhost:5672/"`
}

// Load loads the configuration on bootstrap, this avoid injecting the same config object
// everywhere.
func Load() {
	err := envconfig.Process("", &Cfg)
	if err != nil {
		panic(err)
	}
}
