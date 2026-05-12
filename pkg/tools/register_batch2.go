package tools

import (
	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterBatch2 registers alerting policy CRUD and notification channel tools.
func RegisterBatch2(s *server.MCPServer, client *kentik.Client) {
	registerWriteAlertingTools(s, client)
	registerNotificationTools(s, client)
}
