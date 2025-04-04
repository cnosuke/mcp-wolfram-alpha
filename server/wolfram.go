package server

import (
	"context"

	wolframllm "github.com/cnosuke/go-wolfram-llm"
	"github.com/cockroachdb/errors"
	"go.uber.org/zap"

	"github.com/cnosuke/mcp-wolfram-alpha/config"
)

// WolframServer manages connections to the Wolfram Alpha API
type WolframServer struct {
	client *wolframllm.Client
	cfg    *config.Config
}

// NewWolframServer creates a new Wolfram Alpha server instance
func NewWolframServer(cfg *config.Config) (*WolframServer, error) {
	zap.S().Infow("creating new Wolfram Alpha server",
		"app_id_set", cfg.Wolfram.AppID != "",
		"timeout", cfg.Wolfram.Timeout,
		"use_bearer", cfg.Wolfram.UseBearer)

	if cfg.Wolfram.AppID == "" {
		return nil, errors.New("Wolfram Alpha AppID is required")
	}

	// Create Wolfram Alpha client
	client, err := wolframllm.NewClient(
		cfg.Wolfram.AppID,
		wolframllm.WithTimeout(cfg.Wolfram.Timeout),
		wolframllm.WithBearer(cfg.Wolfram.UseBearer),
		wolframllm.WithDefaultMaxChars(cfg.Wolfram.DefaultMaxChars),
		wolframllm.WithUserAgent("MCP-Wolfram-Alpha-Server/1.0"),
	)
	if err != nil {
		zap.S().Errorw("failed to create Wolfram Alpha client", "error", err)
		return nil, errors.Wrap(err, "failed to create Wolfram Alpha client")
	}

	return &WolframServer{
		client: client,
		cfg:    cfg,
	}, nil
}

// ExecuteQuery executes a query against the Wolfram Alpha API
func (s *WolframServer) ExecuteQuery(ctx context.Context, query string, options *wolframllm.QueryParams) (string, error) {
	zap.S().Debugw("executing Wolfram Alpha query",
		"query", query,
		"max_chars", getMaxChars(options, s.cfg))

	// Create options if nil
	if options == nil {
		options = &wolframllm.QueryParams{}
	}

	// Execute API call
	var response *wolframllm.LLMResponse
	var err error

	if options.Input == "" {
		options.Input = query
		response, err = s.client.QueryWithParams(ctx, query, options)
	} else {
		response, err = s.client.QueryWithParams(ctx, query, options)
	}

	// Handle errors
	if err != nil {
		zap.S().Errorw("Wolfram Alpha query failed",
			"query", query,
			"error", err)

		// Handle specific error types
		if wolframllm.IsAuthError(err) {
			return "", errors.Wrap(err, "authentication error with Wolfram Alpha API")
		} else if wolframllm.IsInvalidInputError(err) {
			return "", errors.Wrap(err, "invalid input for Wolfram Alpha API")
		} else if wolframllm.IsServerError(err) {
			return "", errors.Wrap(err, "Wolfram Alpha server error")
		} else if wolframllm.IsNetworkError(err) {
			return "", errors.Wrap(err, "network error while connecting to Wolfram Alpha")
		}

		return "", errors.Wrap(err, "failed to execute Wolfram Alpha query")
	}

	zap.S().Debugw("received Wolfram Alpha response",
		"query", query,
		"result_length", len(response.Result))

	return response.Result, nil
}

// getMaxChars returns the maximum character count from options or default
func getMaxChars(options interface{}, cfg *config.Config) int {
	if options != nil {
		// Use type assertion to check if options has a MaxChars field
		switch opts := options.(type) {
		case *wolframllm.QueryParams:
			if opts.MaxChars > 0 {
				return opts.MaxChars
			}
		}
	}
	return cfg.Wolfram.DefaultMaxChars
}
