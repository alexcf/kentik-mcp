package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerBGPTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_list_bgp_monitors",
		mcp.WithDescription("List all BGP monitors in the account."),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V6("GET", "/bgp_monitoring/v202210/monitors", nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("List BGP monitors failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_bgp_monitor",
		mcp.WithDescription("Get a BGP monitor by ID."),
		mcp.WithString("monitor_id", mcp.Required(), mcp.Description("Monitor ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("monitor_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("GET", fmt.Sprintf("/bgp_monitoring/v202210/monitors/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get BGP monitor failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_create_bgp_monitor",
		mcp.WithDescription("Create a new BGP monitor for prefix visibility and route change detection."),
		mcp.WithString("monitor_name", mcp.Required(), mcp.Description("Monitor name.")),
		mcp.WithString("prefixes", mcp.Required(), mcp.Description("Comma-separated CIDR prefixes to monitor e.g. '1.2.3.0/24,5.6.0.0/16'.")),
		mcp.WithString("description", mcp.Description("Optional description.")),
		mcp.WithString("notify_channel_ids", mcp.Description("Comma-separated notification channel IDs.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, err := req.RequireString("monitor_name")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		prefixStr, err := req.RequireString("prefixes")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		var targets []map[string]interface{}
		for _, p := range splitAndTrim(prefixStr) {
			afi := "AFI_IP4"
			if len(p) > 0 && strings.Contains(p, ":") {
				afi = "AFI_IP6"
			}
			targets = append(targets, map[string]interface{}{"prefix": p, "afi": afi, "safi": "SAFI_UNICAST"})
		}
		monitor := map[string]interface{}{
			"name":   name,
			"status": "BGP_MONITOR_STATUS_ACTIVE",
			"settings": map[string]interface{}{
				"targets":         targets,
				"health_settings": map[string]interface{}{},
			},
		}
		if v, err := req.RequireString("description"); err == nil && v != "" { monitor["description"] = v }
		if v, err := req.RequireString("notify_channel_ids"); err == nil && v != "" {
			monitor["notificationChannels"] = splitAndTrim(v)
		}
		data, err := client.V6("POST", "/bgp_monitoring/v202210/monitors", map[string]interface{}{"monitor": monitor})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Create BGP monitor failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_update_bgp_monitor",
		mcp.WithDescription("Update a BGP monitor."),
		mcp.WithString("monitor_id", mcp.Required(), mcp.Description("Monitor ID.")),
		mcp.WithString("monitor_name", mcp.Description("New name.")),
		mcp.WithString("description", mcp.Description("New description.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("monitor_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		monitor := map[string]interface{}{"id": id}
		if v, err := req.RequireString("monitor_name"); err == nil && v != "" { monitor["name"] = v }
		if v, err := req.RequireString("description"); err == nil && v != "" { monitor["description"] = v }
		data, err := client.V6("PUT", fmt.Sprintf("/bgp_monitoring/v202210/monitors/%s", id), map[string]interface{}{"monitor": monitor})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Update BGP monitor failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_delete_bgp_monitor",
		mcp.WithDescription("Delete a BGP monitor by ID."),
		mcp.WithString("monitor_id", mcp.Required(), mcp.Description("Monitor ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("monitor_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("DELETE", fmt.Sprintf("/bgp_monitoring/v202210/monitors/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Delete BGP monitor failed: %v", err)), nil }
		if len(data) == 0 { return mcp.NewToolResultText(fmt.Sprintf("BGP monitor %s deleted.", id)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_set_bgp_monitor_status",
		mcp.WithDescription("Pause or resume a BGP monitor."),
		mcp.WithString("monitor_id", mcp.Required(), mcp.Description("Monitor ID.")),
		mcp.WithString("status", mcp.Required(), mcp.Description("'active' or 'paused'.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("monitor_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		status, err := req.RequireString("status")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		apiStatus := "BGP_MONITOR_STATUS_ACTIVE"
		if strings.ToLower(status) == "paused" { apiStatus = "BGP_MONITOR_STATUS_PAUSED" }
		data, err := client.V6("PUT", fmt.Sprintf("/bgp_monitoring/v202210/monitors/%s/status", id), map[string]interface{}{"id": id, "status": apiStatus})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Set BGP monitor status failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_bgp_metrics",
		mcp.WithDescription("Get BGP metric data for a prefix over a time interval."),
		mcp.WithString("monitor_id", mcp.Required(), mcp.Description("Monitor ID.")),
		mcp.WithString("prefix", mcp.Required(), mcp.Description("CIDR prefix e.g. '1.2.3.0/24'.")),
		mcp.WithString("start_time", mcp.Required(), mcp.Description("Start time in RFC3339 format.")),
		mcp.WithString("end_time", mcp.Required(), mcp.Description("End time in RFC3339 format.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, _ := req.RequireString("monitor_id")
		prefix, _ := req.RequireString("prefix")
		start, _ := req.RequireString("start_time")
		end, _ := req.RequireString("end_time")
		body := map[string]interface{}{"id": id, "prefix": prefix, "startTime": start, "endTime": end}
		data, err := client.V6("POST", "/bgp_monitoring/v202210/metrics", body)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get BGP metrics failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_bgp_routes",
		mcp.WithDescription("Get a snapshot of BGP route information for a prefix."),
		mcp.WithString("monitor_id", mcp.Required(), mcp.Description("Monitor ID.")),
		mcp.WithString("prefix", mcp.Required(), mcp.Description("CIDR prefix.")),
		mcp.WithString("start_time", mcp.Required(), mcp.Description("Time in RFC3339 format.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, _ := req.RequireString("monitor_id")
		prefix, _ := req.RequireString("prefix")
		start, _ := req.RequireString("start_time")
		data, err := client.V6("POST", "/bgp_monitoring/v202210/routes", map[string]interface{}{"id": id, "prefix": prefix, "startTime": start})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get BGP routes failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
