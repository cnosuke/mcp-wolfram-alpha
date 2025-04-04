package tools

import (
	mcp "github.com/metoro-io/mcp-golang"
)

// RegisterAllTools - Register all tools with the server
func RegisterAllTools(mcpServer *mcp.Server, greeter Greeter) error {
	// Register greeting/hello tool
	if err := RegisterGreetingHelloTool(mcpServer, greeter); err != nil {
		return err
	}

	return nil
}
