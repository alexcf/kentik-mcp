# Connectivity Costs

Source: https://kb.kentik.com/docs/connectivity-costs

---

This article covers the Connectivity Costs module (Edge » **Connectivity Costs**) of the Kentik portal.

> **Note**:
>
> * Effective use of the Connectivity Costs module requires the entry of business information (e.g. suppliers and rates) that your organization may choose not to share with all of its Kentik users. To restrict access by any given user to the module or to its configuration pages, use the Connectivity Costs settings on the Permissions tab of Settings » **Users**.
> * The tabs of the **Connectivity Costs** page are described in **Cost Breakdown Tabs**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(344).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)

*The Connectivity Costs landing page.*

## About Connectivity Costs

|  |  |
| --- | --- |
| **Purpose:** | Predict costs for external connectivity using provider pricing models and traffic volume measurements. |
| **Benefits:** | * Understand the drivers of overall network spend. * True-up provider invoices and discover billing errors. * Eliminate billing surprises with automatic cost forecasts. |
| **Use Cases:** | * Manage the network edge. * Understand the cost of sending or receiving traffic to or from particular destinations. * Understand the cost of working with particular peering partners or transit providers. * Understand the effective per megabit costs of traffic including fixed monthly or annual costs (such as facilities, IP address allocations, and cross connect costs). |
| **Relevant Roles:** | Network Strategist, Network Architect, Peering Coordinator |

The following topics provide an overview of the **Connectivity Costs** module.

### Connectivity Costs Use Cases

The **Connectivity Costs** module helps you to understand the cost impact of traffic entering or exiting your network via interfaces whose **Network Boundary Attribute** from Interface Classification is "external" (i.e. interfaces used for transit and/or peering). The use cases for this module include the following:

* Check the accuracy of billing statements from transit providers against your own traffic data.
* Visually detect changes in cost, and drill down to understand the root cause of unexpected changes.
* Expose current cost trends for use in network planning and forecasting.
* Surface historical cost trends across providers.
* Establish baselines and benchmarks for interconnection costs.

### Connectivity Costs Concepts

The **Connectivity Costs** module is based on the following basic concepts:

* **Provider**: A service provider that provides network connectivity services to customers (e.g., Century Link). In most cases a provider will correspond to a bill that you receive on a monthly basis.
* **Cost group**: A set of interfaces that connect to a given provider and share a set of common billing characteristics that are established in the contract that governs connectivity for those interfaces.
* **Cost Model**: The billing characteristics of a cost group, which include factors such as fixed periodic charges, the cost of bandwidth up to a given traffic volume, and the rate for volume in excess of that volume (e.g. tiers); for more detail see **Cost Group Model**.

  + A set of interfaces that share an aggregate bandwidth ("bundled") corresponds to a single cost group.
  + Each interface that is metered independently ("unbundled") corresponds to its own separate cost group.

Providers and cost groups are kept distinct to enable different ways of considering costs:

* Providers often differentiate service costs based upon geography, offering one price for circuits in a certain part of the world but pricing those services differently in a different market.
* A given provider may provide multiple types of services to the same user, e.g., a fixed cost transit circuit and also a circuit that is based upon a committed information rate.

The module works by allowing you to enter each of your organization's providers and to configure cost groups for those providers (typically one group per contract) such that they reflect your actual costs for connectivity. Based on this information as well as what we know (from SNMP polling and interface classification) about your utilization of the interfaces connected to your providers, we are able to project current costs and show historical cost trends.

> **Note:** Data entered for cost calculations in the Connectivity Costs modules is currently retained indefinitely (not subject to the retention limits governing flow data retention in your Kentik plans).

### Connectivity Traffic Calculation

The estimation of your costs for connectivity (transit/peering) is based on calculations that we perform on traffic volume data collected via SNMP polling from the interfaces that you have assigned to a cost group (via the **Group Interfaces Tab**). We use the following process to calculate traffic volume:

1. Every data source whose interfaces will be included in the connectivity costs calculation must have SNMP polling enabled (see **Enabling SNMP Polling**), with the polling interval set to Standard (see **SNMP Polling Intervals**). We poll those interfaces and keep a record, at five minute intervals (as per industry standard), of the increase in byte counter values, giving us a count of the bytes in and out during the interval.
2. At a regular interval we use the stored bytes-per-interval data to calculate each interface's in and out volume in Mbps for each five-minute interval.
3. We then find the in and out volume at a given percentile (see Metered Percentile in **Commit Model Settings**) of the volumes calculated for all intervals since the start of the current billing period, either individually for each interface or collectively for all interfaces in a cost group (see Computation Method in **Commit Model Settings**):

   1. By default the percentile is 95th but it may be set to something else if the interface is part of a cost group that uses a Commit cost model (see **Cost Group Model**).
   2. The billing period is specified in the **Provider Settings Pane**.
4. We then calculate the volumes stated in the portal:

   1. The overall ingress or egress volume stated on the **Connectivity Costs Page** is the sum of the volumes for all of the interfaces that have been associated with any cost group.
   2. The ingress or egress volume stated for a given "entity" — provider, connectivity type,  site, or site market — on a **Cost Details Page** is the sum of the volumes for each interface associated with that entity.

Once we arrive at projected traffic volumes we use the costing information you provide in the **Cost Group Details** pane to estimate connectivity costs for the current period.

> **Notes:**
>
> * When a cost group includes multiple interfaces you can choose the method of percentile computation; see Computation Method in **Commit Model Settings**.
> * Cost data for the current billing cycle is an estimate that is computed daily. In practice these figures tend to be effectively final about 36 hours before the end of the cycle. When each new cycle starts the prior cycle is stored as history and the current cost counters are reset.
> * Kentik retains historical cost data as-is indefinitely.

## Enabling Connectivity Costs

To set up the **Connectivity Costs** module:

1. Confirm that when you registered your network's data sources (Setup » **Data Sources**) you permitted Kentik to poll your network interfaces via SNMP V2 or V3 (see **Enabling SNMP Polling**), with the polling interval set to Standard (see **SNMP Polling Intervals**).
2. On the **Connectivity Costs Page** (Edge » **Connectivity Costs**), click the **Manage Providers** button at upper right. You'll be taken to the **Manage Providers Page**, where you'll be able to add a new provider or edit the configuration of an existing provider.
3. Click the **Add Provider** button at upper right, which takes you to a New Provider page (see **Provider Configuration**).
4. On the **Add Provider** page, fill in the Name, Type, and other fields of the **Details** pane, then click the Save button. You'll see a notice that the provider was added successfully, and the main display area of the page will now have a **Cost Groups Pane** with an **Add Cost Group** button.
5. Click the **Add Cost Group** button, which will take you to a Configure Cost Group page to set up your first cost group (see **Group Configuration Page**).
6. On the configuration page, fill in the **Cost Group Details** pane.
7. On the **Cost Model Tab**:

   1. Add any **Global Charges** for this cost group.
   2. If you selected a Commit cost formula in the group details, add any Cost Tiers.
8. On the **Group Interfaces Tab**, click the **Add/Edit Interfaces** button to open the **Configure Cost Group** dialog, enabling you to add interfaces to the cost group (the dialog is the same as in Capacity Planning; see **Configure Interface Groups**). When you close the dialog the interfaces you've added will be listed in the tables of the **Interfaces** tab.
9. If individual charges are associated with any of the interfaces, click the **Add Charges** link for those interfaces and use the form on the resulting **Interface Charges Drawer** to enter those charges.
10. When you've set both the details and interfaces for your cost group, click the Save button at the bottom of the **Cost Group Details** pane. You'll see a notice that the cost group was edited successfully, and you'll be returned to the configuration page for the provider to which you added the cost group. From there you can add additional cost groups as needed.
11. When you're done adding providers and cost groups, use the breadcrumbs at top to return to the **Connectivity** **Costs** page.

## Connectivity Costs Page

The Connectivity Costs page (Edge » **Connectivity Costs**) is covered in the following topics.

### Connectivity Costs UI

The page includes the following main UI elements:

* **Share** (on subnav): Opens the **Share** dialog to enable you to share the current view (see **Sharing via the Share Dialog**).
* **Actions** (on subnav): A drop-down menu with actions you can take involving the current view (see **Connectivity Costs Actions**).
* **Calendar** (on subnav): A button that toggles display of a sidebar containing the **Calendar Pane**, which shows date-related information pertaining to connectivity, including upcoming invoice due dates and contract renewal anniversaries.
* **Month selector**: A drop-down menu from which you can choose the month for which all the information on the page will be displayed. By default the month is set to the current month.
* **Manage Providers**: A button that takes you to the **Manage Providers Page**, where you can manage all of the providers via which traffic enters or exits your network.
* **Cost Metrics pane**: A set of indicators showing high level connectivity cost information (see **Cost Metrics Pane**).
* **Cost Breakdown tabs**: A set of tabs that each include information (charts and tables) related to different ways to break down your costs (see **Cost Breakdown Tabs**), e.g., by Providers, Connectivity Types, Sites, and Site Markets.

#### Connectivity Costs Actions

The following options are available from the Actions menu:

* **Export**: Prepares a visual report (PDF) or data table (CSV) to export as described in **Portal Export Options**. A notification appears when the report is ready to download.
* **Subscribe**: Opens the **Subscribe** dialog enabling you to create a subscription for this Connectivity Costs view. The form in the Subscription dialog is the same as on the **Subscription** tab of the **Share** dialog, which is covered in **Subscription Tab UI**.
* **Unsubscribe**: If you’re subscribed to this Connectivity Costs view, this option opens a dialog enabling you to confirm that you wish to be removed from the subscription.

### Cost Metrics Pane

The Cost Metrics pane displays high-level indicators that summarize your organization's projected connectivity costs for this billing cycle as recomputed daily based on your traffic since the start of the cycle. The pane includes the following indicators:

* **Estimated Cost**: Your total projected cost for transit and peering. This number is computed by summing the costs of all of the cost groups associated with all of your registered providers.
* **Cost per Mbps**: The estimated cost of bandwidth (in Mbps) across all providers, connectivity types, sites, and site markets (based on the cost information provided for each cost group; see **Cost Group Configuration**).
* **Egress** and **Ingress**: The total rate of inbound/outbound traffic as calculated across all of the interfaces identified in all of the cost groups in all of your providers (see **Connectivity Traffic Calculation**). These fields enable you to compare your utilization in each direction.

  > **Notes:**
  >
  > + The ingress/egress direction with greater traffic volume is indicated with a Billing Direction tag.
  > + The ingress/egress figures are not the actual maximum but instead reflect the metered percentile for each cost group (see **Commit Model Settings**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(349).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)

*Cost Metrics are high-level indicators that summarize your organization's connectivity costs.*

### Cost Breakdown Tabs

The cost breakdown tabs each include information (charts and tables) related to a different way of looking at your connectivity costs: by providers, by connectivity types, and by sites.

* **Monthly Costs**: A stacked column chart showing trends in your organization's costs per billing cycle over recent months for the top entities corresponding to the current tab (e.g., top providers on the **Providers** tab).
* **Effective Cost per Mbps**: A line chart showing trends over recent months in your organization's effective cost per Mbps (total paid for connectivity — including at the cost group and interface levels — divided by bandwidth actually used) via the top entities corresponding to the current tab (e.g., top sites on the **Sites** tab).
* **Entities list**: A table listing the individual entities that make up this view of your costs. The columns of the table are covered in **Entities List Columns**. Click on any row of the list to go to a **Cost Details Page** about the corresponding entity (see **Cost Details Entities**).

> **Note:** In the two month-based charts:
>
> * Bars are shown only for months for which data is available (up to 12).
> * The right-most bar, for the current month, is an estimate.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(350).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)

#### Cost Details Entities

A **Cost Details** page is available for entities of the following types:

* **Provider**: Providers as identified via interface classification (see **About Provider Classification**).
* **Connectivity Type**: Connectivity types as identified via interface classification (see **Understanding Connectivity Types**).
* **Site**: Sites as described in **About Sites**.
* **Site Market**: Site markets as described in **About Site Markets**.

#### Entities List Columns

Each breakdown tab includes a table that lists the entities in your organization that correspond to that tab (**Providers**,**Connectivity Types**, or**Sites**), and provides cost and volume estimates for the current billing cycle. The columns are the same for each of these tables:

* **Name**: The name of the provider, connectivity type, site, or site market.
* **Cost**: Estimated total cost for the current billing cycle for this provider, connectivity type, site, or site market.
* **Cost per Mbps**: Estimated cost per Mbps for the current billing cycle for this provider, connectivity type, site, or site market.
* **Ingress (Mbps)**: Estimated volume of ingress traffic for the current billing cycle for this provider, connectivity type, site, or site market.
* **Egress (Mbps)**: Estimated volume of egress traffic for the current billing cycle for this provider, connectivity type, site, or site market.

### Calendar Pane

The **Calendar** pane is a collapsible sidebar located at the right of the Connectivity Costs landing page. The pane appears in a drawer if the width of the browser window is less than 1600px. The **Calendar** pane contains the following UI elements:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(353).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)

* **Calendar**: A calendar for the current month, showing the due dates of any invoices (blue triangle) and renewal dates of any expiring contracts (orange triangle).
* **Today**: A field listing any providers whose invoices are due today.
* **Invoices**: A tab that displays a card for each provider with an invoices due in the next seven days (see **Calendar Provider Cards**).
* **Contracts**: A tab that displays a provider card for each contract with an anniversary in the next 14 days.

#### Calendar Provider Cards

When an invoice is due in the next seven days or a contract renews in the next 14 days then the corresponding tab (**Invoices** or**Contracts**) in the **Calendar** pane will include a provider card for the provider corresponding to that invoice or contract. The provider card includes the following fields:

* **Provider**: The name of the provider. Click the name to go the **Cost Details Page** for this provider.
* **Actions** (menu icon): A drop-down menu with actions you can take involving the current view (see **Connectivity Costs Actions**). The Actions menu additionally gives you the option to open the Share dialog (see **Sharing via the Share Dialog**).
* **Due date**: The due date of the invoice or renewal date of the contract.
* **Cost group**: The cost groups defined in the provider.
* **Amount**: For invoices, the amount due.
* **Trend**: For invoices, the change in amount from the prior billing cycle.

## Cost Details Page

Your organization's costs related to each "entity" — provider, connectivity type,  site, or site market — are detailed on a cost detail page.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(354).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)

*The top portion of a cost details page, including the tags, metrics pane, and rankings charts.*

### Cost Details Access

The cost details pages for an entity of a given type — provider, connectivity type, site, or site market — is accessed from the table at the bottom of the corresponding tab of the Connectivity Costs landing page (see **Cost Breakdown Tabs**). Click on any row of the table to go to a page detailing costs related to the corresponding entity.

### Cost Details UI

The cost details page for each entity includes the following main UI elements:

* **Month selector**: A drop-down menu from which you can choose the month for which all the information on the page will be displayed. By default, the month is set to the current month.
* **Configure Provider** (Provider pages only): A button that links to the Configure Provider page for this provider (see **Provider Configuration**).
* **Tags**: A set of lozenges providing information related to the provider, connectivity type, site, or site market (see **Cost Details Tags**).
* **Metrics pane**: A set of high level indicators providing information on your organization's connectivity costs related to the provider, connectivity type, site, or site market (see **Cost Details Metrics**).
* **Ranking charts**: A set of charts that show how the entity compares to other entities of the same type; see **Rankings Charts**.
* **Details Breakdown Tabs**: A set of tabs that each include graphs and tables detailing traffic and costs for this entity (see **Details Breakdown Tabs**).

### Cost Details Tags

The lozenges at the top center of a cost details page provide the following information related to the entity (provider, connectivity type, site, or site market):

* **Month/year**: The month and year of the current billing cycle (the one for which the information is being displayed).
* **Billing cycle** (Provider pages only): Indicates the number of days left in the current cycle.
* **Connectivity type** (Provider pages only): The provider's connectivity type as identified via **Interface Classification**.

### Cost Details Metrics

The **Metrics** pane on a cost details page includes the following high level indicators, which provide information about your organization's connectivity costs related to this provider, connectivity type, site or site market:

* **Estimated Cost**: Kentik's estimate of the cost for the current billing cycle, as well as the percentage change from the prior cycle.
* **Minimum Cost**: The minimum monthly spend based on the models (e.g., committed or flat rates) used for the cost groups associated with this provider, connectivity type, site, or site market (see **Cost Group Configuration**).
* **Cost per Mbps**: The estimated cost that will be paid per Mbps of bandwidth that is projected to be consumed in the cycle, as well as the percentage change from the prior cycle.
* **Traffic Costs Mbps**: The estimated charge per Mbps for traffic in excess of the volume covered by the committed or flat rates included in your minimum costs, as well as the percentage change from the prior cycle.
* **Egress** and **Ingress**: The estimated volume of traffic in each direction, as well as the percentage change from the prior cycle.

  > **Note:** The direction with greater traffic volume is indicated with a Billing Direction tag.

### Rankings Charts

The rankings charts show Kentik's estimate for the current billing cycle of how the entity detailed on this page compares with other entities in the same category (provider, connectivity type, site, or site market) based on the following metrics:

* **Volume**: The overall volume of your organization's traffic (egress and ingress) associated with this entity.
* **Cost**: Your total cost associated with this entity.
* **Cost per Mbps**: Your cost per Mbps for traffic associated with this entity.

The charts are color-coded and keyed. Click on an item in the key to mute that item in the chart.

### Details Breakdown Tabs

The details breakdown tabs each include information (charts and tables) related to a different way of looking at your connectivity costs associated with this individual provider, connectivity type, site, or site market. As shown in the table below, the available breakdown tabs vary by the category of the entity.

| Breakdown tab | Provider page | Connectivity Type page | Site page | Site Markets page |
| --- | --- | --- | --- | --- |
| Cost Groups | Y | Y | Y | Y |
| Providers | N | Y | Y | Y |
| Sites | Y | Y | N | Y |
| Site Markets | Y | Y | N | N |
| Connectivity Types | N | N | Y | N |

The cost details tabs include the charts and tables covered in the following topics.

#### Monthly Cost Charts

The charts in this pane show monthly trends for two cost metrics:

* **Monthly Costs**: A stacked column chart showing trends in your organization's costs per billing cycle over recent months for the top entities corresponding to the current tab (e.g., top cost groups on the **Cost Groups** tab).
* **Effective Cost per Mbps**: A line chart showing trends over recent months in your organization's effective cost per Mbps (total paid for connectivity — including at the cost group and interface levels — divided by bandwidth actually used) for traffic associated with the top entities corresponding to the current tab (e.g. top connectivity types on the **Connectivity Types** tab).

The **Monthly Costs** and**Effective Costs** charts are color-coded and keyed. Hover on an item in the key to solo that item in the chart, or click on an item to mute/unmute it.

> **Note:** Data points are shown only for months for which data is available (up to 12). The right-most data point is an estimate for the current month.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(355).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)

#### History Chart and Table

The **Traffic History** chart and accompanying table list the individual entities that make up the view shown on this tab, e.g., the cost groups on the **Cost Groups** tab of the details page for a given connectivity type. The chart is a line chart with time (in months) as the horizontal axis and two vertical axes for traffic volume: ingress at left and egress at right. The table also acts as a key for the chart. Hover on an item in the table to solo that item in the chart.  
  
The columns of the table are the same as those in the table on the cost breakdown tabs of the Connectivity Costs landing page (see **Entities List Columns**), with the addition of the following column in the table on the Cost Groups tab on the details page for a connectivity type, site, or site market:

* **Metered Percentile**: The percentile used for the metered percentile calculation on this interface (see **Commit Model Settings**).

The table also acts as a gateway to further details about the listed items. The way these details are presented depends on the details breakdown tab you are on when you click an item in the table:

* **Cost Groups tab**: Clicking on an item in the table will open the **Cost Group Details** drawer for that item.
* **Providers, Sites, Site Markets, or Connectivity Types tab**: Clicking on an item in the table will take you to the details page for that item.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(356).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)

### Cost Group Details

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(357).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)Cost group details are presented in a drawer that slides out from the right of the page when you click on a cost group row in the**Traffic History**table on a **Cost Groups** tab.

The drawer contains the following fields and controls:

* **Name**: The name of the cost group.
* **Export SNMP Samples**: A button that downloads a CSV file of SNMP traffic data associated with this cost group, both in and out at 5 minute intervals. Use the Month Selector from the Cost Details page to select the month for which you'd like to download data.
* **View in Data Explorer**: A button that takes you to a Data Explorer screen that is filtered to show the traffic associated with this cost group.

  > **Note:** The traffic shown in Data Explorer is based on flow data rather than SNMP.
* **Configure**: A button that takes you to the Configure Cost Group page for this cost group (see **Cost Group Configuration)**.
* **Cost Model**: A pane containing a visual breakdown of the estimated charges and per-Mbps costs that will make up the total costs associated with this cost group for the current billing period (see **Cost Group Model**).
* **Cost Group Traffic History**: A line chart with time (in months) as the horizontal axis and two vertical axes for traffic volume: ingress at left and egress at right.
* **Cost Group Interfaces**: A table listing the interfaces, grouped by device, that make up this cost group (see **Cost Group Interfaces**).

#### Cost Group Interfaces

The **Cost Group Interfaces** table lists the interfaces that make up this cost group, grouped by device (with a badge indicating the number of interfaces for each device). The table includes the following columns that provide details related to the traffic on the interfaces:

* **Interface**: The name and description of the interface.
* **Ingress** and **Egress**: Traffic volume (bps) for this interface in this direction.

  > **Note:** For the direction with the greater volume the bps will be presented as a badge.
* **Est. Cost**: An estimate of the costs related to this interface for the current billing cycle.
* **Utilization**: Traffic volume on the interface as a percent of the interface's capacity.

## Manage Providers Page

The Manage Providers page, reached via the **Manage Providers** button on the Connectivity Costs landing page, is covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(359).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)

*Providers are managed in the Connectivity Costs module on the Manage Providers page.*

### Manage Providers UI

The **Manage Providers** page includes the following main UI elements:

* **Configure Currency**: A button that opens the **Configure Currency Dialog**, where you can choose the currency in which the costs for providers will be stated.
* **Add Provider**: A button that takes you to a **New Provider** page, where you can configure the new provider (see **Provider Configuration**).
* **Providers list**: A list of cards (see **Provider Card UI**) that each correspond to a provider and show the provider's current settings. The list is tabbed, with each tab showing providers in one of the following connectivity types (see **Understanding Connectivity Types**)

  + Cloud Interconnect
  + Free Private Peering
  + IX
  + Paid Private Peering
  + Transit
* **Suggested Providers** (on subnav): A button that opens the **Suggested Providers** pane, which lists any providers that are not yet included in any cost group but which your organization connects to on interfaces whose network boundary is classified as external.

### Configure Currency Dialog

When multiple currencies are used across different cost groups, the landing page and all pages that involve multiple cost groups in different currencies will state costs using the default currency configured in the Configure Currency dialog. The dialog includes the following UI elements:

* ![Configure company currency settings, selecting United States Dollar as default option.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CC-configure-currency.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)**Close**: Click the **X** in the upper right corner to exit the dialog and keep the currency the same as it was when the dialog was opened.
* **Select currency**: A dropdown from which you can select the default currency in which costs will be stated (applies across your organization).
* **Cancel**: A button that closes the dialog without saving changes.
* **Save**: A button that saves the current currency setting and exits the dialog.

#### Currency Control

The default currency setting made with the **Configure Currency Dialog** does not necessarily determine the currency used in costing at every level. Settings at the cost group and interface levels enable you to set currencies that are different from the default and different from one another. For example, you can express interface charges in local currencies while expressing data-center cross-connects in another currency.

> **Note:** Exchange rates for every currency are fetched and updated daily as part of cost computation. Because these rates are built into each day’s computation, historical cost figures always reflect the exchange rates in use at the time of computation rather than current exchange rates

### Provider Card UI

The cards in the **Providers** list are categorized into tabs by connectivity type. Each card has the following UI elements:

* **Provider name**: The name of the provider as specified on a **Provider Configuration** page.
* **Edit**: A button that opens the configuration page for this provider (see **Provider Configuration**).
* **View Costs**: A button that takes you to the **Cost Details Page** for this cost group.
* **Billing cycle**: The date each month when the invoice for this provider is due. A Due this week lozenge appears when the invoice date is seven or less days ahead.
* **Cost groups**: A list of this provider's cost groups. Each item in the list includes the following information about a cost group:

  + Name: Click to open the configuration page for this cost group.
  + Cost model: The cost model of this cost group (e.g., commit, flat rate, see **About Cost Groups**).
  + Currency: The currency in which costs for the cost group are stated.
* **Add Cost Group**: A button that takes you to the Configure Cost Group page (see **Cost Group Configuration**).

### Suggested Providers

Suggested providers, identified via **Provider Classification**, are providers with whom your organization connects via external interfaces. The interfaces associated with each listed provider meet the following criteria:

* Their connectivity type matches the connectivity type currently selected in the controls at the top of the **Providers List** (see **Manage Providers UI**).
* They haven't yet been assigned to a cost group.

The**Suggested Providers** pane is a collapsible sidebar at the right of the **Manage Provider****s** page. The pane appears in a drawer if the width of the browser is less than 1600px. The providers are listed as a set of cards and each card includes the following UI elements:

* **Name**: The name of the suggested provider.
* **Connectivity Type**: The type of connection your organization has with this suggested provider (e.g. transit, paid peering, etc.; see **Understanding Connectivity Types**).
* **Interfaces**: The number of your organization's interfaces that connect with this suggested provider.
* **Configure New Provider**: A button that opens a configuration page with the fields populated with information for this suggested provider (see **Provider Configuration**).

The list of suggested providers makes it easy to see interfaces that aren't yet assigned to a cost group and to assign them to either a new or existing group.

## Provider Configuration

The provider configuration pages are covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(360).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)

*The configuration page for an individual provider.*

### Provider Configuration Access

The Connectivity Types module includes two page types for the configuration of providers:

* **Configure Provider page**: Reached via the **Edit** button on a provider card on the **Manage Providers Page**.
* **New Provider page**: Reached via the **Add Provider** button on the Connectivity Types landing page.

The UI for these two pages is effectively identical (see **Provider Configuration UI**), except that the **Cost Groups** list isn't shown on the New Provider page until the Details pane has been successfully saved.

### Provider Configuration UI

The provider configuration pages include the following main sections:

* **Remove Provider**: A button that opens a confirmation dialog that allows you to remove the provider from your organization.
* **Provider Settings**: A pane used to specify general information about a provider (see **Provider Settings Pane**).
* **Cost Groups**: A list made up of cards (see **Cost Groups Pane**) that each correspond to a cost group and show the settings made on the **Group Configuration Page**.

### Provider Settings Pane

The**Provider Settings** pane on a provider configuration page includes the following settings and controls:

* **Name**: The name of the provider.

  > **Note:** To ensure an exact match when filtering, the name you enter for a provider must be the same as the provider value determined by interface classification (see **About Provider Classification**).
* **Type**: A drop-down from which you select the connectivity services provided by this provider (e.g., transit, paid peering, see **Understanding Connectivity Types**).
* **Logo**: A drop-down list of providers whose logos Kentik makes available for use in the Connectivity Costs UI. Choose "Auto-detect" if you want Kentik to attempt to match the logo automatically based on the string in the Name field. If a logo is available based on the current selection, it will be displayed in the field above the drop-down. If not, the field will say "No Logo Match."
* **Save**: A button that saves any changes you've made to the settings in the pane.

### Cost Groups Pane

The**Cost Groups** pane is populated with cards that each correspond to a cost group and have the following UI elements:

* **Name**: The name of the cost group.
* **Interfaces**: A lozenge showing the number of interfaces in the cost group.
* **Currency**: A lozenge showing the currency in which costs are stated.
* **Edit**: A button that takes you to the **Cost Group Configuration** page, where you can specify the group's details and interfaces.
* **Cost details overview**: A table showing key information about the cost group from the Settings pane of the cost group configuration page (see **Group Settings Pane**). The information will vary depending on the cost model (see **Commit Model Settings** or **Flat Rate Settings**).
* **Minimum Monthly Spend**: A lozenge showing Kentik's projection (based on provided information) of the minimum that your organization will spend on this cost group for the current billing cycle.
* **Cost Overview**: A pane containing a visual breakdown of the estimated charges and per-Mbps costs that will make up the total costs associated with this cost group for the current billing period (see **Cost Group Model**).

The**Cost Groups** list also has an **Add Cost Group** button at upper right that takes you to a **Cost Group Configuration** page where you can specify the new group.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(362).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)

*Each Cost Group card shows the settings of a cost group associated with this provider.*

## Cost Group Configuration

The configuration of cost groups is covered in the following topics.

### About Cost Groups

A cost group is a set of interfaces that share common values for a set of cost-related properties. Grouping them together allows you to more easily see how the traffic into and out of your network over these interfaces is contributing to your costs. The primary influencers of these costs are the following:

* **Fixed charges**: Charges paid periodically regardless of the volume of traffic, including per-interface fees, such as for cross connects or for transmission backhaul rent, as well as any global fees for services such as always-on DDoS protection.
* **Traffic-based charges**: Charges that are determined by the volume of traffic. The way these charges are calculated is determined by a **Cost Group Model**, in which a cost group formula is populated with specific rates and then applied to projected traffic volume to arrive at a projected cost for the billing cycle.

> **Note:** The Connectivity Cost module doesn't currently account for setup fees, which are assumed to be  one-time expenses that are not a factor in your ongoing operational expense.

### Cost Group Model

The characteristics of a cost group depend on the group's model, which is determined by the settings on the cost group's configuration page. The factors that go into the cost model depend on the model's cost formula:

* **Commit (Blended)**: Your organization commits to a minimum fee for volume up to a specified bandwidth, whether or not you use the allocated bandwidth. The rate on which this fee is based is also applied to traffic volume in excess of the commit bandwidth but below any tiers that have been defined. Any traffic volume above the lowest tier threshold is charged based on the rate defined for each tier.
* **Commit (Volume)**: Your organization pays the same rate for all traffic, which is based on the rate for the top tier into which your traffic volume falls.
* **Flat rate**: Your organization commits to a minimum fee for volume up to a specified bandwidth, whether or not you use the allocated bandwidth. No provision is made for additional traffic.

The following table shows, for each of the above cost formulae, how cost group charges are summed to reach the amount that Kentik states as the projected cost for the current billing cycle.

| Cost formula | Charges (fixed) | Volume <= Bandwidth setting | Added volume < tier 1 threshold | Added volume > tier 1, < tier 2 | Added volume > tier 2, < tier 3 |
| --- | --- | --- | --- | --- | --- |
| Commit (Blended) | Charges amount | + Commit volume x commit rate | + Added volume x commit rate | + Added volume x tier 1 rate | + Added volume x tier 2 rate |
| Commit (Volume) | Charges amount | + Commit volume x rate of top used tier | + Added volume x rate of top used tier | + Added volume x rate of top used tier | + Added volume x rate of top used tier |
| Flat rate | Charges amount | + Flat fee | N.A. (volume not included in flat fee is not allowed) | N.A. | N.A. |

For Kentik to accurately estimate your costs for a group, you must choose the cost formula setting (see **Group Settings Pane**) that reflects your arrangement with the provider to which the interfaces in the cost group are connected.

> **Notes:**
>
> * For commit models, volumes above the commit volume are calculated based on the Metered Percentile setting (see **Commit Model Settings**).
> * The number of tiers shown above is for example only; 0 to *n* tiers are supported.

### Cost Model Examples

The examples in the topics below illustrate the application of the cost models described in **Cost Group Model**.

#### Commit (Blended) Example

The screenshot below shows an example of the settings on the **Cost Model Tab** for a cost model using the Commit (Blended) cost formula. In addition to the fixed monthly charges shown at left, the overview shows a commitment to pay for 5 Gbps (whether it is used or not) at a rate of $0.80/Mbps. The same rate holds for all traffic up to the lower bound of Tier 1 at 10 Gbps. Traffic between 10 Gbps and 12 Gbps — the lower bound of Tier 2 — would be charged at a discounted rate of $0.70/Mbps. All traffic above the lower bound of Tier 2 would be charged at a further discounted rate of $0.60/Mbps.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(365).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)

#### Commit (Volume) Example

The screenshot below shows the effect of changing the cost formula to Commit (Volume) while keeping all other settings the same. With this formula, the rate applied to all traffic is the same, but that rate depends on the volume used. Your organization is still committed to pay for 5 Gbps at $0.80/Mbps. If the volume is above this committed bandwidth but below the lower bound of Tier 1 (10 Gbps) all traffic will be charged at $0.80/Mbps. If the volume is above the lower bound of Tier 1 but below the lower bound of Tier 2 (12 Gbps), all traffic will be billed at $0.70/Mbps. And if the volume is above the lower bound of Tier 2 all traffic will be billed at $0.60/Mbps.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(366).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)

#### Flat Rate Example

The screenshot below shows the effect of changing the cost formula to Flat Rate while keeping all other settings the same. Your organization is committed to paying the flat fee for the specified bandwidth and has no arrangement with the provider to handle excess volume.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(367).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)

### Group Configuration Access

The Connectivity Types module includes two page types for the configuration of cost groups:

* **Configure Cost Group page**: Reached via the **Edit** button on a cost group card in the **Cost Groups Pane** on a Manage Providers page (see **Provider Configuration**).
* **New Cost Group page**: Reached via the **Add Cost Group** button at the top of the **Cost Groups Pane**.

The UI for these two pages is effectively identical (see **Group Configuration Page**).

### Group Configuration Page

The cost group configuration pages include the following main sections:

* **Remove Cost Group**: A button that opens a confirmation dialog that allows you to remove the cost group from your organization.
* **Cost Group Settings**: A pane used to specify general information about a cost group (see **Group Settings Pane**).
* **Cost Model**: A tab specifying the various charges that result your organization's costs for this cost group (see **Cost Model Tab**).

  > **Note:** Per-interface charges are specified individually via the **Interfaces** tab.
* **Interfaces**: A tab listing, by site and device, the interfaces in this cost group (see **Group Interfaces Tab**).

### Group Settings Pane

The**Cost Group Settings** pane is used to specify the properties of a cost group. Some properties are common to all cost groups while others depend on the group's cost model.

#### Common Cost Settings

The following cost-group details are displayed for all cost group models:

* **Name** (required): The name that you give to the cost group.
* **Currency** (required): The currency in which to state costs.
* **Cost Formula** (see **About Cost Groups**):

  + Commit (Blended): Monthly cost for traffic that falls into each tier is calculated at that tier's rate.
  + Commit (Volume): Monthly cost for traffic in all tiers is calculated at the rate of the highest used tier.
  + Flat Rate: Monthly cost is a flat fee that covers traffic volume only up to the stated bandwidth.
* **Billing Cycle Start Date**: The day of the month on which your billing cycle for this cost group starts.
* **Contract End Date**: The date on which your current contract with this cost group is up for renewal.
* **Bandwidth**: The base traffic volume, included in either your commit or your flat fee, for which you pay the provider regardless of actual usage. This bandwidth is based on the larger of ingress traffic or egress traffic that you expect to send/receive over the interface(s) in the cost group.  
  **Example**: If you expect egress traffic of 10 Gbps and ingress traffic of 40 Gbps, your committed information rate would be based on 40 Gbps.

#### Commit Model Settings

The following fields are present only when Cost Formula is set to Commit (Blended) or Commit (Volume):

* **Metered Percentile**:The metered percentile enables a certain amount of traffic (traffic in excess of the stated percentile) to briefly "burst" past the committed information rate without triggering a financial penalty from the provider. The percentile may be set from 90th to 99th depending on your contract with the provider.  
  **Example***:* Metering at 95th percentile means that the top 5% of all samples considered within a given time period will be discarded, and the max of in and out will be used to compute the cost.
* **Computation Method**: A set of radio buttons that determine the method used to compute the 9xth percentile set with the **Metered Percentile** selector (see **Percentile Computation Methods**).
* **Unit Price per Mbps** (required): The price, negotiated with the provider, by which traffic volume is multiplied to arrive at cost (exclusive of charges). The traffic volume is the greater of the actual utilization of the circuit (after accounting for the metered percentile) or the committed information rate.

#### Percentile Computation Methods

The following examples illustrate the calculations involved for each of the **Computation Method** options:

* **Peak of Sums (single interface sets)**: A less common method for CIR contracts.

  + For each sample interval, separately sum the IN sample and the OUT samples of all interfaces in the set.
  + Separately compute the p9Xth across the time range for IN and OUT.
  + Take the greater of the IN or OUT p9Xth.
* **Peak of Sums (multiple interface sets)**: A rarely used method that addresses interfaces in multiple time zones by allowing one Commitment volume to be shared over separate p9Xth computations for each interface set.

  + Separately compute the Peak of Sums for each interface set.
  + Sum the Peak of Sums results from all of the interface sets.
* **Peak of Sums (max in/out)**: A less common method for CIR contracts (used primarily by TATA Communications, AS6453).

  + For each sample interval, separately sum the IN samples and the OUT samples of all interfaces in the set.
  + Take the greater of the IN or OUT sum for each sample interval.
  + Compute the p9Xth over the time range for the resulting set.
* **Sum of Peaks**: The most common method for CIR contracts.

  + For each interface in the set, separately compute the p9Xth of all samples in the time range for IN and for OUT.
  + For each interface in the set, take the max of the IN p9Xth and the OUT p9Xth.
  + Sum the maxes of all interfaces in the set.

#### Flat Rate Settings

The following field is present only when **Cost Formula** is set to Flat Rate:

* **Flat Rate Charge** (required): The amount your organization will pay for this cost group for this billing period.

#### Saving a Cost Group

The **Save** button in the **Cost Group Settings** pane will be inactive until the required settings, which depend on the cost formula, are specified. Click the **Save** button to save the cost group and return to the provider configuration page, or click **Cancel** to leave the cost group configuration page without saving the current settings.

> **Notes:**
>
> * The **Save** button will be inactive if any cost tier or global charge is only partially specified.
> * If desired a cost group may be saved without making any settings on the**Cost Model** tab (cost tiers and global charges) or the **Interfaces** tab.

### Cost Model Tab

The**Cost Model** tab provides a breakdown of the component parts that add up to your projected spend for this cost group, and it enables you to enter global charges and tiers that factor into the cost of this cost group. The Cost group tab includes the following main elements:

* **Cost Overview**: A pane containing a visual breakdown of the estimated charges and per-Mbps costs that will make up the total costs associated with this cost group for the current billing period (see **Cost Group Model**).
* **Cost Tiers** (not present when Cost Formula is Flat Rate): A pane containing a form for adding one or more tiers to the cost model for this cost group (see **Cost Tiers Pane**).
* **Global Charges**: A pane for entering any fixed periodic charges incurred overall by this cost group (see **Global Charges Pane**) rather than associated with individual interfaces in the group.
* **Minimum Monthly Spend**: A breakdown of the minimum cost you can expect to pay per billing cycle for this cost group.

> **Note:** For examples of the**Cost Model** tab with different **Cost Formula**settings, see **Cost Model Examples**.

#### Cost Tiers Pane

The **Cost Tiers** pane contains a form that enables you to specify tiers for this cost model (see **Cost Group Model**). In its initial state the pane contains only an**Add Tier** button. When the button is clicked the pane expands to show a row including the following fields:

* **Cost Tier Name**: Any name you choose for the tier. If no name is entered the name will default to "Tier 1," "Tier 2," etc.
* **Traffic volume number**: A number for the lower bound of the traffic volume of the tier (e.g. "10" for a tier starting at 10 Gbps).
* **Traffic volume unit**: A drop-down from which you can select the unit of the entered number, either Mbps or Gbps.
* **Rate**: The amount per Mbps that your organization is charged for traffic in this tier over the interfaces in this cost group.
* **Remove** (trash icon): A button that opens a confirmation dialog that allows you to remove the tier from this cost group.

When the fields above are specified, the Cost Overview and Minimum Monthly Spend panes will update to reflect the added tier. Click the**Add Tier** button to add another tier.

#### Global Charges Pane

The **Global Charges**pane enables you to enter any fixed periodic charges incurred overall by this cost group regardless of the volume of traffic, such as for services like always-on DDoS protection, but not including per-interface fees, which would be entered for an individual interface using the **Interface Charges Drawer**. In its initial state the pane contains only an**Add Charge** button. When the button is clicked the pane expands to show a row including the following fields:

* **Charge Name**: Any name you choose to identify the charge.
* **Amount**: The amount of the charge in the shown currency.
* **Select currency**: A dropdown from which you can select the currency in which costs will be stated.
* **Frequency**: The interval at which the provider bills your organization for this charge.
* **Remove** (trash icon): A button that opens a confirmation dialog that allows you to remove the charge from this cost group.

When the fields above are specified, the **Cost Overview** and **Minimum Monthly Spend** panes will update to reflect the added charge. Click the **Add Charge** button to add another charge.

### Group Interfaces Tab

The Interfaces tab contains the following main elements:

* **Add/Edit Interfaces**: A button that opens the **Configure Cost Group** dialog.
* **Interfaces tables**: A set of tables, one for each site that has interfaces in the cost group (see **Group Interfaces Tables**).

#### Group Interfaces Tables

Each site that has interfaces in the cost group is represented as a table that lists those interfaces, grouped by device. The table includes the following columns and controls:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(371).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)**Interface**: The name and description of the interface.
* **Utilization**: Traffic volume on the interface as a percent of the interface's capacity.
* **Connectivity Type**: The connectivity type of the interface as identified by interface classification (see **Interface Classification Dimensions** and **Understanding Connectivity Types**).
* **Provider**: The provider, as identified by interface classification, that provides network connectivity services to your organization (see **About Provider Classification**).

  > **Note:** To ensure an exact match when filtering, the name you enter when configuring a provider for connectivity costs (see **Provider Settings Pane**) must be the same as the provider value determined by provider classification.
* **Monthly Cost**: The projected cost for this interface in this billing cycle.
* **Add/Edit Interface Charges**: A link that slides out the **Interface Charges Drawer** for this interface.
* **Remove** (trash icon): A link that pops up a confirmation dialog allowing you to remove this interface from the cost group.

### Interface Charges Drawer

The **Interface Charges** drawer contains a form that enables you to add costs that are specific to the individual interface that you clicked on to open the drawer. For example, your provider may charge for a cross-connect from the interface on your equipment to a meet-me room where you connect to a connectivity provider’s interface.  
  
The charges added via this drawer for all of the interfaces in the cost group are summed and contribute to the fixed charges shown in the Cost Overview shown on the **Cost Model Tab**, as well as to the cost/Mbps of any cost group whose interfaces have been assigned additional costs.

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(372).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A17Z&se=2026-05-12T10%3A07%3A17Z&sr=c&sp=r&sig=YJ2QArNd%2FyvKa9SSp5knxIif2%2BxAyYdacJgT1i0SczM%3D)**Description**: A string that will help you recall what the charge is for.
* **Amount**: The amount of the charge.
* **Select currency**: A dropdown from which you can select the currency in which costs will be stated.
* **Frequency**: A dropdown from which you can select the interval (monthly or annual) at which the provider bills your organization for this charge. If annual, the amount is divided by 12 before use in monthly computations.

When the fields are filled, the **Add Charge** button will become active. When the button is clicked, the charge will be added to the **Interface Charges Table**.

#### Interface Charges Table

The **Interface Charges**table lists all of the charges that are specific to a given interface. The columns of the table are the same as the fields listed in **Interface Charges Drawer**. The following controls will be present at the right of each row of the table:

* **Edit** (pencil icon): A button that opens the Edit Interface Charge dialog, which allows you to update or delete the charge from the list.
* **Remove** (trash icon): A button that opens a confirmation popup allowing you to remove a charge from the list.

---

© 2014-25 Kentik
