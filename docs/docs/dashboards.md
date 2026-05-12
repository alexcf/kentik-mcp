# Dashboards

Source: https://kb.kentik.com/docs/dashboards

---

This article discusses **Dashboards** in the Kentik portal.

> **Note:** For information on dashboard edit controls, settings, and properties, including adding and editing dashboards, see **Dashboard Configuration**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(374).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A30Z&se=2026-05-12T09%3A44%3A30Z&sr=c&sp=r&sig=7H%2FPDw1OZdaRoUuHVx0UmSfJtqwym7OyyJuca3LC5kc%3D)

*Dashboards display a custom-configured collection of panels that each contain a visualization of your organization's traffic.*

## About Dashboards

Dashboards are accessed from the **Library List** on the portal's Library page. Both your organization’s custom dashboards and Kentik-provided preset dashboards (**Kentik Presets**) are found here. The Library page is also where you'll find the drop-down Add menu, from which you can choose to add a new dashboard.

> **Notes**:
>
> * Any query that can be displayed in **Data Explorer** can be saved to a panel on a dashboard.
> * The contents of any dashboard may be sent as a regularly scheduled report that is emailed to a specified list of subscribers (see **Report Subscriptions**).

### Dashboard State

In its default state a dashboard simply displays the panels that have been assigned to it. To modify dashboard characteristics, such as panel layout or the number and content of individual panels, the dashboard must be in "edit" state. To enter edit state, choose **Edit** from the drop-down Actions menu at the top right of the dashboard. The layout and controls of a dashboard in edit state are covered in **Dashboard Edit Controls**.

> **Note:** There is no edit state for dashboards that are provided by Kentik (see **Kentik Presets**) rather than created by your organization.

### Dashboard Modes

The Kentik portal supports two dashboard modes:

* **Guided mode**: Guided mode lets users see data about a preselected dimension without setting complex query options. When enabled, the dashboard will include a prompt (see **Guided Mode UI**), and the user’s response will determine the value of a predefined filter that is applied to each panel. For information on enabling Guided mode, see **Guided Mode Pane**.
* **Standard mode**: In standard mode, a dashboard has no prompt at the top of the display area. Each panel is independent unless the **Device Behavior** or **Filter Behavior** property on the panel is set to **Controlled by Dashboard** (see **View Panel Dialog Settings**), in which case the panel tracks the current device and/or filter settings of the sidebar.

### Dashboard Visibility

Dashboard visibility determines which users can access a given dashboard:

* **Private**: Dashboards that you create only for your own personal use. No one else has access.
* **Shared**: Dashboards created by any user in your organization. All users in your organization have access. All Kentik preset dashboards are Shared.

### Dashboard Navigation

Dashboard navigation enables users to navigate from a given dashboard panel directly to another dashboard. The destination dashboard would typically be related to the same use case, thus reducing the time spent drilling deeper into a given area.  
  
Using dashboard navigation begins with thinking through a particular use case, including what information you're going to want to start with in a given situation and what would be the next set of information you'd want to see if further investigation is called for.  
  
Once you've settled on a workflow through successive dashboards, and have chosen or created the individual dashboards that will be used, you can begin to designate dashboards as destinations that are reached from a panel on a source dashboard (see **Add Dashboard Navigation**).

> **Note:** The **Dashboard Dependencies Pane** of the Dashboard Properties dialog for a given dashboard shows the dashboards that are navigated to and from that dashboard.       

## Dashboard Layout

A dashboard is made up of the following main areas:

* **Subnav controls**: Page-wide controls in the subnav (see **Dashboard Subnav Controls**).
* **Title bar**: The title bar stretches across the top of the dashboard’s display area, identifying the current dashboard and containing a set of information indicators. For further information, see **Dashboard Title Bar**.
* **Query sidebar**: A popup sidebar (on right) that contains the controls used for overall Dashboard settings, as well as for settings on panels whose time range, devices, and/or filters have been specified as "Controlled by Dashboard." For further information, see **Dashboard Query Sidebar**.
* **Display area**: An area for display of individual panels. For further information, see **Dashboard Panels**.

> **Note:** The layout and controls of a dashboard in edit state are covered in **Dashboard Edit Controls**.       

## Dashboard Subnav Controls

The subnav (silver strip across top of page) contains the following page-wide controls:

* **Refresh**: A button that updates the charts and tables in the dashboard's panels.
* **Auto Update** (down-pointing triangle icon): A drop-down menu from which you can set auto-update to Off(default) or select an interval at which the charts and tables will automatically update:

  + When auto-update is On, the interval (60, 90, or 120 seconds) will be displayed in the subnav to the left of the down-pointing triangle.
  + When auto-update is Off, the down-pointing triangle will be displayed alone (no auto-update indicator).
  > **Note:** The countdown to refresh starts over each time you apply changes and the new result is returned in the display area.
* **Share**: A button that opens the Share dialog to enable you to share the current dashboard (see **Sharing via the Share Dialog**).
* **Actions**: A drop-down with the options covered in **Dashboard Actions**.
* **Dashboards**: A button that toggles visibility of the right-side **Dashboards Sidebar**, which enables you to link directly to other dashboards in your organization.
* **Query**: A button that toggles visibility of the right-side **Query** sidebar, which contains the settings used to control the dashboard (see **Dashboard Query Sidebar**).

#### Dashboard Actions

The **Actions** menu for dashboards includes the following options:

* **Edit**: Puts the dashboard into edit state so you can make changes (see **Edit a Dashboard**).
* **Settings**: Opens the Edit Dashboard dialog where you can change dashboard properties (see **Dashboard Dialogs**).
* **Clone**: Opens the Clone Dashboard dialog, which has the same fields and controls as the Edit Dashboard dialog but without the **Dashboard Dependencies** pane (see **Dashboard Navigation**). When property settings are complete for the clone click the Add Dashboard button to create the cloned dashboard.
* **Export**: Prepares a visual report (PDF) including all panels of the dashboard. A notification appears when the PDF is ready to download.
* **Subscribe**: Opens the Subscribe dialog enabling you to create a subscription for this saved view. The form in the Subscription dialog is the same as on the **Subscription** tab of the Share dialog, which is covered in **Subscription Tab UI**.
* **Unsubscribe**: If you’re subscribed to this saved view, this option opens a dialog enabling you to confirm that you wish to be removed from the subscription.
* **Preview as Tenant**: Opens a submenu from which you can choose an MKP tenant, after which you'll be taken to the dashboard as it would appear to that tenant.
* **Remove**: Opens the Remove View dialog, where you can confirm that you'd like to remove the dashboard from your organization.

> **Note:** The following actions are not available when the dashboard is a Kentik-provided preset (see **Kentik Presets**): **Edit**, **Settings**, and **Remove**.       

## Dashboard Title Bar

The title bar stretches across the top of the dashboard’s display area, identifying the current dashboard and containing a set of information indicators.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(375).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A30Z&se=2026-05-12T09%3A44%3A30Z&sr=c&sp=r&sig=7H%2FPDw1OZdaRoUuHVx0UmSfJtqwym7OyyJuca3LC5kc%3D)

As shown in the table below, the indicators depend on whether the dashboard was reached from the Library or directly from another dashboard (see **Dashboard Navigation**).

| Element | Description | Reached via Library | Reached via Dashboard Navigation |
| --- | --- | --- | --- |
| Breadcrumbs | Indicates the nesting hierarchy (the path used to get to the dashboard from a source panel on another dashboard). Click the left segment to go to the source panel. | N | Y |
| Visibility icon | Indicates the visibility of the dashboard (see **Dashboard Visibility**):   * *Shared*: Multi-head icon. * *Private*: Single-head icon. | Y | N |
| Label | The label (if any) that is assigned to a dashboard in the Library. | Y | N |
| Trending | A flame icon indicating that the dashboard is among the top 10 most-visited dashboards over the last month. | Y | N |
| Dashboard title | The name given by the user when the dashboard was created or edited (see **General Dashboard Settings**). | Y | Y |

## Dashboards Sidebar

The **Dashboards** sidebar allows you to directly browse, search, and find other dashboards without going back to the **Library** page. The sidebar includes filter controls, tabs, and a list of your organization's dashboards (see **Dashboards Sidebar List**), which you can expand/collapse by clicking the hamburger icon at its top right corner. When your browser width is less than 1760 pixels the sidebar displays as a drawer.

> **Note:** The following supported keyboard shortcuts are listed at the bottom of the sidebar:
>
> * Close: Use the **esc** key to close the sidebar.
> * Switch tabs: Use the left/right arrow keys to toggle between the Favorites, Recents, and Browser tab.
> * Select results: Use the up/down arrow keys to select the previous/next dashboard in the list.

### Dashboards Sidebar Filters

The following controls are used to filter the **Dashboards** list:

* **Search**: A field that filters the **Dashboards** list to dashboards whose title matches the entered text.
* **Label**: A filterable drop-down from which you can select one or more labels. The **Dashboards** list will be filtered to dashboards with the selected labels.
* **Privacy**: A drop-down in which checkboxes will include/exclude dashboards from the **Dashboards** list based on their sharing setting (if none are checked then no sharing settings are excluded):

  + Shared: Dashboards that you or others in your organization have designated for use by everyone in the organization.
  + Private: Dashboards that you own and have designated for your own exclusive use (no one else has access).
  + Kentik Preset: Dashboards created by Kentik.
* **Owner**: A filterable drop-down in which checkboxes will include/exclude owners from the **Dashboards** list (if none are checked then no owners are excluded).
* **Reset To Default**: A link that clears any filters set by the above controls.
* **Sort**: A drop-down to select the sorting criteria for the **Dashboards** list: Relevance, Name, Created Date, or Last Updated.

### Dashboards Sidebar Tabs

The following tabs determine which subset of dashboards is displayed on the **Dashboards** list:

* **Favorites**: A tab to display dashboards marked as favorites, with a numerical count of the search results for tab.
* **Recents**: A tab to display most recently viewed dashboards, with a numerical count of the search results for the tab.
* **Browser** (default): A tab to display all dashboards you have access to, with a numerical count of the search results for the tab.

### Dashboards Sidebar List

The **Dashboards** list contains expandable rows, each representing one dashboard, that contain the following elements:

* **Favorite** (star icon): An icon indicating whether the dashboard is favorited. Select the icon to add or remove the dashboard from **Favorites**.
* **Name**: A link to the dashboard.
* **Labels**: A list of labels applied to the dashboard.
* **Description** (when row is expanded): The optional description of the dashboard.
* **Owner** (when row is expanded): The user who created the dashboard.
* **Expand/Collapse**: Toggles between expanded and compact view.

## Dashboard Query Sidebar

The **Query** sidebar contains dashboard-level controls, which are covered in the following topics.

> **Note:** The sidebar pops up from the **Query** button on the subnav. Clicking anywhere outside of the sidebar will hide it.

### Dashboard Query Panes

The sidebar contains a set of panes that each includes the controls used to set dashboard-wide settings. The following sections are included in the sidebar of a dashboard:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(376).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A30Z&se=2026-05-12T09%3A44%3A30Z&sr=c&sp=r&sig=7H%2FPDw1OZdaRoUuHVx0UmSfJtqwym7OyyJuca3LC5kc%3D)**Query buttons**: The Reset and Apply buttons (see **Applying Sidebar Changes**).
* **Guided Mode UI** (guided mode dashboards only): The prompt and value selector for the guided mode filter (see **Guided Mode UI**).
* **Time pane**: Specifies the dashboard-wide time range. The time range will apply only to panels whose time range is set to **Controlled by Dashboard** (see **Panel Time Range**). For a description of the pane's controls, see **Time Pane**.
* **Filters pane**: Specifies dashboard-wide filters. Apply only to panels whose filtering (see **Panel Filtering**) is set to either **Additive** or **Controlled by Dashboard**. For a description of the pane's controls, see **Filters Pane**.
* **Data Sources pane**: Specifies the dashboard-wide data sources (e.g., devices). The data sources will apply only to panels whose data sources are set to **Controlled by Dashboard** (see **Panel Data Sources**). For a description of the pane's controls, see **Data Sources Pane**.

> **Note:** Default values for the above panes can be specified in the **Query** tab of the Add Dashboard and Dashboard Properties dialogs (see **Default Query Settings**).

### Applying Sidebar Changes

The **Query** sidebar includes two buttons located above the individual panes of the sidebar:

* **Apply button**: Applies changes made in the **Time**, **Filters**, and **Data Sources** panes to any sidebar-controlled panels in the display area (see **Dashboard Panels**):

  + When there are changes to apply, the button will be blue and an alert will appear over the display area.
  + When there are no changes to apply the button will be grayed out.
* **Reset button**: Discards any changes made to the sidebar controls since changes were last applied with the **Apply** button.

> **Note:** Sidebar changes are only applied to panels on which the **Controlled by Dashboard** setting is On in the **Data Sources**, **Time Range**, or **Filters** panes on the **View Panel Query Tab** in the Add View Panel or Edit View Panel dialog.

### Guided Mode UI

A guided-mode dashboard includes additional controls that allow a user to quickly specify the value of a filter that is applied to some or all of the dashboard's panels. These controls, located near the top of the **Query** tab (just below the **Reset** and **Apply** buttons), include the following UI elements:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(485).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A30Z&se=2026-05-12T09%3A44%3A30Z&sr=c&sp=r&sig=7H%2FPDw1OZdaRoUuHVx0UmSfJtqwym7OyyJuca3LC5kc%3D)**Prompt**: The prompt specified with the **Prompt Text** field (see **Guided Mode Settings**) when the dashboard was added or edited.
* **Filter value selector**: Used to choose the value of the dashboard's guided-mode filter. The selector type varies depending on the dashboard's guided-mode input type:

  + Freeform: A field into which the value may be entered.
  + Predefined: A drop-down list from which the value must be chosen.

When the above controls are set, use the **Apply** button to set the current **Filter value selector** setting as the value of the dashboard’s predefined filter dimension (see **Dimension to filter by** in **Guided Mode Settings**). The guided-mode filter applies to all panels on the dashboard except those whose **Guided Mode Behavior** is set to Ignore (see **View Panel Dialog Settings**).

### Guided Mode URLs

In addition to specifying the value of the guided mode parameter manually upon arrival at a guided mode dashboard, you can also append the value to the dashboard's URL, which enables Kentik to apply the value as soon as you navigate to the dashboard.  
  
As shown in the following example (with placeholders in italics), a guided mode URL starts with the path to the dashboard, after which you'll add an endpoint and query string made of a key=value pair representing the parameter's dimension and value (see **Guided Mode Dimension Keys**).

* **Dashboard path**: `https://portal.kentik.com/v4/library/dashboards/dashboard_id`
* **Endpoint**: `/guided`
* **Query string**: `?dimension_key=value`

The combination of the above parts creates a URL like the following:  
`https://portal.kentik.com/v4/library/dashboards/dashboard_id/guided?dimension_key=value`

#### Guided Mode Dimension Keys

The dimension keys for guided mode dashboards actually represent dimension families (which encompass multiple actual dimensions) rather than individual directional dimensions. The following keys are supported:

|  |  |  |
| --- | --- | --- |
| application as\_group as\_name as\_number aws\_acc\_id aws\_priv\_dns aws\_pub\_dns aws\_region aws\_subnet\_id aws\_vm\_id aws\_vm\_name aws\_vm\_type aws\_vpc\_id aws\_zone az\_inst\_id az\_inst\_name az\_region az\_rsrc\_grp az\_sub\_id az\_sub\_name | az\_vnet bgp\_aspath bgp\_community bot\_net cdn city cloud\_provider cloud\_service connectivity\_type country device device\_label dns\_query eth\_mac gce\_proj\_id gce\_region gce\_vm\_name gce\_vpn\_snn gce\_zone interface\_capacity | interface\_desc interface\_group interface\_id interface\_name ip market network\_boundary port provider region service\_name service\_provider service\_type site site\_market tag threat\_list traffic\_org vlan vrf |

## Dashboard Panels

The panel display area for both guided-mode and standard-mode dashboards contains one or more panels.

> **Notes:**
>
> * Dashboards in edit state (see **Dashboard State**) are covered in **Dashboard Configuration**.
> * The UI of panels when a dashboard is in edit state is covered in **Panel Edit Controls**.

### Dashboard Panel Types

An individual panel on a dashboard may be any of the following panel types:

* **View panel**: A panel that originated as a view in **Data Explore**r, which may be in any of the supported **Chart Visualization Types**.
* **Raw Flow panel**: A panel that shows the unaggregated values of one or more dimensions ingested into time-stamped KDE flow records for given set of devices over a given time range (see **Raw Flow Table**).
* **Synthetic Test panel**: A panel that shows results from one or more synthetic tests (see **Synthetic Test Panel**).

### Panel UI Visibility

As detailed in the topics below, the UI elements visible at any given moment on a dashboard panel vary depending on several factors:

* **Panel type**: Whether the panel is a view panel, raw flow panel, or synthetic test panel (see **Dashboard Panel Types**).
* **Dashboard state**: Whether the dashboard itself is in edit state or “view” (non-edit) state.
* **Key filtering**: Whether the panel is in its initial state or has been filtered (see **Filtering With the Key**).
* **Dashboard navigation**: Whether the panel enables drill down to another dashboard (see **Dashboard Navigation**).

### Common Panel UI

The following panel UI elements are visible in both edit and non-edit state:

* **Panel title**: The name given to the panel by the user when the chart is created or edited. The title appears at the top of the panel. In view state it acts as a link that takes you to a view of the panel's query in the Data Explorer.
* **Chart** (view panels only): A visual representation (or tabular, if view type is table, matrix, or raw flow) of the traffic specified in the settings for the panel over the defined timespan on the selected devices. For details on chart types see **Chart Visualization Types**.
* **Key** (view panels only): A set of definitions, below the chart, for the values plotted on that chart. Each series in the chart is represented in the key by a same-colored swatch. See also **Filtering With the Key**.
* **Table**: A tabular display of values (e.g., for flow dimensions or from synthetic test results) over a specified time range.

### View Panel UI

The following panel UI elements appear — on view panels only (see **Dashboard Panel Types**) — when the dashboard is not in edit state:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(381).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A30Z&se=2026-05-12T09%3A44%3A30Z&sr=c&sp=r&sig=7H%2FPDw1OZdaRoUuHVx0UmSfJtqwym7OyyJuca3LC5kc%3D)

*Use a panel's drop-down context menu to see a different visualization.*

* **Context menu**: A drop-down (upper right of panel) with the following options:

  + Change Visualization: Used to select the view type (see **Chart Visualization Types**).
  + Export: Opens a submenu from which you can choose among various forms in which the dashboard can be exported (see **Portal Export Options**).
* **Clear Selections** (appears only when a series is selected; see **Filtering With the Key**): A button that resets the chart so that all series are enabled.
* **Drill Down** (appears only when this panel has been assigned a destination dashboard; see **Dashboard Navigation**): A button that enables navigation to a destination dashboard (see **Navigate To Tab**). The number of series currently enabled on the chart is displayed next to the button.

### Filtering With the Key

In view panels only (see **Dashboard Panel Types**) the key accompanying each panel’s chart can be used to filter what's shown in the chart. The key works as follows (as used below, “label” means either the text identifying a given series or the associated colored swatch):

* The initial state of the chart is that all series are shown. The **Drill Down** button (shown if there’s a destination dashboard) will not be accompanied by a number, and the **Clear Selections** button will be hidden.
* Clicking on a label will "solo" the corresponding series; the plots for all other series will be disabled (not shown in the chart) and the vertical axis is re-scaled. The **Drill Down** button will be accompanied by the number 1, and the **Clear** **Selections** button will be shown.
* Clicking the same label again will restore the chart to its initial state.
* Clicking one or more additional labels will enable their corresponding series, adding them back to the chart. The **Drill Down** button will be accompanied by the number of currently enabled series, and the **Clear Selections** button will still be shown.
* Clicking a label for any enabled series will toggle the series back to a disabled state.
* Clicking the Clear Selections button will restore the chart to its initial state.

> **Note:** If Cross-Panel Filtering (see **View Panel Dialog Settings**) is enabled on a given panel, the filtering applied with a key in that panel will affect the other panels on the dashboard and be indicated in the **Filters** pane of the dashboard's sidebar.

### Raw Flow Table

While view panels (see **Dashboard Panel Types**) may contain any of the chart types described in **Chart Visualization Types**, the display area of a raw flow panel always contains a table. The table, based on settings in the Edit Raw Flow Panel dialog (see **Raw Flow Dialogs**), is made up of the following:

* A column for the timestamp of the flow that is represented by each row.
* A column for every dimension chosen from the **Flow Fields** list.
* The number of rows specified with the **Row Count** field.

You can click on any heading to sort the table ascending or descending based on that column.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(382).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A30Z&se=2026-05-12T09%3A44%3A30Z&sr=c&sp=r&sig=7H%2FPDw1OZdaRoUuHVx0UmSfJtqwym7OyyJuca3LC5kc%3D)

*A raw flow view panel contains a table of flow records.*

### Synthetic Test Panel

A synthetic test panel shows results from one or more synthetic tests created in the Kentik portal's Synthetics section. An individual panel may be configured to display either of the following:

* **Single test**: A single test defined on the **Test Settings** page in Synthetics, the display of which will include a health timeline (see **Synthetic Test Health**) and either a table (see **Test Details Table**), a density grid (see **Density Grid**), or a mesh (see **Test Mesh Subtab**).
* **Multiple tests**: Aggregated results from multiple tests of the same test type (see **Synthetics Test Types**).

> **Notes:**
>
> * To learn more about synthetic tests, see **Synthetics**.
> * For details on configuring dashboard panels to display the results of synthetic tests, see **Synthetic Test Dialogs**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(383).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A30Z&se=2026-05-12T09%3A44%3A30Z&sr=c&sp=r&sig=7H%2FPDw1OZdaRoUuHVx0UmSfJtqwym7OyyJuca3LC5kc%3D)

#### Synthetic Test Health

A health timeline, available for Single tests, is a visual representation of the currently selected time range that shows the overall status of the test for each time slice in the range (see **Status Time and Scope**). By default the latest time slice (far right) is selected, meaning that this is the slice for which details are given in the details table (see below). Hovering anywhere in the timeline changes the selection to the time slice at that point, causing a corresponding change to the information in the selected display type and opening a popup over the timeline that gives the start time and overall status of the selected slice.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(384).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A30Z&se=2026-05-12T09%3A44%3A30Z&sr=c&sp=r&sig=7H%2FPDw1OZdaRoUuHVx0UmSfJtqwym7OyyJuca3LC5kc%3D)

*A Synthetics Test View panel with health timeline across the top.*

#### Availability Test

An **Availability** panel can be created for Synthetics test panels whose **Panel Type** is set to Multiple Test (see **Multiple Test Controls**). This panel shows how consistently a test has been able to complete as expected over the time range specified in the panel settings (see **Synthetic Test Dialog UI**), which is indicated just above the panel's table. The table itself is made up of two kinds of rows:

* **Heading row**: A divider that marks the start of a groups of tests that have all been assigned the same label. Heading rows include the following columns:

  + Group Availability: A color-coded lozenge showing, as a percentage, the average availability of all tests in this group.
  + Label: The label by which the tests are grouped, which is the panel's Test Label setting (see **Multiple Test Controls**).
* **Test rows**: An individual test, with the following columns:

  + Availability: A color-coded lozenge showing the availability of this test as a percentage (healthy runs divided by total runs).
  + Test: The first line gives the test type (e.g., IP Addresses, see **Synthetics Test Types**) and the instance tested (e.g., an IP address). The second line gives the name of the test. Click on the name to go to the test's **Test Details Page**.

The following table shows how the color-coding of lozenges correlates with availability percentages:

| Status | Color | Availability |
| --- | --- | --- |
| Excellent | Green | 99-100% |
| Fair | Orange | 95-98.99% |
| Poor | Red | 0-95% |

The subnav is the gray, horizontal bar or strip located below the main navigation bar or navbar in various pages and modules. It typically contains page-wide controls and buttons specific to the page's functionality.
