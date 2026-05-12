# Getting Started

Source: https://kb.kentik.com/docs/aia-mcp-getting-started

---

Follow this guide to set up **Kentik’s AI Advisor MCP Server**.

> 🏃‍♀️ **FAST TRACK:** Already familiar with MCP and have your Kentik API tokenready? Go to **Configuration Examples** and copy the snippet for your preferred agent.

## **Prerequisites**

* An active Kentik account with **AI Advisor enabled**.
* Kentik API credentials (email and **API token**).
* An MCP-compatible AI agent or development tool, e.g., Claude Desktop, GitHub Copilot (see **Configuration Examples**).

## **Get Your API Token**

1. In the Kentik Portal, navigate to your **User Profile** (top-right menu).
2. Select the **Authentication** tab.
3. Copy your **API Token**.

![User profile settings with Single Sign-on information and API token section displayed.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/user-profile-api-token.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A54Z&se=2026-05-12T09%3A27%3A54Z&sr=c&sp=r&sig=lKYpWYQlPKOcLLbF%2BG2cA%2FSjhqEb3v4x4Iq4n8KKVAI%3D)

## Connect the MCP Server

Because the Model Context Protocol is a universal standard, connecting your preferred client is usually as simple as pasting a short JSON snippet or running a single terminal command.

To connect your AI agent or dev tool with AI Advisor MCP Server:

1. Update the appropriate **Configuration Example**with the **endpoint** and **header** information that applies to you.
2. Follow the **instructions** to apply the configuration in the agent or dev tool.
3. Start **asking questions** - remember to mention "Kentik" in your prompts

> **Note**: If you’re using a custom AI agent framework, see **Custom Applications** and **MCP Payload Reference** for additional details.

### Endpoints

Choose the URL for your region/cluster.

| Region | Endpoint URL |
| --- | --- |
| **US Cluster** | `https://api.kentik.com/mcp` |
| **EU Cluster** | `https://api.kentik.eu/mcp` |

### Required Headers

| Header Name | Value |
| --- | --- |
| `X-CH-Auth-Email` | <your-email@company.com> |
| `X-CH-Auth-API-Token` | <your-api-token> |

## Usage

Here are some tips on interacting with the AI Advisor MCP Server via your AI agent.

### Effective Prompting

> **IMPORTANT**: Always explicitly mention "**Kentik**" in your initial prompts.
>
> **Do**: `"Show me vmx devices from Kentik"` or `"What's the status of my Kentik devices?"`
>
> **Don't**: `"Show me all devices."`

If you don't include the “**Kentik**” keyword, the external AI agent may not realize it needs to invoke the **Kentik MCP tools** and might return a generic or incorrect response.

### **Real-time Progress**

While AI Advisor processes your query, you will see real-time progress updates in supported MCP clients (like Claude Desktop and GitHub Copilot). These notifications show the AI’s reasoning process:

* `🔍 Created session...` - Initial session created
* `⏱️ [2s] Analyzing device health metrics...` - Processing updates
* `⏱️ [4s] Checking interface utilization...` - Continued progress
* `✅ Completed` - Final answer ready

Next, let’s check out the **Configuration Examples** for the AI Advisor MCP Server.
