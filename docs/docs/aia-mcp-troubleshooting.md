# Troubleshooting

Source: https://kb.kentik.com/docs/aia-mcp-troubleshooting

---

When things don't work as expected with **Kentik’s AI Advisor MCP Server**, check these common scenarios:

## **Authentication & Connection Failures**

* **Symptoms**: "Authentication failed", "Invalid credentials", or timeout errors.
* **Quick Fixes**:

  + **Verify Your Region:** The most common error is mixing up the US and EU endpoints. Ensure your URL matches your account region (`api.kentik.com` vs. `api.kentik.eu`).
  + **Check Your Token:** Double-check that your API token is active in the Kentik Portal and hasn't been revoked.
  + **Firewall Rules:** Ensure your local network allows outbound HTTPS connections to the Kentik API endpoints.
  + **License Check:** Confirm with your admin that your specific account has AI Advisor access enabled.

## **MCP Server Not Appearing**

* **Symptoms**: The Kentik AI Advisor server doesn't show up in Claude Desktop or your chosen client.
* **Quick Fixes**:

  + **Validate Your JSON:** A missing comma or mismatched quote in the `claude_desktop_config.json` is the #1 cause of this issue. Run your config through a free online JSON validator.
  + **Check Your PATH:** Ensure `npx` is installed and available in your system's PATH (specifically required for Claude Desktop).
  + **Hard Restart:** Completely quit (don't just minimize) and restart your client application to force it to re-read the configuration file.

## AI is Ignoring Kentik Data

* **Symptoms:** The AI responds to your network questions with generic advice, hallucinates answers, or fails to query Kentik.
* **Quick Fixes:**

  + **The "Kentik" Keyword:** Ensure you are explicitly including the word "Kentik" in your prompt (e.g., *"Show me my Kentik devices"* instead of *"Show me my devices"*).
  + **Verify Connection State:** Open your client's developer settings and verify the Kentik MCP server is actually showing as "Connected."

## Slow Responses or Timeouts

* **Symptoms:** Queries are taking a long time to return an answer or are timing out entirely.
* **Quick Fixes:**

  + **Patience for Complex Queries:** Deep network queries *can* take 30-60 seconds to process. This is normal behavior. Watch the real-time progress indicators to ensure it's still working.
  + **Rate Limiting:** You may have hit the API rate limit. Wait 60 seconds and try again.
  + **Prompt Strategy:** Try breaking down massively complex questions into smaller, targeted follow-up questions using the `ask_followup` tool.
