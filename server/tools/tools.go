package tools

import (
	"github.com/mark3labs/mcp-go/server"
	"go.uber.org/zap"
)

// RegisterAllTools - Register all tools with the server
func RegisterAllTools(mcpServer *server.MCPServer, wolframServer WolframQueryer) error {
	zap.S().Infow("registering all tools")

	// Register wolfram_query tool
	if err := RegisterWolframQueryTool(mcpServer, wolframServer); err != nil {
		return err
	}

	zap.S().Infow("all tools registered successfully")
	return nil
}
