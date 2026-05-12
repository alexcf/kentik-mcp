# Interfaces

Source: https://kb.kentik.com/docs/interfaces

---

This article covers the **Interfaces** page in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(616).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A28Z&se=2026-05-12T09%3A46%3A28Z&sr=c&sp=r&sig=EnzjschuyNXJAiieKZM7r95OJW1KpwJG7boRUdDnjOA%3D)

*The Interfaces page enables you to see your device's interfaces and edit their properties.*

## About Interfaces

Each interface on your organization's Kentik-registered devices (routers, switches, or other network hardware; see **About Devices**) is represented within Kentik (e.g., on the portal's Settings » **Interfaces** page) as an individual object with its own properties. The interfaces themselves and the values of their properties are typically discovered via SNMP polling (see **SNMP OID Polling**), but Kentik also allows manual setting of interface information in the following ways (see **Add or Edit an Interface**):

* Interfaces may be added manually, e.g., when SNMP polling is not enabled on the device (see **Manual vs. SNMP Discovery**).
* The SNMP-discovered values of interface properties may be manually overridden (see **Interface Overrides**).

> **Note:** The interfaces managed via the portal's Settings » **Interfaces** page do not include interfaces associated with the "cloud devices" listed on the Public Clouds page.

### Manual vs. SNMP Discovery

Many of the most powerful features of the Kentik platform are enabled by Interface Classification (see **About Interface Classification**), which relies heavily on the recognition of patterns in the names and descriptions of your interfaces. The way those names and descriptions are populated in Kentik depends on whether your organization enables SNMP polling of devices (see **SNMP OID Polling**):

* **SNMP Interface Discovery**: If your organization's devices are configured to allow SNMP polling from Kentik, interface names and descriptions will be populated as your interfaces are discovered automatically. If the name or description of an individual interface is problematic for Interface Classification (e.g., deviates from pattern), you can manually override the value in theEdit Interface dialog. An overridden value is static; it will not be automatically restored to the original value when the device is next polled.
* **Manual Interface Creation**: If your organization doesn't enable SNMP polling from Kentik you will need to create each interface within Kentik, either manually via the **Add Interfaces** page or programmatically via the **Device API**. For optimal results with Interface Classification, be sure to use consistent, parse able patterns in the interface names and descriptions.

> **Note:** To add an interface manually you'll need its SNMP ID (see **Getting an Interface Index**).

### Getting an Interface Index

To add an interface manually you'll need that interface's interface index (`ifIndex`), or SNMP ID. Since the ifIndex is present in a flow record for source and destination interface tuples, naming interfaces based on their actual SNMP ID will allow you to leverage all of the filtering available in Data Explorer and Dashboards, even if SNMP polling is not enabled on the interface's device.  
  
 The method used to determine the correct interface index or SNMP ID value depends on the kind of interface:

* **Interface on a host**:

  + **Get from Data Explorer**: Find this value in Data Explorer using a query whose device is the host that the interface is on, and whose dimensions are `Source:Interface` and `Destination:Interface`. The value will be shown in the Key column at the left of the resulting table.
  + **Derive from MAC address**: Find the MAC address of the interface. The interface index is the decimal value of the last two octets of that address.
* **Interface on a router or switch**:

  + **Device-specific commands**: Log onto the device with a CLI and retrieve the interface index `ifIndex` with the device-specific command shown in the table below.
  + **SNMP walk**: If your device isn't shown in the table below, you can SNMP-walk the IF-MIB, mapping interface names `(IF-MIB::ifName.<ID>)` to displayed IDs `(IF-MIB::ifIndex.<ID>)`.

    > **Note:** Routers must be configured for `ifIndex` persistence across reboots.

Device-specific commands to retrieve `ifIndex`:

| Manufacturer | Operating System | Command |
| --- | --- | --- |
| Cisco | iOS | `show snmp mib ifmib ifindex detail` |
| Cisco | iOS XR | `show snmp interface` |
| Cisco | NX-OS | `show interface snmp-ifindex` |
| Juniper | Junos | `(SNMP ifIndex <ID>)` is part of the standard `show interface <interface_name>` command. |
| Arista | Arista EOS | `show snmp mib ifmib ifindex [<if-range>]` |
| Brocade/Foundry | Network OS | `show snmp ifindex` |

### Flow and SNMP Reporting

Unless otherwise filtered, the Interfaces list on the **Interfaces Page** shows both SNMP-discovered and manually-added interfaces. The information shown in each row of the table varies depending on how the interface was discovered:

* **SNMP-discovered interfaces**: The information is based on the following:

  + Flow reporting: The upper-line values in the **Traffic In** and **Traffic Out** cells are based on flow data reported to Kentik via flow records (NetFlow, sFlow, IPFIX, etc.).
  + SNMP reporting: The lower-line values in the **Traffic In** and **Traffic Out** cells are based on SNMP data collected by polling.
* **Manually-created interfaces**: The information in the **Interfaces** list is based only on flow reporting.

> **Note:** If the **Traffic In** and **Traffic Out** columns aren't shown you can add them to the table via the **Customize** button (see **Interfaces Page UI**).

## Interfaces Page

The Interfaces page is home to the Interfaces list, which is a table that lists all of the interfaces on your Kentik-registered devices and provides information for each.

> **Note:** Member-level users can view the Interfaces page but cannot make changes (the **Classify**, **Customize**, and **Add Interface** buttons are not shown).

### Access the Interfaces Page

You can get to the Interfaces page in either of the following ways:

* Choose **Settings** from the main menu, then click **Interfaces** under **Network Metadata**.
* In the **Device List** on the Devices page (Settings » **Network Devices**), click the row of the device whose interfaces you’d like to see. In the drawer that opens at right, click the **View Interfaces** button (under **Interfaces**). You'll be taken to the Interfaces page, where the **Device** field in the **Filters** pane (see **Interfaces List Filters**) will contain a lozenge representing the device that you clicked on in the **Device** list.

### Interfaces Page UI

The Interfaces page includes the following UI elements:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(617).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A28Z&se=2026-05-12T09%3A46%3A28Z&sr=c&sp=r&sig=EnzjschuyNXJAiieKZM7r95OJW1KpwJG7boRUdDnjOA%3D)

*Interface Classification values can be specified manually with the Classify popup.*

* **Selection indicator**: Indicates how many interfaces are currently selected (checkboxes checked) in the **Interfaces List**. The **Classify** controls apply to the selected interfaces.
* **Classify** (not visible to Member-level users): A button that is inactive until you check at least one interface checkbox. The button opens a popup in which you can set the Network Boundary, Connectivity Type, and Provider for the selected interfaces (see **Interface Classification Dimensions**).
* **PeeringDB IX** (not visible to Member-level users): This button is inactive until you check at least one interface checkbox. The button opens a popup with which you can set the Peering DB IX mapping for the selected interfaces (see **PeeringDB**).
* **Add Interface** (not visible to Member-level users): This button opens the **Add Interface** dialog (see **About Interface Dialogs**), where you can manually register an interface for the device currently shown in the dialog's **Device** selector.
* **Search field**: Filters the **Interfaces** list to show only rows containing the entered text in one of the following columns: Interface, Device, Site, Classification.
* **Filters pane**: A set of controls (fields and checkboxes) with which you can filter the interfaces shown in the **Interfaces** list (see **Interfaces List Filters**).
* **Interfaces List**: A table listing the interfaces on all of your organization's devices or a subset determined by the **Filters** pane (see **Interfaces List**).
* **Customize**: A button that opens a popup in which you can choose the columns to display in the **Interfaces** list (see **Customize Columns Popup**).

## Interfaces List

The main table on the Interfaces page is covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(618).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A28Z&se=2026-05-12T09%3A46%3A28Z&sr=c&sp=r&sig=EnzjschuyNXJAiieKZM7r95OJW1KpwJG7boRUdDnjOA%3D)

*A filterable list of the interfaces on your network devices.*

### About the Interfaces List

The **Interfaces** list is a table listing all of the interfaces on your organization's devices (or a subset determined by the current **Interfaces List Filters**). Each row of the table shows the flow and SNMP data currently reported to Kentik for an individual interface.  
  
The types of information shown for each interface depend on the current **Interfaces List Columns**, which are set in the **Customize Columns Popup**. Click on a column heading to sort the list (ascending or descending).  
  
 In addition to the information shown in the table, you can see additional details about a given interface:

* Click on any white space within an interface’s row to open a **Details** drawer (see **Interface Details**).
* Click the **Edit** icon (pencil) on the right side of an interface’s row, which opens an Edit Interface dialog (see **Interface Settings Dialog**) where you can review and edit settings.

### Interfaces List Columns

The **Interfaces** list can include up to 12 columns of information about the interfaces on your organization's devices. The columns shown at any given moment are set in the **Customize Columns Popup**. The following columns are available to be chosen from the dialog:

* **Interface**: A combination of multiple pieces of information about the interface into a single cell, which is structured as follows:

  + Interface Name and SNMP ID (in parentheses) on the top line;
  + Interface Description on the next line (if available);
  + "(Overridden)" on a third line if any of the interface's SNMP-derived values have been manually overridden.
* **SNMP ID**: Interface ID (ifIndex) as defined in the device itself and retrieved via SNMP.

  > **Note:** Routers must be configured for `ifIndex` persistence across reboots.
* **Name**: The name string defined in the device itself and retrieved via SNMP.
* **Description**: The description string defined in the device itself and retrieved via SNMP.
* **Device**: The device with which this interface is associated.
* **Site**: The physical location of the device.
* **Classification**: The Network Boundary, Connectivity Type, and Provider for the interface (see **Interface Classification Dimensions**).
* **Capacity**: The maximum possible bit-rate (in Mbps) of the interface as reported by SNMP.
* **IP Address**: The IP address for this interface.

  > **Note:** The IP address is not reported for manually-created interfaces.
* **Traffic In**: Indicators that show the detected volume of inbound traffic (traffic entering a port) on the interface within the last 10 minutes, either from flow data (upper line in the cell) or SNMP data (lower line); see **Interface Traffic Indicators**.
* **Traffic Out**: Indicators like those in the Traffic In column, except for outbound traffic (traffic leaving a port).
* **PeeringDB IX Mapping**: The name of the PeeringDB Exchange to which this interface is mapped.
* **Actions**: The far-right column of the list includes icons that function as buttons for the following actions:

  + **Edit**: Opens anEdit Interface dialog that allows you to manually change the values of certain interface properties (see **Interface Overrides**).
  + **Remove**: Opens a confirming dialog that allows you to remove the interface from the list.

#### Interface Traffic Indicators

The table below shows how the indicators in the **Traffic In** and **Traffic Out** columns of the **Interfaces** list vary depending on the flow and SNMP data detected on an interface within the last 10 minutes.

| Data type | Data state | Indicator | Volume |
| --- | --- | --- | --- |
| Flow (upper line) | Detected | Green checkmark | Stated in Mbps |
| Flow (upper line) | Not detected | White exclamation mark in gray triangle | Not shown |
| SNMP (lower line) | Detected; volume within 20% of flow volume | Green checkmark | Stated in Mbps |
| SNMP (lower line) | Detected; volume not within 20% of flow volume (see **Flow SNMP Mismatch**) | White exclamation mark in red triangle | Stated in Mbps |
| SNMP (lower line) | Not detected | White exclamation mark in gray triangle | Not shown |

### Customize Columns Popup

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(619).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A28Z&se=2026-05-12T09%3A46%3A28Z&sr=c&sp=r&sig=EnzjschuyNXJAiieKZM7r95OJW1KpwJG7boRUdDnjOA%3D)The **Customize Columns**popup enables you to choose up to 12 **Interfaces List Columns** to display in the **Interfaces** list. To access the dialog, click the **Customize** button on the top right of the **Interfaces List**.  
  
The dialog includes the following UI elements:

* **Choose columns**: Use the checkbox to the left of a column name to display that column in the **Interfaces List**.
* **Order columns**: Use the handles to the left of the checkboxes to click and drag the columns into the desired order.

Once you've chosen the columns to include, click outside of the dialog to close it. Your column selections will be saved for future visits to the Interfaces page.

### Interfaces List Filters

The **Filters** pane to the left of the **Interfaces List** includes a set of filters that you can use to narrow down the interfaces that appear in the list. This pane includes the following types of controls:

* **Reset to Default** (present only when one or more filters are specified): Click to remove any user-selected filters and reset the Filters pane to its default settings.
* **Selection fields**: Click to drop down a menu from which you can choose an item as a filter criteria, which will appear as a lozenge in the Search field (at the top of the page). Repeat to add more items. To remove a criteria, click the **X** at the right of its lozenge.
* **Checkboxes**: Click checkboxes in a given category (see **Filter Categories**) to choose one or more filter criteria in that category. Once checked, the criteria will appear as a lozenge in the **Search** field. To remove, click the **X** at the right of its lozenge or uncheck the box in the **Filters** pane.

> **Note:** The **Filters** pane shows only controls (checkboxes and selectors) for which there is a match with the interfaces shown in the **Interfaces** list. As the list is filtered the number and type of controls displayed in the **Filters** pane may change.

#### Filter Categories

Filter criteria for the Interfaces list fall into the following categories:

* **Device** (selection field): Include only interfaces on the chosen Kentik-registered devices.
* **Interface** (checkboxes): Narrow the list to interfaces that were manually created or for which one or more SNMP-discovered value(s) has been overridden (see **Interface Overrides**).
* **IP Address** (checkboxes): Include only interfaces matching the checked IP address types.
* **Connectivity Type** (checkboxes): Include only interfaces whose Connectivity Type matches one of the checked values.
* **Network Boundary** (checkboxes): Include only interfaces whose Network Boundary matches one of the checked values.
* **Capacity** (selection field): Include only interfaces whose maximum SNMP-reported capacity (in Mbps) matches the chosen values.
* **Provider** (selection field): Include only interfaces whose Provider matches one of the chosen values.
* **PeeringDB IX** (checkboxes): Include only interfaces whose PeeringDB Internet Exchange matches one of the chosen values.

  > **Notes:**
  >
  > + The IXes that you can select as filters are the IXes already mapped to one or more of your interfaces via either the Peering BD IX menu above the table or the PeeringDB Internet Exchange field in an interface settings dialog (see **Interface Field Definitions**).
  > + Connectivity Type, Network Boundary, and Provider are **Interface Classification Dimensions**.

#### Filter Application Rules

Kentik applies the following rules to filter categories and criteria:

* Criteria are ORed (match any) within categories and ANDed (match all) between categories
* When criteria are selected in more than one category, a row is displayed only if it matches at least one of the selected criteria in all of the categories with criteria selected.
* When no criteria are selected for a category, the row may match any of the category's listed criteria.

### Flow SNMP Mismatch

As noted above, the **Traffic In/Traffic Out** columns of the **Interfaces** list may show a difference between flow-reported traffic volume (upper indicators) and SNMP-reported traffic volume (lower indicators). If the difference is greater than 20 percent and either the flow-reported volume or the SNMP-reported volume is greater than 50 Mbps, instead of a green checkmark in the status column, you’ll see a white exclamation mark inside a red triangle. While this alerts users to the possibility of an issue, it does not mean that there actually is an issue. The following are examples of common scenarios in which there may be a difference between the flow and SNMP traffic data:

* If the network has a number of ports on a router that talk between each other on the same subnet (i.e. switched vs. routed), that inter-port traffic will be counted in SNMP traffic but not in flow traffic, resulting in a percentage difference that reflects the topology of the network and does not indicate a problem.
* If there is a mismatch between the flow sampling rate configured on a router and the sample rate entered when that router is registered with Kentik (i.e. router is configured for a sample rate of 2048 but sample rate in Kentik device setup is 1024), then the Traffic status column will report low but the SNMP status column (which isn't subject to sampling factor) will be correct. To correct the mismatch, change the settings on the router itself or in the Device Details dialog of the Kentik portal.
* If the router is overloaded it will drop flow records, which will be reflected in the flow indicators (upper) but not in the SNMP indicators (lower). You may be able to reduce the extent to which a router drops flow records by changing that router's flow sampling rate (see **Flow Sampling**; note that this involves a change to both the router's configuration and the device settings in the Kentik system). You may also wish to contact the router vendor for optional flow export mechanisms/accelerators.

### Interface Overrides

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(621).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A28Z&se=2026-05-12T09%3A46%3A28Z&sr=c&sp=r&sig=EnzjschuyNXJAiieKZM7r95OJW1KpwJG7boRUdDnjOA%3D)The "original" values of interface properties are either discovered via SNMP polling (see **SNMP OID Polling**) or, for a manually created interface, set in the Add Interface dialog (see **Interface Field Definitions**) when the interface is first created. If a property's original value has been manually changed the value is referred to as being "overridden." Kentik indicates overridden values in two places in the UI:

* **Interfaces list**: Kentik marks an overridden value by inserting the word "Overridden" (in parentheses) in the table cell containing the value.
* **Edit Interface dialog**: Fields containing overridden values are accompanied with an **Original value** indicator and a Restore link. Click the link to restore the field to the original value.

> **Notes:**
>
> * An SNMP-discovered value that is overridden in theEdit Interface dialog is static; it will not be automatically restored to the original SNMP value when the interface's device is next polled.
> * To filter the **Interfaces** list to find overrides, use the checkboxes under **Interface** in the **Filter** pane (see **Interfaces List Filters**).

## Interface Details

The **Details** pane for interfaces shows additional information about the interface whose row in the Interfaces list was most recently clicked. The pane is displayed at the right of the Interfaces page. If the page width is less than 1600px the pane will display in a drawer that is hidden by default but opens from the right when you click on the row.  
  
The **Detail****s** pane contains the following information (when available) about an interface:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(622).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A28Z&se=2026-05-12T09%3A46%3A28Z&sr=c&sp=r&sig=EnzjschuyNXJAiieKZM7r95OJW1KpwJG7boRUdDnjOA%3D)**Name**: The name of the interface.
* **Description**: The description string defined in the device itself and retrieved via SNMP.
* **Classification**: The values, if any, of the **Interface Classification Dimensions** (Network Boundary, Connectivity Type, and Provider) that have been applied to the interface.
* **SNMP ID (ifindex)**: A unique identifying ID for the interface.
* **View in Network Explorer**: A button that takes you to the Details page for this interface (see **About Details Pages**).
* **Traffic**: Indicators that tell you whether flow and SNMP data have been detected from this interface within the last 10 minutes, and if so, at what bit rate (maximum Mbps); see **Interface Traffic Indicators**.
* **IP Address**: The interface’s IP address. Click to access that IP’s Details page (see **About Details Pages**).
* **Secondary IPs**: Any secondary IP addresses that may apply. Click to access each IP’s Details page (see **About Details Pages**).
* **Netmask**: A list of any netmasks used by the interface.
* **Boundary ASNs**: The ASNs of the autonomous systems (AS) to which, so far as Kentik is able to determine based on traffic and BGP data, an edge (External) interface is connected:

  + If one boundary AS is detected, the ASN is stated as well as the percent of the interface's traffic via that AS (in parentheses).
  + If more than one AS is detected, there will also be a View ASNs button. Click the button to pop up a list of the interface’s Boundary ASNs and the percentage of traffic for each.
* **VRF**: If the interface is assigned to a VRF, the ID of that VRF is displayed, otherwise this field does not appear.
* **Cost Group**: A set of interfaces that share common values for a set of cost-related properties (see **About Cost Groups**).
* **Metadata**: SNMP-related interface attributes (type, capacity, parent, status, etc.) derived by polling the ifMib (SNMP global interface data model).
* **PeeringDB IX**: The PeeringDB Exchange to which this interface is mapped.

> **Note**: Click anywhere outside the **Details** drawer to close it (or press the Esc key).

## Interface Settings Dialog

Adding or editing an interface via the Kentik portal involves specifying information in the fields of the interface settings dialogs, which are covered in the following topics.

> **Note:** Interfaces can also be added and edited with the **Device API**.

![Interface settings for device pe2_ord1, including IP address and capacity details.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/IF-edit-interfaces-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A28Z&se=2026-05-12T09%3A46%3A28Z&sr=c&sp=r&sig=EnzjschuyNXJAiieKZM7r95OJW1KpwJG7boRUdDnjOA%3D)

### About Interface Dialogs

The Kentik portal uses interface settings dialogs to collect the information needed to retrieve flow and SNMP data from the interface. The required information is entered into the fields of either of the following dialogs:

* **Add Interface** when registering a new interface with Kentik.
* **Edit Interface** when editing an already-registered interface.

> **Notes:**
>
> * When a setting is changed (overridden) in the Edit Interface dialog, it will be marked with an **Original value** indicator and a **Restore** link. Click the link to restore the field to the original value.
> * An SNMP-discovered value that is overridden in the Edit Interface dialog is static; it will not be automatically restored to the original SNMP value when the interface's device is next polled.

### Interface Dialogs UI

The Add Interface and Edit Interface dialogs share the same layout and the following common UI elements:

* **Add Interface** (Add Interface dialog only): A button that saves the settings for the new interface and exits the dialog.
* **Save** (Edit Interface dialog only): A button that exits the dialog and saves any changes to interface settings.
* **Cancel**: A button — **X** at top right or **Cancel** at bottom right — that exits the dialog without adding an interface or changing interface settings.

### Interface Field Definitions

Each Add or Edit dialog contains the fields and selectors shown in the following table.

| Element | Type | Description |
| --- | --- | --- |
| Device | Drop-down list | Required field; a router or host that has been registered with Kentik. See **About Devices**. After you click Add Interface, the Device is no longer editable. |
| SNMP ID | Editable field | • SNMP-discovered interfaces:  The interface index (ifIndex) as defined in the device and retrieved via SNMP.  • Manually-added interfaces:   * On routers and switches, the interface index (ifIndex) as defined in the device and entered when adding or editing the interface. * On hosts, the decimal value of the last two octets of the MAC address of the interface that Kentik is configured to track.  **Note:** Must be uint32 (see **Getting an Interface Index**). |
| Name | Editable field | The interface name as either defined in the device and retrieved via SNMP or entered manually when adding the interface. Capped at 128 characters. |
| Capacity (Mbps) | Editable field | The maximum capacity in Mbps for the device as either retrieved via SNMP or entered manually when adding the interface. Must be uint32 (0 - 4294967295). |
| Description | Editable field | The interface description as either defined in the device and retrieved via SNMP or entered manually when adding the interface. Capped at 128 characters. |
| IP Address | Editable field | The IP address of this interface. |
| Netmask | Editable field | The netmask configured for the interface (applies to IPv4 interfaces only). |
| PeeringDB Internet Exchange | Drop-down list | The PeeringDB IX to which this interface is mapped. |
| Connectivity Type | Drop-down list | Attribute used to classify interfaces by their role in the overall network. |
| Network Boundary | Drop-down list | Attribute used to classify interfaces as internal or external to the network. |
| Provider/Customer | Editable field | The name of the customer or provider. |

> **Note:** Once an SNMP-discovered value is manually overridden in the Edit Interface dialog it becomes static; it will not be automatically restored to the original SNMP value when the interface's device is next polled.

## Add or Edit an Interface

Device interfaces can be added or edited on the Interfaces page, which is covered in the following sections.

> **Note:** Interfaces can also be added and edited with the **Device API**.

### Add an Interface

Interfaces and hosts that are not configured for SNMP are not automatically detected by Kentik and must be added manually.  
  
To add an interface:

1. Go to the Interfaces page (see **Access the Interfaces Page**).
2. Click the **Add Interface** button, which opens the **Add Interface** dialog.
3. Choose a device from the drop-down **Device** menu.
4. Set the values of the dialog's remaining fields and selectors (see **Interface Field Definitions**).
5. Click the **Add Interface** button to save the interface and return to the Interfaces page.

### Edit an Interface

You can edit the settings of both SNMP-discovered and manually-added interfaces. This enables you to override SNMP-discovered values within Kentik rather than having to modify configurations on a device. For example, if your interface descriptions have been populated automatically by code that you can't or prefer not to modify, you can edit those descriptions manually.  
  
To edit an interface:

1. Go to the Interfaces page (see **Access the Interfaces Page**).
2. On the Interfaces page, click the Edit icon (pencil) at the right of the row of the interface that you'd like to edit. The **Edit Interface** dialog will open.
3. In the dialog, modify the values that you'd like change (see **Interface Field Definitions**).
4. Click the **Save** button.

> **Notes:**
>
> * An SNMP-discovered value that is overridden in the Edit Interface dialog is static; it will not be automatically restored to the original SNMP value when the interface's device is next polled.
> * You can also open the Edit Interface dialog for a given interface from the **Configure** button on the interface's Details page in Network Explorer (see **Details Page UI**).

### Remove an Interface

To remove an interface:

1. Go to the Interfaces page (see **Access the Interfaces Page**).
2. Click the **Remove** icon (trash can) on the interface’s row in the list.
3. In the confirming dialog, click Remove.

> **Note:** While you can remove an SMNP-discovered interface in Kentik, the interface will reappear when its device is next polled.
