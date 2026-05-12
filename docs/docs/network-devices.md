# Networking Devices

Source: https://kb.kentik.com/docs/network-devices

---

The Devices page in the Kentik portal allows you to view information about existing devices registered with Kentik and add or edit devices.

> **Notes:**
>
> * Users with a Kentik role of Member cannot add, edit, or remove devices.
> * Device settings made in the **Device Settings Dialog** are covered in **Device Settings**.
> * Devices may be set up during your organization's initial onboarding with Kentik (see **Onboarding**).
> * For assistance with adding or managing devices, contact Kentik (see **Customer Care**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(549).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A03Z&se=2026-05-12T09%3A47%3A03Z&sr=c&sp=r&sig=XLq%2BbJd7XrV4COYSAkvwDh3To%2FUKvFe0A2tF0gO%2Bve0%3D)

*The Devices page displays device status and enables you to manage devices.*

## About Devices

Devices are one of the primary types of **Data Sources** in Kentik, classified as either routers (including switches and firewalls) or hosts (utilizing Kentik's software host agent). They can send flow data (e.g., NetFlow, sFlow, IPFIX), be polled for SNMP data, and have BGP data read by Kentik. This data is correlated with other information like GeoIP and DNS in the **Kentik Data Engine** (KDE).

#### Enable a Device to Send Data to Kentik

* **Register the Device:** Register a new device in the Kentik portal (see **Manage Devices**).
* **Configure** **the Device**:

  + For routers or networking hardware, see **Router Configuration**.
  + For hosts using Kentik’s software host agent, see **Host Configuration**.

## Supported Device Types

Kentik supports devices in two main categories:

* **Routers**: Includes switches and firewalls.
* **Hosts**: Includes cloud resources.

These categories differ in how they store and report traffic and use different **Device Settings**. The table below lists the device types available in the portal, categorized accordingly. The types are listed in the order in which they appear in the **Type** drop-down on the **General** tab of the **Device** settings dialog (see **Device General Settings**).

> **Notes:**
>
> * All router types are fundamentally "NetFlow-enabled routers”. Choosing a specific device type (e.g., "Cisco IOS XR") allows for the collection of **Device-specific Dimensions**.
> * Selecting "MPLS router" enables data collection for **MPLS Dimensions** but not for device-specific dimensions.

| Device Type | Category | Subtype | Description |
| --- | --- | --- | --- |
| NetFlow-Enabled Router | Router | router | Hardware router or switch. |
| Kentik Host Agent (kprobe) | Host | kprobe | Kentik's software host agent (see **About kprobe**). |
| A10 CGN | Router | a10\_cgn | A10 Thunder Carrier Grade Networking devices (see **A10 Thunder CGN Dimensions**). |
| Advanced sFlow | Router | advanced\_sflow | Any sFlow device from which you want Kentik to ingest a value for the TTL dimension and/or for the Physical Interface dimensions (Src and Dst), which requires use of Kentik's VLAN Mapping API (ask Product Support for assistance). |
| Cisco ASA | Router | cisco\_asa | Cisco Adaptive Security Appliance (see **Cisco ASA Dimensions**). |
| Cisco ASA (Syslog) | Router | cisco\_asa\_syslog | Syslog data from a Cisco Adaptive Security Appliance (see **Cisco ASA Syslog Dimensions**). |
| Cisco NBAR-Enabled Router | Router | cisco\_nbar | Cisco router that supports traffic prioritization using Network Based Application Recognition (**https://www.cisco.com/c/en/us/products/ios-nx-os-software/network-based-application-recognition-nbar/index.html**). |
| Cisco nvzFlow | Router | cisco\_nvzflow | Cisco Network Visibility Flow for AnyConnect VPN client (see **Cisco nvzFlow Dimensions**). |
| Cisco SD-WAN IOS XE | Router | cisco\_sdwan\_xe | Cisco SD-WAN IOS XE Edge devices (see **Cisco SD-WAN vEdge Dimensions**) |
| Cisco Zone-Based Firewall | Router | cisco\_zone\_based\_firewall | Cisco router using a zone-based firewall (see **Cisco Zone-based Firewall**). |
| Darknet Stream | N.A. | N.A. | Reserved for Kentik use. |
| Fortinet FortiGate | Router | fortinet\_fortigate | Fortinet FortiGate firewalls and SD-WAN devices (see **FortiGate Dimensions**). |
| Cisco IOS XR | Router | ios\_xr | Data from routers using the IOS XR operating system (see **Cisco IOS XR Dimensions**). |
| Istio (Beta) | Router | istio | Istio cloud logs |
| Juniper DDoS Flow | Router | juniper\_ddos | Juniper Flow-based Telemetry |
| Process-Aware Telemetry Agent | Host | kappa | Host agent used for Kentik Kube\*. It generates flows from network traffic based on eBPF. |
| kProbe True Origin Tap | N.A. | krobe\_tap | Reserved for Kentik internal use. |
| Cisco Meraki | Router | meraki | A Meraki-managed firewall (see **Cisco Meraki Metrics**). |
| MPLS Router | Router | mpls | An MPLS-enabled router (see **MPLS Dimensions**). |
| Nokia Layer 2 | Router | nokialayer2 | Nokia service routers for IP edge and core applications using L2 flow template. |
| ntop Host Agent (nProbe) | Host | nprobe | Deprecated. |
| Oracle Cloud Infrastructure | Host | oci\_subnet | Oracle Cloud Infrastructure |
| Palo Alto Networks Firewall | Router | paloalto | A PAN firewall (see **Palo Alto Networks Firewall**). |
| Palo Alto Prisma SD-WAN | Router | paloalto\_prismasdwan | Palo Alto Prisma SD-WAN (sdwan) |
| Juniper PFE (Syslog) | Router | pfe\_syslog | Syslog data from a Juniper switch equipped with a Packet Forwarding Engine (see **Juniper PFE Syslog Dimensions**). |
| sFlow Tunnel decode | Router | sflow\_tunnel | Network device sending sFlow logs, with additional decode of the tunneling protocols. This device type is used for decoding VXLAN data into native dimensions (see **sFlow Tunnel Decode Dimensions**). |
| Silver Peak EdgeConnect | Router | silverpeak | Silver Peak appliance running VXOA software (see **Silver Peak Dimensions**). |
| Generic Syslog | Router | syslog | A generic device sending syslogs to kproxy (see **kproxy Syslog Parsing**). |
| Versa | Router | versa | Versa Networks SD-WAN Edge devices (sdwan) |
| Cisco SDWAN vEdge | Router | viptela | IPFIX fields in cflowd records from Cisco vEdge SD-WAN routers (see **Cisco SD-WAN vEdge Dimensions**). |
| VMware SD-WAN | Router | vmware\_velocloud | VMware SD-WAN Edge devices (see **VMware SD-WAN Dimensions**). |
| VMWare vSphere | Router | vmware\_vsphere | VMware vSphere flow log export |
| VXLAN | Router | vxlan | Network device sending sFlow logs, with additional decode of VXLAN protocol fields into separate dimensions (see **VXLAN Dimensions**). |

\*If you don’t see Kentik Kube in your portal, it might not be enabled for your company’s account. Please contact your **Kentik Account Team**.)

> **Note:** In the **Device API**, the above device types are referenced using the `subtype` values listed above.

## Devices Page

The Devices page of the Kentik portal may be reached by choosing **Settings** from the main menu, then clicking **Networking Devices** (under **Data Sources**). All registered devices for your organization are listed on the page, which includes the following main UI elements:

* **Share** (on SubNav): Opens the **Share** dialog to enable you to share the current view (see **Sharing via the Share Dialog**).
* **Actions** (on SubNav): A drop-down menu from which you can choose to export or subscribe to the current view:

  + Export: Exports the page’s content as a visual report (PDF) or data table (CSV). A notification appears when the export is ready to download.
  + Subscribe: Opens the **Subscribe** dialog enabling you to subscribe to regular reports from the page, either by choosing an existing subscription (combination of email address and schedule) or specifying a new one. The form in the **Subscribe** dialog is the same as on the **Subscription** tab of the **Share** dialog, which is covered in **Subscription Tab UI**.
  + Unsubscribe: Opens the **Unsubscribe** dialog. Click the **Subscription** drop-down, select the subscription from which you’d like to unsubscribe, and click **Unsubscribe.**

    > ***Note***: **Unsubscribe** will only appear if you have an existing subscription for this page.
* **Selection indicator**: Indicates how many devices are currently selected in the **Device** list (see the **Select** column in **Device List Columns**). Label controls apply to the selected devices.
* **Label controls**: A set of buttons that allow you to manage labels for devices in the list (see **Label Controls**).
* **Add Device**: A button that opens a **Device Settings Dialog**, which allows you to add a new device.
* **Show/Hide Filters** (filter icon): A button that toggles the **Filters** pane between expanded and collapsed.
* **Group By**: A drop-down from which you can choose how to group the devices on the list (see **Device List Grouping**).
* **Search**: A field that shows lozenges for the filters currently set in the **Filters** pane and also enables you to enter text. The **Device** list will be filtered to show only rows that meet the filter criteria and contain the entered text. Click the **X** at the right of the field to clear entered text, and the **X** in a lozenge to clear the corresponding filter.
* **Filters pane**: A set of controls that enable you to filter the **Device** list (see **Device List Filters**).
* **Device list**: A table listing your organization’s devices (see **Device List**).

#### Device List Grouping

The **Group By** drop-down enables you to group the rows of the **Device** list based on one of the following options:

* **None**: Devices are not grouped.
* **Site**: The devices are grouped by the sites to which each is assigned (see **About Sites**).
* **Plan**: The devices are grouped by their Kentik billing plan (see **About Plans**).
* **Flow Type**: The devices are grouped by the **Flow Protocols** used to send flow data to Kentik.
* **Vendor**: The devices are grouped by their manufacturer.

## Device List

The **Device** list is covered in the following topics.

### About the Device List

The **Device** list is a table that shows information about your organization’s devices as well as available actions. By default, the list is ordered alphabetically by name (ascending). To change the sort order of the list, click on a heading to choose a sorting column, and on the resulting blue up or down arrow to choose the sorting direction (ascending or descending).  
  
Each row of the Device list shows the device's values for each of the **Device List Columns**, as well as **Device List Actions**. You can also show an extensive set of additional details about the device by clicking on the device's row, which slides out the **Device Details** drawer from the right of the page. If the drawer is already open and you click on the row for a different device, the details will be replaced with details for the new device. To close the drawer, click anywhere outside of the Device list.

### Device List Columns

The columns displayed in the **Device** list depend on the settings selected in the **Customize Columns** popup (see **Customize Device List**). The following columns are available:

* **Select All** (checkbox in header row): A checkbox for toggling the selection state of all devices in the **Device** list:

  + If no checkboxes in the list are checked, or only some are checked, then clicking this checkbox will select all checkboxes.
  + If all checkboxes in the list are checked, clicking this checkbox will deselect all checkboxes.
* **Select**: A checkbox at the left of each device row that selects the device for applying or clearing labels (see **Label Controls**).
* **Name**: Includes the following for each device:

  + Type: An icon indicating the category of the device (router, host, etc.; see **Supported Device Types**).
  + Name: Name of the device as specified in the **Device** dialog.
  + Labels: Any labels applied to the device, which display below the device name.
* **ID**: The device's Kentik-assigned device ID.
* **Status**: A set of icons indicating what types of data Kentik is receiving from the device (see **Device Status Icons**).
* **FPS Raw**: The rate at which the device has been generating flow records over the last six hours (before any Kentik downsampling).
* **FPS Sampled**: The rate at which Kentik has been ingesting flow records from the device over the last six hours. This value is after any downsampling Kentik performs pursuant to the terms of the plan to which the device is assigned.
* **BGP Prefixes**: For devices whose **Status** column shows BGP as established, this column shows the IP v4 (top) and v6 (bottom) prefixes currently (within last 5 minutes) in the BGP routing table.
* **Site**: Shows the name of the site to which the device is assigned and provides a link to that site’s Details page (see **About Sites**).
* **Plan**: The billing plan to which the device is assigned(see **About Plans**).
* **Flow Type**: The format of the flow data (e.g., NetFlow v5, v9, IPFIX, or SFlow). See **Flow Protocols**.
* **IP Address**: The IP address(es) from which the device sends data to Kentik.
* **Vendor**: The device’s manufacturer.

#### Customize Device List

To choose which of the columns above are displayed to you in your **Device** list:

1. Click the **Customize** button at the upper right of the **Device** list.
2. In the resulting **Customize Columns** popup, use the checkboxes to select or deselect columns.
3. Use the handle at the left of any column to drag it into the desired order.
4. Click anywhere outside the popup to close it. The list will be updated to reflect your column selections.

#### Device Status

The **Status** column of the **Device** list shows several **Device Status Icons** which indicate the flow status for the following types of data:

* **Flow**: Indicates whether Kentik is receiving flow records from the device.
* **SNMP**: Indicates whether Kentik is able to successfully poll the device for SNMP data.
* **BGP v4**: Indicates whether Kentik is peering with the device to receive BGP v4 routing data.
* **BGP v6**: Indicates whether Kentik is peering with the device to receive BGP v6 routing data.
* **ST**: Indicates whether Kentik is able to successfully receive streaming telemetry (ST) from the device.

#### Device Status Icons

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(551).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A03Z&se=2026-05-12T09%3A47%3A03Z&sr=c&sp=r&sig=XLq%2BbJd7XrV4COYSAkvwDh3To%2FUKvFe0A2tF0gO%2Bve0%3D)The icons in the table below are used to indicate the detected status of various types of data collection for each device.

| Icon | Check in green circle | Exclamation in orange triangle | Three hyphens |
| --- | --- | --- | --- |
| **Flow** | * Direct flow detected * Kentik agent detected | No flow detected | N.A. |
| **SNMP** | Detected | N.A. | Not detected |
| **BGP v4** | Established | Not established | Not configured |
| **BGP v6** | Established | Not established | Not configured |
| **Streaming Telemetry (ST)** | Detected | Not detected | Disabled or not detected |

### Device List Actions

The following actions are available from the icons at the right of each device's row in the Device list:

* **Edit** (pencil icon; not present when the device is archived): Opens a **Device Settings Dialog** in which you can edit the device’s settings (see **Edit a Device**).
* **Archive** (red file box with down arrow; not present when a device is archived): Opens a confirming dialog with which you can archive (deactivate) the device.

  > **Note:** Flow and other data from an archived device will no longer be ingested into Kentik or included in aggregates or evaluations on any portal page or process, including Network Explorer, Insights, Hybrid Map, etc.
* **Restore** (file box with up arrow; only present when a device is archived): A button that opens a confirming dialog enabling you to restore an archived device to active status.

  > **Note:** Archived devices cannot be restored when all of your organization's device licenses are already in use by active devices.
* **Remove** (trash icon; only present when a device is archived): Opens a confirming dialog with which you can permanently remove the device from your organization's collection of Kentik-registered devices.

  > **Note:** Removing a device will remove all of the device's flow records from the **Kentik Data Engine**.

## Device Details

The **Details** pane appears in a drawer that slides out from the right of the page when you click a device's row in the **Device** list. The pane contains the device-related information described in the topics below:

### Device General Details

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(552).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A03Z&se=2026-05-12T09%3A47%3A03Z&sr=c&sp=r&sig=XLq%2BbJd7XrV4COYSAkvwDh3To%2FUKvFe0A2tF0gO%2Bve0%3D)The top part of the **Details** pane contains general information about the device:

* **ID**: The device's Kentik-assigned device ID.
* **Name**: Includes the following:

  + Type: An icon indicating the category of the device (router, host, etc., see **Supported Device Types**).
  + Name: The name of the device as specified when the device was registered or edited.
* **Labels**: The labels, if any, assigned to the device.
* **View in Network Explorer**: A button that takes you to the Details page for this device (see **Core Details Pages**). Present only when Kentik is detecting flow from the device (FPS is greater than 0).
* **View Device Details**: A button that takes you to the Details page for this device in Kentik's NMS Devices module (see **NMS Device Details Page**). Present instead of the View in Network Explorer button when both of the following are true:

  + The device is enabled for NMS on the SNMP or Streaming Telemetry tab of the **Device Settings Dialog**.
  + Kentik is receiving no flow from the device (e.g., the device is not configured to send flow, the device has no traffic, the **Sending IP** is not set in the device's settings dialog, etc.).
* **Manufacturer**: Information about the device, obtained by SNMP polling, including the full name and version identification of the system’s hardware type, software operating-system, and networking software (see `sysDescr` in **System Information OIDs**), if available.
* **Description**: The description string specified in the Device dialog (if provided).
* **Site**: The site to which this device is assigned.
* **Plan**: The billing plan to which the device is assigned (see **About Plans**).
* **Kproxy Agents**: The kproxy agents, if any, via which flow data is sent from the devices at this site to Kentik (see **Kentik Proxy Agent**).
* **Sending IPs**: IP address(es) from which the device sends flow.
* **Created**: The date the device was registered to the Kentik portal.

### Device Flow Details

The **Flow** section of the **Details** pane contains information related to the flow records sent from the device to Kentik (see **Flow Overview**):

* **Flow status**: Indicates whether Kentik is receiving flow from the device. The indicator has the following states:  
   - No Flow Detected (orange): Displayed when no flow is detected from the device.  
   - Direct Flow Detected (green): Displayed when flow is coming directly from a router.  
   - Kentik Agent Detected (green): Displayed when flow is coming via Kentik's `kproxy` software agent (see **Kentik Proxy Agent**).
* **FPS (Sampled)**: The rate at which Kentik has been ingesting flow records from the device over the last six hours. This value is after any downsampling Kentik performs pursuant to the terms of the plan to which the device is assigned.
* **FPS (Raw)**: The rate at which the device has been generating flow records over the last six hours (before any Kentik downsampling).
* **Plan FPS**: The maximum FPS supported by the plan to which this device is assigned.
* **Flow Type**: The format of the flow data, e.g., NetFlow v5, v9, IPFIX, or sFlow (see **Flow Protocols**).
* **Sample Rate (Avg)**: The ratio of total flows per sampled flow (see **Flow Sampling**).

### Device Interface Details

The **Interfaces** section of the **Details** pane contains information related to the interfaces on the device:

* **Interfaces Classified**: The number of classified interfaces, the total number of interfaces, and the percent of interfaces that have been classified (see **About Interface Classification**).
* **View Interfaces**: A button that takes you to the **Interfaces** page where you can manage the interfaces on your devices.

### Device SNMP Details

The **SNMP** section of the **Details** pane contains information related to SNMP polling of the device by Kentik (see **SNMP OID Polling**):

* **SNMP status**: Current status of SNMP polling:

  + No SNMP Detected (orange): Displayed when SNMP is not detected.
  + SNMP Detected (green): Displayed when SNMP is detected.
* **SNMP Polling**: The setting at which SNMP is polling (see **SNMP Polling Intervals**):

  + Standard: SNMP polling is every 5 min for interface counters and 3 hours for descriptions.
  + Minimal: Interface polling is disabled and descriptions are polled every 6 hours.
* **SNMP v3 Auth**: Indicates whether v3 authentication for SNMP polling is on or off (see **About SNMP V3**).

### Device BGP Details

The **BGP** section of the **Details** pane contains information related to BGP on the device (see **BGP Overview**):

* **BGP status**: Current status of BGP peering with the device, e.g., Established, Not Established, or Not Configured.
* **BGP Prefixes**: For devices with BGP enabled, this indicator shows the number of prefixes currently (within last 5 minutes) in the BGP routing table.

### Device ST Details

The **Streaming Telemetry** section of the **Details** pane shows the status of streaming telemetry from the device to Kentik:

* **Disabled** (grey): Streaming Telemetry is not enabled in the device's Kentik settings.
* **No ST detected** (orange): Streaming Telemetry is enabled in the device's Kentik settings, but Kentik is not receiving telemetry from the device.
* **ST Detected** (green): Streaming Telemetry is enabled and Kentik is receiving telemetry.

### Device Features

The **Features** section of the **Details** pane shows information about additional Kentik capabilities on this device:

* **CDN attribution**: Enabled if the device is a host and its **Contribute to CDN Attribution** switch is on in the **General** tab of the **Device** dialog (see **Device General Settings**).
* **Flowspec**: Enabled if the device's **BGP Flowspec Compatible** switch is on in the **BGP** tab of the **Device** settings dialog.
* **RTBH**: Enabled if the device is assigned to an RTBH mitigation platform (see **Configure Platform Devices**).

## Label Controls

Kentik's labeling feature enables you to create a label (essentially a property whose value is text) and apply it to one or more entities represented within the Kentik portal, creating a group that can be referenced collectively rather than individually (e.g., when filtering). The Devices page includes a set of label-related controls that are located above the **Search** field. The controls apply only to devices that are currently selected in the **Device** list:

* **Add/Edit Labels**: A link to the **Labels** page, where you can create, edit, or remove labels that are available to apply to the selected devices.![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(553).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A03Z&se=2026-05-12T09%3A47%3A03Z&sr=c&sp=r&sig=XLq%2BbJd7XrV4COYSAkvwDh3To%2FUKvFe0A2tF0gO%2Bve0%3D)
* **Clear Labels**: Clears the labels that are applied to all currently selected devices.
* **Apply Labels**: A field that shows, applies, and removes individual device labels (see **Apply Labels Field**).

#### Apply Labels Field

The **Apply Labels** field is active only when one or more devices are selected in the **Device** list, at which point the field is populated with all labels that are already applied to any selected device. Clicking in the field will drop down a filterable list of labels; clicking on any listed label will add that label to the field and to all selected devices. Clicking the **X** at the right of a label in the field will remove that label from all selected devices.

## Device List Filters

Device list filters are available in the **Filters** pane, which is covered in the following topics.

### Device Filters Pane

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(554).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A03Z&se=2026-05-12T09%3A47%3A03Z&sr=c&sp=r&sig=XLq%2BbJd7XrV4COYSAkvwDh3To%2FUKvFe0A2tF0gO%2Bve0%3D)The **Filters** pane to the left of the **Device** list includes a set of filters that you can use to narrow the list of devices that appear in the table (using **Filter Application Rules**). This pane includes the following types of controls:

* **Reset to default** (present only when one or more filters are specified): A button that removes all currently set filters.
* **Collapse**: A button that collapses the Filters pane. To expand the pane, click the **Show/Hide Filters** button (see **Devices Page**).
* **Selection field**: Each field displays lozenges for criteria selected from that field’s menu and filters the **Device** list to devices matching those criteria. To select criteria, click in the field and choose an item from the filterable drop-down list. To remove criteria, click the **X** at the right of its lozenge.
* **Checkboxes**: Checkboxes allow you choose one or more filter criteria for that category (see **Filter Categories**). Once checked, the criteria will appear as a lozenge in the **Search** field. To remove, click the **X** at the right of its lozenge or uncheck the box in the **Filters** pane.

### Filter Categories

Filter criteria for the **Device** list fall into the following categories:

* **Site** (selection field): Include only the devices at the selected sites.
* **Flow Status** (checkboxes): Include only devices whose status is:

  + Direct flow: Device’s flow is coming directly from a router or other network hardware.
  + Kentik Agent: Device’s flow is coming via Kentik's kproxy software agent (see **Kentik Proxy Agent**).
  + No flow: No flow is detected for the device.
* **SNMP status** (checkboxes): Include devices based on their current SNMP polling status (Detected or Not Detected).
* **Streaming Telemetry Status**: Include devices based on their current Streaming Telemetry status (Detected or Not Detected).
* **BGP status** (checkboxes): Include devices based on their current BGP peering status (Established, Not Established, or Not Configured).
* **Plan** (selection field): Include only devices with the selected billing plan (see **Licenses**).
* **Label** (selection field): Include only devices with the selected labels applied. After clicking the selection field, the drop-down list stays open to select multiple labels. Click off of the list to close it.
* **Vendor** (checkboxes): Filter devices based on the vendor specified as the device's manufacturer (derived via SNMP polling).
* **Type** (checkboxes): Filter devices of a given type (see **Supported Device Types**).
* **Flow Type** (checkboxes): Filter devices based on the format of the flow data: NetFlow v5, v9, IPFIX, or sFlow (see **Flow Protocols**).
* **Credential** (checkboxes): Filter devices based on the credentials used in the device’s settings (see **Credentials Vault**).

### Filter Application Rules

Kentik applies the following rules to filter categories and criteria:

* Criteria are ORed (match any) within categories and ANDed (match all) between categories.
* When criteria are selected in more than one category, a device is displayed only if it matches at least one of the selected criteria in all of the categories with criteria selected.
* When no criteria are selected for a category, the device may match any of the category’s listed criteria.

## Device Settings Dialog

Adding or editing a device via the Kentik portal involves specifying information with the settings on the tabs of the **Device** dialog that opens when you click either the **Add Device** button or the **Edit** icon in the device's row of the **Device** list. These settings are covered in detail in **Device Settings**.

## Manage Devices

Devices are created and edited via the Devices page of the Kentik portal (choose **Settings** from the main menu, then **Networking Devices** under **Data Sources**).

> **Note:** Settings made in the topics below may also require changes to your device configuration (see **Router Configuration** or **Host Configuration**).

### Add a Device

To add (register) a new device:

1. Choose **Settings** from the main Kentik menu.
2. On the **Settings** page, click **Networking Devices**.
3. Click **Add Device** (at upper right) to open the **Device** dialog.
4. On the **General** tab, enter a name in the **Name** field.
5. Choose the **Type** (see **Supported Device Types**), which determines the tab's remaining settings fields.
6. Specify the values of the remaining fields (see **Device Settings Dialog**).
7. Depending on the type of data you'd like Kentik to collect from the device, specify the settings on the other tabs of the **Device Settings Dialog**.
8. Save the new device by clicking the **Add Device** button.

> **Notes:**
>
> * In addition to registering a device, you must also configure the device itself to send data to Kentik (see **Router Configuration** or **Host Configuration**).
> * For assistance with this process, contact Kentik (see **Customer Care**).

#### Add an SNMP-only Device

If you don't need to collect flow from a given device but you'd still like to poll the device for SNMP, you can add an SNMP-only device. To do so, you'll need to be sure that `kproxy` is installed before performing the following steps (see **kproxy Download and Install**):

1. Follow the steps in **Add a Device** to register one or more new devices. For each device, specify **Device SNMP Settings** on the **SNMP** tab of the **Device** settings dialog.
2. On the Devices page, copy the **ID** of each newly created device, which can be found at the top of the **Device Details** pane that opens when you click on a device's row in the table.
3. Run `kproxy` using the argument `-bootstrap_devices` with a comma-separated list of the IDs for the devices created above:  
   `/usr/bin/kproxy -bootstrap_devices 104884,88726`

> **Notes**:
>
> * For information on SNMP polling, see **About SNMP Polling**.
> * For additional information on the different kproxy command line arguments, see **kproxy CLI Reference** and **kproxy Proxy Agent Arguments**.

### Edit a Device

To edit an existing device:

1. Choose **Settings** from the main menu.
2. On the **Settings** page, click **Networking Devices**.
3. In the **Device** list, click the edit icon in the row for the device that you'd like to edit, which will open a **Device** dialog.
4. Edit the settings that you want to change (see **Device Settings Dialog**).
5. Click **Save** to save the settings and exit the dialog.

### Remove a Device

> **Note:** Removing a device from your organization's collection of Kentik-registered devices will remove the device's existing flow records from your Kentik data.

To remove a device:

1. Choose **Settings** from the main menu.
2. On the **Settings** page, click **Networking Devices**.
3. In the **Device** list, click the **Archive** icon (file box with down arrow) in the row for the device that you'd like to remove.
4. In the resulting confirmation dialog, read the explanation to fully understand the implications of archiving a device.
5. To proceed, click the **Archive** button. In the device's row of the **Device** list, the device will be tagged as "Archived" and the **Edit** and **Archive** buttons will be replaced with **Restore** (file box with up arrow) and **Remove** (trash icon) buttons.
6. Click **Remove**.
7. In the resulting confirmation dialog, click **Cancel** to abort the device removal or **Remove** to complete the removal of the device.

### Enable BGP

To enable the collection of BGP data from a device:

1. Choose **Settings** from the main menu.
2. On the **Settings** page, click **Networking Devices**.
3. In the **Device** list, click the edit icon in the row for the device that you'd like to edit, which will open the **Device** dialog.
4. On the **BGP** tab, choose “Peer with Device” from the drop-down **BGP Type** menu.
5. Specify the remaining settings covered in **Device BGP Settings**.
6. Click **Save** to save the new BGP settings and close the dialog.
