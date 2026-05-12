package tools

import (
	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterBatch1 registers user, tag, and interface write tools.
func RegisterBatch1(s *server.MCPServer, client *kentik.Client) {
	registerWriteUserTools(s, client)
	registerWriteTagTools(s, client)
	registerWriteInterfaceTools(s, client)
}
