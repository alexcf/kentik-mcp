package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerSavedFilterTools(s *server.MCPServer, client *kentik.Client) {
	// List saved filters
	s.AddTool(mcp.NewTool("kentik_list_saved_filters",
		mcp.WithDescription("List all custom saved filters."),
	), makeListSavedFiltersHandler(client))

	// Get saved filter
	s.AddTool(mcp.NewTool("kentik_get_saved_filter",
		mcp.WithDescription("Get a custom saved filter by ID."),
		mcp.WithString("filter_id", mcp.Required(), mcp.Description("Saved filter ID.")),
	), makeGetSavedFilterHandler(client))

	// Create saved filter
	s.AddTool(mcp.NewTool("kentik_create_saved_filter",
		mcp.WithDescription("Create a new custom saved filter."),
		mcp.WithString("filter_name", mcp.Required(), mcp.Description("Name for the saved filter.")),
		mcp.WithString("filter_description", mcp.Description("Optional description.")),
		mcp.WithString("filters_json", mcp.Required(), mcp.Description("Raw JSON string of the filter object.")),
	), makeCreateSavedFilterHandler(client))

	// Update saved filter
	s.AddTool(mcp.NewTool("kentik_update_saved_filter",
		mcp.WithDescription("Update an existing custom saved filter by ID."),
		mcp.WithString("filter_id", mcp.Required(), mcp.Description("Saved filter ID to update.")),
		mcp.WithString("filter_name", mcp.Description("New filter name.")),
		mcp.WithString("filter_description", mcp.Description("New description.")),
		mcp.WithString("filters_json", mcp.Description("New filter object as raw JSON string.")),
	), makeUpdateSavedFilterHandler(client))

	// Delete saved filter
	s.AddTool(mcp.NewTool("kentik_delete_saved_filter",
		mcp.WithDescription("Delete a custom saved filter by ID."),
		mcp.WithString("filter_id", mcp.Required(), mcp.Description("Saved filter ID to delete.")),
	), makeDeleteSavedFilterHandler(client))
}

func makeListSavedFiltersHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V5("GET", "/saved-filters/custom", nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("List saved filters failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeGetSavedFilterHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("filter_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		data, err := client.V5("GET", fmt.Sprintf("/saved-filter/saved-filter/custom/%s", id), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Get saved filter failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeCreateSavedFilterHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		filterName, err := request.RequireString("filter_name")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		filtersJSON, err := request.RequireString("filters_json")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		var parsedFilters interface{}
		if err := json.Unmarshal([]byte(filtersJSON), &parsedFilters); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Invalid filters_json: %v", err)), nil
		}

		filterField := map[string]interface{}{
			"filter_name": filterName,
			"filters":     parsedFilters,
		}
		if v, err := request.RequireString("filter_description"); err == nil && v != "" {
			filterField["filter_description"] = v
		}

		data, err := client.V5("POST", "/saved-filter/custom", map[string]interface{}{"filterField": filterField})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Create saved filter failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeUpdateSavedFilterHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("filter_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		filterField := map[string]interface{}{}
		if v, err := request.RequireString("filter_name"); err == nil && v != "" {
			filterField["filter_name"] = v
		}
		if v, err := request.RequireString("filter_description"); err == nil && v != "" {
			filterField["filter_description"] = v
		}
		if v, err := request.RequireString("filters_json"); err == nil && v != "" {
			var parsedFilters interface{}
			if err := json.Unmarshal([]byte(v), &parsedFilters); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Invalid filters_json: %v", err)), nil
			}
			filterField["filters"] = parsedFilters
		}
		if len(filterField) == 0 {
			return mcp.NewToolResultError("No fields to update provided."), nil
		}

		data, err := client.V5("PUT", fmt.Sprintf("/saved-filter/saved-filter/custom/%s", id), map[string]interface{}{"filterField": filterField})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Update saved filter failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeDeleteSavedFilterHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("filter_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		_, err = client.V5("DELETE", fmt.Sprintf("/saved-filter/saved-filter/custom/%s", id), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Delete saved filter failed: %v", err)), nil
		}
		return mcp.NewToolResultText(fmt.Sprintf("Saved filter %s deleted.", id)), nil
	}
}
