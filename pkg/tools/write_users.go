package tools

import (
	"context"
	"fmt"

	"github.com/alexcf/kentik-mcp/pkg/kentik"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerWriteUserTools(s *server.MCPServer, client *kentik.Client) {
	s.AddTool(mcp.NewTool("kentik_create_user",
		mcp.WithDescription("Create a new user in the Kentik organization."),
		mcp.WithString("email", mcp.Required(), mcp.Description("User email address.")),
		mcp.WithString("username", mcp.Required(), mcp.Description("Username.")),
		mcp.WithString("full_name", mcp.Description("Full display name.")),
		mcp.WithString("role", mcp.Description("Role: 'Member' or 'Administrator'. Default: Member.")),
		mcp.WithString("password", mcp.Description("Initial password.")),
	), makeCreateUserHandler(client))

	s.AddTool(mcp.NewTool("kentik_update_user",
		mcp.WithDescription("Update an existing Kentik user by ID."),
		mcp.WithString("user_id", mcp.Required(), mcp.Description("User ID to update.")),
		mcp.WithString("email", mcp.Description("New email.")),
		mcp.WithString("username", mcp.Description("New username.")),
		mcp.WithString("full_name", mcp.Description("New full name.")),
		mcp.WithString("role", mcp.Description("New role: 'Member' or 'Administrator'.")),
	), makeUpdateUserHandler(client))

	s.AddTool(mcp.NewTool("kentik_delete_user",
		mcp.WithDescription("Delete a Kentik user by ID."),
		mcp.WithString("user_id", mcp.Required(), mcp.Description("User ID to delete.")),
	), makeDeleteUserHandler(client))

	s.AddTool(mcp.NewTool("kentik_reset_user_sessions",
		mcp.WithDescription("Reset all active sessions for a user, forcing re-login."),
		mcp.WithString("user_id", mcp.Required(), mcp.Description("User ID.")),
	), makeResetUserSessionsHandler(client))

	s.AddTool(mcp.NewTool("kentik_reset_user_api_token",
		mcp.WithDescription("Reset the API token for a user, generating a new one."),
		mcp.WithString("user_id", mcp.Required(), mcp.Description("User ID.")),
	), makeResetUserAPITokenHandler(client))
}

func makeCreateUserHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		email, err := request.RequireString("email")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		username, err := request.RequireString("username")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		user := map[string]interface{}{"email": email, "username": username}
		if v, err := request.RequireString("full_name"); err == nil && v != "" {
			user["fullName"] = v
		}
		if v, err := request.RequireString("role"); err == nil && v != "" {
			user["role"] = v
		}
		if v, err := request.RequireString("password"); err == nil && v != "" {
			user["password"] = v
		}
		data, err := client.V6("POST", "/user/v202211/users", map[string]interface{}{"user": user})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Create user failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeUpdateUserHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("user_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		user := map[string]interface{}{"id": id}
		if v, err := request.RequireString("email"); err == nil && v != "" {
			user["email"] = v
		}
		if v, err := request.RequireString("username"); err == nil && v != "" {
			user["username"] = v
		}
		if v, err := request.RequireString("full_name"); err == nil && v != "" {
			user["fullName"] = v
		}
		if v, err := request.RequireString("role"); err == nil && v != "" {
			user["role"] = v
		}
		data, err := client.V6("PUT", fmt.Sprintf("/user/v202211/users/%s", id), map[string]interface{}{"user": user})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Update user failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeDeleteUserHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("user_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		data, err := client.V6("DELETE", fmt.Sprintf("/user/v202211/users/%s", id), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Delete user failed: %v", err)), nil
		}
		if len(data) == 0 {
			return mcp.NewToolResultText(fmt.Sprintf("User %s deleted.", id)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeResetUserSessionsHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("user_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		data, err := client.V6("PUT", fmt.Sprintf("/user/v202211/users/%s/reset_active_sessions", id), map[string]interface{}{})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Reset sessions failed: %v", err)), nil
		}
		if len(data) == 0 {
			return mcp.NewToolResultText(fmt.Sprintf("Sessions reset for user %s.", id)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}

func makeResetUserAPITokenHandler(client *kentik.Client) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("user_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		data, err := client.V6("PUT", fmt.Sprintf("/user/v202211/users/%s/reset_api_token", id), map[string]interface{}{})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Reset API token failed: %v", err)), nil
		}
		return mcp.NewToolResultText(formatJSON(data)), nil
	}
}
