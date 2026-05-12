# Logical Map

Source: https://kb.kentik.com/docs/logical-map

---

This article discusses the **Logical Map** feature of the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(276).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A25Z&se=2026-05-12T09%3A35%3A25Z&sr=c&sp=r&sig=ktfPV9ve9bU%2BrA1CbvNGj511GjB8FA43%2FGhOXzjIzsw%3D)

*The Logical Map shows network elements, links, and traffic on your network.*

> **Note:** The primary topology visualization tool in the Kentik portal is the **Kentik Map**.

## About Logical Map

Kentik's **Logical Map** is a topology diagram that graphically shows the devices, sites, and service providers that make up your network environment, the links between them, and the traffic on those links. As a network engineer, this enables you to see some of the most common conditions without having to run queries or know where to drill down, which helps you (and your less-technical colleagues) to understand what's happening on your network in real time.  
  
 Specifically, logical maps enable users to:

* See the existence of traffic between sites and/or devices and (on hover) the amount of traffic carried by a given link.
* Visually identify sites, devices, or interfaces whose traffic volume is higher or lower than expected.
* Visually identify devices that are experiencing hardware failures or resource exhaustion.
* Right-click a network component to open that component in a portal that enables deeper exploration.

Logical maps are made up of two main types of components:

* **Network Elements**: The devices, sites, and service providers that are part of or link to your network (see **Logical Map Elements**).
* **Links**: The connections between network elements (see **Logical Map Links**).

## Logical Map Page UI

The **Logical** **Map** page includes the following main UI elements:

* **Full width** (in the subnav): A toggle button that expands the map to the maximum horizontal space available within the browser window.
* **View Kentik Map** (in the subnav): A button that takes you to the **Kentik Map**.
* **Logical map**: Occupying the bulk of the page's display area, this diagram graphically shows the devices, sites, and service providers that make up your network environment, as well as the links between them (see **Logical Map Diagram**).
* **Refresh**: Located at the top right of the page, this button grayed-out until you've used the display controls to change what's shown in the diagram. The button causes the diagram to be redrawn to optimize the space based on what's currently shown.
* **Save**: Saves the current layout of nodes (sites, network elements, etc.) in the logical map.

  > **Note:** This button does not save the current settings of the **Map Display Controls**.
* **Display controls**: A set of controls (**Show**, **Filter Sites**, **Highlight Flows**), located just below the **Refresh** and **Save** buttons, that enables you to narrow the portion of your network that is represented in the topology diagram (see **Map Display Controls**).
* **Details pane** (present only when an element is selected in the **Logical Map**): Located below the display controls, this pane displays information about the currently selected device, provider, site, or link (see **Logical Map Details Pane**).

## Map Display Controls

The display controls, located at the top right of the page, include the following UI elements:

* **Show Providers**: A drop-down that determines how the map will show service providers:

  + None: Providers will not be shown.
  + By Name: Providers will be shown and labeled by their names.
  + By Connectivity Type: Providers will be shown and labeled by their connectivity type (see **Understanding Connectivity Types**).
* **Show Unlinked:** Toggle to show/hide elements that are not connected to any other element.
* **Show Sites/Devices**: Toggle on/off the display of nodes (e.g., devices) within the sites of your network:

  + When the button is labeled **Devices** then the map displays the individual nodes within your sites.
  + When the button is labeled **Sites** then the sites are displayed without their individual nodes.
* **Filter Sites**: Click in this field for a drop-down list of sites, then click on an individual site to select it for display in the map. Repeat to select multiple sites to display simultaneously. Each currently selected site will appear as a tag in the field. To remove a field, click **X** in its tag.
* **Highlight Flows**: Select a protocol to highlight only the links carrying traffic in that protocol. Only one protocol can be selected at any given time. To remove protocol-based highlighting, choose None.

## Logical Map Diagram

The controls and usage of the **Logical** **Map** are covered in the following topics.

### Logical Map Elements

The following primary types of network elements are depicted in the map:

* **Devices**: Represented by a filled circle.
* **Sites**: Represented by a semi-transparent rectangle.
* **Service providers**: Represented by a transparent circle.
* **Unlinked**: Represented by a filled rectangle. These elements have no connections to an external element (e.g. provider) or to any other Kentik-registered element within your network (e.g. device).

The above elements are each labeled; for large networks you may need to zoom in (see **Navigating the Diagram**) to read a label.

### Logical Map Links

Links are represented by lines that indicate a flow of data between them. The links are arrayed to reflect the topology of the current (selected) element (typically a star topology for routers or a point-to-point topology for hosts).  
  
 The lines used to represent links have the following characteristics:

* **Arrows**: Arrows on links indicate the direction of traffic flow:

  + A link with no arrows indicates that no traffic is flowing between the two elements.
  + A link with only one arrow indicates that traffic flows in only one direction.
  + A link with two arrows, arrows pointing to both elements, indicates bidirectional traffic flow.
* **Segment length**: For links showing bi-directional traffic flow, the links of the two segments roughly indicates the ratio of traffic flow between the two elements.

  + Two roughly equal segments show that the volume of traffic coming from each element is about the same.
  + If the segments are unequal, the length of the segment connected to each element is roughly proportional to the traffic volume coming from that element.
* **Animation**: A link with traffic flow will, when selected, change from a solid line to a dashed line that is animated to show the direction of traffic.

> **Note:** When a link is selected, more detailed information regarding its traffic flows is provided in the**Logical Map Details Pane**.

### Navigating the Diagram

The following controls enable you to more closely examine inspect certain parts of the diagram or to obtain more of an overall view.

* **Zoom in**: Double-click on the part of the diagram on which you want to zoom in.
* **Zoom Out**: Shift-double-click to zoom out.
* **Pan**: Click (outside of a map element) and drag in the direction you want to pan.

### Logical Map Element Actions

The following actions are available for each element:

* **Select**: Click an element to display information about it in the Details pane (see **Logical Map Details Pane**).

  > **Note:** For sites, click the label to select.
* **Move**: Drag an element to a new location (see **Moving Diagram Elements**).

#### Moving Diagram Elements

Depending on the size and complexity of your network, you may wish to “clean up” areas of the diagram by reducing the overlap of links, making it easier to visualize traffic flow. You can click and drag any network element (device, site, or provider) to move it around in the diagram. When you move an element:

* All links to that element will move along with it.
* If you move a device towards a site boundary, the site's rectangle resizes to encompass the moved device.

## Logical Map Details Pane

The **Details** pane displays information about the currently selected device, provider, site, or link. The information displayed depends on the type of the current element:

* **Device**: Details about a router, host or cloud (see **Logical Map Device Details**).
* **Site**: Details about a site (see **Logical Map Site Details**).
* **Provider**: Details about a service provider (see **Logical Map Provider Details**).
* **Unlinked Details**: Details about a device that is not connected to any other device or provider (see **Logical Map Unlinked Details**).
* **Link Details**: Details about a link (see **Logical Map Link Details**).

### Logical Map Details UI

The **Details** pane includes the following main UI elements:

* **Heading**: Located at top left of the pane, the heading indicates the type of details displayed in the pane (see list in **About Logical Map Details**).
* **Name**: The name (under the heading) of the element or link for which the pane is showing details. The name is a link; click it to go to the **Network Explorer** detail view for the current device (see **Core Details Pages**).

  > **Note:** When a link is selected the name field shows the names of both network elements, separated by a hyphen, as well as the link’s connectivity type (see **Understanding Connectivity Types**).
* **Element-specific details**: The details displayed for a given element in the diagram, which depend on the type of the element (see **Element-specific Details**).

### Element-specific Details

The details displayed for a given element in the diagram, which depend on the type of the element, are covered in the topics below.

> **Note:** The **Total Ingress** and **Total Egress** tables described below are present only when there is ingress or egress traffic.

#### Logical Map Device Details

The following details and controls are provided when the selected network element is a device:

* **Total Ingress**: Traffic entering the device, broken out by protocol.
* **Total Egress**: Traffic leaving the device, broken out by protocol.
* **Type**: Type of device (see **Supported Device Types**).
* **Site**: The name of the site to which the device is assigned (see **About Sites**).

#### Logical Map Site Details

The following details and controls are provided when the selected network element is a site:

* **Location**: The location of the site (see **About Sites**).

#### Logical Map Provider Details

The following details and controls are provided when the selected network element is a provider:

* **Total Ingress**: Traffic entering the device, broken out by protocol.
* **Total Egress**: Traffic leaving the device, broken out by protocol.
* **Description**: The name of the site at which this provider's traffic enters and or leaves your network, as well as the name of the provider itself, which comes from **Provider Classification**.

#### Logical Map Unlinked Details

The following details and controls are provided when the selected network element is a device and that device is unlinked (not connected to any other device or provider):

* **Type**: Type of device (see **Supported Device Types**).
* **Site**: The name of the site to which the device is assigned (see **About Sites**).

#### Logical Map Link Details

The following details and controls are provided when a link is selected:

* **Capacity**: The total capacity of the link (Mbps).
* **Source**: Information about the interfaces used for inbound traffic. The following is provided for each interface on which you receive traffic sent to your network over this link:

  + Device: The name of the source device. Click the device’s name to display its details page (see **Core Details Pages**).
  + Interface: The name (string defined in the device itself and retrieved via SNMP) of the interface from which traffic is being sent over this link. Click the interface's name to display its details view.
  + IP: The IP address at which traffic is being received on this link. Click the address to display its details view.
* **Destination**: Information about the interfaces used for outbound traffic. Device, Interface, and IP (defined above) are provided for each interface on which you send traffic from your network over this link.
* **Volume**: The volume of traffic on the link, shown separately in each direction and broken down by protocol. The percentage value beside the protocol name shows the proportion of traffic based on the protocol, while the value on the right shows bits per second (Mbps).
* **View Flows**: For each direction of the link, the **View Flows** button takes you to Data Explorer, with traffic filtered to show that direction's traffic on the link.

---

© 2014-25 Kentik
