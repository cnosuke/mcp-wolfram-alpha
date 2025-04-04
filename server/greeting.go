package server

import (
	"github.com/cnosuke/mcp-greeting/config"
	"go.uber.org/zap"
)

// GreetingServer - Greeting server structure
type GreetingServer struct {
	DefaultMessage string
	cfg            *config.Config
}

// NewGreetingServer - Create a new Greeting server
func NewGreetingServer(cfg *config.Config) (*GreetingServer, error) {
	zap.S().Infow("creating new Greeting server",
		"default_message", cfg.Greeting.DefaultMessage)

	return &GreetingServer{
		DefaultMessage: cfg.Greeting.DefaultMessage,
		cfg:            cfg,
	}, nil
}

// GenerateGreeting - Generate a greeting message
func (s *GreetingServer) GenerateGreeting(name string) (string, error) {
	if name == "" {
		return s.DefaultMessage, nil
	}
	return s.DefaultMessage + " " + name + "!", nil
}
