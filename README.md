# MCP Wolfram Alpha Server

MCP Wolfram Alpha Server is a Go-based MCP server that provides access to the Wolfram Alpha API. It allows MCP clients (e.g., Claude Desktop) to execute computational queries, access scientific data, and retrieve structured knowledge.

## Purpose and Benefits of this MCP Server

### Purpose

This MCP server delegates **numerical computation tasks** that large language models (LLMs) struggle with to the high-precision computation engine **Wolfram Alpha**.

### Understanding LLM Limitations

LLMs demonstrate remarkable capabilities in natural language processing, text generation, and translation, but their architecture is not optimized for precise numerical calculations or formula processing.

- **Calculation Accuracy:** Beyond basic mental arithmetic (e.g., 2+3), LLMs may make calculation errors, especially with multi-digit numbers or complex formulas (calculus, matrix calculations, etc.).
- **Token Consumption:** When LLMs attempt complex calculations internally, they require many steps and inferences, potentially **consuming valuable tokens (computational resources) inefficiently**. In some cases, they may reach token limits before arriving at the correct answer.

### Role and Benefits of this MCP Server

This MCP server addresses these challenges:

- **High-Precision Calculations:** When an LLM receives a task requiring calculation, it sends the computational part as a request to this MCP server. The server forwards the request to the Wolfram Alpha API and returns accurate calculation results to the LLM, providing users with reliable computational outputs.
- **Reduced LLM Workload:** LLMs no longer need to perform calculations themselves, allowing them to focus on their strengths in language processing and contextual understanding.
- **Optimized Token Usage:** **The server prevents inefficient token consumption that would occur if LLMs attempted uncertain calculations internally.** By delegating computational tasks to an external API (Wolfram Alpha via this MCP server), token usage is limited to API request and response exchanges, enabling more efficient processing overall.

### Recommended Use Cases

LLMs are encouraged to **actively utilize this MCP server** when they need to perform calculations like:

- Arithmetic operations (especially those beyond simple mental arithmetic, such as those with many digits or decimals)
- Solving algebraic equations
- Differentiation and integration
- Linear algebra (matrix calculations, etc.)
- Statistical calculations
- Calculations involving specific formulas in physics or chemistry
- Unit and currency conversions (especially when current rates are needed)

**However, there's no need to query external APIs for extremely simple calculations like `1+1` that are obviously within the LLM's capabilities.** For calculations that LLMs can handle reliably, they should proceed directly.

**In conclusion, this MCP server is an essential support tool that helps LLMs overcome challenges in calculation accuracy and token efficiency, enabling more reliable and efficient responses. When calculations are needed (except for very simple ones), actively use this MCP server API to avoid unnecessary token consumption.**

## Features

- MCP Compliance: Implements a JSON-RPC based interface according to the MCP specification
- Wolfram Alpha Integration: Provides access to mathematical computation, scientific data, and knowledge queries
- Configurable Options: Supports unit systems, regional settings, and language options

## Requirements

- Go 1.24 or later
- Wolfram Alpha API ID (obtainable from [Wolfram Alpha Developer Portal](https://developer.wolframalpha.com/))

## Configuration

The server is configured via a YAML file (default: config.yml):

```yaml
log: 'path/to/mcp-wolfram-alpha.log' # Log file path (empty for no logging)
debug: false # Enable debug mode

wolfram:
  app_id: 'YOUR_WOLFRAM_ALPHA_APP_ID' # Required: Wolfram Alpha API ID
  timeout: 30 # API timeout in seconds
  use_bearer: false # Use Bearer token authentication
  default_max_chars: 2000 # Default maximum characters in responses
```

You can override configurations using environment variables:

- `LOG_PATH`: Path to log file
- `DEBUG`: Enable debug mode (true/false)
- `WOLFRAM_APP_ID`: Wolfram Alpha API ID
- `WOLFRAM_TIMEOUT`: Timeout in seconds
- `WOLFRAM_USE_BEARER`: Use Bearer authentication (true/false)
- `WOLFRAM_DEFAULT_MAX_CHARS`: Default maximum character count

## Building and Running

```bash
# Download dependencies
make deps

# Build the server
make build

# Run the server
./bin/mcp-wolfram-alpha server --config config.yml
```

## MCP Tools

The following MCP tools are implemented:

- `wolfram_query`: Execute Wolfram Alpha queries with options for customization

### Tool Arguments

The `wolfram_query` tool accepts the following arguments:

```json
{
  "query": "integrate x^2",
  "max_chars": 2000,
  "units": "metric",
  "country_code": "JP",
  "language_code": "en",
  "show_steps": true
}
```

- `query` (required): The Wolfram Alpha query to execute
- `max_chars`: Maximum characters in response (default: 2000)
- `units`: Unit system to use (`metric` or `nonmetric`)
- `country_code`: Country code for localization (e.g., 'JP')
- `language_code`: Language code for localization (e.g., 'ja')
- `show_steps`: Request step-by-step solution for math problems (boolean)

## Using with Claude Desktop

To integrate with Claude Desktop, edit your `claude_desktop_config.json` file:

```json
{
  "mcpServers": {
    "wolfram-alpha": {
      "command": "/path/to/bin/mcp-wolfram-alpha",
      "args": ["server", "--config", "/path/to/config.yml"],
      "env": {
        "LOG_PATH": "/path/to/logs/mcp-wolfram.log",
        "WOLFRAM_APP_ID": "YOUR_WOLFRAM_ALPHA_APP_ID"
      }
    }
  }
}
```

## Example Usage

With Claude Desktop properly configured, you can ask Claude questions like:

- "What is the derivative of x^3?"
- "Calculate the distance from Earth to Mars"
- "What is the atomic weight of gold?"
- "Convert 100 kilometers to miles"
- "Solve the equation x^2 + 3x - 4 = 0"

Claude will automatically use the Wolfram Alpha API through this MCP server to compute answers.

## Error Handling

The server provides informative error messages for various failure scenarios:

- Authentication errors (invalid API ID)
- Invalid input errors
- Network connection issues
- Timeout errors
- Server-side Wolfram Alpha errors

All errors are logged with detailed information to help with troubleshooting.

## License

This project is licensed under the MIT License.

## Author

cnosuke (github.com/cnosuke)
