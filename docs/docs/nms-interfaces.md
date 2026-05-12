# NMS Interfaces

Source: https://kb.kentik.com/docs/nms-interfaces

---

This article discusses the **Interfaces** page in the Network Monitoring System (NMS) section of the Kentik portal.

> **Note:** The following articles provide more information about Kentik's Network Monitoring System:
>
> * **Network Monitoring**
> * **Kentik NMS Configuration**
> * **NMS Setup**
> * **NMS Dashboard**
> * **Metrics Explorer**
> * **NMS Devices**
> * **Query Assistant**

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(488).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A20Z&se=2026-05-12T09%3A35%3A20Z&sr=c&sp=r&sig=GKiGhfn09EnoUPdh%2F18K%2FlRMSYjHG2I98qo3HmcCrPI%3D)

*A list of the interfaces on the devices being monitored by Kentik NMS.*

## About NMS Interfaces

|  |  |
| --- | --- |
| **Purpose:** | Monitor and view details of your NMS Interfaces in the Kentik portal. |
| **Benefits:** | - Get a centralized view of all NMS devices. - See what's important to you by filtering and customizing columns. - Speed troubleshooting by seeing key metrics. |
| **Use Cases:** | - Monitor the health and status of all your NMS interfaces. - Diagnose issues by analyzing metrics. |
| **Relevant Roles:** | Network Engineer |

NMS interfaces are interfaces on devices whose data in areas such as status, traffic, and utilization is collected by Kentik's Universal Agent (see **Universal Agents**) using SNMP or Streaming Telemetry, and passed to Kentik NMS for display and querying in portal modules such as the **Network Monitoring System** page and the **Metrics Explorer**. NMS interfaces are managed from the Interfaces page in the Network Monitoring System (NMS) section of the Kentik portal.

## NMS Interfaces Page

The NMS Interfaces page is reached from the main navbar by clicking **Interfaces** under **Network Monitoring System**. The page lists all of the interfaces on your organization's devices that are monitored by Kentik NMS. The list is a customizable table that can include information ranging from status and errors to capacity and utilization.  
  
The NMS Devices page includes the following main UI elements:

* **Breadcrumbs** (on subnav): An indicator of your current location within the Kentik portal.
* **Actions** (on subnav): A drop-down menu from which you can choose multiple actions to take (see **Interfaces Actions**).
* **Show/Hide Filters** (filter icon): A button that toggles the Filters pane between expanded and collapsed.
* **Group by**: A drop-down that determines how the interfaces displayed in the list will be organized into groups, either by Device, Type, or not grouped at all.
* **Search field**: Filters the devices list to show only rows that contain the entered text.
* **Filters pane**: A set of controls that enable you to filter the device list (see **Interface List Filters**).
* **Interface list**: A table listing your organization’s NMS interfaces (see **NMS Interface List**).
* **Customize**: A button that opens a popup from which you can choose the columns to display in the Interface list (see **Customize Columns**).

#### Interfaces Actions

The following actions can be chosen from the drop-down Action menu (in the subnav):

* **Export**: Export the Interface list as tabular data in CSV format. A notification appears when the CSV file is ready to download.
* **Manage Interfaces**: Go to the **Interfaces** page in portal Settings.
* **View in Metrics Explorer**: Go to the **Metrics Explorer** page.

## NMS Interface List

The Interface list is covered in the following topics.

### About the Interface List

The Interface list is a filterable table that shows information about the interfaces on your organization's devices. Each row of the list represents one interface and shows that interface's values for each of the table's columns (see **Interface List Columns**). For additional details about an interface, click on any white space within an interface’s row to open a Details drawer (see **NMS Interface Details**).  
  
You can change the sort order of the list by clicking on the heading of the column you'd like to sort by. Click the resulting blue up or down arrow to choose the sort direction (ascending or descending). In the bottom right corner, you can see when the list was last updated, how many interfaces are currently listed, and the total number of interfaces currently monitored.

### Interface List Columns

The columns displayed in the Interface list depend on the settings in the **Customize Columns** popup. A maximum of 12 columns may be visible at once. The following columns are available:

* **ID**: The interface’s Kentik-assigned device ID.
* **Admin** (default): A lozenge indicating the status of the device:

  + Up (green): Indicates the device is intended to be operational.
  + Down (red): Indicates the device is intentionally not operational.
* **Oper** (default): A lozenge indicating the **Interface OperStatus** of the device.
* **Name** (default): The interface’s name.
* **Description** (default): The interface’s description string.
* **Device** (default): The device where the interface is found. The device name links to its associated **NMS Device Details Page**.
* **Bit Rate** (default): The inbound and outbound bit rate of the interface.
* **Utilization** (default): The inbound and outbound utilization calculated as a percentage of the bit rate over the capacity.
* **Errors** (default): The number of inbound and outbound errors for the interface in the last hour.
* **Discards** (default): The number of inbound and outbound discards for the interface in the last hour.
* **Capacity** (default): The maximum possible bit rate of the interface.
* **MTU**: The maximum packet size, in bytes, that the interface will accept.
* **Index**: The unique identifier for this interface on its device.
* **Physical Address**: The MAC address associated with the interface.
* **Type**: The interface’s network protocol.

#### Interface OperStatus

The operation status lozenges contain the following possible values:

* **Up** (green): The interface is up and operational.
* **Down** (red): The interface is not operational.
* **Not Present** (gray): Indicates that the interface is not present.
* **Lower Layer Down** (gray): Indicates that the interface is down because one or more of the lower layer interfaces it depends on are down.

#### Customize Columns

To choose which of the columns (up to 12) are shown in the Interface list, and in what order:

1. Click the **Customize** button at the upper right of the list.
2. In the resulting **Customize Columns** popup, use the checkboxes to select or deselect columns.
3. Use the handle at the left of any column to drag it into the desired order.
4. Click anywhere outside the popup to close it. The list will be updated to reflect your column selections.

### Interface List Filters

The **Filters** pane to the left of the Interface list includes the following controls:

* **Reset To Default** (present only when one or more filters are specified): Clears all filters.
* **Selection fields:** Click to drop down a menu from which you can choose an item as a filter criteria, which will appear as a lozenge in both the selection field and the Search field (at the top of the page). Repeat to add more items. To remove a criteria, click the **X** at the right of its lozenge.
* **Checkboxes:** Click checkboxes in a given category (see **Filter Categories**) to choose one or more filter criteria in that category. Once checked, the criteria will appear as a lozenge in the Search field. To remove, click the **X** at the right of its lozenge or uncheck the box in the Filters pane.
* **Text fields:** Enter a string to filter the list to interfaces that contain the string in one of their columns.

#### Filter Categories

Filter criteria for the Interface list fall into the following categories:

* **Type** (selection field): Include only interfaces whose type matches one or more of the selected values.
* **OperStatus** (checkboxes): Include only interfaces whose OperStatus matches one or more of the checked values.
* **AdminStatus** (checkboxes): Include only interfaces whose AdminStatus matches one or more of the checked values.
* **Capacity** (selection field): Include only interfaces whose capacity matches the chosen value(s).
* **MTU** (selection field): Include only interfaces whose maximum packet size (a.k.a. Maximum Transmission Unit) matches the chosen value(s).
* **Index** (text field): Include only interfaces whose full index number matches the entered value.

You can use the above filter controls to narrow the set of listed NMS interfaces. The list will update to show only the matching interfaces.

#### Filter Application Rules

Kentik applies the following rules to filter categories and criteria:

* Criteria are ORed (match any) within categories and ANDed (match all) between categories.
* When criteria are selected in more than one category, an interface is displayed only if it matches at least one of the selected criteria in all of the categories with criteria selected.
* When no criteria are selected for a category, the interface may match any of the category’s listed criteria.

## NMS Interface Details

Additional details are available for every interface in the Interface list. These details may be displayed in one of two ways depending on the width of the browser window:

* If the viewport width is less than 1600px, details are presented in a drawer that opens from the right of the page when you click an interface's row in the Interface list.
* If the viewport width is greater, the drawer is displayed as a sidebar showing details for the interface in the most recently clicked row of the list.

The details drawer contains the following controls and additional details (depending on availability for a given interface):

* **Name**: ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(489).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A20Z&se=2026-05-12T09%3A35%3A20Z&sr=c&sp=r&sig=GKiGhfn09EnoUPdh%2F18K%2FlRMSYjHG2I98qo3HmcCrPI%3D)The name of the interface.
* **Close**: An **X** button at the upper right that closes the pane.
* **Admin Status**: A lozenge indicates the status of the interface:

  + **Up** (green): Indicates the interface is intended to be operational.
  + **Down** (red): Indicates the interface is intentionally not operational.
* **Oper Status**: A lozenge indicates the **Interface OperStatus**.
* **Description**: The description string defined in the device itself (if provided).
* **ID**: The interface’s Kentik-assigned device ID.
* **Index**: A unique identifier for each interface on a device.
* **Logical**: A lozenge indicating "false" if the interface is physical and "true" if it is virtual or software-defined.
* **MTU**: The maximum packet size, in bytes, that will be accepted by this interface.
* **Physical Address**: The MAC address associated with the interface.
* **Capacity**: The maximum possible bit rate of the interface.
* **Type**: The interface’s network protocol.
* **Connections**: A collapsible pane containing a table showing information about the interfaces that are directly connected to this interface:

  + **Remote End**: The name and description of the remote (connected) device.
  + **Classification**: The classification (Connectivity and Network Boundary) of the connection (see **Classification Overview**).

    > **Note:** For more detail about connections, or to manually add a connection, see **NMS Device Connections**.
* **Traffic visualization panes**: Panes containing charts showing traffic on the interface (see **Traffic Visualization Panes**).

#### Traffic Visualization Panes

The traffic visualizations for interfaces are presented in collapsible panes that contain the following UI elements:

* **View in Metrics Explorer**: A button to take you to the **Metrics Explorer**, where the **Query** sidebar will be set to run the query shown in the chart.
* **Chart**: A visual representation of traffic for this interface over the last 24 hours. Hover over any point in the graph to view a popup with the timestamp and data point.

The following traffic visualization panes are included in the interface details:

* **In Utilization**: A chart of the interface's inbound utilization in percent.
* **Out Utilization**: A chart of the interface's outbound utilization in percent.
* **In Errors**: A chart of the error count for inbound traffic on the interface.
* **Out Errors**: A chart of the error count for outbound traffic on the interface.
