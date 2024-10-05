package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	RaftId        string   `envconfig:"RAFT_ID" default:"node1" required:"true"`
	RaftAddress   string   `envconfig:"RAFT_ADDRESS" default:"localhost:3001" required:"true"`
	RaftNodes     []string `envconfig:"RAFT_NODES" default:"localhost:3001" required:"true"`
	RaftBootstrap bool     `envconfig:"RAFT_BOOTSTRAP" default:"true" required:"true"`
}

var (
	cfg Config
)

func init() {
	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}
}

func RaftId() string {
	return cfg.RaftId
}

func RaftAddress() string {
	return cfg.RaftAddress
}

func RaftNodes() []string {
	return cfg.RaftNodes
}

func RaftBootstrap() bool {
	return cfg.RaftBootstrap
}
