package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for GreetingHelloArgs
func TestGreetingHelloArgs(t *testing.T) {
	// When name is empty
	argsEmpty := GreetingHelloArgs{
		Name: "",
	}
	assert.Equal(t, "", argsEmpty.Name)

	// When name is set
	argsWithName := GreetingHelloArgs{
		Name: "Test User",
	}
	assert.Equal(t, "Test User", argsWithName.Name)
}
