# Metrics Explorer

Source: https://kb.kentik.com/docs/metrics-explorer

---

This article discusses the features and use of Metrics Explorer in the Kentik portal.

![Graph showing device utilization over time with highlighted interface names and percentages.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/metrics_explorer-main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)

*A query result displayed in Metrics Explorer*

> **Note:** The following articles provide more on other aspects of Kentik's Network Monitoring System:
>
> * **Network Monitoring**
> * **Kentik NMS Configuration**
> * **NMS Setup**
> * **NMS Dashboard**
> * **NMS Devices**
> * **NMS Interfaces**
> * **Query Assistant**

## About Metrics Explorer

Metrics Explorer enables you to query the SNMP or Streaming Telemetry data gathered from your network by Kentik NMS, and to see the results of those queries presented as visualizations and tables that help you better understand the status and activity of your Kentik-monitored network infrastructure over a user-specified timespan.  
  
Queries are typically configured in the panes of the **Metrics Explorer Query Sidebar**, specifying parameters such as measurements, metrics, dimensions, and time ranges. You can filter the returned data (see **Filters Pane**) based on dimensions and/or the metrics values.  
  
Results typically include the following:

* **Visualization**: A chart (may be one of a various different types; see **Visualization Options Pane**).
* **Table**: A tabular listing of query results (see **Table Options Pane**), including **Action** menu options to inspect site/device/interface traffic, when applicable (see **Inspect Traffic Drawer**).

> **Note:** The tables and graphs returned from queries in Metrics Explorer can also be accessed via API (see **Show API Call**).

## Metrics Explorer Layout

Metrics Explorer is made up of the following main areas:

* **Subnav**: Page-wide controls and navigation in a gray strip across the top of the page (see **Metrics Explorer Subnav**).
* **Query Assistant**: A field for entering natural language queries (see **Metrics Query Assistant**).
* **Main display**: The display area for query results, typically including a visualization and an accompanying table (see **Metrics Explorer Main Display**).
* **Query sidebar**: A drawer at the right of the page, toggled by the **Query** button in the subnav, contains panes to specify query settings (see **Metrics Explorer Query Sidebar**).

## Metrics Explorer Subnav

The Metrics Explorer subnav contains the following page-wide controls:

* **Actions**: A button that opens a dropdown with the options detailed in **Metrics Explorer Actions**.
* **Query**: A button that toggles visibility of the **Metrics Explorer Query Sidebar**, which contains controls for specifying the query results shown in the **Metrics Explorer Main Display**.

#### Metrics Explorer Actions

The **Actions** menu for Metrics Explorer includes the following options:

* **Export**: Opens a menu with **Portal Export Options** to download query results for external viewing or sharing.
* **Add to Observation Deck**: Adds the current visualization as a widget to your **Observation Deck**. To remove it, click the **Action** menu (vertical ellipsis) on the widget and choose **Remove**.
* **Add to NMS Dashboard**: Adds the current visualization as a widget the **NMS Dashboard**. To remove it, click the **Action** menu (vertical ellipsis) and choose **Remove.**
* **Create Saved View**: Opens the **Add Saved View** dialog (see **Saved View Dialogs**), to set properties for a new saved view (see **About Saved Views**) accessible via the **Library**.
* **Add To New Dashboard**: Opens the **Add Dashboard** dialog (see **Dashboard Dialogs**) to create a new dashboard with the current Metrics Explorer view as a panel. Click the **Add Dashboard** button to create and then display the new panel.
* **Add To Existing Dashboard**: Opens the **Add Metrics Explorer View Panel** dialog to add the current view as a panel to an existing dashboard. Click the **Add Panel** button to add it to the dashboard and display the updated dashboard.
* **Preview as Tenant**: Present only for accounts that use My Kentik Portal (MKP), this option opens a submenu to select an MKP tenant, allowing you to view the query as it would appear to that tenant.
* **Show API Call**: Choose an option from the submenu (**For Data** or **JSON input**) to open a dialog with cURL or JSON that you can copy. See **Show API Call**, noting that Metrics Explorer has an additional **For Report** option.
* **Reset to Default Query**: Resets the **Query** sidebar settings to the default Metrics Explorer query configuration.

## Metrics Query Assistant

The Query Assistant feature in Metrics Explorer uses Kentik AI. For more details, including data privacy and security measures, see our **Kentik AI Overview**. For example queries or to learn about the Query Assistant page, see **Query Assistant**.

### Query Assistant in Explorer

The **Query Assistant** field, located below the subnav, allows you to enter queries in natural language rather than using **Query** sidebar controls. Kentik interprets your input to configure the sidebar for a relevant query. While you can type any question into this field, some queries will yield more meaningful results than others.

> **Note:** While our initial Query Assistant implementation is primarily focused on English, our engine also recognizes additional languages such as Spanish and French.

### Query Assistant Popup

Click the **Query Assistant** field to open a popup with the following elements:

* **Query Assistant field**: Use natural language to describe your desired query results, then press **Return/Enter** to convert the text into query settings in the **Query** sidebar.
* **Tips**: Opens a dialog with tips from Kentik on creating queries that get the most relevant and insightful results.
* **Submit**: A button at the right of the **Query Assistant** field that submits the query (or press **Return**/**Enter**).
* **Examples**: Tiles with Kentik-provided text for creating queries. Click a tile to enter the text into the **Query Assistant**field.
* **Recent**: Tiles with text from recent queries. Click a tile to enter the text into the **Query Assistant** field.
* **Clear History**: A button to clear the **Recent** queries list.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(441).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)

*A natural language question entered into the Query Assistant popup becomes an NMS query.*

### Using Query Assistant

To use the Query Assistant:

1. Click in the **Query** field to open the Query Assistant popup.
2. Populate the field by doing one of the following:

   * Enter text directly into the field.
   * Click a tile in the **Example Queries** list.
   * Click a tile in the **Recent Queries** list.

     > **Note**: If you click a tile, skip to step 4.
3. Press **Return/Enter** or click the **Submit** button to run the query and close the popup.
4. If the results don’t match your expectations:

   1. Click the **Query** button in the subnav to open the **Query** sidebar.
   2. Check and adjust the sidebar settings to better reflect your intended query.
   3. Click **Run Query** to see updated results.
5. Provide feedback to help us train and improve our AI by doing the following:

   * Click the thumbs up or thumbs down button to share whether the query results met your expectations.
   * Click the **Provide feedback** button (dialog icon) to open the Provide additional feedback dialog and enter your feedback (see **Query Assistant Feedback**).

> **Note:** To use Query Assistant, Kentik AI must be enabled (see **Manage Kentik AI**).

### Query Assistant Feedback

The **Provide additional feedback** dialog for Kentik Query Assistant is accessible from the feedback button (dialog icon) after a query is returned. The dialog includes the following elements:

* ![User interface feedback form for device utilization and user experience improvement.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/provide additional feedback popup.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)**Query feedback**: Radio buttons (Yes and No) to indicate if the response answered your query (displayed above the buttons) or not.
* **Comment**: An input field for additional feedback to help refine Query Assistant capabilities.
* **Cancel**: Buttons (an **X** in the upper right corner and a **Cancel** button at lower right) to close the dialog without submitting feedback.
* **Save**: A button to submit your feedback to Kentik and close the dialog.

## Metrics Explorer Main Display

The main display area shows the results of the most recently run query and includes the following UI elements:

* **Query Assistant**: A field to enter text for Kentik to interpret as a natural language query, populating the **Query** sidebar and returning results (see **Metrics Query Assistant**).
* **Visualization**: Query results shown as a line chart (default). Other chart types are available in the **Query** sidebar's **Visualization Options Pane**.
* **Resize handle**: A handle, located between the visualization and the **Results** table, that allows you to click and drag up or down to resize the chart or table.
* **Results table**: Lists entities and their metric values as configured in the **Metrics Explorer Query Sidebar**. When applicable, the **Inspect Traffic Drawer** is available from the **Action** menu for detailed traffic data.

### Table Sort Priority

The sort order of the **Results** table is primarily determined by the order of the metrics selected in the **Metrics** field in the **Measurement Pane**, and secondarily by the order of the methods selected in the **Aggregation** field in the **Table Options Pane**. For example, if the **Metrics** field contains (from left) metrics 1, 2, and 3, and the **Aggregation** field contains methods A and B, the priority order is: 1A, 1B, 2A, 2B, 3A, 3B.

## Inspect Traffic Drawer

The **Inspect Traffic** drawer provides a bridge between Metrics Explorer (NMS) and Data Explorer (flow, traffic). When the results of a query in Metrics Explorer contain site, device, or interface dimensions, (if flow is enabled) users can inspect traffic directly in Metrics Explorer. If flow is not enabled for a device, a popup (“This device does not have flow data”) will appear when hovering over the **Action** button.

To access the **Inspect Traffic** drawer, click to display a row’s **Action** menu (far right) and from the menu choose one of the following options:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Screenshot%202025-11-28%20at%204.18.15%E2%80%AFPM.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)

* **Inspect Device Traffic**: View traffic for the selected device
* **Inspect Interface Traffic**: View traffic for a source/destination interface (select Inbound by… or Outbound by…)

After confirming dimension selections in the **Show by Dimensions** dialog (see **Dimension Selectors**), the **Inspect Traffic** drawer opens from the bottom of the screen showing a chart and table of query results from Data Explorer.

![Graph showing device traffic and utilization metrics over a specified time period.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/inspect_traffic_drawer.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)

*The Inspect Traffic drawer displayed for a device*

### Inspect Traffic Drawer UI

The **Inspect Traffic** drawer includes the following UI elements:

* **Query title**: The title bar of the drawer displays the query type and the related site/device/interface name (e.g., “Device Traffic: cedge-01-aws” as shown above).
* **Resize handle**: A handle in the middle of the title bar that allows you to drag the top of the drawer up or down to adjust the height of the drawer.
* **Open in Data Explorer**: A button that opens the query in Data Explorer for further analysis.
* **Minimize/Expand**: A button that collapses or expands the drawer, keeping the data on the drawer available.
* **Close**: A button that closes the drawer to show the full **Metrics Explorer Main Display**.
* **Traffic chart**: A visualization of the traffic for the selected device or interface (see **Inspect Traffic Chart**).
* **Traffic table**: A table displaying the traffic data for the selected device, interface, or inbound/outbound external traffic (see **Inspect Traffic Table**).

> Notes
>
> * The table columns displayed vary based on the query settings (dimensions selected from the **Show by Dimensions** dialog).
> * If a query does not include a device/interface with flow data, upon clicking the **Action** menu you’ll receive a message “This device does not have flow data”.

### Inspect Traffic Chart

A **Stacked Area Chart** displays the traffic data available for the selected device/interface and includes the following UI elements (in addition to the standard stacked area chart features):

* **Filter**: Click to display a popup showing all filters applied to the chart.
* **Analyze**: If Kentik AI is enabled for your organization, clicking this button will display the **Cause Analysis Overlay** for the selected device, showing:

  + Any change detection events in the chart (numbered and displayed in a different color)
  + A second tab (**Cause Analysis tab**) in the **Inspect Traffic Table**. Click **X** on this tab to hide the **Cause Analysis Overlay**.

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/hover info.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)**Hover-over data**: Hover over any point of the graph to see exact traffic values for that moment in time, as well as the data available for each of the dimensions selected from the **Show by Dimensions** dialog.
* **Highlight selection**: Click and drag your mouse to highlight a portion of the visualization. A pop-up will display for the highlighted area allowing you to:

  + **Zoom to Selection**: Zoom in so that the visualization only displays the highlighted timeframe.
  + **Analyze Traffic**: Engage **Cause Analysis** to analyze the highlighted traffic.
  + **Compare Traffic**: **Compare traffic** from two separate time windows.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/selection.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)

### Inspect Traffic Table

Columns include traffic information about the selected device, Average, 95th Percentile, Max, Last Datapoint, and any additional dimensions chosen from the **Show by Dimensions** dialog (selected prior to opening the **Inspect Traffic** drawer). Click a device name (or any other column data) to see that dimension’s Details page in a new tab.

When a user clicks the **Analyze** button, the table will display two tabs: **Left +Y Axis** (containing the previously displayed traffic data), and the **Cause Analysis Tab**(see **Cause Analysis****)****.**

#### Cause Analysis Tab

The **Cause Analysis** tab helps you identify and understand unexpected network events (see **Using Automatic Detection of Traffic Changes****)**. The tab displays the following information for each change detection event (one row) displayed in the **Inspect Traffic Chart**:

* **Expand/Collapse**: Indicates whether the **Event Dimensions** row is shown (**V**) or hidden (**>**) in the row beneath (click anywhere on the event’s row to hide or show its dimensions).
* **Event number**: Corresponds to the event number displayed on the **Inspect Traffic Chart**.
* **Type**: The type of event (e.g. a jump or spike in traffic).
* **Change**: The difference in traffic, shown as an increase (green) or decrease (red)
* **Time**: The date (YYYY-MM-DD) and time (HH:MM) the change detection event was first triggered.
* **Summary**: A brief summary of the change detection event.
* **Details**: A bulleted list of further details to provide an analysis of the event.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Screenshot 2025-12-01 at 4.13.07 PM.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)

#### Event Dimensions

When you click within an event’s row, directly below it, an **Event Dimensions** row displays:

* All of that event’s dimensions data
* The difference in traffic that triggered the event
* The estimated percentage of the total traffic represented by this grouping of dimensions.
* A vertical ellipses button allowing you to view this data in Data Explorer (in a new tab).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Dimension Details row.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)

## Metrics Explorer Query Sidebar

The **Query** sidebar is accessible via the **Query** button in the **Metrics Explorer Subnav**.

> **Note:** In the initial state of the Metrics Explorer page, the **Query** sidebar is open but only the **Measurement** selector is visible. Once you select a measurement the rest of the sidebar's fields and controls are displayed.

Queries in Kentik NMS are formatted to align with the **OpenConfig** data model, ensuring consistent data storage from sources such as SNMP or Streaming Telemetry. They are configured the **Query** sidebar, organized into panes, each focusing on a specific aspect of the query. These panes are covered in the topics below.

### Query Sidebar Controls

Beneath the **Measurement**, **Visualization Options**, **Time**, **Filters**, and **Table Options** panes, the following control appears at the bottom of the **Query** sidebar:

* **Run Query**: Executes a query based on the current settings of the **Query** sidebar and returns results (chart and table) to the main display area.

### Measurement Pane

Use the **Measurement** pane to specify the focus of your query.

#### Measurement Pane Parameters

A measurement can be a physical or logical entity:

* **Physical entities**: Examples include `/system` (OpenConfig's name for a device), `/interfaces/counters`, `/components/fans`, and `/components/power supplies`.
* **Logical entities**: Examples include the BGP process, OSPF, and IS-IS.

Once selected, a measurement determines the additional measurement-related parameters available for a query:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(443).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)**Metrics**: How we’re measuring the traffic:

  + **Metric availability***:* Each measurement offers various metrics. For example, the `/components/memory` measurement includes metrics such as available memory, free memory, and utilization percentage, while the `/interface/counters` measurement includes metrics like `In utilization`, `In errors`, `In packet rate`, and `Admin status`.
  + **Metric presentation**: Queries can return multiple metrics, but only the first metric is plotted in the chart unless otherwise specified in the **Visualization Options Pane**. Additional metrics appear only in the table.
* **Group by Dimensions**: Defines how metrics are grouped by dimensions in visualizations and tables.

  + **Example**: To identify interfaces with the highest inbound utilization, set the measurement to `/interfaces/counters` and the metric to `In utilization`, you'd also set the dimensions to `Device Name` and `Interface Name`. This results in chart plots and table rows representing unique combinations of interface and device.
* **Merge series**: Allows aggregation of dimensions in two steps, first via group-by dimensions, then via a method of your choosing (sum/avg/min/max). This allows you to simplify data and highlight key information.

  + **Example**: Selecting both `Device Name` and `Interface Name` as group-by dimensions results in unique combinations for each plot and table row. If instead, `Interface Name` is used as a Merge Series dimension, then each/plot row represents a device, and the **Merge Series** aggregation setting controls how interface values are combined into a single plot per device.

#### Measurement Pane UI

The **Measurement** pane includes the following UI elements:

* **Measurement selection** (required): A dropdown with a filterable list of available measurements.
* **Metrics** (required): A field that opens a list of available metrics for the selected measurement.

  + Includes a filter field to use text to narrow down the list, from which metrics can be selected individually, or as a group (with the **Select All** button).
  + Selected metrics appear as lozenges in the **Metrics** field.
  + Click the field to add more metrics or click the **X** at the right of its lozenge to remove a metric.
* **Group by Dimensions** (required): A field that opens a list of dimensions for grouping results.

  + Includes a filter field to use text to narrow down the list, from which dimensions  can be selected individually, or as a group (with the **Select All** button).
  + Selected dimensions appear as lozenges in the **Group-by Dimensions** field.
  + Click the field to add more dimensions or click the **X** at the right of its lozenge to remove a dimension.
* **Merge Series**: A field that opens a list of dimensions for creating a merge series (see **Measurement Pane Parameters**).

  + Includes a filter field to use text to narrow down the list, from which dimensions  can be selected individually, or as a group (with the **Select All** button).
  + Selected dimensions appear as lozenges in the **Merge Series** field.
  + Click the field to add more dimensions or click the **X** at the right of its lozenge to remove a dimension.
* **Aggregation** (shown only when dimensions are present in the **Merge Series** field): A dropdown that allows you to select a merge series aggregation method (sum/avg/min/max).

> **Note:** Metrics’ order in the **Metrics** field affects the **Table Sort Priority** of the query.

### Visualization Options Pane

This pane allows you to control query parameters affecting the visualization displayed in the **Metrics Explorer Main Display**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(444).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)

The **Visualization** pane includes the following settings:

* **Visualization**: A dropdown from which to choose the chart type. Options include Line (default), Area, Column, Pie, or Table Only (no visualization).
* **Metric**: A dropdown from which to select the metric for visualization. The default is the first metric displayed in the **Metrics** field (see **Measurement Pane UI**).
* **Aggregation**: A dropdown from which to select a method to aggregate data series values into a single value for the duration set with **Sample Size**.
* **Sample Size**: A dropdown to select the time frame of the individual buckets for aggregation. The default is Auto, where Kentik determines a feasible number of data points based on the time duration.
* **Series Count**: A dropdown to choose the number of series to plot. Without group-by dimensions, a query typically only results in a single series.
* **Transformation**: A dropdown that enables you to apply a transformation to the metric values:

  + **None** (default): Plot the metric's absolute value.
  + **Counter**: Plot the metric's rate based on the delta for each time slice, preventing continuous up-counting.

### Time Pane

The time pane specifies the query’s time period.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(445).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)

The pane includes the following controls:

* **Time**: A control set that allows you to define a time range.

  + **Time selector**: Shows the current time range (default: Last Hour) and opens a "Lookback list" of time range presets. To select a time range other than one of the provided presets, click **Custom** to open the **Custom Time Range Popup**.
  + **Forward/Back arrows**: Adjusts the time range while keeping its duration (e.g., shift “Last Hour” back by one hour).
* **Zone**: A dropdown from which to choose the time zone (UTC or Local), defaulting to your User Profile setting (see **User-specific Defaults**).

#### Custom Time Range Popup

The popup used to set custom time ranges includes the following controls:

* ![Date selection interface showing October and November 2023 with highlighted date.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/ME-calendar-selector.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)**Cancel**: A button that closes the popup and leaves the time range as it was before the popup was opened.
* **Apply**: A button that applies the time range from the values in the start and end fields and hides the popup. The applied range will be shown in the Time-range dropdown.
* **Lookback list**: Preset durations back from the current time, listed along the left side of the popup. Click a list item to set the time range. For example, if the current time is 11:00 and the chosen duration is Last 15 Minutes, the resulting time range will be from 10:45 to 11:00.
* **From**: Date and time fields that set the start of the time range.
* **To**: Date and time fields that set the end of the time range.
* **Calendars**: Side-by-side monthly calendars that show the time range (if it spans more than one day) and enable you to change the range (shown in the time range fields above) by clicking on a date. The start and end days of the time range are indicated in blue, and the intervening days are indicated in light gray. The calendar controls include forward and back buttons to change the displayed months as well as dropdown selectors for month and year.

### Filters Pane

![Filters section with options to add dimension and metric filters for data analysis.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/filter pane options.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)The **Filters** pane helps narrow your query scope or results. The number of filters specified in the pane is visible when the pane is collapsed; if no filters are being used, “No Filters” displays.

The pane includes the following controls:

* **Dimension Filters**: Displays applied dimension filters as tiles (see **Filters Dialog**). Remove a listed filter by clicking the red **X** in its tile.
* **Add/Edit Dimension Filters**: Buttons that open a **Filters Dialog** from which you can apply ad hoc filters on available metrics dimensions, similar to Group-by dimensions (see **Measurement Pane**). Example dimensions include `Device Name`, `Site`, and `ifDescr`.

  > **Note**: **Add Dimension Filters** displays when no dimension filters exist; **Edit Dimension Filters** displays when there are existing dimension filters.
* **Metric Filters**: A list of the currently active control sets that have been added with the **Add Metric Filter** button.
* **Add/Edit Metric Filter**: Buttons that open the **Metric Filters Dialog** allowing you to add a control set to filter metrics in the results. For example, filter CPU utilization to show only values over 90%.

  > Note: **Add Metric Filters** displays when no dimension filters exist; **Edit Metric Filters** displays when there are existing dimension filters.

#### Metric Filters Dialog

The **Filters** dialog for Metrics includes the following controls:

* **Group-level Controls**: Controls that apply to all/some filter groups.
* **Metric**: Select the metric for the filter from the dropdown. Available metrics are determined by the values selected in the **Metrics** field of the **Measurement Pane**.
* **Operator**: Apply an operator: equals/does not equal, greater than/less than.
* **Value**: Set the value to include/exclude by the filter.
* **X**: Remove this control set from the filter group.

### Table Options Pane

![Table options for data aggregation, including last option and row selection settings.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/table pane options.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A48%3A04Z&sr=c&sp=r&sig=fiWyvDdZA4YPsFM%2FXMHpSet92jgOuuUTuFC6OCNvCBw%3D)Table options determine how query results are shown in the table in the main display area.

The pane includes the following settings:

* **Aggregation**: A field that specifies method(s) to represent time series values as a single value for the entire time range:

  + Methods are shown as lozenges; default methods are Average and Last.
  + Remove a method  by clicking the **X** in its lozenge.
  + Add a method by clicking in the field to open a selection popup. Each method adds a column to the table.
* **Rows**: A dropdown from which you can choose the number of rows (the top-*X*) to display in the results table (default = 100).  
  **Example**: If × = 10, the table shows the top 10 rows sorted by **Table Sort Priority**.

> **Note:** The ordering of methods in the **Aggregation** field affects the **Table Sort Priority**.

## Run an NMS Query

To run a query with Metrics Explorer's **Query** sidebar:

1. From the main menu, click to open **Metrics Explorer**.
2. If not yet open, click **Query** in the **Metrics Explorer Subnav** to open the sidebar.
3. Click **Measurement** to expand the **Measurement Pane**, then specify **Metrics**, **Group-by dimensions**, and, optionally, Merge series and Aggregation fields (if applicable).
4. Click **Visualization Options** to expand the **Visualization Options Pane**, then specify the fields that determine how visualizations display.
5. Click **Time** to expand the **Time Pane**, then specify query’s time range.
6. Click **Filters** to expand the **Filters Pane**, then specify filters to narrow query’s scope or results.
7. Click **Table Options** to expand the **Table Options Pane**, then specify how results are shown in the table.
8. Click the **Run Query** button to execute the query (see **Query Sidebar Controls**) and view the results in the **Metrics Explorer Main Display**.

> **Note:** You can also specify queries with the **Metrics Query Assistant**.

The subnav is the gray, horizontal bar or strip located below the main navigation bar or navbar in various pages and modules. It typically contains page-wide controls and buttons specific to the page's functionality.
