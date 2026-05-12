package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerWriteInterfaceTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_create_interface",
		mcp.WithDescription("Create a new interface on a device in Kentik."),
		mcp.WithString("device_id", mcp.Required(), mcp.Description("Device ID.")),
		mcp.WithString("snmp_id", mcp.Required(), mcp.Description("SNMP interface index.")),
		mcp.WithString("interface_description", mcp.Description("Interface description.")),
		mcp.WithString("interface_ip", mcp.Description("Interface IP address.")),
		mcp.WithNumber("interface_speed_mbps", mcp.Description("Interface speed in Mbps.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		deviceID, err := req.RequireString("device_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		snmpID, err := req.RequireString("snmp_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		iface := map[string]interface{}{"deviceId": deviceID, "snmpId": snmpID}
		if v, err := req.RequireString("interface_description"); err == nil && v != "" { iface["interfaceDescription"] = v }
		if v, err := req.RequireString("interface_ip"); err == nil && v != "" { iface["interfaceIp"] = v }
		if v, err := req.RequireFloat("interface_speed_mbps"); err == nil && v > 0 { iface["interfaceSpeedMbps"] = int(v) }
		data, err := client.V6("POST", "/interface/v202108alpha1/interfaces", map[string]interface{}{"interface": iface})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Create interface failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_update_interface",
		mcp.WithDescription("Update an existing interface."),
		mcp.WithString("interface_id", mcp.Required(), mcp.Description("Interface ID to update.")),
		mcp.WithString("interface_description", mcp.Description("New description.")),
		mcp.WithString("interface_ip", mcp.Description("New IP address.")),
		mcp.WithNumber("interface_speed_mbps", mcp.Description("New speed in Mbps.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("interface_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		iface := map[string]interface{}{"id": id}
		if v, err := req.RequireString("interface_description"); err == nil && v != "" { iface["interfaceDescription"] = v }
		if v, err := req.RequireString("interface_ip"); err == nil && v != "" { iface["interfaceIp"] = v }
		if v, err := req.RequireFloat("interface_speed_mbps"); err == nil && v > 0 { iface["interfaceSpeedMbps"] = int(v) }
		data, err := client.V6("PUT", fmt.Sprintf("/interface/v202108alpha1/interfaces/%s", id), map[string]interface{}{"interface": iface})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Update interface failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_delete_interface",
		mcp.WithDescription("Delete an interface by ID."),
		mcp.WithString("interface_id", mcp.Required(), mcp.Description("Interface ID to delete.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("interface_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("DELETE", fmt.Sprintf("/interface/v202108alpha1/interfaces/%s", id), nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Delete interface failed: %v", err)), nil }
		if len(data) == 0 { return mcp.NewToolResultText(fmt.Sprintf("Interface %s deleted.", id)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_classify_interfaces",
		mcp.WithDescription("Manually trigger interface classification for a device."),
		mcp.WithString("device_id", mcp.Required(), mcp.Description("Device ID.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("device_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		data, err := client.V6("POST", "/interface/v202108alpha1/manual_classify", map[string]interface{}{"deviceId": id})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Classify interfaces failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
