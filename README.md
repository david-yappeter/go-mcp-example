## Go MCP Example

This is a simple calculator example using the [MCP](https://github.com/mark3labs/mcp-go) library.

## Prerequisites

- [Go](https://golang.org/dl/)
- [MCP](https://github.com/mark3labs/mcp-go)

## Assigning MCP Config to Windsurf

In the `mcp_config.json` file, add the following:

This is for a local development

```json
{
  "mcpServers": {
    "hello-world": {
      "command": "sh",
      "args": ["-c", "cd /home/author/go/go-mcp && go run main.go"],
      "env": {
        "THIS_IS_A_ENV": "ASJGJAJGIAJIiasgjiai"
      }
    }
  }
}
```

For binary version of MCP, add the following:

```json
{
  "mcpServers": {
    "hello-world": {
      "command": "/path/to/mcp_binary"
    }
  }
}
```

