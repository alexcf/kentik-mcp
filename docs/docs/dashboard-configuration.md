# Dashboard Configuration

Source: https://kb.kentik.com/docs/dashboard-configuration

---

This article covers creating and editing dashboards in the Kentik portal, including the settings and properties dialogs.

> **Note:** For general information on dashboards, including their layout and controls, see **Dashboards**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(385).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)

*When a dashboard is in edit state its layout can be changed and panels can be added and edited.*

Because a dashboard is primarily made up of a set of individual panels, dashboard configuration involves settings at the dashboard level and the panel level.

The table below shows the dialogs used for dashboard and panel settings in different situations, and how to access those dialogs:

| Level | Type | Function | Dialog | Access via... |
| --- | --- | --- | --- | --- |
| Dashboard | N.A. | Add | Add Dashboard | * Dashboard on **Add New** menu in Library (see **Library Page**). * **Add to New Dashboard** option on **Actions** menu in Data Explorer (see **Add View to Dashboard**). |
| Dashboard | N.A. | Edit | Edit Dashboard | **Settings** option on the **Actions** Menu in a dashboard (see **Dashboard Actions**). |
| Panel | View | Add | Add Data Explorer Panel | * **Add Data Explorer View** button on dashboard in edit state (see **Dashboard Edit Controls**). * **Add to New Dashboard** or **Add to Existing Dashboard** on **Actions** menu in Data Explorer (see **Add View to Dashboard**). |
| Panel | View | Edit | Edit Data Explorer Panel | **Edit** button on panel in edit state (see **Panel Edit Controls**). |
| Panel | View | Add | Add Metrics Explorer Panel | * **Add Metrics Explorer View** button on dashboard in edit state (see **Dashboard Edit Controls**). * **Add to New Dashboard** or **Add to Existing Dashboard** on **Actions** menu in Metrics Explorer (see **Add View to Dashboard**). |
| Panel | View | Edit | Edit Metrics Explorer Panel | **Edit** button on panel in edit state (see **Panel Edit Controls**). |
| Panel | Raw Flow | Add | Add Raw Flow Panel | **Add Raw Flow View** button on dashboard in edit state (see **Dashboard Edit Controls**). |
| Panel | Raw Flow | Edit | Edit Raw Flow Panel | **Edit** button on panel (in edit state). |
| Panel | Synthetics | Add | Add Synthetic Test Panel | **Add Synthetics View** button on dashboard in edit state (see **Dashboard Edit Controls**). |
| Panel | Synthetics | Edit | Edit Synthetics Panel | **Edit** button on panel (in edit state). |

## Dashboard Edit Controls

When you choose **Edit** from the Actions drop-down (not available on Kentik presets), the **Edit Dashboard** pane is displayed below the dashboard title bar. It contains the following controls:

* **Add Data Explorer View**: A button that adds a panel based on a view in **Data****Explorer**. See the steps described in **Adding Dashboard Panels**.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(387).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)
* **Add Metrics Explorer View**: A button that opens the Add Metrics Explorer Panel dialog, allowing you to view and edit the settings described in **Synthetic Test Dialogs**.
* **Add Raw Flow View**:  A raw flow panel enables you to see the unaggregated values of one or more dimensions ingested into time-stamped KDE flow records for given set of devices over a given time range (see **Raw Flow Table**).
* **Add Synthetic Test View**: A button that opens the Add Synthetic Test Panel dialog, allowing you to view and edit the settings described in **Synthetic Test Dialogs**.

* **Dashboard Properties**: A button that opens the Dashboard Properties dialog (see **Dashboard Dialogs**), allowing you to view and edit the properties of the dashboard.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(388).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)
* **Remove**: A button that opens a confirming dialog in which you can either confirm or cancel removal of the dashboard.
* **Save**: A button that saves the current settings of the dashboard.
* **Cancel**: A button that exits edit state without saving any changes to the dashboard.

## Dashboard Dialogs

> **Note:** For panel-level settings, see **Panel Dialogs**.

Dashboard dialogs are used to collect and display the settings needed for dashboard configuration. These fields include general settings, default query options, guided dashboard options, and dashboard navigation settings covered in the topics below.  
  
The information required for these settings is entered into the fields of the following dialogs:

* Add Dashboard when creating a new dashboard. The contents of this dialog vary depending on where you are when you add the dashboard:

  + **Library**: Choose **Dashboard** from the popup that opens from the **Add New** button on the subnav (see **Library Page**).
  + **Data Explorer**: Save the current **Data Explorer** view to a panel on a new dashboard (see **Add Dashboard from Explorer**).
* Edit Dashboard when editing the properties of an existing dashboard. The dialog opens from the **Edit Properties** option in the Action menu on dashboard's row of the **Library** list (see **Dashboard Action Menu Items**).

### Dashboard Dialogs UI

The Add Dashboard and Edit Dashboard dialogs share the same layout and the following common UI elements:

* **Cancel**: A control — either the **X** at the upper right or the **Cancel** button at the lower right — that exits the dialog while leaving all settings as they were when it was opened.
* **Add Dashboard** (Add Dashboard dialog only): A button to save the settings for the new dashboard and exit the dialog.
* **Save** (Edit Dashboard dialog only): A button to save the changes to the dashboard settings and exit the dialog.

### General Dashboard Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(391).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)General properties are set with the following fields:

* **Name**: A field for the name that will represent this dashboard in the **Library List** and in the dashboard's title bar.
* **Description**: An optional field for the description of the dashboard that will be displayed with the dashboard’s name in the **Library List**.
* **Labels**: A set of controls to add and apply labels to the dashboard (see **Labels Controls**).
* **Visibility**: A set of radio buttons to set the **Dashboard Visibility**.

> **Note:** Modifications made to any shared dashboard will affect that dashboard across all users in your organization.

#### Labels Controls

The **Labels** controls contains the following UI elements:

* **Add Label**: An action link to open a **New Label Dialog** to create a new label.
* **Labels**: A field that shows all labels that are already applied to the dashboard.

  + To remove an applied label, click the **X** in that label.
  + To apply a label, click in the field to pop up a filterable list of labels, then click on any listed label to add it to the dashboard.
  + Multiple labels may be applied.

#### New Label Dialog

The New Label dialog is used to create a new label and contains the following UI elements:

* **Color**: A button that opens a popup to select a color for the new label.
* **Name**: A field to input the name for a label.
* **Cancel**: A button that exits the dialog without creating a new label.
* **Save**: A button that saves the newly created label and exits the dialog.

### Default Query Settings

Default query options, which can be overridden in the panes of the Dashboard sidebar (see **Dashboard Query Panes**), are specified with the following fields:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(392).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)**Time Range**: Sets the duration of the "show last" (counting back from present) time range to which the panels on the dashboard will default.
* **Data Sources**: Selects the data sources (e.g., data sources) to which the panels on the dashboard will default (the default data sources can be overridden on individual panels). The currently selected data sources will be listed. To change the default data sources for this dashboard, click the **Edit****Data****Sources** button to open the **Data Sources Dialog**.
* **Filters**: Sets the filters to be applied by default to the panels on the dashboard (the default filters can be overridden on individual panels). Each currently applied filter group will be listed as a card in this section. To change the default filters for this dashboard, click the **Edit Filters** button to open the **Filters Dialog**.

### Guided Mode Pane

Guided dashboards are dashboards that enable users, guided by a prompt at top, to specify (without getting deep into filtering settings) the value of a filter that can be quickly applied to some or all of the dashboard’s panels.  
  
 Guided mode settings are specified with the following UI elements:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(394).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)**Enabled/Disabled**: A switch to enable or disable guided mode. Standard mode is the default (switch disabled).
* **Filter dimension family**: A drop-down selector from which to choose the category of the dimensions that will be available for filters applied with the **Guided Mode UI**.
* **Prompt text**: A field for the question to use as a prompt at the top of the guided dashboard. The field will be populated with a generic prompt when you choose the dimensions family. If desired, click in the field to enter a different prompt.
* **Input type**: A drop-down selector from which to choose the type of values that the user will be able to enter as a value for the filter:

  + **Predefined**: The user will be able to choose values from a drop-down list.
  + **Freeform**: The user will be able to enter any text.
* **Predefined options** (only shown when **Input Type** is "predefined"): A field for a comma-delimited list of values, intended for the values of the drop-down options (see **Guided Mode UI**).

> **Notes:**
>
> * The way that individual panels in the dashboard are affected by the guided dashboard settings above is determined by settings in the Add Data Explorer Panel or Edit Data Explorer Panel dialog (see **View Panel Dialog Settings**).
> * If the value of the guided mode parameter is appended to a dashboard's URL then Kentik will apply that value automatically upon arrival at the dashboard (see **Guided Mode URLs**).

### Dashboard Dependencies Pane

The **Dashboard Dependencies** pane appears only in the Edit Dashboard dialog. It displays the **Dashboard Navigation**, if any, that has been established for this dashboard:

* **Navigate From**: Dashboards on which this dashboard is the destination specified for dashboard navigation from any panel.
* **Navigate To**: The destination dashboards specified for any panels on this dashboard (see **Navigate To Tab**).

### Add Dashboard from Explorer

A variant of the Add Dashboard dialog appears when creating a new dashboard directly from **Data Explorer** (see **Add Panel From Explorer**). This dialog is used only when saving the current **Data Explorer** view to a panel on a new dashboard. The settings in this dialog are generally the same as in the Add Dashboard dialog (see **Dashboard Dialogs**) but with the following differences:

* This dialog includes a **Dashboard Panel** pane (upper right) with a single Title field into which you enter the name of the panel that will appear on the new dashboard.
* This dialog does not include a **Dashboard Dependencies** pane.

## Panel Edit Controls

The following UI elements, which appear only when a dashboard is in edit state, are used to edit panels:

![Diagram showing WAN client locations with percentages for various global sites.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DC-panel-edit-mode.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)

*A panel in edit mode.*

* **Edit button**: Opens the Edit View Panel dialog (see **Panel Dialogs**), which is used to configure the settings of the panel.
* **Clone to**: A button that opens a pop-up menu from which you can choose where you want the cloned panel to live (current dashboard, existing dashboard, or new dashboard), after which the Add View Panel dialog will open, enabling you to configure the panel (see **View Panel Dialog Settings**). The menu options are:

  + **This Dashboard**: After saving the settings of the resulting Add View Panel dialog the new panel is added to the current dashboard.
  + **Existing Dashboard**: The resulting Add View Panel dialog includes a Dashboard pane in which to choose the dashboard to which the new panel will be added.
  + **New Dashboard**: After saving the settings of the resulting Add View Panel dialog, the Add Dashboard dialog opens, enabling you to configure the new dashboard (see **Dashboard Dialogs**).
* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(404).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)**Remove button**: Remove the panel from this dashboard.
* **Filter** (funnel icon; view panels only): Opens a popup that shows the filters applied to the panel and where those filters originate (dashboard, panel, or both).
* **Resize handle**: A handle at the lower right of the panel that you can click and drag to resize the panel.

> **Note:** For information on panel UI when a dashboard is not in edit state, see **Dashboard Panels**.     

## Panel Dialogs

Panel dialogs are used for the settings of individual panels. The layout and fields of each dialog depends on the type of panel (see **Dashboard Panel Types**) and whether you are editing an existing panel or adding a new one:

* **View Panel Dialogs**:

  + **Add Data Explorer Panel**: Used to configure a new view panel on the current dashboard (see **Adding Dashboard Panels**). To open, click the **Add Data Explorer View** button in the **Edit Dashboard** pane (see **Dashboard Edit Controls**).
  + **Edit Data Explorer Panel**: Used to edit the properties of an existing panel. To open, click the Edit button in the title bar of the panel (see **Editing Dashboard Panels**). The Edit button is only visible when the dashboard itself is in edit state (see **Actions** in **Dashboard Subnav Controls**).
* **Metrics Panel Dialogs**

  + **Add Metrics Explorer Panel**: Used to configure a new metrics panel on the current dashboard. To open, click the **Add Metrics Explorer View** button in the **Edit Dashboard** pane.
  + **Edit Metrics Explorer Panel**: Used to edit the properties of an existing panel. To open, click the Edit button in the title bar of the panel. The **Edit** button is only visible when the dashboard itself is in edit state.
* **Raw Flow Panel Dialogs**

  + **Add Raw Flow Panel**: Used to configure a new raw flow panel on the current dashboard. To open, click the **Add Raw Flow View** button in the **Edit Dashboard** pane.
  + **Edit Raw Flow Panel**: Used to edit the properties of an existing panel. To open, click the **Edit** button in the title bar of the panel when the dashboard itself is in edit state.
* **Synthetic Test Panel Dialogs**

  + **Add Synthetic Test Panel**: Used to configure a new synthetic test panel on the current dashboard. To open, click the **Add Synthetic Test View** button in the **Edit Dashboard** pane.
  + **Edit Synthetic Test Panel**: Used to edit the properties of an existing panel. To open, click the **Edit** button in the title bar of the panel when the dashboard itself is in edit state.

> **Note:** For dashboard-level settings, see **Dashboard Dialogs**.      

## View Panel Dialogs

The Add Data Explorer Panel and Edit Data Explorer Panel dialogs share the same layout and the following common UI elements:

* **Cancel**: Buttons — a **Cancel** button at lower right and an **X** in the upper right corner — that close the dialog without saving any changes to settings.
* **Add Panel** (Add Data Explorer Panel dialog only): A button that saves settings for the new panel and exits the dialog.
* **Save button** (Edit Data Explorer Panel dialog only): A button that saves changes to panel settings and exits the dialog.

### View Panel Dialog Settings

A view panel is configured or edited using the fields and controls of the Add Data Explorer Panel and Edit Data Explorer Panel dialogs:

* **Title field**: A string that will appear in the title bar of the panel.

  > **Note:** In a panel on a guided dashboard, you can create a title that always reflects the panel’s current filter dimension and/or value. To do so, insert placeholders `{{filter_dimension}}` and/or `{{filter_value}}` into your title where you want the actual dimension/value to be.
* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(406).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)**Description**: A user-specified description that explains the purpose for which the dashboard was created.
* **Cross-Panel Filtering**: A switch that, when enabled, results in any key-based filtering (see **Filtering With the Key**) being applied at the dashboard level so that it affects not only the panel in which it is applied but also all other panels on the dashboard.
* **Dashboard pane** (shown only in Add Data Explorer Panel dialog when cloning to existing dashboard; see **Panel Edit Controls**): A drop-down list from which you choose the dashboard to which the clone panel will be added.
* **Query tab**: Specify the settings of the query whose results will be charted or displayed in the panel. See **View Panel Query Tab**.
* **Guided Mode tab** (only on guided-mode dashboards): Settings for panels on guided-mode dashboards (see **Guided Mode Tab**).
* **Navigate To tab**: Settings for specifying a nested dashboard that will be linked to from the panel's **Navigate To** button (see **Navigate To Tab**).

### View Panel Query Tab

The settings on the **Query** tab define the query whose results will be charted or displayed in the panel. The tab includes the following main elements:

* **Query preview area**: The area to the left of the sidebar where query results are previewed (equivalent to the chart display area in **Data Explorer**). The controls in the preview area are a subset of the controls in **Data Explorer** (see **Explorer Chart Display**).
* **Query sidebar controls**: The controls that determine what will be shown in the dashboard panel (see **Query Tab Sidebar Controls**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(407).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)

*When adding a view panel, the Query tab determines the traffic data represented in the chart.*

#### Query Tab Sidebar Controls

The **Query** tab sidebar includes the following controls:

* **Visualization**: Displays the view type for the panel and any controls associated with that view type (see **Visualization Pane**).
* **Data Sources**: Displays the data sources (e.g., devices) currently selected for the query, and enables access to the dialog for selecting data sources; see **Panel Data Sources**.
* **Dimensions pane**: Displays the group-by dimensions currently selected for the query, and enables access to the dialog for selecting dimensions (see **Dimension Selectors**).
* **Metrics pane**: Displays the primary metric currently selected for the query, and enables access to the dialog for selecting primary and secondary metrics (see **Metrics Dialog UI**).
* **Time Range**: Determines the time range covered by the panel. See **Panel Time Range**.
* **Filters** (on standard-mode dashboard only): The source of the filters applied to the data charted in the panel; see **Panel Filtering**.
* **Bracketing**: Determines the bracketing, if any, to be applied to the panel (see **About Bracketing**). To set bracketing, click the gear icon to open the **Bracketing Options Dialog**.
* **Advanced**: Determines the **Advanced Query Settings** for the panel.

> **Notes:**
>
> * The preview visualization in this tab is automatically updated as changes are made to the above settings.
> * When the **Data Sources**,**Time**, or **Filters** panes are set to **Controlled by Dashboard**, the results in the panel will be affected by the settings of the dashboard's **Query** sidebar (see **Dashboard Query Sidebar**).

#### Panel Data Sources

The **Data Sources** pane lists the data sources that will be covered in the view displayed in the panel. This pane also includes a **Controlled by Dashboard** switch:

* On (default): The data sources for the panel will follow the sources set with the Data Sources pane of the dashboard.
* Off: The data sources represented in the panel are fully independent of the overall dashboard; they will not change when data sources are changed in the dashboard's **Query** sidebar. The pane will also include the **Edit Data Sources** button, which opens the **Data Sources Dialog**.

#### Panel Time Range

The time range that will be covered by the view in the panel is shown in the time-range indicator. This pane also includes a**Controlled by Dashboard** switch:

* On (default): The time range will follow the setting in the Time pane in the dashboard sidebar. The time range indicator will be grayed out.
* Off: The panel's time range will be fully independent of the overall dashboard; it will not change when the time range is changed in the **Time** pane of the dashboard sidebar. The time-range indicator will become a selector; click it to open a drop-down menu from which you can choose the time range.

#### Panel Filtering

The **Filters** pane lists the filter groups, if any, that will be applied to the data charted in the panel (see **About Filters**). Unless **Filter Setting** (see below) is **Controlled by Dashboard**, you can set a filter by clicking the gear icon at the upper right of the pane, which opens the **Filters Dialog**.  
  
The source of the panel-level filters that will be applied to the panel is set with**Filter Setting** drop-down:

* **Controlled by Dashboard**: The filters applied to the panel's chart will be the same as the filters, if any, specified in the **Filters** pane of the dashboard's sidebar (see **Dashboard Query Panes**). The settings button (gear icon) at the upper right will be inactive (grayed out).
* **Additive** (default): The filters applied in the **Filters** pane in this dialog will be ANDed with any filters that are applied in the sidebar of the dashboard itself.
* **Locked to Specific Filters**: The filters applied to the panel's chart are fully independent of the overall dashboard; they will not change when filter settings are changed in the **Filters** pane of the dashboard’s sidebar.

#### Panel Filter Interactions

In addition to panel-level filters, the filters applied to a given panel may include a guided-mode filter (see **Guided Mode Tab**) and dashboard-level filters (see **Dashboard Query Panes**). The tables below show how these filters interact depending on panel-level and guided-mode settings.  
  
Filter interaction when the panel's **Filter Setting** is **Additive**:

| Dashboard Filters? | Guided Mode | Resulting Filters |
| --- | --- | --- |
| No | None | Panel only |
| " | Override All | Guided only |
| " | Override specific | Panel only (with Guided value if applicable) |
| " | Add | Panel + Guided |
| " | Ignore | Panel only |
| Yes | None | Panel + Dashboard |
| " | Override All | Dashboard + Guided |
| " | Override specific | Panel (with Guided value if applicable) + Dashboard |
| " | Add | Panel + Dashboard + Guided |
| " | Ignore | Panel + Dashboard |

Filter interaction when the panel's **Filter Setting** is **Controlled by Dashboard**:

| Dashboard Filters? | Guided Mode | Resulting Filters |
| --- | --- | --- |
| No | None | None |
| " | Override All | Guided only |
| " | Override specific | None |
| " | Add | Guided only |
| " | Ignore | None |
| Yes | None | Dashboard only |
| " | Override All | Dashboard + Guided |
| " | Override specific | Dashboard only |
| " | Add | Dashboard + Guided |
| " | Ignore | Dashboard only |

Filter interaction when the panel's **Filter Setting** is **Locked to Specific Filters**:

| Dashboard Filters? | Guided Mode | Resulting Filters |
| --- | --- | --- |
| No | None | Panel only |
| " | Override All | Guided only |
| " | Override specific | Panel only (with Guided value if applicable) |
| " | Add | Panel + Guided |
| " | Ignore | Panel only |
| Yes | None | Panel only |
| " | Override All | Guided only |
| " | Override specific | Panel only (with Guided value if applicable) |
| " | Add | Panel + Guided |
| " | Ignore | Panel only |

### Guided Mode Tab

The **Guided Mode** tab of the Add Data Explorer Panel and Edit Data Explorer Panel dialogs is included only in the dialogs for panels on guided-mode dashboards.

#### About Guided Mode Tab Settings

Upon arrival at a guided-mode dashboard the user is prompted to set a value for the dimensions in the dashboard's guided-mode dimension family (see **Guided Mode UI**), which makes available a set of guided-mode filters that may be applied to any or all of the dashboard's panels. For example, based on the **Dimension Family** setting in the **Guided Mode** pane of the Dashboard Properties dialog (see **Guided Mode Pane**), the user may be prompted to enter an IP address.

For each individual panel of the dashboard, the settings of the**Guided Mode** tab determine if the panel will include a guided-mode filter whose dimension is from the IP dimension family (Source IP, Destination IP, Source Next Hop IP, etc.) and whose value is the IP entered by the user. If a guided-mode filter is applied to a given panel, the settings in the**Guided Mode** tab also determine how that filter will interact with any panel-level filters (see **Panel Filtering**) or dashboard-level filters (see **Dashboard Query Panes**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(409).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)

*Set how a Guided Mode filter will interact with panel-level or dashboard-level filters.*

#### Guided Mode Behavior

The**Guided Mode**tab includes a **Guided Mode Behavior** pane with a set of radio buttons. These buttons determine the interaction between the guided-mode filter and any existing panel-level filters:

* **Override All Filters**: All existing panel-level filters will be ignored. A new (guided-mode) filter will be created and applied to the panel:

  + You choose the new filter's dimension and operator (see **Guided Mode Dimension**).
  + The filter's value will be entered by the user in response to the guided-mode prompt.
* **Override Specific Filters**: No new filter will be created for this panel. All existing panel-level filters will be applied. If any such filter is based on a dimension in the guided-mode dimension family then:

  + You choose the dimension whose existing value will be overridden by the value supplied by the user.
  + The filter's value will be entered by the user in response to the guided-mode prompt.
* **Add Filter Group**: A new (guided-mode) filter will be created and applied to the panel. The new filter will be ANDed with existing panel-level filters:

  + You choose the new filter's dimension and operator.
  + The filter's value will be entered by the user in response to the guided-mode prompt.
* **Ignore**: No new filter will be created for this panel. The value entered by the user in response to the guided-mode prompt will have no effect on this panel.

The **Panel Filter Interactions** table shows how the above settings interact with panel-level filtering (see **Panel Filtering**) and dashboard-level filtering (see **Dashboard Query Panes**) to determine how panel's results are filtered.

#### Guided Mode Dimension

Depending on the **Guided Mode Behavior** setting, the **Guided Mode** tab made include additional settings used to configure the guided-mode filter:

* **Dimension** (present unless behavior is Ignore): A drop-down list of the dimensions in the current dimension family. Choose a dimension for the guided-mode filter.
* **Operator** (present unless behavior is Ignore or Override Specific Filters): A drop-down list of the operators available for the guided-mode filter.

### Navigate To Tab

The **Navigate To** tab includes the following controls used to configure **Dashboard Navigation**:

* **Enable Dashboard Navigation switch**:

  + If on, dashboard navigation will be enabled for this panel, meaning that the panel will include a **Drill Down** button that links to a different dashboard.
  + If off, the rest of the settings on this tab will be hidden.
* **Destination Dashboard selector** (shown only when dashboard navigation is enabled): Choose the dashboard to which this panel should navigate (see **Destination Dashboard Selector**).

![Dashboard navigation settings for configuring data sources and time range options.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DB-Add_panel-Navigate_tab-647h674w.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)

*Dashboard navigation settings for configuring data sources and time range options*.

#### Destination Dashboard Selector

The **Destination Dashboard** selector is a drop-down menu enabling you to choose a destination dashboard from a list of all of the dashboards in your organization. When you choose a destination, additional UI elements appear:

* **Drill Down Button Text**: A field into which you can enter an alternate name for the Drill Down button on this panel (the button at the lower right of the panel that will navigate to the destination dashboard). The default name is "Drill Down."
* **Destination Dashboard Settings** (shown only when a destination dashboard has been selected): Specify how the data sources, time range, and filters of the destination (nested) dashboard will be set (see **Destination Dashboard Settings**).
* **Guided Mode Settings** (only when your destination dashboard is guided-mode): See **Guided Mode Settings**.

#### Destination Dashboard Settings

The **Destination****Dashboard****Settings** section of the**Navigate****To** tab is used to specify how the data sources, time range, and filters of the destination dashboard will be set. The following table shows the options that are available for those settings:

| Use settings from... | Data Sources | Time range | Filters | Description |
| --- | --- | --- | --- | --- |
| Source dashboard | Y | Y | Y | Use the settings from the sidebar of the dashboard containing this panel. |
| Destination dashboard | Y | Y | Y | Use the settings from the sidebar of the nested dashboard that is linked-to from this panel. |
| Source and Destination combined | N | N | Y | Combine the filters from both the source and destination dashboards. Available for filters only. |
| Source panel | Y | Y | Y | Use the settings from this panel. |
| Custom | Y | Y | Y | When this option is selected, controls will appear that enable you to specify a custom setting. |

#### Guided Mode Settings

If the destination you choose from the **Destination Dashboard Selector** is a guided-mode dashboard (see **Dashboard Modes**) the **Guided Mode Settings** appear to the right of the selector. This set of radio buttons determines how the value of the guided-mode dimension will be set when you navigate to the destination dashboard:

* **Use filter value from this guided dashboard**: Enabled only if the panel for which you are setting a destination dashboard is on a guided-mode dashboard. The guided-mode value of the destination dashboard will be set to the guided-mode value of the source dashboard.

  > **Note:** The guided-mode dimensions of the source and destination dashboards are not validated to ensure that they are the same.
* **Use Drill Down selection**: The time-series currently selected in the source panel's chart (see **Filtering With the Key**) will be used to determine the value of the guided mode dimension.

  > **Note:** In the event of a mismatch between the guided mode dimension and the panel dimension, the **Drill Down** button will do nothing when clicked.
* **Prompt user**: When the user clicks the **Drill Down** button (see **View Panel UI**), a prompt will appear asking for the value to use for the guided-mode dimension. This behavior is identical to what happens when a user navigates to a guided-mode dashboard directly (without coming from a source dashboard).

### Add Data Explorer View Panel

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(412).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)The Add Data Explorer View Panel dialog appears when the **Add to Existing Dashboard** option is chosen from the drop-down **Actions** menu in **Data Explorer**, which saves the current view to a panel on an existing dashboard (see **Add Panel From Explorer**).  
  
The settings in this dialog differ as follows from the fields described in **View Panel Dialog Settings**:

* There is no **Query** tab, **Guided Mode** tab, or **Navigate To** tab.
* There is a **Dashboard** pane that contains a drop-down list from which you choose the existing dashboard to which you'd like to add the new panel.

## Metrics Panel Dialogs

Dashboards include two versions of the settings dialog for a **Metrics Explorer** panel:

* **Add Metrics Explorer Panel**: Opens from the **Add Metrics Explorer View** button.
* **Edit Metrics Explorer Panel**: Opens from the **Edit** button in the title bar of an existing panel.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(413).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)

*Metrics panels use the same query settings as Metrics Explorer.*

The settings dialogs for Metrics Explorer panels share the following main UI elements:

* **General dialog UI**: Controls related to the overall operation of the dialog, which are the same (e.g., **Cancel**, **Save**) as in the dialogs for Data Explorer view panels; see **View Panel Dialogs UI**.
* **Title**: A field for a string that will appear in the title bar of the panel.
* **Description** (optional): A user-specified description that (ideally) explains the purpose for which the dashboard was created.
* **Results display**: An area in which the chart resulting from the current settings in the dialog's **Query** sidebar are displayed.
* **Metrics query sidebar**: A set of panes containing controls that define the metrics query. These panes are described in detail in **Metrics Explorer Query Sidebar**.

## Raw Flow Dialogs

The raw flow panel dialogs include the following main UI elements:

* **General dialog UI**: Controls related to the overall operation of the dialog, which are the same (e.g., **Cancel**, **Save**) as in the dialogs for Data Explorer view panels; see **View Panel Dialogs UI**.
* **Title field**: A string that will appear in the title bar of the panel on the dashboard.
* **Description**: A user-specified description that (ideally) explains the purpose for which the dashboard was created.
* **Query Sidebar**: A set of controls that specify the query that determines which raw flow results will be displayed in the panel (see **Raw Flow Query Settings**).
* **Preview Display**: Shows how the panel will display the results of the raw flow query.

### Raw Flow Query Settings

The raw flow **Query** sidebar includes the following controls:

* **Preview Query**: If this button is present at the top of the sidebar, click it to update the table in the **Preview Display**. If the button is not present the table will update automatically as changes are made to the settings.
* **Data Sources**: Displays the data sources (e.g. devices) currently selected for the query, and enables access to the dialog for selecting data sources; see **Panel Data Sources**.
* **Dimensions pane**: Lists the group-by dimensions currently selected for the query, and enables editing that list. See **Raw Flow Dimensions Pane**.
* **Time Range**: Determines the time range covered by the panel. See **Panel Time Range**.
* **Filters**: This pane lists the filters applied to the raw flow data charted in the panel. the source of those filters is determined by the **Filter Setting** drop-down (see **Panel Filtering**). Unless **Filter Setting** is **Controlled by Dashboard**, you can set a filter by clicking the **Edit Filters** button, which opens the **Filtering Options Dialog**.

> **Note:** When the **Data****Sources**, **Time**, or**Filters** panes are set to **Controlled by Dashboard**, the results in the panel will be affected by the settings of the dashboard's **Query** sidebar (see **Dashboard Query Sidebar**).

#### Raw Flow Dimensions Pane

The **Dimensions** pane lists the dimensions currently selected for inclusion in the query. When adding a raw flow panel you'll see a default set of dimensions that represent fields from flow records:

* To remove a dimension from the list, click the **X** at the right of that dimension.
* To add a dimension, click in the list to open the Group By Dimensions dialog (see **Dimension Selectors**).

> **Note:** If a dashboard was built with dimensions that are no longer supported, those dimensions will be indicated in the dimensions list as "invalid."

In addition to the dimensions list, the **Dimensions** pane contains the following controls:

* **Order by**: Set the initial sort order for the rows of the table. The default is Timestamp (chronological order). To sort by a different dimension, click the indicator to open a dimension selector.

  > **Note:** If not timestamp, the order-by dimension must be currently selected for inclusion in the query.
* **Row Count**: Set the number of rows that will return when the table is displayed in the panel.

## Synthetic Test Dialogs

The dialogs for adding and editing synthetic test panels include the following main UI elements:

* ![Interface for adding a synthetic test panel with various configuration options.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/DC-add-synthetic-test-panel.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)**General dialog UI**: Controls related to the overall operation of the dialog, which are the same (e.g. **Cancel**, **Save**) as in the dialogs for Data Explorer view panels; see **View Panel Dialogs UI**.
* **Title field**: A string that will appear in the title bar of the panel on the dashboard.
* **Panel Type**: A drop-down selector from which you choose the type of panel:

  + Single Test: The information displayed in the panel will represent the results from a single test defined on the **Test Settings Page**, the display of which will be determined by the **Display Type** setting.
  + Multiple Test: The information displayed in the panel will represent the aggregate of results from multiple tests.
* **Single-test controls**: Controls used to configure panels for which **Panel Type** is Single Test (see **Single Test Controls**).
* **Multiple-test controls**: Controls used to configure panels for which **Panel Type** is Multiple Test (see **Multiple Test Controls**).
* **Time Configuration**: A switch that determines whether the time range covered by the panel will be **Controlled by Dashboard** (follow the setting in the Time pane in the dashboard sidebar) or be fully independent of the overall dashboard (see **Panel Time Range**).
* **Time Range**: If Time range is not **Controlled by Dashboard**, this control will be active, enabling you to set the time range covered by the panel (see **Time Range Control**).

### Single Test Controls

In addition to the above, the following controls are active when **Panel Type** is set to Single Test:

* **Display Type**: A drop-down for choosing which element of the Test Details page for the specified test will display in the new panel:

  + Table: The panel will display the test's **Test Details Table**.
  + Density Grid: A set of squares that each represent the target of a test by a given agent (see **Density Grid**).
  + Mesh: The panel will display the matrix from the Test Details page's **Mesh** tab (see **Test Mesh Subtab**).
* **Test Type**: A drop-down list of the tests types that can produce a display of the type selected with **Display Type**.
* **Select Test**: A drop-down list of individual tests in your organization of the type selected with **Test Type**. Click to choose a test.

#### Density Grid

A density grid shows the health of the connections between a given agent and the targets (e.g., IP addresses or DNS servers) to which it tests. Each target is represented as a square whose color indicates health status (Healthy, Warning, or Critical) of the subtest toward that target. Hovering over a square opens a popup with additional information (IP address, resolution time) and a **View Details** button that will take you to the Subtest Metrics Tab for the corresponding subtest.  
  
A density grid may be presented in different layouts depending on how many agents are covered by the test that's being shown in the dashboard panel:

* If the number of agents can be accommodated horizontally across the panel then each agent will be presented in a traditional tab that shows all of the targets (squares) for that agent.
* If the test involves a large number of agents then the agents will be listed vertically along the left side of the panel. By default the display area to the right of this list will show the targets for the top-most agent in the list. Click on a different agent to see a different agent's targets.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(417).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A23Z&se=2026-05-12T10%3A03%3A23Z&sr=c&sp=r&sig=uQJFoDyAZGyQAFQZ9U87P5KVEnr%2FOXAWOMJ40A6WdFc%3D)

*A density grid for one agent on one tab of a Synthetics panel for a test with two agents.*

### Multiple Test Controls

In addition to the above, the following controls may be active (depending on the interaction of current settings) when **Panel Type** is set to Multiple Tests:

* **Display Type**: A drop-down for choosing what will display in the new panel (see **Multiple Test Display Types**).
* **Test Type**: A drop-down list of the test types that can produce a display of the type selected with **Display Type**.
* **Metric Type** (not present if Display Type is Density Grid Group): A multi-select box that you click to choose the metrics (e.g., latency) whose values are displayed in the results.

  > **Note:** The available metrics vary by test type.
* **Agent Labels** (only if Display Type is Min/Max/Avg - By Label): A multi-select box that you click to choose the labels whose corresponding agents will be included in the results.

  > **Note:** Labels can be applied to agents to group them by useful attributes (e.g., by region).
* **Select Test** (not present if Display Type is Min/Max/Avg - By Label): A multi-select box that lists individual tests in your organization of the type selected with Test Type. Click to choose one or more tests.
* **Test Label** (only if Display Type is Availability):A drop-down list to choose the labels whose corresponding tests will be displayed in the results.

### Multiple Test Display Types

The following display types are supported when**Panel Type** is Multiple Tests:

* **Min/Max/Avg - By Label**: The panel will display a table showing the Min, Max, and Avg values over the time range for the metrics chosen with **Metric Type** and the agents that have been assigned the **Labels** chosen with **Agent Labels**.
* **Gauge - Min**: The panel will display a gauge (see **Gauge Visualization**) showing the aggregate minimum value, across the tests included with Select Tests, for the metric specified with **Metric Type**.
* **Gauge - Max**: The panel will display a gauge showing the aggregate maximum value, across the tests included with **Select Tests**, for the metric specified with **Metric Type**.
* **Gauge - Avg**: The panel will display a gauge showing the aggregate average value, across the tests included with **Select Tests**, for the metric specified with **Metric Type**.
* **Density Grid Group**: The panel will display a table in which the rows each represent a synthetics agent that is involved in a test specified with **Select Tests**, and the columns each represent those tests. The cells of the table are populated with squares whose color indicates health status (Healthy, Warning, or Critical). Hovering over a square opens a popup with additional information (IP address, resolution time) and a View Details button that will take you to the **Subtest Metrics Tab** for that agent in that test.
* **Availability**: The panel will display a table in which each row shows the availability of one test, which indicates how consistently the test is able to complete successfully (see **Availability Test**). The rows for all tests with the same label will be grouped under header rows.

## Manage Dashboards

Procedures for managing your organizations dashboards are covered in the following topics.

### Add a Dashboard

Creating a dashboard involves adding a new dashboard to your dashboards collection, after which you would typically add and configure at least one panel for that dashboard (see **Adding Dashboard Panels**). To add a new dashboard:

1. Choose **Library** from the portal's navbar menu.
2. On the **Library Page**, click **Dashboard** on the menu that drops down from the **Add New** button in the subnav. The Add Dashboard dialog will open.
3. In the dialog, specify the settings of the dashboard (see **Dashboard Dialogs**), including the dashboard's labels and whether guided mode is enabled or disabled.
4. Click the **Add Dashboard** button. The dialog will close and the dashboard will be listed in the **Library** list.

> **Notes:**
>
> * For information on populating the new dashboard with panels, see **Adding Dashboard Panels**.
> * You can also add a dashboard from Data Explorer; see **Add Dashboard from Explorer**.

### Edit a Dashboard

The details of editing a dashboard vary depending on which aspect of the dashboard you're trying to change.

#### Edit Dashboard Properties

To change a dashboard's properties:

1. Choose **Library** from the portal's navbar menu.
2. In the **Library** list, click **Edit Properties** in the Action menu at the right of the dashboard's row. The Dashboard Properties dialog will open (see **Dashboard Dialogs**).
3. Change the desired settings.
4. Click the **Save** button to save your changes and exit the dialog.

> **Note:** You can also access a dashboard's settings from within the dashboard itself by clicking **Settings** on the dashboard's Actions menu (see **Dashboard Subnav Controls**).

#### Edit Dashboard Contents

To edit a dashboard from the Library:

1. Choose **Library** from the portal's navbar menu.
2. In the **Library** list, click **Open in Edit mode** in the Action menu at the right of the dashboard's row. The dashboard will open in edit mode (see **Dashboard Edit Controls**).
3. Make your desired changes:

   1. To add one or more panels, see **Adding Dashboard Panels**.
   2. To edit a panel, see **Editing Dashboard Panels**.
   3. To change the size or layout of panels on the dashboard, see **Change Dashboard Layout**.
4. Click the **Save** button to save your changes and exit edit mode.

> **Note:** Modifications made to any shared dashboard (see **Dashboard Visibility**) will affect that dashboard across all users in your organization.

#### Change Dashboard Layout

To move and/or resize the panels of a dashboard:

1. Choose **Library** from the portal's navbar menu.
2. In the **Library** list, click **Open in Edit mode** in the Action menu at the right of the dashboard's row. The dashboard will open in edit mode (see **Dashboard Edit Controls**).
3. With the dashboard in edit mode:

   1. Move a panel: Position the cursor on the desired panel. The move cursor will appear. Drag the panel to the desired new location. When the dragged panel is dropped, the other panels will rearrange around it.
   2. Resize a panel: A resize handle will appear at the bottom right of each panel. Drag the handle to change the size of the panel. Surrounding panels may be rearranged to accommodate the new size.
4. Click the **Save** button to save the new layout and exit edit mode.

### Remove a Dashboard

To remove a dashboard from your organization's collection of dashboards:

1. Choose **Library** from the portal's navbar menu.
2. In the **Library** list, click **Remove** on the Action menu at the right of the dashboard's row.
3. In the resulting confirmation popup, click **Remove** to delete the dashboard.

> **Note:** You can also remove a dashboard from the dashboard itself by clicking **Remove** on the dashboard's Actions menu (see **Dashboard Subnav Controls**).

## Add or Edit Panels

As described in **About Dashboards**, dashboards are made up of one or more individual panels that each contain a single chart that represents a query.

> **Note:** The elements that make up each panel are described in **Dashboard Panels**.

### Adding Dashboard Panels

To add a panel to the current dashboard:

1. In the **Dashboard Subnav Controls**, choose **Edit** from the Actions drop-down. The **Edit Dashboard** pane (see **Dashboard Edit Controls**) will appear below the title bar.
2. Click one of the following:

   1. **Add Data Explorer View** (button): The Add Data Explorer Panel dialog will open. Configure the settings described in **View Panel Dialog Settings**.
   2. **Add Raw Flow** (button): The Add Raw Flow Panel dialog will open. Configure the settings described in **Raw Flow Dialogs**.
   3. **Add Synthetic Test** (button): The Add Synthetic Test Panel dialog will open. Configure the settings described in **Synthetic Test Dialogs**.
3. When the dialog settings are complete, click the **Add Panel** button to save the settings and close the dialog. Back on the dashboard, the page will auto-scroll to the new panel, which will appear at the bottom, below the existing panels.
4. In the dashboard title bar, click the **Save** button to exit.

> **Notes:**
>
> * To add additional panels, repeat the above.
> * Panels added to any shared dashboard (see **Dashboard Visibility**) will affect that dashboard across all users in your organization.

### Add Panel From Explorer

To add a dashboard panel directly from **Data Explorer**:

1. In the subnav (see **Explorer Subnav Controls**), click the drop-down Actions menu.
2. Choose one of the following:

   1. **Add To New Dashboard**: Opens the dialog described in **Add Dashboard from Explorer**. Specify settings for the new dashboard, then click the **Add** **Dashboard** button. The new dashboard will open and will display the new panel.
   2. **Add To Existing Dashboard**: Opens the **Add Data Explorer View Panel** dialog. In the **Dashboard** pane, choose the existing dashboard to which you want to add the new panel, then click the **Add** **Panel** button. The dashboard will open with the panel that you just added, and a notification will confirm that the new panel was created.

### Editing Dashboard Panels

The settings of an individual dashboard view panel are edited in the Edit Data Explorer Panel dialog. To edit a panel on the current dashboard:

1. In the **Dashboard Subnav Controls**, choose **Edit** from the Actions drop-down, which places the dashboard in edit state.
2. The title bar of each individual panel will now include an **Edit** button. Click that button in the panel that you want to edit.
3. In the resulting Edit Panel dialog, make your changes to the panel settings referenced in step 2 of **Adding Dashboard Panels**.
4. Click the **Save** button to save your changes and return to the dashboard.

> **Note:** Modifications made to a panel on any shared dashboard (see **Dashboard Visibility**) will affect that dashboard across all users in your organization.  

## Add Dashboard Navigation

**Dashboard Navigation** is set on the dashboard panel from which you will be navigating, which must be a view panel (see **Dashboard Panel Types**). In the following example, we’ll set an existing dashboard (the "destination dashboard") as the destination of an existing panel (the "source panel") on the current dashboard. The process would be similar when using a new dashboard or panel.

1. When the dashboard opens, choose **Edit** from the Actions drop-down (see **Dashboard Subnav Controls**). The title bar of each individual panel will now include an **Edit** button.
2. Click the **Edit** button in the panel on which you want to establish navigation to a different dashboard. The Edit Data Explorer Panel dialog will open (see **Panel Dialogs**).
3. Use the controls of the **Navigate To Tab** to specify the destination dashboard (see **Destination Dashboard Selector**) and set the **Destination Dashboard Settings**.
4. Save the changes you've made to the Edit Data Explorer Panel dialog.
5. Back in the dashboard title bar, click the **Save** button to exit edit state. The panel to which you added navigation will now have a **Drill Down** button at lower right. When you hover over the button, the tooltip will now show the name of the nested dashboard. Click the button to go to that dashboard.
