# Network Explorer

Source: https://kb.kentik.com/docs/network-explorer

---

This article covers the **Network Explorer** page in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(395).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A06Z&se=2026-05-12T09%3A35%3A06Z&sr=c&sp=r&sig=tHe8E81Ug7vKJ2VGcW9aFFcrVJME7iDfLZ9arq6L0bs%3D)

*Network Explorer provides an overview of current network traffic and links to detailed information.*

## Network Explorer UI

The **Network Explorer** page includes the following UI elements:

* **Share** (on subnav): A button that opens the Share dialog to enable you to share the current view (see **Sharing via the Share Dialog**).
* **Actions** (on subnav): A drop-down menu from which you can choose to export the current view as a visual report (PDF) covering the page’s visualizations and tables. A notification appears when the PDF is ready to download.
* **Insights** (on subnav): A button that pops up the **Insights Pane**, which shows insights related to traffic across your network (see **About Insights**).
* **Review your devices**: A callout that appears at the top of the page if any of your organization's devices need attention in order for their flow data to be ingested into Kentik. The **Review your devices** button takes you to the **Device Status** page, where you can address issues with the devices.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(396).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A06Z&se=2026-05-12T09%3A35%3A06Z&sr=c&sp=r&sig=tHe8E81Ug7vKJ2VGcW9aFFcrVJME7iDfLZ9arq6L0bs%3D)
* **Parameter controls**: The parameters of the query whose results are displayed in the page's graph and table (see **Parameter Controls**):

  + Aggregate Selector: The calculation by which traffic data displayed in the graph and table is aggregated.
  + Metric Selector: The metric by which query results are evaluated for top-X and displayed in the graph and table.
  + Time Range Selector: The duration, looking back from the current time, covered by the query whose results are displayed in the graph and table.
* **Traffic graph**: Visualizations of the results returned from the query specified with the **Parameter Controls** (see **Network Explorer Graph**).
* **Clouds and Sites**: A set of cards providing a high-level look at traffic for all of your organization's Clouds and sites (see **Clouds and Sites Pane**).
* **Network Explorer Table**: A tabbed table (see **Network Explorer Table**) in which each tab displays a top-X table corresponding to a dimension. The queries that return results for each tab are affected by the **Parameter Controls**.
* **Network Explorer Sidebar**: A set of cards that each contains high-level information and links related to an individual module/workflow in the portal (see **Network Explorer Sidebar**).

## Network Explorer Graph

The **Network Explorer** graph is a set of visualizations of the results returned from the query specified with the **Parameter Controls**. Depending on your actual traffic volume, there will be a tab for each of the following traffic profiles:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(397).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A06Z&se=2026-05-12T09%3A35%3A06Z&sr=c&sp=r&sig=tHe8E81Ug7vKJ2VGcW9aFFcrVJME7iDfLZ9arq6L0bs%3D)**Total**: All traffic regardless of profile.
* **Internal**: Traffic that both originates and terminates inside your network.
* **Inbound**: Traffic that originates outside your network and terminates inside your network.
* **Outbound**: Traffic that originates inside your network and terminates outside your network.
* **Through**: Traffic that both originates and terminates outside your network.
* **Other**: Traffic that is not classified by the above profiles.

> **Notes:** For a full description, see **Traffic Graph**.

## Clouds and Sites Pane

The **Clouds and Sites** pane is a set of cards providing a high-level look at traffic for each of your organization's clouds and sites. Each card includes the following UI elements:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(398).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A06Z&se=2026-05-12T09%3A35%3A06Z&sr=c&sp=r&sig=tHe8E81Ug7vKJ2VGcW9aFFcrVJME7iDfLZ9arq6L0bs%3D)**Heading**: The name of the cloud or site:

  + If the card is for a cloud, the heading is the name of the cloud provider and the card represents the sum of all traffic involving that provider. The heading links to an Aggregate page for that provider (see **Core Aggregate Pages**).
  + If the card is for a site, the heading is the name of the site and the card represents the sum of all traffic involving that site. The heading links to a Details page for that site (see **Core Detail Pages**).
* **Traffic data**: Statistics, in the current metric, of the traffic over the current time range (see **Traffic Data**).
* **Graph**: A single-area graph showing traffic for the cloud/site over the time-range set with the Time Range selector.

> **Note**: The **See All Sites** link at the top right of the pane takes you to the Aggregate page for your sites (see **Core Aggregate Pages**).

#### Traffic Data

The statistics in an individual card give values, in the current metric, for some or all of the following categories of traffic over the current time range (in descending order by volume):

* **Total**: The sum of traffic in all of the subsets below.
* **Internal**: Traffic whose origin and destination are both within your network.
* **Inbound**: Traffic entering from somewhere outside your network.
* **Outbound**: Traffic leaving to somewhere outside your network.
* **Through**: Traffic that both originates and terminates outside your network.
* **Other**: Traffic that is not classified by the above profiles.

## Network Explorer Table

The **Network Explorer** table is tabbed, with each tab providing information related to a different dimension. The currently displayed tabs (up to 12) are specified with the **Customize Tabs Dialog**.

* For information about the UI of a traffic table, see **Traffic Table UI**.
* For information about a specific dimension that is listed as being available for traffic tables, see **Dimensions Reference**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(399).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A06Z&se=2026-05-12T09%3A35%3A06Z&sr=c&sp=r&sig=tHe8E81Ug7vKJ2VGcW9aFFcrVJME7iDfLZ9arq6L0bs%3D)

> **Notes:**
>
> * The dimensions available for tabs are a subset of the dimensions available for group-by and filtering in **Data Explorer**.
> * For a given Kentik customer, the list of dimensions will be the same in all instances of the Customize Tabs dialog, but may vary over time depending on what types of devices you currently have registered with Kentik.

## Network Explorer Sidebar

The sidebar on the **Network Explorer** page is made up of a set of cards that each provide information related to a specific module/workflow in the Kentik portal and/or a link to that module. (This information is independent of the current settings of the parameter controls described in **Network Explorer UI**.) These sidebar cards include the UI elements listed below:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(400).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A06Z&se=2026-05-12T09%3A35%3A06Z&sr=c&sp=r&sig=tHe8E81Ug7vKJ2VGcW9aFFcrVJME7iDfLZ9arq6L0bs%3D)

* **Core**:

  + Links to the **Kentik Map** module, the **Library**, **Data Explorer**, **Alerting**, and **Insights**.
  + Explore Top Talkers: A button that opens a categorized drop-down menu of Aggregate pages (see **Aggregate Pages Categories**).
* **Synthetics**:

  + Status indicators: The Critical (red), Warning (orange), Healthy (green), and Pending (gray) indicators provide a count of the synthetic tests whose current status falls into each of those states.
  + Synthetics Dashboard: A button that links to the **Synthetics Dashboard**.
* **Capacity Planning**:

  + Status: A count of interfaces on your network whose utilization is in each of the various capacity planning states (healthy, warning, critical) shown both as a ring chart and a list.
  + Capacity Planning: A button that links to the **Capacity Planning** landing page.
* **Connectivity Costs**:

  + Estimated cost: The estimated cost you are paying per billing cycle (monthly) for connectivity.
  + Ingress/Egress: The sums of the Ingress/Egress rates stated for every interface that has been associated with any of your cost groups.
  + Connectivity Costs: A button that links to the **Connectivity Costs** landing page.
* **Peering**:

  + Inbound/Outbound traffic: A table showing the inbound and outbound rates, based on the aggregate, metric, and time range settings in the **Parameter Controls**, of traffic whose connectivity type is IX, Free Private Peering, or Paid Private Peering.

    > **Note:** If there's no traffic for a given connectivity type then that type will not be listed in the table.
  + Discover New Peers: A button that links to the **Discover Peers** landing page.
* **DDoS Defense** (if any attacks or mitigations are active in your organization):

  + Active attacks: The number of attacks currently in progress.
  + Active mitigations: The number of mitigations currently in progress.
  + Attacks within last 24 hours: The number of attacks in the last 24 hours (not affected by the current time range setting on the Network Explorer landing page).
  + DDoS Defense: A button that links to the **DDoS Defense** landing page.
* **CDN Analytics**:

  + Traffic percentage: Indicates the percentage of total subscriber traffic attributed to the top 5 CDNs and lists those CDNs as links that take you to their individual CDN pages.
  + CDN Analytics: A button that links to the **CDN Analytics** landing page.
* **OTT Service Tracking** (if your organization has any OTT traffic):

  + Total traffic: The total bitrate of traffic with which Kentik has associated a service (the sum of Fully-classified and Provider-only).
  + Traffic percentage: Indicates the percentage of total traffic attributed to the top 5 services, and lists those services as links that take you to their individual service pages.
  + OTT Service Tracking: A button that links to the **OTT Service Tracking** landing page.
