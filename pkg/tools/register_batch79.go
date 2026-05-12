package tools

import (
	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterBatch79 registers Pathfinder, KMI, and Credentials Vault tools.
func RegisterBatch79(s *server.MCPServer, client *kentik.Client) {
	registerPathfinderTools(s, client)
	registerKMITools(s, client)
	registerCredentialTools(s, client)
}
