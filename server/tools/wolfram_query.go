package tools

import (
	"context"

	wolframllm "github.com/cnosuke/go-wolfram-llm"
	"github.com/cockroachdb/errors"
	mcp "github.com/metoro-io/mcp-golang"
	"go.uber.org/zap"
)

// WolframQueryArgs defines the arguments for the wolfram_query tool
type WolframQueryArgs struct {
	Query        string `json:"query" jsonschema:"required,description=The Wolfram Alpha query to execute"`
	MaxChars     int    `json:"max_chars" jsonschema:"description=Maximum characters in response (default: 2000)"`
	Units        string `json:"units" jsonschema:"enum=metric,nonmetric,description=Unit system to use (metric or nonmetric)"`
	CountryCode  string `json:"country_code" jsonschema:"description=Country code for localization (e.g., 'JP')"`
	LanguageCode string `json:"language_code" jsonschema:"description=Language code for localization (e.g., 'ja')"`
	ShowSteps    bool   `json:"show_steps" jsonschema:"description=Request step-by-step solution for math problems"`
}

// WolframQueryer defines the interface for Wolfram Alpha query execution
type WolframQueryer interface {
	ExecuteQuery(ctx context.Context, query string, options *wolframllm.QueryParams) (string, error)
}

// RegisterWolframQueryTool registers the wolfram_query tool with the MCP server
func RegisterWolframQueryTool(server *mcp.Server, wolframServer WolframQueryer) error {
	zap.S().Infow("registering wolfram_query tool")

	description := `Execute a Wolfram Alpha query to perform calculations or retrieve knowledge.
This leverages the high-precision Wolfram Alpha API engine to handle numerical calculations and symbolic computations – tasks where LLMs might struggle or produce uncertain results on their own. This allows the LLM to avoid computational errors and focus on delivering reliable information to the user.
For instance, it is effective not only for mathematical calculations like complex arithmetic operations, solving algebraic equations, and calculus, but also for referencing scientific knowledge and factual data such as physical constants, chemical properties, statistical data, and geographical information. Attempting these tasks internally within the LLM can lead to inefficient token consumption. Except for very simple calculations, when computation or accurate data retrieval is required, actively utilize this function to optimize token consumption and maximize the LLM's core capabilities.`

	err := server.RegisterTool("wolfram_query", description,
		func(args WolframQueryArgs) (*mcp.ToolResponse, error) {
			zap.S().Debugw("executing wolfram_query tool",
				"query", args.Query,
				"max_chars", args.MaxChars,
				"units", args.Units,
				"country_code", args.CountryCode,
				"language_code", args.LanguageCode,
				"show_steps", args.ShowSteps)

			// Input validation
			if args.Query == "" {
				err := errors.New("query cannot be empty")
				zap.S().Errorw("empty query provided", "error", err)
				return nil, err
			}

			// Create context
			ctx := context.Background()

			// Set query options
			options := &wolframllm.QueryParams{
				MaxChars:     args.MaxChars,
				Units:        args.Units,
				CountryCode:  args.CountryCode,
				LanguageCode: args.LanguageCode,
			}

			// Handle step-by-step option
			query := args.Query
			if args.ShowSteps {
				// Wolfram Alpha API has a special syntax for requesting steps
				query = "show steps " + query
			}

			// Execute query
			result, err := wolframServer.ExecuteQuery(ctx, query, options)
			if err != nil {
				zap.S().Errorw("failed to execute Wolfram Alpha query",
					"query", query,
					"error", err)
				return nil, errors.Wrap(err, "failed to execute Wolfram Alpha query")
			}

			zap.S().Debugw("wolfram_query executed successfully",
				"query", query,
				"result_length", len(result))

			// Return response
			return mcp.NewToolResponse(mcp.NewTextContent(result)), nil
		})

	if err != nil {
		zap.S().Errorw("failed to register wolfram_query tool", "error", err)
		return errors.Wrap(err, "failed to register wolfram_query tool")
	}

	return nil
}
