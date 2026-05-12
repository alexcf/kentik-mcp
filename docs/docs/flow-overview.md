# Flow Overview

Source: https://kb.kentik.com/docs/flow-overview

---

This article provides an introduction to **network device flow**, the protocols used to transport it, and how Kentik leverages **flow sampling** to deliver deep, actionable network visibility.

> **Notes:**
>
> * For information on configuring devices to send flow to Kentik, see **Router Configuration** and **Universal Agents**.
> * For a guide on calculating baseline Flows Per Second (FPS) on your network devices, see **Estimating Flow Rate**.

## About Flow

Flow records are metadata about IP conversations that traverse network devices like routers, switches, or hosts.

* **Mechanism**: When enabled, a device collects these records in a cache and exports them to Kentik at specified intervals.
* **Key Data Points**: A standard flow record includes essential fields such as:

  + **Source/Destination IP**: The endpoints of the conversation.
  + **Protocol & Ports**: L4 details (TCP/UDP) that identify the type of traffic.
  + **Counters**: Byte and packet counts to measure traffic volume.
  + **SNMP Indices**: Input and output interface IDs for path tracking.

> **TIP:** For a list of devices that can send flow to Kentik, see **Supported Device Types**.

## Flow Protocols

Choosing the right flow protocol is a balance between your **hardware’s native export capabilities** and the **granularity of visibility** required. Use the table below to compare protocol support for key dimensions.

* **sFlow**: Statistical monitoring tool designed by **sFlow.org**.
* **IPFIX**: The IETF universal standard for flow export.
* **NetFlow v9**: Cisco's template-based protocol, nearly identical to IPFIX.
* **NetFlow v5**: Legacy protocol (called "JFlow" on Juniper).

### Protocol Comparison

The following table highlights the natively supported features by protocol. Ensure your hardware is compatible to successfully populate data to the **Kentik Data Engine**.

| Protocol | Best For | Basic Fields | Embedded Sampling Rate | IPv6 Support | Layer-2 (MAC/VLAN) | Payload Sample |
| --- | --- | --- | --- | --- | --- | --- |
| **sFlow** | **High-speed switching environments.** | Yes | Yes | Yes | Yes | Yes |
| **IPFIX** | **Modern, vendor-agnostic deployments.** | Yes | Yes | Yes | Custom Template\* | No |
| **NetFlow v9** | **Cisco-heavy infrastructure.** | Yes | No | Yes | Custom Template\* | No |
| **NetFlow v5** | **Older hardware with limited IPv6 needs.** | Yes | Yes | **No** | No | No |

*\*Support for MAC and VLAN fields requires hardware support for "****Option Templates****" via* ***NetFlow v9*** *or* ***IPFIX****.*

### Important Technical Considerations:

* **Layer-2 vs. Layer-3**: Note that some hardware will only report flow data for traffic that triggers a Layer-3/routing decision, meaning local Layer-2 traffic may not be reported.
* **Sampling Algorithms**: Different vendors use various mathematical algorithms for sampling; however, Kentik is designed to support and normalize data from all of them.
* **sFlow Advantage**: Unlike the other protocols, **sFlow** includes a sample of the actual packet payload, which can provide even deeper packet-level context.

## Flow Sampling

Flow sampling means exporting a metadata record for only one in every X flows (e.g., **1:10,000**). This balance allows you to maintain deep network visibility with efficient resource utilization by:

* **Reducing Infrastructure Load**: Lowers CPU/memory overhead on your routers and switches.
* **Optimizing Bandwidth**: Decreases the traffic required to export flow data to Kentik.
* **Ensures Real-Time Visibility**: Prevents devices from being overwhelmed during traffic spikes like DDoS attacks which could otherwise cause data gaps.
* **Scales Plan Capacity**: Helps you monitor more infrastructure within your Flows Per Second (FPS) limits.

### **Understanding the Two Layers of Sampling**

It is important to distinguish between where sampling occurs to avoid confusion in the Kentik UI:

* **Device-Level Sampling (Proactive):** This is configured by you on your routers or switches (e.g., `1:1000`). It reduces the load on your hardware and is also the rate Kentik uses to calculate total traffic.
* **Kentik Downsampling** **(Reactive):** This is triggered automatically by Kentik only if your incoming flow exceeds your **licensed FPS limit**.

Aim to set your device-level sampling high enough so that Kentik remains in Full Resolution (Green), avoiding the need for reactive platform downsampling.

### Flow Sample Rates

Kentik will accept all flow packets sent by customer routers up to the maximum limit permitted by the current active FPS license. If Kentik receives more flow packets per second than the license covers, Kentik will begin downsampling to maintain the contracted maximum.

> **TIP:** Kentik highly recommends configuring flow-exporting devices to send sampled flow data for more efficient downsampling.

Kentik will not sample less than a 5:1 ratio, meaning a minimum 20% of the flows will be saved. Devices with excessive flow records that require downsampling will be marked with a yellow caution indicator in the device listing. The actual flow rate will be displayed next to the configured rate. If the maximum flow has been reached and Kentik has reached the maximum sampling rates for proper resolution, the indicator will turn red and the super admin will be notified. There are multiple areas where downsampling indicators will appear in the following locations:

* **Devices Page** **»** **NMS Device List**
* **Devices Page** **»** **Telemetry Tab**
* **Organization** **Settings** **»** **Licenses Page**

#### Inventory Flow Statuses

The three inventory flow status indicators are:

* **Direct Flow Detected**: A green status that indicates successful **direct** flow connection coming from the router to Kentik, and the device is configured to send flow directly to Kentik’s collectors (not through `kproxy`).
* **Agent Detected**: A green status that indicates successful **agent-based** flow connection coming from Kentik’s `kproxy` software agent, and the device sends flow to `kproxy`, which then forwards to Kentik. This is common for devices behind firewalls or using private IPs.
* **No Flow Detected**: A yellow status that indicates a connection/configuration issue that needs attention. Kentik is not receiving any flow records from the device, and it occurs when the device is not configured to send flow, has no traffic, or the Sending IP is not set in the **Device Settings**.

#### Data Resolution Statuses

The three data resolution status indicators are:

* **Full Resolution**: A green status that indicates no downsampling is occurring. Flow is configured and being received at or below your licensed FPS rate, and all flow data sent by your device is stored in Kentik at full fidelity.

  + **Threshold**: **Received FPS** **≤ Licensed Max FPS**
* **Downsampled**: A yellow status that indicates Kentik is performing moderate **downsampling** to stay within your plan limits. Flow is configured but received FPS exceeds your licensed rate. Some flows are discarded, and the remaining flows are adjusted to represent the full traffic volume.

  + **Threshold**: **Received FPS > Licensed Max FPS AND downsampling ratio < 1:5**
* **Low Resolution**: A red status that indicates Kentik is performing excessive downsampling (1:5 or greater). Flow is configured but received FPS significantly exceeds your licensed rate, and substantial data fidelity loss due to heavy flow discarding.

  + **Threshold**: **Received FPS > Licensed Max FPS AND downsampling ratio** **≥****1:5**

> **Note:** If devices are over their max FPS, please contact your **Account Team** or adjust device sampling.

#### Downsampling

Downsampling is the process where Kentik reduces the number of flows per second (FPS) it ingests when you exceed your plan’s maximum FPS limit. Downsampling indicators are displayed in the UI as **Data Resolution Statuses**.

The risks of downsampling include:

* Reduced visibility of the details about your network traffic patterns.
* Data accuracy issues such as missing critical traffic patterns and anomalies that could indicate security threats or performance problems.
* Data loss triggered by exceeding your FPS license limits.

> **TIP:** Configure flow-exporting devices to send pre-sampled flow data rather than relying on Kentik’s automatic downsampling. This gives you more control over what data is retained.

#### Devices List Flow Status

The inventory flow status will display in the **Flow/Enrichment Status** column of the Devices list in the Manage Tab.

Click on a row to open the sidebar. The Flow Status from the Flow/Enrichment Status column and the Data Resolution status will be displayed next to **Flow** under the Telemetry section.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Flow Enrichment Status.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A44Z&se=2026-05-12T09%3A34%3A44Z&sr=c&sp=r&sig=Z0dfHDkwTEbIJRG04Ja2dimqL47eLFz88bjAENwLNSc%3D)

*Example of a the Flow/Enrichment Status column and Sidebar displaying the full Flow Status “Direct Flow Detected”*

#### Device Telemetry Tab Flow Status

The Telemetry Tab will display a tooltip icon to indicate there is a downsampling status. The flow status will display next to **Flow**, and the full data resolution status will display next to the **Stored Sample Rate**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image-1772480061575.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A44Z&se=2026-05-12T09%3A34%3A44Z&sr=c&sp=r&sig=Z0dfHDkwTEbIJRG04Ja2dimqL47eLFz88bjAENwLNSc%3D)

*Example of a Device’s Telemetry Tab indicating the Flow as Direct Flow Detected and a Downsampled Stored Sample Rate*

#### Licenses Page Flow Status

The **Licenses** page only displays data resolution statuses that will display on the details cards for CloudPak, DevicePak, and FlowPak plans. Kentik provides intelligent downsampling by dynamically adjusting data resolution to ensure you stay below your max FPS.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(854).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A44Z&se=2026-05-12T09%3A34%3A44Z&sr=c&sp=r&sig=Z0dfHDkwTEbIJRG04Ja2dimqL47eLFz88bjAENwLNSc%3D)

*Example of the Low Resolution and Downsampled Data Resolution status indicators*

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(855).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A44Z&se=2026-05-12T09%3A34%3A44Z&sr=c&sp=r&sig=Z0dfHDkwTEbIJRG04Ja2dimqL47eLFz88bjAENwLNSc%3D)

*Example of the Full Resolution Data Resolution Status Indicator*

### Sampling Best Practices

To ensure high data integrity, Kentik recommends following these configuration standards based on your specific ingestion method. Kentik categorizes data ingestion into two primary **Paks** based on where your traffic originates.

> **Notes:**
>
> * While it may seem counterintuitive, higher sampling rates (e.g., **1:1000** vs. **1:10,000**) provide **highly accurate measurements** for most common use cases.
> * For assistance with calculating sampling rates for your specific architecture, contact **Customer Care**.

#### Recommended Sampling Rates (Standard vs. Enhanced)

This table provides a range of flow sample rates based on device role, resolution, and link speed.

| Device Role | Resolution | 1 Gbps | 10 Gbps | 100 Gbps | 1 Tbps |
| --- | --- | --- | --- | --- | --- |
| **Edge/Internet** | Standard | N/A | 3k – 7k | 8k – 10k | 11k – 15k |
|  | Enhanced | N/A | 2k – 4k | 5k – 7k | 8k – 10k |
| **Data Center/Core** | Standard | 400 – 800 | 1k – 1.5k | 10k – 20k | 25k – 50k |
|  | Enhanced | 200 – 400 | 500 – 800 | 5k – 14k | N/A |

> **IMPORTANT**: To maintain high-fidelity analytics, Kentik will not downsample below a 5:1 ratio, preserving a minimum of 20% of flows.

### Ingestion Best Practices: FlowPak & CloudPak

Understanding the difference between these ingestion methods ensures you apply the correct sampling logic for your environment. For a visual representation of FlowPak & CloudPak, see **Kentik Plans**.

#### FlowPak (Physical & Virtual Infrastructure)

**FlowPak** refers to traditional flow data exported from physical or virtualized networking hardware (routers, switches, firewalls) using standard protocols like NetFlow, IPFIX, or sFlow.

* **Ingress Only (default)**: Enable flow sampling on **all interfaces** for **ingress traffic only**. Kentik automatically learns egress data from the ingress records of adjacent interfaces.
* **Avoid Double Counting**: Never enable both ingress and egress sampling on the same interface, as this will artificially double your reported traffic and FPS.
* **SNMP Validation**: Ensure flow traffic volume stays within **20%** of your **SNMP-recorded interface traffic** to verify configuration accuracy.
* **Egress Exceptions**: While ingress is the standard, egress flow creation may be required for specific scenarios like VPN services or traffic compression.

#### CloudPak (Public Cloud Environments)

**CloudPak** is the method used to ingest flow logs from public cloud providers. Unlike hardware-level configuration, this data is typically pulled via APIs or native **cloud log exports**.

* **AWS**: Fixed **1:N** sampling (e.g., **1:10**), configured in the **Public Clouds** settings in the Kentik portal.
* **Azure**: No sampling is applied by default. You must configure a fixed **1:N** sampling rate per exporter within the **Public Clouds** settings in the Kentik portal.
* **GCP**: Managed in the Google Cloud console per subnet. Kentik treats rates of **67%** or higher as "unsampled" and automatically converts percentage-based rates to a **1:N** format.
* **OCI**: Sampling rate is configured in OCI and applied per capture filter.

> **TIP**: You can see your sampling ratios on the **Public Clouds** page, with its graphs of the raw FPS, received FPS, and stored FPS.

### **Sample Rate** **Checklist**

Use this checklist to ensure your flow configuration is optimized for accuracy and performance within the Kentik platform.

* **Determine Target Sampling Rate**: Use the **Recommended Sampling Rates** table to identify the ideal rate based on device role, link speed, and desired resolution.
* **Configure Sampling**: Apply your chosen sampling rates to your devices in the Kentik portal under **Device Settings** and **Public Clouds**.
* **Validate and Adjust**:

  + Use the **Estimating Flow Rate** guide to measure your current Flows Per Second (FPS) levels and estimate requirements for peak events and future growth.
  + In the Kentik portal, check the **Device Telemetry** tab of **Device Settings** for **downsampling** indicators, which appear as yellow or red warnings.

FPS stands for flows per second, which is the rate at which flow metadata is sent to Kentik from a given customer. It directly correlates with the resources required to provide that customer with the Kentik service.

The process of reducing the number of flow records by sampling, where only one flow record is exported for every X flows.
