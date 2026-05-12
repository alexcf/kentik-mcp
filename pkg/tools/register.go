package tools

import (
	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterAll registers every Kentik tool on the given MCP server.
func RegisterAll(s *server.MCPServer, client *kentik.Client) {
	registerDeviceTools(s, client)
	registerInterfaceTools(s, client)
	registerQueryTools(s, client)
	registerTopTalkersTools(s, client)
	registerMultiSiteTools(s, client)
	registerCapacityPlanTools(s, client)
	registerSNMPTools(s, client)
	registerAlertingTools(s, client)
	registerSyntheticsTools(s, client)
	registerLabelTools(s, client)
	registerSiteTools(s, client)
	registerUserTools(s, client)
	registerTagTools(s, client)
	registerAIAdvisorTools(s, client)
	registerDimensionTools(s)
	registerContextTools(s)
	registerAccountTools(s, client)
	registerDocsTools(s)
	registerWriteDeviceTools(s, client)
	registerWriteLabelTools(s, client)
	registerWriteSiteTools(s, client)
	registerWriteSyntheticsTools(s, client)
	RegisterBatch1(s, client)
	RegisterBatch2(s, client)
	RegisterBatch3(s, client)
	RegisterBatch4(s, client)
	RegisterBatch5(s, client)
	RegisterBatch6(s, client)
	RegisterBatch79(s, client)
}
