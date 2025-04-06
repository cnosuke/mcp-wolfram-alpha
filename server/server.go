package server

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"go.uber.org/zap"

	"github.com/cnosuke/mcp-wolfram-alpha/config"
	"github.com/cnosuke/mcp-wolfram-alpha/server/tools"
	"github.com/cockroachdb/errors"
)

// Run - Execute the MCP server
func Run(cfg *config.Config, name string, version string, revision string) error {
	zap.S().Infow("starting MCP Wolfram Alpha Server")

	// Format version string with revision if available
	versionString := version
	if revision != "" && revision != "xxx" {
		versionString = versionString + " (" + revision + ")"
	}

	// Create Wolfram Alpha server
	zap.S().Debugw("creating Wolfram Alpha server")
	wolframServer, err := NewWolframServer(cfg)
	if err != nil {
		zap.S().Errorw("failed to create Wolfram Alpha server", "error", err)
		return err
	}

	// Create custom hooks for error handling
	hooks := &server.Hooks{}
	hooks.AddOnError(func(ctx context.Context, id any, method mcp.MCPMethod, message any, err error) {
		zap.S().Errorw("MCP error occurred",
			"id", id,
			"method", method,
			"error", err,
		)
	})

	// Create MCP server with server name and version
	zap.S().Debugw("creating MCP server",
		"name", name,
		"version", versionString,
	)
	mcpServer := server.NewMCPServer(
		name,
		versionString,
		server.WithHooks(hooks),
	)

	// Register all tools
	zap.S().Debugw("registering tools")
	if err := tools.RegisterAllTools(mcpServer, wolframServer); err != nil {
		zap.S().Errorw("failed to register tools", "error", err)
		return err
	}

	// Start the server with stdio transport
	zap.S().Infow("starting MCP server")
	err = server.ServeStdio(mcpServer)
	if err != nil {
		zap.S().Errorw("failed to start server", "error", err)
		return errors.Wrap(err, "failed to start server")
	}

	// ServeStdio will block until the server is terminated
	zap.S().Infow("server shutting down")
	return nil
}
