package consul

import (
	"fmt"
	"log"
	"time"

	"sgago/thestudyguide-causal/envcfg"

	consulapi "github.com/hashicorp/consul/api"
)

type Client interface {
	Register() error
	StartTTL()
	StopTTL()
	Self() string
	Services() []string
}

type ConsulCli struct {
	api       *consulapi.Client
	serviceId string
	checkId   string
	ticker    *time.Ticker
}

var _ Client = (*ConsulCli)(nil)

func New() (*ConsulCli, error) {
	config := consulapi.DefaultConfig()
	config.Address = envcfg.ConsulAddr()

	client, err := consulapi.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("error creating Consul client: %v", err)
	}

	return &ConsulCli{
		api:       client,
		serviceId: envcfg.HostName(),
		checkId:   fmt.Sprintf("%s-check", envcfg.HostName()),
		ticker:    time.NewTicker(5 * time.Second),
	}, nil
}

func (consul *ConsulCli) Register() error {
	registration := &consulapi.AgentServiceRegistration{
		ID:      consul.serviceId,
		Name:    envcfg.ServiceName(),
		Port:    envcfg.Port(),
		Address: envcfg.HostName(),
		Tags:    []string{"http"},
		Check: &consulapi.AgentServiceCheck{
			CheckID:                        consul.checkId,
			TTL:                            "10s",
			DeregisterCriticalServiceAfter: "1m",
		},
	}

	if err := consul.api.Agent().ServiceRegister(registration); err != nil {
		return fmt.Errorf("error registering service with Consul: %v", err)
	}

	log.Println("Service registered successfully with Consul")
	return nil
}

func (consul *ConsulCli) StartTTL() {
	go func() {
		for range consul.ticker.C {
			if err := consul.api.Agent().UpdateTTL(consul.checkId, "TTL OK", consulapi.HealthPassing); err != nil {
				log.Printf("Error updating TTL: %v", err)
			}
		}
	}()
}

func (consul *ConsulCli) StopTTL() {
	consul.ticker.Stop()
}

func (consul *ConsulCli) Self() string {
	return fmt.Sprintf("http://%s:%d", envcfg.HostName(), envcfg.Port())
}

func (consul *ConsulCli) Services() []string {
	services, err := consul.api.Agent().Services()
	if err != nil {
		fmt.Printf("Error getting services: %v", err)
	}

	svc := make([]string, 0, 3)

	for _, service := range services {
		if service.Service == envcfg.ServiceName() {
			svc = append(svc, fmt.Sprintf("http://%s:%d", service.ID, service.Port))
		}
	}

	return svc
}
