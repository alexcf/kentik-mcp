# Network Monitoring

Source: https://kb.kentik.com/docs/nms-overview

---

This article provides an introduction to Kentik’s **Network Monitoring System** (NMS).

![Kentik NMS gathers and presents key metrics on the availability, health, and performance of your network infrastructure.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/NO-NMS_diagram-624h1120w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A19Z&se=2026-05-12T09%3A30%3A19Z&sr=c&sp=r&sig=5YTIDW9Dnt1ZYkNBzUR6h88S6HelaDuPUvucuSJso04%3D)

*Kentik NMS shows key metrics on the availability, health, and performance of your network infrastructure.*

## About Kentik NMS

Kentik's **Network Monitoring System** (NMS) provides a base layer of visibility for your network by discovering and monitoring network infrastructure. Here are some of the features:

* Supports traditional use cases like detecting if a device goes down, graphing interface statistics, sending alerts, and creating dashboards.
* Supports data collection from SNMP and Streaming Telemetry (see **NMS via Streaming Telemetry**).
* Normalizes the collected data for consistency across dashboards, queries, and alerts regardless of the source.

## NMS Documentation

The following articles provide detailed information about various aspects of Kentik NMS:

* **Kentik NMS Configuration**: Learn how to set up Kentik's Universal Agent for NMS.
* **NMS Dashboard**: With this landing page for the NMS section of the Kentik portal, you can quickly understand your network status/performance via a customizable set of visualizations and tables.
* **SSH Command Access for Kentik NMS**: Use Kentik’s AI Advisor to analyze NMS device configuration changes alongside flow data, metrics, and network topology.
* **Metrics Explorer**: Query any aspect of your infrastructure to get the big picture about your network. Dive deep into the details of individual data sources.
* **NMS Devices**: See availability, health, and performance information for all devices that you've configured (with  the Universal Agent) to report NMS data to Kentik. Get extensive details about individual devices.
* **NMS Interfaces**: See availability, health, and performance information for all interfaces on your NMS-monitored devices. Get extensive details about individual interfaces.
* **Query Assistant**: Get answers about your infrastructure from queries that Kentik derives from your natural language questions.
* **NMS Setup**: A wizard to discover network devices to monitor with Kentik NMS.

The role of NMS is also covered in articles about Kentik's **Alerting** system, including the following:

* **About NMS Policies**
* **Alert Details Drawer**
* **NMS Alert Details Page**
* **NMS Policy Settings**
* **Add an NMS Policy**
* **Add an NMS Threshold Policy**

## NMS Quick Start

Adding NMS to Kentik involves the following steps:

1. **Deploy & Run Kentik’s Universal Agent**: NMS metrics come from your infrastructure to Kentik via a collector agent that is installed in the environment whose infrastructure you'd like to monitor. The agent can be deployed via Docker or as a Linux package. For further information and step-by-step instructions, see **NMS Setup Wizard**.
2. **Choose Devices to Monitor**: When the Universal agent is deployed and run it begins "discovery", which is the process of finding all SNMP-enabled devices within the specified IP ranges. You can then choose which devices to monitor and start monitoring those devices (see **NMS Discovery**).
3. **Monitor Your Network Metrics**: The metrics sent from your infrastructure to Kentik will be visible on the pages of the NMS section of the Kentik portal (listed in **NMS Documentation** above).

## Expanding Your NMS Capacity

Every Kentik customer subscribes to a Kentik edition that is augmented with one or more Kentik plans (see **About Licenses**), and every Kentik edition includes some NMS capacity (see **Metrics Limits by Edition**). To expand your organization's use of NMS beyond the below-listed limits, you'll need to supplement the included NMS capacity of your edition by purchasing an NMS Metrics plan. Once an NMS Metrics plan is added, you'll see it on your **Licenses Page**.

#### Metrics Limits by Edition

The table below shows the NMS capacity, in metrics per second (MPS), included with Kentik's various editions.

| Edition | MPS | Network Devices |
| --- | --- | --- |
| Platform Essentials | 100 | 3 to 10 |
| Pro | 250 | 8 to 25 |
| Premier | 250 | 8 to 25 |
| Classic | 100 | 3 to 10 |

## NMS Supported Devices

Use Kentik’s **NMS Supported Devices** tool to explore whether your devices are supported by Kentik NMS.

> **Tip:** Kentik is able to add device-specific datapoints to our data model. If datapoints for your devices are missing in Kentik NMS or you'd like to add unique data points, see **Custom Device Profiles** or reach out to us for assistance (see **Customer Care**).

## Using NMS via API

The tables and graphs returned from queries in the **Metrics Explorer** module of Kentik NMS can also be accessed via API:

1. In Metrics Explorer, click the **Actions** button in the subnav, then hover on **Show API Call** from the drop-down menu.
2. In the resulting submenu, choose one of the following:

   1. **For Data**: Opens the Data API Call via cURL dialog, which contains cURL that will request the Metrics Explorer's current table from a CLI such as Terminal.
   2. **JSON Input**: Opens the Data API JSON Input dialog, which contains JSON that can be used in a request body.
3. Copy the code in the dialog (manually or using the **Copy to Clipboard** button).
4. Use the code to get NMS results from a call made with the **Query Data Method** of the Kentik Query API.

Example JSON response snippet for when the **Query** sidebar is set to its default settings:

```
{
  "data": [
    {
      "timeseries": [
        {
          "timestampMillis": 1719497880000,
          "in-utilization": 0,
          "out-utilization": 0
        },
        {
          "timestampMillis": 1719497940000,
          "in-utilization": 2354,
          "out-utilization": 2420
        },
       // etc
      ],
      "last_in-utilization": 2289,
      "last_out-utilization": 3136,
      "device_name": "qfx_iad2_kentik_com",
      "name": "irb.130",
      "ifindex": 687
    },
    // etc.
  ]
}
```
