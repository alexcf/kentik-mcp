package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerWriteAlertingTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_list_alert_policies",
		mcp.WithDescription("List all alert policies configured in Kentik."),
		mcp.WithString("policy_type", mcp.Description("Filter by type: 'flow', 'nms', 'synthetic'. Leave empty for all.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		body := map[string]interface{}{}
		if v, err := req.RequireString("policy_type"); err == nil && v != "" {
			body["policyType"] = v
		}
		data, err := client.V6("POST", "/v202505/policies/list", body)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("List alert policies failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_alert_policy",
		mcp.WithDescription("Get a specific alert policy by ID."),
		mcp.WithString("policy_id", mcp.Required(), mcp.Description("Policy ID.")),
		mcp.WithString("policy_type", mcp.Required(), mcp.Description("Policy type: 'flow', 'nms', or 'synthetic'.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("policy_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		ptype, err := req.RequireString("policy_type")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("GET", fmt.Sprintf("/v202505/policies/%s/%s", ptype, id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get alert policy failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_disable_alert_policy",
		mcp.WithDescription("Disable an alert policy (stop it from generating alerts)."),
		mcp.WithString("policy_id", mcp.Required(), mcp.Description("Policy ID.")),
		mcp.WithString("policy_type", mcp.Required(), mcp.Description("Policy type: 'flow', 'nms', or 'synthetic'.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("policy_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		ptype, err := req.RequireString("policy_type")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("POST", fmt.Sprintf("/v202505/policies/%s/%s/disable", ptype, id), map[string]interface{}{})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Disable alert policy failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_enable_alert_policy",
		mcp.WithDescription("Enable a previously disabled alert policy."),
		mcp.WithString("policy_id", mcp.Required(), mcp.Description("Policy ID.")),
		mcp.WithString("policy_type", mcp.Required(), mcp.Description("Policy type: 'flow', 'nms', or 'synthetic'.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("policy_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		ptype, err := req.RequireString("policy_type")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("POST", fmt.Sprintf("/v202505/policies/%s/%s/enable", ptype, id), map[string]interface{}{})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Enable alert policy failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_start_manual_mitigation",
		mcp.WithDescription("Start a manual DDoS mitigation for an active alert."),
		mcp.WithString("alert_id", mcp.Required(), mcp.Description("Alert ID to mitigate.")),
		mcp.WithString("ip_cidr", mcp.Required(), mcp.Description("IP/CIDR to mitigate e.g. '1.2.3.4/32'.")),
		mcp.WithString("comment", mcp.Description("Mitigation comment.")),
		mcp.WithNumber("platform_id", mcp.Description("Mitigation platform ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		alertID, err := req.RequireString("alert_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		cidr, err := req.RequireString("ip_cidr")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		body := map[string]interface{}{"id": alertID, "ip_cidr": cidr}
		if v, err := req.RequireString("comment"); err == nil && v != "" { body["comment"] = v }
		if v, err := req.RequireFloat("platform_id"); err == nil && v > 0 { body["platform_id"] = int(v) }
		data, err := client.V5("POST", "/alerts/manual-mitigate", map[string]interface{}{"manual-mitigate": body})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Manual mitigation failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
