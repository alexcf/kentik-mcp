# Configuration Examples

Source: https://kb.kentik.com/docs/aia-mcp-configuration-examples

---

Once you have your Kentik API credentials, you are ready to wire up your AI agent to the **Kentik AI Advisor MCP Server**.

Choose your preferred environment below to get connected:

## **Claude Desktop**

Claude Desktop supports MCP servers on both free and paid plans through manual JSON configuration.

### **Configuration File Location**

* **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
* **Windows**: `%APPDATA%\\Claude\\claude_desktop_config.json`

**Quick Access**: Claude Desktop menu → Settings → Developer → Edit Config

### **Configuration**

Add the following to your  `claude_desktop_config.json`:

```
{
  "mcpServers": {
    "kentik-ai-advisor": {
      "command": "npx",
      "args": [
        "mcp-remote",
        "https://api.kentik.com/mcp",
        "--header",
        "X-CH-Auth-Email: ${KENTIK_API_EMAIL}",
        "--header",
        "X-CH-Auth-API-Token: ${KENTIK_API_TOKEN}"
      ],
      "env": {
        "KENTIK_API_EMAIL": "your-email-here",
        "KENTIK_API_TOKEN": "your-api-token-here"
      }
    }
  }
}
```

**After Configuration**:

1. Restart Claude Desktop.
2. Open Settings to verify the MCP server is connected.
3. Start asking questions - remember to mention "Kentik" in your prompts.

## **GitHub Copilot / VS Code**

For **VS Code with GitHub Copilot** or other MCP extensions:

```
{
  "servers": {
    "kentik-ai-advisor": {
      "url": "https://api.kentik.com/mcp",
      "type": "http",
      "headers": {
        "X-CH-Auth-Email": "your-email-here",
        "X-CH-Auth-API-Token": "your-api-token-here"
      }
    }
  },
  "inputs": []
}
```

## **Custom Applications**

For custom integrations, use any MCP-compatible client library. The server supports:

* **Protocol**: MCP over HTTP (Streamable HTTP transport)
* **Port**: Standard HTTPS (443)
* **Format**: JSON-RPC 2.0 (see **MCP Payload Reference**)
* **Tools**: See **MCP Tools Reference**

Having an issue? See our **Troubleshooting** guide for connecting to the AI Advisor MCP Server.
