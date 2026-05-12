# Core

Source: https://kb.kentik.com/docs/core-overview

---

This article covers the Core section of the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(390).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A59Z&se=2026-05-12T09%3A35%3A59Z&sr=c&sp=r&sig=7KfwzSmLoS8pKeV2zgbH9jXkfQMxHJ2IhZivYLEFnpg%3D)

*Network Explorer, the landing page for Core, shows traffic conditions and links to other key portal modules.*

## About Core

The Core section of the Kentik portal enables you to understand the day-to-day operation of your network in environments including cloud, data center, WAN, and/or campus. The included modules are focused on traffic flows and their impact on the core elements of your network. These powerful tools enable you to easily browse and search through network data using pre-defined views, to drill deep into network flows, to visualize your network topology, and to create custom views (dashboards).

## Core Modules

The Core section of the portal includes the following modules and views:

* **Network Explorer**: The home module of Core is the place to go for a high-level overview of the current traffic on your network, including breakdowns by direction, by individual sites and clouds, and by up to 12 parameters (see **Network Explorer**). This page provides direct access to the views described in **Network Explorer Pages**. It also provides links to other pages in the Core section of the portal, as well as to workflows in other sections (i.e. Edge, Protect, Service Provider).
* **Data Explorer**: Data Explorer is our primary interface for manually exploring stored network data (flow records, BGP, SNMP, etc.). Data Explorer settings are translated into queries that return “views” made up of tables and graphs about traffic dimensions and metrics on specified devices during a specified timespan. **Data Explorer** is a powerful, deep-featured network analytics tool for troubleshooting, forensics, and BI.
* **Capacity Planning**: This workflow enables service providers and other network operators to quickly assess the utilization of interfaces relative to capacity so that they can take action before links become overloaded. You can create capacity plans for various groups of interfaces and evaluate utilization independently for each plan according to custom criteria. You’ll see current traffic volume as a percent of capacity, an estimate (based on current trends) of when traffic is likely to reach a target utilization range, and suggestions for rerouting traffic to optimize the use of your infrastructure (see **Capacity Planning**).
* **Insights**: This module enables the viewing and configuration of ready-made insights that automatically notify you of potentially interesting and/or concerning developments surfaced by our insights engine (see **Insights** and **Insights Reference**).
* **Forensic Raw Flow**: A viewer for raw flows from a set of data sources over a given time range, including, for each flow, the timestamp and the values of selected flow fields (device, protocol, source/destination IP, and source/destination port).

  > **Note:** This feature is not included by default in any Kentik edition. For further information, contact Kentik (see **Customer Care**).
* **Raw Flow Explorer**: This module enables you to directly examine flow records stored in Kentik Data Engine (KDE), Kentik's backend datastore (see **Raw Flow Explorer**).

## Network Explorer Pages

While some modules in the Core section (e.g., Data Explorer, Kentik Map, etc.) feature a unique interface that is specialized for a specific purpose, the pages in this section feature a consistent layout built around of set of standard components that, while customizable, are fundamentally similar from page to page (see **Main Panel Components**). This layout is used for two similar but distinct kinds of pages:

* **Aggregate Pages**: Reached via the **Explore Top Talkers** drop-down in the **Core** card at the upper right of the **Network Explorer** home page, these pages provide a complete picture of an entire class of traffic that is defined by a specific parameter (e.g., data source, dimension, or cloud provider). All instances of the class are listed, and aggregated network traffic data about those instances is displayed in statistics, graphs, and tables. For example, the page at Core » Network Explorer » **Applications** presents your network traffic broken out by the applications to which Kentik attributes the traffic. The structure of Aggregate pages is described in **Core Aggregate Pages**.
* **Details Pages**: These pages are structured similarly to Aggregate pages but focus on one individual instance of a class of traffic rather than aggregating information about the entire class. These pages are primarily reached via links in the tabs of the **Traffic Table** on the Network Explorer page and the Aggregate pages. The structure of Details pages is described in **Core Details Pages**.

## Main Panel Components

The following topics provide a high-level overview of the components that recur on the Network Explorer pages (home, Aggregate, and Detail) of the Core section of the portal (for more detail, see **Core Page Components**).

### Parameter Controls

Set the parameters of the query whose results are displayed in the page's graph and table. These controls typically include:

* **Filter**: Apply filters (specified values for selected dimensions) that narrow the traffic whose data is displayed. See **Filters Control**.
* **Aggregate**: Choose the calculation by which the included traffic data is aggregated. See **Aggregate Selector**.
* **Metric**: A metric is a combination of a unit (e.g. a bit) with a method of calculation (e.g., average per second) to create a quantifiable measurement (average bits/second). The measurements made on network traffic are used for counts, rankings (e.g. in a top-X list), and thresholds (e.g., in alerting). See **Metric Selector**.
* **Time Range**: Sets the time range for a preset duration back from the current time. See **Time Range Selector**.

### Traffic Graph

The **Traffic** graph is a visualization (see **Stacked Area Chart**) of the results of the network traffic query whose parameters are defined with the **Parameter Controls**. It shows the sum for all of the individual items listed in the table (e.g., devices, applications, etc., depending on the page). The graph is tabbed so that you can quickly see traffic broken into major subsets. In most cases the tabs include:

* **Total**: The sum of traffic in all of the subsets below.
* **Internal**: Traffic whose origin and destination are both within your network.
* **Inbound**: Traffic entering from somewhere outside your network.
* **Outbound**: Traffic leaving to somewhere outside your network.
* **Ingress** (Interfaces views only): Traffic coming into a device.
* **Egress** (Interfaces views only): Traffic going out of a device.
* **Through**: Traffic that both originates and terminates outside your network.
* **Other**: Traffic that is not classified as one of the above types of traffic (see **Traffic Classified as Other**).

> **Notes**:
>
> * Each flow is counted only once. For example, traffic that has been counted as inbound upon entering your network is not counted again as internal as it transits the network or as outbound when it leaves the network.
> * For more specific information on the **Traffic Graph** UI, see **Traffic Graph**.

### Traffic Table

The **Traffic** table displays the data returned from the network traffic query whose parameters are defined with the **Parameter Controls**. Two main types of tables are used, depending on the type of page you are on (see **Network Explorer Pages**):

* **Aggregate page table**: The table on an Aggregate page lists all of the individual instances of the "thing" that is being aggregated on the page. On the Devices Aggregate page, for example, the table lists devices. The columns of Aggregate page tables vary depending on the "thing" that the page is about.
* **Details page table**: The table on Details pages (e.g., a page devoted to a specific device) is tabbed. Each tab corresponds to a parameter by which traffic data can be ranked (top-X). The set of tabs that's visible on a given Details page is set with the **Customize** button at the top right of the table (see **Customize Tabs Dialog**). The columns shown in the table depend on the currently selected tab.
