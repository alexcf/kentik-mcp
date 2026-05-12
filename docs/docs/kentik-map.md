# Kentik Map

Source: https://kb.kentik.com/docs/kentik-map

---

This article discusses the **Kentik Map** module in the Kentik portal.

> **Note:** For information related to the contents of the **Details** drawer of the Kentik Map, see **Kentik Map Details**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(432).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A52Z&se=2026-05-12T09%3A58%3A52Z&sr=c&sp=r&sig=aK9HyBKT4hQmTy20PnTHGgtU9JjObd2PkWAWxC6EW%2Bg%3D)

*The Kentik Map shows the on-prem, cloud, and Internet elements that handle your network's traffic.*

## About Kentik Map

|  |  |
| --- | --- |
| **Purpose:** | Visualize every aspect of network infrastructure, both on-prem and cloud, to enable fast understanding of how components are interconnected and how that affects traffic patterns, network health, and performance, including application delivery and customer experience. |
| **Benefits:** | * Unified view into traffic, performance, and health between and within cloud, on-prem, internet and WAN networks. * See connections between on-prem networks and VPCs, as well as between different cloud providers, to understand patterns, investigate problems, discover application dependencies, and reveal unintended Internet traffic. * Get insight into expensive, brittle, or bandwidth-constrained flow connections. * Answer questions about traffic in and between any environment, to and through ASNs, and out to the Internet, as well as north/south and east/west flows in data centers. * Identify at a glance links that are down or interfaces whose health or utilization are in a critical state, then identify potential causes with a few clicks. |
| **Use Cases:** | * Hybrid network architecture visualization and mapping * Network health visualization * Network traffic visualization |
| **Relevant Roles:** | Network Admin/Engineer, Network Architect, Site Reliability Engineering (Traffic Engineer, Net SRE, NetOps Engineer) |

The Kentik Map module of the Kentik portal illustrates the relationship between three main aspects of your Hybrid IT network infrastructure:

* **Clouds**: The cloud providers you use for compute and/or storage:

  + AWS: See **Kentik for AWS**.
  + Azure: See **Kentik for Azure**.
  + GCP: See **Kentik for GCP**.
  + OCI: See **Kentik for OCI**.
* **Internet**: The external sources and destinations of traffic to and from your network are broken down by Origin Network, Next Hop Network, and Provider (see **Provider Classification**).
* **On Prem**: The sites where your data center infrastructure is located (see **About Sites**).

In the main (top-level) Kentik Map, each of the above areas of your network is represented as a grey rectangle, referred to as a "block," in which you can drill down to get further details about the network's structure and traffic. This helps network engineers better understand real-time network activity, compare current and historical traffic for specific network entities, and identify common adverse conditions without having to run queries.

### Kentik Map Views

The Kentik Map views represent the structure of your overall network, the components within that structure, and the traffic between those components. These views enable you to quickly drill down into your infrastructure at any level, where you can see information including the following:

* **Traffic Volume**: View existing traffic volume between your sites, such as data centers and branch offices (**On Prem** block), your clouds (**Clouds** block), and external networks (**Internet** block).
* **On-prem**: The sites that make up your on-prem infrastructure, including:

  + Weather Map: The location of sites on a zoomable world map, with multiple layers showing information such as links, utilization, and health.
  + Topology: The architecture and individual devices of each site, as well as details about the traffic on the individual network entities or between entities.
* **Network Health**: Assess the health of network entities their connections.
* **Cloud providers**: View region breakdowns for each cloud provider.
* **Top ASNs and Service Providers**: Identify top ASNs (origin and next-op) and service providers (transit, peering, or IX) that account for your network's incoming and outgoing traffic.

### Kentik Map Health

The Kentik Map provides real-time network health indicators for network infrastructure, monitored via SNMP polling of network device and interface metrics. Metrics are assessed against expected ranges to classify health status as "healthy," "warning," or "critical." Categories include:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(435).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A52Z&se=2026-05-12T09%3A58%3A52Z&sr=c&sp=r&sig=aK9HyBKT4hQmTy20PnTHGgtU9JjObd2PkWAWxC6EW%2Bg%3D)

* **Site Health** - Based on devices and interfaces within the site.
* **Device Health** - Evaluated from device metrics and interface metrics:

  + **Device Availability**: Checks if metrics are available for this device.
  + **Device Metrics**: Monitors CPU and memory utilization.
* **Interface Health** - Assessed from interface’s metrics:

  + **Interface Availability**: Checks if metrics are available for this interface.
  + **Interface Metrics**: Monitors input and output interface utilizations.

> **Notes:**
>
> * In order to be evaluated for health, your devices must allow SNMP polling from Kentik (see **Enabling SNMP Polling**), with a “Standard” polling interval (see **SNMP Polling Intervals**).
> * The health status indicators in the Network Map UI reflect the state each element at the time the map is accessed.

### Kentik Map Prerequisites

To effectively use Kentik Maps, ensure that your overall Kentik setup is as complete as possible:

* **Register all devices**: Physical devices (e.g., routers and switches) and host agents must be registered with Kentik in order for us to receive traffic data:

  + To register devices via the Kentik **Onboarding** wizard, see **Device Setup**.
  + To register via the portal's device admin UI, see **Add a Device**.
* **Configure SNMP on all devices**: SNMP polling enables Kentik to enrich traffic data (flow records) with data about the interfaces via which traffic is entering, transiting, and leaving your network:

  + To activate SNMP on a device via the Kentik **Onboarding** wizard, see **Device SNMP Setup**.
  + To activate SNMP via the portal's device admin UI, enable polling on the device with the device-specific SNMP configs provided in our **Device Configs Directory**, and also set the device's **Device SNMP Settings**.
  > **Note:** To take advantage of the health status feature of Kentik maps, set the SNMP polling interval to Standard (see **SNMP Polling Intervals**).
* **Register all clouds**: A cloud export in Kentik represents one or more cloud resources (e.g., VPCs or subnets) used by your network on a given cloud provider (e.g., AWS, GCP, Azure). To gain visibility into those resources you must register them in Kentik (see **Cloud Overview**).
* **Assign data sources to sites**: A site is a specific user-defined physical location (e.g., the address of a data center) to which one or more data sources (devices, hosts, or clouds) may be assigned (see **About Sites**).
* **Exclude interfaces from the map**: Your organization may have interfaces that, to reduce clutter and improve clarity, you don't want to show up on the map. Before running interface classification, you can set an IC rule that excludes these interfaces (see Exclude Interfaces from Map).
* **Run Interface Classification** (see **Using Interface Classification**): Interface Classification assigns a Network Boundary and Connectivity Type value to every interface in the network:

  + **Network Boundary**: Classifies interfaces as Internal or External, to determine if the source and destination of the traffic are both fully within your network or if the traffic crossed a network boundary (came from or went to a different AS; see **Network Boundary Attribute**).
  + **Connectivity Type**: Classifies interfaces by their role in the overall network (see **Connectivity Type Attribute**), such as Transit, IX, or Cloud Interconnect, etc. (see **Understanding Connectivity Types**).
  > **Note:** Links are drawn between two sites in the **On Prem** block if the interfaces that connect them are assigned a Connectivity Type of either Backbone or Datacenter Interconnect.

#### Exclude Interfaces from Map

To exclude interfaces from the Kentik Map:

1. Go to Settings » Interface Classification, and click **Add Rule** to open the **Add Rule** dialog (see **Rule Dialogs UI**).
2. In the rule's **IF** settings, define the conditions to match any interfaces for exclusion from the map.
3. In the rule's **THEN** settings:

   1. Set the **Connectivity Type** to **Host/Access**.
   2. Ensure **Network Boundary** is set to **Internal**, (turn off **Auto** if needed).
4. Click **Add Rule**. The dialog will close and the new rule will appear in the **Rules** list, showing that **Host/Access | Internal**, and matching interfaces will be excluded from the map.

## Kentik Map Page

The **Kentik Map** page includes elements common to many page types and others that are found only on the pages for **Topology Views**.

### Common Map Elements

The following elements are common to most Kentik Map pages:

* **Breadcrumbs** (subnav): Indicates of your current location within the Kentik portal. As you drill down deeper you can click on a breadcrumb to go back to a higher level.
* **Full width** (subnav): Expands the map to the maximum horizontal space available within the browser window.
* **View Kube**: Links to **Kentik Kube**for Kubernetes clusters (not present on **Cloud Topology** views). Activated by request to Kentik; see **Customer Care**.

  > **Note**: If you don’t see this option in your map, Kentik Kube might not be enabled for your company’s account. Please contact your Kentik Account Team.
* **View Logical Map**: Links to the **Logical Map**, formerly known as the Network Map (not present on cloud topology views).
* **Details** (subnav): Toggles visibility of the right-side **Details** drawer, which contains details about the currently selected map element (see **Kentik Map Elements**). If no element is selected this button is inactive.
* **Filters**: Opens a popup to view or edit the filters applied to the displayed data. The popup’s elements include:

  + Cards: Existing filters appear as a cards. Remove filter(s) using the red **X** at the right of its card.
  + Apply: Applies any changes made in the **Filters** dialog to the Kentik Map.
  + Edit Filters: Add new filters or edit existing ones using the **Edit Filters** button (opens a **Filtering Options Dialog**). Once your changes are saved on the **Filters** dialog, click **Apply** in the popup.

* **Time range**: Indicates the current time range of the data displayed on this page and when clicked, pops up a calendar where you can specify the time range (see **Time Range Control**).
* **Health indicator** (heart icon): Indicates the number of issues identified by Kentik on your network. Click to open a popup giving a breakdown of different problem types (e.g. high inbound utilization, high outbound utilization, high memory utilization, etc.). Click **View Problems** to go to the **Health Problems Page**.
* **Gbits/s legend**: A horizontal bar showing traffic bit rates with colors representing different networks in the **Internet** block (see **Kentik Map Blocks**).
* **Map display area**: The main section displaying area your network environment, including network entities their connections (see **Kentik Map Display**).

### Topology View Elements

These elements are specific to **Topology Views**:

* **Search**: Filters **Topology Views** based on the entered text (see **Topology Search Controls**).
* **Color by** (**AWS Topology** only): Represents AWS VPS traffic volume by color intensity (greater intensity indicates greater volume), based on bits/second inbound, outbound, or total bits/second.
* **Show Default VPCs and Subnets** (**AWS Topology** only): A checkbox to toggle the display of default AWS VPCs and subnets.
* **Show Default Networks** (**GCP Topology** only): A checkbox to toggle the display of default GCP networks, which are automatically provisioned virtual networks provided to new projects unless user-disabled.
* **Subscriptions** (**Azure Topology** only): A dropdown to filter the topology views by selected subscriptions (see **Azure Subscription Filter**).

#### Topology Search Controls

The **Search** control set in the **Topology Search Controls** enables you to filter and save searches in **Topology Views**. UI elements include:

* **Search by**: Enter one or more comma-separated values to match entities:

  + Click the field to open the **Search** popup and enter your search criteria.
  + Click the **X** at the right of the field to clear the search and close the popup.
* **Search Entries**: A pane displaying a categorized list of matching entities, varying by cloud provider (see **Topology Search Entries**).
* **Name**: Enter a name to save the current search.
* **Save**: Saves the current search.
* **Saved Searches**: A list of previously saved searches (see **Topology Saved Searches**).

#### Topology Search Entries

The **Search Entries** pane of the **Topology Search Controls** includes a table of exact matches for entries across various categories depending on the cloud provider:

| Cloud Provider | Account IDs | IDs | CIDRs | Tags | Tenancy IDs | Compartment IDs |
| --- | --- | --- | --- | --- | --- | --- |
| AWS | Yes | Yes | Yes | Yes | No | No |
| Azure | Yes | Yes | Yes | Yes | No | No |
| GCP | Yes | Yes | Yes | Yes | No | No |
| OCI | Yes | Yes | Yes | Yes | Yes | Yes |

The pane includes the following UI elements for each category:

* **Search Entries**: Displays a lozenge for each match, which contains the matched string. Click the **X** on a lozenge to remove it from the **Search** field.

  > **Note:** Lozenge colors are randomly assigned.
* **Clear**: A button to clear all strings in the category.

#### Topology Saved Searches

The **Saved Searches** pane of **Topology Search Controls** includes a table that lists all of your previously save topology view searches, any of which can be used for any cloud provider. Click on a listed search to use it to filter the current topology view.  
  
 The **Saved Searches** table contains the following UI elements:

* **Name**: Displays the name of the saved search. Click to apply it.
* **Edit** (edit icon): A button that allows you to modify the saved search’s name.
* **Remove** (trash icon): A button that opens a confirmation dialog to remove the saved search.

## Kentik Map Display

The main parts of the Kentik map display area are covered in the following topics.

### Kentik Map Elements

A Kentik map diagram consists of the following main components:

* **Blocks**: Gray rectangles that each enclose different buckets of infrastructure (see **Kentik Map Blocks**).
* **Network Entities**: Individual network components within blocks:

  + **Physical entities**: The sites and data sources (devices and hosts) in your data centers.

    > **Note:** Physical entities are each marked with a health indicator; see **Element Health Indicators**.
  + **Virtual entities**: The regions, gateways, VPCs, and subnets in your clouds.
  + **Logical entities**: The ASes and service providers for network traffic connections.
  > **Note:** Click an entity to either open the **Details** drawer for that entity (see **Kentik Map Details**) or pop up a menu listing possible **Network Element Actions**.
* **Links**: Lines showing connections between blocks and network entities, with arrows indicating traffic direction. Click a link to view traffic details in the **Details** sidebar.
* **Weather Map**: The default view for the **On Prem** block (see **Kentik Weather Map**).
* **Topology**: Shows the internal architecture and relationships of map elements. **Topology Views** are accessed via the **View Topology** option in a **Details** popup (click any element to see the popup). Topology views are available for the following entities:

  + **On Prem**: Shows the interconnections of on-premises network resources, allowing you to select individual sites to see their internal topology.
  + **Site**: Shows device architecture within a site (see **Site Topology**).
  + **Devices**: Illustrates device relationships and interface connevtions (see **Device Topology**).
  + **Cloud provider**: Displays regions and counts of active VPCs and instances (AWS, Azure, OCI) or subnets and VMs (GCP).

> **Note:** In the **AWS Topology**, **Azure Topology**, and **OCI Topology** views, each VPC is represented as an expandable card.

#### Element Health Indicators

If health status information (see **Kentik Map Health**) is available for a physical entity (a site, device, or interface), then that element’s overall health status will be marked in topology views as follows:

* **Critical** (red): One or more metrics are in the critical range.
* **Warning** (orange): One or more metrics are out of the normal range, but not critical.
* **Healthy** (green disc): All metrics are within normal ranges.
* **Unknown** (gray): Health status is unknown (e.g. SNMP polling is not configured).

Further detail about the health status of a given element is available in the following locations:

* On the **Health** tab in the element’s **Details** popup (see **Health Details**).
* On the **Health Problems Page** for warning or critical statuses.

#### Network Element Actions

Clicking a network element in the On Prem Topology view pops up a menu with these options:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KM-Onprem_links-590h594w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A52Z&se=2026-05-12T09%3A58%3A52Z&sr=c&sp=r&sig=aK9HyBKT4hQmTy20PnTHGgtU9JjObd2PkWAWxC6EW%2Bg%3D)**View Topology**: Access the topology view for the element (not present for logical entities; see **Topology Views**).
* **Show Details**: View element information in the **Details** drawer.
* **Show Connections**: Display traffic lines between the entities and blocks. In AWS maps, it also draws connections between subnets and gateways within a VPC.
* **Show in AWS Console** (**AWS Topology** view): Open the selected element in the AWS console to make configuration changes.
* **Show in Azure Console** (**Azure Topology** view): Open the selected element in the Azure console to make configuration changes.
* **Show Path To** (AWS and Azure subnets): Displays the traffic path from a subnet to its destination, with directional arrows and displays path details when you hover over it.

> **Note:** In the **AWS Topology** view, clicking on Direct Connection, Customer Gateway, Direct Connect Gateway, or VPN Connection opens the Details drawer directly. See **AWS Interconnection Elements** and **Kentik Map Details**.

### Kentik Map Blocks

Kentik map blocks categorize network entities as follows:

* **Clouds**: Contains for cloud providers (AWS, Azure, GCP, OCI).
* **Internet**: Contains external traffic sources and destinations (ASNs and service providers).
* **On Prem**: Contains one of the following:

  + **Weather Map** (default): A zoomable world map showing site locations with separate layers for links, utilization, and health.
  + **Topology**: Displays sites with data center infrastructure (see **About Sites**).
* **Site** (site topology map): Contains devices within a specific site.
* **Other Sites** (site topology map): Contains sites other than the site shown in the **Site** block.

### Inter-block Traffic Volume

Traffic volume between blocks on the main (top-level) Kentik map is calculated as follows:

* **On-Prem » Internet**: Includes all flows leaving on-prem infrastructure via External Network Boundary (see **Interface Classification Dimensions**), excluding:

  + Cloud Interconnect flows (see **Understanding Connectivity Types**);
  + Flows with a Traffic Profile of From Inside to Cloud (see **Network Classification Dimensions**).
* **Internet » On-Prem**: Includes flows entering on-prem infrastructure via an External Network Boundary excluding:

  + Cloud Interconnect flows;
  + Flows wotjTraffic Profile of From Cloud to Inside.
* **On-Prem » Cloud**: Includes the following traffic leaving on-prem:

  + Cloud Interconnect flows;
  + Flows with a Traffic Profile of From Inside to Cloud.
* **Cloud » On-Prem**: Includes the following traffic entering on-prem:

  + Cloud Interconnect flows;
  + Flows with a Traffic Profile of From Cloud to Inside.
* **Cloud » Internet**: Includes all flows leaving any cloud with a Traffic Profile of From Cloud to Outside.
* **Internet » Cloud**: Includes all flows entering any cloud with a Traffic Profile of From Outside to Cloud.
* **Regions » On Prem**: For Azure, GCP, and OCI, a line is drawn between the Cloud Region Block and the On Prem block.

  > **Note:** These lines do not display data rates.
* **Regions » Other Clouds**: Visualizations for Azure, GCP, and OCI include a line between the selected cloud’s region block and Other Clouds.

  > **Note:** These lines do not display data rates.

### Kentik Weather Map

The weather map is the default view of the **On Prem** block on the Kentik Map landing page. This zoomable world map shows the location of network entities — typically sites (see **About Sites**) or cloud regions — as well as the links between them. Set the type of the links using the **Links** settings in the **On Prem Controls**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(448).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A52Z&se=2026-05-12T09%3A58%3A52Z&sr=c&sp=r&sig=aK9HyBKT4hQmTy20PnTHGgtU9JjObd2PkWAWxC6EW%2Bg%3D)

*The weather map shows the physical location of network entities.*

#### Weather Map Clustering

By default, the number of entities in each location on the map determines how the entities are displayed:

* **Clustered marker**: Represented by a circled number indicating multiple nearby entities.

  + If the **Health** switch is on in the **On Prem Controls**, a marker’s circumference shows health segments for each entity.
  + As you zoom further in on the map (either with your scroll wheel or using the zoom buttons; see **On Prem Controls**), multi-site markers may split off into their own separate markers.
  + Hover over the marker to open a popup with information about the entities (see **Weather Map Popups**).
  + Click a clustered marker to open a **Kentik Map Details** drawer about the entities it contains.
* **Individual marker**: A circle with a label indicating the name of a site or cloud region.

  + Click an individual marker to see information about the site in **Details** drawer (not available for cloud regions).

> **Note:** Turn clustering on/off in the **On Prem Controls**.

#### Weather Map Popups

The **Weather Map** features the following popups that open upon hover:

* **Links**: If the **Links** switch is On in the **On Prem Controls**, links will be drawn between the entities shown on the map. Hover over a link to display the traffic volume in each direction, expressed in bps and also (if the **Utilization** switch is On in the **On Prem Controls**) as a percentage of capacity.
* **Entities**: If the **Clustering** switch is On in the **On Prem Controls**, hovering over a clustered marker opens a popup shows:

  + **Health**: If enabled in the **On Prem Controls**, the popup will contain a breakdown of the health status of the clustered entities.
  + **Sites**: A section listing the names of the clustered entities that are sites.
  + **Regions**: If enabled in the **On Prem Controls**, the popup will contain a section listing the names of the clustered entities that are regions.

### On Prem Controls

The On Prem controls, accessible from the Layers icon in the **On Prem** block, configure the **Weather Map** and **Topology** views.

* **Links**: Controls how lines are drawn between entities.

  + **Traffic Type**: Selects connection types to render (see **Traffic Type Options**).
  + **Draw Connections Using**: Chooses traffic layer for rendering links (All Layers, Layer 3, or Layer 2).
  > **Note:** Links are solid lines if healthy, or a dashed yellow line if its health is degraded or not currently working.
* **Cloud Regions**: Determines whether cloud regions are rendered.

  + **Cloud Backbone Traffic**: Enable or disable links between cloud provider regions.
* **Utilization**: Turn on/off coloring of links in the **Weather Map** based on utilization:

  + **On**: Link colors are based on the utilization of the link (traffic as a percent of capacity), and a link legend (Utilization & Traffic) displays at the bottom of the map.
  + **Off**: All links will be rendered as blue lines.
  > **Note:** Capacity information isn’t available for inter-region cloud links.
* **Health**: If on, displays segments on each clustered marker’s circumference that indicate the health of one of the represented entities, and a health legend displays at the bottom of the map.
* **Clustering**: Reduces clutter on the **Weather Map** by using a single marker for multiple nearby entities (see **Weather Map Clustering**).

#### Traffic Type Options

The **Traffic Type** settings determine how lines appear between sites in the **On Prem** block:

* **Connected Interfaces**: Connects sites based on interfaces with IPs in the same subnet at different sites that are configured with IP addresses inside of the same subnet. For example, interface 1 in site A with IP address 192.168.1.2 will be connected to interface 2 in site B configured with an IP address of 192.168.1.3. You can choose whether the connections are based on all layers (default), Layer 2 only, or Layer 3 only.
* **Site IP**: Connects sites by querying traffic from each site’s IP range to others. The results of the queries are used to determine the traffic volume between sites. Connections are drawn as arrow links.
* **Ultimate Exit**: Connects sites by querying traffic using Kentik’s BGP Ultimate Exit feature. Each site connected to another site represents a volume of traffic ingressing a source site and egressing the ultimate exit site.

## Topology Views

> **Note:** The topology view for AWS is distinct from that of other cloud providers.

Topology views illustrate the relationships between sites, devices, and interfaces in your on-prem infrastructure, as well as cloud resources across providers. They are available for:

* **Site**: Shows the architecture of the data sources within each site (see **Site Topology**).
* **Devices**: Shows device relationships and interface connections to other devices (see **Device Topology**).
* **Cloud provide**r: Displays regions within a provider, and counts of active VPCs and instances (AWS, Azure, OCI) or subnets and VMs (GCP).

The UI elements in topology views are similar to those of the main Kentik map, as detailed in **Kentik Map Page**.

### Site Architecture

To define a site’s architecture for a meaningful on-prem topology view in Kentik, follow these steps:

1. From the main Kentik Map, click on a site to open the site’s **Details** drawer.
2. Click the **View Topology** button, which will take you to the site’s topology view.
3. Assuming that the site architecture hasn’t already been defined, the block for the site will show a number of devices under the heading **Unassigned**, beneath which is a **Configure Site** link. Click the link, which will open the **Edit Site** dialog.
4. In the **Type** section (below the address field), click the button that most closely corresponds to your overall concept of how the site is organized (e.g., Data Center, Cloud, etc.).
5. In the **Architecture** section, click the **Edit Architecture** button, which will open the **Edit Site Architecture** dialog.
6. The dialog contains multiple tabs, each of which provides modifiable templates for different “typical” architectures. Click on the template closest to your situation (or choose **Custom** from the **Other** tab), which will open the **Architecture Edit UI** for that template.
7. Use the architecture edit UI to assign devices to the layers and to rename each layer as needed, then click the **Save Architecture** button to save your changes.

#### Architecture Edit UI

The **Edit Site Architecture** dialog enables you to customize a site architecture template to the specifics of your site. The dialog includes the following UI elements:

* **Close** (**X**): Close the dialog without saving any changes to the architecture.
* **Layers**: Represents a layer in the architecture and contains the fields described in **Layer Fields**.
* **Add Layer**: Adds a box for a new layer between existing layers.
* **Add Parallel Laye**r: Adds a box for a new layer that at the same level as an existing layer.
* **Cancel**: Closes the dialog without saving any changes to the architecture.
* **Save Architecture**: Saves all changes to the architecture and closes the dialog.

#### Layer Fields

Each layer of the architecture is represented as a box containing the following fields:

* **Layer**: The name of the layer.
* **Devices**: Select one or more devices for the layer from a drop-down list of the Kentik-registered devices assigned to this site.
* **Handle**: Drag to reorder layers.
* **Remove** (trash icon): Remove this layer from the architecture.

### Site Topology

After specifying a site’s architecture (see **Site Architecture**), its layers and device relationships are displayed in the site’s topology view. Access this by clicking the **View Site Topology** button in the site’s **Details** drawer (see **Details Panes**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(451).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A52Z&se=2026-05-12T09%3A58%3A52Z&sr=c&sp=r&sig=aK9HyBKT4hQmTy20PnTHGgtU9JjObd2PkWAWxC6EW%2Bg%3D)

#### Site Topology Blocks

The topology view includes the following blocks (see **Kentik Map Blocks**):

* **Site**: Shows the site as a block similar to the **On Prem** block in the Kentik Map view (see **Site Block**).
* **Other Sites**: Lists all other sites in your infrastructure.
* **Clouds**: Shows your cloud providers (AWS, GCP, Azure, and OCI).
* **Internet**: Displays external traffic sources and destinations (ASNs and service providers).

#### Site Block

The site block is structured according to the layers defined in the site architecture:

* Each device is represented by a labeled icon in its assigned layer:

  + Click the icon to open the device’s **Details** drawer (see **Kentik Map Details**).
  + Click the **View Device Topology** button in the drawer to see the **Device Topology**.
* Lines indicate links between connected devices; hover over a device to highlight all of its links.
* An **Unassigned** section lists devices that not yet assigned to a layer. Use the **Configure Site** link to assign them (see **Site Architecture**).

### Device Topology

The device topology view is organized into the following blocks:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(452)(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A52Z&se=2026-05-12T09%3A58%3A52Z&sr=c&sp=r&sig=aK9HyBKT4hQmTy20PnTHGgtU9JjObd2PkWAWxC6EW%2Bg%3D)**Upstream Connected Devices**: Devices in the same site that are connected to this device and assigned to a higher layer.
* **Parallel Connected Devices**: Devices in the same site that are connected to this device and assigned to the same layer.
* **Device**: Displays information about the device (see **Device Block Information**) and its interfaces (see **Device Block Interfaces**). Click the device name to see the Network Explorer Details page for this device.
* **Downstream Connected Devices**: Devices in the same site, connected to this device and assigned to a lower layer.

#### Device Block Information

The left side of the device block provides the following information related to the main device of the topology view:

* **Status**: The device health status (see **Element Health Indicators**).
* **IP Address**: The IP from which this device sends flow to Kentik.
* **Site**: The site where the device is located (click to see the Network Explorer Details page for this site).
* **Sample rate**: The flow sampling rate for the device (see **Flow Sampling**).
* **Machine Type**: The device type (e.g., router, host, etc.).
* **Device ID**: The device’s Kentik-assigned ID.
* **Metrics**: Device metrics gathered via SNMP (see **Device Metrics Information**).

#### Device Metrics Information

The **Metrics** section of the device block includes the following information (gathered via SNMP) and controls:

* **View Details**: Opens a detailed chart view of device metrics.
* **CPU Utilization**: A chart showing CPU utilization for the last 24 hours, including the peak value.
* **Memory Utilization**: A chart showing memory utilization for the last 24 hours, including the peak value.

#### Device Block Interfaces

The main device block area displays the total number of known interfaces on the device and provides a breakdown of those interfaces based on the device layer to which those interfaces connect:

* **Upstream Connected Interfaces**: Shows interfaces linked to higher-layer devices.
* **Parallel Connected Interfaces**: Shows interfaces connected same-layer devices.
* **Unknown Connected Interfaces**: Shows interfaces that are:

  + Unmonitored by Kentik.
  + Part of a logical bundle (multiple physical interfaces defined as a single logical interface).
  + Connected to a Layer 2 device.
* **Downstream Connected Interfaces**: Shows interfaces connected to a lower-level devices.

### Cloud Topology

The Cloud topology view is consistent from the provider level down to subnets, organized into blocks: **Common Cloud Topology Blocks** and **Level-specific Topology Blocks**.

> **Notes:** Cloud topology views are covered in these provider-specific topics:  
>  - **AWS Topology**  
>  - **Azure Topology**  
>  - **GCP Topology**  
>  - **OCI Topology**

#### Common Cloud Topology Blocks

The following blocks appear in cloud topology views at all levels:

* **On Prem**: Represents on-premises infrastructure connected to the cloud provider.
* **Other Clouds**: Shows other cloud providers registered with Kentik, and any traffic links between them.
* **Internet**: Shows the external traffic sources and destinations (origin networks, service providers, and next-hop networks).

#### Level-specific Topology Blocks

These blocks appear only in cloud topology views at the indicated levels:

* **Regions** (cloud provider level only): Displays regions with resources, showing the number of VPCs and subnets. Click a region and choose **View Topology** to see its detailed topology.
* **VPC**: Displays VPCs within the region, showing the number of subnets and VMs. Click a VPC and choose **View Topology** to view its detailed topology.
* **Subnets** (region level in GCP): Displays subnets within a VPC or region. Click a subnet and select **Show Details** for traffic data details.

> **Note**: Shared VPCs, allowing centralized management of resources across multiple accounts/projects, are only supported for AWS and GCP. See the GCP docs on **Shared VPC** for details.

### AWS Topology

The **AWS Topology** view is organized into the following blocks:

* **On Prem**: Represents on-premises infrastructure connected to AWS resources.
* **Internet**: Shows external traffic sources and destinations (origin networks, service providers, and next-hop networks).
* **Regions** (cloud provider level only): Blocks representing regions with AWS resources. Each block shows:

  + VPCs in the region (VPCs are expanded unless more than nine row are present).
  + Transit Gateway for traffic entry/exit.
  + Lines indicating inter-region traffic volumes.
* **VPCs**: Click to expand and view subnets and the VPC connections (see **AWS Interconnection Elements**) for that VPC. Display of VPCs varies in the following ways:

  + **Expanded**: Labeled card showing VPC name, ID, and configured CIDR block.
  + **Collapsed**: Square with color intensity indicating status (see **Color** **by** control in **Kentik Map Page**). Hover to see the name, ID, and CIDR.
* **Subnets**: Cards showing the subnet name and IP /CIDR, grouped by Availability Zones (AZ), shown as dashed outlines around each subnet. An AZ represents a physically isolated datacenter in Amazon’s ecosystem. Click a subnet and choose **Show Details** to for traffic details.
* **Connections**: Various connection types within AWS and between AWS on-premises infrastructure (see **AWS Interconnection Elements**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(453).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A52Z&se=2026-05-12T09%3A58%3A52Z&sr=c&sp=r&sig=aK9HyBKT4hQmTy20PnTHGgtU9JjObd2PkWAWxC6EW%2Bg%3D)

#### AWS Interconnection Elements

The AWS topology view includes these interconnection elements:

* **Customer Gateway**: Terminates one or more site-to-site VPN connections extended from VPC virtual gateways.
* **Link Aggregation Group**: A logical interface that uses the Link Aggregation Control Protocol (LACP) to combine multiple connections at an AWS Direct Connect endpoint, for unified management (see **Link aggregation groups**in AWS docs).
* **Direct connection**: Entry/exit point for traffic to/from the On Prem block that transits an AWS Direct Connect circuit. Kentik can visualize AWS Direct Connects whose virtual interfaces extend from a given VPC directly through to an on-prem router or those connected to Transit Gateway routing devices.
* **Direct Connect Gateway**: Aggregates one or more direct connect circuits and allows for easy connectivity between VPCs and multiple on-prem connections (see AWS **Direct Connect** docs).
* **VPN Connection**: Links Amazon VPCs to remote networks and users (see AWS **VPN connections** docs).
* **VPC connection**: When expanded, a VPC displays its connection gateways as labeled squares across the bottom of the VPC block. Supported types include internet gateway, peering connection, virtual gateway, TGW attachment, NAT gateway, VPC endpoint interface, and transit gateway.
* **Transit Gateway**: An AWS-managed regional network transit hub for high availability and scalability, interconnecting VPCs and customer networks (see AWS **Transit Gateway** docs).

> **Note**: Traffic links between these interconnections are shown by default, except for Internet, NAT Gateway, and Virtual Gateway links, which appear only when **Show Connections** is chosen from the drop-down **Network Element Actions** menu.

### Azure Topology

The Azure topology view is organized into the following blocks:

* **On Prem**: Represents your on-premises infrastructure connected Azure resources.
* **VPN Sites**: Represents VPN Site links connected to your cloud environment.
* **Internet**: Shows external traffic sources and destinations (origin networks, service providers, and next-hop networks).
* **Regions**: Shows the regions, each represented as a block, where you have resources within this cloud provider. Each region block shows the VNets in that region. Click the **Show Details** link to open the region’s **Details** drawer (see **Kentik Map Details**).
* **VNets**: Virtual Networks (VNets) in each region are initially collapsed. Click a VNet to expand and view its subnets. The display of the VNets varies as follows:

  + **Collapsed**: Displays VNet name, ResourceId, and CIDR.
  + **Expanded**: Shows subnets, interconnection elements (gateways, etc.), and a **Show Details** link. Other VNets appear as labeled cards showing VNet name, ResourceId, and CIDR.

* **Subnets**: The subnets of a VNet are each represented as a rectangle giving the subnet name and IP/CIDR. Click the subnet to open its **Details** drawer.
* **VNet connections**: When expanded, VNet interconnections appear as labeled rectangles across the bottom of the VNet block (see **Azure Interconnection Elements**).

> **Notes:**  
>  - For an overview of Azure networks, see **Azure networking services overview**.  
>  - For details on Azure VNets, see Microsoft documentation at **What is Azure Virtual Network**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(454).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A52Z&se=2026-05-12T09%3A58%3A52Z&sr=c&sp=r&sig=aK9HyBKT4hQmTy20PnTHGgtU9JjObd2PkWAWxC6EW%2Bg%3D)

*A region, showing VNets and interconnection elements, in an Azure topology view.*

#### Azure Interconnection Elements

The Azure topology view includes the following interconnection elements:

* **VNet Gateway**: A virtual network (VNet) gateway is composed of two or more virtual machines (VMs) automatically configured and deployed to a specific subnet you create called the gateway subnet. The gateway VMs contain routing tables and run specific gateway services. Direct VM configuration isn’t possible, but gateway settings impact VM creation. See Microsoft documentation at **What is a virtual network gateway**.
* **VNet Peering**: Virtual network peering connects two or more VNets, making them appear as one for connectivity purposes. Traffic between VMs in peered VNets uses Microsoft’s private network only. See Microsoft documentation at **Virtual network peering**.
* **NAT Gateway**: A fully managed Network Address Translation (NAT) service simplifies outbound Internet connectivity for virtual networks. Configured on a subnet, it uses static public IPs for all outbound traffic. See Microsoft documentation at **What is Virtual Network NAT**.

> **Note:** For a comparison of Azure peering and gateway options, see Microsoft documentation at **VNet peering and VPN gateways**.

#### Azure Subscription Filter

The Azure **Subscription Filter** is a dropdown tool for managing the Azure subscriptions displayed in the topology view (check to display, uncheck to hide). It contains the following UI elements:

* **Search**: Filters the **Azure Subscription List** by name, ID, or CIDR.
* **View**: Shows all or selected subscriptions.
* **Group**: Filters subscriptions by selected group.
* **Select All** (checkbox): Toggles the selection of all subscriptions, with a count of selected subscriptions compared to the total available.
* **Create Group**: Opens the **Create Group Dialog**.
* **Clear**: Clears the subscription selections.
* **Subscriptions List**: Displays Azure subscriptions for filtering the topology view (see **Azure Subscription List**).

#### Create Group Dialog

The **Create Group** dialog allows you to form a group from any subscriptions selected in the **Subscription** list and includes the following UI elements:

* **Name**: A text field for the group name.
* **Count**: The number of subscriptions selected to include in the group.
* **Cancel**: Exits the dialog without creating a new group.
* **Confirm**: Creates a new group and exits the dialog.

#### Azure Subscription List

The Azure **Subscription** list displays all the Azure subscriptions within the organization and is used to filter the topology view to display only the selected resources associated with the subscriptions. Each row in the subscription list contains the following UI elements:

* **Select** (checkbox): A checkbox that selects this subscription for the topology view.
* **Name**: The name of the Azure subscription.
* **Subscription ID**: The ID of the Azure subscription.
* **Cloud Entities**: Icons indicating the different Azure cloud entities along with a numerical count of each entity type associated with the subscription.

### GCP Topology

The Google Cloud Platform (GCP) topology view is organized into the following blocks:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(456).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A52Z&se=2026-05-12T09%3A58%3A52Z&sr=c&sp=r&sig=aK9HyBKT4hQmTy20PnTHGgtU9JjObd2PkWAWxC6EW%2Bg%3D)**On Prem**: Displays your on-premises infrastructure that is connected GCP resources.
* **Internet**: Shows the external traffic sources and destinations (origin networks, service providers, and next-hop networks).
* **VPCs**: Each VPC block indicates the regions a VPC spans. GCP VPCs are global and can span multiple regions. Click the **Show Details** link to open a VPC’s **Details** drawer (see **Kentik Map Details**).
* **Regions**: Displayed within the VPC block. By default, regions are collapsed. Click to expand and view subnets. The display of regions varies as follows:

  + **Collapsed**: Shown as a rectangle with the region’s name.
  + **Expanded**: Appears as a green block, displaying subnets and interconnection elements (gateways, etc.), with a **Show Details** link to open its **Details** drawer.
* **Subnets**: Each subnet in a region is a rectangle showing the subnet name and IP/CIDR. Click a subnet and select **Show Details** to open its **Details** drawer.
* **VPC connections**: Supports various interconnection types (see **GCP Interconnection Elements**).

#### GCP Interconnection Elements

The GCP topology view includes the following interconnection elements:

* **Cloud Router**: Facilitates dynamic routing to facilitate traffic between VPCs and external networks. See GCP documentation at **Cloud Router Overview**.
* **External VPN Gateway**: Connects external networks to your VPC via VPN tunnels. See GCP documentation at **Cloud VPN Overview**.
* **Interconnect Attachment**: Connects on-prem networks to VPS cetworks through a specific interconnect (also known as Dedicated Interconnect attachment). See GCP documentation at **Dedicated Interconnect Overview**.
* **VPN Gateway**: Connects VPC networks to external networks via VPN tunnels. See GCP documentation at **Cloud VPN Overview**.
* **VPN Tunnel**: A VPN tunnel provides a path for network traffic between your GCP VPC and external networks. See GCP documentation at **Key Terms**.
* **NAT Gateway**: Provides network address translation for outbound traffic to the internet, VPCs, on-prem networks, and other cloud provider networks. See **Cloud Nat Overview**.

### OCI Topology

The OCI topology contains the following blocks:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(457).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A52Z&se=2026-05-12T09%3A58%3A52Z&sr=c&sp=r&sig=aK9HyBKT4hQmTy20PnTHGgtU9JjObd2PkWAWxC6EW%2Bg%3D)**On Prem**: Represents your on-premises infrastructure that is connected to OCI resources.
* **Internet**: Displays external traffic sources and destinations (origin networks, service providers, and next-hop networks).
* **Regions**: Each block represents a region with OCI resources, showing the Virtual Cloud Networks (VCNs) within. Use the **Show Details** link to access the region's **Details** drawer (see **Kentik Map Details**).
* **VCNs**: Displayed within the region block. By default, all VCNs are collapsed. Click a VCN to expand it and view its subnets. The display of the VCN varies as follows:

  + **Collapsed**: Shown as a rectangle with the VCN name, Oracle Cloud Identifier (OCID), and CIDR.
  + **Expanded**: Appears as a red block showing the VCN’s subnets and interconnection elements (gateways, etc.), with a **Show Details** link. Other VCNs appear as labeled cards showing VCN name, OCID, and configured CIDR block.
* **Subnets**: Each VNet subnet is shown as a rectangle with its name and IP/CIDR. Click a subnet, then click **Show Details** to open its **Details** drawer.
* **VCN connections**: OCI supports multiple interconnection types (see **OCI Interconnection Elements**).

> **Notes:** For an OCI network overview, see **Oracle Cloud Infrastructure Networking Overview**.

#### OCI Interconnection Elements

The OCI topology view includes the following interconnection elements:

* **Dynamic Routing Gateway**: A virtual router for traffic between your VCN and on-prem network. See OCI documentation at **Dynamic Routing Gateways**.
* **Internet Gateway**: Manages direct internet access for VCNs, enabling public resource access and connection initiation. See OCI documentation at **Internet Gateway**.
* **IPSec Connection**: A site-to-to VPN connecting an on-prem network and VCN. See OCI documentation at **Site-to-Site VPN Overview**.
* **Local Peering Gateway**: Connects two VCNs in the same region for private IP communication, similar to Azure VNet Peering. See OCI documentation at **Overview of Local VCN Peering**.
* **NAT Gateway**: Simplifies outbound internet connectivity VCN for resources, withoutdirect inter-network connections. See OCI documentation at **NAT Gateway**.

## Health Problems Page

The Kentik Map features a dedicated health assessment system that evaluates entities’ health when you open the map (see **Kentik Map Health**). Clicking the **Health** indicator (heart icon) at the upper right of the Kentik Map Page opens a popup, where a **View Problems** button will display the Health Problems page. The page is built around a table that shows information about each entity that has a health issue (status of either Warning or Critical).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(458).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A52Z&se=2026-05-12T09%3A58%3A52Z&sr=c&sp=r&sig=aK9HyBKT4hQmTy20PnTHGgtU9JjObd2PkWAWxC6EW%2Bg%3D)

*Devices and interfaces with health issues are listed on the Health Problems page.*

### Health Problems Page UI

The **Health Problems** page includes the following UI elements:

* **Filter field**: Enter text to filter issues by alarm type, entity name, or site.
* **Group By**: Select a property (e.g., Site) from the drop-down menu to group the issues in the table by the value of that property. The table supports grouping by alarm type, site, device name, and device label.
* **Health Problems List**: Displays a list of health issues (see **Health Problems List**).

### Health Problems List

The**Health Problems** page features a table that lists the issues identified by the Kentik Map health assessment. To change the sort order of the list, click a heading to select a column, and click the resulting blue up or down arrow to choose the sort direction (ascending or descending).  
  
 The**Health Problems** list includes the following columns (left to right):

* **Alarm Type**: The nature of the health issue (for descriptions, see **Kentik Map Health**).
* **Entity name**: Device and/or interface name. Click a names to go to that entity’s **Details** page (see **Core Details Pages**).
* **Site**: The site in which the entity is located.
* **Current Value**: Triggered value for the alarm (e.g., "150%" for an Alarm Type of "Interface Inbound Utilization").
* **Actions**: Options for further investigation (see **Health Problems Actions**).
* **Total**: Sum of all current health problems is displayed in the lower right of the **Health Problems** list.

### Health Problems Actions

The following actions are available in the far-right columns of the **Health Problems List**:

* **View Entity in Kentik Map**:

  + Device: Opens the topology view for the site’s device, along with the device’s **Details** sidebar (see **Kentik Map Details**).
  + Interface: Opens the topology view for the device containing the interface. The **Details** sidebar will be open for that interface.
* **View Entity Settings**:

  + Device: Opens the **Device Settings Dialog**.
  + Interface: Opens the **Interface Settings Dialog**.
