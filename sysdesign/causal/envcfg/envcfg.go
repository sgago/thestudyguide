package envcfg

import (
	"github.com/kelseyhightower/envconfig"
)

type envcfg struct {
	Port        int    `envconfig:"PORT" default:"8080" required:"true"`
	HostName    string `envconfig:"HOSTNAME" default:"localhost" required:"false"`
	ServiceName string `envconfig:"SERVICE_NAME" default:"goginair" required:"false"`
	ServiceAddr string `envconfig:"SERVICE_ADDR" default:"8080" required:"false"`
	ConsulAddr  string `envconfig:"CONSUL_ADDRESS" default:"consul:8500" required:"false"`
	IsLocal     bool   `envconfig:"IS_ALONE" default:"true" required:"false"`
}

var (
	e envcfg
)

func init() {
	if err := envconfig.Process("", &e); err != nil {
		panic(err)
	}
}

func Port() int {
	return e.Port
}

func ServiceName() string {
	return e.ServiceName
}

func ServiceAddr() string {
	return e.ServiceAddr
}

func ConsulAddr() string {
	return e.ConsulAddr
}

func HostName() string {
	return e.HostName
}

func IsLocal() bool {
	return e.IsLocal
}
