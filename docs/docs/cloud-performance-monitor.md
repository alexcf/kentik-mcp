# Cloud Performance Monitor

Source: https://kb.kentik.com/docs/cloud-performance-monitor

---

This article explains how to use Kentik’s **Cloud Performance Monitor** to visualize and manage connectivity across hybrid and multi-cloud environments.

> **Notes:**
>
> * The Cloud Performance Monitor page is visible only to customers that have **at least one AWS cloud** registered in **Settings »** **Public Clouds**.
> * For general information about monitoring your cloud resources with Kentik, start with **Cloud Overview**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(166).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A57Z&se=2026-05-12T09%3A34%3A57Z&sr=c&sp=r&sig=4ymRxofdSBL8krn5Y0NcO4uCN3uhtYurrxzXJ3O6VL0%3D)

*The Interconnects tab of Cloud Performance shows paths from VPCs to your on-prem network.*

## About Cloud Performance Monitor

The **Cloud Performance Monitor** page provides deep visibility into connectivity across hybrid and multi-cloud environments. It helps network engineers visualize and manage performance paths between cloud VPCs and on-prem infrastructure.

The page feature two primary tabs that focus on different aspects of cloud health:

* **Interconnects** **(AWS Only)**: Visualize and monitor **Direct Connect** and **Site-to-Site VPN** connections. It automatically maps paths between cloud and on-premises elements.
* **Services**: Manage **synthetic testing** for critical AWS and Azure services (e.g., S3, EC2, Front Door). It uses flow data to identify service utilization and guide agent deployment.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(167).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A57Z&se=2026-05-12T09%3A34%3A57Z&sr=c&sp=r&sig=4ymRxofdSBL8krn5Y0NcO4uCN3uhtYurrxzXJ3O6VL0%3D)

## Cloud Performance Monitor UI

The **Cloud Performance Monitor** page includes the following controls and indicators:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(168).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A57Z&se=2026-05-12T09%3A34%3A57Z&sr=c&sp=r&sig=4ymRxofdSBL8krn5Y0NcO4uCN3uhtYurrxzXJ3O6VL0%3D)**Time Range**: Adjust the lookback period for performance metrics and status (see **Time Range Control**).
* **Show** (Services tab only): Toggle visibility between AWS, Azure, or both cloud providers.
* **Manage Agents** (button): Opens the **Manage Agents Dialog** to deploy or configure `ksynth` private agents within your cloud infrastructure.
* **Filter** (Interconnects tab only): Quickly find services by AWS Account ID, Entity ID, CIDRs, or Tags.

### Agent Status Counters

These indicators show the real-time installation status of agents across your VPCs and VNets:

| Status | Meaning |
| --- | --- |
| Unmonitored | Resources with no agent deployed. |
| Pending | Agents deployed but not yet active. |
| Fully Installed | Active agents currently collecting data. |

> **Note:** On the **Services Tab**, the count shown in these counters depends on the services that have been selected for monitoring with the **Configure Service Monitoring** dialog.

## Manage Agents Dialog

Use this dialog enables to deploy and configure `ksynth` private agents (see **Kentik Synthetics Agents**). The list of VPCs/VNets filters dynamically based on your current tab (Interconnects vs. Services).

> **Note:** The VPCs listed in this dialog are **not** affected by the services selected for monitoring with the **Configure Service Monitoring** dialog on the **Services Tab**.

![List of unmonitored VPCs with options to deploy agents for monitoring.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CPM-manage-agents-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A57Z&se=2026-05-12T09%3A34%3A57Z&sr=c&sp=r&sig=4ymRxofdSBL8krn5Y0NcO4uCN3uhtYurrxzXJ3O6VL0%3D)

| Tab | Description |
| --- | --- |
| Unmonitored VPCs | Lists VPCs needing agents. Click **Deploy Agent** to launch the standard **Private Agent Setup**. |
| Pending Agents | Agents installed but not yet active. |
| Private Agents | Active agents. Click **Configure Agent** to edit settings via the **Configure Agent Dialog**. |

> **Note:** Use the **Remove** button in the Synthetics Agent Management module of the portal to uninstall agents (see **Details Sidebar UI**).

## Interconnects Tab

The **Interconnects** tab visualizes paths from AWS VPCs (left) to on-prem Customer Gateways (right), grouped by region. Depending on how your VPC connections are configured, this tab shows you either or both of the following:

* **Direct Connect**: Each Direct Connect is a dedicated network connection between your network and public AWS resources through a virtual private gateway.
* **Site-to-Site VPN**: Each VPN refers to a connection between a VPC and your own on-premises network.

> **TIP:** For diagrammed examples of the different types of connections supported by AWS, see **Connection Examples**.

#### Health Legend

| Line Style/Color | Meaning |
| --- | --- |
| **Dotted Gray** | Links from VPCs without an active agent. |
| **Solid Gray** | Agent is installed, but no health test has been created (see **Create Path Health Test**). |
| **Solid Green** | Healthy path status. |
| **Solid Orange/Red** | Warning or Critical status determined by synthetic tests (see **Test Status Levels**). |

#### Create Path Health Test

To create a health test for the path between a VPC and a customer gateway, an agent must already be installed on the VPC (see **Manage Agents Dialog**).

1. Click on any of the gray links between the entities on the path.
2. The link popup will open. Choose **Create Test**.
3. Testing will start. After about a minute the links in the path will be colored based on health status.

### Using the Interconnects Tab

The Interconnects tab supports a number of actions that give you a clear picture of your connections to your cloud resources:

* For each region, a **View in Map** link takes you to the **Kentik Map**.
* Hover over a VPC to see the path that the traffic takes to your on-prem infrastructure.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(170).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A57Z&se=2026-05-12T09%3A34%3A57Z&sr=c&sp=r&sig=4ymRxofdSBL8krn5Y0NcO4uCN3uhtYurrxzXJ3O6VL0%3D)

  *The path from a VPC at left to a Customer Gateway at right.*
* Hover over a link between network elements in the path (e.g., gateways and connections) to see the upstream (left) and downstream (right) connections that make up the remainder of that path.
* Click a network element to open a Details drawer from the right, providing further information in areas such as Traffic, Performance, and Route Table.
* Click a link (segment of a path) to open a informational popup. The type of line representing each link depends on whether the link is being monitored with performance testing:

  + **Colored line**: A green, orange, or red line indicates the health status of a tested link (see **Test Status Levels**). The popup will indicate the current status of the link for latency, packet loss, and jitter.
  + **Dashed gray line**: The path is not currently being monitored. The popup includes a **Deploy Agent** link, which opens the Deploy an Agent dialog so you can begin testing the link.

    ![Instructions for deploying an agent in AWS VPC with activation button.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CPM-deploy-an-agent-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A57Z&se=2026-05-12T09%3A34%3A57Z&sr=c&sp=r&sig=4ymRxofdSBL8krn5Y0NcO4uCN3uhtYurrxzXJ3O6VL0%3D)
* In a link popup:

  + Click **Show Details** to open the Details drawer, which contains details about the link in areas such as Traffic and Performance. These details parallel those shown in the Details drawer on the Kentik Map (see **Kentik Map Details**).
  + Click **Create Test** to activate health testing on the path that includes this link (see **Create Path Health Test**).
* The **Interconnects** tab shows only Direct Connects from accounts (or child accounts using **AssumeRole**) that are configured within a Kentik cloud export. The No Interconnection Found message will display if **either** of the following is true:

  + The account(s) covered by a cloud export don't include any instance of Direct Connect.
  + The account(s) covered by a cloud export include an instance of Direct Connect but the relevant account permissions are incorrectly configured.

## Services Tab

The **Services** tab helps you manage synthetic testing of important AWS and Azure cloud services (S3, EC2, Storage, Front Door, etc.) to evaluate their performance. Using flow data gathered by monitoring your cloud resources, Kentik can identify utilization of various services based on IP address, either public or private (VPC or VNet endpoints that expose cloud services), and guide you as to which resources to test performance. You can also easily launch into deploying synthetics agents on those resources.

### Service Monitoring UI

The **Services** tab includes the following main UI elements:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(172).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A57Z&se=2026-05-12T09%3A34%3A57Z&sr=c&sp=r&sig=4ymRxofdSBL8krn5Y0NcO4uCN3uhtYurrxzXJ3O6VL0%3D)

* **Search**: A field that filters the **Services List** to services in which the text matches values one of the following fields in the list's cards and tables: Service name, Region, Connection Type, HTTP status, HTTP Latency.
* **Update tests**: A button that directs Kentik to automatically configure appropriate synthetic tests to the services that you've chosen to check (see **Configure Service Monitoring**) and Kentik has discovered that you use.

  > **Note:** This button is only active if agent deployment has changed since the last time tests were updated (see **Manage Agents Dialog**).
* **Test coverage**: An indicator stating (a) as a percent of (b):  
   (a) the number of AWS VPCs or Azure VNets with agents that are active and configured to run tests;  
   (b) the count of your organization's VPCs or VNets that Kentik has "discovered" as using any AWS or Azure services.
* **Settings** (gear icon): A button that opens the **Configure Service Monitoring** dialog, where you can include or exclude any discovered service from being tested.

  > **Note:** Changing the services that are monitored will change the counts shown in the agents counters at the upper right when you're on the Services tab of the Performance Monitor page.
* **Services list**: A list of the AWS services used by your organization (see **Services List**).

### Services List

The Services list is a set of cards that each provide information about your organization's usage of a cloud service (e.g., S3, EC2, Front Door, Storage, etc.). Each card includes the following UI elements:

* **Expand/collapse**: An arrow that you click on to expand or collapse the **Services Table**.
* **Service name**: The name of the monitored cloud service (CloudFront, EC2, S3, Global Accelerator, etc.).
* **Connection type**: An indicator stating the connection type and count of VPCs or VNets using the service via public and private connections.
* **From service**: An indicator showing, with numbers and a sparkline, the traffic volume in bps from the service to the monitored VPCs or VNets.
* **To service**: An indicator showing, with numbers and a sparkline, the traffic volume in bps to the service from the monitored VPCs or VNets.
* **Actions menu** (vertical ellipsis): The following option is available from the menu:

  + **View in Data Explorer**: Opens (in a new tab) **Data Explorer**, with the Query sidebar set to show the cloud  traffic for this service.
* **Services table**: A list of VPCs/VNets using the service, grouped by region (see **Services Table**).

> **Note:** If the health status of any of a service's regions, VPCs, or VNets is Warning or Critical then the card for that service will be outlined in red.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(173).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A57Z&se=2026-05-12T09%3A34%3A57Z&sr=c&sp=r&sig=4ymRxofdSBL8krn5Y0NcO4uCN3uhtYurrxzXJ3O6VL0%3D)

*The tables in the Services list shows the regions and VPCs that use the services that you're monitoring with synthetic testing.*

#### Services Table

The card for each service includes a table in which traffic related to the service is broken down into VPCs (AWS) or VNets (Azure) grouped by region, each of which is represented in a heading row. This table includes the following columns:

* **Expand/collapse**: An arrow that you click on to expand or collapse the VPC rows for a region.
* **Region**:

  + **Region rows**: The name of the region, with the number of VPCs given in parentheses.
  + **VPC/VNet rows**: The name and ID of the VPC/VNet.
* **Connection Type**:

  + **Region rows**: The connection type and count of VPCs or VNets.
  + **VPC/VNet rows**: The connection type (public or private).
* **From Kbits/s**: The bitrate of traffic from the service to the region or individual VPC/VNet.
* **To Kbits/s**: The bitrate of traffic to the service from the region or individual VPC/VNet.
* **Status Code** (only for entities with an activated agent): The HTTP Return Code returned by AWS or Azure from requests to the monitored entity.
* **Avg HTTP Latency** (only for entities with an activated agent): The elapsed time (average during the current time range) from making an HTTP request to receiving the last byte of the response.
* **Actions** (vertical ellipsis): A popup menu with the following options:

  + **Synthetics** (only if an agent is activated in the region/VPC): **Configure Test** (region only) takes you to the **Test Settings Page** for the synthetics test in this region that has been activated via the **Update Tests** button, while **View Test Results** takes you to the **Test Details Page** for that test.
  + **Open Quick-view**: Opens (in a new tab) the Network Explorer details page for this region/VPC in Kentik's Core module (see **Core Details Page**).
  + **View in Kentik Map**: Opens (in a new tab) the **AWS Topology** or **Azure Topology** page in Kentik Map.
  + **View in Data Explorer**: Opens (in a new tab) **Data Explorer**, with the **Query** sidebar set to show the traffic for this region/VPC.
  + **Copy to Clipboard**: Copies the region name or VPC ID to the clipboard.

### Configure Service Monitoring

The Configure Service Monitoring dialog, which opens from the Settings button (gear icon), lets you choose the AWS or Azure services for which you'd like to enable performance monitoring via synthetic testing. The dialog is structured as two columns:

* **Available Services**: A list of services that Kentik has identified as being used by your AWS or Azure cloud resources. Use the Search field to filter the list to services whose name matches entered text.
* **Selected Services**: A list of the services that have been chosen from the **Available Services** list. An **X** at the right of each service enables you to deselect that service.
* **Cancel**: A **Cancel** button at lower right and an **X** in the top right corner that close the dialog without saving changes.
* **Save**: A button that saves the services monitored and returns you to the Cloud Performance Monitoring page.

To use the dialog:

* To select all discovered services, click **All Services** in the **Available Services** list.
* To select individual services

  + In the **Selected Services** list, deselect **All Services**.
  + In the **Available Services** list, click each individual service that you'd like to select.
