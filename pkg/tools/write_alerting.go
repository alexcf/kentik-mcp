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
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V5("GET", "/alerts", nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("List alert policies failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_alert_policy",
		mcp.WithDescription("Get a specific alert policy by ID."),
		mcp.WithString("policy_id", mcp.Required(), mcp.Description("Policy ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("policy_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V5("GET", fmt.Sprintf("/alert/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get alert policy failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_create_alert_policy",
		mcp.WithDescription("Create a new alert policy in Kentik."),
		mcp.WithString("policy_name", mcp.Required(), mcp.Description("Policy name.")),
		mcp.WithString("policy_description", mcp.Description("Policy description.")),
		mcp.WithString("severity", mcp.Description("Alert severity: 'critical', 'major', 'minor', 'warning'.")),
		mcp.WithBoolean("is_active", mcp.Description("Whether policy is active. Default: true.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, err := req.RequireString("policy_name")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		policy := map[string]interface{}{"policy_name": name, "status": "A"}
		if v, err := req.RequireString("policy_description"); err == nil && v != "" { policy["policy_description"] = v }
		if v, err := req.RequireString("severity"); err == nil && v != "" { policy["activate_when"] = v }
		if v, err := req.RequireBool("is_active"); err == nil && !v { policy["status"] = "I" }
		data, err := client.V5("POST", "/alert", map[string]interface{}{"policy": policy})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Create alert policy failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_update_alert_policy",
		mcp.WithDescription("Update an existing alert policy."),
		mcp.WithString("policy_id", mcp.Required(), mcp.Description("Policy ID to update.")),
		mcp.WithString("policy_name", mcp.Description("New policy name.")),
		mcp.WithString("policy_description", mcp.Description("New description.")),
		mcp.WithString("severity", mcp.Description("New severity.")),
		mcp.WithBoolean("is_active", mcp.Description("Set active/inactive.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("policy_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		policy := map[string]interface{}{}
		if v, err := req.RequireString("policy_name"); err == nil && v != "" { policy["policy_name"] = v }
		if v, err := req.RequireString("policy_description"); err == nil && v != "" { policy["policy_description"] = v }
		if v, err := req.RequireString("severity"); err == nil && v != "" { policy["activate_when"] = v }
		if v, err := req.RequireBool("is_active"); err == nil { if v { policy["status"] = "A" } else { policy["status"] = "I" } }
		if len(policy) == 0 { return mcp.NewToolResultError("No fields to update."), nil }
		data, err := client.V5("PUT", fmt.Sprintf("/alert/%s", id), map[string]interface{}{"policy": policy})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Update alert policy failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_delete_alert_policy",
		mcp.WithDescription("Delete an alert policy by ID."),
		mcp.WithString("policy_id", mcp.Required(), mcp.Description("Policy ID to delete.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("policy_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V5("DELETE", fmt.Sprintf("/alert/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Delete alert policy failed: %v", err)), nil }
		if len(data) == 0 { return mcp.NewToolResultText(fmt.Sprintf("Alert policy %s deleted.", id)), nil }
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
