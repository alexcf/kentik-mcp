package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerPathfinderTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_create_pathfinder_report",
		mcp.WithDescription("Create a Pathfinder network path analysis report. Traces the network path between two IPs with hop-by-hop analysis."),
		mcp.WithString("src_ip", mcp.Required(), mcp.Description("Source IP address.")),
		mcp.WithString("dst_ip", mcp.Required(), mcp.Description("Destination IP address.")),
		mcp.WithString("start_time", mcp.Required(), mcp.Description("Start time in RFC3339 format e.g. '2024-01-01T00:00:00Z'.")),
		mcp.WithString("end_time", mcp.Required(), mcp.Description("End time in RFC3339 format.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		srcIP, err := req.RequireString("src_ip")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		dstIP, err := req.RequireString("dst_ip")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		start, err := req.RequireString("start_time")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		end, err := req.RequireString("end_time")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		body := map[string]interface{}{
			"src":       map[string]interface{}{"ip": srcIP},
			"dst":       map[string]interface{}{"ip": dstIP},
			"startTime": start,
			"endTime":   end,
		}
		data, err := client.V6("POST", "/pathfinder/v202505beta1/create", body)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Create pathfinder report failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
