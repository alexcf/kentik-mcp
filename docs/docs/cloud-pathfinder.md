# Cloud Pathfinder

Source: https://kb.kentik.com/docs/cloud-pathfinder

---

This article covers the **Cloud Pathfinder** page in the Kentik portal.

> **Note:** For general information about monitoring your cloud resources with Kentik, start with **Cloud Overview**.        

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(174).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A46Z&se=2026-05-12T09%3A34%3A46Z&sr=c&sp=r&sig=Ag13STw5NVr78gZSaCgwh5L3OU5zxlR0GwuRkP0sWn8%3D)

*The Cloud Pathfinder page, which helps troubleshoot cloud connection problems*

## About Cloud Pathfinder

|  |  |
| --- | --- |
| **Purpose:** | Analyze AWS or Azure metadata within your cloud environments to monitor the forward and returns paths between two points. |
| **Benefits:** | * Provides a detailed map of the connection, highlighting any issues causing the connection to fail. * Offers a comprehensive view of routing and NACL/SG configurations. * Facilitates faster troubleshooting by providing direct links to misconfigured gateways or security policies. |
| **Use Cases:** | * Monitoring cloud network paths for AWS and Azure. * Ad hoc network connectivity tests within your cloud environments. |
| **Relevant Roles:** | Network Engineer |

Cloud Pathfinder analyzes AWS or Azure metadata collected from Kentik-monitored cloud environments and inventories subnets, instances, and VPCs to determine how they communicate. It then visualizes the forward and return paths between two points on your cloud network, helping you identify potential issues like dropped connections so that your network teams can quickly fix them.

> **Note**: Cloud Pathfinder does not establish network connections. Instead, it analyzes metadata from your cloud environment to simulate and illustrate the communication paths between two points.

### Cloud Pathfinder Workflow

The workflow is composed of two main parts:

* **Create Report**: Use the **Create Report Page** to define parameters for your analysis. You can run an ad hoc analysis or save it as a report for future reference.
* **View Report**: Saved reports appear in the **Previous Reports** table. Each listed report links to an individual **Pathfinder Report Page** that provides a detailed connection analysis, including visualizations.

## Cloud Pathfinder Page

The Cloud Pathfinder for AWS and Azure landing page enables access to the Cloud Pathfinder Reports page for each of the Kentik-supported cloud providers. The page includes the following UI elements:

* **Preview**: A tabbed pane showing a preview of the different Cloud Pathfinder pages:

  + **Report List**: An image showing one of the **Cloud Pathfinder Reports Page**.
  + **Report Detail**: An image showing a **Pathfinder Report Page**.
* **Description**: A description of Cloud Pathfinder.
* **Show AWS Pathfinder Report**: A button that takes you to the AWS Pathfinder Reports page.
* **Show Azure Pathfinder Report**: A button that takes you to the Azure Pathfinder Reports page.

## Cloud Pathfinder Reports Page

The Cloud Pathfinder Reports pages for AWS and Azure each include the same main UI elements:

* **Favorite**: A star to the left of the page title that allows you to add this page to the **Favorites Tab** of the **Portal Search**.
* **Description**: A description of how you can use the Pathfinder report.
* **Create Report**: A button that takes you to the **Create Report Page**.
* **Previous Reports**: A table that lists your **Previous Reports**.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(175).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A46Z&se=2026-05-12T09%3A34%3A46Z&sr=c&sp=r&sig=Ag13STw5NVr78gZSaCgwh5L3OU5zxlR0GwuRkP0sWn8%3D)

  *The AWS Pathfinder Reports page.*

#### Previous Reports

The **Previous Reports** table lists the reports that you have previously generated and saved. Click the heading of any column to sort the list, which includes the following elements:

* **Reachability**: Indicates whether there is a connectivity issue between the source and destination:

  + **Reachable** (green checkmark): Connection is successful between the source and destination.
  + **Not reachable** (red X): Connection is unsuccessful between the source and destination.
* **Report Name**: The name of the report.
* **Source**: The source entity from which the connection path originates.
* **Destination**: The destination entity at which the connection path terminates.
* **Last Analysis Date**: The date the connection analysis was most recently performed.
* **Topology Date**: The end time of the time range covered by the report (see **Pathfinder Report UI**).
* **View Report**: A button to take you to the **Pathfinder Report Page**.
* **Remove** (trash icon): Opens a confirmation popup from which you can remove the report from the list of saved reports.

## Create Report Page

The **Create Report** page is used to run a connectivity check between two endpoints. The page contains the following UI elements:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(176).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A46Z&se=2026-05-12T09%3A34%3A46Z&sr=c&sp=r&sig=Ag13STw5NVr78gZSaCgwh5L3OU5zxlR0GwuRkP0sWn8%3D)**Name**: The name of this report. If left blank, a report name will be autogenerated using the source and destination values.
* **Source type**: A drop-down that determines the type of the sources listed in the **Source** drop-down. Source types vary by cloud provider as follows:

  + **Network Interfaces**(AWS only): AWS network interfaces.
  + **Subnets**: AWS or Azure subnets.
  + **Instance**(AWS only): AWS instances.
* **Source**: A filterable drop-down listing cloud entities of the type selected with **Source Type**. Each entity is represented by a card that includes its name (in bold) and additional **Entity Details** that vary depending on the **Source Type**. Click on a card to choose the entity to use as the start point of the connectivity check.
* **Target Type**: A drop-down that determines the type of the destinations listed in the **Destination** drop-down.
* **Destination**: A filterable drop-down listing cloud entities of the type selected with **Target Type**. Each entity is displayed on a card that includes its name (in bold) and additional **Entity Details** that vary depending on the Target Type. Click on a card to choose the entity to use as the destination of the connectivity check.
* **Time Range**: A drop-down that pops up the **Custom Time Range Settings**.
* **Protocol**: A filterable drop-down listing protocols (e.g., TCP and UDP). Click a protocol to use it for the connection check.
* **Port**: A field in which to enter the port used to establish the connection.
* **Cancel**: A button to cancel the report creation and exit the page.
* **Run**: A button to execute the analysis and generate the report, which opens in the **Save Report Page**.

#### Entity Details

As shown in the table below, the details included in an entity's card in the **Source** or **Destination** drop-down vary depending on the type of entity selected in the associated **Type** drop-down.

| Detail | Network Interfaces | Subnets | Instance | Description |
| --- | --- | --- | --- | --- |
| Name | Yes | Yes | Yes | The name of the source or destination. |
| PrivateIpAddress | Yes | No | Yes | The private IP address. |
| PublicIpAddress | No | No | Yes | The public IP address. |
| PrivateDnsName | No | No | Yes | The private DNS name. |
| PublicDnsName | No | No | Yes | The public DNS name. |
| SubnetId | Yes | No | Yes | The subnet from which the connection path originates or terminates. |
| VpcId | Yes | Yes | Yes | The Virtual Private Cloud from which the connection path originates or terminates. |
| CidrBlock | No | Yes | No | The CIDR block from which the subnet originates or terminates. |

## Save Report Page

The Save Report page is shown immediately after a report is created on the Create Report page. Except for an added **Save Report** button, the contents of this page will be the same as a **Pathfinder Report Page**. Click the button to save the report and display the Pathfinder Report page. The report will be added to the **Previous Reports** list for the cloud provider.

## Pathfinder Report Page

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(177).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A46Z&se=2026-05-12T09%3A34%3A46Z&sr=c&sp=r&sig=Ag13STw5NVr78gZSaCgwh5L3OU5zxlR0GwuRkP0sWn8%3D)

*The Pathfinder report page shows the results of a connectivity check.*

The Pathfinder Report page includes the following main UI elements:

* **Pathfinder report information**: A set of fields across the top of the page that provide information about what was checked for the report (see **Pathfinder Report Information**).
* **Kentik AI Summary**: A pane containing Kentik AI's summary of the Pathfinder report (see **AI Connection Summary**).
* **Paths**: Depictions of the sequence of entities traversed by traffic on the checked connection (see **Connection Paths**):

  + **Forward**: Traversed entities from source to destination.
  + **Return**: Traversed entities from destination to source.

### Pathfinder Report Information

The following fields across the top of the Pathfinder Report page provide information about what was checked for the report:

* **Source**: The type (e.g., subnet) and name of the source entity from which the connection originates.
* **Destination**: The type and name of the destination entity at which the connection terminates.
* **Port**: The port used to establish the connection.
* **Protocol**: The protocol used to establish the connection.
* **Last Run**: The timestamp for when the report was created.
* **Start Time**: The timestamp for the start of the connection assessment period.
* **End Time**: The timestamp for the end of the connection assessment period.

> **Note:** The names shown in the **Source** and **Destination** fields are links that open the entity's Network Explorer Details page in a new browser tab (see **Core Details Pages**).

### AI Connection Summary

The pane at the left of the Pathfinder Report page contains a summary by Kentik AI of the connection between the selected source and destination entities. The summary includes the following parts:

* **Overview**: A statement of the overall conclusion of the analysis.
* **Analysis**: A more detailed breakdown of the connection.
* **Resolution**: A statement of actions, if any, that could be taken to resolve any issues revealed by the connectivity check.
* **Close Summary**: A button that hides the summary pane.

### Connection Paths

The **Forward Paths** and **Return Paths** panes each include the following UI elements:

* **Path**: A sequence of **Transit Points** for the connection, from source to destination or vice versa.
* **Connection Visualization**: A set of shaded rectangles representing a region and containing that region’s path entities (see **Connection Visualization**).

#### Transit Points

Each transit point on the path contains the following information:

* **Type**: The entity type for the transit point.
* **ID**: The ID of the entity.
* **Copy**: A button that copies the ID to the clipboard.

> ***Note:***
>
> If the connection check discovers an issue involving a transit point, the transit point will be marked with a red caution icon (triangle with exclamation point). Hovering on the icon will pop up a tool tip that indicates the nature of the problem.  

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(178).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A46Z&se=2026-05-12T09%3A34%3A46Z&sr=c&sp=r&sig=Ag13STw5NVr78gZSaCgwh5L3OU5zxlR0GwuRkP0sWn8%3D)

#### Connection Visualization

This visualization represents the sequence of paths traversed by the connection, and includes the following UI elements:

* **Region**: Shaded rectangles that segment the visualization by region and show when a connection crosses regional boundaries.
* **VPC** (AWS only): Outlined rectangles within regions that group the connection points by VPC and show any entry or exit points.
* **VN** (Azure only): Outlined rectangles within regions that group the connection points by their VN.![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(179).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A15%3A46Z&se=2026-05-12T09%3A34%3A46Z&sr=c&sp=r&sig=Ag13STw5NVr78gZSaCgwh5L3OU5zxlR0GwuRkP0sWn8%3D)
* **Transit Point**: A rectangle for each entity. Click to open a popup with options for more information:

  + **AWS Map**: Takes you to the AWS **Cloud Topology** view of **Kentik Map**.
  + **Azure Map**: Takes you to the Azure **Cloud Topology** view of **Kentik Map**.
  + **Show in AWS Console**: Takes you to the AWS console (sign-in page unless you are already signed into the AWS account).

> **Note:** If the connection check discovers an issue involving a transit point, the transit point's outline will be red.

## Manage Reports

The following topics describe how to manage Cloud Pathfinder reports.

### Run a Pathfinder Report

You can start a Pathfinder report from two locations in the Kentik portal:

* Directly from the **Cloud Pathfinder Reports Page** of a given cloud provider (AWS or Azure); or
* From the provider's **Cloud Topology** view in the **Kentik Map**.

Once run, a report started in either location can be saved by clicking the **Save** button on the **Save Report Page**. The saved report will open a **Pathfinder Report Page**, and the report will be added to the **Previous Reports** list on one of the **Cloud Pathfinder Reports Page** for the cloud provider.

#### Run from Cloud Pathfinder

To run a Pathfinder report from Cloud Pathfinder:

1. On the portal main menu, click **Cloud Pathfinder** under **Cloud**.
2. On the Cloud Pathfinder for AWS and Azure page, click **Show Pathfinder Report** for either AWS or Azure.
3. On the **Cloud Pathfinder Reports Page**, click **Create Report** to open the Create Report page.
4. Specify the values of the fields in the **Create Report Page**.
5. Click **Run** to run the analysis, which will open on the **Save Report Page**.

#### Run from Kentik Map

To run a Pathfinder report from the Kentik Map:

1. On the portal main menu, click **Kentik Map** in the list of featured modules at the left.
2. In the **Clouds** block at the upper left of the Kentik Map, click the cloud provider (AWS or Azure) and choose **View Topology** from the resulting popup.
3. On the **Cloud Topology** view, click a VPC or subnet object and choose **Cloud Pathfinder** from the resulting popup. This will take you to the **Create Report Page** with the **Source type** and **Source** fields already specified.
4. Specify the values of the remaining fields (see **Create Report Page**).
5. Click the **Run** button to run the analysis, which will open on the **Save Report Page**.

### Save a Pathfinder Report

To save a Pathfinder report:

1. Complete the steps to**Run a Pathfinder Report**.
2. On the **Save Report Page**, click **Save Report**.

> **Note**: If you do not save the report, it will not be available on the **Previous Reports** table of the Cloud Pathfinder Reports page.

### Remove a Pathfinder Report

To remove a saved **Cloud Pathfinder** report:

1. On the portal main menu, click **Cloud Pathfinder** under **Cloud**.
2. On the Cloud Pathfinder for AWS and Azure page, click the **Show Pathfinder Reports** button corresponding to the Cloud provider (AWS or Azure).
3. Locate the report in the **Previous Reports** list.
4. Click the **Remove** button (trash icon) at the right of the row.
5. In the confirmation popup, click **Remove** to delete the report. The popup will close, and the report will be removed from the list.
