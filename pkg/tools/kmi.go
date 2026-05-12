package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerKMITools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_list_kmi_markets",
		mcp.WithDescription("List available KMI geographic markets for internet routing intelligence."),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, err := client.V6("GET", "/kmi/v202212/markets", nil)
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("List KMI markets failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_kmi_network",
		mcp.WithDescription("Get KMI network data showing customers, providers, or peers for an ASN in a market."),
		mcp.WithString("market_id", mcp.Required(), mcp.Description("KMI market ID.")),
		mcp.WithString("asn", mcp.Required(), mcp.Description("AS number.")),
		mcp.WithString("network_type", mcp.Required(), mcp.Description("Relationship type: 'customers', 'providers', or 'peers'.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		marketID, err := req.RequireString("market_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		asn, err := req.RequireString("asn")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		netType, err := req.RequireString("network_type")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		path := fmt.Sprintf("/kmi/v202212/market/%s/network/%s/%s", marketID, asn, netType)
		data, err := client.V6("POST", path, map[string]interface{}{"limit": 50})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get KMI network failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})

	s.AddTool(mcp.NewTool("kentik_get_kmi_rankings",
		mcp.WithDescription("Get KMI ASN rankings in a market by customer base, transit degree, or peer count."),
		mcp.WithString("market_id", mcp.Required(), mcp.Description("KMI market ID.")),
		mcp.WithString("rank_type", mcp.Required(), mcp.Description("Ranking type: 'customer_base', 'transit_degree', or 'peer_count'.")),
		mcp.WithNumber("limit", mcp.Description("Number of results. Default: 20.")),
	), func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		marketID, err := req.RequireString("market_id")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		rankType, err := req.RequireString("rank_type")
		if err != nil { return mcp.NewToolResultError(err.Error()), nil }
		limit := 20.0
		if v, err := req.RequireFloat("limit"); err == nil && v > 0 { limit = v }
		path := fmt.Sprintf("/kmi/v202212/market/%s/rankings/%s", marketID, rankType)
		data, err := client.V6("POST", path, map[string]interface{}{"limit": int(limit)})
		if err != nil { return mcp.NewToolResultError(fmt.Sprintf("Get KMI rankings failed: %v", err)), nil }
		return mcp.NewToolResultText(formatJSON(data)), nil
	})
}
