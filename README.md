# MCP Greeting Server

MCP Greeting Server is a Go-based MCP server implementation that provides basic greeting functionality, allowing MCP clients (e.g., Claude Desktop) to generate greeting messages.

## Features

* MCP Compliance: Provides a JSON‐RPC based interface for tool execution according to the MCP specification.
* Greeting Operations: Supports generating greeting messages, with options for personalization.

## Requirements

* Go 1.24 or later

## Configuration

The server is configured via a YAML file (default: config.yml). For example:

```yaml
log: 'path/to/mcp-greeting.log' # Log file path, if empty no log will be produced
debug: false # Enable debug mode for verbose logging

greeting:
  default_message: "こんにちは！"
```

Note: The default greeting message can also be injected via an environment variable `GREETING_DEFAULT_MESSAGE`. If this environment variable is set, it will override the value in the configuration file.

You can override configurations using environment variables:
- `LOG_PATH`: Path to log file
- `DEBUG`: Enable debug mode (true/false)
- `GREETING_DEFAULT_MESSAGE`: Default greeting message

## Logging

Logging behavior is controlled through configuration:

- If `log` is set in the config file, logs will be written to the specified file
- If `log` is empty, no logs will be produced
- Set `debug: true` for more verbose logging

## MCP Server Usage

MCP clients interact with the server by sending JSON‐RPC requests to execute various tools. The following MCP tools are supported:

* `greeting/hello`: Generates a greeting message, with an optional name parameter for personalization.

### Using with Claude Desktop

To integrate with Claude Desktop, add an entry to your `claude_desktop_config.json` file. Because MCP uses stdio for communication, you must redirect logs away from stdio by using the `--no-logs` and `--log` flags:

```json
{
  "mcpServers": {
    "greeting": {
      "command": "./bin/mcp-greeting",
      "args": ["server"],
      "env": {
        "LOG_PATH": "mcp-greeting.log",
        "DEBUG": "false",
        "GREETING_DEFAULT_MESSAGE": "こんにちは"
      }
    }
  }
}
```

This configuration registers the MCP Greeting Server with Claude Desktop, ensuring that all logs are directed to the specified log file.

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests for improvements or bug fixes. For major changes, open an issue first to discuss your ideas.

## License

This project is licensed under the MIT License.

Author: cnosuke ( x.com/cnosuke )
