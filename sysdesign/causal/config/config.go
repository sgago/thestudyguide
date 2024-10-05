package config

import "github.com/kelseyhightower/envconfig"

type config struct {
	MyName     string   `envconfig:"MY_NAME" default:"localhost" required:"true"`
	OtherNames []string `envconfig:"OTHER_NAMES" default:"localhost" required:"false"`
}

var (
	cfg config
)

func init() {
	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}
}

func MyName() string {
	return cfg.MyName
}

func OtherNames() []string {
	return cfg.OtherNames
}
