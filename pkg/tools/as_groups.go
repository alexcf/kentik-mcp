package tools

import (
	"context"
	"fmt"
	"strconv"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerASGroupTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_list_as_groups",
		mcp.WithDescription("List all AS groups. AS groups aggregate multiple ASNs into named sets for traffic analysis."),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V6("GET", "/as_group/v202212/as_group", nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("List AS groups failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_as_group",
		mcp.WithDescription("Get an AS group by ID."),
		mcp.WithString("group_id", mcp.Required(), mcp.Description("AS group ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("group_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("GET", fmt.Sprintf("/as_group/v202212/as_group/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get AS group failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_create_as_group",
		mcp.WithDescription("Create a new AS group to aggregate ASNs for traffic analysis."),
		mcp.WithString("name", mcp.Required(), mcp.Description("Group name e.g. 'hyperscalers'.")),
		mcp.WithString("asns", mcp.Required(), mcp.Description("Comma-separated AS numbers e.g. '15169,8075,16509'.")),
		mcp.WithString("description", mcp.Description("Optional description.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, err := req.RequireString("name")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		asnsStr, err := req.RequireString("asns")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		var asns []map[string]interface{}
		for _, a := range splitAndTrim(asnsStr) {
			if n, err := strconv.Atoi(a); err == nil {
				asns = append(asns, map[string]interface{}{"asn": n})
			}
		}
		group := map[string]interface{}{"name": name, "asns": asns}
		if v, err := req.RequireString("description"); err == nil && v != "" { group["description"] = v }
		data, err := client.V6("POST", "/as_group/v202212/as_group", map[string]interface{}{"asGroup": group})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Create AS group failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_update_as_group",
		mcp.WithDescription("Update an AS group."),
		mcp.WithString("group_id", mcp.Required(), mcp.Description("AS group ID.")),
		mcp.WithString("name", mcp.Description("New name.")),
		mcp.WithString("description", mcp.Description("New description.")),
		mcp.WithString("asns", mcp.Description("New comma-separated AS numbers.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("group_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		group := map[string]interface{}{"id": id}
		if v, err := req.RequireString("name"); err == nil && v != "" { group["name"] = v }
		if v, err := req.RequireString("description"); err == nil && v != "" { group["description"] = v }
		if v, err := req.RequireString("asns"); err == nil && v != "" {
			var asns []map[string]interface{}
			for _, a := range splitAndTrim(v) {
				if n, err := strconv.Atoi(a); err == nil { asns = append(asns, map[string]interface{}{"asn": n}) }
			}
			group["asns"] = asns
		}
		data, err := client.V6("PUT", fmt.Sprintf("/as_group/v202212/as_group/%s", id), map[string]interface{}{"asGroup": group})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Update AS group failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_delete_as_group",
		mcp.WithDescription("Delete an AS group by ID."),
		mcp.WithString("group_id", mcp.Required(), mcp.Description("AS group ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("group_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("DELETE", fmt.Sprintf("/as_group/v202212/as_group/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Delete AS group failed: %v", err)), nil }
		if len(data) == 0 { return mcp.NewToolResultText(fmt.Sprintf("AS group %s deleted.", id)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
