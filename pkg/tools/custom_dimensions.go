package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerCustomDimensionTools(s *server.MCPServer, client *kentik.Client) {
	// List custom dimensions
	s.AddTool(mcp.NewTool("kentik_list_custom_dimensions",
		mcp.WithDescription("List all custom dimensions in Kentik. Custom dimensions let you tag flows with business-relevant metadata."),
	), makeListCustomDimensionsHandler(client))

	// Get custom dimension
	s.AddTool(mcp.NewTool("kentik_get_custom_dimension",
		mcp.WithDescription("Get a custom dimension by ID."),
		mcp.WithString("dimension_id", mcp.Required(), mcp.Description("Custom dimension ID.")),
	), makeGetCustomDimensionHandler(client))

	// Create custom dimension
	s.AddTool(mcp.NewTool("kentik_create_custom_dimension",
		mcp.WithDescription("Create a new custom dimension. The name must start with 'c_'."),
		mcp.WithString("name", mcp.Required(), mcp.Description("Dimension name (must start with 'c_').")),
		mcp.WithString("display_name", mcp.Required(), mcp.Description("Human-readable display name.")),
		mcp.WithString("type", mcp.Required(), mcp.Description("Dimension type: 'string' or 'uint32'.")),
	), makeCreateCustomDimensionHandler(client))

	// Update custom dimension
	s.AddTool(mcp.NewTool("kentik_update_custom_dimension",
		mcp.WithDescription("Update an existing custom dimension by ID."),
		mcp.WithString("dimension_id", mcp.Required(), mcp.Description("Custom dimension ID to update.")),
		mcp.WithString("display_name", mcp.Description("New display name.")),
		mcp.WithString("type", mcp.Description("New type: 'string' or 'uint32'.")),
	), makeUpdateCustomDimensionHandler(client))

	// Delete custom dimension
	s.AddTool(mcp.NewTool("kentik_delete_custom_dimension",
		mcp.WithDescription("Delete a custom dimension by ID. This is irreversible."),
		mcp.WithString("dimension_id", mcp.Required(), mcp.Description("Custom dimension ID to delete.")),
	), makeDeleteCustomDimensionHandler(client))

	// Create populator
	s.AddTool(mcp.NewTool("kentik_create_populator",
		mcp.WithDescription("Create a populator rule that assigns a custom dimension value to matching flows."),
		mcp.WithString("dimension_id", mcp.Required(), mcp.Description("Custom dimension ID to add the populator to.")),
		mcp.WithString("value", mcp.Required(), mcp.Description("The tag value to assign to matching flows.")),
		mcp.WithString("direction", mcp.Description("Flow direction to match: 'src', 'dst', or 'either'.")),
		mcp.WithString("addr", mcp.Description("IP address or CIDR to match.")),
		mcp.WithString("port", mcp.Description("Port number or range to match.")),
		mcp.WithString("protocol", mcp.Description("Protocol to match.")),
		mcp.WithString("tcp_flags", mcp.Description("TCP flags to match.")),
		mcp.WithString("device_type", mcp.Description("Device type to match.")),
		mcp.WithString("site", mcp.Description("Site name to match.")),
		mcp.WithString("interface_name", mcp.Description("Interface name to match.")),
		mcp.WithString("device_name", mcp.Description("Device name to match.")),
		mcp.WithString("bgp_as_path", mcp.Description("BGP AS path to match.")),
		mcp.WithString("bgp_community", mcp.Description("BGP community to match.")),
		mcp.WithString("country", mcp.Description("Country code to match.")),
	), makeCreatePopulatorHandler(client))

	// Delete populator
	s.AddTool(mcp.NewTool("kentik_delete_populator",
		mcp.WithDescription("Delete a custom dimension populator rule."),
		mcp.WithString("dimension_id", mcp.Required(), mcp.Description("Custom dimension ID.")),
		mcp.WithString("populator_id", mcp.Required(), mcp.Description("Populator ID to delete.")),
	), makeDeletePopulatorHandler(client))
}

func makeListCustomDimensionsHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V5("GET", "/customdimensions", nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("List custom dimensions failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeGetCustomDimensionHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("dimension_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		data, err := client.V5("GET", fmt.Sprintf("/customdimension/%s", id), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Get custom dimension failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeCreateCustomDimensionHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, err := request.RequireString("name")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		dimType, err := request.RequireString("type")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		// V6 /v1/customdimension with body: "dimension" — send dimension fields directly
		body := map[string]interface{}{
			"name": name,
			"type": dimType,
		}
		if displayName, err := request.RequireString("display_name"); err == nil && displayName != "" {
			body["description"] = displayName
		}
		data, err := client.V6("POST", "/v1/customdimension", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Create custom dimension failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeUpdateCustomDimensionHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("dimension_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		dim := map[string]interface{}{}
		if v, err := request.RequireString("display_name"); err == nil && v != "" {
			dim["display_name"] = v
		}
		if v, err := request.RequireString("type"); err == nil && v != "" {
			dim["type"] = v
		}
		if len(dim) == 0 {
			return mcp.NewToolResultError("No fields to update provided."), nil
		}

		data, err := client.V5("PUT", fmt.Sprintf("/customdimension/%s", id), map[string]interface{}{"dimension": dim})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Update custom dimension failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeDeleteCustomDimensionHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("dimension_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		_, err = client.V5("DELETE", fmt.Sprintf("/customdimension/%s", id), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Delete custom dimension failed: %v", err)), nil
		}
		return mcp.NewToolResultText(fmt.Sprintf("Custom dimension %s deleted.", id)), nil
	}
}

func makeCreatePopulatorHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		dimensionID, err := request.RequireString("dimension_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		value, err := request.RequireString("value")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		populator := map[string]interface{}{
			"value": value,
		}
		if v, err := request.RequireString("direction"); err == nil && v != "" {
			populator["direction"] = v
		}
		if v, err := request.RequireString("addr"); err == nil && v != "" {
			populator["addr"] = v
		}
		if v, err := request.RequireString("port"); err == nil && v != "" {
			populator["port"] = v
		}
		if v, err := request.RequireString("protocol"); err == nil && v != "" {
			populator["protocol"] = v
		}
		if v, err := request.RequireString("tcp_flags"); err == nil && v != "" {
			populator["tcp_flags"] = v
		}
		if v, err := request.RequireString("device_type"); err == nil && v != "" {
			populator["device_type"] = v
		}
		if v, err := request.RequireString("site"); err == nil && v != "" {
			populator["site"] = v
		}
		if v, err := request.RequireString("interface_name"); err == nil && v != "" {
			populator["interface_name"] = v
		}
		if v, err := request.RequireString("device_name"); err == nil && v != "" {
			populator["device_name"] = v
		}
		if v, err := request.RequireString("bgp_as_path"); err == nil && v != "" {
			populator["bgp_as_path"] = v
		}
		if v, err := request.RequireString("bgp_community"); err == nil && v != "" {
			populator["bgp_community"] = v
		}
		if v, err := request.RequireString("country"); err == nil && v != "" {
			populator["country"] = v
		}

		data, err := client.V5("POST", fmt.Sprintf("/customdimension/%s/populator", dimensionID), map[string]interface{}{"populator": populator})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Create populator failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeDeletePopulatorHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		dimensionID, err := request.RequireString("dimension_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		populatorID, err := request.RequireString("populator_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		_, err = client.V5("DELETE", fmt.Sprintf("/customdimension/%s/populator/%s", dimensionID, populatorID), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Delete populator failed: %v", err)), nil
		}
		return mcp.NewToolResultText(fmt.Sprintf("Populator %s deleted.", populatorID)), nil
	}
}
