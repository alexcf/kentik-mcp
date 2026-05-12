package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerWriteLabelTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_create_label",
		mcp.WithDescription("Create a new device label in Kentik."),
		mcp.WithString("name", mcp.Required(), mcp.Description("Label name.")),
		mcp.WithString("color", mcp.Description("Label color as hex, e.g. '#FF5733'.")),
		mcp.WithString("description", mcp.Description("Optional description.")),
	), makeCreateLabelHandler(client))

	s.AddTool(mcp.NewTool("kentik_update_label",
		mcp.WithDescription("Update an existing device label."),
		mcp.WithString("label_id", mcp.Required(), mcp.Description("Label ID to update.")),
		mcp.WithString("name", mcp.Description("New label name.")),
		mcp.WithString("color", mcp.Description("New color as hex.")),
		mcp.WithString("description", mcp.Description("New description.")),
	), makeUpdateLabelHandler(client))

	s.AddTool(mcp.NewTool("kentik_delete_label",
		mcp.WithDescription("Delete a device label by ID."),
		mcp.WithString("label_id", mcp.Required(), mcp.Description("Label ID to delete.")),
	), makeDeleteLabelHandler(client))
}

func makeCreateLabelHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, err := request.RequireString("name")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		label := map[string]interface{}{"name": name}
		if v, err := request.RequireString("color"); err == nil && v != "" {
			label["color"] = v
		}
		if v, err := request.RequireString("description"); err == nil && v != "" {
			label["description"] = v
		}
		data, err := client.V5("POST", "/deviceLabels", map[string]interface{}{"label": label})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Create label failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeUpdateLabelHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("label_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		label := map[string]interface{}{}
		if v, err := request.RequireString("name"); err == nil && v != "" {
			label["name"] = v
		}
		if v, err := request.RequireString("color"); err == nil && v != "" {
			label["color"] = v
		}
		if v, err := request.RequireString("description"); err == nil && v != "" {
			label["description"] = v
		}
		if len(label) == 0 {
			return mcp.NewToolResultError("No fields to update provided."), nil
		}
		data, err := client.V5("PUT", fmt.Sprintf("/deviceLabels/%s", id), map[string]interface{}{"label": label})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Update label failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeDeleteLabelHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("label_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		data, err := client.V5("DELETE", fmt.Sprintf("/deviceLabels/%s", id), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Delete label failed: %v", err)), nil
		}
		if len(data) == 0 {
			return mcp.NewToolResultText(fmt.Sprintf("Label %s deleted successfully.", id)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}
