package config

import "github.com/kelseyhightower/envconfig"

type config struct {
	HostName  string   `envconfig:"HOSTNAME"`
	ZkServers []string `envconfig:"ZOO_SERVERS" required:"true"`
}

var (
	cfg config
)

func init() {
	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}
}

func HostName() string {
	return cfg.HostName
}

func ZkServers() []string {
	return cfg.ZkServers
}
