# NMS Devices

Source: https://kb.kentik.com/docs/nms-devices

---

The **Devices** page in the Network Monitoring System (NMS) section of the Kentik portal is covered in this article.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(459).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A33Z&se=2026-05-12T09%3A51%3A33Z&sr=c&sp=r&sig=SiyhPnUjVZVlIIPbDPMUC8Wur1mr3f9V82pNv%2B7WN50%3D)

*The NMS Devices page lists your devices being monitored by Kentik NMS.*

> **Note:** The following articles provide more on other aspects of Kentik's Network Monitoring System:
>
> * **Network Monitoring**
> * **Kentik NMS Configuration**
> * **NMS Setup**
> * **NMS Dashboard**
> * **Metrics Explorer**
> * **NMS Interfaces**
> * **Query Assistant**

## About NMS Devices

|  |  |
| --- | --- |
| **Purpose:** | See and manage all of your NMS devices. |
| **Benefits:** | * Get a centralized view of all NMS devices. * See what's important to you by filtering and customizing columns. * Speed troubleshooting by seeing key metrics. |
| **Use Cases:** | * Monitor the health and status of all your NMS devices. * Diagnose issues by analyzing device metrics. |
| **Relevant Roles:** | Network Engineer |

NMS devices contain data that is collected by Kentik’s Universal agent (see **Universal Agents**) using SNMP or Streaming Telemetry. The data can be viewed and queried in portal modules like the **Network Monitoring System** page and the **Metrics Explorer**. NMS devices are managed from the NMS Devices page.

## NMS Devices Page

In the main menu, click **Devices** under **Network Monitoring System**. The page lists all of your organization's devices that are monitored by Kentik NMS. The list is a customizable table that includes information such as status, utilization, and make and model.

### Monitoring Templates

Monitoring templates manage the data you poll, polling frequencies, and polling interfaces. Only Kentik administrators can access the Monitoring Templates page. Other capabilities are:

* Configure, apply, and change monitoring targets and intervals across multiple devices at scale.
* Prevent the polling of “admin down” interfaces, which will never be operational.
* Stop polling virtual/stub interfaces that aren’t “real”.
* Manage Metric-per-Second (“MPS”) licensing consumption (applies particularly to Streaming Telemetry).
* Control the fidelity of the data available to build graphs and other visualizations (e.g., 1-minute polling versus 5-minute polling).
* Influence the load created on devices by SNMP/ST data collection activities.

There are two ways to access the Monitoring Templates page. From the main menu, select Network Monitoring System » Devices (the breadcrumb appears as Infrastructure » Devices when the page loads) to access the Devices page.

* The first method is to click **Manage** (top right) to display a drop-down menu and select **Monitoring Templates**.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(742).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A33Z&se=2026-05-12T09%3A51%3A33Z&sr=c&sp=r&sig=SiyhPnUjVZVlIIPbDPMUC8Wur1mr3f9V82pNv%2B7WN50%3D)

  *Monitoring Templates access via the Devices page*
* The second method is to open the Device dialog from the Action menu of a particular device:

  1. Click the **Edit** icon located on the right side of the selected device row to display the Action menu.
  2. Click the **SNMP** tab.
  3. Click the **Monitoring Template** drop-down to display the menu or click the **Create Template** button at the bottom of the dialog.
  4. Click **Manage templates** to open the Settings page (see **Settings Page**).

     > **Note:** Each device can only have one monitoring template assigned at a time.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(741).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A33Z&se=2026-05-12T09%3A51%3A33Z&sr=c&sp=r&sig=SiyhPnUjVZVlIIPbDPMUC8Wur1mr3f9V82pNv%2B7WN50%3D)

*Submenu drop-down selections for Monitoring Template*

#### Settings Page

Kentik administrators can:

* View the list of preset and custom monitoring templates.
* See the number of devices and the template used for each device.
* Add monitoring templates (see **Add Monitoring Template**).
* View, copy, edit, or delete existing templates using the pencil icon.

> **Notes:**
>
> * Devices that existed prior to October 2024 will not have a template applied, and will function as though they have the “Everything” preset, but will not show up on this settings page.
> * New devices will automatically have the “Everything” template applied.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(743).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A33Z&se=2026-05-12T09%3A51%3A33Z&sr=c&sp=r&sig=SiyhPnUjVZVlIIPbDPMUC8Wur1mr3f9V82pNv%2B7WN50%3D)

*The Device Monitoring Templates page*

#### Add Monitoring Template

Enter a name for the template, add a description, and complete the interface selection to choose which interfaces will be monitored. Interfaces can be selected statically or dynamically.

![Template for adding monitoring settings, including polling intervals and SNMP options.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ND-add-monitoring-template-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A33Z&se=2026-05-12T09%3A51%3A33Z&sr=c&sp=r&sig=SiyhPnUjVZVlIIPbDPMUC8Wur1mr3f9V82pNv%2B7WN50%3D)

*The Add Monitoring Templates popup*

> **Note:** Monitoring templates also support Interface Classification (Settings > Interface Classification) configuration for dynamic interface selection.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(745).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A33Z&se=2026-05-12T09%3A51%3A33Z&sr=c&sp=r&sig=SiyhPnUjVZVlIIPbDPMUC8Wur1mr3f9V82pNv%2B7WN50%3D)

*The Interface Classification page that can be used for configuring dynamic interface selections.*

#### Advanced Measurement Selection

This workflow allows administrators to choose which metrics to collect from devices with the applied template.

Although Monitoring Templates are an **SNMP** tab setting, there are also some Streaming Telemetry settings. If the “Specify streaming telemetry intervals” option is enabled, then a new “Streaming telemetry interval” column will be available on the template.

![Monitoring template settings with various metrics and SNMP intervals for data collection.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ND-add-monitoring-template-measurements.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A33Z&se=2026-05-12T09%3A51%3A33Z&sr=c&sp=r&sig=SiyhPnUjVZVlIIPbDPMUC8Wur1mr3f9V82pNv%2B7WN50%3D)

*Measurement selection in the Monitoring Template. The “Specify streaming telemetry intervals” option is not enabled.*

![Configuration settings for monitoring template with SNMP and telemetry intervals displayed.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ND-add-monitoring-template-st-on.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A33Z&se=2026-05-12T09%3A51%3A33Z&sr=c&sp=r&sig=SiyhPnUjVZVlIIPbDPMUC8Wur1mr3f9V82pNv%2B7WN50%3D)

*The “Specify streaming telemetry intervals” option is enabled, and displays the “Streaming telemetry interval” column.*

### NMS Devices UI

The NMS Devices page includes the following main UI elements:

* **Breadcrumbs** (on subnav): Indicates your current location within the Kentik portal.
* **Actions** (on subnav): A drop-down menu for actions related to NMS devices (see **NMS Devices Actions**).
* **Manage Discoveries**: A button that takes you to the **NMS Discoveries Page**

  + View and manage the device discoveries that have been performed within your organization.
  + Each discovery involves the process described in **NMS Setup Wizard**, where you select an agent, specify an IP range, and choose devices to monitor with that agent.
* **Add Devices**: Opens the **Add NMS Devices** dialog.
* **Filters** (filter icon): Toggles visibility of the **Filters** pane. Hiding the pane does not change the currently set filters.
* **Group By**: A drop-down for choosing to group the rows of the Device list by category (see **Device List Grouping**).
* **Search**: A field that shows lozenges for the currently filters and also enables you to enter text.

  + The Device list will be filtered to show only rows that match the filters and also contain the entered text.
  + Click the **X** at the right to clear any text from the field.
* **Device List**: A table listing your organization’s NMS devices (see **NMS Device List**).
* **Filters pane**: Controls that filter the device list (see **NMS Device List Filters**).

### NMS Devices Actions

A drop-down menu that contains the following actions:

* **Remote SNMP Walks**: Opens the Remote SNMP Walks dialog, which contains the **Allow remote SNMP walks** switch.

  + Turning on this "opt-in" switch grants Kentik support permission to trigger a remote SNMP walk on a device in order to troubleshoot issues with NMS-monitored devices.

    > **Note:** This is an organization-wide setting.
* **Export**: Export the Device list as tabular data in CSV format. A notification appears when the CSV file is ready to download.
* **View in Metrics Explorer**: A link to the NMS Metrics Explorer page, where the Query sidebar shows the availability of your devices (up to 1000).

### Add NMS Devices

The Add NMS Devices dialog initiates NMS monitoring on a device using one of the following approaches:

* **Full monitoring**: Monitor device metrics including CPU, interfaces, routing protocols, etc. using SNMP, ICMP, and/or Streaming Telemetry. This option takes you to the **NMS Setup** page.
* **Ping Only**: Monitor up/down status and latency using ICMP on devices to which you have limited access. This option takes you to the **Add ICMP-only Devices Page**.

## NMS Device List

The Device list is covered in the following topics.

### About the NMS Device List

The Device list is a filterable table that shows information about the listed devices. Each row of the list represents one device and shows that device's values for each of the table's columns (see **NMS Device List Columns**). For additional details about a device, click the device's name to go to its **NMS Device Details Page**.  
  
You can change the sort order of the list by clicking on the heading of the column. Click the resulting blue up or down arrow to choose the sort direction (ascending or descending). In the bottom right corner, you can see how many devices are currently listed.

### NMS Device List Columns

The columns displayed in the Device list are chosen in a **Customize Columns** popup. A maximum of 12 columns may be visible at once. The following columns are available:

* **ID**: The device's Kentik-assigned device ID.
* **Status**: A set of lozenges indicating the status of the device:

  + **Up**: Indicates the device is intended to be operational.
  + **Down**: Indicates the device is intentionally not operational.
* **Name**: The device’s name as specified when the device was registered or edited. The name is a link to the Details page for the device (see **NMS Device Details Page**).
* **Labels**: The labels that have been applied to this device.
* **IP Address**: The IP address(es) from which the device sends data to Kentik.
* **Location**: The data center where the device is located.
* **Site**: The site where the device is located.
* **Object ID**: The globally unique identifier used in SNMP.
* **Description**: The device’s description string.
* **CPU**: The device’s CPU utilization.
* **MEM**: The device’s memory utilization.
* **Vendor**: The device’s manufacturer.
* **Model**: The device’s model number.
* **Product** **Name**: The device’s product name.
* **OS Name**: The operating system the device runs on.
* **OS Version**: The version number of device’s operating system.
* **Serial Number**: The device’s serial number.
* **Edit** (pencil icon): A button that opens a settings dialog for the device. The settings dialog for NMS-monitored devices is the same as that for non-NMS devices; see **Device Settings**.

#### Customize Columns

To choose which of the available columns (up to 12) are shown in the table, and in what order:

1. Click the **Custom** button at the upper right of the list.
2. In the resulting **Customize Columns** popup, use the checkboxes to select or deselect columns.
3. Use the handle at the left of any column to drag it up or down.
4. Click anywhere outside the popup to close it. The list will be updated to reflect your column selections.

### Device List Grouping

The **Group By** drop-down enables you to group the rows of the Device list based on one of the following categories:

* **None**: The devices are not grouped.
* **Labels**: The devices are grouped by their labels.
* **Location**: The devices are grouped by their location.
* **Model**: The devices are grouped by their manufacturer’s model.
* **Site**: The devices are grouped by their site.
* **Status**: The devices are grouped by their status.
* **Vendor**: The devices are grouped by their manufacturer.

Unless the selected category is None, a group will consist of all device rows that share a common instance of the selected grouping category (e.g., if the category is Site then all devices in the same site). The rows in each group will start with a heading row, and clicking the table's column headings will sort the rows within each group rather sorting across the entire table.

### NMS Device List Filters

The **Filters** pane to the left of the Device list includes a set of filters that you can use to narrow the devices that appear in the list:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(483).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A33Z&se=2026-05-12T09%3A51%3A33Z&sr=c&sp=r&sig=SiyhPnUjVZVlIIPbDPMUC8Wur1mr3f9V82pNv%2B7WN50%3D)**Reset To Default** (present only when one or more filters are specified): Clears all filters.
* **Status**: Checkboxes that filter for devices with a specified status.
* **Site**: Checkboxes that filter for a specific site (see **About Sites**).
* **Location**: Checkboxes that filter for devices at a specified SNMP location.
* **Vendor**: Checkboxes that filter for devices from a specified vendor.
* **Model**: A field that shows a lozenge for any device models for which the Devices list is currently being filtered. When clicked, the field drops down a filterable list from which you can choose one or more additional models to filter. To remove the filter for a model, click the **X** at the right of its lozenge in the field.
* **Label**: A field that shows a lozenge for any labels for which the Devices list is currently being filtered. When clicked, the field drops down a filterable list from which you can choose one or more additional labels to filter. To remove the filter for a label, click the **X** at the right of its lozenge in the field.
* **Label**: Checkboxes that filter for a specific label (see **About Labels**).
* **Agent**: Checkboxes that filter for a specific Universal agent (see **Universal Agents**).

When a filter is applied in a given category, that category is represented in the **Search** field (see **NMS Devices UI**) as a single lozenge. You can remove a filter individually by unchecking it in the **Filters** pane, or remove all filters in the category by clicking the **X** at the right of the category's lozenge in the **Search** field.

#### Filter Application Rules

Kentik applies the following rules to filter categories and criteria:

* Criteria are ORed (match any) within categories and ANDed (match all) between categories.
* When criteria are selected in more than one category, an interface is displayed only if it matches at least one of the selected criteria in all of the categories with criteria selected.
* When no criteria are selected for a category, the interface may match any of the category’s listed criteria.

## Add ICMP-only Devices Page

This page opens from the **Ping Only** option in the **Add NMS Devices** dialog and enables you to designate devices for ICMP-only NMS monitoring of up/down status and latency.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(484).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A33Z&se=2026-05-12T09%3A51%3A33Z&sr=c&sp=r&sig=SiyhPnUjVZVlIIPbDPMUC8Wur1mr3f9V82pNv%2B7WN50%3D)

*Bulk import or individually add devices for ICMP-only monitoring.*

The page includes the following UI elements:

* **Devices list**: A table listing devices and enabling you to select devices using the checkbox at the left of each row. The table includes the following information fields for each device:

  + **Device name**: The name of the device.
  + **IP Address**: The IP address of the device.
  + **Agent**: The Universal agent that is monitoring this device  (see **Universal Agents**).
  + **Remove**: A button (trash icon) that opens a confirming popup that enables you to remove this device from the list.
* **Add Row**: A button that adds one row to the list. The fields for that row will be empty, requiring you to enter the information for a device.
* **Choose file**: A field with which you can browse for a CSV file to enable bulk import of multiple devices, each into a new row in the table. The file must contain `name`, `address`, and `agent_id` columns that provide the required information for each device (the first available agent will be used for any device for which no `agent_id` is specified).
* **Agent**: A drop-down from which you can select a Universal agent that will be set as the agent for all currently selected devices.
* **Remove**: A button that will remove all currently selected devices from the list.
* **Add Devices**: A button that will initiate ICMP-only monitoring of all devices that are currently selected in the list, after which you’ll be taken back to the NMS Devices page.

## NMS Device Details Page

The Device Details page, reached by clicking the device's name in the **NMS Device List**, provides detailed information about an individual device.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(486).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A33Z&se=2026-05-12T09%3A51%3A33Z&sr=c&sp=r&sig=SiyhPnUjVZVlIIPbDPMUC8Wur1mr3f9V82pNv%2B7WN50%3D)

*An NMS Device Details page provides information about an individual device.*

### NMS Device Details UI

The Device Details page contains the following UI elements:

* **Breadcrumbs** (on subnav): An indicator of the current device you are viewing. Click on the **Devices** breadcrumb to go back to the Devices page.
* **Actions** (on subnav): A drop-down menu from which you can choose a tab from which to export data as a CSV file (see **Device Details Actions**).
* **Configure** (on subnav): A button that opens the settings dialog for the device, whose tabs and controls are covered in **Device Settings**.
* **Status**: A lozenge showing the current status of the device.
* **Favorite** (star icon): An icon indicating whether the device is favorited. Click the icon to add or remove the designation as a favorite.
* **Name**: The device’s name.
* **IP**: The device’s IP address.
* **Manufacturer**: The logo of the device’s manufacturer.
* **Uptime**: A graph showing the device's uptime information over a time range:

  + **Time range**: A drop-down from which you can choose the time range shown by the graph.
  + **Percent**: An indicator stating the device's uptime percentage over the specified time range.
  + **Graph**: A set of segments that are colored green for uptime and gray for downtime. Hover over a segment to open a popup giving the device's status at the segment's timestamp.
* **Tabs**: A set of tabs that each present details related to the device (see **NMS Device Details Tabs**).

#### Device Details Actions

The **Actions** drop-down on the subnav enables you to choose from the following actions:

* **Interfaces**: Export the table on the **Interfaces** tab as a CSV.
* **Hardware**: Export the table on the **Hardware** tab as a CSV.
* **BGP Neighbors**: Export the table on the **BGP Neighbors** tab as a CSV.

### NMS Device Details Tabs

The following tabs are used to present NMS information about a given device:

* **Overview**: The **Details Overview Tab** shows general information and visualizations for the device.
* **Traffic**: The **Details Traffic Tab** provides a high-level overview of the traffic on your device.
* **Interfaces**: The **Details Interfaces Tab** shows a table giving information about each interface associated with the device. A counter next to the tab name displays the number of associated interfaces.
* **Hardware**: The **Details Hardware Tab** shows a table giving information about each hardware component on the device. A counter next to the tab name displays the number of components.
* **BGP Neighbors**: The **Details BGP Neighbors Tab** shows a table giving information about each BGP neighbor associated with the device. A counter next to the tab name displays the number of BGP neighbors for the device.
* **IS-IS**: The **Details IS-IS Tab** is present when the router is running the "Intermediate System to Intermediate System" (IS-IS) protocol. The tab lists all of the neighbor routers and the state of those neighbor relationships.

## Details Overview Tab

The **Overview** tab of the Details page for NMS devices is covered in the following topics.

### About Overview Details

The **Overview** tab contains general information as well as six different visualizations that focus on a specific aspect of your device. Each visualization is generated by its respective query and contains a graph displaying the query results (see **NMS Device Overview Charts**).  
  
 The general information section in the **Overview** tab contains the following device details:

* **Details**: An expandable pane showing general device information. See **NMS Device Details** for the list of information displayed.
* **Capabilities**: An expandable pane showing which types of monitoring are enabled on the device (see **NMS Device Capabilities**).
* **OSPF States**: An expandable pane showing Open Shortest Path First information (see **NMS Device OSPF States**).
* **LLDP States**: An expandable pane showing Link Layer Discovery Protocol information (see **NMS Device LLDP States**).
* **Routing**: A pane showing BGP details (see **NMS Device Routing**).
* **Charts**: A set of visualizations displaying metrics about the device (see **NMS Device Overview Charts**).

### NMS Device Details

The **Details** pane contains the following information:

* **Labels**: The labels (if any) that have been applied to this device.
* **Manufacturer**: The device’s manufacturer.
* **Model**: The device’s model number.
* **Serial Number**: The device’s serial number.
* **IP Address**: The IP address(es) from which the device sends data to Kentik.
* **Boot Time**: The time since the device was last initialized.
* **Location**: The data center where the device is located.
* **OID**: The device’s Object Identifier.
* **System Description**: The device’s system description.
* **OS Name**: The operating system the device runs on.
* **OS Version**: The version number of device’s operating system.
* **Expand**: A button to expand or collapse the pane.

### NMS Device Capabilities

The **Capabilities** pane contains the following information:

* **Flow**: Indicates whether flow is enabled on the device.
* **SNMP**: Indicates whether SNMP is enabled on the device.
* **BGP**: Indicates whether BGP is enabled on the device.
* **ST**: indicated whether Streaming Telemetry is enabled on the device.
* **Expand**: A button to expand or collapse the pane.

### NMS Device OSPF States

The **OSPF States** pane contains the following information:

* **Router ID**: The ID of the router.
* **ABR**: Indicates whether the router is an Area Border Router.
* **ASBR**: Indicates whether the router is an Autonomous System Boundary Router.
* **# of Interfaces**: The number of interfaces associated with the device.
* **# of Neighbors**: The number of neighbors associated with the device.
* **Expand**: A button to expand or collapse the pane.

### NMS Device LLDP States

The**LLDP States** pane contains the following information:

* **Chassis ID Type**: The device’s chassis identifier type.
* **Chassis ID**: The device’s chassis identifier.
* **System Name**: The device’s user-friendly name.
* **Expand**: A button to expand or collapse the pane.

### NMS Device Routing

The **Routing** pane contains the following information:

* **BGP Local AS**: The AS of the local device.
* **BGP Router ID**: The ID of the BGP router.
* **BGP Neighbors**: The number of BGP neighbors associated with the device.

### NMS Device Overview Charts

The **Overview** tab displays visualizations that focus on different aspects of your device. It contains the following graphs:

* **Highest Inbound Utilization**: Metrics for the highest inbound utilization from the device’s interfaces.
* **Highest Outbound Utilization**: Metrics for the highest outbound utilization from the device’s interfaces.
* **CPU**: Metrics for CPU utilization.
* **Memory**: Metrics for memory utilization.
* **BGP Prefixes Sent**: Metrics for the number of prefixes sent.
* **Temperature**: Metrics for the temperature of the device.
* **Connections**: A table showing information about the devices that are directly connected to the device (see **NMS Device Connections**).

#### NMS Device Connections

The **Connections** table on the **Overview** tab shows information about the devices that are directly connected to the device. The table includes the following columns:

* **Local Interface**: The name and description of the local interface. The name is a link that takes you to the **Details Interfaces Tab**, where the Interfaces list will be filtered to this interface.
* **Remote End**: The name of the remote (connected) interface (bottom line) and its device (top line). If the device is in your network and registered with Kentik, the device name will link to its **NMS Device Details Page** and the interface name will link to the device's **Details Interfaces Tab**, where the Interfaces list will be filtered to the interface.
* **Classification**: The classification (Connectivity and Network Boundary) of the connection (see **Classification Overview**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(487).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A33Z&se=2026-05-12T09%3A51%3A33Z&sr=c&sp=r&sig=SiyhPnUjVZVlIIPbDPMUC8Wur1mr3f9V82pNv%2B7WN50%3D)

*The Connections table shows direct connections to the device.*

#### Manual Connections

By default, connected devices are discovered automatically via Link Layer Discovery Protocol (LLDP), but they may also be identified manually with the dialog that opens from the **Add Connection** button at the upper right of the table. Manually identified connections are indicated with an **M** at the far left of their row in the Connections list. The row will also include a pencil icon indicating that the settings in the Add Connection dialog can be edited.  
  
The **Add Connection** dialog includes the following settings:

* **Local Device**: A filterable drop-down listing local devices.
* **Local Interface**: A filterable drop-down listing local interfaces.
* **Remote End Type**: The type of the connection endpoint, either a monitored device or an external link from your domain to another.
* **Remote Device**: A filterable drop-down listing remote devices.
* **Remote Interface**: A filterable drop-down listing remote interfaces.

## Details Traffic Tab

The **Traffic** tab provides a high-level overview of the traffic on your device. The tab mirrors the UI elements from the **Traffic Tab** of the Details page for an individual device in the Kentik portal's Network Explorer module:

* **Parameter controls**: Set the parameters of the query whose results are displayed in the page's graph and table (see **Parameter Controls**).
* **Traffic tab graph**: Displays visualizations of the results returned from the query specified with the Parameter Controls (see **Traffic Tab Graph**).
* **Traffic table**: A tabbed table (see **Traffic Tab Table**) in which each tab displays a top-X table corresponding to a dimension. The queries that return results for each tab are affected by the parameter controls.

## Details Interfaces Tab

The **Interfaces** tab mirrors **NMS Interfaces** and shows a filterable list of interfaces associated with the device.

## Details Hardware Tab

The **Hardware** tab of the Details page for NMS devices is covered in the following topics.

### About Hardware Details

The **Hardware** tab features a table showing details about the physical hardware components that make up the device. Each row of the list represents one hardware component and shows its values for each of the table's columns (see **Hardware List Columns**). You can also view additional details about the component by clicking on its row to open the **Hardware Details Drawer**.  
  
The**Hardware** tab contains the following UI elements:

* **Filters** (filter icon): A button that toggles the **Filters** pane.
* **Group By**: A drop-down from which you can choose how to group the hardware components in the list. This operation of this control is identical to that of the Group By drop-down described in **Device List Grouping**.
* **Search**: A field that shows lozenges for the currently set filters and also enables you to enter text. The Hardware list will be filtered to show only rows that contain the entered current filters and the entered text. Click the **X** at the right to clear any text from the field.
* **Hardware List**: A table listing the components of your organization’s NMS-monitored devices.
* **Filters pane**: A set of controls that enable you to filter the device list (see **Hardware List Filters**).

### Hardware List Columns

The columns displayed in the Hardware list are chosen in a **Customize Columns** popup. A maximum of 12 columns may be visible at once. The following columns are available:

* **ID**: The hardware component’s Kentik-assigned device ID.
* **Status**: A set of lozenges indicating the OperStatus of the component.
* **Name**: The component’s type and name.
* **Description**: The component’s description string.
* **Hardware Version**: The component’s hardware version.
* **Firmware Version**: The component’s firmware version.
* **Software Version**: The component’s software version.
* **Manufacturer**: The component’s manufacturer.
* **Mode**l: The component’s model number.
* **Serial #**: The component’s serial number.

### Hardware List Filters

The **Filters** pane to the left of the Hardware list includes a set of filters that you can use to narrow the hardware components that appear in the list:

* **Type**: Checkboxes that include only components of a specified type.
* **OperStatus**: Checkboxes that include only components of a specified operation status.

### Hardware Details Drawer

The hardware **Details** drawer slides out from the right of the page when you click on the hardware component’s row in the list. The drawer contains the following elements:

* **Close**: An **X** button at the upper right to close the drawer.
* **Type**: The hardware component’s type (e.g., CPU, chassis, etc.).
* **Name**: The component’s name.
* **OperStatus**: A lozenge indicating the operation status of the component.
* **Details**: Information on the component found in **Hardware List Columns**.
* **Metrics charts**: Collapsible panes displaying different metrics in graph view (see **Hardware Metrics Charts**).

#### Hardware Metrics Charts

The charts in the collapsible panes of the **Details** drawer are visualizations of query results. The charts included in the drawer depend on the hardware component type. Each pane contains the following UI elements:

* **Expand**: An icon that opens the chart in popup view.
* **Graph**: Hover anywhere over the chart to view a popup giving the data at the timestamp.

The following visualizations may be included in these panes:

* **CPU Utilization**: Metrics for CPU utilization.
* **Memory Utilization**: Metrics for MEM utilization.
* **Temperature**: Metrics for temperature.

## Details BGP Neighbors Tab

The **BGP Neighbors** tab of the Details page for NMS devices is covered in the following topics:

### About BGP Neighbors Details

The **BGP** **Neighbors** tab features a table showing details about the peering neighbors of the device. Each row of the list represents one neighbor and shows the peers values for each of the table's columns (see **BGP Neighbors List Columns**). You can also view additional details about the neighbor by clicking on its row to open the **BGP Neighbors Details Drawer**.  
  
The **BGP Neighbors** tab contains the following UI elements:

* **Filters** (filter icon): A button that toggles the **Filters** pane.
* **Search**: A field that shows lozenges for the currently set filters and also enables you to enter text. The BGP Neighbors list will be filtered to show only rows that contain the current filters and entered text. Click the **X** at the right to clear any text from the field.
* **BGP Neighbors List**: A table listing your organization’s NMS interfaces.
* **Filters pane**: A set of controls that enable you to filter the BGP Neighbors list (see **BGP Neighbors List Filters**).

### BGP Neighbors List Columns

The columns displayed in the BGP neighbors list depend on the settings in the **Customize Columns** popup (see **Customize Columns**). A maximum of 12 columns may be visible at once. The following columns are available:

* **Session**: The **Session State** of the BGP connection.
* **Peer AS**: The AS of a neighboring or peer router with which the network device is forming a BGP session.
* **Neighbor IP**: The IP address of the BGP peer.
* **Local AS**: The AS of the network device.
* **Prefixes Sent**: The number of prefixes sent to that BGP peer.
* **Prefixes Received**: The number of prefixes received from that BGP peer.
* **Network Instance**:
* **Hold time**: Indicates (in seconds) how long the device will wait before declaring the connection to the peer lost.
* **Keepalive**: Indicates (in seconds) the frequency at which keepalive messages are sent to the BGP peer.
* **Connect Retry**: Indicates (in seconds) how long the device waits before initiating a new connection attempt.
* **Local IP**: The IP address used by the device to establish a connection.
* **Local Port**: The port number used be the device.
* **Remote IP**: The IP address of the BGP peer.
* **Remote IP Type**: The type of IP address used by the BGP Peer.

#### Session State

The following session states indicate the progress and health of a BGP session:

* **Idle**: The initial state of a BGP connection.
* **Connect**: The device is waiting for its connection to complete.
* **Active**: The device was unsuccessful in establishing a connection. It will continue to retry the connection.
* **Open Sent**: The device has sent a message and is waiting from its peer.
* **Open Confirm**: The device has received a message and is now waiting for its peer.
* **Established**: The BGP peering session is established and routing information is shared.

### BGP Neighbors List Filters

The **Filters** pane to the left of the neighbors list includes a set of filters that you can use to narrow the BGP neighbors that appear in the list:

* **Session State**: Checkboxes that include only neighbors with a specified connection state.
* **Local AS**: Checkboxes that include only neighbors with a specified local AS.
* **Peer AS**: Checkboxes that include only neighbors with a specified peer AS.

### BGP Neighbors Details Drawer

The BGP neighbors **Details** drawer slides out from the right of the page when you click on the neighbor’s row in the list. The drawer contains the following elements:

* **Close**: An **X** button on the upper right to close the pane.
* **Name**: The peer AS of the BGP neighbor.
* **Session State**: A lozenge indicating the state of the BGP connection.
* **Session Details**: A collapsible pane displaying the BGP neighbor details found in **BGP Neighbors List Columns**.
* **Metrics charts**: Collapsible panes displaying different metrics in graph view (see **BGP Neighbors Metrics Charts**).

#### BGP Neighbors Metrics Charts

The charts in the collapsible panes of the **Details** drawer are visualizations of query results. Each pane contains the following UI elements:

* **View in Metrics Explorer**: A button that takes you to the **Metrics Explorer**, where the Query sidebar will be set to show the query.
* **Graph**: Hover anywhere over the chart to view a popup giving the data at the timestamp.

The BGP Neighbors metric visualizations contain the following graphs:

* **Session State**: Metrics for session state.
* **Prefixes Received**: Metrics for prefixes received.
* **Prefixes Sent**: Metrics for prefixes sent.
* **Messages In**: Metrics for messages in.
* **Messages Out**: Metrics for messages out.

## Details IS-IS Tab

The **IS-IS** tab of the Details page for NMS devices is covered in the following topics.

### About IS-IS Details

The **IS-IS** tab is present when the router is running the "Intermediate System to Intermediate System" (IS-IS) protocol, a link-state interior gateway protocol (IGP) that uses the shortest-path-first (SPF) algorithm to determine routes. The tab includes a table that lists all of the neighbor routers and the state of those neighbor relationships.

### IS-IS Tab UI

The **IS-IS** tab contains the following UI elements:

* **Filters** (filter icon): A button that toggles the **Filters** pane.
* **Search**: A field that shows lozenges for the currently set filters and also enables you to enter text. The **IS-IS** list will be filtered to show only rows that contain the current filters and entered text. Click the **X** at the right to clear any text from the field.
* **IS-IS list**: A table listing your organization’s NMS interfaces (see **IS-IS List Columns**).
* **Filters pane**: A set of controls that enable you to filter the **IS-IS** list. The filters currently include only **Adjacency State (Up)**.

### IS-IS List Columns

The **IS-IS** list includes the following columns:

* **State**: The IS-IS state of the neighbor (see **IS-IS States**).
* **Neighbor Device**: The neighbor device's hostname. If that device is registered in Kentik the name will be a link.
* **Neighbor IP**: The neighbor's v4 IP.
* **Neighbor IP (v6)**: The neighbor's v6 IP (if applicable).
* **Interface**: The interface participating in the relationship.

#### IS-IS States

IS-IS neighbor states are defined as follows:

* **Down**: The initial state, meaning no hello packets have been received from the neighbor
* **Initializing**: The local router has received hello packets from the neighbor, but it's not sure if the neighbor has received the local router's hello packets
* **Up**: It's confirmed that the neighbor router receives the local router's hello packets
