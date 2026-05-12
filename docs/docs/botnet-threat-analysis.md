# Botnet & Threat Analysis

Source: https://kb.kentik.com/docs/botnet-threat-analysis

---

This article covers **Botnet & Threat Analysis** dashboard in the Kentik portal.

## About Botnet & Threat Analysis

Kentik's Botnet & Threat Analysis helps you find traffic from infected or compromised hosts by enriching flow records in the Kentik Data Engine with IP reputation data from Spamhaus. Based on this information, we expose two dimensions that you can use for group-by or filtering, Botnet C&C (Command & Control) and Threat-list Host (see **Threat Feed Dimensions**).  
  
The **Botnet & Threat Analysis** page is a dashboard of pre-configured charts and tables using these dimensions to show traffic associated with known Botnet C&C hosts and other threats. By default, your network's impacted traffic is shown for the last hour, but you can quickly customize this using the settings in the **Query** sidebar (see **Botnet & Threat Subnav**). The charts and tables on this page make it easy to understand the extent to which your network traffic is associated with known threats, enabling you to make faster and better-informed decisions about such traffic.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(380).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A57Z&se=2026-05-12T09%3A39%3A57Z&sr=c&sp=r&sig=6ATc%2B8dA31CIGsYUaFUJxnouIzzyvWhxVkw1VSEm0hg%3D)

*The Botnet & Threat Analysis dashboard.*

## Botnet & Threat Analysis Page

The following topics cover the UI features of the **Botnet & Threat Analysis** page, highlighting any differences from a standard Kentik dashboard.

> **Note:** For further information about standard Kentik dashboards, see **Dashboards**.

### Botnet & Threat Page Layout

The **Botnet & Threat Analysis** page contains the following main areas:

* **Subnav**: A silver strip across top of page that contains page-wide controls (see **Botnet & Threat Subnav**).
* **Title bar**: The controls associated with the page's title (see **Botnet & Threat Title Bar**).
* **Display area**: The area where the page's panels are displayed (see **Botnet & Threat Panels**).

### Botnet & Threat Subnav

The subnav (silver strip across top of page) contains the following page-wide controls (see **Dashboard Subnav Controls**):

* **Refresh**: A button that updates the charts and tables in the dashboard’s panels.
* **Auto** **Update** (down-pointing triangle icon): A drop-down menu from which you can set auto-update to Off (default) or select an interval at which the charts and tables will automatically update:

  + When auto-update is on, the interval (5, 10, 60, 90, or 120 seconds) will be displayed in the subnav to the left of the down-pointing triangle.
  + When auto-update is off, the down-pointing triangle will be displayed alone (no auto-update indicator).

    > **Note**: The countdown to refresh starts over each time you apply changes and the new result is returned in the display area.
* **Share**: A button that opens a dialog enabling you to share the contents of the page (see **Portal View Sharing**). For this page, the only available **Report File Type** in the **Email** and **Subscriptions** tabs of the **Share** dialog is PDF.

  > **Notes**:
  >
  > + To download the entire page, use the **Export** option from the **Actions** drop-down in the subnav.
  > + To export an individual visualization, select Export from that panel's context menu (see **Botnet & Threat Panel UI**).
* **Actions**: A drop-down with multiple page-wide options (see **Botnet & Threat Actions**).
* **Dashboards**: A button that opens the **Dashboards Sidebar**, with which you can search or browse for any other dashboards available to your organization.
* **Query**: A button that opens the **Dashboard Query Sidebar**, where you can use the Time, Filtering, and Data Sources panes to specify your query.

  > **Note:** Since the **Botnets & Threat Analysis** page is not a guided mode dashboard, the Guided Mode UI is not available in this query sidebar.

#### Botnet & Threat Actions

The **Actions** drop-down includes the following options:

* **Clone**: Opens the **Clone Dashboard** dialog (see **Dashboard Navigation**).
* **Export**: Prepares a visual report (PDF) including all panels of the dashboard. A notification appears when the PDF is ready to download.

**Subscribe**: Opens the **Subscribe** dialog enabling you to create a subscription for this page (see **Subscription Tab UI**).

* **Unsubscribe**: Visible only if you’re subscribed to this page, this option opens a dialog enabling you to confirm that you wish to be removed from the subscription.
* **Preview as Tenant**: Opens a submenu from which you can choose a MKP tenant, after which you'll be taken to the page as it would appear to that tenant.

### Botnet & Threat Title Bar

The page title bar, located directly under the subnav, includes the following elements:

* **Favorite** (star): An icon indicating whether this dashboard has been designated as a favorite dashboard. Click the star to designate the page as a favorite (see **Library Bookmarks**).
* **Title**: Botnets & Threat Analysis (the name of the page).
* **View Description** (comment icon): A button that opens a **Dashboard Description** dialog containing a short description of the Botnet & Threat Analysis dashboard.

### Botnet & Threat Panel UI

The following UI elements are common to all panels on the Botnet & Threat Analysis page:

* **Title**: Located at the top left of the panel, the title is a link that opens the panel's visualization in **Data Explorer**, enabling you to customize the visualization with the **Query Sidebar Controls**.
* **Action** (vertical ellipsis): A  button at the panel's top right that pops up a menu with the following options:

  + Change Visualization: Opens a submenu from which you can select the visualization type (see **Chart Visualization Types**).
  + Export: Opens a submenu from which you can export the panel's data in one of the listed forms (see **Portal Export Options**).

> **Notes:**
>
> * The panel **Associated Raw Flows** doesn't have a linkable title or an Action menu.
> * The remaining panels that contain tables include a menu at the right of each table row. These menus are not currently functional.

## Botnet & Threat Panels

The following topics cover the visualizations included in the panels of the main display area of the **Botnet & Threat Analysis** page.

> **Note:** A panel for which no threat or botnet traffic has been detected by Kentik during the time range specified in the Time pane of the **Query** sidebar will display “No Results” instead of a visualization.

### Timeseries Comparison

The **Botnet & Threat-feed Timeseries Comparison** line chart plots the Kentik-detected activity of the two **Threat Feed Dimensions** (Botnet C&C and Threat-list Host) over the current time range (e.g., last hour). When you hover your pointer over the chart a popup shows the precise activity rate in bits per second for each dimension at that point in the time range. Clicking the labels “Botnet C&C” or “Threat-Feed” in the chart's legend will hide/show the plot for a given dimension.

### Botnet Route Prefix

The **Botnet Source & Destination Route Prefix** **Sankey Diagram** traces suspected botnet traffic over the current time range (e.g., last hour) for each unique combination of source and destination route prefixes. It features four columns representing dimensions of the traffic (left to right):

* **Src Route Prefix/LEN**: The prefix of the source block of IP addresses followed by the prefix length in bits, e.g., 141.193.37.0/24.
* **Protocol**: The name of the internet protocol being used, e.g., TCP.
* **Dest Botnet C&C**: The name of the malware used in the Botnet C&C, e.g., TrickBot.
* **Dest Route Prefix/LEN**: The prefix of the destination block of IP addresses followed by the prefix length in bits, e.g., 190.0.0.0/18.

Each column includes one or more color-coded segments representing an individual value for the corresponding dimension. Hovering on a segment opens a popup showing the dimension name, value, and traffic volume in Kbps.  
  
The segments are connected by horizontal gray bands that each represent botnet traffic for a unique combination of source and destination route prefixes, with the height of the bands representing the volume of that traffic. Depending on the number of segments present in each column, the bands may join or split in between. For example, the bands for TCP traffic with differing source route prefixes will join at the Protocol column but then might split at the Dest Botnet C&C column.  
  
Hovering on a band in between two columns darkens the band to distinguish it from the other traffic and opens a popup stating the band's traffic volume.

### Botnet Geographical Sources

This **Geo HeatMap** shows the geographic distribution of suspected sources of botnet traffic. Each source city is overlaid with a circle whose size corresponds to the relative volume of botnet traffic, and whose color indicates the percentile bracket corresponding to that volume.  
  
Use pointer scrolling to zoom the map view in or out and pointer dragging to navigate the map view in any direction. If no botnet traffic exists, the text “No results were returned” will be displayed and no circles will be shown on the map.

### Top Botnet ASes

This table shows you traffic, grouped by Autonomous System (AS), that crossed or terminated in your network during the time range set in the Query sidebar whose destination IP has been identified as that of a Botnet C&C host.  
  
The table can be sorted by clicking the header for the desired sort column; clicking the same header again will change the sort direction (ascending vs. descending). The numerical values in each column are totaled in the bottom row of the table.  
  
Clicking anywhere in a table row toggles on/off the selection of that row. Multiple rows can be selected at once, each of which will be highlighted in blue. If one or more rows are selected, the page's other panels will automatically refresh to show data only from the ASes corresponding to the selected rows. To clear all selected rows at once, click **Clear Selections** at the bottom left of the panel.  
  
 The table includes the following columns:

* **Destination AS Number**: The name of the destination AS, followed by a lozenge containing the ASN.

  + The name is a link to the Core details page (see **Core Details Pages**) for that AS in Network Explorer.
  + The lozenge is a button that opens a popup showing the registration country and providing links to the following views for the ASN: Kentik Market Intelligence, PeeringDB Record, and ASN Traffic View.
  + Hovering on either the name or the lozenge opens a tooltip with the name, country code of registration, and ASN.
* **Max Unique Dst IPs**: The count in this AS of unique destination IP addresses that are associated with a known Botnet C&C host.
* **Max Kbits/s**: The maximum rate of botnet traffic observed over the current time range (e.g., last hour) in kilobits per second. By default, the table is sorted by this column in descending order.
* **Max packets/s**: The maximum number of packets per second observed over the current time range (e.g., last hour).
* **Last Datapoint Kbits/s**: The rate, in kilobits per second, of botnet traffic observed in the most recent time slice of the current time range.

### Botnet Traffic Total

This gauge shows the maximum rate of botnet traffic observed over the current time range (e.g., last hour), along with the average and 95th percentile rates, all in kilobits per second. The color of the gauge will vary depending on which percentile bracket the total botnet traffic falls into.

### Threat-feed Route Prefix

The **Threat-feed Source & Destination Route Prefix** Sankey diagram visually traces suspected threat-list traffic over the current time range (e.g. last hour) for each unique combination of source and destination route prefixes. It features four vertical bars from left to right representing dimensions of the traffic, as follows:

* **Src Route Prefix/LEN**: The prefix of the source block of IP addresses followed by the prefix length in bits, e.g., 141.193.37.0/24.
* **Src Country**: The two-letter country code associated with the source route prefix.
* **Dest Threat-list Host**: The name of the threat-list host, e.g., SINKHOLE.
* **Dest Country**: The two-letter country code associated with the destination route prefix.
* **Dest Route Prefix/LEN**: The prefix of the destination block of IP addresses followed by the prefix length in bits, e.g., 190.0.0.0/18.

Hover over any vertical bar segment to bring up a popup showing the count of unique IP addresses for that route prefix, country code, or host name.

> **Note**: See **Botnet Route Prefix** for more on the common features of the Sankey diagrams on this page.

### Threat-list Geo Sources

The **Threat-list Geographical Sources** heat map shows the source cities for any suspected threat-list traffic identified by Kentik.

> **Note**: See **Botnet Geographical Sources** for more on the common features of geographical heat maps on this page.

### Top Threat-list ASNs

The **Top Threat-list ASNs** table shows you traffic, grouped by Autonomous System (AS), that crossed or terminated in your network during the time range set in the Query sidebar whose destination IP has been found on a threat-list. The columns of the table, as well as the table's behavior, are the same as those of the table in the **Top Botnet ASes** panel.

> **Note**: Traffic rates in this table are expressed in Mbits/s (megabits per second) while the rates on the **Top Botnet ASes** are in Kbits/sec (kilobits per second).

### Threat-list Traffic Total

This gauge shows the maximum rate of threat-list traffic observed over the current time range (e.g., last hour), along with the average and 95th percentile rates, all in megabits per second.

> **Note**: See **Botnet Traffic Total** for more on the common features of gauges on this page.

### Botnet Type by FPS

This bar chart shows the traffic rate in flows per second (FPS) for all types of botnets over the current time range (e.g., last hour). The names of all botnet types are listed in the vertical axis with the corresponding traffic rates represented by bars in the horizontal axis. An accompanying horizontal scale is presented at the bottom of the chart showing you the magnitude of each bar in FPS. The botnet types with the highest FPS are sorted to the top of the chart and the horizontal bars are color-coded based on the percentile range the FPS values fall into.  
  
Hovering over any horizontal bar brings up a popup showing the botnet type name, the FPS rate rounded to two decimal places, and the percentile range that the FPS rate falls into. On the far right of the chart is a vertical line, showing where the total FPS (all botnet types combined) falls on the horizontal FPS scale.  
  
Clicking anywhere on a horizontal bar highlights it in blue and all other panels on the page will automatically refresh with only the data for the selected botnet type. Click on additional bars to add them to the selection or click a selected bar to remove it from the selection. With each selection change, all panels on the page will automatically refresh.

> **Note**: To clear all selected bars at once, click the **Clear Selections** button at the bottom left of the panel.

### Threat-list Type By FPS

This bar chart shows the traffic rate in flows per second (FPS) for all types of threat-lists over the current time range (e.g., last hour).

> **Note**: See **Botnet Type by FPS** for more on the common features of bar charts on this page.

### Active Botnet Overview

This horizon chart shows an overview of the botnet types that were active over the current time range (e.g., last hour), separated into time slices of one minute. The names of the botnet types are listed vertically on the left side of the chart inside a corresponding swim-lane. Each swim-lane shows the traffic over time (if any) from left to right for a particular botnet type.  
  
Listed at the bottom of the chart in the horizontal axis are the UTC times over the current time range (e.g., last hour), separated into five-minute increments from left to right. Traffic is shown in vertical bars, each being the height of the botnet type’s swim-lane and the width of a one-minute time slice.  
  
Hovering your pointer on the graph reveals a vertical line and crosshair as well as time (YYYY-MM-DD HH:MM format) and traffic rate (kilobits per second) labels along the bottom right of each swim-lane. The time values in each label update as you move the vertical line left and right. The rate values update in the corresponding swim-lane label as you move the vertical line over a vertical bar.

### Active Threat-feed Traffic

This horizon chart shows an overview of threat-feed types that were active over the current time range (e.g., last hour), separated into time slices of one minute. The chart's behavior is the same as described for the chart in the **Active Botnet Overview** panel.

> **Note**: Traffic rates on this chart are expressed in Mbits/s (megabits per second) while the rates on the **Active Botnet Overview** are in Kbits/sec (kilobits per second).

### Associated Raw Flows

This table shows which raw flows are the data source for the other visualization panels on this page, including the following columns:

* **Timestamp**: The date-time of the flow record in YYYY-MM-DD HH:MM:SS format.
* **Destination Botnet CC**: The name of malware used in the Botnet C&C, e.g., TrickBot.
* **Destination Threat-list Host**: The name of the threat-list host, e.g., SINKHOLE.
* **Protocol**:The Kentik internal ID (integer) of the protocol used.
* **Device Name**:The device name associated with the flow record.
* **Source IP/CIDR**: The IP address or CIDR of the source machine for the flow record.
* **Source Port Number**:The port number of the source machine for the flow record.
* **Destination Port Number**: The port number of the destination machine for the flow record.
* **Destination IP/CIDR**:The IP address or CIDR of the destination of the flow record.
* **Sampled at Ingress bytes**: The number of bytes sampled at ingress when the flow record was captured.
* **Sampled at Ingress packets**: The number of packets sampled at ingress when the flow record was captured.
