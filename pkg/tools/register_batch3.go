package tools

import (
	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterBatch3 registers batch 3 tools.
func RegisterBatch3(s *server.MCPServer, client *kentik.Client) {
	registerBGPTools(s, client)
}
