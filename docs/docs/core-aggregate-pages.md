# Core Aggregate Pages

Source: https://kb.kentik.com/docs/core-aggregate-pages

---

This article covers the aggregate pages in the **Network Explorer** module of the Kentik portal.

> **Note:** For details pages in the Network Explorer module, see **Core Details Pages**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(418).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A44Z&se=2026-05-12T09%3A30%3A44Z&sr=c&sp=r&sig=7tGJgsOpzzcK06mtRDHpyWlKyZmV%2BUYR3lJSzqfMCIY%3D)

*The Aggregate page in Network Explorer for sites.*

## About Aggregate Pages

The **Network Explorer** aggregate pages are part of the Core section of the portal. To access these pages, go to Core » Network Explorer and click the **Explore Top Talkers** button in the Core card at the top of the right sidebar, then pick an aggregate from the drop-down menu.

The Network Explorer aggregate pages are each focused on one specific aspect (e.g., dimension or data source) of your overall network traffic. For example, the Applications aggregate page (Core » Network Explorer » **Applications**) presents your network traffic broken out by the applications to which Kentik attributes the traffic. These applications are listed in the **Aggregate Traffic Table**. By clicking on an individual application listed in that table you can drill down to the Details page for that specific application (see **Core Details Pages**).

## Aggregate Pages UI

The main content area in an aggregate page includes the following UI elements:

* **Share** (on subnav): Opens the **Share** dialog to enable you to share the current page (see **Sharing via the Share Dialog**).
* **Actions** (on subnav): A drop-down menu from which you can choose to export the current page as a visual report (PDF) covering the page’s visualizations and tables. A notification appears when the PDF is ready to download.
* **Insights** (on subnav): A button that pops up the **Insights Pane**, which shows insights related to the traffic type of this page (e.g., to applications on the Applications aggregate page; see **About Insights**).
* **Add Device** (Devices aggregate page only): Opens the Add Device dialog in a new tab (see **Device Settings Dialog**).
* **Parameter controls**: Sets the parameters of the query whose results are displayed in the page's graph and table (see **Parameter Controls**):

  + Filters control: Applies filters that narrow the traffic covered by the query results.
  + Aggregate selector: Chooses the calculation by which traffic data displayed in the graph and table is aggregated.
  + Metric selector: Sets the metric by which query results are evaluated for top-X and displayed in the graph and table.
  + Time Range selector: Sets the duration, looking back from the current time, covered by the query whose results are displayed in the graph and table.
* **Aggregate Traffic Graph**: Displays visualizations of the results returned from the query specified with the **Parameter Controls** (see **Aggregate Traffic Graph**).
* **Aggregate Traffic Table**: A top-X table presenting the results returned from the query specified with the **Parameter Controls**. The group-by dimension for the query is the dimension corresponding to the current page, e.g., when the page is Applications, the dimension is applications. The columns of the table are listed in **Aggregate Traffic Table**.

## Aggregate Traffic Graph

The **Aggregate Traffic** graph displays visualizations of the results returned from the query specified with the **Parameter Controls**. Depending on your actual traffic volume, there will be a tab for each of the following types of traffic:

* **Total**: The sum of traffic in all of the subsets below.
* **Internal**: Traffic whose origin and destination are both within your network.
* **Inbound**: Traffic entering from somewhere outside your network.

  > **Note:** In the Interfaces aggregate page, Inbound is replaced with Ingress.
* **Outbound**: Traffic leaving to somewhere outside your network.

  > **Note:** In the Interfaces aggregate page, Outbound is replaced with Egress.
* **Through**: Traffic that both originates and terminates outside your network.
* **Other**: Traffic that is not classified by the above profiles.

> **Note:** For a full description, see **Traffic Graph**.

## Aggregate Traffic Table

The **Aggregate Traffic Table** lists your network's top-X instances of the dimension corresponding to the current page. On the **Applications** aggregation page, for example, the table is a top-X list of applications on your network. The table, whose underlying query can be tailored with the **Parameter** controls, provides information and actions for each of the listed instances. The columns in a table vary depending on the specific aggregate page, but generally include the columns detailed in **Common Table Columns**.

> **Note:** For a full description of table UI, see **Traffic Table**.

## Aggregate Pages Categories

Other than the landing page, which provides an overall view of traffic, the **Network Explorer** aggregate pages are grouped into the following categories:

* **Network & Traffic**: Information related to non-cloud data sources including interface names and descriptions, port IDs, how and by whom it’s transported, etc:

  + Sites
  + Devices
  + Interfaces
  + Providers
  + Connectivity Types
  + Network Boundaries
* **IP & BGP Routing**: IP addresses (IPv4 or IPv6), protocol (e.g., TCP or UDP), TCP flags, and ToS, as well as routing information including source and destination AS, AS path, AS names, community, prefixes, and hops:

  + ASNs
  + AS Paths
  + BGP Community
  + INET Family
  + IP Addresses
  + Next-Hop ASNs
  + Packet Size
  + Protocols
  + Route Prefixes
* **Geographic**: Properties related to physical location:

  + Countries
  + Regions
  + Cities
* **Host Monitoring**: Metrics from host agents:

  + TCP Traffic
  + DNS Traffic
* **Application Context**: Factors related to context, e.g., whether a flow originated or terminated with a commercial CDN, or what “service” (port and protocol) it represents:

  + Applications
  + Services
* **Cloud**: Information sourced from VPC flow logs from cloud providers:

  + Amazon Web Services (AWS)
  + Google Cloud Platform (GCP)
  + Microsoft Azure
  + Oracle Cloud Infrastructure (OCI)

> **Note:** For more information on the dimensions represented by the above views, see **Dimensions Reference**.

---

© 2014-25 Kentik
