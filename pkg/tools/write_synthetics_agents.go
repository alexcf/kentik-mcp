package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerWriteSyntheticsAgentTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_list_agent_alerts",
		mcp.WithDescription("List synthetic agent alert configurations."),
		mcp.WithString("agent_id", mcp.Description("Filter by agent ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		path := "/synthetics/v202309/agentAlerts"
		if v, err := req.RequireString("agent_id"); err == nil && v != "" {
			path += "?agentId=" + v
		}
		data, err := client.V6("GET", path, nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("List agent alerts failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_agent_alert",
		mcp.WithDescription("Get a specific agent alert configuration by ID."),
		mcp.WithString("alert_id", mcp.Required(), mcp.Description("Agent alert ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("alert_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("GET", fmt.Sprintf("/synthetics/v202309/agentAlerts/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get agent alert failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_create_agent_alert",
		mcp.WithDescription("Create a new synthetic agent alert configuration with latency and packet loss thresholds."),
		mcp.WithString("agent_id", mcp.Required(), mcp.Description("Agent ID.")),
		mcp.WithNumber("latency_warning_ms", mcp.Description("Latency warning threshold in ms.")),
		mcp.WithNumber("latency_critical_ms", mcp.Description("Latency critical threshold in ms.")),
		mcp.WithNumber("packet_loss_warning_pct", mcp.Description("Packet loss warning threshold %.")),
		mcp.WithNumber("packet_loss_critical_pct", mcp.Description("Packet loss critical threshold %.")),
		mcp.WithString("notification_channel_ids", mcp.Description("Comma-separated notification channel IDs.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		agentID, err := req.RequireString("agent_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		thresholds := map[string]interface{}{}
		if v, err := req.RequireFloat("latency_warning_ms"); err == nil && v > 0 { thresholds["latencyWarning"] = int(v) }
		if v, err := req.RequireFloat("latency_critical_ms"); err == nil && v > 0 { thresholds["latencyCritical"] = int(v) }
		if v, err := req.RequireFloat("packet_loss_warning_pct"); err == nil && v > 0 { thresholds["packetLossWarning"] = v }
		if v, err := req.RequireFloat("packet_loss_critical_pct"); err == nil && v > 0 { thresholds["packetLossCritical"] = v }
		alert := map[string]interface{}{"agentId": agentID, "thresholds": thresholds}
		if v, err := req.RequireString("notification_channel_ids"); err == nil && v != "" {
			alert["notificationChannels"] = splitAndTrim(v)
		}
		data, err := client.V6("POST", "/synthetics/v202309/agentAlerts", map[string]interface{}{"agentAlert": alert})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Create agent alert failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_update_agent_alert",
		mcp.WithDescription("Update an existing agent alert configuration."),
		mcp.WithString("alert_id", mcp.Required(), mcp.Description("Agent alert ID.")),
		mcp.WithNumber("latency_warning_ms", mcp.Description("New latency warning threshold ms.")),
		mcp.WithNumber("latency_critical_ms", mcp.Description("New latency critical threshold ms.")),
		mcp.WithNumber("packet_loss_warning_pct", mcp.Description("New packet loss warning %.")),
		mcp.WithNumber("packet_loss_critical_pct", mcp.Description("New packet loss critical %.")),
		mcp.WithString("notification_channel_ids", mcp.Description("Comma-separated notification channel IDs.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("alert_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		thresholds := map[string]interface{}{}
		if v, err := req.RequireFloat("latency_warning_ms"); err == nil && v > 0 { thresholds["latencyWarning"] = int(v) }
		if v, err := req.RequireFloat("latency_critical_ms"); err == nil && v > 0 { thresholds["latencyCritical"] = int(v) }
		if v, err := req.RequireFloat("packet_loss_warning_pct"); err == nil && v > 0 { thresholds["packetLossWarning"] = v }
		if v, err := req.RequireFloat("packet_loss_critical_pct"); err == nil && v > 0 { thresholds["packetLossCritical"] = v }
		alert := map[string]interface{}{"id": id, "thresholds": thresholds}
		if v, err := req.RequireString("notification_channel_ids"); err == nil && v != "" {
			alert["notificationChannels"] = splitAndTrim(v)
		}
		data, err := client.V6("PUT", fmt.Sprintf("/synthetics/v202309/agentAlerts/%s", id), map[string]interface{}{"agentAlert": alert})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Update agent alert failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_delete_agent_alert",
		mcp.WithDescription("Delete an agent alert configuration."),
		mcp.WithString("alert_id", mcp.Required(), mcp.Description("Agent alert ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("alert_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("DELETE", fmt.Sprintf("/synthetics/v202309/agentAlerts/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Delete agent alert failed: %v", err)), nil }
		if len(data) == 0 { return mcp.NewToolResultText(fmt.Sprintf("Agent alert %s deleted.", id)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_update_synthetic_agent",
		mcp.WithDescription("Update a synthetic agent's display name."),
		mcp.WithString("agent_id", mcp.Required(), mcp.Description("Agent ID.")),
		mcp.WithString("alias", mcp.Required(), mcp.Description("New display name for the agent.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("agent_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		alias, err := req.RequireString("alias")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("PUT", fmt.Sprintf("/synthetics/v202309/agents/%s", id), map[string]interface{}{"agent": map[string]interface{}{"id": id, "alias": alias}})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Update agent failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_delete_synthetic_agent",
		mcp.WithDescription("Delete a synthetic agent from the account."),
		mcp.WithString("agent_id", mcp.Required(), mcp.Description("Agent ID to delete.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("agent_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("DELETE", fmt.Sprintf("/synthetics/v202309/agents/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Delete agent failed: %v", err)), nil }
		if len(data) == 0 { return mcp.NewToolResultText(fmt.Sprintf("Synthetic agent %s deleted.", id)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
