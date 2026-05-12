# Dimension Settings

Source: https://kb.kentik.com/docs/dimension-settings

---

This article discusses **Dimensions** settings for queries in **Data Explorer** and dashboard panels in the Kentik portal.

## Dimension Panes

The Query sidebar controls in **Data Explorer** and dashboard panels include two panes for choosing the dimensions that are used for queries:

* **Dimensions** pane: Indicates the currently selected group-by dimensions, and enables access to the Group By Dimensions dialog, which has two tabs:

  + **Preset**: See **About Dimensions**
  + **Filter-Based**: See **Filter-Based Dimensions**.
* **Matrix With** pane: Visible only when the visualization type is Matrix (see **Matrix Visualization**). Indicates the currently selected matrix-by dimensions, and enables access to the Matrix By Dimensions dialog.

> **Note:** The Group By Dimensions and Matrix By Dimensions dialogs are effectively identical (see **Dimension Selectors**).

### Dimensions Pane UI

The **Dimensions** pane includes the following UI elements:

* **Dimensions**: A list of the dimensions that are currently selected for the query:

  + When no dimensions are selected, the list is replaced with "No dimensions selected".
  + Each item in the list has a handle that you can click and drag to change the order of the items.
  + An individual item in the list can be removed by clicking the **X** at right.
* **Edit Dimensions**: A button that opens the **Group By Dimensions** selector (see **Dimension Selectors**).
* **CIDR**: A pair of fields that appear only when the selected dimension includes a CIDR component (e.g., Source IP/CIDR):

  + **v4 CIDR**: Use to specify the number of bits of the routing prefix in IPv4. Default is 32.
  + **v6 CIDR**: Use to specify the number of bits of the routing prefix in IPv6. Default is 128.
  > **Note:** If you have only v4 traffic, the v6 CIDR field will be ignored (can be left at 128).

### Matrix With Pane UI

The **Matrix With** pane appears only when the **Chart Visualization Type** is Matrix. The pane includes the following UI elements:

* **Dimensions**: A list of the **Matrix With** dimensions that are currently selected for the query:

  + When no dimensions are selected, the list is replaced with "No dimensions selected".
  + Each item in the list has a handle that you can click to change the order of the items.
  + An individual item in the list can be removed by clicking the **X** at right.
* **Edit Dimensions**: A button that opens the **Matrix By Dimensions** selector (see **Dimension Selectors**).

## Dimension Selectors

The Matrix By Dimensions dialog and the Preset tab of the Group By Dimensions dialog are used to choose preset dimensions for the **Matrix With** selector and **Group By Dimensions** selector, respectively.

> **Notes:**
>
> * To query on total traffic, don't select a group-by dimension (the **Dimensions** pane will say "No dimensions selected").
> * For fastest results, where possible use native dimensions (rather than virtual) as indicated for each dimension in the **Dimensions Reference**).
> * A query can use no more than eight group-by dimensions and eight matrix-with dimensions.
> * The Group By Dimensions dialog also includes a tab that can be used to choose a filter-based dimension, which will override the preset dimensions in the **Group By Dimensions** selector (see **Filter-Based Dimensions**).

### Group By Dimensions Dialog UI

The query dimension dialogs share the following common UI elements:

* **Close** (button): Click the **X** in the upper right corner to close the dialog. All elements will be restored to their values at the time the dialog was opened.
* **Filter** (field): Filters the list of available dimensions to those containing the entered text.
* **Selected Dimensions**: A list of selected dimensions (up to eight). Use the handle at the left of each selected dimension to change the order in which the dimensions are applied.
* **Available dimensions**: A grid containing of all of the dimensions that are currently supported by Kentik, listed by group (see **Dimension Selection Groups**).
* **Clear Selections** (button): Clears all dimensions from the **Selected Dimensions** list.
* **Cancel** (button): Cancel changes to the selected dimensions and exit the dialog. The current query's dimensions will be restored to what they were when the dialog was opened.
* **Save** (button): Save changes to the selected dimensions and exit the dialog.

### Dimension Selection Groups

The dimensions available in the dimension selector dialogs are based on the columns of the main tables of the KDE (see **KDE Tables**), which each represent a minute of flow records for a given device. The types of dimensions available for querying are discussed in **Dimension Categories**, which includes links to categorized lists giving a description of each individual dimension.

#### Dimension Directionality

In the dimension selector dialogs (as well as in the dimension selector for filters; see **Ad Hoc Filter Controls**), the available dimensions are organized into a grid with the following columns:

* **Source**: Source traffic.
* **Destination**: Destination traffic.
* **Source or Destination** (Filter controls only): Matches on either source or destination traffic (saves adding two filters).
* **Non-Directional/Other**: Traffic attributes that are non-directional.

#### Dimension Functional Categories

The rows of the grid organize the available dimensions into collapsible panes by functional category:

* **Network & Traffic Topology**: Interface, connectivity, provider, etc.; see **Network and Traffic Topology**.
* **IP & BGP Routing**: Dimensions related to IP address, protocol, TCP flags, and ToS, as well as routing information related to ASes, paths, community, prefixes, and hops; see **IP and BGP Routing** and **BGP Overview**.
* **Cloud**: AWS, GCP, Azure, etc.; see **Cloud Dimensions** and **Public Clouds**.
* **Geolocation**: Country, region, city, etc.; see **Geolocation Dimensions** and **Custom Geos**.
* **Application Context & Security**: Dimensions related to services (port and protocol) such as CDN and OTT as well as to threats; see **Threat Feed Dimensions**.
* **Container Networking**: Kubernetes properties, PATA dimensions, etc.; see **Container Networking Dimensions**.
* **Device-Specific Dimensions**: A series of panes for Cisco ASA, Juniper PFE, etc.; see **Device-Specific Dimensions**.
* **Application Decodes**: DNS lookup, HTTP, TLS, DHCP, Radius; see **Host Traffic Dimensions**.
* **Device Metrics**: Synthetic, SNMP, and Streaming Telemetry; see **Device Metrics Dimensions**.
* **Azure Network Metrics**: Express Route Metrics
* **Application Attribution**: DNS
* **Events**: SNMP Traps, Device Config Change, Syslog
* **Custom**: **Custom Dimensions**.

### Choosing Query Dimensions

To choose group-by dimensions (the same procedure applies for matrix-with dimensions):

1. By default (e.g., when you first navigate to **Data Explorer**), there are no dimensions selected in the **Dimensions** pane. Click the **Edit Dimensions** button to open the Group By Dimensions dialog.
2. On the **Preset** tab of the dialog, you'll see the **Available Dimensions** list (see **Group By Dimensions Dialog UI**). At right above the list is a filter box that you can use to narrow the dimensions in the list.
3. Click the checkboxes for the dimensions that you want to use in the query, which will then appear in the **Selected Dimensions** list.
4. To change the order in which the dimensions are applied, drag the dimensions in the **Selected Dimensions** list into the desired order.
5. Click **Save** to run the query using the currently selected dimensions.

> **Notes:**
>
> * When more than one group-by dimension is selected, the combination of dimensions is evaluated together to determine the rows that are included in the results; see **Using Multiple Dimensions**.
> * The **Matrix With** selection box operates the same as the **Group By Dimension** selection box (see **Matrix Visualization**).

### Using Multiple Dimensions

The following example illustrates how multiple group-by dimensions combine to determine the results returned from a query. The example shows a common use case for multiple dimensions, which is when an organization that generates traffic wants to see where it's going and which links and devices it's using to get there (which enables you to see if the traffic is going to the expected geographic locations in the expected proportions):

1. Using the **Metrics Pane**, set the metric to Bits/s and set **Primary Display & Sort Metric** to Average.
2. Use the **Time Pane** to set the time range.
3. Use the **Filters Pane** to filter so that you are seeing only outbound traffic. (Assuming that **Interface Classification** shows that at least 75 percent of your interfaces are classified, the easiest way to do this is to set a Destination Network Boundary filter to External.)
4. In the chart display settings (see **Chart Display UI**) set the **Visualization Type** to Sankey Flow Diagram.
5. Use the **Data Sources Pane** to select multiple devices.
6. Use the **Group-By Dimension** selector to set following group-by dimensions:

   * **Full**: Device
   * **Destination**: Next Hop AS Number
   * **Destination**: AS Number
   * **Destination**: Region
   * **Destination**: Country
7. Click **Save** to exit the dialog. The graph and table will update to show traffic categorized by the specified dimensions.

In the results:

* The **Data Explorer Table** will include a column for each dimension in the key (the five group-by dimensions specified above).
* All traffic with the same dimension value for each component of the key will be measured in Mb/second and summed onto a single table row, meaning that each row will represent traffic that has the same `device`, `dst_next-hop-asn`, `dst_as`, `dst_region`, and  `dst_country`.
* The rows will be listed in descending order of highest average bits/second.
* The Sankey diagram will show the paths of the traffic represented by the top table rows.

## Filter-Based Dimensions

Filter-based dimensions allow a single query result (graph and table) to combine a set of time series that can each have different filters applied. You can specify a nominally unlimited number of series, each with its own specific filters; collectively this set is referred to as a single filter-based dimension (of which you can have only one at a time). As shown in the screenshot below, for example, you might want to look at the traffic across your network to two specific ASes and compare the balance of traffic to each AS from two different source countries.

Note that a query can’t mix filter-based dimensions with regular (“preset”) group-by dimensions. Any dimensions that you already have in the group-by selector will be overridden when you save a filter-based dimension.

### Filter-Based Tab

Filter-based dimensions are configured on the **Filter-Based** tab of the Group By Dimensions dialog. The tab includes the following elements:

* **Enable filter-based dimension**: A switch that turns on/off filter-based dimensions. When the switch is on, the dimension configured on the **Filter-Based** tab overrides the group-by dimensions set with the **Preset** tab.
* **Dimension Name**: A user-specified name string for the dimension. The name will be displayed as the query title in the display area (see **Chart Display UI**) and in the **Group By Dimensions** selector in the **Dimensions** pane of the Query sidebar.
* **Auto-Add Other Series**: Include a series for all traffic that has been filtered out of the series configured in the Series section.
* **Series**: A set of configuration UI for the series that make up the filter-based dimension (see **Series UI**).

### Series UI

The controls of the **Series** pane of the Group By Dimensions dialog are similar to those of the **Ad-Hoc Filter Groups** pane of the **Filters Dialog**. The interface includes the controls described in the topics below.

#### Series Controls

The following UI elements are high-level controls for the set of series that make up a filter-based dimension:

* **Remove All Series**: Removes all filter groups. Located at top right of **Series** pane.
* **Add Series button**: Adds a new filter group to the set. Located below the last series.

#### Series-level Controls

A single container for one or more individual filters, which includes the following controls:

* **Name**: A user-specified string for the name that will be used to identify the series in results table and in the key of the graph.
* **Include/Exclude**: A drop-down selector that determines whether results that match the group are included or excluded from the results of queries to which the filter set is applied.
* **All/Any**: A drop-down selector that, if there are multiple filters in a filter group, determines the conjunctive operator used to join those filters:

  + **All** (default): Used to AND the filters.
  + **Any**: Used to OR the filters.
* **Remove Series**: Removes the entire filter group.
* **Add Nested Group**: Adds a filter group that is nested within the series.
* **Add Condition**: Adds a filter to the group.

#### Filter Controls

An individual filter in a group, which includes the following controls:

* **Dimension selector**: Opens a dialog enabling you to choose a dimension on which to filter (see **Dimension Selector Dialog**).
* **Direction button** (shown only when the filter dimension supports directionality): toggles between source and destination.
* **Operator**: The operator to apply in the filter. Options vary depending on dimension, but may include:

  + equals/does not equal
  + greater than/less than
  + contains/does not contain (case insensitive)
  + matches regex/does not match regex
* **Value**: The value to match.
* **Remove** (**X**): Removes the filter from the group.

#### Dimension Selector Dialog

The dimension selector presents a list from which you can choose a dimension on which to filter. The dialog includes the following elements:

* **Filter field**: Filters the available dimensions to show only those whose name contains the entered text.
* **Dimension lists**: Lists of the available dimensions in each category, e.g., source, destination (see **Dimensions Reference**).

### Setting a Filter-based Dimension

To set a filter-based dimension for a query in **Data Explorer**:

1. In the Dimensions pane of the **Query** sidebar, click the **Edit Dimensions** button.
2. In the resulting Group By Dimensions dialog, click the **Filter-Based** tab.
3. Click the **Enable filter-based dimension** switch to activate the controls of the tab.
4. Fill in the Dimension Name field, and specify the **Auto-Add Other Series** button (see **Filter-Based Tab**).
5. Configure the series that you want to appear in the graph and table (see **Series UI**).
6. Click **Save** to apply the Filter-based dimension, which will appear in the **Group By Dimensions** selector. The filter-based dimension will override any preset dimensions specified for the query.
