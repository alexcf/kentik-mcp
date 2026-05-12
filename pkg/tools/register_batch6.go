package tools

import (
	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterBatch6 registers capacity plan, connectivity cost, and cloud export tools.
func RegisterBatch6(s *server.MCPServer, client *kentik.Client) {
	registerCapacityAPITools(s, client)
	registerCostTools(s, client)
	registerCloudExportTools(s, client)
}
