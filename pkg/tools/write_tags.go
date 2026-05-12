package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerWriteTagTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_create_tag",
		mcp.WithDescription("Create a new flow tag in Kentik for traffic classification."),
		mcp.WithString("name", mcp.Required(), mcp.Description("Tag name.")),
		mcp.WithString("value", mcp.Required(), mcp.Description("Tag value.")),
		mcp.WithString("description", mcp.Description("Optional description.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, err := req.RequireString("name")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		value, err := req.RequireString("value")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		tag := map[string]interface{}{"name": name, "value": value}
		if v, err := req.RequireString("description"); err == nil && v != "" { tag["description"] = v }
		data, err := client.V6("POST", "/flow_tag/v202404alpha1/tag", map[string]interface{}{"flowTag": tag})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Create tag failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_update_tag",
		mcp.WithDescription("Update an existing flow tag."),
		mcp.WithString("tag_id", mcp.Required(), mcp.Description("Tag ID to update.")),
		mcp.WithString("name", mcp.Description("New name.")),
		mcp.WithString("value", mcp.Description("New value.")),
		mcp.WithString("description", mcp.Description("New description.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("tag_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		tag := map[string]interface{}{"id": id}
		if v, err := req.RequireString("name"); err == nil && v != "" { tag["name"] = v }
		if v, err := req.RequireString("value"); err == nil && v != "" { tag["value"] = v }
		if v, err := req.RequireString("description"); err == nil && v != "" { tag["description"] = v }
		data, err := client.V6("PUT", fmt.Sprintf("/flow_tag/v202404alpha1/tag/%s", id), map[string]interface{}{"flowTag": tag})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Update tag failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_delete_tag",
		mcp.WithDescription("Delete a flow tag by ID."),
		mcp.WithString("tag_id", mcp.Required(), mcp.Description("Tag ID to delete.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("tag_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("DELETE", fmt.Sprintf("/flow_tag/v202404alpha1/tag/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Delete tag failed: %v", err)), nil }
		if len(data) == 0 { return mcp.NewToolResultText(fmt.Sprintf("Tag %s deleted.", id)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
