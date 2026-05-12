# Chart Visualization Types

Source: https://kb.kentik.com/docs/chart-visualization-types

---

This article covers the types of visualizations available in the Kentik portal.

## About Visualization Types

Kentik supports a variety of visualization (chart) types to view traffic data stored in the Kentik Data Engine (KDE). Most visualizations are based on the "top-X rows" as measured by the metric selected in the **Metrics Pane** of the **Query** sidebar. Many are based on time-series data, plotted over a time range (horizontal axis; see **Time Pane**) and the metric represented on the vertical axis. For most visualization types (excluding Gauge) the chart is accompanied by a results table (see **Data Explorer Table**).  
  
Visualizations are used throughout the Kentik portal along with data tables that also serve as legends for the chart. Most portal pages have a default chart type that best illustrates the data. In the following areas of the portal you can choose the chart type for the visualization:

* **Data Explorer**:The visualization that appears in the **Explorer Chart Display** can be chosen from the **Visualization Pane** in the **Query** sidebar.
* **Dashboard panels**: Each panel displayed on a preset or custom dashboard includes a drop-down menu to select chart type (see Change Visualization in **View Panel UI**).
* **Saved views**: When you create a saved view in Data Explorer you can choose the chart type for your view (see **Create Saved View**).

> **Notes:**
>
> * The number of top rows displayed in a visualization depends on the Visualization Depth (see **Visualization Pane**).
> * The number of rows in the results table that accompanies visualizations depends on the Visualization Depth (maximum of 350 rows).
> * If Historical Overlay is On (see **Advanced Query Settings**) then visualization types will plot historical values as a dashed gray line.

## Time-based Visualizations

The following charts are visualizations of time-series data.

### Stacked Area Chart

This is the default chart type. The top rows in the returned data are stacked to show the contribution of each row to the combined total. This visualization helps you understand how the top series changes over time in relation to one another.

**Example:** This visualization works well to address the following question, "What are the TOS categories for the traffic on this interface, and how much bandwidth are they using?"

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CVT-Stacked.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

*Visualization Type: Stacked Column Chart; Visualization Depth: 8, Dimensions: Source City, Destination City, Site Country; Metrics: bits/s; Time: Last Hour; Filters: None; Brackets: Off*

### 100 Percent Stacked Area Chart

This chart type is similar to a standard stacked graph, but it provides an indication of relative value rather than absolute value. Each row is shown as a percentage of the total traffic returned from the query. The entire area of the graph is always filled, but the proportion of vertical space represented by each row of results changes over time. This is useful for seeing changes over time in the top series, not only in relation to one another but also in relation to the total.

**Example:** This visualization answers the question, "What percent of the traffic we're serving is going to each Country, or to each Custom Geo?"

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CVT-100 Stacked.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

*Visualization Type: Stacked Column Chart; Visualization Depth: 8, Dimensions: Source City, Destination City, Site Country; Metrics: bits/s; Time: Last Hour; Filters: None; Brackets: Off*

### Stacked Column Chart

The top rows in the returned data are plotted against an absolute scale over the specified time range, showing the contribution of each row to the combined total. The duration represented by each bar is determined by the query's aggregation step boundaries, which depends on the width of the query timespan.

**Example:** The aggregation of values into segments of time can answer the question, "Over the last week, how much traffic per hour was outbound from the cloud networks that we're monitoring?"

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CVT-Stacked Column.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

*Visualization Type: Stacked Column Chart; Visualization Depth: 8, Dimensions: Source City, Destination City, Site Country; Metrics: bits/s; Time: Last Hour; Filters: None; Brackets: Off*

### Line Chart

The top rows in the returned data plotted over a specified time range against an absolute scale. Since the values are overlapping rather than stacked, this visualization emphasizes the absolute value of an individual series rather than the total resulting from the combination of all plotted series. The combined total is Off by default, but it can also be plotted.

**Example:** This visualization answers the question, "How much traffic are we sending on each of our outward-facing interfaces?"

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CVT-Line.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

*Visualization Type: Line Chart; Visualization Depth: 8, Dimensions: Source City, Destination City, Site Country; Metrics: bits/s; Time: Last Hour; Filters: None; Brackets: Off*

### Horizon Chart

A horizon chart combines time-based and density-based information into a compact visualization that displays the metric (e.g., traffic volume in bits/second) over time for the top-X rows returned from the query. Each top row is represented by one “lane,” and Visualization Depth determines the number of lanes (see **Advanced Query Settings**). The height of all lanes is the same regardless of the amplitude of the represented data.

Within each lane, the overall amplitude range is divided into five “bands” that are assigned different colors. Bands are overlaid within the lane (rather than stacked), starting with the lowest-volume band at the back. The highest-volume band is always visible in the foreground. This approach is useful for identifying series or time slices that are outliers (high or low). It is also more informative than a stacked area chart when there are numerous series.

**Example:** A horizon chart can answer the question, "Are any of the devices we're monitoring sending more flows than normal, and have any of them stopped sending flows?"

The **Sync Density Scale Across Series** switch in the Visualization pane of the **Query** sidebar determines how the bands for each lane are scaled:

* **Compare density**: If the switch is On, then all lanes will use the same scale, and the bands each represent one fifth of the total range from zero to the highest value that is present in any of the lanes. This visualization is best for comparing density (e.g., traffic volume) across the lanes.
* **Identify patterns**: If the switch is Off, each lane uses its own scale, and the bands each represent one fifth of the total range from zero to the highest value in that lane. This visualization makes it easier to identify patterns across the lanes (e.g., do they all have a peak at the same time).

Hover over the chart at a given point in the timeline to display the value (measured in Gbits/s) of each lane at that point.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CVT-Horizon.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

*Visualization Type: Horizon Chart; Visualization Depth: 8, Dimensions: Source Country, Source City, Destination City, Site Country; Metrics: bits/s; Time: Last Hour; Filters: None; Brackets: Off*

## Comparative Visualizations

The following visualizations compare total cumulative values over the entire time-range of a query.

### Bar Chart

The top rows in the returned data are plotted as individual bars against a horizontal axis representing the metric being counted, and the combined total is also plotted. This visualization is useful for comparing series aggregates over time, where series vs. series is more important than series vs. total.

**Example:** This visualization answers the question, "How much traffic did we send on each of our peering and transit links this month?"

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CVT-Bar.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

*Visualization Type: Bar Chart; Visualization Depth: 8, Dimensions: Destination Country, Source Country; Metrics: bits/s; Time: Last 1 Day; Filters: None; Brackets: On*

### Pie Chart

In this visualization, the top rows are shown individually as a colored segment of the pie, with the sum of all other rows shown together as "Other." This is useful for comparing series aggregates over time, where series vs. total is more important than series vs. series.

**Example:** This chart addresses the question, "What's the ratio of IPv4 to IPv6 in the past quarter?"

> **Notes:**
>
> When the query is run for a Pie Chart visualization type:
>
> * Display and Sort By will be automatically set to Average.
> * If Historical Overlay is On (see **Advanced Query Settings**), a Historical Total row will display in the **Data Explorer Table**, but it will not display in the pie chart.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CVT-Pie.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

*Visualization Type: Pie Chart; Visualization Depth: 8, Dimensions: Source Country, Source City, Destination City, Site Country; Metrics: bits/s; Time: Last Month; Filters: None; Brackets: Off*

## Hierarchical Visualizations

The following visualizations show the break-out of subsets within sets of data.

### Sankey Diagram

A Sankey diagram is used to see connections across the various dimensions that make up a key (see **About Keys**), such as how much of a given source IP's traffic is going to each of its various destination IPs via various intermediate hops. It's useful when visualizing these relationships across three or more related dimensions. Each dimension in the diagram is represented by a colored vertical bar, and the width of the gray bands between the bars is proportional to the quantity of traffic in common between those dimensions.

**Example:** This visualization answers complex questions, such as, "For the traffic that is being sent from my network to the Netflix AS, what protocols and ports are being used, how much of it originated inside my network, and what was the source AS of the outside traffic?"

> **Note:**Unless the group-by dimension is Destination BGP AS Path, you must have at least two group-by dimensions to use the Sankey visualization type.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CVT-Sankey.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

*Visualization Type: Sankey; Visualization Depth: 8; Dimensions: Source City, Source Traffic Origination, Destination Traffic Termination, Destination City; Metrics: bits/s; Time: This month; Filters: None; Brackets: Off*

### Sunburst Chart

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CVT-Sunburst.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

*Visualization Type: Sunburst; Visualization Depth: 8, Dimensions: Source Country, Source City, Destination City, Destination Country; Metrics: bits/s; Time: This month; Filters: None; Brackets: None*

A Sunburst chart is similar to a Sankey in terms of visualizing relationships among dimensions in a key (see **About Keys**), but better for visualizing each dimension's contribution to the total. The dimensions that comprise the key definition are represented as concentric rings (from inner to outer). They are segmented into wedges representing the top rows returned from the query (segments with the same value for a given dimension are joined into one). The number of segments is determined by the **Visualization Depth** (see **Advanced Query Settings**).

**Example:** A sunburst chart can answer the question, "What's the relative distribution of traffic volume by TOS/DSCP, and for each value, what are the destination protocols, ports, and IP/CIDR subnets?"

## Density-based Visualizations

The following types of density-based visualizations illustrate the volume of traffic.

> **Note:** Matrix visualization is also part of this category (see **Matrix Visualization**).

### Gauge Visualization

A gauge presents a single primary metric value without a time-series graph or a table, which is useful for presenting a top-level KPI or metric (e.g., total bits/sec in and out for the entire network) on a dashboard. To use the Gauge visualization type, configure and enable Brackets (see **Bracketing Pane**), including the following settings in the Brackets Options dialog:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CVT-Gauge.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

*Visualization Type: Gauge; Visualization Depth: 8, Dimensions: Source Country, Source City, Destination City, Destination Country; Metrics: bits/s; Time: Last 1 Day; Filters: None; Brackets: On*

* The **Use Last Datapoint Value** switch determines which metric is displayed as the primary value (large type in left column), as well as which values are displayed in the secondary column at the right of the gauge.
* The **Brackets Range** settings determine the background color of the gauge, which is the color assigned to the range into which the current value of the primary metric falls.

### Geo HeatMap

A Geo HeatMap is a zoomable map that displays the geographic distribution of a metric, for example the traffic volume associated with various physical locations. Each geography is colored according to its rank relative to other geos.  
  
Map visualization supports one group-by dimension at a time. Supported dimensions are shown in the following table:

| Category | Source | Destination | Non-Directional/Other |
| --- | --- | --- | --- |
| Network & Traffic Topology | — | — | Site |
| Network & Traffic Topology | — | — | Ultimate Exit Site |
| Geolocation | Custom Geo | Custom Geo | — |
| Geolocation | Country | Country | — |
| Geolocation | Region | Region | — |
| Geolocation | City | City | — |
| Geolocation | — | — | Site Country |
| Geolocation | — | — | Ultimate Exit Site Country |

Brackets are not required, but if they are enabled, then the map colors will correspond to the ranges of brackets (see screenshot below).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CVT-GeoHeat.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

*Visualization Type: Geo HeatMap; Visualization Depth: 8, Dimensions: Source Country; Metrics: bits/s; Time: Last Week; Filters: None; Brackets: On*

## Matrix Visualization

Matrix visualizations are covered in the following topics.![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DE-View_matrix-674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

### About Matrix Visualizations

When the visualization type is set to Matrix, the display area will show a matrix, which is a table in which:

* The vertical axis (rows) represents the currently chosen group-by dimensions.
* The horizontal axis (columns) represents the currently chosen matrix-with dimensions.
* The values in the cells are expressed in the currently set metric.
* The number of rows/columns is determined by the **Visualization Depth** setting (see **Advanced Query Settings**), with an upper limit of 15.

### Matrix Visualization Queries

The table in a matrix visualization is populated (behind the scenes) using three successive queries (where X represents the Visualization Depth):

1. Get the top-X instances, measured by current metric (e.g., average pkts/sec), of the currently chosen group-by dimension (e.g., Src cities) across all currently selected devices. These become the rows of the matrix.
2. Get the top-X instances, measured by current metric, of the currently chosen matrix-with dimensions (e.g., Dst cities) across all currently selected devices and filtered to include only the results of the first query. This becomes horizontal axis (columns).
3. Get the traffic volume, in currently selected metric, between the group-by and matrix-with dimensions at each row/column intersection. These values populate the corresponding cells of the matrix.

> **Notes:**
>
> * A matrix may include multiple group-by and/or matrix-with dimensions.
> * In addition to the filtering described in query 2 above, all other filters specified in the **Filters** pane of the sidebar will also be applied to the queries.
> * Variations in cell background color correspond to the scale at the right of the matrix itself.

### Matrix Visualization Example

In the following example:

* The metric is packets/second, Visualization Depth is set to 8
* The Display and Sort By is Average
* The Group-by Dimension is Source: City
* The matrix-with dimension is Destination: City
* The vertical axis shows the top 8 source cities as measured in average packets/second
* The horizontal axis shows the top 8 destination cities filtered by the source cities
* The cells of the table are populated with average packets/second between the cities on the two axes

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DE-Matrix_view-674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

#### Using Matrix Visualization

To use the Matrix visualization type in Data Explorer:

1. Click the Query button to open the **Query** sidebar.
2. Use the panes of the sidebar to set the desired data sources, metrics, time range, and filters.
3. Click the **Edit Dimensions** button in the sidebar’s Dimensions Pane to open the Group-by Dimensions dialog.
4. Choose one or more dimensions for the rows in the Matrix table (see **Dimension Selectors** and **Using Multiple Dimensions**), then click the Save button.
5. In the sidebar's **Visualization Pane**, choose Matrix from the Visualization Type drop-down. The Matrix-by Dimensions dialog opens.
6. Choose one or more dimensions for the columns in the Matrix table, then click the Matrix by Selected Dimensions button. The dialog will close, the Matrix With pane appears in the Query sidebar, and the chart in the main display area is now a matrix.

### Matrix Detail Graph

Once the matrix is rendered, you can click an individual cell to open a pop-up with a visualization of the corresponding data. You can select an alternate visualization type from the drop-down menu in the Visualization pane, and you can save the chart by choosing **Create Saved View** from Data Explorer's Actions menu.

The following example shows a detail graph of the Los Angeles row of a matrix, rendered with a visualization type of Sankey diagram.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DE-Matrix_detail-674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

> **Notes:**
>
> * When creating a detail graph from an individual cell, don't click directly on the cell value.
> * The visualization type setting for a detail graph is sticky until the overall matrix is reset by applying sidebar changes, at which point detail graphs will once again be rendered using the default visualization type.
> * The effect of the Total Overlay switch on a detail graph depends on visualization type (see **Chart Display UI**).

## Table Visualization

The table is displayed without a graph or chart. The table can be exported (see **Export Chart or Table**) or added to a dashboard (see **Add View to Dashboard**). Tables are useful for queries where there are many keys with relatively equal weight.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/CVT-Table.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A55Z&se=2026-05-12T09%3A41%3A55Z&sr=c&sp=r&sig=ZY%2B2Kh2iN7xlmZPFIPSzoNEc31bzlntN6nr4gtGmlAU%3D)

*Visualization Type: Table; Visualization Depth: 100, Dimensions: Source Country; Metrics: bits/s; Time: Last Week; Filters: None; Brackets: Off*
