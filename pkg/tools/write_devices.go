package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerWriteDeviceTools(s *server.MCPServer, client *kentik.Client) {
	// Create device
	s.AddTool(mcp.NewTool("kentik_create_device",
		mcp.WithDescription("Create a new device in Kentik."),
		mcp.WithString("device_name", mcp.Required(), mcp.Description("Unique device name.")),
		mcp.WithString("device_type", mcp.Required(), mcp.Description("Device type: 'router', 'host', 'aws_subnet', 'azure_subnet', 'gce_subnet'.")),
		mcp.WithString("device_subtype", mcp.Description("Device subtype, e.g. 'router', 'advanced_sflow'. Defaults to device_type.")),
		mcp.WithString("sending_ip", mcp.Description("IP address that sends flow to Kentik.")),
		mcp.WithString("device_description", mcp.Description("Optional description.")),
		mcp.WithNumber("device_sample_rate", mcp.Description("Flow sample rate. Default: 1.")),
		mcp.WithNumber("plan_id", mcp.Description("Plan ID to assign the device to.")),
		mcp.WithNumber("site_id", mcp.Description("Site ID to assign the device to.")),
		mcp.WithString("device_snmp_ip", mcp.Description("SNMP polling IP address.")),
	), makeCreateDeviceHandler(client))

	// Update device
	s.AddTool(mcp.NewTool("kentik_update_device",
		mcp.WithDescription("Update an existing Kentik device by ID."),
		mcp.WithString("device_id", mcp.Required(), mcp.Description("Device ID to update.")),
		mcp.WithString("device_name", mcp.Description("New device name.")),
		mcp.WithString("device_description", mcp.Description("New description.")),
		mcp.WithString("device_snmp_ip", mcp.Description("New SNMP IP.")),
		mcp.WithNumber("device_sample_rate", mcp.Description("New sample rate.")),
		mcp.WithNumber("site_id", mcp.Description("New site ID.")),
	), makeUpdateDeviceHandler(client))

	// Delete device
	s.AddTool(mcp.NewTool("kentik_delete_device",
		mcp.WithDescription("Delete a device from Kentik by ID. This is irreversible."),
		mcp.WithString("device_id", mcp.Required(), mcp.Description("Device ID to delete.")),
	), makeDeleteDeviceHandler(client))

	// Assign labels to device
	s.AddTool(mcp.NewTool("kentik_set_device_labels",
		mcp.WithDescription("Set (replace) all labels on a device. Removes existing labels and applies the provided list."),
		mcp.WithString("device_id", mcp.Required(), mcp.Description("Device ID.")),
		mcp.WithString("label_ids", mcp.Required(), mcp.Description("Comma-separated label IDs to assign.")),
	), makeSetDeviceLabelsHandler(client))
}

func makeCreateDeviceHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, err := request.RequireString("device_name")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		dtype, err := request.RequireString("device_type")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		device := map[string]interface{}{
			"device_name": name,
			"device_type": dtype,
		}
		if v, err := request.RequireString("device_subtype"); err == nil && v != "" {
			device["device_subtype"] = v
		} else {
			device["device_subtype"] = dtype
		}
		if v, err := request.RequireString("sending_ip"); err == nil && v != "" {
			device["sending_ips"] = []string{v}
		}
		if v, err := request.RequireString("device_description"); err == nil && v != "" {
			device["device_description"] = v
		}
		if v, err := request.RequireFloat("device_sample_rate"); err == nil && v > 0 {
			device["device_sample_rate"] = fmt.Sprintf("%d", int(v))
		} else {
			device["device_sample_rate"] = "1"
		}
		if v, err := request.RequireFloat("plan_id"); err == nil && v > 0 {
			device["plan_id"] = int(v)
		}
		if v, err := request.RequireFloat("site_id"); err == nil && v > 0 {
			device["site"] = map[string]interface{}{"id": int(v)}
		}
		if v, err := request.RequireString("device_snmp_ip"); err == nil && v != "" {
			device["device_snmp_ip"] = v
		}

		data, err := client.V5("POST", "/device", map[string]interface{}{"device": device})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Create device failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeUpdateDeviceHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("device_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		device := map[string]interface{}{}
		if v, err := request.RequireString("device_name"); err == nil && v != "" {
			device["device_name"] = v
		}
		if v, err := request.RequireString("device_description"); err == nil && v != "" {
			device["device_description"] = v
		}
		if v, err := request.RequireString("device_snmp_ip"); err == nil && v != "" {
			device["device_snmp_ip"] = v
		}
		if v, err := request.RequireFloat("device_sample_rate"); err == nil && v > 0 {
			device["device_sample_rate"] = fmt.Sprintf("%d", int(v))
		}
		if v, err := request.RequireFloat("site_id"); err == nil && v > 0 {
			device["site"] = map[string]interface{}{"id": int(v)}
		}
		if len(device) == 0 {
			return mcp.NewToolResultError("No fields to update provided."), nil
		}

		data, err := client.V5("PUT", fmt.Sprintf("/device/%s", id), map[string]interface{}{"device": device})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Update device failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeDeleteDeviceHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("device_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		data, err := client.V5("DELETE", fmt.Sprintf("/device/%s", id), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Delete device failed: %v", err)), nil
		}
		if len(data) == 0 {
			return mcp.NewToolResultText(fmt.Sprintf("Device %s deleted successfully.", id)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeSetDeviceLabelsHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("device_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		labelIDsStr, err := request.RequireString("label_ids")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		var labels []map[string]interface{}
		for _, lid := range splitAndTrim(labelIDsStr) {
			labels = append(labels, map[string]interface{}{"id": lid})
		}
		body := map[string]interface{}{"labels": labels}
		data, err := client.V5("PUT", fmt.Sprintf("/device/%s/labels", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Set labels failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}
