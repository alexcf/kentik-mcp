package tools

import (
	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterBatch4 registers batch 4 tools.
func RegisterBatch4(s *server.MCPServer, client *kentik.Client) {
	registerWriteSyntheticsAgentTools(s, client)
}
