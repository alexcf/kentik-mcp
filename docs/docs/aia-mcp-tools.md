# MCP Tools Reference

Source: https://kb.kentik.com/docs/aia-mcp-tools

---

Kentik’s **AI Advisor MCP Server** exposes three tools for interacting with **AI Advisor**:

* **ask\_question**
* **ask\_followup**
* **get\_session\_history**

> **Note**: This reference and the **MCP Payload Reference** are specifically for developers building custom AI agents or integrations with the Kentik MCP Server. Standard MCP clients like Claude Desktop handle these tools automatically in the background.

## **Rate Limits**

Because the MCP server uses the **AI Advisor REST API**, standard **rate limits** apply. The MCP server automatically handles periodic polling to respect these limits.

> **IMPORTANT:** Use the `ask_followup` tool instead of creating new sessions via `ask_question` whenever possible to minimize API calls and maintain context.

## **ask\_question**

Submit a new question to AI Advisor and create a new conversation session.

**Parameters**:

* `prompt` (string, required): Your question about the network

**Example**:

```
{
  "prompt": "Show me all devices that are currently down"
}
```

**Returns**:

* AI-generated answer in Markdown format
* Session ID for follow-up questions
* Real-time progress updates during processing

**Typical Processing Time**: 5-30 seconds

## **ask\_followup**

Ask a follow-up question in an existing conversation, maintaining context.

**Parameters**:

* `session_id` (string, required): UUID of the existing chat session.

  > **Note**: This value is dynamic and must be retrieved from the previous response.
* `prompt` (string, required): Your follow-up question.

**Example**:

```
{
  "session_id": "019ab80d-f7a1-739b-99e0-803f501fdea9",
  "prompt": "What about their interface utilization?"
}
```

**Returns**:

* AI-generated answer with conversation context
* Updated session information

**Note**: The `session_id` is provided in the response from `ask_question`. Save it to continue the conversation.

## **get\_session\_history**

Retrieve the complete conversation history for a session.

**Parameters**:

* `session_id` (string, required): UUID of the chat session

**Example**:

```
{
  "session_id": "019ab80d-f7a1-739b-99e0-803f501fdea9"
}
```

**Returns**: Formatted conversation history with all questions and answers.
