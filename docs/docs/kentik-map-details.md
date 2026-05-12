# Kentik Map Details

Source: https://kb.kentik.com/docs/kentik-map-details

---

This article covers the **Details** drawer of the **Kentik Map** module in the Kentik portal.

> **Note:** For information about the rest of the Kentik Map module, see **Kentik Map**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(699).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A58Z&se=2026-05-12T09%3A50%3A58Z&sr=c&sp=r&sig=cgYi3NAiE48Xfscjc1qtLGWQE%2BecpCeVxN6%2BFtjxJ64%3D)

*The Details drawer shows details about an entity on the Kentik Map.*

## About Kentik Map Details

The Kentik Map Details drawer displays information about the currently selected  map element. Details are available for the following entity types:

* **Site entities**: Details about network entities in your physical (on-prem) infrastructure:

  + Site: Overall traffic to and from one of your sites.
  + Device (via site topology): Traffic to and from an individual router or host.
  + Interface (via device topology): Traffic to and from an individual interface.
* **Cloud entities**: Details about traffic to and from your organization’s resources in a cloud provider such as AWS, Azure, GCP, or OCI:

  + Provider: Overall traffic to and from the cloud resources.
  + Region (via cloud topology): Traffic to and from your resources in one of the provider’s cloud regions.
  + Subnet (via cloud region topology): Traffic to and from an individual subnet in the cloud region.
* **Internet entities**: Details about your traffic to and from network entities beyond your own physical infrastructure or cloud resources:

  + Origin Network: Traffic to and from an origin AS.
  + Provider: Traffic to and from a service provider.
  + Next-hop Network: Traffic to and from a next-hop AS.
* **Links**: Details about a direct link between two individual map entities, e.g., sites, devices, or interfaces.

> **Note:** To close the drawer, click anywhere outside it or click the Details button in the subnav (see **Kentik Map Page**).

## Details Types

The information displayed in a Kentik Map **Details** drawer varies depending on the specifics of the current entity (see **Entity-specific Details**). Two main categories of details are currently displayed:

* **Link details**: Contains details about the link between two map entities (see **Link Details**). Hovering anywhere on the line directly connecting two entities will cause the line to become bold, and clicking on the line will open the drawer.
* **Entity Details**: Contains details (see **Entity Details**) about an individual entity (one of the types listed in **About Kentik Map Details**). To open the **Details** drawer for an entity, click directly on the map element and choose **Show Details** from the popup **Network Element Actions** menu.

## Link Details

The **Details** drawer for links includes the following main parts:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(700).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A58Z&se=2026-05-12T09%3A50%3A58Z&sr=c&sp=r&sig=cgYi3NAiE48Xfscjc1qtLGWQE%2BecpCeVxN6%2BFtjxJ64%3D)**Details type**: Located at the top left of the drawer, this “Traffic Details” label indicates that the details displayed in the drawer are for a link.
* **Link entities**: Also at top left, this field indicates the two entities at either end of the link.
* **Query Results panes**: A set of one or more panes with graphs and visualizations showing the results returned from queries that are automatically run on the link being detailed (see **Details Panes**).

  > **Note:** The results reflect the Kentik Map’s current time range and filters settings (see **Kentik Map Page**).

## Entity Details

The **Details** drawer for entities includes the following main parts:

* **Entity type**: Located at the top left of the drawer, this field indicates the type of entity for which details are displayed in the drawer (e.g., Site, Cloud, ASN, Data Center, see list in **About Kentik Map Details**).
* **View Topology** (sites only): A button that takes you to a topology view for the entity (see **Topology Views**).
* **Name**: The name (just under the type) of the entity for which the drawer is showing details. The name is a link; click it to open, in a new tab, the **Network Explorer** detail view for the current device (see **Core Details Pages**).
* **Entity-specific details**: Additional information that varies depending on the type of the entity (see **Entity-specific Details**).
* **Query Results panes**: A set of one or more panes with graphs and visualizations showing the results returned from queries that are automatically run on the entity being detailed (see **Details Panes**).

  > **Note:** The results reflect the Kentik Map’s current time range and filters settings (see **Kentik Map Page**).

#### Entity-specific Details

In addition to the general information above, the following additional details or links may be included in the entity information, depending on the type of the entity:

* **Site entities**:

  + Site (Kentik Map only): The Type (e.g., Data Center) and Address (physical location).
  + Device (Kentik Map and topology views): The Type (e.g., Cisco ASA), device name, device ID, and any **Labels** assigned to the device.
  + Interface(topology views only): A drop-down Connection list from which you can choose the link for which information will be displayed in the panes of the Details sidebar (see **Details Panes**).
* **Cloud entities** (non AWS):

  + Cloud (Kentik Map only): Name.
  + Region (topology views only): Name.
  + Subnet (topology views only): Name.
  + VNet (Azure topology only): Name.
* **Cloud entities** (AWS):

  + Cloud (Kentik Map only): Name.
  + Interconnections (topology views only; see **AWS Interconnection Elements**): Name plus details that vary depending on type of interconnection, including ID, Account ID, VPC ID, State, Tags, Destinations.
  + VPC (topology views only): Name, ID, Account ID, CIDR, State, Tags.
  + Subnet (topology views only): Name, ID, Account ID, CIDR, State, Tags.
* **Internet entities**: Details about your traffic to and from network entities beyond your own physical infrastructure or cloud resources:

  + Origin Network: Name and **View Peering Analysis** link, which opens a new tab showing the **Potential Peer Page** for the ASN.
  + Provider: Name.
  + Next-hop Network: Name and **View Peering Analysis** link, which opens a new tab showing the **Potential Peer Page** for the ASN.

> **Note:** The entity name at the top of a **Details** pane is a link that takes you to that entity’s **Details** page in the Core section of the portal (see **Core Details Pages**).

## Details Panes

The panes of the **Details** drawer are covered in the following topics.

### About Details Panes

Every **Details** drawer includes at least one pane that displays the results of queries that Kentik runs regarding the entity or link that is the subject of the drawer. The panes included for a given entity, which depend on the entity type, are listed in the topics below.

#### Universal Details Panes

The following panes are found in the **Details** drawer for all types of network entities:

* **Traffic** (all detail types): Ingress, egress, and top-X for traffic on the selected network element or link (see **Traffic Details**).

#### Infrastructure Details Panes

The panes in the table below are found in the Details drawer for network entities in physical infrastructure.

| Pane | Entity | Description |
| --- | --- | --- |
| Internal Map Details | Sites | A diagram of the site topology (see **Internal Map Details**). |
| Devices | Sites | Information about the devices in a site, including health status and a list of devices with details on each (see **Device Details**). |
| Health | Devices, Interfaces, Sites | Health status based on SNMP metrics (see **Health Details**). |
| Metrics | Devices, Interfaces, Links | Contents vary depending on the type of entity (see **Metrics Details**):   * Links: SNMP counter data (see **Counter SNMP OIDs**) on input and output traffic between interfaces on the devices at each end of the link, including bitrate, errors, and discards. * Devices: SNMP-derived information on utilization over time. * Interfaces: SNMP counter data (see **Counter SNMP OIDs**) on input and output traffic over the interface. |
| Interface Metadata | Links | Information about the interfaces (see **Interface Metadata Details**). |

#### Cloud Details Panes

The panes in the table below are found in the Details drawer for Kentik-supported cloud providers.

| Pane | Provider | Description |
| --- | --- | --- |
| Cloud Details | AWS, Azure, GCP, OCI | Details about various cloud entities and connections (see **Cloud Details**). |
| Cloud Metrics | AWS, Azure | Information and metrics for cloud entities (see **Cloud Metrics Details**). |
| Route Table | AWS, Azure, GCP, OCI | Indicators and controls related to route tables on a VPC or transit gateway (see **Route Table Details**). |
| Security Groups & Network ACLs | AWS | Information about flows that were denied as a result of a rule in a Security Group or a Network ACL (see **Security Groups & NACLs Details**). |
| Network Security Groups | Azure | Information about flows that were denied as a result of a rule in a Security Group (see **NSG Details**). |
| Denied Traffic | Azure | Information about flows that were denied as a result of a firewall rule (see **Denied Traffic Details**). |
| Firewall Policies | GCP | Information about flows that were denied as a result of a firewall rule (see **Firewall Rules Details**). |
| Security Lists | OCI | Information about flows that were denied as a result of a firewall rule (see **Security Lists Details**). |
| Load Balancers | GCP | Information about the load balancers deployed in a given GCP region or VPC (see **Load Balancers Details**). |

### Internal Map Details

If the architecture of the site has already been defined (see **Site Architecture**), the **Internal Map Details** pane will feature a diagram of that topology. Click the expand icon at the upper right of the pane to open the topology in Internal Map Details dialog. In that dialog or in the pane itself you can click on any device in the topology to open a menu with the following options:

* **View Full Map**: Takes you to a topology view for the site containing the device (see **Site Topology**).
* **Show Connections**: Highlight (in blue) the links between the selected entity and other blocks.

If the architecture of the site hasn’t already been defined you can click the **Configure Site** link, which opens the Edit Site dialog (see **Site Settings**).

### Cloud Details

The **Details** pane for an entity in AWS, Azure, GCP, or OCI includes various information about the entity that varies depending on both the cloud provider and entity type. The tables in the following topics show which details are included for the various entity types.

#### AWS Entity Details

Details for AWS regions, VPCs, and subnets.

| Detail | Region | VPC | Subnet |
| --- | --- | --- | --- |
| Transit Gateways | Y | N | N |
| VPN Gateways | Y | N | N |
| VPCs | Y | N | N |
| ID | N | Y | Y |
| Account ID | N | Y | Y |
| CIDR | N | Y | Y |
| State | N | Y | Y |
| Tags | N | Y | Y |
| VPC ID | N | N | Y |

#### AWS-to-On Prem Connection Details

Details for interconnections between on premises infrastructure and resources in AWS:

| Detail | Customer Gateway | LAG | Direct connection | Direct Connect Gateway | VPN Connection |
| --- | --- | --- | --- | --- | --- |
| ID | Y | Y | Y | Y | Y |
| State | Y | N | Y | Y | Y |
| Tags | Y | N | Y | N | Y |
| Destination | Y | N | N | N | N |

#### AWS Internal Connection Details

Details for interconnections between and within AWS regions:

| Detail | Internet Gateway | Peering Connection | Virtual Gateway | TGW Attachment | NAT Gateway | VPC Interface Endpoint | Transit Gateway |
| --- | --- | --- | --- | --- | --- | --- | --- |
| ID | Y | Y | Y | Y | Y | Y | Y |
| State | N |  | Y | Y | Y | N | Y |
| Tags | N | Y | Y | Y | N | N | Y |
| Destination | N | N | N | N | N | N | N |
| Account ID | Y | N | N | N | N | Y | Y |
| VPC ID | Y | N | Y | Y | Y | N | N |
| Subnet ID | N | N | N | N | Y | Y | N |
| Service Name | N | N | N | N | N | Y | N |
| ENI ID | N | N | N | N | N | Y | N |

#### Azure Entity Details

Details for Azure regions, VNets, and subnets.

| Detail | Region | VNet | Subnet |
| --- | --- | --- | --- |
| Tenant ID | Y | Y | Y |
| Subscription ID | Y | Y | Y |
| Location | Y | Y | N |
| Latitude | Y | N | N |
| Longitude | Y | N | N |
| Resource group | N | Y | Y |
| Name | N | Y | Y |
| CIDRs | N | Y | Y |
| State | N | Y | Y |
| Tags | N | Y | N |
| Local Network | N | N | N |
| Remote Network | N | N | N |
| Subnets | N | N | N |

#### Azure Interconnection Details

Details for Azure interconnections:

| Detail | VNet Gateway | VNet Peering | NAT Gateway |
| --- | --- | --- | --- |
| Tenant ID | Y | Y | Y |
| Subscription ID | Y | Y | Y |
| Location | N | N | N |
| Latitude | N | N | N |
| Longitude | N | N | N |
| Resource group | Y | N | N |
| Name | N | N | N |
| CIDRs | N | N | N |
| State | N | N | N |
| Tags | N | N | N |
| Local Network | N | Y | N |
| Remote Network | N | Y | N |
| Subnets | N | N | Y |

#### GCP Entity Details

Details for GCP VPCs and subnets.

| Detail | VPC | Subnet |
| --- | --- | --- |
| Resource URL | Y | Y |
| Project | Y | Y |
| ID | Y | Y |
| Name | Y | Y |
| Routing Mode | Y | N |
| Firewall Policy Enforcement | Y | Y |
| CIDR | Y | Y |
| Network | N | Y |
| Region | N | Y |
| Purpose | N | Y |
| Sampling Enabled | N | Y |
| Sampling Rate | N | Y |
| Gateway | N | Y |
| Secondary CIDRS | N | Y |

#### OCI Entity Details

Details for OCI regions, VCN, and subnets.

| Detail | Region | VCN | Subnet |
| --- | --- | --- | --- |
| Tenancy ID | N | Y | N |
| Compartment ID | N | Y | Y |
| Region Name | Y | Y | N |
| Region Key | Y | N | N |
| Home Region | Y | N | N |
| Name | N | Y | Y |
| CIDRs | N | Y | Y |
| ID | N | Y | N |
| Created At | N | Y | N |
| State | N | N | Y |
| Virtual Router IP | N | N | Y |

### Traffic Details

The **Traffic** pane includes the following UI elements:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(701).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A58Z&se=2026-05-12T09%3A50%3A58Z&sr=c&sp=r&sig=cgYi3NAiE48Xfscjc1qtLGWQE%2BecpCeVxN6%2BFtjxJ64%3D)**Expand/Collapse**: Toggles visibility between the title bar only and the full pane.
* **Information** (icon): A tooltip that describes how the traffic information is calculated. The tooltip is displayed only in the Traffic pane for the links between top-level blocks (Clouds, Internet, and On Prem).
* **Open in modal** (diagonal arrows icon): Opens the pane in a modal.
* **Traffic selector**: A drop-down menu from which to choose the set of traffic that will be evaluated for the query whose results will be displayed in the pane. Options (e.g., All Traffic, External Traffic, etc.) will vary depending on the entity type of the details drawer (see list in **About Kentik Map Details**).
* **Dimension selector**: A drop-down from which to choose the dimension of the query.
* **Metric selector**: The metric used to quantify the query results shown in the pane.
* **Sync Chart Scales**: A checkbox that causes the vertical axis of the Ingress and Egress charts to be the same scale.
* **Traffic charts**: Two time-series stacked charts (or one if traffic is set to All Traffic), the top for ingress and the bottom for egress, showing traffic over the time period specified in the **Time** pane of the **Options** sidebar. The charts include the following elements:

  + Heading: Shows the direction and volume of the traffic to or from this entity.
  + Time scale (bottom chart only): A time scale representing the time period specified in the **Time** pane of the **Options** sidebar.
  + Time-point details: A popup, which opens when hovering over either chart, that shows values for the Total and Historical Total (7 days prior) at that point in the time range.
  + View in Data Explorer (icon; on hover only): Opens a new browser tab to show the query in the portal’s **Data Explorer** module.
* **Traffic table** (not shown for total traffic queries): A list of the top-X results returned from the current query. This table is similar to the traffic table in Data Explorer (see **Explorer Table Overview**).

> **Note:** The query results returned in this pane are affected by the settings in the Kentik Map’s **Time** and **Filter** controls (see **Kentik Map Page**).

### Device Details

Information about the devices in a site, including:

* A summary of device health showing the number of devices with various health statuses.
* A list of devices in the site, each of which can be expanded to show the following:

  + Information table: IP address, site, sample rate, machine type, and device ID.
  + Utilization charts: SNMP-discovered average utilization over time for CPU and memory.

### Health Details

The **Health** pane gives an entity’s current health status (see **Kentik Map Health**):

* If the status of all health metrics for the entity is Healthy then the tab will contain a single indicator stating that all is well.
* If status of any health metric is not Healthy then tab will contain a card, with values and sparkline, for each metric whose status is either Warning or Critical.

This pane also includes the following controls:

* **Expand/Collapse**: Toggles visibility between the title bar only and the full pane.
* **Open in modal** (diagonal arrows icon): Opens the pane in a modal.

> **Note:** If the total number of health issues for the entity is greater than zero the count of issues will be displayed in an indicator in the **Health** tab head.

### Metrics Details

The contents of the **Metrics** pane vary depending on whether it’s in the Details drawer for a network entity that is in physical infrastructure or an entity on the Topology page for a cloud provider (see **Cloud Metrics Details**).

#### Metrics for Links

In physical infrastructure, the **Metrics** pane for links shows the following information from SNMP polling of the devices on which the interfaces exist (for descriptions see **SNMP Interface Metrics**). The information is presented in two columns, one each for the interfaces at either end of the link):

* SNMP Bits/s Out![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(702).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A58Z&se=2026-05-12T09%3A50%3A58Z&sr=c&sp=r&sig=cgYi3NAiE48Xfscjc1qtLGWQE%2BecpCeVxN6%2BFtjxJ64%3D)
* Input Errors
* Output Errors
* Input Discards
* Output Discards

The links version of this pane also includes the following controls:

* **Expand/Collapse**: Toggles visibility between the title bar only and the full pane.
* **Open in modal** (diagonal arrows icon): Opens the pane in a modal.
* **View in Data Explorer**: A button that opens the portal’s **Data Explorer** module in a new browser tab to show a query illustrating traffic over the link.

#### Metrics for Interfaces

For interfaces, the pane contains the same SNMP information listed above for links, but for only one interface. For devices, the pane contains charts showing SNMP-derived information on utilization over time, including average CPU utilization and average memory utilization.

### Cloud Metrics Details

For a network entity that is on the Topology page for AWS and Azure, the **Metrics** pane shows the metrics associated with network entities in the cloud. AWS CloudWatch metrics are displayed for AWS entities and Azure Monitor is used for Azure entities. (There’s no **Metrics** pane in GCP or OCI topology.) The pane includes the following UI elements:

* **Expand/Collapse**: Toggles visibility between the title bar only and the full pane.
* **Interval**: The frequency (e.g., 1 minute) at which the data is sampled.
* **Refresh**: Updates the table to show the most recent data available.
* **Open in modal** (diagonal arrows icon): Opens the pane in a modal (see **Metrics Modal**).
* **Filter**: A field at the top of the **Metrics** table into which you can enter text to filter the listed metrics by name.
* **Metrics table**: A table that displays the different types of metrics depending on the entity type (see **Metrics Pane Table**).

#### Metrics Pane Table

This table can be sorted by column and includes the following headings:

* **Metric**: The measurement being analyzed. Hover over the metric icon to display a sparkline of the metric.

  > **Note**: The metric icon is greyed out if there is no data collected for the metric.
* **Avg**: The average value over the specified interval.
* **Sum** (Azure only): The sum value over the specified interval.

#### Metrics Modal

The metrics modal contains the following fields and controls:

* **Title**: The name of the entity and the metric.
* **Close**: (**X** at the upper right): Closes the modal and returns to the main view.
* **Time and interval**: Shows a timestamp and the frequency (e.g., 1 minute) at which the metric is reported.
* **Filter**: A field at the top of the **Metrics** table into which you can enter text to filter the listed metrics by name.
* **Metrics table**: A table that displays the different types of metrics depending on the entity type (see **Metrics Modal Table**).

#### Metrics Modal Table

The table in the metrics modal can be sorted by column and includes the following columns:

* **Metric**: The measurement being visualized.

  > **Note**: The metric icon is greyed out if there is no data collected for the metric.
* **Trend**: A sparkline of the metric.
* **Avg**: The average value over the specified interval.
* **Sum** (Azure only): The sum value over the specified interval.

### Interface Metadata Details

The **Interface Metadata** pane includes information about the link whose details are currently shown in the **Details** drawer, as well as about the interfaces at each end of the link, details for which are shown in two columns (one for each interface):

* **Layer**: Indicates the layer (2 or 3) of this link’s connection. If the **Draw Links Using** drop-down (see **Kentik Map Page**) is set to “All Layers” and connections between the entities at either end of the link exist on both layers then this pane will include a metadata section for Layer 2 and a section for Layer 3.
* **Interface name**: The name of the interface.
* **Interface description**: The interface description as either defined in the device and retrieved via SNMP or specified manually. Capped at 128 characters.
* **Device**: The name of the device to which this interface belongs.
* **IP Address**: The primary IP address of this interface.
* **Capacity**: The maximum capacity in Mbps as reported by SNMP.
* **Network Boundary**: The network boundary value assigned to the interface by interface classification (see **Network Boundary Attribute**).
* **Connectivity Type**: The network boundary value assigned to the interface by interface classification (see **Connectivity Type Attribute**).

![Interface metadata showing device details, IP addresses, and capacities for network connections.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/KMD-interface-metadata-pane.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A58Z&se=2026-05-12T09%3A50%3A58Z&sr=c&sp=r&sig=cgYi3NAiE48Xfscjc1qtLGWQE%2BecpCeVxN6%2BFtjxJ64%3D)

*Using the Expand button, a Details pane can be opened in a modal.*

This pane also includes the following controls:

* **Expand/Collapse**: Toggles visibility between the title bar only and the full pane.
* **Open in modal** (diagonal arrows icon): Opens the pane in a modal.

### Route Table Details

The **Route Tables** details pane includes indicators and controls related to route tables on a VPC or transit gateway. The contents of the pane is similar for all Kentik-supported cloud providers, but with some variation between AWS, Azure, GCP, and OCI.  
  
The pane’s title bar includes the following:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(704).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A58Z&se=2026-05-12T09%3A50%3A58Z&sr=c&sp=r&sig=cgYi3NAiE48Xfscjc1qtLGWQE%2BecpCeVxN6%2BFtjxJ64%3D)**Expand/Collapse**: Toggles visibility between the title bar only and the full pane.
* **State**: Lozenges that indicate the number and state of the following:

  + Tables (blue): The number of tables (if more than one).
  + Active routes (green): The number of routes with a good destination.
  + Blackholed (red): The number of routes that are programmed in a table but whose destination can’t be reached.
* **Azure Console** (Azure only): A button that takes you to your organization’s console in Azure.
* **Open in modal** (diagonal arrows icon): Opens the pane in a modal (with identical UI).

Below the title bar is a table made up of the following:

* **Filter**:A field in which you can enter text to filter the routes shown in the table.

  + In AWS and Azure this matches against the Route Target column.
  + In GCP and OCI this matches against the Destination column.
* **Header rows** (AWS/Azure/OCI only): A collapsible/expandable header row precedes the list of routes for one route table. A header row includes the following:

  + State (AWS/Azure only): An icon representing the state of the table’s routes (see State above).
  + Name: The name of the route table.

    > **Note:** Header rows are not used in GCP because there is only one group per network.
* **Route rows**: Rows that each represent an individual route in a route table.

The information in the individual route rows is presented in the following columns:

* **State**: Active (checkmark) or blackholed (exclamation point).
* **Destination**: The destination CIDR block against which traffic is evaluated to determine the route target to which it should be forwarded.
* **Route target** (AWS/Azure/GCP only): Content depends on the cloud provider and the type of network entity:

  + *V*PC (AWS): The ID of the gateway that will handle the different routing functions within a VPC.
  + Transit gateway (AWS): The Attachment (the transit gateway extension that “lives” in a VPC) and the Next Hop Resource (the next resource that the traffic will enter).
  + VPC (GCP): The route target VPC.
  + Default internet gateway (GCP): The entity’s default internet gateway.
* **Network Entity** (OCI only): The name of the OCI network entity associated with the route table.
* **Next Hop** (Azure only): The network entity’s next hop in the route table.

> **Notes:**
>
> * Subnets can either use a main route table (indicated by “Main” in parentheses after the table’s name in the header row) or a dedicated route table.
> * For additional information on route tables in AWS, see AWS docs on **VPC Route Tables**.

### Security Groups & NACLs Details

This pane includes the following UI elements:

* **Expand/Collapse**: Toggles visibility between the title bar only and the full pane.
* **Open in modal** (diagonal arrows icon): Opens the pane in a modal.
* **Filter**: A field in which you can enter text to filter the **Denied Traffic** table. The table displays any rows with a match in the **Source** or **Destination** column.
* **Denied Traffic**: A table showing flows that were denied as a result of a rule in a security group or a network ACL (see **Denied Traffic Table**).
* **View Security Groups**: Open the **Security Groups Tab** of the Security Groups & Network ACLs dialog.
* **View Network ACLs**: Opens the **Network ACLs Tab** of the Security Groups & Network ACLs dialog.

#### Denied Traffic Table

Each row of this table shows a set of flows that were denied as a result of a rule in a security group or a network ACL. The flows in a given row share a direction, source, and destination. Click on any row to open the Security Groups & Network ACLs dialog (see **Security Groups & NACLs**).  
  
The table can be sorted by its column headings, which include the following:

* **Direction**: The direction of traffic.
* **Source**: The IP address of the source.
* **Destination**: The IP address of the destination.
* **Total Flows**: The number of denied flows.
* **View in Data Explorer**: An icon that opens **Data Explorer**, where the query settings will be set to show these flows.

### NSG Details

The **Network Security Groups** pane includes information showing flows that were denied as a result of a rule in a security group. The pane includes the following UI elements:

* **Expand/Collapse**: Toggles visibility between the title bar only and the full pane.
* **Open in modal** (diagonal arrows icon): Opens the pane in a modal.
* **Filter**: A field in which you can enter text to filter the **Denied Traffic** table. The table displays any rows with a match in the **Source** or **Destination** column.
* **Denied Traffic Table**: A table showing flows that were denied as a result of a rule in a security group (see **Denied Traffic Table**).

### Denied Traffic Details

The **Denied Traffic** pane has the same UI elements as the Network Security Group pane (see **NSG Details**) but displays information on flows denied as a result of a firewall rule.

### Firewall Rules Details

The **Firewall Rules** pane includes indicators and information related to GCP firewall rules. The pane includes the following UI elements:

* **Expand/Collapse**: Toggles visibility between the title bar only and the full pane.
* **Rule count indicators**: A lozenge that indicates the number of active rules.
* **Open in modal** (diagonal arrows icon): Opens the pane in a modal.
* **Firewall Rules Table**: A table that displays the firewall rules (see **Firewall Rules Table**).

#### Firewall Rules Table

The following controls are found at the top of the table:

* **Group By**: A drop-down to group the rules based on column. Each group will start with a header row that states the value that’s common to the rules in the group and the number of rules in the group.
* **Filter**: A field with which you can filter the table to rows containing the entered text in the **Resource** or **Protocol** columns.

The table can be sorted by column headings, which include the following:

* **Direction**: The direction of traffic.
* **Rule Action**: The policy of the rule.
* **Resource**: The IP address of the destination.
* **Port**: The port number.
* **Protocol**: The traffic protocol.
* **View in Data Explorer**: An icon to open the traffic information in **Data Explorer**.

### Firewall Policies Details

The **Firewall Policies** pane enables you to see information about firewall policies in each GCP region or VPC. In GCP, firewall policies are groups of firewall rules that can be applied across your cloud organization. The following fields are included in the pane:

* **Policy Name**: The name of the firewall policy.
* **Tuple Count**: The total count of firewall policy rule tuples.
* **Policy Created**: The timestamp indicating the date and time when the firewall policy was initially established.
* **Scope**: The level at which the firewall policy is applied within the GCP organization.
* **Associations**: Indicates where the rules within the policy are actively enforced.
* **Firewall Policies Table**: A table that displays the firewall rules (see **Firewall Policies Table**).

#### Firewall Policies Table

The following controls are found at the top of the table:

* **Group By**: A drop-down to group the rules based on column. Each group will start with a header row that states the value that’s common to the rules in the group and the number of rules in the group.
* **Filter**: A field with which you can filter the table to rows containing the entered text in the **Source**, **Destination**, and **Protocol and Port** columns.

The table can be sorted by column headings, which include the following:

* **Direction**: The direction of traffic.
* **Rule Action**: The policy of the rule.
* **Source**: The IP address of the source.
* **Destination**: The IP address of the destination.
* **Protocol and Port**: The traffic protocol and port number.

### Security Lists Details

The **Security Lists** pane enables you to see information about traffic rules in each OCI region or VCN. In OCI, security lists consist of a set of ingress and egress traffic rules that apply to all Virtual Network Interface Cards (VNICs) in any subnet that the security list is associated with. The pane displays two tables, **Ingress Security Rules** and **Egress Security Rules**, that share similar UI elements.  
  
 The pane's title bar includes the following:

* **Expand/Collapse**: Toggles visibility between the title bar only and the full pane.
* **Open in modal** (diagonal arrows icon): Opens the pane in a modal.

Below the title bar is a table made up of the following:

* **Filter**:A field in which you can enter text to filter the rules shown in the table.
* **Header rows**: A collapsible/expandable header row precedes the list of rules for one security list. A header row includes the following:

  + State: Indicates whether the rule is active.
  + Name: The name of the security list.
  + Number: The count of rules in the security list.
* **Security List rows**: Rows that each represent an individual rule in a security list.

The information in the individual rule is presented in the following columns:

* **Source** (ingress only): The CIDR block from which incoming traffic is allowed.
* **Source Type** (ingress only): The source type indicating a type of IP address.
* **Destination** (egress only): The destination CIDR block against which traffic is evaluated to determine the route target to which it should be forwarded.
* **Destination Type** (egress only): The destination type indicating a type of IP address.
* **Protocol**: The type of network protocol allowed by the rule.
* **Description**: A description of the rule.

### Load Balancers Details

The **Load Balancers** pane enables you to see — without navigating to the Google Cloud console — information about the load balancers deployed in a given GCP region or VPC (see **Cloud Load Balancing overview**). The fields included in the pane vary depending on the specifics of your load balancing setup. The following fields may be included:

* **Load Balancer Name**: The name of the load balancer.
* **Load Balancer Type**: The type of load balancer, either Network or Application.
* **Front End IP**: The IP address of the load balancer.
* **Protocol**: The transport layer protocol used to direct traffic, such as TCP, UDP, ESP, GRE, ICMP, and ICMPv6.
* **Port Range**: The port on which to access the load balancer.
* **Scope**: The scope of the load balancer:

  + Region: The balancer operates at the region level.
  + Global: The balancer operates at the level of a VPC (which GCP refers to as a network).
* **Load Balancing Scheme**: An attribute on the forwarding rule and the backend service of a load balancer that indicates whether the load balancer can be used for internal or external traffic: EXTERNAL, EXTERNAL\_MANAGED, INTERNAL, or INTERNAL\_MANAGED.
* **Network Tier**: Network Service Tiers enable you to optimize network quality and performance vs. cost for your resources and projects. Premium is optimized for performance, Standard is optimized for cost.
* **Target Proxy**: Proxies that terminate incoming connections from clients and create new connections from the load balancer to the backends.
* **Instance Group Count**: The number of instance groups (grouped virtual machine (VM) instances) used in backend services or target pools by this load balancer.
* **Network Endpoint Group Count**: The number of network endpoint groups (configuration objects that specify a group of backend endpoints or services), which are used for more granular control over the distribution of traffic to the load balancer’s backends.

#### Backend Pools Table

The **Load Balancers** pane also includes a table for **Backend Pools** that shows the groups of resources that will serve traffic for the balancer’s load-balancing rules. The table, which can be filtered by IP with the **Filters** field, includes the following columns:

* **Primary Internal Ip**: The IP address of a backend VM or container in the VPC network that is connected to from the load balancer.
* **External Ip**: The IP address used to connect to the load balancer.
* **Nat Name**: The name of the NAT that the load balancer is behind.

## Security Groups & NACLs

The Security Groups & Network ACLs dialog, which opens from the **Security Groups & Network ACL****s** details pane in an AWS topology view (see **Security Groups & NACLs**), provides information related to flows that were denied as a result of a security rule. The dialog contains two tabs, one for security groups and the other for NACLs.

#### Security Groups Tab

The **Security Groups** tab contains the following UI elements:

* **Security Groups**: A drop-down selector from which to choose a security group.
* **Security group information**: A set of fields giving the name, description, account ID, security group ID, and VPC ID of the selected security group.
* **Denied Flows**: A table listing flows that were denied by the rules of this security group:

  + The table gives the direction, resource, protocol, port/range, IP version, and description for each listed flow.
  + The flows can be grouped by direction, protocol, or IP version.
  + You can enter text in the filter field to show only rows in which the contents of a column matches the entered text.
  + A **View in Data Explorer** button at the right of each row will take you to **Data Explorer**, where the query settings will be set to show the flow.

#### Network ACLs Tab

The **Network ACLs** tab contains the following UI elements:

* **Network ACLs**: A drop-down selector from which to choose a security group.
* **Network ACL information**: A set of fields giving the account ID, and VPC ID of the selected network ACL.
* **Denied Flows**: A table listing flows that were denied by the rules of this network ACL:

  + The table gives the direction, rule action, rule number, resource, protocol, port/range, and IP version for each listed flow.
  + The flows can be grouped by direction, protocol, IP version, or rule action.
  + You can enter text in the filter field to show only rows in which the contents of a column matches the entered text.
  + A **View in Data Explorer** button at the right of each row will take you to **Data Explorer**, where the query settings will be set to show the flow.

---

© 2014-25 Kentik
