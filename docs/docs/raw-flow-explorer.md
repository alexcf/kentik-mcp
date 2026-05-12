# Raw Flow Explorer

Source: https://kb.kentik.com/docs/raw-flow-explorer

---

The **Raw Flow Explorer** module of the Kentik portal enables you to directly examine the flow data stored in Kentik Data Engine (KDE).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(437).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A05Z&se=2026-05-12T09%3A34%3A05Z&sr=c&sp=r&sig=zKhwZgwsy5kmA3bJtGeu3U14znWKU2td3dM89w30Ktg%3D)

*Raw Flow Explorer enables you to examine flow records from Kentik's KDE backend.*

## About Raw Flow Explorer

KDE is the repository of the network traffic data that is collected, stored, and queried by Kentik (see **Kentik Data Engine**). The core of this data is flow data that is exported from network devices that support NetFlow, sFlow, or other flow protocols (see **About Flow**). At time of ingest into KDE, the flow records containing this data are timestamped, augmented with a wide variety of related information, and used to populate the columns of each record in the KDE. Raw Flow Explorer makes it possible to find and examine the value stored in KDE for one or more KDE columns in an individual time-stamped flow record.

> **Notes:**
>
> * The retention period for raw flow data is seven days.
> * For an overview of the various types of data collected in KDE, see **What traffic data is collected?**.
> * For a list, with descriptions, of the filtering and group-by dimensions based on KDE data, see **Dimensions Reference**.

## Raw Flow Explorer Page

The Raw Flow Explorer page, reached via the main navbar (Core » **Raw Flow Explorer**), is made up of the following main areas:

* **Subnav controls**: Miscellaneous page-wide controls (see **Raw Flow Explorer Subnav**).
* **Explorer Query Sidebar**: An area at the right (or left depending on your configured Layout Preferences) that contains a set of control panes (see **Raw Flow Query Sidebar**).
* **Display area**: An area for display of a table containing raw flow records (see **Raw Flow Explorer Table**).

Settings are made in the sidebar and then applied to update the results in the display area.

## Raw Flow Explorer Subnav

The Raw Flow Explorer subnav (silver strip across top of the page) contains the following page-wide controls:

* **Share** (appears only after Query is run): Opens the Share dialog to enable you to share the current view (see **Sharing via the Share Dialog**).
* **Actions**: A drop-down menu with actions you can take involving the current view:

  + Refresh: Updates the currently displayed table based on the current query settings.
  + Export CSV: Exports the information represented by the table to an external CSV.
* **Query**: Toggle visibility of the **Query** sidebar (see **Sidebar Panes**), which contains the controls used to specify the query whose results are returned in the display area. By default the sidebar is visible when the page is opened.

## Raw Flow Query Sidebar

The **Query** sidebar contains the following UI elements:

* **Run Query button**: A button that applies changed sidebar settings to the table in the display area.
* **Sidebar**: A set of panes that are used to set values for the queries whose results are shown in the display area (see **Sidebar Panes**).

### Sidebar Panes

The **Query** sidebar contains the following panes that control the query whose results are displayed in the display area:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(439).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A05Z&se=2026-05-12T09%3A34%3A05Z&sr=c&sp=r&sig=zKhwZgwsy5kmA3bJtGeu3U14znWKU2td3dM89w30Ktg%3D)

* **Data Sources pane**: Displays the data sources (e.g., devices) currently selected for the query, and enables access to the dialog for selecting data sources (see **Data Source Settings**).
* **Query pane**: Displays the dimensions currently selected for the query, and enables access to the dialog for selecting dimensions. It is nearly identical to the Dimensions pane in Data Explorer (see **Dimension Settings**).
* **Time pane**: Sets the time range for the query (see **Time Pane**).
* **Filters pane**: Displays the filter groups and conditions currently applied to the query, and enables access to the dialog for specifying filters (see **Filters Settings**).

#### Query Pane

The query pane includes the following UI elements:

* **Dimensions**: A list of dimensions that are currently selected for the query (see **Dimensions Pane UI**).
* **Edit Dimensions**: A button that opens the **Group By Dimensions** selector (see **Dimension Selectors**).
* **Order By**: A popup used to choose the dimension by which the results of the query will be ordered (see **Order By Dropdown**).
* **Row Count**: A field to specify the number of records to display in the **Raw Flow Explorer Table**. The maximum row count is 15000.

#### Time Pane

Time-range settings in the Query sidebar's **Time** pane define the timespan covered by the query whose results are displayed in the **Raw Flow Explorer Table**. When the **Time** pane is expanded, it includes the following controls:

* **Time Range**: A three-part control to display and change the time range of the query:

  + Shift back: The arrow at the left end of the control shifts the time range back by the range's current duration, e.g., if the range is one hour starting at 11 AM today it will shift back to one hour starting at 10 AM today.
  + Range drop-down: The currently defined time range is displayed in the central part of the control. Click the control to open the Time Range popup to choose a different lookback duration or define a different custom range
  + Shift forward: The arrow at the right end of the control shifts the time range forward by the range's current duration. If shifting forward would extend the range beyond the present the control will be inactive.
* **Use AWS Timestamps**: Choose the timestamps to use for flow records originating in a cloud export from AWS (see **Cloud Exports and Devices**). By default the switch is off, meaning that flow records are timestamped upon intake by Kentik. When the switch is on, the timestamp for records from AWS flow logs will instead be AWS's timestamp.
* **Unapplied Changes**: A lozenge informing the user that the changes have not been applied.

#### Order By Dropdown

The **Order By** dropdown contains the following UI elements:

* **Search**: Filters the grid to show dimensions containing the entered text.
* **Dimensions list**: A grid containing dimensions organized by directionality and functional categories. the list includes:

  + Categories belonging only to Raw Flow Explorer (see **Order By Functional Categories**).
  + Standard dimension categories (see **Dimension Selection Groups**).

#### Order By Functional Categories

The rows of the grid organize the available dimensions by functional category:

* **Time**: Dimensions related to time such as Timestamp and Flow Start.
* **Metrics**: Dimensions related to sampling such as Sampled at Ingress packets and Sampled at Egress bytes.

## Raw Flow Explorer Table

The Raw Flow Explorer table is used to examine the value stored in KDE for one or more KDE columns in an individual time-stamped flow record. The table displays the results of the query and is made up of the following:

* **Search**: Filters the Raw Flow Explorer table to rows that contain the entered text in any column.
* **Columns**: A column for every dimension selected with the **Dimensions** selector in the **Query** pane. You can click on any heading to sort the table ascending or descending based on that column.
* **Rows**: A row for each result returned. The number of rows displayed is specified with the **Row Count** selector in the **Query** pane.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(440).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A05Z&se=2026-05-12T09%3A34%3A05Z&sr=c&sp=r&sig=zKhwZgwsy5kmA3bJtGeu3U14znWKU2td3dM89w30Ktg%3D)

*The table displays the results of the query, enabling you to see the dimension values in individual flow records.*

---

© 2014-25 Kentik
