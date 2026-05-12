package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/alexcf/kentik-mcp/pkg/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	email := os.Getenv("KENTIK_EMAIL")
	apiToken := os.Getenv("KENTIK_API_TOKEN")
	region := os.Getenv("KENTIK_REGION")
	profileName := os.Getenv("KENTIK_PROFILE")

	// Load profiles from config file (always loaded for runtime switching).
	profileCfg, err := kentik.LoadProfileConfig(kentik.DefaultConfigPath())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not load profile config: %v\n", err)
		profileCfg = &kentik.ProfileConfig{}
	}

	// Credential priority:
	//   1. KENTIK_PROFILE env var (selects a named profile from config file)
	//   2. KENTIK_EMAIL + KENTIK_API_TOKEN env vars (direct credentials)
	//   3. First profile in config file
	activeProfileName := "default"

	if profileName != "" {
		// Find named profile in config file.
		var found *kentik.ProfileEntry
		for i, p := range profileCfg.Profiles {
			if strings.EqualFold(p.Name, profileName) {
				found = &profileCfg.Profiles[i]
				break
			}
		}
		if found == nil {
			fmt.Fprintf(os.Stderr, "Error: profile %q not found in %s\n", profileName, kentik.DefaultConfigPath())
			if len(profileCfg.Profiles) > 0 {
				fmt.Fprintln(os.Stderr, "Available profiles:")
				for _, p := range profileCfg.Profiles {
					fmt.Fprintf(os.Stderr, "  %s (%s, %s)\n", p.Name, p.Email, p.Region)
				}
			} else {
				fmt.Fprintln(os.Stderr, "No profiles found in config file.")
			}
			os.Exit(1)
		}
		email = found.Email
		apiToken = found.APIToken
		region = found.Region
		activeProfileName = found.Name
	} else if email == "" || apiToken == "" {
		// Fall back to first profile in config file.
		if len(profileCfg.Profiles) == 0 {
			fmt.Fprintln(os.Stderr, "Error: no credentials found.")
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprintln(os.Stderr, "Option 1 — environment variables:")
			fmt.Fprintln(os.Stderr, "  export KENTIK_EMAIL=user@example.com")
			fmt.Fprintln(os.Stderr, "  export KENTIK_API_TOKEN=your_api_token")
			fmt.Fprintln(os.Stderr, "  export KENTIK_REGION=US  # optional, US or EU")
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprintln(os.Stderr, "Option 2 — profile config file (~/.kentik-mcp.json):")
			fmt.Fprintln(os.Stderr, `  {"profiles":[{"name":"prod","email":"user@example.com","api_token":"xxx","region":"US"}]}`)
			fmt.Fprintln(os.Stderr, "  export KENTIK_PROFILE=prod  # optional, uses first profile if unset")
			os.Exit(1)
		}
		first := profileCfg.Profiles[0]
		email = first.Email
		apiToken = first.APIToken
		region = first.Region
		activeProfileName = first.Name
	}

	client := kentik.NewClient(kentik.Config{
		Email:    email,
		APIToken: apiToken,
		Region:   region,
	})

	// Register all profiles so the switch-account tool can use them at runtime.
	if len(profileCfg.Profiles) > 0 {
		profiles := make([]kentik.Profile, 0, len(profileCfg.Profiles))
		for _, p := range profileCfg.Profiles {
			profiles = append(profiles, kentik.Profile{
				Name:     p.Name,
				Email:    p.Email,
				APIToken: p.APIToken,
				Region:   p.Region,
			})
		}
		client.AddProfiles(profiles)
	}
	// Set the active profile name so kentik_list_accounts marks it correctly.
	if activeProfileName != "default" {
		_ = client.SwitchProfile(activeProfileName)
	}

	s := server.NewMCPServer(
		"Kentik MCP Server",
		"1.0.0",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
		server.WithInstructions("Kentik MCP Server provides access to the Kentik network observability platform. "+
			"Available capabilities: query network flow data (traffic by source/dest IP, AS, geography, protocol, etc.), "+
			"list and inspect devices, interfaces, sites, labels, tags, and users, "+
			"run and inspect synthetic monitoring tests, agents, and results, "+
			"ask Kentik's AI Advisor natural language questions about your network, "+
			"save and reuse named query contexts, and switch between multiple Kentik accounts. "+
			"API docs: https://kb.kentik.com/docs/apis-overview"),
	)

	tools.RegisterAll(s, client)

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
