package tools

import (
	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterV5Gaps registers custom dimensions, saved filters, plans, and audit log tools.
func RegisterV5Gaps(s *server.MCPServer, client *kentik.Client) {
	registerCustomDimensionTools(s, client)
	registerSavedFilterTools(s, client)
	registerPlanTools(s, client)
}
