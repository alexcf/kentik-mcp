package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerCloudExportTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_list_cloud_exports",
		mcp.WithDescription("List all cloud flow export configurations (AWS VPC Flow Logs, Azure NSG, GCP VPC, etc.)."),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V6("GET", "/cloud_export/v202210/exports", nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("List cloud exports failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_cloud_export",
		mcp.WithDescription("Get a cloud export configuration by ID."),
		mcp.WithString("export_id", mcp.Required(), mcp.Description("Cloud export ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("export_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("GET", fmt.Sprintf("/cloud_export/v202210/exports/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get cloud export failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_create_cloud_export",
		mcp.WithDescription("Create a new cloud flow export configuration."),
		mcp.WithString("name", mcp.Required(), mcp.Description("Export name.")),
		mcp.WithString("cloud_provider", mcp.Required(), mcp.Description("Cloud provider: 'aws', 'azure', 'gce', 'ibm'.")),
		mcp.WithString("description", mcp.Description("Optional description.")),
		mcp.WithNumber("plan_id", mcp.Description("Kentik plan ID to associate.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, err := req.RequireString("name")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		provider, err := req.RequireString("cloud_provider")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		export := map[string]interface{}{
			"name":          name,
			"cloudProvider": strings.ToUpper(provider),
			"enabled":       true,
		}
		if v, err := req.RequireString("description"); err == nil && v != "" { export["description"] = v }
		if v, err := req.RequireFloat("plan_id"); err == nil && v > 0 { export["planId"] = int(v) }
		data, err := client.V6("POST", "/cloud_export/v202210/exports", map[string]interface{}{"export": export})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Create cloud export failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_update_cloud_export",
		mcp.WithDescription("Update a cloud export configuration."),
		mcp.WithString("export_id", mcp.Required(), mcp.Description("Export ID.")),
		mcp.WithString("name", mcp.Description("New name.")),
		mcp.WithString("description", mcp.Description("New description.")),
		mcp.WithBoolean("enabled", mcp.Description("Enable or disable the export.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("export_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		export := map[string]interface{}{"id": id}
		if v, err := req.RequireString("name"); err == nil && v != "" { export["name"] = v }
		if v, err := req.RequireString("description"); err == nil && v != "" { export["description"] = v }
		if v, err := req.RequireBool("enabled"); err == nil { export["enabled"] = v }
		data, err := client.V6("PUT", fmt.Sprintf("/cloud_export/v202210/exports/%s", id), map[string]interface{}{"export": export})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Update cloud export failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_delete_cloud_export",
		mcp.WithDescription("Delete a cloud export configuration."),
		mcp.WithString("export_id", mcp.Required(), mcp.Description("Export ID to delete.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("export_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("DELETE", fmt.Sprintf("/cloud_export/v202210/exports/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Delete cloud export failed: %v", err)), nil }
		if len(data) == 0 { return mcp.NewToolResultText(fmt.Sprintf("Cloud export %s deleted.", id)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
