# AI Advisor MCP Server

Source: https://kb.kentik.com/docs/ai-advisor-mcp-server

---

Kentik’s **AI Advisor MCP Server** enables you to seamlessly integrate **AI Advisor** into external AI agents.

## **What is MCP?**

The **Model Context Protocol** **(MCP)** is an emerging industry standard for AI agents to access data and interact with external systems. It's supported by major AI platforms including:

* **Claude Desktop** **(Anthropic)**
* **GitHub Copilot**
* **OpenAI Agents**
* **Custom AI agent frameworks**: LangGraph, Google ADK, LlamaIndex, etc.

> **TIP**: Already familiar with MCP? Jump straight to **Getting Started** and start setting up `kentik-ai-advisor`.

## **Key Features**

This Kentik-hosted service provides a standardized interface for AI agents to query your network data and maintain conversational context.

* **Natural Language Queries**: Ask questions about your network in conversational language
* **Multi-turn Conversations**: Maintain context across multiple questions
* **Real-time Progress**: See AI Advisor's reasoning process as it works
* **Asynchronous Processing**: Handles long-running queries efficiently
* **Session Management**: Reviews and continues previous conversations

## **Use Cases**

To get the most out of the AI Advisor MCP Server, tailor your prompts to specific operational goals. Here are a few **NetOps**-related examples of how to format your requests to ensure the AI agent correctly routes the query to Kentik:

> **IMPORTANT**: Always explicitly mention "**Kentik**" in your prompts to ensure your agent uses its **MCP** **tools**.

| Operational Goal | Example Prompt |
| --- | --- |
| **Health Monitoring** | "Show me all devices with high CPU usage from **Kentik**." |
| **Troubleshooting** | "Check in **Kentik** which interfaces have errors in the last hour." |
| **Traffic Analysis** | "Ask **Kentik** what the traffic volume is on my external interfaces." |

## **Next Steps**

* Follow the **Getting Started** guide to go over prerequisites.
* Modify one of the **Configuration Examples**for your agent/env of choice and connect this MCP server.
* Study the additional background information:

  + **Kentik AI Advisor**
  + **Kentik AI Advisor API**
  + **MCP Specification**
