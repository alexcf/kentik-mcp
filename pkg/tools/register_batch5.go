package tools

import (
	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterBatch5 registers AS group and network classification tools.
func RegisterBatch5(s *server.MCPServer, client *kentik.Client) {
	registerASGroupTools(s, client)
	registerNetworkClassificationTools(s, client)
}
