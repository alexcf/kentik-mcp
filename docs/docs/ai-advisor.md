# AI Advisor

Source: https://kb.kentik.com/docs/ai-advisor

---

This article discusses the **AI Advisor** features of the Kentik portal.

> **IMPORTANT**:
>
> * Kentik AI must be enabled to use AI Advisor (see **Kentik AI**).
> * AI Advisor is **read-only**. Kentik does not request, nor want, write access to your infrastructure (see **Security & Privacy Commitments**).

![AI Advisor in overlay mode after being launched from the Alerting page sidebar.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AAD-overlay-alert-troubleshooting.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A37%3A16Z&sr=c&sp=r&sig=tRTDGqgYpMD%2BkV130Zbae3WdrTpL0VBjOT6hklbfwfI%3D)

*AI Advisor in overlay mode after being launched from the Alerting page sidebar.*

> **Note**: For more information about Kentik AI, see **Kentik AI Overview**and **Kentik AI Settings****.**

## About AI Advisor

Kentik’s AI Advisor is an advanced AI agent designed to act as your Network Intelligence partner. It guides triage, troubleshoots, and suggests remediation for network operations issues by combining rich data with natural language interaction:

* **Interpretation**: Takes your text-based question/prompt, determines the intent (e.g., check interface status, troubleshoot an alert), and, when applicable, makes a plan and includes its reasoning in the conversation.
* **Tool Execution**: Selects and executes the necessary internal/external data exploration and diagnostic tools (e.g., querying interface inventory, fetching device metrics, running a flow analysis, see **Supported Tools**). All progress, including steps taken and data gathered including visualizations, is shown in the conversation in real time.

### **Customize AI Advisor**

Make AI Advisor even more powerful by tailoring it to your organization’s needs with **Runbooks** and **Custom Network Context**.

#### **Runbooks**

Runbooks are predefined, Markdown-formatted text recipes to guide AI Advisor through the diagnostic steps for specific alerts. They can:

* Be assigned to one or more alert policies
* Help to start alert investigations with all relevant information already in place
* Help to minimize human error by ensuring a pre-defined, systematic troubleshooting approach for specific situations

#### **Custom Network Context**

Add Markdown-formatted text (up to 100K characters) of any additional information you choose to provide about your network environment. Here are some considerations:

* AI Advisor will use this custom network context in all conversations in your organization.
* Examples include:

  + Network design and architecture (campus, WAN, data center, cloud)
  + IP addressing schemas
  + Naming conventions for devices, interfaces, and sites
  + Device types and vendors
  + Maintenance window schedules
  + What the word “customer” means in your context
  + Critical applications and traffic patterns
  + Any specific operational procedures or constraints

> **Note**: Runbooks and Custom Network Context can be configured by a Super Administrator user in Organization Settings » **Kentik AI** (see **Kentik AI Settings**).

### UI Modes

AI Advisor is available in the Kentik portal in several UI modes: **Full-Page Mode**, **Overlay Mode**, and **Overlay Mode (Additional Context)**, as described below.

#### **Full-Page Mode**

![Kentik main nav menu showing the AI Advisor link highlighted.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AAD-main-nav-link-highlighted.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A37%3A16Z&sr=c&sp=r&sig=tRTDGqgYpMD%2BkV130Zbae3WdrTpL0VBjOT6hklbfwfI%3D)

* Select **AI Advisor** from the portal’s main nav menu.
* Conversation occupies the entire page for maximum readability.

![Interface status report indicating a loose cable connection issue and recommended actions.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AAD-fullpage-active(2).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A37%3A16Z&sr=c&sp=r&sig=tRTDGqgYpMD%2BkV130Zbae3WdrTpL0VBjOT6hklbfwfI%3D)

*An AI Advisor conversation in full-page mode.*

#### **Overlay Mode**

![Highlighted 'Ask' button in the Kentik portal header navigation.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AAD-ask-button-in-nav-highlighted.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A37%3A16Z&sr=c&sp=r&sig=tRTDGqgYpMD%2BkV130Zbae3WdrTpL0VBjOT6hklbfwfI%3D)

* The **Ask** button on most portal pages opens a resizable conversation window.
* Navigate the portal normally while the conversation remains open.
* Same agent and same UI features as full-page mode.
* Click the “full screen” icon to move the conversation to full-page mode.

![Device status overview with diagnostics and recommended actions for network troubleshooting.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AAD-overlay-active(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A37%3A16Z&sr=c&sp=r&sig=tRTDGqgYpMD%2BkV130Zbae3WdrTpL0VBjOT6hklbfwfI%3D)

*The same AI Advisor conversation in overlay mode.*

> **TIP**: In overlay mode on certain portal pages, e.g., **Devices** and **Alerting**, AI Advisor can automatically start troubleshooting based on what alert/device you are viewing; see **Overlay Mode (Additional Context)**.

#### Overlay Mode (Additional Context)

When you open AI Advisor on the following Kentik portal pages, it can automatically use the prompt “Troubleshoot this alert for me” and start investigating based on the device/alert you are viewing. A blue lozenge in the chat indicates the name of the alert that AI Advisor is using as context for its investigation.

![DDoS alert analysis showing traffic patterns and recommended actions for mitigation.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AAD-overlay-alert-details-page-context.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A37%3A16Z&sr=c&sp=r&sig=tRTDGqgYpMD%2BkV130Zbae3WdrTpL0VBjOT6hklbfwfI%3D)

*AI Advisor in overlay mode with additional context, in this case the ‘Alert Details’ for a particular alert, included in the conversation.*

| Portal Page | How to Launch AI Advisor |
| --- | --- |
| Alerting | Click **Investigate with AI Advisor** in the sidebar under Take Action (troubleshooting conversation starts automatically). |
| Alert Details | Click **Investigate with AI Advisor** in the sidebar under Take Action (troubleshooting conversation starts automatically). |
| Device Details | Click **Ask** in the page header (enter a prompt to start the conversation). |

## Get Started

Switching on AI Advisor in the Kentik portal is easy — just enable it in the settings and start using it:

1. Enable Kentik AI for your entire organization via the **Kentik AI Settings** page.

   > **Note**: Must be done by a Super Administrator user (see **Kentik AI Settings** for a step-by-step procedure).
2. Refresh any open browser tabs with the Kentik portal.
3. **Start a Conversation** with AI Advisor in the Kentik portal using one of the following:

   * Select **AI Advisor** in the main nav menu.
   * Click the **Ask** button in the header of most pages.
   * Click **Investigate with AI Advisor** in certain pages; see **Overlay Mode (Additional Context)**.

## Example Uses

Here are examples of the deeper, agentic actions that Kentik’s AI Advisor can perform:

> Note: Kentik feeds NMS device configurations into AI Advisor, allowing you to analyze configuration changes alongside flow data, metrics, and network topology (see **SSH Command Access for Kentik NMS**).

### Rapid Network Troubleshooting

AI Advisor accelerates incident response by allowing you to ask natural language questions about network issues instead of manually querying multiple dashboards and tools.

**Example Scenario**:

* An engineer notices degraded application performance and asks: "Show me interfaces with high packet loss in the last hour"
* AI Advisor queries the NMS metrics, identifies problematic interfaces, and presents the data in an easy-to-read format
* Follow-up questions like "What BGP neighbors are affected?" or "Analyze the traffic spike on interface ae0.3" provide deeper context

This conversational approach reduces mean time to resolution (MTTR) by eliminating the need to navigate multiple screens and construct complex queries.

### Cloud Network Visibility and Optimization

AI Advisor helps teams understand and optimize their multi-cloud network architecture by providing instant insights into AWS, Azure, and GCP environments.

**Example Scenario**:

* A cloud architect asks: "Show me all VPCs in my AWS account and their connectivity"
* AI Advisor retrieves the cloud inventory and visualizes the topology
* Follow-up questions like "What's the traffic between my AWS and Azure environments?" or "Find security group misconfigurations blocking traffic to subnet-abc123" help identify optimization opportunities

AI Advisor can trace network paths, identify cost-saving opportunities (like peering candidates), and surface security risks—all through simple conversational queries.

> Did You Know? Multi-Lingual Support
>
> AI Advisor uses a powerful language model that supports conversations in multiple languages, including Spanish, German, French, Japanese, Hindi, Arabic, and others. Feel free to ask your questions and receive answers in your preferred language.

## Using AI Advisor

This section covers how to use AI Advisor, along with guidelines and best practices.

### Start a Conversation

![AI Advisor interface displaying a query about interface status.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AAD-main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A37%3A16Z&sr=c&sp=r&sig=tRTDGqgYpMD%2BkV130Zbae3WdrTpL0VBjOT6hklbfwfI%3D)Start new conversations with AI Advisor in the following portal locations using these steps:

1. **Launch AI Advisor:** In the Kentik portal, either:

   1. Select **AI Advisor** from the main nav menu to open it in **Full-Page Mode**.
   2. Click **Ask** in the header of most pages to open it in **Overlay Mode**.
   3. On certain portal pages, click **Investigate with AI Advisor** to start a context-aware troubleshooting session (see **Overlay Mode (Additional Context)**.
   4. While already in a conversation, click **+ New**.![AI Advisor conversation where the agent is investigating agent failure with checklist for diagnostics and system checks.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AAD-conversation-reasoning.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A37%3A16Z&sr=c&sp=r&sig=tRTDGqgYpMD%2BkV130Zbae3WdrTpL0VBjOT6hklbfwfI%3D)
2. **Ask**: Type in your question/request into the chat field (see **Best Practices**), then press Enter to start your interactive conversation.

   1. In **Overlay Mode**, optionally click a suggested prompt to start the conversation.
3. **Watch**: AI Advisor starts generating a response, providing updates on steps being taken.
4. **Reasoning Panes**: When applicable, AI Advisor shows its plan-making reasoning in a collapsible pane at the top of each message. Use the **Hide reasoning**/**Show reasoning** buttonsto open and close the pane.
5. **Completed Task Panes:** Finished tasks are shown in panes with green check marks. The **View Details** button opens a sidebar with detailed task results, e.g., a visualization from Kentik’s Data Explorer, or search results from the Kentik Knowledge Base.

   > **TIP**: See **Best Practices**and**Guidelines** from this article for more tips on working with AI Advisor.

### Manage Conversation History

To view and manage conversation history, open AI Advisor in either overlay or full-page mode and follow these steps.

1. Click **History** (or **View history** in overlay mode) to open the conversation history.
2. From the **Recent** or **Company** conversation lists, select a conversation to view (or continue) it in the main AI Advisor pane.
3. ![Recent AI Advisor conversations list showing options to pin, rename, or delete.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AAD-fullpage-history-actions-menu.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A37%3A16Z&sr=c&sp=r&sig=tRTDGqgYpMD%2BkV130Zbae3WdrTpL0VBjOT6hklbfwfI%3D)Enter keywords in the **Search…** field, then click the conversation to open it in the main AI Advisor pane.
4. For your conversations (**Recent** list), click the vertical dots icon for a conversation and select from the following options:

   1. **Pin**/**Unpin**: Click **Pin** to keep the conversation at the top of the Recent list, and **Unpin** to remove from the top.
   2. **Rename**: Edit the Kentik-assigned conversation name to your preferred name.
   3. **Delete**: Remove the conversation from your AI Advisor history.
   > **Note**: Company conversations are read-only and can only be managed by the users who created them.

### Guidelines

Using AI can be confusing, so here are some guidelines for what you can expect AI Advisor to do (and not do):

| AI Advisor can… | AI Advisor cannot… |
| --- | --- |
| Interpret complex, natural language questions in multiple languages. | Make changes to your network device configurations (it is read-only). |
| Plan and execute multi-step troubleshooting and data-gathering workflows with an increasing number of Kentik-managed tools. | Access or understand proprietary knowledge, apart from what you configure in Runbooks and Custom Network Context (see **Customizations**). |
| Provide step-by-step reasoning for its answers. | Perform complex Network Intelligence tasks that are outside the scope of its defined tools (see **Supported Tools**). |
| Maintain a “memory” within a particular conversation. | Maintain a persistent "memory" across multiple conversations, apart from what you configure in Runbooks and Custom Network Context (see **Customizations**). |
| Ask clarifying questions, suggest next steps to take, and sometimes offer remediation advice. | Provide the same exact response to a question/prompt asked at different times. Gen AI inherently can produce a different, yet plausible, response every time it receives the same input. |
| Use predefined Runbooks to automate diagnostic procedures (e.g., how to troubleshoot an alert). |  |
| Occasionally make errors or make mistakes (known as "hallucinations") and provide confident but incorrect information. |  |

> **Note**: See Kentik’s **Security & Privacy Commitments**for more, and reach out to **Customer Care** with any lingering questions.

### Best Practices

Since AI Advisor is an agent, it thrives on clear instructions, context, and iterative feedback. Here are some tips for interacting with this highly capable, yet literal, digital colleague.

#### 1. Be Specific, Not Just Descriptive

Whenever possible, tell AI Advisor exactly what to do and what format you need.

| Goal | Less-Effective Prompt | Effective, Agent-Ready Prompt |
| --- | --- | --- |
| Get specific data | "Show me interface health. | "Find all interfaces with packet discards > 100/sec in the last 4 hours on devices tagged as 'core'." |
| Define a time | "What about last week?" | "Compare the top talkers from the last 24 hours to the top talkers from the same period last Tuesday." |
| Set a role | "Explain the alert." | "Act as a Senior Network Engineer and provide a root cause summary for the BGP alert in 3 bullet points." |

#### 2. Check and Redirect to Refine Execution

Use AI Advisor's reasoning panes (available within most conversations) to debug the workflow and improve your next prompt.

![Paying attention to AI Advisor's reasoning panes can help refine the troubleshooting strategy when necessary](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AAD-conversation-reasoning(2).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A16Z&se=2026-05-12T09%3A37%3A16Z&sr=c&sp=r&sig=tRTDGqgYpMD%2BkV130Zbae3WdrTpL0VBjOT6hklbfwfI%3D)

*Paying attention to AI Advisor's reasoning panes can help refine prompts (when applicable).*

| Principle | Technique | How to Use It |
| --- | --- | --- |
| Verify intent | Check the reasoning | If the answer seems wrong, check the reasoning pane(s). Did the AI correctly identify the device or the time window? Did it choose the wrong data tool? |
| Correct the plan | Redirect AI Advisor | Refine your prompt and have the AI try again: "Re-run the analysis but filter by interface description containing 'IXP'." |
| Confirm the tool | Reference specific data points | If the AI used the "Interface Inventory Tool" but missed the operational status, ask: "The last result was missing the operational status. Please call the Interface Inventory Tool again and ensure the `ifOperStatus` field is included." |

#### 3. Provide Context to Avoid Hallucination

Any AI Agent’s biggest risk is confident fabrication (hallucination). Whenever possible, provide the necessary context to anchor its response in reality.

* **Anchor to Data**: When asking for a definition or general concept, try to follow up with a local data question: "What is BGP session flapping? Now, check if we had any BGP session flaps on our core devices in the last 72 hours."
* **Define Unknowns**: If you use a custom term (like a non-standard device tag), define it in the prompt: "Find flow data for the 'Customer\_Gold' group (this group is defined by ASN 64511 and BGP Community 1234:56). Which country generated the most traffic to them in the last hour?"

## Supported Tools

The Kentik-managed “tools” that AI Advisor can use includes, but is not limited to the following:

| Name | Comment |
| --- | --- |
| On-Demand Connectivity Test | On-demand ping and traceroute. Requires Universal Agent capability installed |
| Data Explorer | Query On-Prem and Cloud Flows. Includes Cause Analysis. Does not support Legacy SNMP queries, kprobe and kappa flows |
| Metrics Explorer | Query NMS and other metrics using Metrics Explorer |
| Alerts List | Filtered list of Kentik Alerts based on time, status, type |
| Syslog | Syslog search in Data Explorer received through Kentik NMS |
| SNMP Trap | SNMP Trap search in Data Explorer received through Kentik NMS |
| Devices Inventory | Search Devices (works for legacy (flow only) and NMS-based devices) |
| Interfaces Inventory | Search Device Interfaces (works for legacy (flow only) and NMS-based devices) |
| Sites Inventory | Search the list of Sites with all relevant metadata |
| Cloud Pathfinder | Perform Cloud Pathfinder query and provide results back |
| Cloud Inventory - AWS | Search AWS inventory items and their metadata |
| Cloud Inventory - Azure | Search Azure inventory items and their metadata |
| Device Neighbors | List device neighbors detected by Kentik NMS feature |
| Kentik Knowledge Base | Search Kentik KB articles |
| Whois | Query whois service based on ASN or prefix, using RIPE Stat database |
| DNS lookup | Resolve DNS queries based on name or IP, supporting A, AAAA, PTR records |
| Saved Views | Search and view Saved Views |
| Runbooks | Get Runbook content to get instructions on how to handle user’s request or troubleshoot the problem |
| Site Map | Provides the context about the current Kentik Portal URL with descriptions and KB article. |
| ASN Lookup | Search for ASNs by number (or network name) and get the names and country codes. |
| BGP Mitigation Announcements | Retrieves BGP update announcements for RTBH (Remote Triggered Black Hole) and Flowspec mitigations. |
| Custom Applications | Query custom applications configured for the company using optional filters |
| Mitigation Config | This tool retrieves configuration details for Mitigation Platforms and Methods. |
| Mitigation Events | This tool retrieves and summarizes mitigation events (Active, Inactive, Failed, and Waiting). |
| Network Providers | The tools provides a search for “Provider” names used in Interface classification |
| Synthetics Test Results | Returns test results metrics for a given test |
| Synthetics Test Configuration | Returns test settings, thresholds, and alerting configuration in CSV format. |
| Synthetics Alarms | Returns alarm/incident data for synthetics tests. |
| Synthetics Agents | Returns agent details including location, site, status, and health information. |
| Synthetics Agent Down Alert | Returns private agent downtime alerts. |
| Azure Resource Graph Tool | Execute Azure Resource Graph (Kusto) queries against the customer's Azure subscriptions |
