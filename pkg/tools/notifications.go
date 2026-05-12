package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerNotificationTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_list_notification_channels",
		mcp.WithDescription("List all configured notification channels (email, Slack, PagerDuty, webhook, etc.)."),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V6("GET", "/notification_channel/v202210/notification_channels", nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("List notification channels failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_notification_channel",
		mcp.WithDescription("Get a specific notification channel by ID."),
		mcp.WithString("channel_id", mcp.Required(), mcp.Description("Channel ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("channel_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("GET", fmt.Sprintf("/notification_channel/v202210/notification_channels/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get notification channel failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_search_notification_channels",
		mcp.WithDescription("Search notification channels by name or type."),
		mcp.WithString("search_query", mcp.Description("Search term to filter by name.")),
		mcp.WithString("channel_type", mcp.Description("Filter by type: 'email', 'slack', 'pagerduty', 'webhook', etc.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		body := map[string]interface{}{}
		if v, err := req.RequireString("search_query"); err == nil && v != "" { body["search"] = v }
		if v, err := req.RequireString("channel_type"); err == nil && v != "" { body["type"] = v }
		data, err := client.V6("POST", "/notification_channel/v202210/notification_channels/search", body)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Search notification channels failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
