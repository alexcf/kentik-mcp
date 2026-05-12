# Core Page Components

Source: https://kb.kentik.com/docs/core-page-components

---

This article discusses common UI components of **Core** pages in the Kentik portal.

## Parameter Controls

The parameter controls, located at the top of a Core page below the title bar, set the parameters of the query whose results are displayed in the page's graph and table.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(401).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A46Z&se=2026-05-12T09%3A41%3A46Z&sr=c&sp=r&sig=PmZTqf4yz%2Bfxw1w3yN9c%2BNCzrwf6JeQ0bNgDoXI0OsI%3D)

### Filters Control

The **Filters** control enables you to apply filters that narrow the traffic covered by the query results displayed in the graph and table on a given aggregate or details page. On the Devices aggregate page, for example, by default the results would show all traffic — within the time range specified with the **Time Range Selector** — that is related to all devices. With the filter control you can specify one or more filters that result in the graph and table being updated to show only a subset of the traffic.  
  
Filters are applied by selecting a dimension and then selecting the value you want to filter on. You can further refine the results by choosing additional dimension/value pairs. When no filters are applied, "No filters applied" appears on the **Filters** control. If one or more filters are applied, the number of applied filters is displayed (for example, "1 filter applied").

#### Filters Control Dimensions

Available filters currently include:

* **ASN**: The origin ASN associated with the source/destination IP of the flow.
* **Next Hop ASN**: The BGP next-hop IP address, either IPv4 or IPv6, for the source/destination IP of the flow (see **About BGP**).
* **Provider**: The name of a provider, identified by Kentik using **Provider Classification**, via which traffic from a given externally facing interface reaches the Internet.
* **AS Path**: The BGP ASPATH for the flow’s source/destination IP.
* **BGP Community**: The set of BGP communities associated with the flow’s source/destination IP.
* **Site Country**: The name of a country in which your organization has sites; enables the grouping, with a single dimension, of traffic from all sites in a given country.
* **Country**: A two-letter country code associated with the source/destination IP of the flow (see **Geolocation Dimensions**).
* **Region**: The full-text English name of the region (state or province, e.g., “California”) associated with the source IP of the flow.
* **City**: Full-text English name of the city (e.g., “San Francisco”) associated with the source IP of the flow.
* **Site**: The name of a specific user-defined physical location (e.g., a data center) to which you can assign one or more devices (see **About Sites**).
* **Device**: The name of a network asset such as a router, switch, or host (see **Supported Device Types**).
* **Interface Name**: The vendor-defined name (e.g., “GigabitEthernet0/1”) of the device interface (physical or logical) through which flow ingressed/egressed.
* **Interface Description**: A user-provided description (e.g., “Connected to upstream ISP”) of the device interface (physical or logical) through which flow ingressed/egressed.
* **Application**: The name of a service, determined based on the factors described in **About Applications**.
* **Protocol**: An IP protocol number. When Protocol is chosen as the filter type, protocol names and numbers are listed in the **Select value** menu.
* **Port**: Layer 4 source/destination port (e.g., 80, 443).
* **Connectivity Type**: The role of an interface in the overall network, such as transit, ix, paid peering, etc. (see **Connectivity Type Attribute**).
* **Network Boundary**: An interface classification attribute that designates a given interface as internal (no direct connection outside your network) or external (see **Network Boundary Attribute**).
* **INET Family**: The address family of the flow, either 4 (IPv4) or 6 (IPv6).

> **Note:** Some of the filters listed above aren't available in all pages.

#### Filters Control Popup

The **Filters** control popup opens when you click on **Filters** in the parameter controls. The popup contains the following UI elements:

* **Dimension selector**: A drop-down menu from which you can choose the dimension of the filter (see **Filters Control Dimensions**).
* **Value selector**: **Select value** displays a drop-down menu from which you can choose a value to filter on. The available values vary depending on the dimension. To narrow the list, enter text into the filter field at top.
* **Remove**: The **X** at the right side of any filter removes the filter.
* **Add a Filter**: Adds a new filter for which you can specify dimension and value.

#### Using the Filters Control

To filter the results in the table:

1. On one of the **Core Aggregate Pages**, click the **Filters** control, then click **Add a Filter** in the resulting popup. The popup will expand to include the full set of filtering controls (see **Filters Control Popup**).
2. Choose a dimension (e.g., Country) from the drop-down **Dimensions** selector.
3. Choose the value you want to use to filter the results (e.g., Canada). The filter will be applied, and the graph and table will automatically update to show only traffic from or to Canada.
4. To add another filter, click **Add a Filter** again, then choose a dimension (e.g., Protocol) and a value (e.g., UDP). The filter will be applied, and the graph and table will automatically update to show only UDP datagrams from or to Canada.

### Aggregate Selector

The **Aggregate** selector is used to choose the calculation by which traffic data displayed in the graph and table is aggregated:

* **Average**: Calculates the top-X based on average traffic volume over the specified time range.
* **95th Percentile**: Calculates the top-X based on the time slice whose volume was at or below 95% of the time-slices within the aggregation period.
* **Maximum**: Calculates the top-X based on the time-slice whose volume was the highest of all time-slices within the aggregation period.

### Metric Selector

A metric is a combination of a unit (e.g., a bit) with a method of calculation (e.g., average) and a unit of time (e.g. one second) to create a quantifiable measurement (average bits/second). The measurements are made on network traffic and used for counts, rankings (e.g., in a top-X list), and thresholds (e.g., in alerting). The following metrics are available from the **Metric** selector:

* **Bits/s**: Bits per second
* **Packets/s**: Packets per second
* **Flows/s**: Flows per second

### Time Range Selector

The **Time Range** selector sets the duration, looking back from the current time, covered by the query whose results are displayed in the graph and table. The following time ranges are available:

* Last 5 Minutes
* Last 15 Minutes
* Last Hour
* Last 3 Hours
* Last 6 Hours
* Last 1 Day
* Last 2 Days
* Last 3 Days
* Last Week
* Last 2 Weeks
* Last 30 Days
* This Month
* Last Month

## Traffic Graph

Located toward the top of the page on both **Core Aggregate Pages** and **Core Details Pages**, the traffic graph (also called a tabbed graph) presents visualizations (see **Stacked Area Chart**) of the results returned from the query specified with the **Parameter Controls**. The traffic data that is plotted in the graph corresponds to the current page (e.g., on the Devices aggregate page the data pertains to traffic on all devices, while on a Device details page the traffic pertains to an individual device).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(402).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A46Z&se=2026-05-12T09%3A41%3A46Z&sr=c&sp=r&sig=PmZTqf4yz%2Bfxw1w3yN9c%2BNCzrwf6JeQ0bNgDoXI0OsI%3D)

#### Traffic Graph Tabs

The graph is tabbed to allow you to quickly see various subsets of the data (see **Simple Traffic Profile**). In most situations (see notes below), the available tabs show data for the following types of traffic:

* **Total**: The sum of traffic in all of the subsets below.
* **Internal**: Traffic whose origin and destination are both within your network.
* **Inbound** (not on Interfaces pages): Traffic entering from somewhere outside your network.
* **Outbound** (not on Interfaces pages): Traffic leaving to somewhere outside your network.
* **Ingress** (Interfaces pages only): Traffic coming into a device.
* **Egress** (Interfaces pages only): Traffic going out of a device.
* **Through**: Traffic that both originates and terminates outside your network.
* **Other**: Traffic that is not classified as one of the above types of traffic (see **Traffic Classified as Other**).

> **Notes:**
>
> * The traffic represented by some of the above tabs may in certain situations represent an insignificant portion of overall traffic. A given tab may not present if there is insufficient traffic to display.
> * The division of traffic into the subsets described above is based on the Network Boundary attribute assigned to interfaces during Interface Classification (see **Classification Overview**).
> * Each flow is counted only once. For example, traffic that has been counted as inbound upon entering your network is not counted again as internal as it transits the network or as outbound when it leaves the network.

### Traffic Graph UI

The traffic graph UI includes the following elements:

* **Tab indicator**: A blue bar displayed at the top of the graph indicates the currently displayed tab.
* **Traffic information**: The following information is shown above the graph for each of the data sets (Total, Inbound, Outbound, etc.) that has a tab in the graph:

  + **Color indicator**: Shows the color used to represent the traffic data in the graph (e.g., cyan for Total, orange for Inbound, green for Outbound), thereby serving as a legend.
  + **Traffic label**: Indicates the type of traffic (Total, Inbound, Outbound, etc.), which is also the name of the tab.
  + **Current volume**: The current volume of each type of traffic, expressed in the chosen metric.
* **Graph**: A visualization of the data specified with the **Parameter Controls** (see **Traffic Visualization**).
* **View in Explorer**: Opens **Data Explorer** for further exploration of the device’s traffic. The Data Explorer controls will be set to query about the traffic shown on the current tab of the graph.

#### Traffic Visualization

A visualization of the data specified with the **Parameter Controls**. The horizontal axis corresponds to the **Time** **Range** setting and the vertical axis corresponds to the **Metric** setting. The components are plotted in the graph depending on the tab:

* **Total tab**: In addition to the Total, the graph also plots all types of traffic that are included in the data specified with the **Parameter Controls** (Internal, Inbound, Outbound, etc.). The data from every tab above the graph is displayed as a part of the Total tab’s graph.
* **Remaining tabs**: The graph plots the top-X components that make up the overall traffic corresponding to the tab. On the Devices aggregate page, for example, if the tab is **Inbound**, then the areas of the stacked area chart each represent inbound traffic on one of the top-X devices listed in the table beneath. Click any of the remaining tabs **(Internal**, **Inbound**, **Outbound**, **Through**, **Other**, **Ingress**, or **Egress**) to see the graphical data for that portion of traffic isolated from all other traffic data.

> **Notes:**
>
> * On Interfaces pages, the **Inbound/Outbound** tabs are replaced with **Ingress/Egress** tabs (see **Traffic Graph Tabs**).
> * The tabs displayed on the graph depend on the traffic data available (for example, if there is no Internal traffic in the data specified with the **Parameter Controls**, the **Internal** tab will not be displayed).

### Traffic Graph Actions

The following actions will change the display of the graph:

* **Change tab**: The default tab shown is **Total**. To view a different tab, click the tab you’d like to view (see **Traffic Graph UI**).
* **Solo a plot**: In the legend below the graph, hover over the label for one plot. The other plots will be muted.

  + **Example**: On the IP Addresses aggregate page, in the legend below the graph on the **Inboun**d tab, hover over **Total**. Everything on the graph except the Total line will be muted, making it easier to see just the Total traffic.

    > **Note**: Not all graphs displayed on a Core details page display a legend under the graph. This feature is only available for those that do.
* **Hide a plot**: Hide the display of an individual area plotted on the graph, by clicking on the corresponding label or color swatch in the legend. Click again to toggle the plot back on.

  + **Example**: On the IP Addresses aggregate page, in the legend below the graph on the **Inboun**d tab, click **Total**. The area for Total will be hidden, and the graph will redraw to adjust to the new scale in the absence of that area.

    > **Note**: Not all graphs displayed on a Core details page display a legend under the graph. This feature is only available for those that do.
* **Display point-in-time values**: Display values for all plots at a specific point in the time range by hovering over the graph at that point. The resulting popup will display the date, time, and traffic data values.

  + **Example**: Hover at any point within the graph on the **Total** tab. The popup will show the volumes of the total, inbound, outbound and all other available traffic tabs at that point in the time range.

## Traffic Table

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(403).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A46Z&se=2026-05-12T09%3A41%3A46Z&sr=c&sp=r&sig=PmZTqf4yz%2Bfxw1w3yN9c%2BNCzrwf6JeQ0bNgDoXI0OsI%3D)

Like the **Traffic Graph**, the **Traffic** table presents the results returned from the query specified with the **Parameter Controls** (traffic over the current time range, expressed in the current metric, and calculated with the current aggregation). Rather than a visualization, however, the data is presented as a top-X table. The data shown in the rows of the table vary depending on the type of the current page:

* **Aggregate page table**: If the table is on an aggregate page, each row corresponds to an individual instance of the network entity whose traffic data is being aggregated on the page. The rows are ordered according to traffic volume (top-X).  
  **Example**: On the Devices aggregate page, each Kentik-registered device in your network is represented as a row in the table.
* **Details page table**: If the table is on a details page (or on the **Traffic Overview** page), the table is tabbed. Each tab corresponds to a dimension by which traffic data can be ranked (top-X). A **Customize** button (see **Customize Tabs Dialog**) at the top right of the table allows you to choose the dimensions for which a tab will be displayed (up to 12). The table columns shown for each tab vary depending on the dimension.

  > **Note*:*** For a list of dimensions for which traffic data can be displayed in a traffic table, see **Dimensions Reference**.

### Traffic Table UI

The traffic table includes the following UI elements:

* **Tabs** (details pages only): Up to 12 tabs, each of which contains a table for one dimension. The current tab is indicated in blue text and with a blue underline. To view the table for a different dimension, click the name of that dimension.
* **Customize button** (details pages only): Opens the **Customize Tabs Dialog** (click anywhere outside the popup to close it).

  > **Note:**For information about a specific dimension that is listed in the dialog as being available for traffic tables, see **Dimensions Reference**.
* **Dimension table**: Each tab contains a table for one dimension. What each table row represents varies depending on the type of table:

  + **Aggregate page** (untabbed table): The rows always correspond to the subject of the page itself. On the Devices aggregate page, for example, each row corresponds to one device, with the left-most column listing the names of those devices and the other columns providing information about those devices.
  + **Details page** (tabbed table): The rows correspond to the dimension of the current tab. On the details page for an individual Device, for example, the table in the Applications tab lists the applications with traffic on that device. Each row corresponds to one application, with the left-most column listing the names of those applications and the other columns providing information about those applications.
* **Table columns**: The columns in a table vary depending on the type of table:

  + **Aggregate page** (untabbed table): The columns include the **Common Table Columns** as well as additional columns providing relevant information, which vary depending on the page (the Devices aggregate page has different additional columns than the Sites or Applications aggregate pages).
  + **Details page** (tabbed table): The table on a given tab will include the **Common Table Columns**.

    > **Note:**On Interface details pages, the Inbound/Outbound columns are replaced with Ingress/Egress.
* **Show Top 50 Results**: This link (bottom left of table) expands the table to show rows for the top 50 results (based on total volume of traffic). The link appears only when there are more than 50 results to display.

### Common Table Columns

The following columns are common to the **Traffic** table on both aggregate pages and details pages (clicking on a column heading sorts the list, ascending or descending):

* **Dimension** (actual column heading varies by tab/page): The name of the dimension/instance being viewed (see **Dimension Column**).
* **Internal**: Traffic whose origin and destination are both within your network.
* **Inbound**: Traffic entering from somewhere outside your network.
* **Outbound**: Traffic leaving to somewhere outside your network.
* **Ingress** (Interfaces aggregate page only): Traffic coming into a device.
* **Egress** (Interfaces aggregate page only): Traffic going out of a device.
* **Through**: Traffic that both originates and terminates outside your network.
* **Other**: Traffic that is not classified as one of the above types of traffic (see **Traffic Classified as Other**).
* **Total**: The sum of Internal, Inbound, and Outbound traffic.
* **View in Explorer** (vertical ellipses): Opens **Data Explorer** for further exploration of the traffic. The Data Explorer controls will be set to query about the traffic in this row of the table. Several options are available:

  + Open ‘Internal’ in Data Explorer
  + Open ‘Inbound’ in Data Explorer
  + Open ‘Ingress’ in Data Explorer (Interfaces only)
  + Open ‘Outbound’ in Data Explorer
  + Open ‘Egress’ in Data Explorer (Interfaces only)
  + Open ‘Through’ in Data Explorer
  + Open ‘Other’ in Data Explorer

> **Notes:**
>
> * The Total column is not shown on aggregate pages for Sites, Devices, or Interfaces.
> * On the Interfaces tabs/pages, the only traffic columns shown are the Ingress and Egress columns.
> * The units displayed in the table depend on the setting of the **Metric Selector**.

#### Dimension Column

While the actual column heading varies by tab and by page, the dimension column is the name of the dimension or instance being viewed. For aggregation pages, the column heading matches the page (on the Applications aggregate page, the column name would be "Application"), while for details pages, the column heading would match the selected tab (on the **Cities** tab, the column’s heading would be "City"). Depending on the tab or page selected, the following items may appear beside items listed in this column:

* **ASN lozenge** (only on ASNs and Next-hop ASNs tabs/pages): Beside each individual AS in the table is a lozenge displaying the AS number. Click it to see the **ASN Links Popup**, which gives you access to that ASN’s registration country, and links to the ASN’s KMI details, its PeeringDB information, and its traffic data.
* **AS Group lozenge** (only on ASNs and Next-hop ASNs tabs/pages): Hover over an AS Group lozenge to see a popup displaying all of the ASNs within that group. On the popup, click any ASN lozenge to see its associated **ASN Links Popup**.
* **Breakdown icon** (only on BGP Paths tabs/pages and AS Paths columns): Click the numbered list icon to the right of the path to open a popup showing additional hops in the path (next, 2nd, 3rd, etc.). Click any of the links to go to that AS’s details page.

### Customize Tabs Dialog

The popup Customize Tabs dialog enables you to configure the display of tabs in the **Traffic Table** on details pages. The dimensions available for tabs are a subset of the dimensions available for group-by and filtering in **Data Explorer**. For a given Kentik customer, these dimensions will be the same in all instances of the dialog, but they may vary over time depending on what types of devices you currently have registered with Kentik.  
  
To access the Customize Tabs dialog, click **Customize** just above the table (on the right). The dialog includes the following UI elements:

* **Choose tabs**: Each checkbox to the left of a dimension name determines if that tab is displayed (checked) or hidden (unchecked). Select up to 12 dimensions for which you'd like to display tabs.
* **Order tabs**: Handles to the left of the checkboxes allow you to click and drag the dimensions into the desired tab order.

Once you’ve chosen the tabs to include, click outside the dialog to close it. Your selections will be saved for future visits to any Core details pages of the same type (for example, if you select specific dimensions and order them in a certain way for an Application details page, all other Application details pages will display with those tabs ordered the same way, but Device details pages will be unaffected).

> **Note:** For information about a specific dimension that is listed as being available for traffic tables, see **Dimensions Reference**.
