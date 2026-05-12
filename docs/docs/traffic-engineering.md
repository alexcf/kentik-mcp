# Traffic Engineering

Source: https://kb.kentik.com/docs/traffic-engineering

---

The **Traffic Engineering** workflow in the Kentik portal is covered here.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(287).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A15Z&se=2026-05-12T09%3A44%3A15Z&sr=c&sp=r&sig=VGtkkwKTq9rvk0IhuTL5ICax4c0g0jMiuelu2J4GeUk%3D)

*The Inbound tab of the Traffic Engineering landing page.*

## About Traffic Engineering

|  |  |
| --- | --- |
| **Purpose:** | Support BGP traffic engineering by identifying groups of prefixes whose traffic is impacting congested interfaces. |
| **Benefits:** | * Automate data collection and collation for traffic engineering tasks. * Set target utilization for custom-defined groups of interfaces. * Precisely measure the impact of traffic associated with any set of prefixes on an individual interface's utilization. |
| **Use Cases:** | * Assess traffic volume relative to capacity. * Avert capacity crises by taking advance action based on dynamic data and intelligent analysis. |
| **Relevant Roles:** | Network Engineer, Network Planner, Network Architect |

Traffic engineering involves using BGP policies to shift traffic from one network peer to another, which network operators do routinely to satisfy operational and business requirements. The reasons for engineering traffic typically vary depending on the network type:

* An eyeball network that has an uplink to a primary provider may also have, for redundancy, another uplink to a different provider, which is not to be used under normal circumstances but is instead intended — and budgeted — to serve only as a failover. To contain overall costs, eyeball networks try to minimize the volume of traffic over backup links (a.k.a. inbound traffic engineering).
* A transit AS may want to optimize performance by balancing the load on their routers and minimizing traffic across their backbone links. To do so, they may shift traffic within a PoP, e.g., from a peer to a customer (a.k.a. inbound and outbound traffic engineering).
* Content providers will be concerned mostly with allocating outgoing traffic to the most performant peer.

Regardless of the type of network or the intention of the operator, traffic engineering involves updating BGP policies to shift traffic between upstream or downstream interfaces. The inbound and outbound BGP UPDATE messages exchanged between a given device and its directly connected networks are intercepted by the device’s policy and can be modified to make routes appear more or less attractive prior to being installed in the device's routing table and processed by the routing engine. The main challenge is often to determine the prefixes that route maps or route policies should be updated to favor or exclude on one or more individual routers.

> **Note:** Route modification techniques such as AS Path prepending are often ignored upstream, making it much harder to influence where inbound traffic comes from than to influence where outbound traffic goes. 

### Traffic Engineering with Kentik

Kentik's Traffic Engineering workflow helps you determine how changes in traffic routing will impact utilization of a specific interface. The key controls for this process are found on the Engineer Interface Traffic page for an individual interface, where you can set a target for interface utilization (traffic volume as a percent of capacity). Based on this target, Kentik suggests traffic — grouped by prefix (default) or ASN (2nd or 3rd hop) — to route away from the interface. By deselecting these suggested traffic groups in a list you can see in a chart the effect of shifting the traffic elsewhere. Once you've decided which groups to reroute, you can export a list of those groups, which you can then use to implement your decisions in a network management tool (Kentik does not have the ability to actually change the routing of traffic).

### Traffic Engineering Scope

Kentik's implementation of traffic engineering is focused exclusively on the interfaces that handle edge traffic  rather than traffic inside your network. Kentik defines edge interfaces based on the values of the following Interface Classification dimensions (see **About Interface Classification**):

* **Network Boundary**: External (see **Network Boundary Attribute**).
* **Connectivity Type**: One of the following types (see **Understanding Connectivity Types**): Transit, IX, Free Private Peering, Paid Private Peering, Customer, Cloud Interconnect.

> **Note:** If Interface Classification on your network hasn't yet been run or is incomplete, the Traffic Engineering page won't necessarily list all of the interfaces that fit the classifications described above.

### Accessing Traffic Engineering

There are two points of entry to the **Traffic Engineering** workflow:

* **From the main menu**: In the main menu (click the **Menu** button at top left), click the **Traffic Engineering** link under Edge, which takes you to the Traffic Engineering landing page (see **Traffic Engineering Page**), where you can do one of the following:

  + Choose an interface that Kentik has identified as having had high inbound or outbound utilization in the previous 24 hours, which makes it a likely candidate for traffic engineering.
  + Search for specific interfaces by site, device, connectivity type, or name.

    > **Note:** Only edge interfaces can be accessed via the Traffic Engineering page (see **Traffic Engineering Scope**).
* **From an Interface Details page**: If an individual interface meets the criteria described in **Traffic Engineering Scope** then the detail quick view for that interface in the Core section of the portal (see **Core Details Pages**) will include an **Engineer Traffic** button at the upper right. Click the button to go to the traffic engineering page for that interface.

## Traffic Engineering Page

The layout and usage of the Traffic Engineering page is covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(290).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A15Z&se=2026-05-12T09%3A44%3A15Z&sr=c&sp=r&sig=VGtkkwKTq9rvk0IhuTL5ICax4c0g0jMiuelu2J4GeUk%3D)

*The Manual Search tab lets you find interfaces by filtering.*

### Traffic Engineering Page UI

The Traffic Engineering page enables you to find the interfaces whose traffic you'd like to engineer. Only edge interfaces can be accessed via this page (see **Traffic Engineering Scope**). The page includes the following tabs, each of which includes an **Interfaces** table:

* **Highest Inbound Utilization**: Your organization's edge interfaces, listed in descending order of inbound utilization (SNMP-derive inbound traffic volume as a percent of interface capacity). The **Filter** field (magnifying glass icon) narrows the list to rows containing the entered string in one of the table's columns (see **Interfaces Table Columns**).
* **Highest Outbound Utilization**: Your organization's edge interfaces, listed in descending order of outbound utilization. The **Filter** field (magnifying glass icon) narrows the list to rows containing the entered string in one of the table's columns.
* **Manual Search**: Find edge interfaces that match specified filter criteria:

  + Site: The name of the site in which the interface is located (see **About Sites**).
  + Device: The name of the device to which the interface belongs.
  + Connectivity type: The type of connection (transit, IX peering, etc.; see **Understanding Connectivity Types**).
  + Search: Enter a string in the field to find all interfaces that contain the string in the value of one of the columns in the **Interfaces** table.

### Interfaces Table Columns

The **Interfaces** table on the tabs of the Traffic Engineering page includes the following columns:

* **Interface**: The name and description of the interface.
* **Device**: The name of the device to which the interface belongs.
* **Site**: The name of the site in which the interface is located (see **About Sites**).
* **Connectivity Type**: The type of connection (see **Understanding Connectivity Types**).
* **Provider**: The provider (e.g., ISP) via which source/destination traffic over this interface reaches the Internet (see **Interface Classification Dimensions**).
* **95th Percentile Bit Rate** (not shown on Manual Search tab): The 95th percentile bitrate over the last 24 hours.
* **95th Percentile Utilization** (not shown on Manual Search tab): The 95th percentile of interface capacity, in percent, over the last 24 hours.
* **Engineer Traffic**: Buttons that take you to the Engineer Interface Traffic page (inbound or outbound) for an individual interface.

## Engineer Interface Traffic

The traffic engineering pages for individual interfaces, which are accessed as described in **Accessing Traffic Engineering**, are covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(291).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A15Z&se=2026-05-12T09%3A44%3A15Z&sr=c&sp=r&sig=VGtkkwKTq9rvk0IhuTL5ICax4c0g0jMiuelu2J4GeUk%3D)

*The Engineer Interface Traffic page identifies traffic to shift from the interface and graphs*

*the effect of doing so.*

### Engineer Interface Traffic UI

The individual interface pages in Traffic Engineering each include the following UI elements:

* **Interface information**: Information about the interface (see **Interface Information Group**).
* **Utilization chart**: A chart showing interface capacity, utilization, and traffic volume (see **Interface Utilization Chart**).
* **Traffic Group List**: A table listing, in descending order of traffic volume, the groups resulting from the current traffic grouping setting (see **Traffic Group List**).
* **Interface Utilization Settings**: The factors that determine how traffic over the interface is characterized for display in the **Interface Utilization Chart** and the **Traffic Group List**, as well as what traffic is suggested for rerouting away from the interface (see **Traffic Routing Options**).
* **Routing Options UI**: UI, in the right sidebar, for the routing changes suggested by Kentik to meet your utilization target for the interface (see **Traffic Routing Options**).

### Interface Information Group

This set of indicators provides the following information about the interface:

* **Interface name**: The interface name as either defined in the device and retrieved via SNMP or entered manually when adding the interface.
* **Interface description**: The interface description as either defined in the device and retrieved via SNMP or entered manually when adding the interface.
* **Interface details**:

  + Classification: A concatenation of the interface's Network Boundary and Connectivity Type.
  + Provider: The value of the interface's Provider dimension.
  + Traffic direction: Inbound or outbound.

### Interface Utilization Chart

This stacked area chart shows the following (in bps):

* **Capacity** (solid red line): The capacity of the interface.
* **Target** (dashed orange line): The current setting of the Target Utilization control.
* **Total volume** (solid blue line): The total traffic volume on the interface.
* **Group volumes** (areas): Each filled area (colored band) represents the traffic volume attributable to an individual group listed in the **Traffic Group List** below (e.g., to a prefix if the **Group Traffic By** setting is Raw Prefix).

### Traffic Group List

This table (below the utilization chart) lists traffic groups by volume (high to low). The composition of the groups varies depending on the mode selected with the **Group****Traffic****By** radio button (outbound traffic only; for inbound the mode is always Raw Prefix). To see how traffic over the time range would have been different if various groups had been rerouted away from this interface, use the checkboxes to unselect one or more groups.  
  
 The table includes the following columns:

* **Group by**: The column name follows the current **Group By** setting in **Interface Utilization Settings**.
* **AS Path** (outbound only): The BGP ASPATH for the destination IP of the flows representing the traffic in this group.
* **Destination ASN**: The origin ASN associated with the destination IP of the flows representing the traffic in this group.
* **Max Mbit/s**: The volume of the traffic in this group.

### Interface Utilization Settings

This pane is used to configure how traffic over the interface is characterized for the **Interface Utilization Chart** and the **Traffic Group List**, as well as what options are offered for rerouting traffic away from the interface (see **Traffic Routing Options**). The pane includes the following settings:

* **Target Utilization**: A field in which you can specify the target utilization as a percent of the interface's capacity. The "Target" line on the chart corresponds to this value. Default is 50%.
* **Update**: A button, active only if the target utilization value has been changed, that refreshes the charts and tables of the page based on the new value.

  > **Note:** This action will clear the **Filtered Traffic List**.
* **Group Traffic By** (outbound only): A radio button enabling you to select the method by which Kentik groups traffic for the purpose of quantifying traffic volume (see **Group Traffic By Modes**).
* **Time Range**: A selector with which you can choose the duration, looking back from the current time, over which traffic volume will be quantified for the utilization chart and table.

### Group Traffic By Modes

The **Group****Traffic****By** modes, listed below, determine how traffic is grouped for display in the **Interface Utilization Chart** and the **Traffic Group List**:

* **Raw Prefix**: The graph will show outbound traffic grouped by destination prefix and netmask length as reported by the router.
* **Dest 2nd Hop AS Path** (outbound only): The graph will show outbound traffic grouped by the first digit in the 2nd hop path element, e.g., all 2nd hops that begin with 7 will be in the same group. This mode will only return results if the router has a BGP connection enabled.
* **Dest 3rd Hop AS Path** (outbound only): The graph will show outbound traffic grouped by the first digit in the 3rd hop path element, e.g., all 3rd hops that begin with 7 will be in the same group. This mode will only return results if the router has a BGP connection enabled.

## Traffic Routing Options

The right sidebar includes UI related to Kentik suggestions for routing changes to meet your utilization target for the interface, which is covered in the following topics.

### About Traffic Routing Options

Traffic routing options are suggestions made by Kentik in response to your **Interface Utilization Settings** for the interface. These recommendations indicate how traffic over the interface could be redirected (BGP route filtered) to achieve your specified **Target****Utilization**. The recommendations appear in the**Suggested****Filters** table, which is visible only when total traffic volume (solid blue line on chart) is greater at any point in the timeline (last 24 hours) than the target (dashed orange line).  
  
Each suggestion represents a set of traffic — grouped by prefix or hop, depending on group-by mode — that could be routed over a different interface. To see the effect of a suggested BGP route filtering on the utilization of the interface, you click the filter icon at the right of the suggestion to add it to the Filtered Traffic list; the traffic represented by the suggestion will then be removed from the **Interface Utilization Chart**. Once you're satisfied with the projected result of the suggested filtering, you can copy the Filtered Traffic list and execute the route filtering in your network management system.

### Routing Options UI

The UI for exploring Kentik's suggested changes to your traffic routing is found in the right sidebar and includes the following elements:

* **Suggested Filters list**: A table listing suggested collections of groups that could be routed away from this interface in order to reduce traffic volume to the target utilization. See **Suggested Filters List**.

  > **Note:** This table is visible only when total traffic volume (solid blue line on **Interface Utilization Chart**) is greater at any point in the timeline (last 24 hours) than the target (dashed orange line).
* **Manual Prefix/Len Filtering** (present only in Raw Prefix mode): Populate the **Filtered Traffic** list using CIDR (see **Manual CIDR-based Filtering**).
* **Filtered Traffic list**: A list of groups (e.g., prefixes when **Group Traffic By** is set to Raw Prefix mode) that are candidates for rerouting traffic in order to meet your target utilization on this interface. See **Filtered Traffic List**.

#### Manual CIDR-based Filtering

The **Manual Prefix/Len Filtering** field enables you to populate the Filtered Traffic list using CIDR. Enter an aggregate prefix (e.g., 10.0.0.0/16), then click the + icon, which results in the following:

* A list will appear below the **Manual Prefix/Len Filtering** field and the prefix/len will be added to this list. A trash icon at the right of the added item will enable you to remove it from the list.
* If there's been traffic during the time range to this interface via one or more prefixes that fall within the specified range then those prefixes will be added to the **Filtered Traffic** list, which will trigger an update to the **Interface Utilization** chart.

  > **Note:** The existing prefixes in the Filtered Traffic list, if any, will be replaced by these prefixes.
* If there's been no traffic to this interface via any prefix in the specified range then the **Filtered Traffic** list will be empty (existing contents, if any, will be gone).

### Suggested Filters List

Each row in the **Suggested****Filters** table represents a collection of groups (e.g., prefixes when**Group****Traffic****By** is set to Raw Prefix mode) that could be routed away from this interface in order to reduce traffic volume to the level specified with the **Target****Utilization** setting. The table includes the following columns:

* **% of Target**: The traffic volume, as a percent of target utilization, that would have occurred over the specified time range if the traffic for this collection of groups had been routed via a different interface.  
  **Example**: If **Target Utilization** is 50% and **% of Target** is indicated as 90% then if the traffic for that collection of groups had been routed to a different interface the utilization would have been 45% (90% of 50%).
* **# Groups**: The number of groups (e.g., prefixes) in the collection of groups represented by this row in the table.
* **Filter button**: Clicking this button for a given row has the following effect:

  + Removes from the **Filtered Traffic** list any previously selected collection of groups (only one collection can be active at a time).
  + Adds the collection of groups represented by this row to the **Filtered Traffic** list.
  + Updates the **Interface Utilization Chart** to show what the effect would be of routing traffic in the **Filtered** **Traffic** list via a different interface

### Filtered Traffic List

The **Filtered Traffic** list is a list of groups (e.g., prefixes when **Group Traffic By** is set to Raw Prefix mode) representing traffic that could be routed elsewhere to bring actual utilization of the interface in line with target utilization. The traffic represented by the items in the list is removed from the traffic represented in the **Interface Utilization Chart**, so you can see what traffic utilization would have been if the traffic had been rerouted away from the interface. Once you've decided which traffic should be rerouted, the list can be copied and used to implement your decisions in a network management tool.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(321).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A15Z&se=2026-05-12T09%3A44%3A15Z&sr=c&sp=r&sig=VGtkkwKTq9rvk0IhuTL5ICax4c0g0jMiuelu2J4GeUk%3D)

*Groups added to the Filtered Traffic list are deselected in the Traffic Groups list.*

When the Engineer Interface Traffic page for a given interface is first opened the **Filtered****Traffic** list is not populated. The list can be populated in one of the following ways:

* In the **Traffic Group List**, uncheck the checkbox for a group. The group will be added to the list.
* In the **Suggested Filters List**, click the filter icon in the row for a given collection. The individual components of the previously selected collection, if there was one, will be removed from the **Filtered Traffic** list, and the individual components of the clicked collection will be added.
* In the **Prefix/Len Filtering** field (present only in Raw Prefix mode), enter an aggregate prefix (e.g., 10.0.0.0/16), then click the + icon. Every identified prefix that falls within the supplied range will be added to list.

The**Filtered****Traffic** list includes the following additional UI elements:

* **Reset**: Empties the list, returning all groups in the **Traffic Group List** to their checked state.
* **View in Explorer**: Takes you to **Data Explorer**, with the query sidebar set to show the traffic represented by the groups currently in the Filtered Traffic list.

## Using Traffic Engineering

To engineer traffic for a specific interface:

1. Navigate to the Traffic Engineering home page (Edge » **Traffic Engineering**).
2. Choose the interface from the list in one of the tabs (see **Traffic Engineering Page UI**).
3. Click the **Engineer Traffic** button for the interface, which will take you to the individual **Engineer Interface Traffic** page.
4. In the **Target Utilization** field at upper right (see **Interface Utilization Settings**), set the desired traffic volume as a percent of interface capacity.
5. If you're engineering outbound traffic, use the **Group Traffic By** radio button to set the traffic grouping mode (see **Group Traffic By Modes**).
6. Set the time range to the duration, looking back from the current time, over which you want the traffic volume quantified for the **Interface Utilization Chart** and the **Traffic Group List**.
7. Using one or more of the methods described in **Filtered Traffic List**, choose a group whose traffic you may want to redirect away from this interface.
8. Click in the **Filtered Traffic List**, select all, and copy the list to the clipboard. You can now copy the list into your traffic engineering tool to reroute the traffic.
