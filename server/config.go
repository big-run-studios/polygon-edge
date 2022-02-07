package server

import (
	"github.com/hashicorp/go-hclog"
	"net"

	"github.com/0xPolygon/polygon-edge/chain"
	"github.com/0xPolygon/polygon-edge/network"
	"github.com/0xPolygon/polygon-edge/secrets"
)

const DefaultGRPCPort int = 9632
const DefaultJSONRPCPort int = 8545
const DefaultBlockTime = 2 // in seconds

// Config is used to parametrize the minimal client
type Config struct {
	Chain *chain.Chain

	JSONRPCAddr *net.TCPAddr
	GRPCAddr    *net.TCPAddr
	LibP2PAddr  *net.TCPAddr

	PriceLimit uint64
	MaxSlots   uint64
	BlockTime  uint64

	Telemetry *Telemetry
	Network   *network.Config

	DataDir     string
	RestoreFile *string

	Seal bool

	SecretsManager *secrets.SecretsManagerConfig

	LogLevel hclog.Level
}

// DefaultConfig returns the default config for JSON-RPC, GRPC (ports) and Networking
func DefaultConfig() *Config {
	return &Config{
		JSONRPCAddr:    &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: DefaultJSONRPCPort},
		GRPCAddr:       &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: DefaultGRPCPort},
		Network:        network.DefaultConfig(),
		Telemetry:      &Telemetry{PrometheusAddr: nil},
		SecretsManager: nil,
		BlockTime:      DefaultBlockTime,
	}
}

// Telemetry holds the config details for metric services
type Telemetry struct {
	PrometheusAddr *net.TCPAddr
}
