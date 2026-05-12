package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerWriteSiteTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_create_site",
		mcp.WithDescription("Create a new site in Kentik."),
		mcp.WithString("site_name", mcp.Required(), mcp.Description("Site name.")),
		mcp.WithNumber("lat", mcp.Description("Latitude.")),
		mcp.WithNumber("lon", mcp.Description("Longitude.")),
	), makeCreateSiteHandler(client))

	s.AddTool(mcp.NewTool("kentik_update_site",
		mcp.WithDescription("Update an existing site."),
		mcp.WithString("site_id", mcp.Required(), mcp.Description("Site ID to update.")),
		mcp.WithString("site_name", mcp.Description("New site name.")),
		mcp.WithNumber("lat", mcp.Description("New latitude.")),
		mcp.WithNumber("lon", mcp.Description("New longitude.")),
	), makeUpdateSiteHandler(client))

	s.AddTool(mcp.NewTool("kentik_delete_site",
		mcp.WithDescription("Delete a site by ID."),
		mcp.WithString("site_id", mcp.Required(), mcp.Description("Site ID to delete.")),
	), makeDeleteSiteHandler(client))
}

func makeCreateSiteHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, err := request.RequireString("site_name")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		site := map[string]interface{}{"site_name": name}
		if v, err := request.RequireFloat("lat"); err == nil {
			site["lat"] = v
		}
		if v, err := request.RequireFloat("lon"); err == nil {
			site["lon"] = v
		}
		data, err := client.V5("POST", "/site", map[string]interface{}{"site": site})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Create site failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeUpdateSiteHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("site_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		site := map[string]interface{}{}
		if v, err := request.RequireString("site_name"); err == nil && v != "" {
			site["site_name"] = v
		}
		if v, err := request.RequireFloat("lat"); err == nil {
			site["lat"] = v
		}
		if v, err := request.RequireFloat("lon"); err == nil {
			site["lon"] = v
		}
		if len(site) == 0 {
			return mcp.NewToolResultError("No fields to update provided."), nil
		}
		data, err := client.V5("PUT", fmt.Sprintf("/site/%s", id), map[string]interface{}{"site": site})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Update site failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeDeleteSiteHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("site_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		data, err := client.V5("DELETE", fmt.Sprintf("/site/%s", id), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Delete site failed: %v", err)), nil
		}
		if len(data) == 0 {
			return mcp.NewToolResultText(fmt.Sprintf("Site %s deleted successfully.", id)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}
