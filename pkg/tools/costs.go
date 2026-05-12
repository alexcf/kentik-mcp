package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerCostTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_list_cost_providers",
		mcp.WithDescription("List all connectivity cost providers (transit, peering, etc.)."),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V6("GET", "/cost/v202308/cost/providers", nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("List cost providers failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_cost_summaries",
		mcp.WithDescription("Get cost summaries across all providers showing traffic volume and estimated cost."),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V6("GET", "/cost/v202308/cost/summary", nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get cost summaries failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_cost_summary",
		mcp.WithDescription("Get cost summary for a specific provider."),
		mcp.WithString("provider_id", mcp.Required(), mcp.Description("Cost provider ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("provider_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("GET", fmt.Sprintf("/cost/v202308/cost/summary/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get cost summary failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
