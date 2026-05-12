# Traffic Costs

Source: https://kb.kentik.com/docs/traffic-costs

---

Kentik’s Traffic Costs workflow enables estimation and analysis of costs for various network traffic slices. It handles complex cost structures and traffic patterns, aiding in data-driven decisions to optimize network spending and maximize revenue.

![The Traffic Costs workflow enables you to create cost estimates and track them from a single view.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/TC-Main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A18Z&se=2026-05-12T09%3A42%3A18Z&sr=c&sp=r&sig=NKxTM2wPlRBL3nEg%2BZsHmFMOySay94desPcE554UDiE%3D)

*The Traffic Costs workflow enables you to create cost estimates and track them*   
*from a single view.*

> **Note**: To effectively use the Traffic Costs workflow, business information (e.g., suppliers and rates) must be entered. To restrict user access to its configuration pages, adjust the Connectivity Costs settings on the RBAC Roles & Permissions tab under Organization Settings » Users.

## About Traffic Costs

|  |  |
| --- | --- |
| **Purpose****:** | * Estimate and track edge traffic costs for specific ASNs or AS Groups * Assist NSPs in estimating traffic costs to/from any Customer Port or CDN Provider * Calculate costs by specific CDNs, OTT services and providers, geographic regions or markets |
| **Benefits:** | * Understand ingress/egress costs for slices of traffic like ASNs, OTT Providers, or geographic areas * Determine costs for downstream customers or upstream content |
| **Use Cases:** | * Identify costs for downstream customers * Save money by re-routing traffic from high-cost connections * Minimize costs for Content Providers on eyeball networks * Optimize peering strategies |
| **Relevant Roles:** | Network Planner, Peering Coordinator, Traffic Engineer, Broadband ISP Sales/FinOps |

The following features are available on the **Traffic Costs Page**:

* **Cost Estimation**: Estimate costs for various traffic slices, even across shared interfaces, by applying previously defined cost groups (see **About Cost Groups**). See **Understanding Cost Estimation**.
* **Traffic Slicing**: Analyze costs for different traffic slices like ASNs, AS Groups, AS Paths, or Customer Ports. See **Create a New Estimate Pane**.
* **Monthly Tracking and Snapshots**: Automatically track monthly cost trends and build a historical record of estimated charges. Use your snapshot credits to save estimates/snapshots and their associated flow data. See **Monthly Tracking**.
* **Visualizations**: Explore cost trends over time with the **Monthly Snapshot History Chart**.
* **Reporting**: Access detailed cost breakdowns of each estimate/snapshot in the **Traffic Costs Snapshots Table**.

### Understanding Cost Estimation

Kentik estimates your monthly traffic costs by integrating data from three sources to determine the cost of transporting traffic to or from another network:

* **Flow data**: Identifies traffic traversing an interface.
* **SNMP data**: Measures traffic volume through an interface.
* **Contract data**: Specifies cost per Mbps for an interface’s traffic, as defined in cost groups (see **About Cost Groups**).

The process Kentik takes to estimate your costs includes the following:

* For each interface in a cost group, Kentik retrieves the current price per Mbps from the Connectivity Cost workflow.
* If the traffic follows the billing direction, the cost is calculated by multiplying the peak Mbps by the price per Mbps. Traffic that is not following the billing direction is considered non-billed.

> **Note**: These estimates depend on the level of detail included in your provided network information, so they might not perfectly align with actual costs.

## Enable Traffic Costs

To enable Traffic Costs, Connectivity Costs must be enabled (see **Enabling Connectivity Costs**). Once you’ve enabled Connectivity Costs, navigate to the Traffic Costs page (in the main menu, under Edge) to **Create a New Estimate Pane**.

> **Note**: Adding external interfaces to multiple cost groups may affect Traffic Costs calculations.

## Traffic Cost Estimates

When you manually create an estimate for traffic costs, Kentik uses the current month’s billing information up until the date the estimate was generated. To compare costs over time, save an individual estimate with its associated flow data as a single snapshot or automatically capture estimated monthly traffic costs snapshots (see **Monthly Tracking**). Each snapshot created (manually or automatically) requires one snapshot credit.  
  
All snapshots are listed in the **Traffic Costs Snapshots Table** and viewable individually on each snapshot’s **Traffic Costs Details Page**.

### Estimate Traffic Costs

To manually estimate your traffic costs (rather than automatically generating a **Monthly Tracking** snapshot), follow these steps:

1. In the **Create a New Estimate** pane of the **Traffic Costs Page**, select a traffic type from the categories (ASN, OTT, Geography) on the left side. The list of traffic slices on the right side updates based on your selection.

   ![Select the traffic slice for which you want a cost estimate created.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/TC-Create-Estimate.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A18Z&se=2026-05-12T09%3A42%3A18Z&sr=c&sp=r&sig=NKxTM2wPlRBL3nEg%2BZsHmFMOySay94desPcE554UDiE%3D)

   *Select the traffic slice for which you want a cost estimate created.*
2. Enter search terms as needed to locate the desired slice in the list.
3. Click **Estimate** to create the estimate and open the **Traffic Costs Details Page** for the selected slice.

> **Notes**:
>
> * In order to save the estimate to view in the future, you must **Create a Snapshot** from the estimate.
> * You can view, share (internal/public share or email), or export (visual report pdf) an estimate from the **Traffic Costs Details Page** before or after saving it as a snapshot

### Create a Snapshot

To create snapshot of your traffic costs, follow these steps:

1. Create an estimate of your traffic costs by following the steps in **Estimate Traffic Costs**.
2. On the **Traffic Costs Details Page**, save the snapshot by completing one of the following options:

   1. Click **Save** in the banner at the top of the screen.
   2. Click **Save Snapshot** in the subnav.

![After clicking the Estimate button, a screen shows the details along with an option to save a snapshot.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/TC-Details-Estimate.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A18Z&se=2026-05-12T09%3A42%3A18Z&sr=c&sp=r&sig=NKxTM2wPlRBL3nEg%2BZsHmFMOySay94desPcE554UDiE%3D)

*After clicking the Estimate button, a screen shows the details along with an option to save a snapshot.*

To start monthly tracking, see **Enable Monthly Tracking**.

> **Notes**:
>
> * Manually-created snapshots are always for the current month and year.
> * The process of saving one snapshot uses one snapshot credit, even if the snapshot is later removed (see **Remove a Snapshot**).
> * You can save multiple snapshots for the same traffic slice without monthly tracking. However, if you take more than one snapshot in a month for the same slice, only keep the latest one. It will be more accurate due to having more monthly data.

### View a Snapshot

To view an existing snapshot on the Traffic Costs page, follow these steps:

1. Click the name of the traffic slice in the **Traffic Costs Snapshots Table** to open the snapshot’s **Traffic Costs Details Page**.![View your cost estimates in the Traffic Costs Snapshots table.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/TC-Snapshots-Table.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A18Z&se=2026-05-12T09%3A42%3A18Z&sr=c&sp=r&sig=NKxTM2wPlRBL3nEg%2BZsHmFMOySay94desPcE554UDiE%3D)
2. Click a snapshot date (displayed along the horizontal axis) in the **Traffic Cost History Chart** to view its ingress and egress data. The selected date will be highlighted in blue, and its data will be displayed below the **Traffic Cost History Chart** in the relevant tab(s) and table(s).

   ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/TC-TC-History-Table(2).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A18Z&se=2026-05-12T09%3A42%3A18Z&sr=c&sp=r&sig=NKxTM2wPlRBL3nEg%2BZsHmFMOySay94desPcE554UDiE%3D)

> **Note**: You can view total estimated costs for a saved snapshot in the **Monthly Snapshot History Chart** even if monthly tracking is not enabled.

### Remove a Snapshot

When you remove a traffic slice from the Traffic Costs Snapshots table, you remove all snapshots related to that traffic slice and prevent future monthly tracking for the snapshot.  
  
 To remove a traffic slice and its snapshots:

1. On the Traffic Costs page, select a traffic slice from the **Traffic Costs Snapshots Table** to open that its **Traffic Costs Details Page**.
2. On the subnav, click **Actions**, and from the dropdown menu, select **Remove**.
3. In the confirmation dialog, click **Remove** to remove all snapshots associated with this traffic slice and return to the Traffic Costs page.

> **Note**: Removing a traffic slice or its snapshot does not restore the credits that were spent generating those snapshots.

## Monthly Tracking

Enable monthly tracking of a traffic slice from its Traffic Costs Details Page. This creates a new snapshot at 04:00 UTC on the first day of the next month and estimates the previous month’s traffic costs.

### Enable Monthly Tracking

To enable monthly tracking:

1. Click the name of a traffic slice in the **Traffic Costs Snapshots Table** to open its Traffic Costs Details page.
2. In the subnav, click**Start Monthly**. A confirmation message will briefly appear at the top of the page confirming that monthly tracking is enabled.

   ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(326).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A18Z&se=2026-05-12T09%3A42%3A18Z&sr=c&sp=r&sig=NKxTM2wPlRBL3nEg%2BZsHmFMOySay94desPcE554UDiE%3D)
3. Kentik will capture 12 monthly snapshots per year for the selected traffic slice.

### Pause Monthly Tracking

To pause monthly tracking of a traffic slice:

1. Click the name of a traffic slice in the **Traffic Costs Snapshots Table** to open its Traffic Costs Details page.
2. In the subnav, click **Pause Monthly**. A confirmation message will briefly appear at the top of the page confirming that monthly tracking is disabled.

## Traffic Costs Page

The Traffic Costs page offers a concise view of cost estimates, history, and tools.

### Traffic Costs Page UI

The Traffic Costs page has the following UI elements:

* **Breadcrumbs** (subnav): An indicator of your current location within the Kentik portal. As you drill down deeper, you can click on a breadcrumb to go back to a higher level.
* **Share** (subnav): Opens a dialog to share a page link (internal/public share) or email a PDF summary.
* **Actions** (subnav): A button that allows you to export a PDF of the information on the Traffic Costs page.
* **Title bar**: Displays beneath the subnav and includes the following controls:

  + **Favorite** (star icon): A star to the left of the page name that allows you to add this page to the Favorites tab of the portal search (see **Portal Search Tabs**).
  + **Title**: Traffic Costs (the name of the page).
  + **Manage Providers**: A button that opens the **Manage Providers Page**.
* **Info pane**: A brief overview of how traffic costs work.
* **Create a New Estimate**: A pane where you can select a traffic slice and generate an estimate for that slice’s traffic costs for the month (see **Create a New Estimate Pane**).
* **Monthly Traffic Cost Snapshot History**: A chart that shows the total estimated traffic costs for your saved monthly snapshots (see **Monthly Snapshot History Chart**).
* **Traffic Cost Snapshots**: A table showing a list of your saved snapshots (see **Traffic Costs Snapshots Table**).

### Create a New Estimate Pane

Use this pane to create an estimate for your traffic costs. It includes the following UI elements:

* **Traffic Type Selector**: Select a traffic type from the list on the left, grouped by category:

  + **ASN**

    - **Source/Destination ASN**: The ASN of the network originating (source) or receiving (destination) the traffic.
    - **Source/Destination AS Group**: A logical grouping of ASNs  (e.g., Akamai, Google, Comcast).
    - **ASN in Destination AS Path**: Any ASN that appears in the BGP AS path to the destination, useful for understanding routing and transit relationships.
    - **Customer Port**: The physical or logical port on a network device through which customer traffic enters or exits the network.
    - **CDN Provider**: The Content Delivery Network (CDN) responsible for delivering the content associated with the traffic flow (e.g., Akamai, Fastly).
  + **OTT**

    - **OTT Service**: An individual over-the-top (OTT) content service (e.g., A&E TV, BBC).
    - **OTT Service Category**: The type of content provided by the OTT service (e.g., Ads, Cloud, Gaming, Video).
    - **OTT Service Provider**: The company or entity that operates the OTT service (e.g., Apple, CBS).
  + **Geography**: The country, region, or city associated with the source or destination of the traffic.

    - Country
    - Region
    - City
* **Searchable List of Traffic Slices**: List on the right that updates based on the selected traffic type. Enter search terms to filter the list as needed to find your traffic slice.
* **Estimate**: Located on each row of the traffic slices list, click the button to generate the cost estimate and open the **Traffic Costs Details Page** for the traffic slice.

### Monthly Snapshot History Chart

The Monthly Traffic Cost Snapshot History chart shows total estimated traffic costs for your saved monthly snapshots, and has the following elements:

* **Horizontal axis**: Month and year of snapshots.
* **Vertical axis**: Estimated cost range in the currency defined for the cost group. The range auto-adjusts based on snapshot totals.
* **Plot line**: Color-coded lines for each traffic slice with shape-coded points for each snapshot. Hover on a snapshot to see its month/year, name, and total estimated cost.
* **Legend**: Displays the name of each traffic slice shown in the chart with its associated line color and shape-coded points.

  + Below the table, click a name to toggle that traffic slice’s visibility in the chart. Excluding all snapshots hides axis values.

### Traffic Costs Snapshots Table

This table offers a concise view of your saved snapshots with the following sortable columns:

* **Name**: The traffic slice value selected from the Filter Value dropdown in the **Create a New Estimate Pane**. Click to access the Traffic Costs Details page for that traffic slice.
* **Slice**: Type of traffic slice (e.g., “AS Group”).
* **Snapshots**: Number of saved snapshots for this particular traffic slice.
* **Last checked**: Time since the last calculation.
* **Ingress Mbits/s**: Ingress traffic rate to the interfaces in the cost group for this estimate.
* **Egress Mbits/s**: Egress traffic rate from the interfaces in the cost group for this estimate.
* **Saved by**: User who last saved a snapshot for this traffic slice.
* **Total cost**: The estimated traffic costs calculated for the latest snapshot.
* **Automation**: Monthly snapshot automation status (“Monthly” with pause icon or “None” with play icon).

## Traffic Costs Details Page

The Traffic Costs Details page (accessible from the Traffic Costs page by creating a new estimate or by clicking a link in the Name column of the **Traffic Costs Snapshots Table**) provides in-depth cost estimates for a specific traffic slice.

### Traffic Costs Details UI

The Traffic Costs Details page has the following UI elements:

* **Breadcrumbs** (subnav): An indicator of your current location within the Kentik portal. As you drill down deeper, you can click on a breadcrumb to go back to a higher level.
* **Start/Pause Monthly** (subnav): Toggle (enable/disable) monthly snapshot capture for the displayed traffic slice.
* **Share** (subnav): A button that allows you to open the Share dialog so that you can share a page link (internal/publish share) or email a PDF summary of the page’s data.
* **Actions** (subnav): A button that opens a dropdown menu where you can choose from the following options:

  + **Recompute**: Update cost analysis after cost group configuration changes (auto-updates daily).
  + **Export**: Export a PDF of the estimate or snapshot data.
  + **Remove**: Remove all snapshots for this traffic slice and return to the Traffic Costs page (snapshot credits that were used to generate the removed snapshots are not restored).
* **Title bar**: Displays beneath the subnav and includes the following controls:

  + **Favorite** (star icon): A star to the left of the page name that allows you to add this page to the Favorites tab of the portal search (see **Portal Search Tabs**).
  + **Logo**: If applicable, the logo for the traffic slice value (e.g., a company).
  + **Title**: The traffic slice value for the estimate or snapshot (e.g., the AS Group name).
  + **Lozenge** (AS-related slices only): Provides details for an ASN (click to see **ASN Links Popup**), an AS Group, or an ASN in Destination AS Path.

    - Hover over an AS Group lozenge to see a popup displaying all of the ASNs within that group. On the popup, click any ASN lozenge to see its associated **ASN Links Popup**.
* **Traffic Costs Metrics**: A pane that displays beneath the title bar and contains a summary of the data from traffic slice’s estimate or snapshot (see **Traffic Costs Metrics Pane**).
* **Traffic Cost History**: A chart that displays a traffic slice’s traffic costs snapshots over time (see **Traffic Cost History Chart**).
* **Directional diagrams**: Two tabs (Ingress and Egress) each display a **Sankey Diagram** of the snapshot’s respective traffic. If no traffic is available, “No Results” will display.
* **Directional Traffic tables**: Two collapsible tables (Ingress and Egress) each display the device-level traffic and costs for the month (see **Directional Traffic Tables**).
* **Total Cost**: Displays the totals for Cost/Mbps, Ingress Mbits/s, Egress Mbits/s, and Est. cost of the traffic tables at the bottom of the page.

### Traffic Costs Metrics Pane

This pane summarizes your organization’s estimated costs and traffic volume for a specific traffic slice in the selected monthly snapshot. It contains the following information:

* **Estimated Cost**: Total estimated cost for this traffic slice for the selected monthly snapshot.
* **Cost per Mbps**: Estimated cost (per Megabit per second) for this traffic slice.
* **Ingress traffic volume**: Total incoming traffic volume.
* **Egress traffic volume**: Total outgoing traffic volume.
* **Billing Direction**: A tag indicating whether ingress or egress traffic had greater volume.
* **Trend indicator**: Shows percentage change from the previous month; red for increase, green for decrease.

### Traffic Cost History Chart

Visualize your estimated costs over time for a traffic slice. The chart has the following elements:

* **Monthly snapshots**: Horizontal axis with month/year buttons.
* **Cost**: Vertical axis with estimated costs. A trend line connects the data points, taller green segments indicate higher costs. Hover for a summary of metrics.

  > **Notes**:
  >
  > + The latest monthly snapshot displays by default. Selecting another snapshot in the timeline updates the page with that month’s data.
  > + This chart is only visible when accessing a **Traffic Costs Details Page** by clicking a snapshot name in the **Traffic Costs Snapshots Table** (on the **Traffic Costs Page**). It is not present for a newly generated estimate or snapshot.

### Directional Traffic Diagrams

Two tabs (Ingress and Egress) allow you to toggle between the snapshot’s ingress and egress traffic views. By default, the tab that is designated as the billing direction is displayed. If no traffic is available, “No Results” will display. On each tab, the following elements display:

* **View Flow in Data Explorer**: A button that opens Kentik’s Data Explorer in a new tab, pre-populated with this traffic slice’s data allowing you to analyze the traffic flow in detail.
* **Diagram**: Each Ingress and Egress **Sankey Diagram** provide a visual breakdown of traffic for the displayed monthly snapshot. Each “column” of the diagram represents different dimension combinations for visual traffic grouping which may differ between traffic slice types. For example, the columns for an ASN or AS Group diagram are:

  + **Device**: Device name.
  + **Src/Dst Interface**: Source or destination interface name.
  + **Src/Dst Provider**: Source or destination provider name (e.g., “equinix”).
  + **Src/Dst Connectivity Type**: Interface connection type (e.g., “Transit” or “Backbone”; see **Understanding Connectivity Types**).
  + **Src/Dst AS Number**: Source or destination ASN or AS Group (e.g., “Amazon”).

    ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(328).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A18Z&se=2026-05-12T09%3A42%3A18Z&sr=c&sp=r&sig=NKxTM2wPlRBL3nEg%2BZsHmFMOySay94desPcE554UDiE%3D)

> **Note**: Hover over nodes or connections for detailed popups.

### Directional Traffic Tables

By default, the billing direction determines which table (Ingress or Egress) is expanded and which is collapsed when the snapshot is first viewed. The table billing direction is expanded, and the other directional table is minimized. Each table details either the ingress or egress traffic for the selected month and is sortable by column headers.

#### Directional Traffic Header

The following elements are available in the header row of the traffic tables:

* **Snapshot month**: A red lozenge that indicates the month and year the estimate or snapshot was generated.
* **Snapshot name**: The name derived from the direction and current traffic slice.
* **Hide empty cost groups** (Ingress table only): A box that, once checked, hides all empty cost groups in both the Ingress and Egress tables.
* **How did we calculate this?** (Ingress table only): An informational popup that provides more information about how Kentik calculates the traffic costs.
* **Expand/collapse**: An icon at the far right of the header expands or collapses each individual table.

> **Note****:** The non-billing direction table is minimized by default. Use the expand/collapse icon to show or hide either table.

#### Directional Traffic Columns

The Ingress and Egress traffic tables contain the following column headers:

* **Device**: Name of the device; click to open the Details page for this device (see **Core Details Pages**).
* **Interface**: Name of the interface; click the link to open the Details page for this interface (see **Core Details Pages**). The interface’s alias is displayed below its name.
* **Connectivity Type**: Type of connection for the interface (see **Understanding Connectivity Types**); click the link to open the aggregate page for that connectivity type (see **Core Aggregate Pages**.
* **Provider**: Name of the provider (typically matches your monthly bill); click the link to open the Details page for that provider (see **Core Details Pages**).
* **Cost Group**: Provider cost group for the device/interface combination; click to open the **Cost Details Page** for this cost group.
* **Cost/Mbps**: Cost per megabit per second of traffic volume defined for this cost group.
* **Ingress Mbits/s**: Traffic volume entering the interface.
* **Egress Mbits/s**: Traffic volume exiting the interface.
* **Est. cost**: The device’s estimated total cost for the month.
* **Ingress/Egress Total**: The last row of each table shows the ingress/egress totals for Cost/Mbps, Ingress Mbits/s, Egress Mbits/s, and Est. cost.

> **Note**: An orange triangle icon with an exclamation mark indicates that the listed interface is not part of a cost group. Hover over the icon to display a popup with more information and an Add a Cost Group button that enables you to open the **Manage Providers Page** in order to add a cost group

### Total Cost Pane

Displays the totals for Cost/Mbps, Ingress Mbits/s, Egress Mbits/s, and Est. cost of the traffic tables at the bottom of the page. These are the same totals displayed in the **Traffic Costs Metrics Pane**at the top of the page.

---

© 2014-25 Kentik
