package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerCredentialTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_list_credential_groups",
		mcp.WithDescription("List all credential groups in the Kentik vault. Used for NMS SNMP and device access credentials."),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V6("GET", "/credential/v202407alpha1/group", nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("List credential groups failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_credential_group",
		mcp.WithDescription("Get details of a specific credential group."),
		mcp.WithString("group_id", mcp.Required(), mcp.Description("Credential group ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("group_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("GET", fmt.Sprintf("/credential/v202407alpha1/group/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get credential group failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
