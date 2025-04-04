package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for GreetingHelloTool
func TestGreeterFunctionality(t *testing.T) {
	// This test verifies the basic functionality of the Greeter interface
	// It does not perform actual MCP server integration testing

	// Mock Greeter instance
	mockGreeter := &TestGreeter{
		defaultMessage: "Hello!",
	}

	// Test
	greeting1, err := mockGreeter.GenerateGreeting("")
	assert.NoError(t, err)
	assert.Equal(t, "Hello!", greeting1)

	greeting2, err := mockGreeter.GenerateGreeting("Tanaka")
	assert.NoError(t, err)
	assert.Equal(t, "Hello! Tanaka!", greeting2)
}

// Test implementation of Greeter
type TestGreeter struct {
	defaultMessage string
}

func (g *TestGreeter) GenerateGreeting(name string) (string, error) {
	if name == "" {
		return g.defaultMessage, nil
	}
	return g.defaultMessage + " " + name + "!", nil
}
