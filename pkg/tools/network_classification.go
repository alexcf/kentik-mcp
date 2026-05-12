package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerNetworkClassificationTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_get_network_classification",
		mcp.WithDescription("Get current network classification rules defining which IPs/prefixes are internal vs external."),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V6("GET", "/network_class/v202109alpha1/network_class", nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get network classification failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_update_network_classification",
		mcp.WithDescription("Replace the entire network classification configuration. Pass full classification JSON."),
		mcp.WithString("classification_json", mcp.Required(), mcp.Description("Full network classification JSON to replace the current config.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		raw, err := req.RequireString("classification_json")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		var body interface{}
		if err := json.Unmarshal([]byte(raw), &body); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Invalid JSON: %v", err)), nil
		}
		data, err := client.V6("POST", "/network_class/v202109alpha1/network_class", body)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Update network classification failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
