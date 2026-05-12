# SNMP Traps & Syslogs

Source: https://kb.kentik.com/docs/snmp-traps-syslogs

---

This article discusses Kentik's Network Monitoring System (NMS) support of SNMP trap and syslog events.

## About SNMP Trap & Syslog

Kentik NMS supports **SNMP Traps** and **Syslog Messages,** two of the most widely used protocols for real-time network event communication. Ingesting these events into Kentik enables you to set up real-time alerts for things like hardware failures, interface status changes, and critical software log messages. You can also use Kentik’s Data Explorer to query this event data for troubleshooting or investigating unusual server logs.  
  
With SNMP trap and syslog event ingestion, you can centralize and correlate even more telemetry within Kentik’s Network Intelligence platform.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(581).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A42Z&se=2026-05-12T09%3A35%3A42Z&sr=c&sp=r&sig=YiQ%2BYwkkyrps02n2QmJUUEWig%2FgGWpMqF0VuYJ2Y4jQ%3D)

*Query/browse trap and syslog events in Data Explorer*

### SNMP Trap Features

SNMP traps are events immediately pushed by SNMP-enabled devices when specific conditions occur, bypassing both SNMP polling and Streaming Telemetry data collection intervals.  
  
With Kentik’s SNMP trap integration, you can:

* Receive SNMP traps from routers, switches, firewalls, and servers
* Filter and search trap events by name and OID
* Create alerts and notification policies based on SNMP traps
* Visualize trap events with other telemetry for contextually broader analysis

### Syslog Features

Syslog messages capture detailed system-level events across a wide range of devices.  
  
Kentik’s syslog integration enables you to:

* Collect syslog events from routers, switches, firewalls, and servers
* Filter and search syslog events by name, severity, and message content
* Create alerts and notification policies based on syslog messages
* Visualize syslog events with other telemetry for contextually broader analysis

## SNMP Trap & Syslog Setup

The steps for setting up SNMP trap and syslog for both existing and new NMS customers are as follows:

### Existing Kentik NMS Customers

Follow these steps to set up SNMP trap and syslog:

1. **Install and Enable Capabilities**:

   1. Navigate to Settings » **Universal Agents** from the main nav menu.
   2. Click the pencil icon to open the **Agent Details** drawer.
   3. Proceed to the next screen by clicking **Cancel.**
   4. Under Capabilities » Available**,** click **Install** for **SNMP Trap Receiver** and/or **Syslog Server.**

      ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(582).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A42Z&se=2026-05-12T09%3A35%3A42Z&sr=c&sp=r&sig=YiQ%2BYwkkyrps02n2QmJUUEWig%2FgGWpMqF0VuYJ2Y4jQ%3D)
   5. Enable each capability under Capabilities » Installed by clicking the switch ON (turns blue).
   6. Repeat for additional agents.
2. **Configure Devices**:

   1. Direct devices to send traps and syslogs to the configured Universal Agent(s).

> **Notes**:
>
> * SNMP trap and syslog ingestion require messages to come from a licensed NMS device.
> * When using a syslog forwarder, spoof the original device IPs in the forwarded logs to associate logs with their originating devices.
> * To change the TCP/UDP port numbers for SNMP trap and syslog ingestion, contact support. UI functionality will be added in the future.
> * Use the filter sidebar to show agents with **SNMP Trap Receiver** and/or **Syslog Server** capabilities enabled.
> * For more on working with Universal Agents, see **Universal Agents**.

### New Kentik NMS Customers

Follow these steps to set up SNMP trap and syslog:

1. Complete the steps in **NMS Setup**.
2. Follow **Existing Kentik NMS Customers** steps to install and enable the SNMP Trap Receiver and/or Syslog Server capabilities.

## Query Events with Data Explorer

Kentik supports querying SNMP trap and syslog metrics and dimensions in Data Explorer.

| Metric | Calculated as… | Default Dimensions | Description |
| --- | --- | --- | --- |
| SNMP Traps: Events | Average, 95th Percentile, Max | Timestamp, Trap Type, Device Name | Events from SNMP-enabled devices sent to a Kentik Universal Agent |
| Syslog: Events | Average, 95th Percentile, Max | Timestamp, Severity, Message, Device Name | Log messages from devices sent to a Kentik Universal Agent |

> **Note:** For complete lists of the related dimensions, see **SNMP Traps Dimensions** and **Syslog Dimensions**.

**Steps to Query SNMP Trap and Syslog Events**:

1. Navigate to Core » **Data Explorer** from the main nav menu.
2. In the Metrics pane of the Query drawer, select **SNMP Traps: Events** or **Syslog: Events** from the dropdown.
3. To optionally add dimensions:

   1. Click **Edit Dimensions** in the Dimensions pane.
   2. Use the search field to filter dimensions by “traps "or “syslog”, for example.
   3. Select dimensions and click **Save** to display the query results table.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(584).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A42Z&se=2026-05-12T09%3A35%3A42Z&sr=c&sp=r&sig=YiQ%2BYwkkyrps02n2QmJUUEWig%2FgGWpMqF0VuYJ2Y4jQ%3D)

*Select trap and syslog dimensions in Data Explorer.*

### Create Alert Policy from Query

Take your Data Explorer query for SNMP trap or syslog events further by creating a new alert policy directly from the query context.

1. After running your query, click Actions » **Create Alert Policy**.
2. Click **Continue** on the confirmation dialog to access the **Add NMS Alert Policy** page.
3. Follow the steps in **Create SNMP Trap & Syslog Alerts** to finalize the new policy.

> **Note:** For more, see **Add a Query-based Policy**.      

## Create SNMP Trap & Syslog Alerts

In addition to the ability to **Create Alert Policy from Query**, you can also create SNMP trap and syslog alert policies from scratch.

1. Navigate to Settings » **Alert Policies** and click Add Alert Policy.
2. In the dialog, select **NMS** and click **Continue.**
3. Enter a **Name,** optional **Description,** and set the policy **On/Off**. Click the right arrow to proceed.
4. Choose **Event** from the Policy Type dropdown.
5. Select **Syslog** or **SNMP Trap** under Event Type. Optionally, edit Dimensions or Filters, then click the right arrow to proceed.
6. Choose a **Severity** level for alerts generated from this policy.
7. Click **+ Add Condition Group** under Alert Conditions.
8. Set **Dimension,** **Operator,** and **Values** (e.g., "Trap Type", "is", "ciscoConfigManEvent").
9. Set Acknowledgement Required to **On** or leave as **Off.**
10. Select or add a **Notification Channel**.
11. Click **Create** to finalize the alert policy and return to the Alert Policies list.

> **Notes**:
>
> * Event type alerts require manual clearing.
> * For more details, see **Alert Policies**.

## SNMP Traps Dimensions

| Dimension Name (Portal) | Description |
| --- | --- |
| Dataset | The name of the dataset or log source. |
| Trap Type | The type of SNMP trap received. |
| Uptime | The duration for which the system has been operational. |
| Interface | Identifies a network interface. |
| Oper Status | The current operational status of an interface (e.g., up, down). |
| Admin Status | The configured administrative status of an interface (e.g., up, down). |
| BGP Peer | Identifies a BGP neighbor or peer. |
| BGP Peer State | The current state of the BGP peering session. |
| BGP Last Error Code | The last error code reported by BGP. |
| BGP Last Error Subcode | The subcode for the last BGP error. |
| Event ID | A unique identifier for a specific event. |
| Device ID | A unique identifier for a network device. |
| Timestamp | The date and time of the event. |
| Event Type | The type or category of the event. |
| Device Name | The name of the network device. |
| Trap OID | The Object Identifier (OID) of the SNMP trap. |
| Raw PDUs | Raw Protocol Data Units. |
| License Feature License ID | The ID of a license feature. |
| License Feature Name | The name of a license feature. |
| Command Source | The source of a configuration command. |
| Config Source | The source of a configuration. |
| Config Dest | The destination of a configuration. |
| Terminal User | The user interacting with the terminal. |
| Status Change Reason | The reason for a status change. |
| Reason | A general reason or description. |
| BGP Peer Last State | The last known state of the BGP peer. |
| BGP Last Error Reason | A description of the last BGP error. |
| Terminal User | The user interacting with the terminal. |
| Config Last Changed | The timestamp of the last configuration change. |

## Syslog Dimensions

| Dimension Name (Portal) | Description |
| --- | --- |
| Dataset | The name of the dataset or log source. |
| Transport | The method used to transmit log data, such as syslog, UDP, or HTTPS. |
| Format | The format of the log data, such as plain text, JSON, or XML. |
| Framing | The encapsulation or structure of the log data within a larger message or frame. |
| Trailer | Additional information or metadata appended to the end of a log message. |
| ParseError | An error that occurred during parsing or processing of the log data. |
| Priority | The level of importance or urgency associated with a log event. |
| Facility | The system or application that generated the log message. |
| Severity | The level of seriousness of a log event. |
| Hostname | The name of the host or machine that generated the log message. |
| Application Name | The name of the application or software that generated the log message. |
| Process ID | The unique identifier of the process that generated the log message. |
| Message | The main content or body of the log message, often containing specific details about the event or error. |
| Version | The version number of the software or application that generated the log message. |
| Message ID | A unique identifier for the log message, used for tracking and correlation. |
| Structured Data (JSON) | Additional structured data included within the log message, often in JSON format. |
| Event ID | A unique identifier for a specific event or occurrence. |
| Device ID | A unique identifier for a specific device. |
| Timestamp | The date and time when the log event occurred. |
| Event Type | The type or category of the log event, such as information, warning, or error. |
| Device Name | The name of the device that generated the log event. |
| Raw Log | The original, unprocessed log entry. |
