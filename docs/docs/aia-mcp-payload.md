# MCP Payload Reference

Source: https://kb.kentik.com/docs/aia-mcp-payload

---

**Kentik’s AI Advisor MCP Server** uses standard **JSON-RPC 2.0** over HTTP. If you are building a custom integration rather than using an off-the-shelf client like Claude Desktop, refer to the payload structures below.

> **Note**: This reference and the **MCP Tools Reference** are specifically for developers building custom AI agents or integrations with the Kentik MCP Server. Standard MCP clients like Claude Desktop handle these tools automatically in the background.

### Request Format

When your agent invokes a tool like `ask_question` or `ask_followup`, it sends a `tools/call` request containing the specific tool name and necessary arguments.

```
{
  "jsonrpc": "2.0",
  "method": "tools/call",
  "id": 1,
  "params": {
    "name": "ask_question",
    "arguments": {
      "prompt": "Your network question here"
    }
  }
}
```

### Response Format

Successful tool executions return the AI's generated response mapped into a content array. This is typically returned as raw Markdown text, which your client application will need to render.

```
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "AI-generated answer in Markdown format."
      }
    ]
  }
}
```

### Progress Notifications

Because deep network queries can take up to 60 seconds, the server pushes asynchronous progress updates. Listen for these notifications to update your client's UI so the user knows the request hasn't hung.

```
{
  "jsonrpc": "2.0",
  "method": "notifications/progress",
  "params": {
    "progressToken": "unique-request-token",
    "progress": 50,
    "total": 100
  }
}
```
