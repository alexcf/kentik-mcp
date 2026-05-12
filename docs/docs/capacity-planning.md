# Capacity Planning

Source: https://kb.kentik.com/docs/capacity-planning

---

This article discusses the Capacity Planning workflow in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(707).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A03Z&se=2026-05-12T09%3A56%3A03Z&sr=c&sp=r&sig=cHbPcD6rTKyYOeyi8EAoyK8PJKTqLvJ8VSWe0A5S1zM%3D)

*The Capacity Planning page tells you the utilization and projected runout of capacity for groups of interfaces.*

## About Capacity Planning

|  |  |
| --- | --- |
| **Purpose:** | Automate network capacity planning tasks and prioritize actions using growth forecasts and projected run-out dates. |
| **Benefits:** | * Automate data gathering and correlation to prioritize the most critical capacity issues. * Eliminate complex spreadsheets and manual planning processes. * Prevent congestion surprises that impact applications or users. |
| **Use Cases:** | Capacity planning |
| **Relevant Roles:** | Network Architect, Procurement Team |

Capacity planning for a network involves assessing whether and when the volume of traffic on one or more interfaces is likely to exceed the available bandwidth of those interfaces. This assessment can be expressed in terms of two main metrics:

* **Utilization**: The bitrate of traffic volume as a percent of capacity.
* **Runout**: The time remaining before traffic, at current utilization trends, will exceed bandwidth.

Capacity planning has traditionally been a manual process involving complex spreadsheets that are used to collect data, aggregate it, run statistics on it, and trend it over time. In Kentik, the Capacity Planning workflow (Core » **Capacity Planning**) automates this process, enabling you to quickly assess the utilization of links relative to link capacity so that you can take action before links become overloaded.  
  
Kentik's capacity planning tools show you not only current traffic volume, but also utilization trends and how soon your links are likely to reach runout. You can create multiple "capacity plans" and assign different groups of interfaces to each plan. The properties of a plan, such as utilization and runout thresholds, are used to evaluate the status of the plan's interfaces. These capabilities enable service providers and other network operators to optimize the use of their infrastructure, delivering the best possible service while avoiding overspend on capacity.

### Capacity Planning Use Cases

The following table lists some use cases of capacity planning for various roles within a network organization.

| Role | Scenarios |
| --- | --- |
| Network Architect | * Determine when to purchase or design upgraded bandwidth. * View capacity to individual network providers to determine when to make network upgrades or decommission underutilized capacity for a given provider. * View capacity to groups of customer interfaces to determine when to recommend network upgrades to a given customer. |
| Procurement Team | * View capacity to individual network providers to determine when to make network upgrades or decommission underutilized capacity for a given provider. |

### Capacity Status

The status of a capacity plan and its interfaces reflects the settings for the plan's **Capacity Plan Thresholds**:

* **Interface status**: The status displayed for a given interface is the metric — utilization or runout — that is currently most severe. For example, if the interface's utilization status is Warning and its runout status is Critical then the overall status of the interface will be indicated as Critical.
* **Capacity plan status**: The status displayed for a given capacity plan is the status of the most severe interface in the interface groups assigned to the plan. For example, if a plan's groups include 10 interfaces whose status is Normal, five whose status is Warning, and one whose status is Critical, the overall status of the plan will be indicated as Critical.

## Capacity Planning Page

The main elements of the Capacity Planning page include the following:

* **Share** (on subnav): Opens the Share dialog to enable you to share the current view (see **Sharing via the Share Dialog**).
* **Actions** (on subnav): A drop-down menu with actions you can take involving the current view (see **Capacity Planning Actions**).
* **Insights** (on subnav): A button that pops up the **Insights Pane**, which shows insights related to traffic across your network (see **About Insights**).
* **Status filter controls**: Controls that enable you to determine what's currently shown in the **Capacity Plans** list (see **Status Filter Controls**).
* **Plan sort controls**: Controls that enable you to determine what's currently shown in the **Capacity Plans** list (see **Plan Sort Controls**).
* **Configure New Plan**: A button that opens the configuration page for a new capacity plan (see **Configure Capacity Plan**).
* **Capacity Plans list**: The main display area of the page contains a set of cards that each provide information about one of your organization’s existing capacity plans (see **Capacity Plans List**).

#### Capacity Planning Actions

The following options are available from the Action menu:

* **Export**: Prepares a visual report (PDF) or data table (CSV) to export as described in **Portal Export Options**. A notification appears when the report is ready to download.
* **Subscribe**: Opens the Subscribe dialog enabling you to create a subscription for this capacity planning view. The UI of the Subscribe dialog is the same as that of the **Subscription** tab on the Share dialog, which is covered in **Subscription Tab UI**.
* **Unsubscribe** (only if you’re already subscribed to this capacity planning view): Opens a dialog enabling you to confirm that you wish to be removed from the subscription.

#### Status Filter Controls

The **Capacity Plans** list can be filtered to show only plans whose health status (see **Capacity Status**) matches the statuses selected with this control. By default a lozenge representing each of the three statuses (Healthy, Warning, and Critical) is displayed in the control, meaning that plans of all statuses will be included in the **Capacity Plans** list.

To filter the list:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(465).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A03Z&se=2026-05-12T09%3A56%3A03Z&sr=c&sp=r&sig=cHbPcD6rTKyYOeyi8EAoyK8PJKTqLvJ8VSWe0A5S1zM%3D)**Remove plans of a given status**: Click the **X** at the right of the lozenge representing that status.
* **Add plans of a given status**: Click within the control to open a popup from which you can select a status to add.

#### Plan Sort Controls

The **Capacity Plans** list can be sorted to determine the order in which plans are displayed. Click the **Sort by** drop-down to choose one of the following sort orders:

* **Utilization**: Sorts the cards from highest to lowest utilization.
* **Runout**: Sorts the cards from earliest to latest runout.
* **Name**: Sorts the cards alphabetically by name.

### Capacity Plans List

Your organization’s existing capacity plans are listed as a set of cards in the **Capacity** **Plans** list. The width of the browser viewport determines whether the cards are arrayed one, two, or three across.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(466).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A03Z&se=2026-05-12T09%3A56%3A03Z&sr=c&sp=r&sig=cHbPcD6rTKyYOeyi8EAoyK8PJKTqLvJ8VSWe0A5S1zM%3D)

*Each card in the Capacity Plans list provides details on the utilization and runout of a capacity plan.*

Each card in the list contains the following UI elements:

* **Overall Status**: An icon indicating the overall status of the capacity plan (see "Capacity plan status" in **Capacity Status**), which is the status of the interface with the worst status (either runout or utilization) in the plan.  
  **Example*:*** A single interface in the group has a runout status of Critical. The entire capacity plan will be indicated as Critical.
* **Name**: The name of this capacity plan.
* **Action**: A menu that drops down from the vertical ellipsis icon at the top right of the card and offers a set of possible actions for this capacity plan (see **Capacity Plan Actions**).
* **Description**: Optional description text for this capacity plan.
* **Status bar**: A horizontal bar composed of segments that each represent one of the interfaces assigned to this capacity plan.

  + The color of each segment is determined by the corresponding interface's status (see "Interface status" in **Capacity Status**).
  + The width of each segment is determined by the capacity of the corresponding interface as a proportion of the capacity of all of the interfaces in the plan.  
    **Example*:*** In a plan with 20 interfaces whose total capacity adds up to 100 Gbps, an interface whose capacity is 1 Gbps has a status of Critical and two interfaces whose combined capacity is 10 Gbps have a status of Warning. One percent of the status bar will be red, 10 percent will be orange, and the remaining 89 percent will be green.
* **Interface Count**: The total number of interfaces in all of the capacity groups of this capacity plan.
* **Total Capacity**: The sum of the capacities of all interfaces in the group.
* **Runout**: The runout of the interface whose runout is soonest.
* **Utilization**: The utilization of the interface with the highest utilization.

### Capacity Plan Actions

The drop-down Action menu (vertical dots icon) on each capacity plan's card includes the following actions for that plan:

* **Edit Plan**: Opens the configuration page for the plan (see **Configure Capacity Plan**).
* **Export**: Exports one of the following:

  + Visual Report: A PDF of the current state of the plan. Kentik will prepare the report and display a notification indicating that the report is ready. When you click the provided link your browser will handle the download to your local machine.
  + Data Table: Information about the capacity plan as a CSV file. Open the file in a spreadsheet or other program for further analysis.
* **Subscribe**: Opens the Subscribe dialog enabling you to create a subscription for this capacity planning view. The form in the Subscribe dialog is the same as on the **Subscription** tab of the Share dialog (see **Subscription Tab UI)**.
* **Unsubscribe**: If you’re subscribed to this capacity planning view, this option opens a dialog enabling you to confirm that you wish to be removed from the subscription.
* **Share:** Opens the Share dialog which allows you to share via link, email, or subscription (see **Sharing via the Share Dialog**).
* **Remove**: Opens a confirmation dialog that enables you to remove this plan from your organization's collection of capacity plans.

## Capacity Plan Page

A Capacity Plan page provides access to information and settings for an individual capacity plan.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(467).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A03Z&se=2026-05-12T09%3A56%3A03Z&sr=c&sp=r&sig=cHbPcD6rTKyYOeyi8EAoyK8PJKTqLvJ8VSWe0A5S1zM%3D)

*Individual capacity plans are managed on a Capacity Plan page.*

### Capacity Plan Page UI

The **Capacity Plan** page includes the following main UI elements:

* **Plan name**: The name of this capacity plan.
* **Plan status**: An indicator, just to the right of the name, stating the current status of the plan (see **Capacity Status**).
* **Description**: A description of this capacity plan.
* **Edit Capacity Plan**: A button that opens the **Configure Plan** page for this plan (see **Configure Capacity Plan**).
* **Remove Capacity Plan**: A button that opens a confirming popup that enables you to remove the plan from your organization’s collection of plans.
* **Capacity plan indicators**: A pane containing four panels of indicators for this plan (see **Capacity Plan Indicators**).
* **Interfaces list filters**: A set of filters that can be applied to narrow the list of interfaces (see **Interfaces List Filters**).
* **Interfaces list**: A list of the interfaces assigned to this capacity plan (see **Plan Interfaces List**).

#### Capacity Plan Indicators

The indicator panels contain the following counts, status, and configuration information for this plan:

* **Total Capacity**:

  + Total: The sum of the capacities of all interfaces in this capacity plan.
  + % Utilized: Actual traffic volume on these interfaces as a percentage of their total capacity.
* **Interfaces**: Shows you the total number of interfaces and the number of interfaces whose **Capacity Status** is Critical, Warning, and Healthy.
* **Runout**: A summary of the runout status of the interfaces in the plan:

  + Runout settings (question mark): An icon that you can hover over to open a popup stating the calculation strategy (either Week-over-Week or Month-over-Month) and threshold settings used for Critical (red dot) and Warning (orange dot) in this plan (see **Capacity Plan Thresholds**).
  + Runout status: An icon showing the runout status of this plan, which represents the worst runout status of any interface assigned to the plan.
  + Runout status count: The number of interfaces whose runout status is critical, warning, and healthy.
  + Earliest: The soonest date on which any interface in the plan is predicted to run out of capacity.
* **Utilization**: A summary of the utiization status of the interfaces in the plan:

  + Utilization settings (question mark): An icon that you can hover over to open a popup stating the method (e.g. average, maximum, 95th percentile, 98th percentile) used to evaluate traffic volume and the utilization settings (in percent) for Critical (red dot) and Warning (orange dot) in this plan (see **Capacity Plan Thresholds**).
  + Utilization status: An icon showing the utilization status of this plan, which represents the worst utilization status of any interface assigned to the plan.
  + Utilization status count: The number of interfaces whose utilization status is Critical and Warning.
  + Highest: The highest utilization of any interface in the plan.

### Interfaces List Filters

A set of filtering controls just above the **Interfaces** list can be applied to narrow the list of interfaces and change the way the interfaces are organized. The following filter controls are included:

* **Search field**: Enter text to narrow the **Interfaces** list to rows in which one of the columns contains the entered text.
* **Status filters**: A set of controls that enable you to filter the list to show only interfaces with one or more health statuses (see **Status Filter Controls**).
* **Group By**: A drop-down from which you can choose how to group the interfaces in the list. Each group is preceded with a heading row identifying the group:

  + None: Interfaces are presented, without grouping, in the sort order determined by the currently selected column heading. Site and Device columns are both present in the table.
  + Site: Group the interfaces by their respective sites, with a heading row for each site.
  + Device (Site): Group the interfaces by the device on which they exist, with a heading row for each device that also states the site. Site and Device will no longer be included as columns.
  + Connectivity Type: Group the interfaces by their **Connectivity Type Attribute**, with a heading row for each connectivity type.
  + Provider (Connectivity Type): Group the interfaces by their provider, with a heading row for each provider that also states the connectivity type.

### Plan Interfaces List

The **Plan Interfaces** table lists the interfaces currently assigned to this capacity plan (see **Configure Interface Groups**). When clicked, each row will expand and show an **Interface Utilization Chart**. The table contains the following columns:

* **Sev.**: An indicator stating the severity of the current status of the interface (see **Capacity Status**).
* **Interface**: This column contains the following information for each interface:

  + Interface name: Name as either defined in the device and retrieved via SNMP or as entered manually when adding the interface in Kentik (see **Name** in **Interface Field Definitions**). Click the name to go to that interface’s Details page (see **Core Details Pages**).
  + Interface description: Description as either defined in the device and retrieved via SNMP or as entered manually when adding the interface in Kentik (see **Description** in **Interface Field Definitions**).
  + Utilization meter: A 10-segment meter in which the color of the segments reflects the utilization threshold settings for this plan (see **Capacity Plan Thresholds**) and the percent of actual utilization is represented by the segments that are solid rather than outline.
* **Capacity**: Capacity in Mbps for the interface either as retrieved via SNMP or as entered manually when the interface was added in Settings » **Interfaces**.
* **Site** (not present when the **Group By** setting is Site or Device): The site at which the device containing the interface exists.
* **Device** (not present when the **Group By** setting is Device): The device on which the interface exists.
* **Network Boundary**: The interface's **Network Boundary Attribute**.
* **Connectivity Type**: The interface's **Connectivity Type Attribute**.
* **Provider**: The provider via which source/destination traffic over a given interface reaches the Internet (see **Provider Classification**).
* **Traffic**: The bitrate, inbound (In) and outbound (Out), of the actual traffic volume on the interface (in Mbps) calculated over the period specified with the current Runout Calculation Strategy (last month or last week) and using the current Utilization Aggregate (see **Capacity Plan Thresholds**), which is stated in parentheses in the heading.
* **Utilization**: The actual traffic volume, inbound (In) and outbound (Out), as a percent of the interface capacity. The utilization aggregate used for the calculation is stated in parentheses in the heading.
* **Runout**: The projected runout of capacity on the interface, inbound (In) and outbound (Out). If the projected runout date is more than a year in the future the displayed value is “Over a year.” If it’s too far out to be calculated (e.g. the utilization trend is declining) the value is stated as “N/A.” If capacity has already been exceeded, the displayed value is “Over Capacity.”

  > **Note:** Runout dates are computed using double or triple exponential smoothing.
* **Trend**: The trend in actual traffic, inbound (In) and outbound (Out), over the most recent period (week or month depending on this capacity plan’s Runout Calculation Strategy setting).
* **Action**: A drop-down menu from which you can navigate directly to other areas of the portal where information about the interface is available, which may include:

  + Interface Details: Goes to the Details page for the interface in Network Explorer.
  + Traffic Engineering (only on interfaces whose Network Boundary is External): Goes to the individual page for the interface in Traffic Engineering.
  + Show in Data Explorer: Opens Data Explorer in a new tab, with the **Query** sidebar set to show Ingress or Egress traffic on this interface.

#### Interface Utilization Chart

When you click a row of the **Plan Interfaces** list, that row expands to show a bidirectional line chart representing traffic on the interface in bits/s over the last 30 days. The chart includes the following UI elements.

* **Bits/s In and Out**: Solid lines that plot the interface's traffic volume. Inbound traffic is plotted above the X axis and labeled on the left; outbound traffic is plotted below the X axis and labeled on the right. The lines are color-coded based on threshold settings:

  + Red: The traffic utilization is greater than the Critical threshold.
  + Orange: The traffic level is between the Warning and Critical thresholds.
  + Green: The traffic level is less than the Warning threshold.
  > **Note**: The bitrate calculation is based on the current **Utilization Aggregate** setting (see **Capacity Plan Thresholds**).
* **Thresholds**: Dashed lines that indicate your current Critical (red) and Warning (orange) threshold settings for interface utilization for both inbound and outbound traffic (see **Capacity Plan Thresholds**).
* **Reset Zoom**: A button that restores the chart to its original zoom. The button appears in the top right corner of the chart when you click and drag on the chart zoom into a shorter time range.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(468).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A03Z&se=2026-05-12T09%3A56%3A03Z&sr=c&sp=r&sig=cHbPcD6rTKyYOeyi8EAoyK8PJKTqLvJ8VSWe0A5S1zM%3D)

*The bidirectional utilization chart plots in (above) and out (below) interface traffic.*

## Configure Capacity Plan

The **Configure Plan** page may be accessed from the following locations:

* **Capacity Planning**: On the Capacity Planning landing page, click the **Configure New Plan** button at the upper right.
* **Capacity Plan**: On the Capacity Plan page for an individual plan, click the **Edit Capacity Plan** button.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(469).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A03Z&se=2026-05-12T09%3A56%3A03Z&sr=c&sp=r&sig=cHbPcD6rTKyYOeyi8EAoyK8PJKTqLvJ8VSWe0A5S1zM%3D)

*The Configure Plan page is used to specify the settings for an individual capacity plan.*

### Configure Plan Page UI

The Configure Plan page for a capacity plan includes the following main fields and controls:

* **Plan Name**: The name of the capacity plan.
* **Cancel**: A button that exits the page without saving the changes made to the settings.
* **Remove Capacity Plan** (only when editing existing plan): A button that opens a confirmation popup that enables you to remove the plan from your organization’s collection of plans.
* **Save Capacity Plan**: A button that saves the current settings and exits the page.
* **Capacity Computation**: A tab containing utilization and runout settings for this capacity plan (see **Capacity Computation Settings**).
* **Interfaces**: A tab listing the interface groups assigned to this capacity plan and enabling you to edit the interfaces included in those groups (see **Capacity Plan Interfaces**).

### Capacity Computation Settings

The **Capacity Computation** tab includes the following settings for a capacity plan.

* **Plan Name**: A field in which you can specify a name for the capacity plan.
* **Description**: A field in which you can enter a description of the capacity plan.
* **Thresholds**: Settings for the thresholds used to define the current utilization and runout status (Normal, Warning, Critical) of this plan (see **Capacity Plan Thresholds**).

### Capacity Plan Thresholds

Each capacity plan includes threshold settings for the following:

* **Runout**: The number of months before traffic will exceed capacity if present trends continue.
* **Utilization**: Interface traffic volume relative to capacity as reported by SNMP.

The settings for these thresholds let you define the conditions under which the status of a given interface becomes Warning or Critical. The following controls are provided for these settings:

* **Runout Calculation Strategy**: A drop-down to set the comparison period used when calculating, based on the current **Utilization Aggregate** method, the runout (assuming that current trends continue):

  + Week over Week (WoW): Compare the most recent week (back in time from the present) with the week before that.
  + Month over Month (MoM): Compare the most recent month (back in time from the present) with the month before that.
* **Runout**: A dual-slider control that sets the following:

  + Warning threshold (left slider): The number of weeks or months below which the runout will trigger a Warning state.
  + Critical threshold (right slider): The number of weeks or months below which the runout will trigger a Critical state.
  > **Note**: The unit of measurement is determined by the time period type selected for the **Runout Calculation Strategy**.
* **Utilization Aggregate**: A drop-down to set the aggregation method (average, maximum, 95th percentile, 98th percentile, or 99th percentile) used to measure traffic volume.
* **Utilization**: A dual-slider control that sets the following:

  + Warning threshold (left slider): The percent of capacity above which traffic volume will trigger the Warning state.
  + Critical threshold (right slider): The percent of capacity above which traffic volume will trigger the Critical state.

### Capacity Plan Interfaces

The **Interfaces** tab of a Configure Plan page lists the interface groups assigned to this capacity plan and enables you to edit the interfaces included in those groups.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(470).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A03Z&se=2026-05-12T09%3A56%3A03Z&sr=c&sp=r&sig=cHbPcD6rTKyYOeyi8EAoyK8PJKTqLvJ8VSWe0A5S1zM%3D)

*The Interfaces tab is used to manage the interfaces assigned to a capacity plan.*

#### About Interface Groups

Interfaces are assigned to capacity plans via **interface groups**, which may be either of the following types:

* **Static Group**: Created by assigning individual interfaces directly. Assigned interfaces will remain part of a static group unless they are removed. A maximum of one static group can be defined per capacity plan.
* **Dynamic Groups**: Created by filtering the list of available interfaces so that the only interfaces in the group are those that match the specified criteria. Interfaces are dropped from the group if they no longer meet the group's defined conditions, and other interfaces may be added as they meet those conditions. Multiple dynamic groups may be defined per capacity plan.

#### Interfaces Tab UI

The **Interfaces** tab includes the following UI elements:

* **Add/Edit Interfaces**: A button that opens the Configure Capacity Group dialog (see **Configure Interface Groups**), where you can specify individual interfaces for a static group and filter criteria for dynamic groups.
* **Capacity groups list**: A set of tables that each represent an interface group that is currently defined in the Configure Capacity Group dialog (see **Interface Groups List**).

### Interface Groups List

The **Interface Groups** list is a set of tables that each contain interfaces assigned to this capacity plan. One table lists the interfaces (if any) in the plan's static group (see **About Interface Groups**). The rest of the tables each represent one dynamic group and show all of the interfaces that currently meet the criteria defined for that group in the **Filters** pane of the Configure Capacity Group dialog (see **Configure Interface Groups**). This may include interfaces that didn't exist or didn't meet the criteria when the dynamic group was created.  
  
 Each of the tables in the **Interface Groups** list include the following columns:

* **Interface**: The name (top line) and description of the interface, either as defined in the device and retrieved via SNMP or as entered manually when adding the interface in Kentik (see **Interface Field Definitions**).
* **Capacity**: Capacity in Gbits/s for the interface either as retrieved via SNMP or as entered manually when the interface was added in Settings » **Interfaces**.
* **Network Boundary**: The interface’s **Network Boundary Attribute**, if any.
* **Connectivity Type**: The interface’s **Connectivity Type Attribute**.
* **Provider**: The provider, if any, via which source/destination traffic over a given interface reaches the Internet (see **Provider Classification**).

The tables in the capacity groups list include the following additional features:

* **Table header**: A row across the top of the table (above the column headings) that includes:

  + Count: The number of interfaces in the group.
  + Filters (for dynamic groups only): The filter criteria by which the group is defined.
  + Remove (for dynamic groups only): A button that opens a confirmation popup allowing you to remove the dynamic group from this capacity plan.
* **Heading rows**: Interfaces are grouped by device. The heading row for each device includes:

  + Expand/collapse: An arrow at the left of the row enabling you to show or hide the rows for individual interfaces.
  + Device: The name of the device on which the interfaces exist.
  + Site: The site at which the device exists.
  + Count: The number of interfaces on the device that match the criteria for this dynamic group.

## Configure Interface Groups

The Configure Capacity Group dialog, which opens from the **Add/Edit Interfaces** button on a capacity plan's Configure Plan page, is used to define, either directly or via dynamic filtering, the plan's static and dynamic interface groups.

![Configuration interface for managing dynamic and static capacity groups in network settings.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CP-configure-capacity-group-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A03Z&se=2026-05-12T09%3A56%3A03Z&sr=c&sp=r&sig=cHbPcD6rTKyYOeyi8EAoyK8PJKTqLvJ8VSWe0A5S1zM%3D)

*The Configure Capacity Group dialog is used to assign interfaces to the static and dynamic groups of a capacity plan.*

### Configure Capacity Group UI

The Configure Capacity Group dialog includes the following main UI elements:

* **Filters**: A set of fields, drop-downs, radio buttons, and checkboxes that define filter criteria, which are used to narrow the set of interfaces that are available to add to a static or dynamic interface group (see **Filters Pane**).
* **Group By**: A set of buttons that determine how the interfaces in the **Matching Interfaces** list will be organized into groups, either by device, connectivity type, or network boundary (see **Interface Classification Dimensions**).
* **Create Dynamic Group**: A button, active only when one or more filters are set in the **Filters** pane, that creates a dynamic group made up of all interfaces that match the current filter criteria.

  > **Note:**The composition of a dynamic group will change in response to changes in the interfaces that match the group's filter conditions. For example, if a group is defined as interfaces on a given device whose network boundary is external, then additional external interfaces added on that device will become members of that group.
* **Matching Interfaces**: A list of the individual interfaces that match all criteria currently defined in the **Filters** pane. If no filters are set the list will include the interfaces on all of your Kentik-registered devices. This list serves two purposes:

  + It shows all interfaces that are available to add manually to the **Static Group** at the right of the dialog.
  + If any filters are set, it shows all interfaces that would be added to a new dynamic group if you were to create one with the **Create Dynamic Group** button.
* **Dynamic Groups**: A list of the dynamic groups defined for this capacity plan.
* **Static Group**: A list of the interfaces that have been individually added to this capacity plan's static group.
* **Cancel**: A button — **X** at top right or **Cancel** at bottom right — that exits the dialog, leaving the interface groups for this capacity plan as they were when the dialog was opened.
* **Save**: A button that exits the dialog, updating the interface groups for this capacity plan to the current dynamic groups and static group.

### Filters Pane

The **Filters** pane (on the left) includes a set of filters that you can use to narrow down the interfaces that appear in the **Matching Interfaces** list. This pane includes the following types of controls:

* **Reset to Default**: A button that resets the **Filters** pane to its default setting, removing any user-selected filters.
* **Text fields**: Enter a full or partial text string on which to match interfaces in the field's labeled category (name, description, etc.).
* **Selection fields**: Click to drop down a filterable list of values on which to match interfaces in the field's labeled category (device, label, site, etc.). A selected value will appear as a lozenge in the field. Repeat to add more values. To remove a value, click the **X** at the right of its lozenge.
* **Checkboxes**: Click checkboxes in a given category (see **Filter Categories**) to choose one or more filter criteria in that category. To remove, uncheck the box.
* **Radio buttons**: Click only one of the available options.

> **Note:** The **Filters** pane shows only controls (checkboxes and selectors) for which there is a match with the interfaces shown in the **Matching Interfaces** list. As the list is filtered, the number and type of controls displayed in the **Filters** pane may change.

#### Filter Categories

Filter criteria for the **Matching Interfaces** list fall into the following categories:

* **Interface Name** (text field): Match interfaces whose name contains the entered string. Regular expressions are allowed.
* **Interface Description** (text field): Match interfaces whose description contains the entered text (case insensitive). Regular expressions are allowed.
* **IP** (text field): Match interfaces with the entered IP/CIDR address range.
* **Device** (selection field): Match interfaces to the device selected from the drop-down menu.
* **Device Name** (text field): Match interfaces whose device name contains the entered string. Regular expressions are allowed.
* **Device Label** (selection field): Match interfaces to the device label selected from the drop-down menu.
* **Site** (selection field): Match interfaces to the site selected from drop-down menu
* **Capacity** (selection field): Match interfaces whose capacity matches the selected value.
* **Interface Type** (radio buttons): Select whether interfaces are to match All Interfaces, Logical Interfaces (interfaces that have been assigned an IP address), or Physical Interfaces.
* **Provider** (selection field): Match interfaces whose source/destination traffic reaches the Internet via the selected provider (see **Provider Classification**).
* **Connectivity Type** (checkboxes): Match interfaces whose connectivity type matches the checked box(es); see **Understanding Connectivity Types**.
* **Network Boundary** (checkboxes): Match interfaces whose network boundary matches the checked box(es); see **Network Boundary Attribute**.

You can use the above filter controls to narrow the set of available interfaces. The **Matching Interfaces List** will update to show only interfaces that match the current filters, and the **Create Dynamic Group** button will become active.

#### Filter Application Rules

Kentik applies the following rules to filter categories and criteria:

* Criteria are ORed (match any) within categories and ANDed (match all) between categories.
* When criteria are selected in more than one category, an interface is displayed only if it matches at least one of the selected criteria in all of the categories with criteria selected.
* When no criteria are selected for a category, the interface may match any of the category’s listed criteria.

### Matching Interfaces List

The **Matching Interfaces** list is populated based on the current filters in the **Filters Pane** and grouped according to the **Group By** setting. Interfaces in the list can be manually designated for inclusion in the **Static Group** list by clicking on the **+** icon at the right of each row.  
  
 The list includes the following columns and controls:

* **Interface**: In each row, the top line is the interface’s SNMP alias and the bottom line is the interface’s description (if provided); see **Information SNMP OIDs**.
* **Information**: Indicators showing the interface’s capacity, connectivity type, and network boundary are shown when available (see **Interface Classification Dimensions**).
* **Add** (**+**): An icon at the right of each individual row. Click to add the corresponding interface to the **Static** **Group** list. The icon will change to a circled checkmark.

### Interface Group Lists

The right side of the Configure Capacity Group dialog shows the groups of interfaces that will be assigned to the current capacity plan when the dialog is closed with the **Save** button. Any listed item may be manually removed from the capacity plan with the red **X** at the right of each row.  
  
 Two types of interface groups are represented in this area of the dialog:

* **Dynamic Groups**: Each listed item represents the set of criteria that was specified in the **Filters Pane** when the **Create Dynamic Group** button was clicked (one dynamic group per click of the button). The group is dynamic because all interfaces on your network that currently meet the criteria for a listed dynamic group will be included in the capacity group even if those interfaces didn't exist or didn't meet the criteria when the dynamic group was created.
* **Static Group**: A list of interfaces that have been individually selected from the **Matching Interfaces List** for inclusion in this capacity plan. In each row, the top line is the interface’s SNMP alias and the bottom line is the interface’s description (see **Information SNMP OIDs**).

## Using Capacity Planning

Using capacity planning involves setting up one or more capacity plans and assigning interfaces to the interface groups of each plan. The following steps illustrate how to create a new capacity plan:

1. From the main menu, navigate to the **Capacity Planning Page** (Core » **Capacity Planning**).
2. Click the **Configure New Plan** button at upper right.
3. In the **Capacity Computation** tab of the Configure Plan page (see **Configure Capacity Plan**), enter a name and description for the plan.
4. In the **Thresholds** pane of the tab (see **Capacity Plan Thresholds**), set the methods by which traffic will be compared (**Runout Calculation Strategy**) and aggregated (**Utilization Aggregate**) when it is evaluated for runout and utilization status.
5. Use the **Runout** control to set the number of months below which the runout will trigger a Warning state (left slider) and Critical state (right slider).
6. Use the **Utilization** control to set the percent of capacity above which traffic volume will trigger the Warning state (left slider) and Critical state (right slider).
7. Switch to the page's **Interfaces** tab, then click the **Add/Edit Interfaces** button.
8. In the resulting Configure Capacity Group dialog (see **Configure Interface Groups**), the **Matching Interfaces List** will show all of the interfaces (on Kentik-registered devices) in your network.
9. Use the controls of the **Filters Pane** to narrow the **Matching Interfaces** list.
10. Do one of the following to define an interfaces group for this capacity plan:

    1. Add a dynamic group, based on current **Filters** pane settings, to the **Dynamic Groups** list (see **Interface Group Lists**) by clicking the **Create Dynamic Group** button.
    2. Add an individual interface directly to the **Static Group** list by clicking the **+** on the interface's row in the **Matching Interfaces** list.
11. Repeat step 10 until all of the interfaces that you want in this capacity plan are included in either a dynamic group or the static group, then click the **Save** button to close the dialog.
12. The tables on the **Interfaces** tab of the Configure Plan page will now list the interfaces in the dynamic groups and the static group that you defined in the dialog.
13. Click the **Save Capacity Plan** button to save your new plan. You’ll be returned to the Capacity Planning page, where you can see your new plan as a card in the **Capacity Plans List**.
