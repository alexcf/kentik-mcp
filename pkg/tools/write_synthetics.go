package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerWriteSyntheticsTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_create_synthetic_test",
		mcp.WithDescription("Create a new Kentik synthetic monitoring test."),
		mcp.WithString("name", mcp.Required(), mcp.Description("Test name.")),
		mcp.WithString("type", mcp.Required(), mcp.Description("Test type: 'ip', 'hostname', 'agent', 'flow', 'url', 'page_load', 'dns', 'dns_grid', 'bgp_monitor', 'network_grid', 'throughput'.")),
		mcp.WithString("target", mcp.Description("Target hostname, IP, or URL depending on test type.")),
		mcp.WithString("agent_ids", mcp.Description("Comma-separated agent IDs to run the test from.")),
		mcp.WithNumber("period", mcp.Description("Test period in seconds. Default: 60.")),
		mcp.WithString("protocol", mcp.Description("Protocol for network tests: 'icmp', 'tcp', 'udp'.")),
		mcp.WithNumber("port", mcp.Description("Port for TCP/UDP tests.")),
	), makeCreateSyntheticTestHandler(client))

	s.AddTool(mcp.NewTool("kentik_delete_synthetic_test",
		mcp.WithDescription("Delete a synthetic test by ID."),
		mcp.WithString("test_id", mcp.Required(), mcp.Description("Test ID to delete.")),
	), makeDeleteSyntheticTestHandler(client))

	s.AddTool(mcp.NewTool("kentik_set_synthetic_test_status",
		mcp.WithDescription("Pause or resume a synthetic test."),
		mcp.WithString("test_id", mcp.Required(), mcp.Description("Test ID.")),
		mcp.WithString("status", mcp.Required(), mcp.Description("New status: 'active' or 'paused'.")),
	), makeSetSyntheticTestStatusHandler(client))
}

func makeCreateSyntheticTestHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, err := request.RequireString("name")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		testType, err := request.RequireString("type")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		period := 60.0
		if v, err := request.RequireFloat("period"); err == nil && v > 0 {
			period = v
		}

		settings := map[string]interface{}{
			"period": int(period),
		}

		// Set target based on type
		if target, err := request.RequireString("target"); err == nil && target != "" {
			switch testType {
			case "hostname":
				settings["hostname"] = map[string]interface{}{"target": target}
			case "ip":
				settings["ip"] = map[string]interface{}{"targets": []string{target}}
			case "url", "page_load":
				settings["http"] = map[string]interface{}{"target": target}
			case "dns", "dns_grid":
				settings["dns"] = map[string]interface{}{"target": target}
			default:
				settings["hostname"] = map[string]interface{}{"target": target}
			}
		}

		if proto, err := request.RequireString("protocol"); err == nil && proto != "" {
			networkSettings := map[string]interface{}{"protocol": strings.ToUpper(proto)}
			if port, err := request.RequireFloat("port"); err == nil && port > 0 {
				networkSettings["port"] = int(port)
			}
			settings["network_grid"] = networkSettings
		}

		var agentIDs []string
		if ids, err := request.RequireString("agent_ids"); err == nil && ids != "" {
			agentIDs = splitAndTrim(ids)
		}

		test := map[string]interface{}{
			"name":     name,
			"type":     testType,
			"status":   "TEST_STATUS_ACTIVE",
			"settings": settings,
			"deviceId": agentIDs,
		}

		data, err := client.V6("POST", "/synthetics/v202309/tests", map[string]interface{}{"test": test})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Create synthetic test failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeDeleteSyntheticTestHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("test_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		data, err := client.V6("DELETE", fmt.Sprintf("/synthetics/v202309/tests/%s", id), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Delete test failed: %v", err)), nil
		}
		if len(data) == 0 {
			return mcp.NewToolResultText(fmt.Sprintf("Synthetic test %s deleted successfully.", id)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeSetSyntheticTestStatusHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("test_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		status, err := request.RequireString("status")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		// Map friendly status to API enum
		apiStatus := "TEST_STATUS_ACTIVE"
		if strings.ToLower(status) == "paused" {
			apiStatus = "TEST_STATUS_PAUSED"
		}
		body := map[string]interface{}{"id": id, "status": apiStatus}
		data, err := client.V6("POST", fmt.Sprintf("/synthetics/v202309/tests/%s/status", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Set test status failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}
