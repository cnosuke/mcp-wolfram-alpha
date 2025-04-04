package server

import (
	"testing"

	"github.com/cnosuke/mcp-wolfram-alpha/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

// TestNewWolframServer - Test initialization of WolframServer with valid config
func TestNewWolframServer(t *testing.T) {
	// Set up test logger
	logger := zaptest.NewLogger(t)
	zap.ReplaceGlobals(logger)

	// Test configuration with valid AppID
	cfg := &config.Config{}
	cfg.Wolfram.AppID = "test-app-id"
	cfg.Wolfram.Timeout = 30
	cfg.Wolfram.DefaultMaxChars = 2000

	// Create server
	server, err := NewWolframServer(cfg)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, server)
}

// TestNewWolframServerMissingAppID - Test initialization with missing AppID
func TestNewWolframServerMissingAppID(t *testing.T) {
	// Set up test logger
	logger := zaptest.NewLogger(t)
	zap.ReplaceGlobals(logger)

	// Test configuration with missing AppID
	cfg := &config.Config{}
	cfg.Wolfram.AppID = "" // Missing AppID

	// Create server
	server, err := NewWolframServer(cfg)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, server)
	assert.Contains(t, err.Error(), "AppID is required")
}

// TestGetMaxChars - Test the getMaxChars function
func TestGetMaxChars(t *testing.T) {
	// Set up configuration
	cfg := &config.Config{}
	cfg.Wolfram.DefaultMaxChars = 2000

	// Test with nil options
	assert.Equal(t, 2000, getMaxChars(nil, cfg))

	// Import QueryParams directly from wolframllm
	// But for testing without actually calling the API
	wolframOptions := &struct{ MaxChars int }{MaxChars: 0}
	
	// Since our options is not the real wolframllm.QueryParams,
	// it will return the default value
	assert.Equal(t, 2000, getMaxChars(wolframOptions, cfg))
}
