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
		mcp.WithDescription("Create a new flow tag rule in Kentik. Flow tags classify traffic matching the rule and assign the tag value."),
		mcp.WithString("flow_tag", mcp.Required(), mcp.Description("Tag value to assign to matching traffic e.g. 'internal', 'cdn', 'voip'.")),
		mcp.WithString("device_name", mcp.Description("Device name(s) to apply tag to. Comma-separated or 'all'.")),
		mcp.WithString("addr", mcp.Description("Source/dest IP or CIDR to match e.g. '10.0.0.0/8'.")),
		mcp.WithString("port", mcp.Description("Port(s) to match e.g. '443' or '80,443'.")),
		mcp.WithString("protocol", mcp.Description("Protocol to match e.g. 'tcp', 'udp'.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		flowTag, err := req.RequireString("flow_tag")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		tag := map[string]interface{}{"flow_tag": flowTag}
		if v, err := req.RequireString("device_name"); err == nil && v != "" { tag["device_name"] = v }
		if v, err := req.RequireString("addr"); err == nil && v != "" { tag["addr"] = v }
		if v, err := req.RequireString("port"); err == nil && v != "" { tag["port"] = v }
		if v, err := req.RequireString("protocol"); err == nil && v != "" { tag["protocol"] = v }
		data, err := client.V5("POST", "/tag", map[string]interface{}{"tag": tag})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Create tag failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_update_tag",
		mcp.WithDescription("Update an existing flow tag rule."),
		mcp.WithString("tag_id", mcp.Required(), mcp.Description("Tag ID to update.")),
		mcp.WithString("flow_tag", mcp.Description("New tag value.")),
		mcp.WithString("device_name", mcp.Description("New device name.")),
		mcp.WithString("addr", mcp.Description("New IP/CIDR to match.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("tag_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		tag := map[string]interface{}{}
		if v, err := req.RequireString("flow_tag"); err == nil && v != "" { tag["flow_tag"] = v }
		if v, err := req.RequireString("device_name"); err == nil && v != "" { tag["device_name"] = v }
		if v, err := req.RequireString("addr"); err == nil && v != "" { tag["addr"] = v }
		if len(tag) == 0 { return mcp.NewToolResultError("No fields to update."), nil }
		data, err := client.V5("PUT", fmt.Sprintf("/tag/%s", id), map[string]interface{}{"tag": tag})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Update tag failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_delete_tag",
		mcp.WithDescription("Delete a flow tag rule by ID."),
		mcp.WithString("tag_id", mcp.Required(), mcp.Description("Tag ID to delete.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("tag_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V5("DELETE", fmt.Sprintf("/tag/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Delete tag failed: %v", err)), nil }
		if len(data) == 0 { return mcp.NewToolResultText(fmt.Sprintf("Tag %s deleted.", id)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
