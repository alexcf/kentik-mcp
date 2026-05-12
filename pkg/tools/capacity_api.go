package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerCapacityAPITools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_list_capacity_plans",
		mcp.WithDescription("List all capacity plans from the Kentik Capacity Planning API."),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V6("GET", "/capacity_plan/v202212/capacity_plan", nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("List capacity plans failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_capacity_plan",
		mcp.WithDescription("Get a specific capacity plan by ID."),
		mcp.WithString("plan_id", mcp.Required(), mcp.Description("Capacity plan ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("plan_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("GET", fmt.Sprintf("/capacity_plan/v202212/capacity_plan/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get capacity plan failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_capacity_plan_summary",
		mcp.WithDescription("Get summary metrics across all capacity plans."),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V6("GET", "/capacity_plan/v202212/capacity_plan/summary", nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get capacity plan summary failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_capacity_plan_detail",
		mcp.WithDescription("Get detailed capacity summary for a specific plan."),
		mcp.WithString("plan_id", mcp.Required(), mcp.Description("Capacity plan ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("plan_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("GET", fmt.Sprintf("/capacity_plan/v202212/capacity_plan/%s/summary", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get capacity plan detail failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
