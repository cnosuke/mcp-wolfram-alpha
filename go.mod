module github.com/cnosuke/mcp-wolfram-alpha

go 1.24.0

require (
	github.com/cnosuke/go-wolfram-llm v0.0.0-20250403152242-3894e63c16bd
	github.com/cockroachdb/errors v1.11.3
	github.com/jinzhu/configor v1.2.2
	github.com/metoro-io/mcp-golang v0.8.0
	github.com/stretchr/testify v1.9.0
	github.com/urfave/cli/v2 v2.27.6
	go.uber.org/zap v1.27.0
)

// This is a tentative version to be used as an MCP server on Cline.
// cf. https://github.com/cnosuke/mcp-golang/pull/1
replace github.com/metoro-io/mcp-golang v0.8.0 => github.com/cnosuke/mcp-golang v0.8.1

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cockroachdb/logtags v0.0.0-20241215232642-bb51bb14a506 // indirect
	github.com/cockroachdb/redact v1.1.6 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.6 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/getsentry/sentry-go v0.31.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/invopop/jsonschema v0.13.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.8 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
