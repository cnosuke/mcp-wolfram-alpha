package server

import (
	"testing"

	"github.com/cnosuke/mcp-greeting/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

// TestNewGreetingServer - Test initialization of GreetingServer
func TestNewGreetingServer(t *testing.T) {
	// Set up test logger
	logger := zaptest.NewLogger(t)
	zap.ReplaceGlobals(logger)

	// Test configuration
	cfg := &config.Config{}
	cfg.Greeting.DefaultMessage = "Test greeting"

	// Create server
	server, err := NewGreetingServer(cfg)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, server)
	assert.Equal(t, "Test greeting", server.DefaultMessage)
}

// TestSetupServerComponents - Test server setup logic
func TestSetupServerComponents(t *testing.T) {
	// Set up test logger
	logger := zaptest.NewLogger(t)
	zap.ReplaceGlobals(logger)

	// Test configuration
	cfg := &config.Config{}
	cfg.Greeting.DefaultMessage = "Test greeting"

	// Create and test server
	greetingServer, err := NewGreetingServer(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, greetingServer)

	// Test greeting generation functionality
	greeting, err := greetingServer.GenerateGreeting("Test User")
	assert.NoError(t, err)
	assert.Equal(t, "Test greeting Test User!", greeting)
}
