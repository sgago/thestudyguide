package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Port       string   `envconfig:"PORT" default:"8080" required:"true"`
	MyName     string   `envconfig:"MY_NAME" default:"localhost" required:"true"`
	MyId       int      `envconfig:"MY_ID" default:"1" required:"true"`
	OtherNames []string `envconfig:"OTHER_NAMES" default:"localhost" required:"false"`
	Hosts      []string
}

var (
	cfg config
)

func init() {
	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}

	for _, other := range OtherNames() {
		cfg.Hosts = append(cfg.Hosts, fmt.Sprintf("http://%s:%s", other, cfg.Port))
	}
}

func MyName() string {
	return cfg.MyName
}

func MyId() int {
	return cfg.MyId
}

func OtherNames() []string {
	return cfg.OtherNames
}

func Port() string {
	return cfg.Port
}

func Hosts() []string {
	return cfg.Hosts
}
