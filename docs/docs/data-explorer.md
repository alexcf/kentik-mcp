# Data Explorer

Source: https://kb.kentik.com/docs/data-explorer

---

This article covers the features and use of Kentik’s Core » **Data Explorer** module.

> **Note:** The functionality of Data Explorer is also available via API (see **About V5 Query API** and **Show API Call**).

![Data Explorer enables query-based forensic and real-time exploration of traffic flows across your network.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DE-Main(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A45Z&se=2026-05-12T09%3A42%3A45Z&sr=c&sp=r&sig=hxOdZI5FOcutwh2Me66DOvuM90HxebL30Rh%2BdUIr684%3D)

*Data Explorer enables query-based forensic and real-time exploration of traffic flows across your network.*

## About Data Explorer

Data Explorer is Kentik's primary interface for exploring the network data (flow records, BGP, SNMP, etc.) stored in the main tables of the Kentik Data Engine (KDE, see **KDE Tables**). You can customize queries that display information about the traffic on specific data sources (routers, hosts, clouds, etc.) during a specific timespan.

Data Explorer “views” are customized using the **Query** sidebar (see **Query Sidebar Controls**). These controls specify query parameters such as time range, data sources, and dimensions, and narrow the returned data by filtering on dozens of dimensions (see **Dimensions Reference**). Views can be saved and reloaded at a later time (see **About Saved Views**). Results display as graphs and tables (see **Explorer Chart Display** and **Data Explorer Table**.

## Data Explorer Layout

Data Explorer contains the following areas:

* **Subnav controls**: Miscellaneous page-wide controls (see **Explorer Subnav Controls**).
* **Title bar**: An area just below the subnav with the name of the current view and various other information (see **Explorer Title Bar**).
* **Main display**: Shows query results in the form of a graph (see **Data Explorer Chart**), and an accompanying table (see **Data Explorer Table**).
* **Explorer Query sidebar**: A vertical sidebar that expands/collapses by clicking the **Query** button in the subnav. Use the sidebar contains the controls to specify the query, which displays results in the main display area (see **Query Sidebar Controls**).

#### Explorer Title Bar

The title bar in Data Explorer includes the following information:

* **Query title**: Automatically derived name based on the custom query. The title updates when changes are applied to the query sidebar.
* **Query settings summary**: Information about the query results, including the time range, the number of data sources, and the number of filters.

  > ***Note:*** To view and change the query settings, click the **Query** button (in subnav) to open the **Query** sidebar.

* **Analyze:** Click to enable Cause Analysis detection to identify the most likely causes of traffic fluctuations (see **Cause Analysis Overlay**).

## Explorer Subnav Controls

The Data Explorer subnav (silver strip across top of page) contains the following page-wide controls:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DE-SubNav_controls(2).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A45Z&se=2026-05-12T09%3A42%3A45Z&sr=c&sp=r&sig=hxOdZI5FOcutwh2Me66DOvuM90HxebL30Rh%2BdUIr684%3D)

* **Refresh**: Updates the currently displayed graph and table based on the existing query settings.
* **Auto Update** (drop-down arrow icon): Allows you to set Auto Update to Off (default) or select an interval at which the chart and table will automatically update (see **Auto Update Mode**):

  + When Auto Update is On, the interval (60, 90, or 120 seconds) will be displayed in the subnav next to the Auto Update arrow.
  + When Auto Update is set to Off, the Auto Update arrow will display alone.

    > **Note:** Auto Update is available only when the **Time Range** selector in the **Query** sidebar is set to Lookback (see **Time Pane**).
* **Share**: Enables you to share the current view via link, email, or subscription (see **Sharing via the Share Dialog**).
* **Saved Views**: Expands/Collapses the right sidebar to show other saved views in your organization (see **Saved Views Sidebar**).
* **Actions**: A drop-down menu that shows available actions for the current view (see **Data Explorer Actions**).
* **Query**: Expands/Collapses the **Query** sidebar, which contains the controls for customizing the query shown in the display area (see **Query Sidebar Controls**).

## Explorer Chart Display

The chart display area shows a visual representation of the current query results.

![The chart display area presents traffic data in a wide variety of time-series visualizations.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DE-Chart.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A45Z&se=2026-05-12T09%3A42%3A45Z&sr=c&sp=r&sig=hxOdZI5FOcutwh2Me66DOvuM90HxebL30Rh%2BdUIr684%3D)

*The chart display area presents traffic data in a wide variety of time-series visualizations.*

### Chart Display UI

The chart display area has the following UI components:

* **Restore**: A back-curved arrow icon that appears at the upper right of the chart when the page is in “period-over-period” mode and the Compare icon has been clicked in a table row (see **Compare Periods**). This button restores the chart from a period-over-period comparison back to its normal state.
* **Chart**: A visualization of query results. Choose from multiple visualization options in the **Visualization** pane of the **Query** sidebar (see **Chart Visualization Types**).
* **Table**: A corresponding table listing the query results in tabular form (see **Data Explorer Table**).
* **Resize handle** (between the chart and the table): Drag up or down to change the vertical allocation of the display area between the chart and the table.

### Data Explorer Chart

The Data Explorer chart is a visualization based on traffic flow records stored in the Kentik Data Engine (KDE). Available visualization types are detailed in **Chart Visualization Types**.

Most portal visualizations show the top results from the current query, ranked by the metric selected in the **Metrics Pane** of the **Query** sidebar. Many are based on time-series data, plotted over a time range (x-axis), for a given metric (y-axis; see **Time Pane**).

Each plot in the visualization corresponds to a row in the table below the chart (see **Data Explorer Table**). The number of table rows plotted in the chart depends on the visualization depth (see **Advanced Query Settings**). Each row corresponds to a colored disc in the table.  
  
 The chart in the display area is dynamic:

* Hover over any line in a line chart or area upper boundary in a time series stacked graph to see a popup containing data for a specific record at a specific point in time.
* Click, drag, and release in the chart to select a portion of the time range to zoom in on. When zooming:

  + The Time pane is automatically set to From + To with the start and end times defined based on the zoomed region;
  + The graph and table in the display area, along with its associated URL, update so that the zoomed range can be shared; and
  + A **Zoom out** button appears at the upper right of the graph, which can be clicked to zoom out to the previous time range.
* Clicking on the colored disc at the left of any row in the table will hide the area or line corresponding to that row from the chart. The disc will turn into a circle. Click the circle to restore the line or area in the chart.

### Auto Update Mode

Auto Update mode allows you to optionally set an interval — 60, 90, or 120 seconds — at which the Data Explorer graph and table will automatically be refreshed. The countdown to refresh starts over each time you apply changes, and the new result is returned in the display area. To enter Auto Update, choose the desired update interval from the drop-down **Auto Update** menu (see **Explorer Subnav Controls**). To exit, choose Off.

> **Note:** Auto Update mode is available only when the **Query** sidebar's **Time Pane** is set to **Lookback**.

### Cause Analysis Overlay

Cause Analysis is an AI-powered tool that automatically analyzes network traffic for the most impactful dimensions contributing to traffic patterns. When enabled in Data Explorer, it automatically creates an overlay on the graph to identify anomalies (spikes/drops in the chart), summarize likely cause(s), and show confidence percentages per dimension. This view supports three user workflows:

* **Traffic analysis**: Finds the most contributing traffic dimensions in a single time selection window.
* **Traffic comparison analysis**: Finds changes in traffic patterns between two selected time windows.
* **Automatic detection and analysis**: Automatically detects and analyzes spikes or drops in the time series data and compares traffic before and after each event.

This feature can be enabled one of two ways:

* Click the **Analyze** button in the chart display area (see **Explorer Title Bar**).
* Switch the “Enable Cause Analysis” toggle On in the **Kentik AI** pane of the **Query** sidebar.

The chart will automatically highlight identified events (spikes) in purple on the graph with a corresponding table below (see **Cause Analysis Table**). The chart display will change as Kentik AI options are customized in the **Kentik AI** pane of the **Query** sidebar (see **Kentik AI Pane**).

![The Data Explorer chart displaying Cause Analysis spikes related to traffic data. The numbers above the purple spikes are automatically numb](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DE-CA Chart.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A45Z&se=2026-05-12T09%3A42%3A45Z&sr=c&sp=r&sig=hxOdZI5FOcutwh2Me66DOvuM90HxebL30Rh%2BdUIr684%3D)

*The Data Explorer chart displaying Cause Analysis spikes related to traffic data. The numbers above the purple spikes are automatically numbered on the chart.*

## Data Explorer Table

The query results displayed in the chart display area are also presented as a table.

> **Note:**The structure of the table when the **Compare over previous period** switch is on in the **Time Pane** is covered in **Compare Periods**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(433).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A45Z&se=2026-05-12T09%3A42%3A45Z&sr=c&sp=r&sig=hxOdZI5FOcutwh2Me66DOvuM90HxebL30Rh%2BdUIr684%3D)

*The table provides a top-X list of the keys whose traffic matches the current query settings.*

### Explorer Table Overview

The Data Explorer table lists the values of selected metrics for the results returned from the current query in descending order. The last row will show the combined total of all records returned from the query. The table also doubles as a legend for the chart above; the rows that are marked with a colored disc at left are those that are plotted in the chart (the number of plotted rows is determined by the **Visualization Depth** setting (see **Advanced Query Settings**).  
  
 The location in which the table is displayed depends on the current visualization type (see **Chart Visualization Types**):

* When the visualization type is a graph or chart, the table is shown below the chart display area (see **Explorer Chart Display**).
* When the visualization type is set to Table, the table alone is displayed without a graph or chart. In this mode, the table itself can still be exported (see **Export Chart or Table**) or added to a dashboard (see **Add View to Dashboard**).
* The table is not shown when the visualization type is Matrix.

> **Notes:**
>
> * The number of rows in the results table that accompanies visualizations is dependent on the **Visualization Depth** setting and limited to a maximum of 350 unless the visualization type is Table (may include up to 50,000 rows depending on group-by dimension and metric).
> * When displaying results from a compound query (see **Compound Queries**), multiple tables are used, each on a separate tab corresponding to one axis (left/right) and/or direction (positive/negative).

#### Table AS Grouping

When the following conditions are met, then results from all ASes in each AS group are combined for top-X evaluation, graphs, and tables:

* Your organization has at least one Autonomous System (AS) group (see **About AS Groups**),
* The **Use AS Groups** switch is On in **Advanced Query Settings**, and
* A query's group-by dimensions include destination and/or source ASN.

If a table row represents a group, it will include a group icon at the left of the group name. Click the icon or name to open a list of the ASes in the group.

### Explorer Table Columns

The left-most columns of the table always correspond to the dimensions selected in the **Dimensions** pane of the **Query** sidebar (see **Dimension Panes**).  
  
The other columns depend on the metrics selected in the **Metrics Pane** (see **Metrics Settings**). The dialog allows you to customize which columns are shown. The default columns for a given metric will include the following:

* Average
* 95th Percentile
* Max
* Last Datapoint

> **Notes:**
>
> * The table will include a row (at bottom) for the combined total of all records returned from the query.
> * The table will include a row for historical values if Historical Overlay is On (see **Advanced Query Settings**).
> * The Last Datapoint column shows the value of the final datapoint in the time series displayed in the chart/table.

### Compare Periods

When the **Compare over previous period** switch is On in the **Time Pane**, the **Data Explorer Table** changes to a multi-tab structure used to compare periods ("period-over-period" comparison). The following tabs are included:

* **Current Period**: Contains a top-X table showing traffic data for the time range specified with the **Current** control in the **Time** pane.
* **Previous Period**: Contains a top-X table showing traffic data for a period specified with the **Previous** control in the **Time** pane.
* **Comparison Summary**: Contains a top-X table giving comparisons of the traffic during the two periods.

The columns of the table on the **Current Period** tab and the **Previous Period** tab are identical. They include all the columns present when the switch is Off, plus the **Compare** column. When you click a **Compare** icon (left and right arrows) in this column, the visualization in the **Data Explorer Chart** will show only two plots—solid for the Current period and dashed for the Previous period. To restore the full chart, you can either click the **Compare** icon again in the table or click the **Restore** icon at the top-right corner of the chart.  
  
The only columns on the comparison tab that are similar to the other tabs are for the query’s dimensions. The columns below are unique to this tab:

* **Percentage change**: The percentage by which the primary metric value for the current period differs from that of the previous period.
* **Current Average metric**: The primary metric value for the current period.
* **Previous Average metric**: The primary metric value for the previous period.

### Explorer Table Actions

Table actions can change the display of information in the table and the corresponding chart (see **Data Explorer Chart**):

* **Only display this row**: Click this action to display the selected row on the chart.
* **Display all rows but this one**: Click this action to display all rows but the selected row.

> **Notes:**
>
> * The far left column of the table contains the circles mentioned above.
> * Click each circle to hide/show the selected row(s).
> * The colored circle indicates that the row is display on the chart, and the empty circle indicates that the row is not being displayed on the chart.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(434).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A45Z&se=2026-05-12T09%3A42%3A45Z&sr=c&sp=r&sig=hxOdZI5FOcutwh2Me66DOvuM90HxebL30Rh%2BdUIr684%3D)

* **Include this in a new query**: The filter will include matching traffic using either the = or LIKE operator and rerun the query (see **Row Filter Actions**).
* **Exclude this in a new query**: The filter will exclude matching traffic using either the <> or NOT LIKE operator and rerun the query (see **Row Filter Actions**).
* **Show By**: Clicking this action results in two actions:

  + Opens a **Show By Dimension** dialog. The dimensions selected in this dialog will replace the current dimensions previously selected in the **Dimensions** pane of the Query sidebar.
  + Adds an **Include** filter as described in "Add filter to query" above.
* **Compare over** (shown only when the **Compare over previous period** switch in the **Time** pane is Off): Perform a "period-over-period" comparison (see **Compare Periods**) for the traffic data in this table row:

  + Hover over **Compare** over to open a drop-down from which you can select the period (day, week, or month) with which previous traffic should be compared.
  + The **Compare over previous period** switch will be turned On in the Time pane.
  + The chart will show a plot for current (solid) and previous (dashed) traffic over the currently specified time range.
  + The table will include **Current Period**, **Previous Period**, and **Comparison Summary** tabs.

#### Row Filter Actions

The filter actions for an individual row will rerun the existing query but with added filters that are based on the value of that row's dimensions:

* **Include this in a new query**: The filter will include matching traffic using either the = or LIKE operator.
* **Exclude this in a new query**: The filter will exclude matching traffic using either the <> or NOT LIKE operator.

The operation of the control depends on the number of dimensions in the query:

* If the query has one dimension, a filter to include/exclude that dimension can be applied directly from the table row's drop-down **Actions** menu.
* If the query has multiple dimensions, a submenu that lists each dimension individually will appear when you hover over the include/exclude option in the **Action** menu:

  + To apply filters for all of the dimensions, click in the menu.
  + To apply a filter for just one dimension, click on that dimension in the submenu.

Filters added via a row's **Actions** menu will appear in the **Filters Pane** in the **Query** sidebar.

#### Cause Analysis Table

The Data Explorer table for Cause Analysis displays information based on the type of workflow (see **Cause Analysis Overlay**). There are two tabs in the Cause Analysis table:

* **Left +Y Axis**: Displays the same table columns as the normal Explorer table view.
* **Cause Analysis**: Displays information about the type, change (measured in Gbits/s), time, and a summary of the analysis.

  + Click on an individual row to expand more information about the dimension, change (Gbits/s), and the estimated percentage of the total traffic represented by this grouping of dimensions. The Cause Analysis tab in the table will change as Kentik AI options are customized in the **Kentik AI** pane of the **Query** sidebar (see **Kentik AI Pane**).

![The Cause Analysis table contains auto-generated information by Kentik AI. The table rows are organized in order of the spikes displayed in](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DE-CA Table.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A45Z&se=2026-05-12T09%3A42%3A45Z&sr=c&sp=r&sig=hxOdZI5FOcutwh2Me66DOvuM90HxebL30Rh%2BdUIr684%3D)

*The Cause Analysis table contains auto-generated information by Kentik AI. The table rows are organized in order of the spikes displayed in the chart.*

## Data Explorer Actions

The **Actions** menu (in subnav) provides multiple ways to use query results outside of the Data Explorer module (e.g., save to a panel on a Dashboard or share with other Kentik users). These capabilities are covered in the following topics:

### Export Chart or Table

The Export action shows the Export submenu allowing you to export the information represented by the chart and/or table to an external file (PDF, PNG, SVG, or CSV). The available export options are detailed in **Portal Export Options**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DE-Actions(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A45Z&se=2026-05-12T09%3A42%3A45Z&sr=c&sp=r&sig=hxOdZI5FOcutwh2Me66DOvuM90HxebL30Rh%2BdUIr684%3D)

> **Note:** If Data Explorer is currently displaying the results of a compound query (see **Compound Queries**), then the Export submenu will list the Chart Data and Legend Data options individually for each axis (e.g., positive and negative).

After clicking the export option, a popup notification will display noting that the export is processing. When the file is ready, another popup notification indicate the export is complete and provide a link to download the file.

![](https://lh7-rt.googleusercontent.com/docsz/AD_4nXfVIRGXo5r3X47JpfXV94DKGTJPMzpwGf1kjinAtHnP2Ep-hjiJj4vR5QNJZ6v52rCBavvlq-z5aKOE8kWNhAVlOLo9JUU_GZn-RXoWRUWDpdzLXsj6sR-84672ErordY88QZBjJ_1mFmTdt1dB3Q?key=RZDV-dY2pN4o86jAvsLgPw)

*Example of the two popup notifications for exporting a chart or table.*

### Add to Observation Deck

This action takes you to your **Observation Deck** and adds the current Data Explorer visualization to the page as a widget. To undo, click the kebab menu (vertical ellipsis) in the top-right corner of the widget, then choose Remove.

### Create Saved View

Create Saved View opens the **Saved View** dialog (see **Saved View Dialogs**), used to set the properties of a new saved view (see **About Saved Views**). Once created, a Saved View can be accessed from the **Saved Views** tab of the Library (see **Library**).

### Create Alert Policy

The Create Alert Policy action opens the Add Traffic Alert Policy page. Current criteria specified in Data Explorer's **Query** sidebar (dimensions, metrics, filters) will be used to create a Traffic-type policy (see **Policy Types**).

> ***Note:*** If dimensions are not selected in Data Explorer, then this action will result in display of the **Unsupported Alert Policy Settings** dialog. Click **Cancel** to return to the main Data Explorer page.

### Add View to Dashboard

There are two options that enable you to add a the current Data Explorer view to a panel on a dashboard:

* **Add to New Dashboard**: Opens the **Add Dashboard** dialog (see **Add Dashboard from Explorer**). Specify settings for the new dashboard, then click the Add Dashboard button. The new dashboard will open and will display the new panel.
* **Add to Existing Dashboard**: Opens the **Add View Panel** dialog (see **View Panel Dialog Settings**). In the **Dashboard** pane, choose an existing dashboard, then click the **Add View Panel** button. The dashboard will open the newly added panel, and a confirmation notification.

### Preview as Tenant

This action opens a submenu to choose a MKP tenant. Settings of the Data Explorer view will be modified to show the chart and table as it would appear to that tenant .

### Show API Call

This action displays a dialog of code (cURL or JSON) that can be used to return the current view from the Kentik **Query API**. Copy and past the code to enable access to Kentik functions programmatically rather than through the portal.  
  
 Query API code is accessed via the following dialogs:

* **For Chart** (cURL): Opens a dialog containing the cURL for returning an image of the Data Explorer's current chart from a CLI such as Terminal. Equivalent to the **Query Chart Method** of the Kentik Query API.
* **For Data** (cURL): Opens a dialog containing the cURL for returning the Data Explorer's current table from a CLI such as Terminal. Equivalent to the **Query Data Method** of the Kentik Query API.
* **JSON Input**: Opens a dialog containing JSON that can be used in the Query Data Method.

When using the cURL, the following placeholders must be replaced with the appropriate information:

* Replace `<YOUR_EMAIL_HERE>` with the email address used to register you as a Kentik user.
* Replace `<YOUR_API_TOKEN_HERE>` with your API token, which you'll find on your **User Profile**.
* If the cURL is for a chart, replace `<CHOOSE ONE OF:pdf|png>` with the desired file type.

### Reset to Default Query

This action reverts all settings on the **Query** sidebar (Visualization, Data Sources, Dimensions, etc.) to their default settings to quickly customize a new query without checking individual settings in the sidebar's panes.
