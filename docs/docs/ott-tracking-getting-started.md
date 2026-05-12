# Getting Started

Source: https://kb.kentik.com/docs/ott-tracking-getting-started

---

This article guides you through the setup and configuration of **OTT Service Tracking**, including how to enable DNS collection and customize your tracking thresholds.

## OTT Service Tracking Access

Using Kentik's OTT Service Tracking workflow requires the following:

* **Enabling Access**: Contact Customer Success (**see Customer Care**) to enable your organization to access OTT Service Tracking features.
* **Deploying DNS Taps**: Taps should be deployed on every possible DNS used by subscribers (see **Enable OTT DNS Collection**).

  + Taps can be deployed directly or via a DNS port span from the connected router using promiscuous mode.
  + Taps allow Kentik to map source and destination IPs against hostnames from the OTT service dictionary.
  + The more DNS coverage, the more accurate the OTT identification process

## Enable OTT DNS Collection

Using OTT Service Tracking requires a Kentik agent that taps into DNS traffic on your network. This deterministic data provides the “Fully Classified” baseline that Kentik’s Intelligent Classification model builds upon to detect the OTT services being consumed by your subscribers, and to estimate your total OTT traffic.

The process for enabling OTT DNS collection is summarized as follows:

1. Deploy Kentik Universal Agents on host machines in your network (see **Deploy a Universal Agent**).
2. Install and enable the **DNS OTT Tap** capability on each agent (see **Agent Capabilities**).

## OTT Configuration Page

The **OTT Service Tracking****Configuration** page, which is accessed by clicking the **Configure** button on the OTT Service Tracking page, allows you to configure both OTT thresholds and advanced filtering for the OTT module.

It includes the following UI elements:

* **Favorite**: A star next to the page title to add this page to the **Favorites** tab in portal search (see **Portal Search Tabs**).
* **Tabs**: Three tabs for configuring OTT Service Tracking:

  + **Thresholds**: Set interface utilization thresholds.
  + **Advanced Filtering**: Configure OTT workflow filtering.
  + **OTT Taps**: View agents with taps enabled, deploy new Universal Agents.

### OTT Thresholds Tab

The **Thresholds** tab lets you set utilization points for an interface, expressed as a percentage of its capacity, to trigger Warning or Critical statuses.

* **Warning threshold** (left slider): Sets the percentage above which utilization triggers a Warning status, shown in orange on the **Capacity** tab.
* **Critical threshold** (right slider): Sets the percentage above which utilization triggers a Critical status, shown in red on the **Capacity** tab.
* **Save**: Click to save changes; unsaved changes will be lost if you navigate away.

![Configuration settings for OTT service tracking interface capacity thresholds and utilization levels.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/OTT-configuration-main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A25Z&se=2026-05-12T09%3A32%3A25Z&sr=c&sp=r&sig=Y44KNWgNmTFWFGuLeVu81LNKlUHmdyknXCW6MGU6kXw%3D)

### OTT Advanced Filtering Tab

> **Note:** Kentik advises consulting your **Customer Success** representative before making any changes to this tab.

### OTT Taps Tab

The **OTT Tabs** tab displays the list of agents with the DNS OTT Tap capability deployed in your organization and provides a button to launch the wizard to **Deploy an OTT DNS Tap**.

* **Filter**: Filters the agent list based on a text search.
* **Deploy an OTT DNS Tap** (button): Opens the wizard to deploy a Universal Agent with the DNS OTT Taps capability (see **Deploy the Universal Agent**).
* **Agent List** (table): Lists agents with the DNS OTT Taps capability enabled (if any) and includes the following columns:

  + **Agent Status**: The agent's status (e.g., "Up").
  + **Name**: The agent's name.
  + **Agent Version**: The agent's software version.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(258).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A25Z&se=2026-05-12T09%3A32%3A25Z&sr=c&sp=r&sig=Y44KNWgNmTFWFGuLeVu81LNKlUHmdyknXCW6MGU6kXw%3D)

> **Note**: You can also view and install new Universal Agents via Settings » **Universal Agents**, accessible from the portal's main nav menu.

Now that your tracking is configured, let's move on to exploring your top-level traffic views in **Navigating OTT Dashboards**.
