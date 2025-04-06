package tools

import (
	"context"

	wolframllm "github.com/cnosuke/go-wolfram-llm"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"go.uber.org/zap"
)

// WolframQueryArgs defines the arguments for the wolfram_query tool (kept for testing compatibility)
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
func RegisterWolframQueryTool(mcpServer *server.MCPServer, wolframServer WolframQueryer) error {
	zap.S().Infow("registering wolfram_query tool")

	description := `Execute a Wolfram Alpha query to perform calculations or retrieve knowledge.
This leverages the high-precision Wolfram Alpha API engine to handle numerical calculations and symbolic computations – tasks where LLMs might struggle or produce uncertain results on their own. This allows the LLM to avoid computational errors and focus on delivering reliable information to the user.
For instance, it is effective not only for mathematical calculations like complex arithmetic operations, solving algebraic equations, and calculus, but also for referencing scientific knowledge and factual data such as physical constants, chemical properties, statistical data, and geographical information. Attempting these tasks internally within the LLM can lead to inefficient token consumption. Except for very simple calculations, when computation or accurate data retrieval is required, actively utilize this function to optimize token consumption and maximize the LLM's core capabilities.
When requesting a calculation, a query should be structured for maximum reliability, preferably by embedding numerical values directly into the formula whenever possible.

If variables are used for conciseness, the query should follow the format "compute <formula> where <single-char variable>=<value>, <single-char variable>=<value>, ...", using only single-character variable names and strongly avoiding characters with potentially specific meanings in scientific computing (like 'i', which can be misinterpreted as the imaginary unit).
Variables should be a single character; more than two characters are error-prone.
Example: "compute x^3+y^2+z where x=3, y=2, z=1"`

	// Define the tool with parameters
	tool := mcp.NewTool("wolfram_query",
		mcp.WithDescription(description),
		mcp.WithString("query",
			mcp.Description("The Wolfram Alpha query to execute"),
			mcp.Required(),
		),
		mcp.WithNumber("max_chars",
			mcp.Description("Maximum characters in response (default: 2000)"),
		),
		mcp.WithString("units",
			mcp.Description("Unit system to use (metric or nonmetric)"),
			mcp.Enum("metric", "nonmetric"),
		),
		mcp.WithString("country_code",
			mcp.Description("Country code for localization (e.g., 'JP')"),
		),
		mcp.WithString("language_code",
			mcp.Description("Language code for localization (e.g., 'ja')"),
		),
		mcp.WithBoolean("show_steps",
			mcp.Description("Request step-by-step solution for math problems"),
		),
	)

	// Add the tool handler
	mcpServer.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Extract parameters
		var query string
		var maxChars int
		var units string
		var countryCode string
		var languageCode string
		var showSteps bool

		// Get query parameter (required)
		if queryVal, ok := request.Params.Arguments["query"].(string); ok {
			query = queryVal
		} else {
			zap.S().Errorw("query parameter is required")
			return mcp.NewToolResultError("query parameter is required"), nil
		}

		// Get optional parameters
		if maxCharsVal, ok := request.Params.Arguments["max_chars"].(float64); ok {
			maxChars = int(maxCharsVal)
		}
		if unitsVal, ok := request.Params.Arguments["units"].(string); ok {
			units = unitsVal
		}
		if countryCodeVal, ok := request.Params.Arguments["country_code"].(string); ok {
			countryCode = countryCodeVal
		}
		if languageCodeVal, ok := request.Params.Arguments["language_code"].(string); ok {
			languageCode = languageCodeVal
		}
		if showStepsVal, ok := request.Params.Arguments["show_steps"].(bool); ok {
			showSteps = showStepsVal
		}

		zap.S().Debugw("executing wolfram_query tool",
			"query", query,
			"max_chars", maxChars,
			"units", units,
			"country_code", countryCode,
			"language_code", languageCode,
			"show_steps", showSteps)

		// Input validation
		if query == "" {
			errMsg := "query cannot be empty"
			zap.S().Errorw(errMsg)
			return mcp.NewToolResultError(errMsg), nil
		}

		// Set query options
		options := &wolframllm.QueryParams{
			MaxChars:     maxChars,
			Units:        units,
			CountryCode:  countryCode,
			LanguageCode: languageCode,
		}

		// Handle step-by-step option
		if showSteps {
			// Wolfram Alpha API has a special syntax for requesting steps
			query = "show steps " + query
		}

		// Execute query
		result, err := wolframServer.ExecuteQuery(ctx, query, options)
		if err != nil {
			errMsg := "failed to execute Wolfram Alpha query: " + err.Error()
			zap.S().Errorw(errMsg,
				"query", query,
				"error", err)
			return mcp.NewToolResultError(errMsg), nil
		}

		zap.S().Debugw("wolfram_query executed successfully",
			"query", query,
			"result_length", len(result))

		// Return response
		return mcp.NewToolResultText(result), nil
	})

	return nil
}
