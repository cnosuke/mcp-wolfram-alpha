package tools

import (
	"github.com/cockroachdb/errors"
	mcp "github.com/metoro-io/mcp-golang"
	"go.uber.org/zap"
)

// GreetingHelloArgs - Arguments for greeting/hello tool
type GreetingHelloArgs struct {
	Name string `json:"name" jsonschema:"description=Optional name for personalized greeting"`
}

// Greeter defines the interface for greeting generation
type Greeter interface {
	GenerateGreeting(name string) (string, error)
}

// RegisterGreetingHelloTool - Register the greeting/hello tool
func RegisterGreetingHelloTool(server *mcp.Server, greeter Greeter) error {
	zap.S().Debugw("registering greeting/hello tool")
	err := server.RegisterTool("greeting/hello", "Generate a greeting message",
		func(args GreetingHelloArgs) (*mcp.ToolResponse, error) {
			zap.S().Debugw("executing greeting/hello",
				"name", args.Name)

			// Generate greeting
			greeting, err := greeter.GenerateGreeting(args.Name)
			if err != nil {
				zap.S().Errorw("failed to generate greeting",
					"name", args.Name,
					"error", err)
				return nil, errors.Wrap(err, "failed to generate greeting")
			}

			return mcp.NewToolResponse(mcp.NewTextContent(greeting)), nil
		})

	if err != nil {
		zap.S().Errorw("failed to register greeting/hello tool", "error", err)
		return errors.Wrap(err, "failed to register greeting/hello tool")
	}

	return nil
}
