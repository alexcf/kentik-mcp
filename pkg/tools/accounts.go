package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/awlx/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerAccountTools(s *server.MCPServer, client *kentik.Client) {
	listAccounts := mcp.NewTool("kentik_list_accounts",
		mcp.WithDescription("List all configured Kentik account profiles and show the currently active one. Profiles are loaded from ~/.kentik-mcp.json."),
	)
	s.AddTool(listAccounts, makeListAccountsHandler(client))

	switchAccount := mcp.NewTool("kentik_switch_account",
		mcp.WithDescription("Switch to a different Kentik account profile. The switch takes effect immediately for all subsequent API calls. Use kentik_list_accounts to see available profiles."),
		mcp.WithString("profile_name",
			mcp.Required(),
			mcp.Description("Name of the profile to switch to (as defined in ~/.kentik-mcp.json)."),
		),
	)
	s.AddTool(switchAccount, makeSwitchAccountHandler(client))
}

func makeListAccountsHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		profiles := client.ListProfiles()
		current := client.CurrentProfile()

		if len(profiles) == 0 {
			return mcp.NewToolResultText(fmt.Sprintf(
				"No named profiles configured.\n\nActive account: %s (from environment variables)\n\n"+
					"To configure profiles, create ~/.kentik-mcp.json:\n\n```json\n"+
					"{\n  \"profiles\": [\n    {\n      \"name\": \"prod-us\",\n      \"email\": \"you@company.com\",\n"+
					"      \"api_token\": \"your_token\",\n      \"region\": \"US\"\n    }\n  ]\n}\n```",
				client.CurrentEmail(),
			)), nil
		}

		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("## Kentik Account Profiles (%d configured)\n\n", len(profiles)))
		sb.WriteString(fmt.Sprintf("| %-20s | %-35s | %-6s | %-8s |\n", "Profile", "Email", "Region", "Active"))
		sb.WriteString("|" + strings.Repeat("-", 22) + "|" + strings.Repeat("-", 37) +
			"|" + strings.Repeat("-", 8) + "|" + strings.Repeat("-", 10) + "|\n")

		for _, p := range profiles {
			region := p.Region
			if region == "" {
				region = "US"
			}
			active := ""
			if strings.EqualFold(p.Name, current) {
				active = "YES"
			}
			sb.WriteString(fmt.Sprintf("| %-20s | %-35s | %-6s | %-8s |\n",
				p.Name, p.Email, region, active))
		}

		sb.WriteString(fmt.Sprintf("\n*Active profile: **%s** (%s)*\n", current, client.CurrentEmail()))
		return mcp.NewToolResultText(sb.String()), nil
	}
}

func makeSwitchAccountHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		profileName, err := request.RequireString("profile_name")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		if err := client.SwitchProfile(profileName); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(fmt.Sprintf(
			"Switched to profile **%s** (%s). All subsequent Kentik API calls will use this account.",
			client.CurrentProfile(), client.CurrentEmail(),
		)), nil
	}
}
