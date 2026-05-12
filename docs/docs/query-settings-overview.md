# Query Settings

Source: https://kb.kentik.com/docs/query-settings-overview

---

This article provides an overview of the settings for queries in the Kentik portal, including Data Explorer and dashboard panels.

![Settings for queries are made int he panes of the sidebar that opens from the Query button.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/QS-Main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A47%3A04Z&sr=c&sp=r&sig=nK1k0QXyawHBDm03myCTsIlDIxI96YutmzDHu1XqxP8%3D)

*Query settings are made in the panes of the Query sidebar.*

## Query Sidebar Controls

Several sections of the Kentik portal include a Query sidebar that contains a set of controls presented as a series of panes. The controls in the panes are used to set values for queries whose results (graph and table) are shown in the main display area of the page. The standard set of controls includes the following:

* **Visualization**: Sets the chart type (see **Chart Visualization Types**) for the query as well as various result display options that are specific to individual views (see **Visualization Pane**).
* **Data Sources** pane: Displays the data sources (e.g., devices) currently selected for the query, and enables access to the dialog for selecting data sources (see **Data Source Settings**).
* **Dimensions** pane: Displays the group-by dimensions currently selected for the query and enables access to the dialog for selecting dimensions (see **Dimension Settings**).

  > **Note:** If no dimensions are specified, then the query will run on the total traffic (less traffic filtered out, if any).
* **Matrix With** (shown only when the Visualization Type is set to Matrix; see **Matrix Visualization**): Displays the matrix-with dimensions for a matrix query. This pane is nearly identical to the **Dimensions** pane (see **Matrix Pane UI**).
* **Metrics** pane: Displays the primary metric currently selected for the query and enables access to the dialog for selecting primary and secondary metrics (see **Metrics Settings**).
* **Time** pane: Sets the time range for the query (see **Time Pane**).
* **Filters** pane: Displays the filter groups and conditions currently applied to the query and enables access to the dialog for specifying filters (see **Filters Settings**).
* **Brackets** pane: Displays the upper bound and the color of each defined range (see **Bracketing Settings**).
* **Advanced** pane: A collection of controls used to fine-tune the query (see **Advanced Query Settings**).
* **Kentik AI**: Controls that enable and customize Cause Analysis (see **Kentik AI Pane**).

> **Note:** Some of the above controls are documented in this article while more complicated settings are covered in other sections.

#### Query Settings Locations

The following table shows which of the above-listed panes appear in the various portal locations that include sidebars with query settings:

| Portal section | Data Sources | Time | Filters | Other |
| --- | --- | --- | --- | --- |
| Data Explorer | Yes | Yes | Yes | • Visualization • Dimensions • Metrics • Brackets • Advanced  • Kentik AI |
| Dashboards, main sidebar | Yes | Yes | Yes | N.A. |
| Dashboards, Panel dialog | Yes | Yes | Yes | • Visualization • Dimensions • Metrics • Brackets • Advanced |

## Visualization Pane

The Visualization pane is used to specify how query results will be displayed as a visualization. It includes the following controls:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/QS-Visualization.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A47%3A04Z&sr=c&sp=r&sig=nK1k0QXyawHBDm03myCTsIlDIxI96YutmzDHu1XqxP8%3D)

* **Visualization Type**: A drop-down menu used to set the type of visualization (see **Chart Visualization Types**) in the chart display area.
* **Sync Axes**: Sets the scale of both left and right axes to the same numeric values (even if the metric is different). This option is present only when **Bi-Directional Charting** is on (see **Advanced Pane UI**) or **Secondary Metric** is set to not None (see **About the Metrics Dialog**).

  > **Note:** The greater the difference in value between the scale of the two axes, the greater the likelihood that syncing the axes will effectively hide data plotted against the lower-value axis.
* **Use Logarithmic Scale**: When On, this switch causes one or more axis of the visualization to be plotted against a logarithmic scale (see **Logarithmic Scale Options**).
* **Visualization depth**: Determines how many rows (1-40) will be plotted in the graph. The higher the visualization depth, the more detailed the chart. There will also be a decrease in the gap between the individual plotted data and the blue line representing Total.

  > **Note:** When Visualization Type is set to Table, the Visualization Depth determines how many rows will be displayed in the table.

#### Logarithmic Scale Options

Logarithmic scale switches are present for the following view types:

* **Stacked Area Chart**:

  + Use Logarithmic Scale on Primary Axis
  + Use Logarithmic Scale on Secondary Axis
* **Line Chart**:

  + Use Logarithmic Scale on Primary Axis
  + Use Logarithmic Scale on Secondary Axis
* **Stacked Column chart**:

  + Use Logarithmic Scale on Primary Axis

## Time Pane

Time-range settings in the sidebar's **Time** pane define the timespan covered by the query whose results are displayed in the graph and table in the display area. When the Time pane is expanded its initial control set includes the following:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(479).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A47%3A04Z&sr=c&sp=r&sig=nK1k0QXyawHBDm03myCTsIlDIxI96YutmzDHu1XqxP8%3D)

* **Compare over previous period**: A switch that enables a "period over period" comparison. It compares the results from the current time range with results from the same previous period. The switch is Off by default. When the switch is On, the **Time Range** drop-down is labelled "Current" and the **Previous** controls are displayed.

  > **Notes:**
  >
  > + This switch results in multiple changes to the content and behavior of the Data Explorer table and charts. For information about these changes, see **Compare Periods**.
  > + When Cause Analysis is enabled, “Compare over previous period” is disabled.
* **Time Range**: A three-part control to display and change the time range of the query:

  + **Shift back**: The arrow at the left end of the control shifts the time range back by the range's current duration (e.g., if the range is one hour starting at 11 AM today it will shift back to one hour starting at 10 AM today).
  + **Range drop-down**: The time range is displayed in the central part of the control. Click the control to open the **Time Range** popup to choose a different lookback duration or define a different custom range.
  + **Shift forward**: The arrow at the right end of the control shifts the time range forward by the range's current duration. The control will become inactive if the shift extends beyond the range.
* **Previous** (shown only when **Compare over previous period** is On): A two-part control for setting the time range against which the current time range will be compared:

  + **Count**: A field in which you can enter the number of units back from the current range.
  + **Unit**: A drop-down from which you can select a duration unit, either hour, day, week, or month.
* **Zone**: The time zone is either UTC or local.

  > **Note:** To change the default for this setting, see user **User-specific Defaults**.
* **Use AWS Timestamps**: Choose the timestamps to use for flow records originating in a cloud export from AWS (see **Cloud Exports and Devices**). The switch is Off by default, meaning that flow records are timestamped upon intake by Kentik. When the switch is On, the timestamp for records from AWS flow logs will instead be AWS's timestamp.

> **Note:** When zooming in to the display area (see **Data Explorer Chart**), the time range is automatically set to the start and end times of the zoomed region.

#### Custom Time Range Settings

When Custom is chosen from the **Time Range** drop-down, a popup opens to show the following controls that defines a range:

* **Cancel**: A button that closes the popup and discards time range changes.
* **Apply**: A button that applies the time range changes and closes the popup. The applied range will be shown in the Time-range Drop-down.
* **Lookback list**: Preset durations back from the current time, listed along the left side of the popup. Click a list item to set the time range. For example, if the current time is 11:00 and the chosen duration is ‘Last 15 Minutes’, the resulting time range will be from 10:45 to 11:00.
* **From**: Date and time fields that set the start of the time range.
* **To**: Date and time fields that set the end of the time range.
* **Calendars**: Two monthly calendars are shown side-by-side. You can change the range by clicking on a ‘from’ date and ‘to’ date. The calendar controls include forward and back buttons to change the displayed months and drop-down selectors to change the month and year.

## Advanced Query Settings

Kentik supports advanced options that tailor queries to return more precise visualization. The controls for these options are in the **Advanced** pane of the **Query** sidebar in Data Explorer and in the configuration dialog for individual dashboard panels.

### Advanced Pane UI

The **Advanced** pane includes the following controls:

* **Dataseries**: A drop-down menu to set the resolution of the KDE dataseries for running the query (Auto, Full, or Fast; see **Resolution Overview**).

  > **Note:** The available dataseries settings depend on the current time range.
* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(480).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A47%3A04Z&sr=c&sp=r&sig=nK1k0QXyawHBDm03myCTsIlDIxI96YutmzDHu1XqxP8%3D)**Show Total Overlay**: Enables/disables plotting of the total traffic returned from the query. The effect varies depending on visualization type:

  + **Stacked and bar charts**: Total appears in the graph as a blue line above the other plots.
  + **Line chart**: Total appears in the graph as a dashed blue line above the other plots.
  + **Comparison bar chart**: Total appears as a blue line at right.
  + **Pie chart**: If total overlay is On, the chart includes a ring segment labeled Other (traffic not plotted on the other segments); if Off, the Other segment is not included.
  + **Sankey diagram**: If On, a row for total is included in the table below the diagram.
  + **Matrix**: If On, a total line will be plotted on the detail graphs that are rendered below the matrix when you click on a matrix cell, column heading, or row heading (depending on the view type selected for those graphs; see **Matrix Detail Graph**).
  > ***Note:*** Switching **Show Total Overlay** to Off rescales the axis, which may improve the visibility of smaller-value plots.
* **Show Historical Overlay**: Enables/disables plotting of the total from the same query run on a time range from a number of days earlier. Historical values are plotted as a dashed gray line. **Historical Overlay** is not available if the **Visualization Type** is set to Time Series Line Graph and is automatically switched Off when the type is Pie Chart.
* **Days Back**: If **Historical Overlay** is On, this field sets how many days back to show in the historical display. For example, if the time range is the last 6 hours, the current time is 11:00, and the ‘days back’ is set to 7 (default), the historical plot will show the total from 5:00 to 11:00 seven days ago.
* **Enable Reverse DNS Lookups**: Determines whether reverse DNS (rDNS) lookup will be performed to determine associated host names when querying IPs.

  > **Notes:**
  >
  > + When this switch is On and the query includes an IP/CIDR dimension, the host name for each IP will appear in parentheses in the IP/CIDR column of the results table.
  > + To use an alternate DNS server (instead of the default) for reverse lookup, see **Custom DNS**.
  > + Queries return faster when this option is Off.
* **Use AS Groups**: When this switch is On the results from all ASes of each AS group (see **About AS Groups**) will be summed for top-X evaluation, graph plotting, and display in the results table.

  > **Note:** The switch is visible when at least one AS group exists in your organization. When visible, the switch is enabled by default.
* **Bi-Directional Mode**: Enables simultaneous charting of two graphs, one based on the current group-by dimensions and the other based on the opposite of those dimensions (see **Compound Queries**). This is only shown if visualization type is stacked or line (see **Chart Visualization Types**).
* **Generate One Chart Per Series**: Changes the visualization in the display area from a single chart plotting multiple top-X series to a set of charts that each show a single top-X series (see **Generated Charts**).
* **Extract From**: Enables group-by on substrings in certain DNS/WWW dimension values (see **DNS/WWW Extract Function**). This is only shown for specific group-by dimensions.
* **Time Series**: Settings enabling greater control over the aggregation of returned data into a time series in the query's chart and top-X table (see **Time Series Settings**).

### Compound Queries

The **Bi-Directional Mode** switch (**Advanced Pane UI**) enables compound queries that include graphs resulting from multiple simultaneous underlying queries. Compound queries fall into two general categories, but in some cases the categories can be applied simultaneously to create a chart incorporating four graphs:

* **Bi-Directional**: An “original” graph of traffic is based on the current **Group-by Dimension** setting, and an “opposite” graph is based on the opposite of those dimensions. For example, if the group-by dimensions are Source Country and Destination AS Number, the opposite graph would show traffic based on Destination Country and Source AS Number.

  > **Note:** Filters are also reversed for the opposite view, meaning that filters on source in the original are on destination in the opposite, and vice versa.
* **Secondary Metric**: A “primary” graph of traffic is based on the primary **Metric** setting, and a “secondary” graph is based on the **Secondary Display Metric** setting (see **Metrics Dialog UI**).

> **Note:** The two categories of compound queries are independent. You can create either a Bi-Directional query with no secondary metric or a secondary metric query that's not Bi-Directional (see **Compound Query Variations**).

#### Compound Query Variations

As detailed in the table below, the results (graphs and tables) that are returned from a compound query depend on the interaction of several settings, including:

* The **Bi-Directional Mode** switch in the **Query** sidebar’s **Advanced** **pane**
* The **Secondary Display Metric** in the **Metrics** dialog
* The **Visualization** setting in the chart display area (see **Chart Visualization Types**)

| Visualization | Secondary display metric | Bi-Directional mode | Chart axes | Chart directions | Table tabs |
| --- | --- | --- | --- | --- | --- |
| Stacked or Line | None | Off | Left only | Positive only | Primary only |
| Stacked | None | On | * Left: group-by dimensions * Right: opposite dimensions | * Positive: group-by dimensions * Negative: opposite dimensions | * Original: group-by dimensions * Opposite: opposite dimensions |
| Stacked | not None | Off or On  **Note**: This setting is ignored unless Secondary Display Metric is None. | * Left: primary metric * Right: secondary metric | * Positive: primary metric * Negative: secondary metric | * Primary metric * Secondary metric |
| Line | None | On | Left only | * Positive: group-by dimensions * Negative: opposite dimensions | * Original: group-by dimensions * Opposite: opposite dimensions |
| Line | not None | Off | * Left: primary metric * Right: secondary metric | Positive only | * Primary metric * Secondary metric |
| Line | not None | On | * Left: primary metric * Right: secondary metric | * Positive: group-by dimensions * Negative: opposite dimensions | * Original group-by, primary metric * Original group-by, secondary metric * Opposite dimensions, primary metric * Opposite dimensions, secondary metric |

#### Compound Query Considerations

The following considerations apply when using compound queries:

* In-line charts, plots against the left axis are drawn with a solid line, while plots against the right axis are dashed.
* On Bi-Directional charts, the flipping of dimensions (e.g., from Source ASN to Destination ASN) for opposite graphs applies only to dimensions in the Source and Destination groups (see **Dimension Selection Groups**). Dimensions in the Full and DNS/WWW categories are treated the same on both original and opposite graphs.
* Dimensions in the Custom category (see **Custom Dimensions**) will be reversed only if there are two dimensions with identical names except in one of the following ways:

  + One includes "src" where the other has "dst."
  + One includes "in" where the other has "out."
  + One includes "to" where the other has "from."

### Generated Charts

Generated charts are visualizations made up of a set of individual charts, each of which represents one top-X series in the current query. This feature is enabled by the **Generate One Chart Per Series** switch (see **Advanced Pane UI**), and is used to define a dashboard panel (see **Add View to Dashboard**).

Like standard single-chart visualizations, generated charts are shown in the Data Explorer display area (see **Explorer Chart Display**) above the query’s results table. The number of query's visualization depth sets the number of generated charts. If brackets are set (see **Bracketing Pane**) then each graph will be outlined in the color of the brackets range corresponding to the value (in the brackets metric; e.g., average Gbits/s) of the series graphed in that chart.

#### Generated Charts UI

The following UI elements are shown in the **Advanced** pane of the **Query** sidebar when the **Generate One Chart Per Series** switch is On:

* **Chart Titles**: The title that will appear at the upper left of each chart.

  + By default, the title of each chart is set with the `generator_series_name` variable, which is based on values for dimensions that make up the query's key. For example, if the query's group-by dimensions are Source Country and Destination Country then each chart name is a combination of a source country code and a destination country code (e.g., "US ---- GB").
  + If you enter a string into this field, then all of the generated charts will have the same title (the entered string).
* **Charts per Row**: The number of charts (up to four) that will appear (horizontally) in each row of the display area.
* **Chart Height**: The vertical space allocated to a row of charts. The settings (Short, Medium, and Tall) are relative; their actual effect depends on the height of the browser window and the position of the Resize bar (see **Chart Display UI**) that divides the display area from the results table.
* **Chart-level Group By Dimensions**: A dimension selector that opens a **Group By Dimension** dialog (see **Dimension Selectors**) that enables you to choose one or more group-by dimensions. The top-X for the chosen dimensions will be calculated and plotted individually for each chart rather than across the entire query.
* **Chart Visualization Depth** (shown only when there is a panel dimension): Determines the number of series (top-X combinations of panel group-by dimensions) that will be shown in each chart.

> **Note:** If there isn’t enough space to show all the lines in a chart’s key, Up and Down arrow icons will appear below the chart. Click these icons to scroll through the key lines.

### DNS/WWW Extract Function

The dimensions available in the **Group By Dimension** selector vary depending on the device type. If any device currently selected in the **Data Sources** pane (see **Data Source Settings**) is a host of type “Kentik Host Agent”, the dimension selector will include a set of DNS/WWW dimensions (see **Host Traffic Dimensions**).

When some of these DNS/WWW dimensions are selected, the **Extract****From** settings will appear in the **Advanced** pane of the **Query** sidebar. These **Extract****From** settings are available for the following dimensions:

* DNS Query
* HTTP URL
* HTTP Host Header

The Extract settings change how Kentik evaluates the values of the dimensions listed above:

* With no extraction, each unique value in the column will be treated separately.
* With extraction, a regex-defined pattern will be used to find matching substrings within the values, and all values that match the same substring will be grouped together.

For example, if you run a Top-X query with DNS Query as the dimension and set the extract function to Domain, then all subdomains of the same domain (e.g., *x.domain.com*, *y.domain.com*, and *z.domain.com*) will be counted together as one value, instead of separately.

To apply the extract function, choose the type of substring that you want to match using **Extract****From** settings drop-down menu. The Regex and Selector fields will populate with the suggested POSIX-style regex (shown in the following table) and selector for that type of match. You can edit the regex and selector as needed to achieve the desired result.

| Substring type | Regex: DNS Query, HTTP Host Header | Regex: HTTP URL |
| --- | --- | --- |
| TLD | ['.]+\.(['.]+)$ | N.A. |
| Domain | (['.]+\.['.]+)$ | N.A. |
| Subdomain | (['.]+\.['.]+\.['.]+)$ | N.A. |
| Host | ('.\*)\.['.]+\.['.]+$ | N.A. |
| Path | N.A. | ('.\*)/.\*(\?)+ |
| Filename | N.A. | '.\*(/.\*)(\?)+ |
| File Path | N.A. | ('.\*)(\?)+ |
| Query String | N.A. | '.\*[\?]+(.\*) |

As explained in **About Regex Capture Groups**, the selector (e.g., `$1`) will give the index (1-based) of a “capture group,” which is a substring that is surrounded by parentheses in the regex. For example, in the regex for Path in the table above, `$1` would refer to a substring that matches the pattern `'.*` (the first substring is surrounded by parentheses).

> **Note:**For additional information on or assistance with using the extract function, contact support@kentik.com.

### Time Series Settings

Queries are performed on Kentik-stored network traffic data (flow records) in two stages:

* **Initial flow aggregation**: For all flows matching the query’s filters, the total number of bytes or packets over the selected time is added up. The total is then divided by the value of the flow aggregation window setting and grouped by the query's selected dimensions. This results in a set of sample series that represent an average bit/packet rate across all of the summed flows.
* **Time series construction**: The top-X sample series from the initial flow aggregation are aggregated into time series buckets of a specified width using the specified aggregation technique.

The settings of the **Time Series** controls in the **Query** sidebar's **Advanced** pane enable you to customize how aggregation will be performed during the two stages described above:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(481).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A04Z&se=2026-05-12T09%3A47%3A04Z&sr=c&sp=r&sig=nK1k0QXyawHBDm03myCTsIlDIxI96YutmzDHu1XqxP8%3D)

* **Flow Aggregation Window**: The time window used for initial flow aggregation (see **Flow Aggregation Window**). The default setting is Auto.
* **Bucket Aggregation**: The function used to construct time series buckets from flow aggregation samples:

  + **None**: No further aggregation of the initial flow aggregation into time series buckets.
  + **Average**: Calculates the average value of the flow aggregation samples in each bucket.
  + **Maximum**: Selects the flow aggregation sample with the maximum value in each bucket.
  + **Minimum**: Selects the flow aggregation sample with the minimum value in each bucket.
* **Width** (hidden when **Bucket Aggregation** is set to None): The width (in time) of the "buckets" to group flow data when creating a time series for top-X results:

  + If set to Auto, the bucket width will be one hour.
  + If set to any other value (10/20/30 min, 1/6/12 hours, 1 day, or 1 week) the bucket width must be longer than the flow aggregation window.

#### Flow Aggregation Window

A query's flow aggregation window depends on multiple factors including the type of dataseries (see **About Dataseries Resolution**) and the duration of the query's time range. The table below shows the windows that are available for various time ranges.

| Dataseries | Time range | Window |
| --- | --- | --- |
| Full | * up to 3 days * up to 14 days * up to 30 days * up to 60 days | * 1 minute * 5 minutes * 10 minutes * 20 minutes |
| Fast | Any | * 1 hour * 6 hours * 12 hours |

## Kentik AI Pane

The Kentik AI pane includes the following controls:

* **Enable Cause Analysis**: Turns Cause Analysis On/Off.
* **Analysis Type**: A brief description automatically generated by Kentik AI. For example, Change Point Detection identifies points of major increase or decrease and determines what traffic primarily contributed to that change.
* **Min Support Percent**: Ranges from 0-100% and sets the minimum percentage an itemset must appear to be considered frequent. Higher percentages filter out rare itemsets.
* **Change Point Detection (CPD) Limit**: Ranges from 1-15 and limits how many change points (major traffic shifts) the system will detect.
* **CPD Change Threshold**: Ranges from 0-100% and sets the minimum percentage change needed for a change point to be considered significant, filtering out small, less important changes.
* **Show/Hide Advanced Options**: Shows/hides more options for Kentik AI in Data Explorer (see **Kentik AI Advanced Options**).

### Kentik AI Advanced Options

The following are additional options for Kentik AI in the **Query** sidebar:

* **Dimensions**: Choose dimensions in the dialog for Kentik AI to analyze in addition to the previously selected dimensions (see **Dimension Settings**).

  + **Only use selected dimensions in analysis**: When switched On, only selected dimensions will be analyzed.
* **Pruning Percent:** Ranges from 0-100% and removes similar or overlapping item sets if their frequency difference is less than this percentage, keeping only the. most unique ones.
* **Max Itemset Length**: Ranges from 0-20 and limits how many data points in an itemset. Lower values reduce complexity but may exclude larger, meaningful patters.
* **CPD Window Size (samples)**: Ranges from 0-20 and sets how many consecutive data points are analyzed together when detecting significant changes. Larger windows may smooth out noise but delay detection.
* **CPD Percentage Threshold**: Ranges from 0-100% and only displays changes within this percentage of the max.
