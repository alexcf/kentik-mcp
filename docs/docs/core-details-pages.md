# Device Details Pages

Source: https://kb.kentik.com/docs/core-details-pages

---

This article covers the Details pages in the **Network Explorer** module of the Kentik portal.

> **Note:**For Aggregate pages in the Network Explorer module, see **Core Aggregate Pages**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(419).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A08Z&se=2026-05-12T09%3A56%3A08Z&sr=c&sp=r&sig=mJVMoin%2FhYgYBxbSpXw8lAX1mvPiq04SNFqoSfnN3sk%3D)

*The Details page for a site, showing the Traffic tab.*

## About Details Pages

Details pages are part of the **Network Explorer** module of the portal. While **Core Aggregate Pages** present all traffic related to a given data source or dimension (e.g., all devices or all ASNs), Details pages present information for an individual entity (e.g. one device or one ASN).

#### Accessing Details Pages

There are a few different ways to drill down to a Details page:

* **Directly from Network Explorer**: In the tabbed **Traffic Table** at the bottom of the Network Explorer page (on the main menu, select **Network Explorer** under **Core**), the left column corresponds to the currently visible tab. The values in that column are each links that take you to the Details page for an individual instance.  
  **Example**: In the **Application**s tab, click on a value in the Application column to go to the Details page for an application.
* **From an Aggregate page**: From Network Explorer, click the **Explore Top Talkers** button in the **Core** card at the top of the right sidebar, then pick an Aggregate page from the drop-down menu. In the traffic table at the bottom of the resulting view, click on the name of the individual instance you’d like to see.  
  **Example**: On the Applications Aggregate page, each row of the table lists an individual application. Click the application's name in the table to go to its Application Details page.
* **From a different module**: In some cases, the results from modules outside of Network Explorer will include links to Details pages within Network Explorer:  
  **Example**: On the **Discover Peers** page (in the Edge section of the portal), the **Potential Peers** table includes an AS Name column. Click a value in this column to go to the Details page for that ASN.

## Details Information Types

All Details pages include traffic information and may — depending on the type of network entity (e.g., device or dimension) — also include one or more additional types of information. If more than one information type is included, each type is presented in its own tab.  
  
 The following information types may be included on a Details page:

* **Traffic**: While many elements of the Traffic view are common to all types of network entities (see **Common Traffic UI**), some elements are displayed only for particular types of entities (see **View-specific Traffic UI**).
* **More Info**: The Details pages for certain network entities (Devices, IPs, Interfaces, Sites, and AS groups) include a **More Info Tab** that provides additional details about the network entity.
* **Metrics**: The **Metrics Tab**, shown only for interfaces, displays five visualizations (Utilization, Bits/s, Packets/s, etc.) if SNMP polling is enabled for the interface, and an additional five if Streaming Telemetry is also enabled.
* **Capacity**: The **Capacity Tab** is shown only for interfaces. It uses a box and whisker plot to display the interface’s minimum and maximum inbound and outbound bitrates.
* **PeeringDB Info**: The **PeeringDB Info Tab**, shown only for ASNs, is a PeeringDB integration that provides details and visualizations of an ASN’s network infrastructure.

## Details Page UI

The UI elements on a Core details page are covered in the following topics.

### Details Subnav

The subnav of Core details pages is a silver-gray strip under the main navbar that includes the following UI elements:

* **Breadcrumbs**: An indicator of your current location within the Kentik portal. The last segment represents the entity (e.g., an individual AS) whose details page you are currently viewing.

  > **Note:** If a down-pointing triangle shows at the right of the segment, click it to drop down a filterable list of other entities of the same type (e.g., other ASes if you're on an AS page). Click on an entity in the list to go to the details page for that entity.
* **My Peering Details** (present only for ASes): A button that opens the **My Peering Details Dialog**, from which you can generate — and copy to the clipboard — a report on the peering details of an ASN (see **PeeringDB Info Overview**).
* **Share**: Opens the Share dialog to enable you to share the current view (see **Sharing via the Share Dialog**).
* **Actions**: A drop-down menu from which you can choose to export the current view as a visual report (PDF) covering the page’s visualizations and tables. A notification appears when the PDF is ready to download.
* **Insights**: A button that pops up the **Insights Pane**, which shows insights related to the traffic detailed on this page (e.g. on an Applications Details page, traffic related to that application; see **About Insights**).

### Details Display Area

The main display area of a Core details page may, depending on the type of network entity it represents, include the following UI elements:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(420).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A08Z&se=2026-05-12T09%3A56%3A08Z&sr=c&sp=r&sig=mJVMoin%2FhYgYBxbSpXw8lAX1mvPiq04SNFqoSfnN3sk%3D)

  *The top-level controls and information fields for a Details page, including the tabs for interface views.*

  **Name**: The name of the individual instance about which this view provides information (e.g., “https” for a Details page of the https application).
* **ASN** (ASN only): Located to the right of the ASN’s name, the lozenge displays the ASN of this AS. Click it to see the AS’s registration country and access a link to this AS’s **Kentik Market Intelligence** page.
* **AS Group** (AS Groups only): Located to the right of the AS Group’s name, the lozenge identifies this page as an AS Group Details page (see **AS Groups**).
* **Device labels** (Devices only): Labels that have been applied to the device (see **Labels**) display to the right of the Device name.
* **Classification Indicator** (Interfaces only): Located to the right of the interface name, this indicates the Interface Classification status of this interface (see **Interface Classification**), either Classified or Not Classified.
* **Description** (Device, Site, and Interface only): Located just below the entity’s name, the contents of the description depend on the type of entity:

  + Device: The description text entered when the device was registered with Kentik or subsequently edited.
  + Interface: The interface description as either defined in the device and retrieved via SNMP or specified manually. Capped at 128 characters.
  + Site: The street address of the site.
* **Configure** (Device, Interface, and Site Details pages only): A button that opens the configuration dialog for the individual device (see **Network Devices**), interface (see **Interfaces**), or site (**Site Settings**) where you can edit its settings.
* **Tabs**: A set of tabs (determined by the type of entity), listed in **Details Information Types**, that present different views related to this Details page's individual network entity (e.g., device or dimension).

  > **Note:**All Details pages include a **Traffic** view. If the page includes a **More Info** tab then the **Traffic** view is also presented as a tab.

## Traffic Tab

The various areas of the **Traffic** tab are covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(421).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A08Z&se=2026-05-12T09%3A56%3A08Z&sr=c&sp=r&sig=mJVMoin%2FhYgYBxbSpXw8lAX1mvPiq04SNFqoSfnN3sk%3D)

*The controls and chart of the Traffic tab.*

> **Note:** If the Details page is of a type that has no other tabs, then the Traffic information is not presented within a tab.

### Common Traffic UI

The following main UI components are common to the **Traffic** tab of all Details pages in Network Explorer:

* **Parameter controls**: Sets the parameters of the query (Filters, Aggregate, Metric, and Time Range) whose results are displayed in the page's graph and table (see **Traffic Parameter Controls**).
* **Traffic graph**: Displays visualizations of the results returned from the query specified with the **Parameter** **Controls** (see **Traffic Tab Graph**).
* **Traffic table**: A top-X table presenting the results returned from the query specified with the **Parameter** **Controls**. The group-by dimension for the query is the dimension corresponding to the current page (e.g., when the page is Applications, the dimension is applications). The columns of the table are listed in **Traffic Tab Table**.

#### Traffic Parameter Controls

The parameter controls, located below the title bar, set the parameters of the query whose results are displayed in the page’s graph and table (see **Parameter Controls**):

* **Filters control**: Apply filters that narrow the traffic covered by the query results.
* **Aggregate selector**: Choose the calculation by which traffic data displayed in the graph and table is aggregated.
* **Metric selector**: Set the metric by which query results are evaluated for top-X and displayed in the graph and table.
* **Time Range selector**: Set the duration, looking back from the current time, covered by the query whose results are displayed in the graph and table.

### View-specific Traffic UI

The following additional UI elements of the **Traffic** tab are found only on specific Details pages:

* **Engineer Traffic** (Interfaces only): A button on the **Traffic** tab that opens the **Engineer Interface Traffic** page for an individual interface.
* **Logical Map** (Providers only): A Logical Map displays below the **Traffic** graph and above the **Traffic** table (see **Logical Map Pane**).

### Traffic Tab Graph

The graph in the **Traffic** tab displays visualizations of the results returned from the query specified with the **Traffic Parameter Controls**. Depending on your actual traffic volume, there will be a tab for each of the following types of traffic:

* **Total**: All traffic regardless of profile.
* **Internal**: Traffic that both originates and terminates inside your network.
* **Inbound**: Traffic that originates outside your network and terminates inside your network.
* **Outbound**: Traffic that originates inside your network and terminates outside your network.
* **Through**: Traffic that both originates and terminates outside your network.
* **Other**: Traffic that is not classified by the above profiles.
* **Ingress** (Interface views only): Traffic that originates outside your network.
* **Egress** (Interface views only): Traffic that originates inside your network.
* **View in Data Explorer**: A link that takes you to Data Explorer, where the Query controls will be set to show the traffic displayed in the Traffic tab (see **Data Explorer**).

> **Notes:**
>
> * For a full description, see **Traffic Graph**.
> * For an explanation of how traffic is assigned to the above categories, see **Simple Traffic Profile**.

### Traffic Tab Table

The table in the **Traffic** tab is tabbed, with each tab providing information related to a different dimension. The currently displayed tabs (up to 12) are specified with the **Customize Tabs Dialog**.

* For information about the UI of the traffic table (also called a tabbed table), see **Traffic Table UI**.
* For information about common traffic table columns, see **Common Table Columns**.
* For information about a specific dimension that is listed as being available for traffic tables, see **Dimensions Reference**.

> **Notes:**
>
> * The dimensions available for tabs are a subset of the dimensions available for group-by and filtering in **Data Explorer**.
> * For a given Kentik customer, the list of dimensions will be the same in all instances of the Customize Tabs dialog but may vary over time depending on what types of devices you currently have registered with Kentik.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(422).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A08Z&se=2026-05-12T09%3A56%3A08Z&sr=c&sp=r&sig=mJVMoin%2FhYgYBxbSpXw8lAX1mvPiq04SNFqoSfnN3sk%3D)

*The traffic table gives you a breakdown of the profile of traffic associated with a given network entity.*

### Logical Map Pane

The **Logical Map** pane, which appears only on Details pages whose type is Provider, is located between the **Traffic Tab Graph** and the **Traffic Tab Table**. The contents of the pane are essentially the same as in the portal’s **Logical Map** module, except that the pane does not have a settings sidebar. An "open in new" icon at the upper left of the pane will open the logical map for this provider in the module.

## More Info Tab

On Device Details pages, the **More Info** tab displays a wide variety of information about the device, which is covered in the topics below.

#### Device General Details

General information about the device:

* **Manufacturer**: If SNMP polling is enabled and the device has a value for the `ifDescr` OID (see **Information SNMP OIDs**), the value is displayed in this field.
* **Description**: User-provided descriptive text entered in the **Description** field of the device's configuration dialog. The field is not shown if no description has been provided.
* **Site**: The site to which this device is assigned (see **Sites**).
* **Plan**: The billing plan to which the device is assigned (see **Licenses**).
* **Kproxy Agents**: The kproxy agents, if any, via which this device sends flow data to Kentik (see **Kentik Proxy Agent**).
* **Sending IPs**: IP address(es) from which the router sends flow.
* **ID**: The device’s Kentik-assigned device ID.
* **Created**: The date that the device was added.

#### Device Flow Details

Details related to the flow records sent from the device to Kentik (see **Flow Overview**):

* **Flow**: Indicates whether Kentik is receiving flow from the device. The indicator has the following states:

  + No Flow Detected (red): Displayed when no flow is detected from the device.
  + Direct Flow Detected (green): Flow is coming directly from a router.
  + Kentik Agent Detected (green): Flow is coming via Kentik's kproxy software agent (see **Kentik Proxy Agent**).
* **FPS (Sampled)**: The rate at which Kentik has been ingesting flow records from the device over the last six hours. This value is after any downsampling Kentik performs pursuant to the terms of the plan to which the device is assigned.
* **FPS (Raw)**: The rate at which the device has been generating flow records over the last six hours (before any Kentik downsampling).
* **Plan FPS**: The maximum FPS supported by the plan to which this device is assigned.
* **Flow Type**: The format of the flow data (e.g., NetFlow v5, v9, IPFIX, or sFlow). See **Flow Protocols**.
* **Sample Rate**: The ratio of total flows per sampled flow (see **Flow Sampling**).

#### Device Interface Details

Details related to the interfaces on the device:

* **Interfaces Classified**: The number of classified interfaces, the total number of interfaces, and the percent of interfaces that have been classified (see **About Interface Classification**).
* **View Interfaces**: A button that takes you to the **Interfaces Page** where you can manage the interfaces on your devices.

#### Device SNMP Details

Details related to SNMP polling of the device by Kentik (see **SNMP OID Polling**):

* **SNMP Status**: Current status of SNMP polling (e.g., Detected or Not Detected).
* **SNMP Polling** (see **SNMP Polling Intervals**):

  + If Standard, SNMP polling is every 10 minutes for interface counters and three hours for descriptions.
  + If Minimal, interface polling is disabled and descriptions are polled every six hours.
* **SNMP V3 Auth**: Indicates whether V3 authentication for SNMP polling is on or off (see **About SNMP V3**).

#### Device BGP Details

Details related to BGP on the device (see **BGP Overview**):

* **BGP status**: Current status of BGP peering with the device (e.g., Established, Not Established, or Not Configured).
* **BGP Prefixes**: For devices with BGP enabled, this indicator shows the number of prefixes currently (within last 5 minutes) in the BGP routing table.

#### Device Streaming Telemetry

Details related to Streaming Telemetry (ST) on the device (see **Device ST Details**).

* **Streaming Telemetry status**: Current status of ST for the device (e.g. Disabled, No ST detected, or ST Detected).

#### Device Features

Additional Kentik capabilities enabled on this device (display an **X** if disabled):

* **CDN attribution**: Enabled if the device is a host and its Contribute to CDN Attribution switch is on in the **General** tab of the Device Settings dialog (see **Network Devices**).
* **Flowspec**: Enabled if the device's **BGP Flowspec Compatible** switch is on in the **BGP** tab of the Device Settings dialog.
* **RTBH**: Enabled if the device is assigned to an RTBH mitigation platform (see **Configure Platform Devices**).

### Interface More Info Tab

The **More Info** tab on Details pages for interfaces includes the following information:

* **Metadata**: SNMP-related interface attributes derived by polling the ifMib (SNMP global interface data model). See **Interface Metadata**.
* **Status**: The interface’s statuses:

  + Admin Status: The configured status of the interface (see **ifAdminStatus**).
  + Operating Status: The interface’s operational status (see **ifOperStatus**).
* **LLDP Neighbors**: This information is only present when the associated key:value store includes `lldpRemPortId` (see **Interface LLDP Neighbors**).
* **Classification**: Information about the interface’s status, network boundary, connectivity type, and provider (see **More Info Classification**).

#### Interface Metadata

SNMP-related interface attributes are derived by polling the ifMib (SNMP global interface data model). The following metadata is available on the **More Info** tab for Interface Details pages:

* **Device**: The name of the device to which this interface belongs.
* **Site**: The site location of this device.
* **SNMP ID**: The SNMP ID (polled or manually specified) of this interface.
* **Lower Sublayer Interface**: If the interface uses a parent/child relationship, and this interface is a child, this field shows the parent value. Otherwise, the field is not present.
* **Parent Interface**: This displays if the interface uses a parent/child relationship, this interface is a parent, and will show the value(s) of any children that may exist.
* **Alias**: The interface’s SNMP description (alias).
* **Type**: The underlying technology of the interface (see **Interface Types**)
* **Speed**: The network speed in Mbps.
* **MTU**: The interface’s maximum transmission unit, which is the largest packet size, in bytes, that the interface will accept.
* **Physical Address**: The MAC address of the interface.
* **Connector Present**: If the interface sublayer has a physical connector (port), the value “physical” is displayed. Otherwise, it displays “logical”.
* **IP**: The primary IP address of this interface.
* **Additional IPs**: Additional IP address(es) assigned to this interface (if applicable).
* **VRF**: The VRF that the interface belongs to. If it doesn’t belong to a VRF, it displays “None.”
* **Boundary ASNs**: The ASNs of the autonomous systems to which — so far as Kentik is able to determine based on traffic and BGP data — an edge (External) interface is connected. If there's more than one AS, the percentage of traffic for each is also indicated.
* **Capacity** (present only if the interface is monitored by a capacity plan): The maximum capacity in Mbps as reported by SNMP.

> **Note**: Any values that show as “Overridden” have been changed by the user in the **Interface Settings Dialog** for that particular interface (see **Interfaces Page**).

#### Interface LLDP Neighbors

The **LLDP Neighbors** section on the **More Info** tab for Interface Details pages is only visible when the associated key:value store includes `lldpRemPortId` (see https://oidref.com/1.0.8802.1.1.2.1.4.1.1). The following pieces of information display:

* **Remote SysName**: The remote system’s name.
* **Remote Port ID**: The port component associated with the remote system.
* **Remote Port ID SubType**: The port identifier type used.
* **Remote Port Description**: The port associated with the remote system.

#### More Info Classification

The **Classification** section of the **More Info** tab for an Interface Details page contains the following pieces of information:

* **Status**: This indicates the Interface Classification status of this interface (see **Interface Classification**), either Classified or Not Classified.
* **Network Boundary**: The network boundary value assigned to the interface by interface classification (see **Network Boundary Attribute**).
* **Connectivity Type**: The network boundary value assigned to the interface by interface classification (see **Connectivity Type Attribute**).
* **Provider**: The provider value assigned to the interface by provider classification (see **About Provider Classification**).

### IP More Info Tab

The **More Info** tab on Details pages for IP addresses includes the following information:

* **AS Details**: The name and number of the AS that contains the IP address (click to access that ASN Details page).
* **Location**: The physical location of the registered owner of the AS that contains the IP address.
* **CDNs**: Indicates whether or not the IP is contained in an AS that is a CDN.
* **Threat Feeds**: Indicates whether the IP appears on any of the lists that Kentik checks for known threats (see **Using Threat Feeds**).
* **Network Classification**: Indicates whether the IP address is internal to your network or external (click to access the **Network Classification** page).

### Site More Info Tab

The **More Info** tab on Site Details pages includes the following information:

* `kproxy` **Agents**: The `kproxy` agents, if any, via which flow data is sent from the devices at this site to Kentik (see **Kentik Proxy Agent**).

### AS Group More Info Tab

The **More Info** tab on an AS Group Details page includes the following information:

* **AS Group Members list**: A listing of all of the ASes within this AS Group.
* **ASN lozenge**: Beside each AS, a blue lozenge is displayed with the AS’s AS number, which, when clicked, opens the **ASN Links Popup**.

## Metrics Tab

The **Metrics** tab appears only for **Interface Details** pages. The tab contains five stacked area charts displaying visualizations of SNMP-polled data. For interfaces that have Streaming Telemetry (ST) enabled, an additional five charts, based on ST data, display beside the initial five for side-by-side comparison (see **SNMP vs. Streaming Telemetry**). Indicators at the top-right corner of each chart state traffic for the last 24 hours using In/Out averages, maximums, and last observed levels, measured in the units employed in the graph.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(423).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A08Z&se=2026-05-12T09%3A56%3A08Z&sr=c&sp=r&sig=mJVMoin%2FhYgYBxbSpXw8lAX1mvPiq04SNFqoSfnN3sk%3D)

### Metrics Tab Visualizations

The following visualizations are included on the **Metrics** tab:

* **Utilization**: The utilization of the interface statecd as a percentage. The left vertical axis of the graph measures the interface’s input utilization, while the right vertical axis measures the interface’s output utilization. Hover over any part of the graph to see the values of each for any point of the last 24 hours (measured in 10 minute intervals).

  > **Notes**:
  >
  > + The **SNMP** visualization displays by default (on the left); Streaming Telemetry only displays if that is enabled for the interface.
  > + On the **SNMP** graph, click **View in Data Explorer** to see the Total by Average Utilization In, Average Utilization Out query displayed in Data Explorer.
  > + The **Streaming Telemetry** visualization (on the right) is the only place to see this information displayed in graph form and, unlike the SNMP graph, does not currently link to a more detailed visualization in Data Explorer.
* **Bits/s**: The interface’s input and output bit rate, measured in bits/s. The left vertical axis of the graph measures the interface’s input in bits/s and the right vertical axis measures the interface’s output in bits/s. Hover over any part of the graph to see the values of each for any point of the last 24 hours (measured in 10 minute intervals). Click **View in Data Explorer** to see the Total by Average Bits/s In, Average Bits/s Out query displayed (for either SMNP or Streaming Telemetry) in Data Explorer.
* **Packets/s**: The interface’s input and output packet rate. The left vertical axis of the graph measures the number of packets/s in, while the right vertical axis measures the number of packets/s out. Hover over any part of the graph to see the values of each for any point of the last 24 hours (measured in 10 minute intervals). Click **View in Data Explorer** to see the **Total by Average Packets/s In**, **Average Packets/s Out** query displayed (for either SMNP or Streaming Telemetry) in Data Explorer.
* **Discards/s**: The interface’s input and output packets discard rate. The left vertical axis of the graph measures the number of input packets discarded, while the right vertical axis of the graph measures the number output packets discarded. Hover over any part of the graph to see the values of each for any point of the last 24 hours (measured in 10 minute intervals). Click **View in Data Explorer** to see the **Total by Average Input Discards**, **Average Output Discards** query displayed (for either SMNP or Streaming Telemetry) in Data Explorer.
* **Errors/s**: The interface’s input and output error rates. The left vertical axis measures the number of input errors for the interface, while the right vertical axis measures the number of output errors. Hover over any part of the graph to see the values of each for any point of the last 24 hours (measured in 10 minute intervals). Click **View in Data Explorer** to see the **Total by Average Input Errors**, **Average Output Errors** query displayed (for either SMNP or Streaming Telemetry) in Data Explorer.

## Capacity Tab

The **Capacity** tab appears only on **Interface Details** pages. The box and whisker plot shows the minimum and maximum bitrates for both the inbound and outbound traffic. The chart only appears for interfaces monitored by a capacity.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(424).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A08Z&se=2026-05-12T09%3A56%3A08Z&sr=c&sp=r&sig=mJVMoin%2FhYgYBxbSpXw8lAX1mvPiq04SNFqoSfnN3sk%3D)

### Capacity Tab UI

The **Capacity** tab contains the following UI elements:

* **Capacity Plan selector**: A drop-down that lists the capacity plans that monitor this interface. Select a plan to see its box and whisker plot chart. The lozenges beside the capacity plan name are the thresholds selected for that capacity plan (see **Capacity Plan Thresholds**).
* **Box and whisker chart**: This logarithmic graph shows the inbound and outbound bitrate of the selected interface, measured by either weeks or months (see **Box and Whisker Chart**).
* **Point-in-time values**: Displays values for each bitrate at a specific point in the time range when you hover over a point on the graph. The resulting popup will display the P95th (95th percentile), P75th, P50th, P25th, and P5th for both the inbound and outbound bitrates.
* **View Capacity Plan**: This link takes you to the page for the selected capacity plan monitoring this interface.

> **Note**: If the interface is not being managed by a capacity plan, this tab will show a link to the main **Capacity Planning** page.

#### Box and Whisker Chart

A box and whisker chart is a logarithmic graph (where scale is non-proportional) that shows the inbound and outbound bitrate of the selected interface, measured by one of the following:

* Weeks if the capacity plan's **Runout** setting is Week over Week (see **Capacity Plan Thresholds**).
* Months if **Runout** is Month over Month.

The contents of the chart are as follows:

* The orange boxes and whiskers in the diagram represent inbound traffic.
* The green boxes and whiskers represent outbound traffic.
* The top whisker is the highest bitrate detected during that time period (95th percentile or P95th), and the bottom whisker is the lowest (5th percentile or P5th).
* The top of the box represents the 75th percentile (P75th) and the bottom of the box is the 25th percentile (P25th).
* The horizontal line within the box is the median (50th percentile or P50th).
* If the traffic approaches the interface’s capacity, it will display in the chart as a horizontal red line.

## PeeringDB Info Tab

The **PeeringDB** Info tab is found only on **ASN Details** pages.

### PeeringDB Info Overview

PeeringDB is an organization that provides a platform for the Autonomous Systems (ASes) that make up the Internet to declare their footprint (locations of physical and virtual network infrastructure) and to evaluate peering requests from other networks. PeeringDB has been widely adopted by the peering community and is recognized as an important tool for strategic peering planning.  
  
Kentik’s PeeringDB integration facilitates the following:

* Evaluates the details of the inbound and outbound traffic of your AS.
* Evaluates the conditions for peering with a network.
* Finds common locations between another AS and your own.
* Identifies the network operations center (NOC) and peering contacts for a network.
* Examines and analyzes network infrastructure and services shared by multiple networks (common footprint analysis) in conjunction with **Data Explorer** or **Kentik Market Intelligence**.

### PeeringDB Info Panes

The **PeeringDB** Info tab appears on **ASN Details** pages and displays a wide variety of PeeringDB information related to the AS. The tab includes the following panes:

* **Traffic Profile**: Information about the volume and connectivity type of the AS's traffic (see **Traffic Profile Pane**).
* **General Information**: Information about the AS, including network attributes and peering policy (see **General Information Pane**).
* **Visualization**: A map and/or table showing the AS's peering sites and IX exchanges (see **Visualization Pane**).

### Traffic Profile Pane

The **Traffic Profile** pane, located at the upper left of the tab, contains the following UI elements:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(425).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A08Z&se=2026-05-12T09%3A56%3A08Z&sr=c&sp=r&sig=mJVMoin%2FhYgYBxbSpXw8lAX1mvPiq04SNFqoSfnN3sk%3D)

* **View in Data Explorer**: A button (graph icon) that opens Data Explorer in a new tab, where it will display a Source Connectivity Type data visualization for the AS.
* **In : Out Ratio**: The AS's ratio of inbound traffic to outbound traffic.
* **Inbound**: A bar graph displaying how much and where inbound traffic is coming from.

  + Total Traffic: The total amount of inbound network traffic.
  + Bar graph: A breakdown of the total inbound traffic shows as percentages on the bar graph.
  + Key: Color-coded key mapping the different inbound traffic on the bar graph.
* **Outbound**: A bar graph displaying how much and where outbound traffic is going to.

  + Total Traffic: The total amount of outbound network traffic.
  + Bar graph: A breakdown of the total outbound traffic shows as percentages on the bar graph.
  + Key: Color-coded key mapping the different outbound traffic on the bar graph.
* **See KMI Rankings**: A link to the Kentik Market Intelligence page for the ASN (see **KMI Page**).

### General Information Pane

The **General Information** pane located at the bottom left of the tab contains the following UI elements:

* **Info**: ASN name and website.
* **Network Attributes**: Network attribute details such as traffic information and IP prefix information.
* **Notes**: Notes as provided by the ASN. The notes can be expanded with the expand/collapse button.
* **Peering Policy**: Peering policy as provided by the ASN.
* **Contacts**: Contact information provided by the ASN.

> **Note**: This information can be found by searching for the ASN at **PeeringDB**.

### Visualization Pane

The main display area toward the right of the tab is used for a geographic map and/or a table displaying the peering sites (facilities) and IX exchanges in which the AS has a presence, which can help you identify overlaps between this AS and your own:

* **Filter**: A set of controls to choose what points of presence appear in the visualizations (see **PeeringDB Filters**).
* **Visualization**: A drop-down to choose to display the map only, table only, or both.
* **Map**: An interactive map displaying the location of facilities and exchanges (see **PeeringDB Map**).
* **Table**: A tabbed table listing facilities and exchanges and providing information for each. The table contains the following tabs:

  + Peering Sites (Facilities): Information about facilities. See **Peering Sites Table**.
  + IX (Exchanges): Information about exchanges. See **IX (Exchanges) Table**.
  > **Note**: On the right of each tab’s label is the number of facilities or exchanges listed within each respective tab, out of the total number available for that ASN. If no filters are applied, only the total will appear.

#### PeeringDB Filters

On the **Visualization** pane, you can filter the points of presence that appear in the visualizations with the following controls:

* **Type**: A drop-down to choose either facilities, exchanges, or both.
* **Country**: A filterable drop-down to choose one or more countries in which the points exist.
* **Common**: A drop-down to choose all points, points in common with your own AS, or points not in common.
* **Search**: A field in which to enter text that must be matched by a point in order for it to be shown.

#### PeeringDB Map

The **PeeringDB** map visualization contains the following UI elements:

* **Facility** (pin marker): A map marker showing the location of a facility.
* **Exchange** (square marker): A map marker showing the location of an exchange.
* **Information Popup**: Markers display a popup with information about the facility or exchange when you hover over them (see **PeeringDB Map Popup**).

> **Note**: Open the respective **Facility Details Drawer** or **Exchange Details Drawer** for that site by clicking on its map marker.

#### PeeringDB Map Popup

Each marker on the **PeeringDB** map displays a popup (when the cursor hovers over it) with the following information about the AS's presence in that location:

* **Name**: The name of the facility or exchange.
* **Address** (facility only): The address of the facility.
* **Networks**: The number of networks located at the facility or exchange.
* **Facilities** (exchange only): The number of facilities at the exchange.
* **Exchange** (facility only): The number of exchanges at the facility.

#### Peering Sites Table

The **Peering Sites (Facilities)** table contains the following columns:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(426).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A08Z&se=2026-05-12T09%3A56%3A08Z&sr=c&sp=r&sig=mJVMoin%2FhYgYBxbSpXw8lAX1mvPiq04SNFqoSfnN3sk%3D)**Common**: Indicates whether this AS and your own are both present in the facility represented in this row of the table:

  + Blue check icon: Your AS and this AS are both present in this facility.
  + Unchecked: Your AS and this AS are not both present in this facility.
* **Facility**: The name of the facility.
* **Organization**: The name of the organization that operates the facility.
* **Address**: The address of the facility.
* **City**: The city where the facility is located.
* **State/Region**: The state/region where the facility is located.
* **Country**: The country where the facility is located.
* **Networks**: The number of ASes operating at the facility.

> **Notes**: Open the **Facility Details Drawer** by clicking on any row of the table.

#### IX (Exchanges) Table

The **IX (Exchanges)** table contains the following column headers:

* **Expand** (icon): See the facilities operating at this exchange. The IX row moves up to the top of the table and a **Peering Sites Table** opens directly beneath it, listing all of the facilities operating at this exchange (IX). Click the icon again to hide the facilities table.
* **Common**: Indicates whether you can use this IX to peer with this AS:

  + Blue checkmark: Your AS and this AS are both subscribed to this IX. If you haven't already, you can peer with the AS by setting up a BGP session through the IX.
  + Faded blue checkmark: Your AS and this AS are both at a facility that is served by this IX, and this AS is already subscribed to that IX. Peer with this AS by subscribing to the IX. Click the expand icon at the left of the IX’s row to reveal the peering sites (facilities) that your AS is at where this IX is available.
  + No checkmark: Your AS has no presence at this IX or at any of the facilities served by this IX.
* **Exchange Point**: The name of the exchange point (IX).
* **Speed**: The network capacity for peering connections.
* **IP Address**: The IP address of the exchange.
* **City**: The city the exchange is located in.
* **Country**: The country the exchange is located in.
* **Facilities**: The number of facilities operating in at the exchange.
* **Networks**: The number of ASes operating at the exchange.

> **Notes**: Open the **Exchange Details Drawer** by clicking on any row of the table.

### Facility Details Drawer

Access the **Facility Details** drawer in the following ways:

* Click on any row on the **Peering Sites** table.
* Click on any facilities marker on the **PeeringDB** map.

The drawer contains the following UI elements:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(428).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A08Z&se=2026-05-12T09%3A56%3A08Z&sr=c&sp=r&sig=mJVMoin%2FhYgYBxbSpXw8lAX1mvPiq04SNFqoSfnN3sk%3D)**Name**: The name of the facility.
* **Organization**: The organization the facility belongs to.
* **Close**: The X in the upper right corner closes the drawer.
* **Location**: The geographic location of the facility.
* **Networks**: The number of ASNs operating at the facility.
* **Exchanges**: The number of exchanges operating at the facility.
* **Filter controls**: A set of controls to filter for ASNs.

  + Policy: A drop-down to filter by peering policy constraints (see **Peering Policy**).
  + Traffic Ratio Type: A drop-down to filter by the inbound/outbound traffic ratio.
  + Search: A field to search and match the string input on any column.
* **AS List**: A list of ASes operating at this exchange (see **Facility AS List**).

#### Facility AS List

The AS list in the **Facility Details** drawer shows the ASes that are operating at the facility and contains the following columns:

* **Name**: The name of the AS.
* **ASN**: The AS number, which, when clicked, opens the **ASN Links Popup**.
* **Policy**: The peering policy constraints (see **Peering Policy**).
* **Traffic Ratios**: The inbound/outbound traffic ratio.

#### ASN Links Popup

The ASN links popup includes the following information and links for the AS whose ASN is clicked:

* **Registration Country**: The country in which the ASN is registered.
* **Kentik Market Intelligence**: A link to the **AS Details Page** for this AS in KMI.
* **PeeringDB record**: A link to the **PeeringDB Info Tab** of the Core Details page for this AS.
* **AS Traffic View**: A link to the **Traffic Tab** of the Core Details page for this AS.

### Exchange Details Drawer

Access the **Exchange Details** drawer in the following ways:

* Click on any row on the **IX** table.
* Click on any exchange marker on the **PeeringDB** map.

The drawer contains the following UI elements:

* **Name**: The name of the exchange.
* **Close**: The X in the upper right corner closes the drawer.
* **Peers**: The number of ASNs connected to the exchange.
* **Connections**: The number of connections to the exchange.
* **% with IPv6**: The percentage of ASNs connected to the exchange with Ipv6.
* **Open peers**: The number of ASNs with an open peering policy.
* **Total speed**: The current total operating speed of the exchange.
* **Facilities**: The number of facilities located at the exchange.
* **Filter controls**: A set of controls to filter for ASNs.

  + Policy: A drop-down to filter by peering policy constraints (see **Peering Policy**).
  + Traffic Ratio Type: A drop-down to filter by the inbound/outbound traffic ratio.
  + Search: A field to search and match the string input on any column.
* **AS List**: A list of ASes operating at the facility (see **Exchange AS List**).

#### Peering Policy

ASes have different approaches to establishing and managing peering relationships with other networks. A peering policy outlines the terms and conditions under which a given network is willing to peer. These policies fall into the following general categories:

* **None**: No peering policy. Any network may peer with this network.
* **Open**: Networks with an open peering policy are willing to peer with networks that meet basic technical and operation requirements.
* **Restrictive**: Networks with a restrictive peering policy impose specific conditions on peering partners. These conditions can be based on multiple factors such as traffic volume, network size, and geographical location.
* **Selective**: Networks with a selective peering policy evaluate potential peering partners based on their own specific criteria. Similar to a restrictive policy, these conditions can be based on multiple factors such as traffic volume, network size, and geographical location.

#### Exchange AS List

The AS list in the **Exchange Details** drawer shows the ASes that are operating at the exchange and contains the following columns:

* **Name**: The name of the AS.
* **ASN**: The AS number, which, when clicked, opens the **ASN Links Popup**.
* **Policy**: The peering policy constraints (see **Peering Policy**).
* **Traffic Ratios**: The inbound/outbound traffic ratio.
* **Speed**: The network capacity for peering connections.
* **IP Address**: The IP address of the ASN.

### My Peering Details Dialog

The My Peering Details dialog, opened via the button on the subnav, generates a peering information report for your AS. The dialog contains the following fields and controls:

* **Close**: The **X** in the upper right corner closes the dialog.
* **PeeringDB Record URL**: A switch that includes a URL for your AS.
* **Exchanges**: A switch that includes information about the exchanges with which you are peering. A count of those exchanges is shown to the right of the switch.
* **Facilities**: A switch that includes information about the facilities in which your AS is present. A count of those facilities is shown to the right of the switch.
* **Show Data from ASNs**: A set of switches to choose which ASes are included in the report.
* **Report Display**: A display showing the content of the report based on the settings of the switches above.
* **Copy to Clipboard**: A button that copies the report to the clipboard. The report will be copied as text in the format shown in the **Report Display**. When the button is clicked, its label will briefly change to confirm that the report has been copied.

---

© 2014-25 Kentik
