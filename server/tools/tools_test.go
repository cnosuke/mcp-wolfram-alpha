package tools

import (
	"context"
	"testing"

	wolframllm "github.com/cnosuke/go-wolfram-llm"
	"github.com/stretchr/testify/assert"
)

// Test for WolframQueryer functionality
func TestWolframQueryerFunctionality(t *testing.T) {
	// This test verifies the basic functionality of the WolframQueryer interface
	// It does not perform actual API calls to Wolfram Alpha

	// Mock WolframQueryer implementation
	mockWolframServer := &TestWolframQueryer{
		responses: map[string]string{
			"integrate x^2":  "x^3/3 + C",
			"population of Tokyo": "13.96 million people (2023 estimate)",
			"distance from Earth to Mars": "54.6 million kilometers (minimum), 401 million kilometers (maximum)",
		},
	}

	// Test basic query
	result1, err := mockWolframServer.ExecuteQuery(context.Background(), "integrate x^2", nil)
	assert.NoError(t, err)
	assert.Equal(t, "x^3/3 + C", result1)

	// Test with options
	options := &wolframllm.QueryParams{
		MaxChars: 1000,
		Units:    "metric",
	}
	result2, err := mockWolframServer.ExecuteQuery(context.Background(), "population of Tokyo", options)
	assert.NoError(t, err)
	assert.Equal(t, "13.96 million people (2023 estimate)", result2)

	// Test non-existent query
	result3, err := mockWolframServer.ExecuteQuery(context.Background(), "non existent query", nil)
	assert.NoError(t, err)
	assert.Equal(t, "No result found", result3)
}

// TestWolframQueryer is a mock implementation of WolframQueryer
type TestWolframQueryer struct {
	responses map[string]string
}

// ExecuteQuery implements the WolframQueryer interface
func (q *TestWolframQueryer) ExecuteQuery(_ context.Context, query string, _ *wolframllm.QueryParams) (string, error) {
	// Return pre-defined response if exists
	if result, exists := q.responses[query]; exists {
		return result, nil
	}
	
	// Default response for unknown queries
	return "No result found", nil
}
