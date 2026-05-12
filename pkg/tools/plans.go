package tools

import (
	"context"
	"fmt"
	"net/url"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerPlanTools(s *server.MCPServer, client *kentik.Client) {
	// List plans
	s.AddTool(mcp.NewTool("kentik_list_plans",
		mcp.WithDescription("List all Kentik plans (service tiers) associated with your account."),
	), makeListPlansHandler(client))

	// Get audit log
	s.AddTool(mcp.NewTool("kentik_get_audit_log",
		mcp.WithDescription("Get audit log entries showing account activity. Optionally filter by date range or user."),
		mcp.WithString("start_date", mcp.Description("Start date filter in YYYY-MM-DD format.")),
		mcp.WithString("end_date", mcp.Description("End date filter in YYYY-MM-DD format.")),
		mcp.WithString("user_id", mcp.Description("Filter audit log entries by user ID.")),
	), makeGetAuditLogHandler(client))
}

func makeListPlansHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V5("GET", "/plans", nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("List plans failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeGetAuditLogHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := url.Values{}
		if v, err := request.RequireString("start_date"); err == nil && v != "" {
			params.Set("start_date", v)
		}
		if v, err := request.RequireString("end_date"); err == nil && v != "" {
			params.Set("end_date", v)
		}
		if v, err := request.RequireString("user_id"); err == nil && v != "" {
			params.Set("user_id", v)
		}

		path := "/audit-log"
		if len(params) > 0 {
			path = path + "?" + params.Encode()
		}

		data, err := client.V5("GET", path, nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Get audit log failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}
