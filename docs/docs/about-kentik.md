# About Kentik

Source: https://kb.kentik.com/docs/about-kentik

---

This article provides a basic introduction to Kentik, with answers to the following questions:

* **What is Kentik?**
* **What traffic data is collected?**
* **What synthetic testing is supported?**
* **How is flow data collected?**
* **How do I access my data?**
* **Anything else I should know?**

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(5).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A41Z&se=2026-05-12T09%3A28%3A41Z&sr=c&sp=r&sig=XTHHgwjal2AnMhXyDCMkZioj47enHDw8ByBcUqj9b78%3D)

*Kentik achieves unrivaled network observability by integrating flow-based analytics, NMS, and synthetic testing.*

## What is Kentik?

Kentik is an open, scalable platform for collecting, analyzing, and visualizing data about your networks. It supports both on-prem and cloud resources, correlating real and synthetic traffic data for real-time and historical insights.  
  
Kentik's purpose-built platform sets up quickly, identifying and explaining unusual activity with real-time alerts for performance issues and attacks. The web-based portal offers advanced traffic analytics, synthetic testing, and network protection with alerts and mitigations. It also integrates with other tools and systems via Kentik’s REST APIs.

## What traffic data is collected?

Kentik collects two main types of traffic data:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(6).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A41Z&se=2026-05-12T09%3A28%3A41Z&sr=c&sp=r&sig=XTHHgwjal2AnMhXyDCMkZioj47enHDw8ByBcUqj9b78%3D)

*Kentik supports flow data export from a huge variety of devices.*

* **Flow data**: Includes packets traversing physical or virtual devices (see **Supported Device Types**), collected using protocols like protocols sFlow, IPFIX, and NetFlow (see **About Flow**):

  + **Data Center**: Physical devices cache and send flow data to Kentik at a specified interval.
  + **Cloud Provider**: Publish flow logs to cloud buckets for Kentik to access.
* **Metrics**: Kentik’s Network Monitoring System (NMS) collects metrics from entities supporting SNMP and/or Streaming Telemetry

  + Normalizes data for consistent dashboards, queries, and alerts.
  + Supports traditional device monitoring, graphing, alerting, and dashboard creation, while leveraging other data for deeper analytics.

* **SNMP** (non-NMS): Provides interface names/descriptions and flow level validation (see **SNMP OID Polling**).
* **GeoIP**: Determines the country, region, and city of flow sources and destinations.
* **BGP**: Extracts source/destination AS Path and community information (see **BGP Overview**), enabling features like **Discover Peers**.
* **Host traffic data**: Correlates host traffic with flow to provide host information, including URLs, DNS queries, and performance metrics (see **Host Traffic Dimensions** and **Host Traffic Metrics**).
* **Classification data**: Provides business intelligence on the role of interfaces for traffic ingress and egress (see **Interface Classification**).
* **Threat feeds**: Correlates daily data from Spamhaus with flow data to identify security threats.

> **Note**: For a more details on the data stored in Kentik, refer to the **Dimension Categories** and **Dimensions Reference**.

## What synthetic testing is supported?

Kentik's Synthetics workflows are easy to set up and cost-effective to run. Enable testing by deploying Kentik's `ksynth` software agent in public or private contexts (see **About Synthetics Agents**):

* **Public Agents**: Part of the Kentik Global Agent Network, located in major Internet hubs and cloud regions (AWS, GCP, Azure, OCI). Accessible to all Kentik customers.
* **Private Agents**: Deployed in your physical infrastructure or cloud resources, accessible only to your organization.

Run continuous ping and traceroute tests with public or private agents to generate key metrics (latency, jitter, loss) for network health and performance. Kentik intelligently guides synthetic testing based on actual traffic patterns, optimizing resources for maximum impact (see **Synthetics**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(7).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A41Z&se=2026-05-12T09%3A28%3A41Z&sr=c&sp=r&sig=XTHHgwjal2AnMhXyDCMkZioj47enHDw8ByBcUqj9b78%3D)

*The Synthetics Dashboard is the landing page for Synthetics.*

## How is flow data collected?

Kentik collects flow data from various sources:

* **Direct**: Routers and switches send flow data directly to Kentik servers (see **About Devices**).
* **Host agent**: The Flow Proxy capability of Kentik’s **Universal Agent** collects data from hosts, aiding in performance issue debugging by analyzing TCP retransmits per flow.
* **Cloud providers**: Supported cloud resources (e.g., AWS, Azure, GCP, OCI) generate flow logs and publish them to cloud buckets, which Kentik accesses via provider APIs (see **Cloud Overview**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Proxy_data_flows-674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A41Z&se=2026-05-12T09%3A28%3A41Z&sr=c&sp=r&sig=XTHHgwjal2AnMhXyDCMkZioj47enHDw8ByBcUqj9b78%3D)       

## How do I access my data?

Access your data in Kentik with the following methods:

* **Portal**: Use the portal’s UI (Data Explorer, Dashboards, Query Editor) to view your data.
* **APIs**: Access your data programmatically via Kentik APIs (see **About Kentik APIs**).
* **Firehose**: Use Firehose, supported by the `ktranslate` agent, to integrate enriched flow records into non-Kentik analytics systems or data lakes (see **Using Kentik Firehose**).

## Anything else I should know?

To get started with Kentik, explore the **Portal Overview** for key features of Kentik's v4 portal.

For questions, contact **Customer Care**.
