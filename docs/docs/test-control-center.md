# Test Control Center

Source: https://kb.kentik.com/docs/test-control-center

---

Use Kentik’s **Test Control Center** (Synthetics) to both configure tests, as covered in **Synthetics Test Settings**, and to view test results.

> **Notes:**
>
> * For information about the types of tests that can be managed in the Test Control Center, see **Synthetics Test Types**.
> * For information about how test status (Healthy, Warning, Critical) is derived from test results, see **Synthetics Test Status**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(503).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*The Test Control Center landing page.*

## About Test Control Center

The **Test Control Center** is used to manage your organization's synthetics tests. Cards across the top of the page show your current (month to date) and projected (full month, at current trends) consumption of test credits (purchased from Kentik in SyntheticPaks, see **Synthetics Test Credits**). A **Tests List**, filterable by name, type, label, and by agent, provides information about your currently configured tests and gives you access to an edit page for any existing test. You can also use the **Add Test** button to go to the **Add Test Page**, where you can choose and configure a new test.  
  
The Test Control Center is structured as follows to allow successive drill-down into the details of each test and its component subtests (see **Tests and Subtests**):

* **Test Control Center page**: Provides an overview of all tests (see **Test Control Center Page**).
* **Test Details page**: Provides an overview of a single test including information and status for all of the test's subtests (see **Test Details Page**).
* **Subtest Details page**: Provides information and status details for an individual subtest (see **Subtest Details Page**).

> **Note:** If your organization's test credit balance (see **About Test Credits**) is below zero, a notification will appear to the right of the title on the main **Test Control Center Page** as well as on any **Add Test Page** or **Test Settings Page**. To add more credits, click the link in the notification (see **Adding Test Credits**).

## Test Control Center Page

The Test Control Center page is the home page of the test control workflow. The page includes the following main UI elements:

* **Share** (on subnav): Opens the Share dialog to enable you to share the current view (see **Sharing via the Share Dialog**).
* **Actions** (on subnav): A drop-down menu from which you can choose to export the current view as a visual report (PDF) covering the page’s visualizations and tables. A notification appears when the PDF is ready to download.
* **Add Test**: A button that takes you to the **Add Test Page**.
* **Credits notification**: A notification — visible only when your organization has run out of **Synthetics Test Credits** — that alerts you to that fact. The notification includes a **Submit a Request** button, which opens a version of the **Contact Support** popup (see **Support Request**) that is pre-populated with a request for a quote for additional credits.
* **Credit Utilization**: Two cards that indicate your organization's current (this month) and projected (next month) usage of test credits (see **Credit Utilization**).
* **Tests list**: A list of the currently configured tests in your organization (see **Tests List**).
* **Filters pane**: A set of fields and checkboxes used to filter the Tests List (see **Tests List Filters**).

#### Credit Utilization

The **Credit Utilization** section is made up of two cards that indicate your organization's current (this month) and projected (next month) usage of test credits (see **About Test Credits**).

* **Current Credit Utilization**: A card that indicates your organization's current (month to date) usage of test credits:

  + Monthly Allocation: The total credits that are available to your organization this month (e.g., the credits-per-month purchased from Kentik in a SyntheticPak).
  + Used: The credits your organization has actually used so far this month.
  + Projected Total: The credits Kentik expects your organization to use during this entire month based on the credits consumed by the tests that are currently active.
  + Projected Remaining*:* The credits from your organization's allocation for this month that will go unused if your actual usage matches your projected usage.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(504).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

* **Next Month Projection**: A card that indicates your Kentik-projected consumption of credits for the month following the current month:

  + Projected Total: The credits Kentik expects your organization to use during the entire next month based on the credits consumed by the tests that are currently active.
  + Projected Remaining: The credits from your organization's allocation for next month that will go unused if your actual usage matches your projected usage.

  ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(505).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

## Tests List

The **Tests** list on the Test Control Center page, which provides information about the synthetic tests currently configured in your organization, is covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(506).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*The Tests List is a list of all tests currently configured in your organization.*

### Label Controls

Kentik's labeling feature enables you to create a label (essentially a property whose value is text) and apply it to one or more of your synthetic tests, creating a group of tests that can be referred to (e.g., filtered for) collectively rather than individually. The Test Control Center includes the following label-related controls for tests, which are associated with the tests list:

* **Selection indicator**: Indicates how many tests are currently selected in the tests list (see the Select checkbox in **Tests List Columns**).
* **Add/Edit Labels**: Links to the settings page for **Labels**, where you can create or remove the labels that are available to apply to the selected tests (as well as to data sources such as routers).
* **Clear Labels**: A button that clears the labels that are applied to all currently selected tests in the list.
* **Apply Labels**: A drop-down list of labels from which you can choose a label to apply to all tests that are currently selected in the tests list.
* **Select All** (checkbox): A checkbox for toggling the selection state of all tests in the tests list:

  + If either no checkboxes in the list itself are checked or only some are checked then clicking this checkbox will select all tests.
  + If all checkboxes in the list are checked, clicking this checkbox will deselect all tests.

> **Note:** Labels may also be applied to and removed from tests in the **Test Information** tab of the Test Settings page.

### Tests List Columns

The **Tests** list provides the following types of information for each listed test:

* **Select** (checkbox): Select this test to clear or apply labels (see **Label Controls**).
* **Status**: The status of the test and its results (see **Synthetics Test Status**).
* **Test**: The type (see **Synthetics Test Types**) and name of the test.
* **Created by**: The user in your organization who created the test, as well as the date created.
* **Last updated**: The user in your organization who most recently modified the test, as well as the date modified.
* **Agents**: The number of agents (see **About Synthetics Agents**) involved in the test.
* **Monthly Credits**: The projected monthly consumption of test credits (see **About Test Credits**) by this test.
* **ID**: A Kentik-assigned ID for the test.
* **Pause/Resume** (icon): Suspend or resume this test.
* **Edit** (icon): Takes you to the **Test Settings Page** for this test.
* **Copy** (icon): Duplicate the test so that you can make a new test by modifying (editing) the duplicate rather than starting from scratch.
* **Remove** (trash icon): Opens a confirming popup from which you can remove the test from your organization's collection of tests.

### Tests List Filters

The **Tests** list can be filtered using the **Filters** pane at the right. The pane includes the following settings:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(507).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

* **Clear all**: Clear all currently applied filters.
* **Search**: A field into which you can enter text to filter the Tests list to tests whose name contains the entered string or whose ID is an exact match (partial matches are excluded).
* **Type**: A list of test types. Check one or more types to see only tests of those types (when no types are checked, all tests are listed). Only test types that are currently used are displayed.
* **Labels**: A field that displays a lozenge for each selected label and filters the Tests list to tests with those labels. To select a label, click in the field and choose the label from the filterable drop-down list. To remove a label, click the **X** at the right of its lozenge.
* **Agents**: A field that displays a lozenge for each selected agent and filters the Tests list to tests involving those agents. To select an agent, click in the field and choose the agent from the filterable drop-down list. To remove an agent, click the **X** at the right of its lozenge.
* **Credentials**: A field that displays a lozenge for each selected credential and filters the Tests list to tests using those credentials. To select a credential, click in the field and choose the credential from the filterable drop-down list. Only credentials in current use are listed. To remove a credential, click the **X** at the right of its lozenge.

  > **Note:** For information about the listed credentials, choose **Credentials Vault** from the main navbar's **Organization Settings** menu to go to the Credentials Vault page, where the credentials will be shown in the **Credential List**.

> **Note:**
>
> The filters applied via the **Filters** pane interact as follows:
>
> * Filters are ORed within each of the categories (Name, Type, Label, Agent), so tests matching any filter value within an individual category are considered a match for that category.
> * Filters are ANDed between categories, so tests must match in all specified filter categories to be included in the **Tests Lis****t**.

## Test Details Page

The Test Details page for a given test, which is accessed by clicking on the name of the test in the **Tests List**, is covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(508).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*The Test Details page provides detailed information about a given test over the specified time-range.*

### Test Details Subnav Controls

The Test Details page includes the following UI elements on the subnav (light gray strip at the top of the page):

* **Edit Test**: A button that takes you to the **Test Settings Page** for this test.
* **Pause/Resume Test**: A button that does one of the following:

  + Pause: If the test is active, click the button to stop all testing related to this test.
  + Resume: If the test is paused, click the button to resume all testing related to this test.
* **Share**: Opens the **Share** dialog to enable you to share the current view (see **Sharing via the Share Dialog**).
* **Actions**: A drop-down menu from which you can choose to export the current view as a visual report (PDF) covering the page’s visualizations and tables. A notification appears when the PDF is ready to download.

### Main Test Details

The common area of the Test Details page (above the tabs) includes the following information:

* **Test type**: The type of the test (above the test name); see **Synthetics Test Types**.
* **Test name**: The name of the test, presented as the heading at the upper left of the page.
* **Time Range**: An indicator that shows the current local time range and also acts as a button that opens a popup that enables you to define the time range covered by the timeline (see **Time Range Control**).
* **Notes**: If the **Description** field is populated in the test’s Settings (Edit Test » **Test Information** tab), you can click or hover over this element to see the provided description. This element is present only for tests with a description.
* **Details tabs**: Different views of the test results for the currently selected time slice (see **Test Details Tabs**).

## Test Details Tabs

The tabbed views of the Test Details page are covered in the following topics.

> **Note:** The Details pages for BGP Monitor tests are structured differently from other test types and don't have the same tabs (see **BGP Monitor Details Page**).

### About Test Details Tabs

Each Test Details page shows the results of a test that has been configured with a **Test Settings Page**. The results are organized into the tabs listed below. The presence of a given tab on a given Test Details page depends on the type of test (see **Test Details by Test Type**):

* **Results**: A timeline, subtabs, and tables that provide health and performance information for the test (see **Test Results Details**).  
  **Path View**: A graphical representation of traceroute data from the subtests in this test (see **Test Path View**).
* **BGP**: The contents of this tab are similar to those of the Details page for a BGP Monitor test (see **Test Details BGP Tab**).

#### Test Details BGP Tab

The **BGP** tab appears on a Test Details page only when:

* The test type is either Page Load, or HTTP(S) or API.
* The **BGP Monitoring** switch is on in the test's settings.

The contents of the tab include the following main elements:

* A **Health timeline** (see **Results Tab Elements**)
* The **Chart Display Controls**
* An **Identified Issues** table (see **Results Tab Elements**)
* Visualizations for **Reachability/Visibility** and **AS Path Changes** (see **Upper Pane Elements**)
* An **AS Path Visualization Pane**

### Test Details by Test Type

As shown in the table below, the tabs and subtabs of a Test Details page vary depending on the test type (see **Synthetics Test Types**).

| Test Type | Results:  Sankey | Results:  Mesh | Results:  Map | Path View | BGP |
| --- | --- | --- | --- | --- | --- |
| BGP Monitor  (NO TABS) | No | No | No | No | No |
| **Network: Agent-to-Agent** | | | | | |
| Agent-to-Agent | No | No | Yes | Yes | No |
| Network Mesh | No | Yes | Yes | Yes | No |
| **Network: Agent-to-Server** | | | | | |
| Server IP Addresses | No | No | Yes | Yes | No |
| Server Hostname | No | No | Yes | Yes | No |
| Network Grid | No | No | Yes | Yes | No |
| **Network: Autonomous** | | | | | |
| ASN | Yes | No | Yes | Yes | No |
| CDN | Yes | No | Yes | Yes | No |
| Country | Yes | No | Yes | Yes | No |
| Region | Yes | No | Yes | Yes | No |
| City | Yes | No | Yes | Yes | No |
| **Application: DNS** | | | | | |
| DNS Server Monitor | No | No | Yes | No | No |
| DNS Server Grid | No | No | Yes | No | No |
| **Application: HTTP** | | | | | |
| HTTP(S) or API | No | No | Yes | Yes | Yes |
| Page Load | No | No | Yes | Yes | Yes |
| Transaction | No | No | Yes | No | No |

## Test Results Details

The **Results** tab of the **Test Details** pages is covered in the following topics.

### Results Tab Elements

The **Results** tab uses the following main UI elements to convey health status and performance information about the test:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(509).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

* **Health timeline**: A visual representation of the currently selected time range that shows the overall status of the test for each time slice in the range (see **Status Time and Scope**). By default, the latest time slice (far right) is selected, meaning that this is the slice for which details are given in the Details tabs (see below). Hovering anywhere in the timeline changes the selection to the time slice at that point, causing a corresponding change to the information in the Details tables and tabs (e.g., **Table** tab, **Mesh** tab) and opening a popup over the timeline that gives the start time and overall status of the selected slice.
* **Subtabs**: A set of subtabs that display, in graphical and tabular form, the metrics of test results for the specified time range (see **Results Tab Subtabs**).
* **Identified Issues**: A table listing the subtests in this test whose health status was either Warning or Critical at the currently selected time slice in the timeline. If all subtests were healthy at that point, the table is not displayed. The columns of the table are the same as the **Test Details Table**.
* **Test Timestamp**: An indicator giving the timestamp of the currently selected time slice in the timeline.
* **Test Details**: A table listing the subtests in this test (see **Tests and Subtests**) and providing performance metrics and other details about each subtest at the currently selected time slice in the timeline (see **Test Details Table**).

### Results Tab Subtabs

Depending on the type of test (see **Test Details by Test Type**), the **Results** tab of a Test Details page may include subtabs with the following views:

* **Sankey**: A Sankey diagram showing the site, connectivity type, provider, and target IP of the tested path.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(510).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*The Sankey subtab of the Results tab.*

* **Mesh**: A matrix providing the latency, packet loss, and jitter status, over the selected time range, of the individual tests between all of the agents in the matrix (see **Test Mesh Subtab**).
* **Map**: A map showing the physical location of hosted agents (private or public) used for this test and the cloud region for cloud agents.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(511).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*The Map subtab shows the geographic location of the agents involved in your tests.*

### Test Mesh Subtab

The **Mesh** subtab of the **Results** tab on a Test Details page shows a matrix representing the set of agents you've chosen to monitor and for which Kentik has automatically configured the To and From testing. As with any matrix, these agents are presented on the tab in a grid in which every agent is tested to and from every other agent. Inside each cell at the intersection of two agents is a representation of the health status (latency, packet loss, and jitter) of the subtest between those agents over the currently selected time slice.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(512).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*Mesh subtab buttons drop down menus that control mesh display.*

The following buttons above the grid determine how the subtest information in each cell is represented:

* **Source Agent**: A button that drops down a filterable list of agents from which you can choose one source agent. The button will then be labeled with the selected agent's name and row number, and the corresponding row will be highlighted in the mesh. To clear the label and highlight, click the **X** at the right of the button.
* **Target Agent**: A button that drops down a filterable list of agents from which you can choose one target agent. The button will then be labeled with the selected agent's name and column number, and the corresponding column will be highlighted in the mesh. To clear the label and the highlight, click the **X** at the right of the button.
* **Metrics**: A button that drops down a list of metrics — latency, jitter, and packet loss — in which you can click checkboxes to include/exclude individual metrics from the mesh. The button label lists the currently included metrics, which by default is all three.

  > **Notes**:
  >
  > + You can't uncheck all three individual metrics at once.
  > + In a Regular Density mesh (see **Density** below), unless a source agent and target agent are both highlighted in the mesh, only the dots representing selected metrics will be displayed in the cells.
* **Density**: A button that drops down a list of densities from which you can choose the display density of the mesh:

  + Regular Density: The names of all agents are displayed along both axes of the matrix and individual health dots are displayed for latency, packet loss, and jitter.
  + High Density: Each row and column of the matrix is labeled with a number, and the health of each cell is represented by a single dot showing whichever health status is worst between latency, packet loss, and jitter.
  + Max Density: The same as High Density except that the matrix is more compact and has no row or column headers.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(513).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*A mesh, without highlights, showing all metrics at the three different densities.*

> **Note:** The test types for which the **Mesh** subtab is included on the **Results** tab of the **Test Details** page are listed in **Test Details by Test Type**.

#### Mesh Information Popup

In the mesh itself, hovering over a cell (or a dot at High or Max density) opens a popup with the following information about the agents involved and the subtest results in both directions (forward and reverse):

* **Name**: The names of the agents.
* **Location** (in lozenge): The regions of the agents.
* **Metrics**: Latency, packet loss, and jitter metrics for the current time slice.
* **View Details**: A button that takes you to the corresponding **Subtest Details Page**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(514).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

While the most detailed information about a cell (intersection of a source and target agent) is available from the popup, the tables in the topics below may help you to better understand the information you can see without hovering.

#### Regular Density Meshes

If **Density** is set to Regular then the subtest results for each cell can be determined based on the characteristics in the below table.

| Status | Dot Size | Dot Color | Cell Shading |
| --- | --- | --- | --- |
| All metrics healthy | Small | Green | None |
| One or more metrics are missing | Small | Gray | Gray |
| One or more metrics have a warning | Medium | Orange | Orange |
| One or more metrics are critical | Medium | Red | Red |

> **Notes**:
>
> * In this context, "all metrics" refers to the all of the individual metrics currently selected with the **Metrics** drop-down (if only one metric is selected then "all metrics" refers to just that metric).
> * For shaded cells, the cell shading color (gray, orange, or red) corresponds to the worst health status (missing, warning, or critical) of the metrics that are currently selected with the Metrics control.

#### High/Max Density Meshes

If **Density** is set to either **High Density** or **Max Density** then the subtest results for each cell can be determined based on the characteristics in the below table.

| Status | Dot Size | Dot Color | Dot Halo | Cell Shading |
| --- | --- | --- | --- | --- |
| All metrics healthy | Small | Green | None | None |
| 1/3 or 1/2 metrics are missing | Small | Gray | Gray | None |
| 2/3 or all metrics are missing | Small | Gray | None | Gray |
| 1/3 metrics has a warning | Medium | Orange | Orange | None |
| 2/3 or all metrics have warnings | Medium | Orange | None | Orange |
| 1/3 metrics is critical | Medium | Red | Red | None |
| 2/3 or all metrics are critical | Medium | Red | None | Red |

> **Note**: In the cells of a High or Max density mesh, the overall health of each agent-to-agent subtest is represented by a single dot showing the worst health status (normal, missing, warning, or critical) of the metrics that are currently selected with the Metrics control.

## Test Path View

The **Path View** tab of a Test Details page, which shows a graphical representation of traceroute data from the subtests in a test, is covered in the following topics.

> **Note:** The test types for which this tab is included on the Test Details page are listed in **Test Details by Test Type**.

### Path View UI

The **Path View** tab includes the following UI elements:

* **Path View Controls**: The top pane in the tab displays the current settings for the Path View (see **Path View Controls**).
* **Path Hops and Geo Distance**: A timeline of the currently selected time range in which the average number of hops in the traceroutes for this test, as well as the average physical distance (estimated based on Geo IP) of those traceroutes, is plotted over time. Hovering over a point in the timeline shows a popover stating the average hops and distance at that point, and also determines the time slice for which paths are displayed in the Path View.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(515).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*Path Hops and Geo Distance shows the average hops and distance in the traceroutes for a given time slice.*

* **Path View**: Shows the hops taken by the traceroute tests, each represented as a colored circle whose AS is identified in the key, as well as the links between those hops. See **Path View Diagram**.
* **Legend** The key indicates, by color, the ASes to which each of the hops shown in the Path View belongs.

### Path View Controls

The Path View control pane enables you to customize the way the **Path****View** is displayed. The following settings are included:

* **Show**: A drop-down that allows you to select which agents to include in the diagram (default is the first 15 agents alphabetically).
* **Group hops by**: Choose how nodes are grouped when collapsed for display in the diagram:

  + No Grouping: Nodes are collapsed regardless of their AS or site.
  + ASN: Nodes within the same AS will be jointly represented by a single rectangle rather than as individual hops. The color of the rectangle will correspond to the legend below the diagram. A slider enables you to determine where in the path the setting applies by choosing the number of ASes from each end of the diagram (after source on left; before destination at right) that will be collapsed.
  + Site: All nodes within the same site will be jointly represented by a single circle rather than as individual hops.
* **Highlight Latency**: A drop-down that shows a field that allows you to specify the milliseconds of latency above which diagram elements will be highlighted.
* **Highlight Packet loss**: A drop-down that shows a field that allows you to specify the percent of loss above which diagram elements will be highlighted.
* **Collapse Timeouts**: A switch, that if enabled (default), nodes from which there is no response are collapsed together in the diagram.
* **Highlight links exceeding geo latency estimate**: A switch that turns on highlighting for links whose measured latency is greater than expected based on the estimated distance between hops (calculated using Geo IP).
* **Reset**: A button that restores the **Path View** settings to defaults.

> **Note:** The above settings persist for your subsequent visits to this Details page.

### Path View Diagram

The **Path View** diagram is a visualization of the nodes and links in the path taken from source to destination by the traceroutes for the subtests (e.g., from one agent to one IP in a given AS) in the test corresponding to this Test Details page. The time slice covered by the diagram is determined by the current time slice of the **Path Hops and****Geo****Distance** timeline at the top of the tab (see **Path View UI**). You can change the time slice by hovering over that timeline.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(516).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*The Path View diagram shows the traceroutes from the subtests in a test at a given time slice.*

The diagram is structured as follows:

* The sources (e.g., agents) are represented in the diagram as labeled rectangles at the far left.
* The target (e.g., the IP in an IP Addresses test) is represented as a circle at the right.
* Each hop is represented as a circle whose color identifies it as belonging to one of the ASes listed in the **Legend**.
* A hop that is indicated as an empty circle is "unidentified" (non-responsive).

  > **Note:** If the **Collapse** **Timeouts** switch is On (default), a contiguous series of unidentified hops that are in the middle of the path (neither the first two after the source nor the last two before the target) will be represented with a single circle stating the number of unidentified hops it represents.
* Each link between hops is represented as a line, the type and color of which indicates the characteristics of the link (see **Path View Lines**).
* Depending on the complexity of the diagram, when the diagram is loaded some paths may be collapsed to enhance visual clarity. Each such path is represented by a green circle that is labeled "Collapsed Path". To view the individual nodes of a collapsed path, click the circle to expand the path.

> **Notes:**
>
> * Links that exceed user-specified latency and hops that exceed user-specified loss may be highlighted (see **Path View Controls**) to draw your attention to problem areas.
> * Hover over a source, target, hop, or link to open a popup with details about that link (see **Path View Popups**).

#### Path View Lines

The type and color of the line representing a link between hops indicates the following characteristics of the link:

* **Both nodes identified**: We have detected the IP address of the nodes at both ends of the link.
* **Threshold exceeded**: The latency or packet loss on the link is in excess of a threshold determined in one of the following ways:

  + Compare packet loss with the **Highlight** threshold set manually by the user (see **Path View Controls**).
  + Compare latency with the **Highlight** threshold set manually by the user.
  + If the **Highlight** **links exceeding geo latency estimate** switch is On, compare latency with expected values calculated based on the distance between hops.

The table below shows how the conditions above determine the type and color of the line.

| Line type and color | Both nodes identified | Threshold exceeded |
| --- | --- | --- |
| Solid gray | Yes | No |
| Dotted gray | No | No |
| Solid red | Yes | Yes |
| Dotted red | No | Yes |

### Path View Popups

Hovering over a source, node (hop), link, or target will open an information popup. Popup contents vary depending on the type of element. For nodes the following additional factors influence the information we have available to show:

* If this is a customer node that is sending flow data and SNMP to Kentik.
* If it is a customer node, the SNMP metrics it is configured to send.
* If we’re able to do a reverse DNS lookup to get the hostname.

The information popups may contain the following types of information:

* **Source information**:

  + Agent name: The name of the agent.
  + Details: AS name, ASN, location, and IP address.
  + Traces: The number of traces from this agent.
  + Collapse other paths: If nodes are expanded on paths other this one, collapse them.
  + Collapse this path: If nodes are expanded on this path, collapse them.
  + Show raw trace data: Opens a popup that displays the actual traceroute text on which the traces for this agent are plotted in the diagram.
* **Node information**:

  + IP: The IP Address of the node.
  + Metrics: Loss and latency (the round trip time of the probe for this hop over the current time slice).
  + Details: AS name, ASN, location, and IP address.
  + Traces: The number of traces using this node.
  + More Info: Depending on availability, this pane at the right of the popup may include device name and utilization (CPU and memory), site name, ingress and egress IPs, interface information (e.g. name, type, and capacity), and host*n*ame.

    > **Note:** The information in the More Info pane is always as of the present moment, rather than as of the time slice that's currently selected in the timeline.
* **Link information**:

  + Metrics: Loss and link latency (the average round trip time of the later hop less the average round trip time of the earlier hop).
  + From/To: IP/CIDR of the nodes at either end of the link.
  + Traces: The number of traces using this link.
* **Target information**:

  + IP: The IP Address of the target.
  + Metrics: Loss and latency (the sum of each hop’s trip time).
  + Details: AS name, ASN, location, and IP address.
  + Traces: The number of traces toward this target.

## Test Details Table

The table on the **Results** tab of a Test Details page for a given test is covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(517).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*The Test Details table shows metrics about the subtests in the test.*

### About the Test Details Table

The **Results** tab of each Test Details page includes a **Test Details** table listing the subtests used in this test and providing information about them and their individual results during the current time slice (see **Status Time and Scope**). The content and layout of the table varies depending on the test type (see **Synthetics Test Types**).

> **Notes:**
>
> * Subtests with no results are not included in the table.
> * A**Test Details** table is not included when the test type is BGP Monitor.

### Details Table Columns

The columns of the Test Details table vary depending on the type of the test. The tables below show which columns are included for each test type. For descriptions of the columns, see **Standard Detail Columns** and **Additional Detail Columns**.

> **Notes:**
>
> * For information about the heading rows shown in the following tables, see **Details Table Heading Rows**.
> * BGP Monitor tests do not include a Test Details table.

#### Network: Agent-to-Agent Tests

| Test type | Heading row info | Agent | Target | Avg Latency | Avg Jitter | Packet Loss | Additional columns |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Agent-to-Agent | * Agent * Agent location * Target count | No (if heading row displays)  Yes (no heading row) | Yes | Yes | Yes | Yes | No |
| Network Mesh | * Agent * Agent location * Target count | No | Yes | Yes | Yes | Yes | No |

#### Network: Agent-to-Server Tests

| Test type | Heading row info | Agent | Target | Avg Latency | Avg Jitter | Packet Loss | Additional columns |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Server IP Addresses | * Agent * Agent location * Target count | No (if heading row displays)  Yes (no heading row) | Yes | Yes | Yes | Yes | No |
| Server Hostname | * Agent * Agent location * Target count | Yes | Yes | Yes | Yes | Yes | No |
| Network Grid | * Agent * Agent location * Target count | No | Yes | Yes | Yes | Yes | No |

#### Network: Autonomous Tests

| Test type | Heading row info | Agent | Target | Avg Latency | Avg Jitter | Packet Loss | Additional columns |
| --- | --- | --- | --- | --- | --- | --- | --- |
| ASN | * Agent * Agent location * Target count | No | Yes | Yes | Yes | Yes | * Provider * Connectivity Type * Potentially Impacted Traffic |
| CDN | * Agent * Agent location * Target count | No | Yes | Yes | Yes | Yes | * Provider * Connectivity Type * Potentially Impacted Traffic |
| Country | * Agent * Agent location * Target count | No | Yes | Yes | Yes | Yes | * Provider * Connectivity Type * Potentially Impacted Traffic |
| Region | * Agent * Agent location * Target count | No | Yes | Yes | Yes | Yes | * Provider * Connectivity Type * Potentially Impacted Traffic |
| City | * Agent * Agent location * Target count | No | Yes | Yes | Yes | Yes | * Provider * Connectivity Type * Potentially Impacted Traffic |

#### Application: DNS Tests

| Test type | Heading row info | Agent | Target | Avg Latency | Avg Jitter | Packet Loss | Additional columns |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Server Monitor | * Agent * Agent location * Target count | Yes | Yes | No | No | No | * - Resolution Time  - DNS Results  - DNSSEC |
| DNS Server Grid | * Agent * Agent location * Target count | No | Yes | No | No | No | * - Resolution Time  - DNS Response Code  - DNS Results  - DNSSEC |

#### Application: HTTP Tests

| Test type | Heading row info | Agent | Target | Avg Latency | Avg Jitter | Packet Loss | Additional columns |
| --- | --- | --- | --- | --- | --- | --- | --- |
| HTTP(S) or API | N.A. | Yes | No | Yes | Yes | Yes | * Certificate Expiry * Status Code * Domain Lookup Time * Connect Time * Response Time * Avg HTTP Latency |
| Page Load | N.A. | Yes | No | Yes | Yes | Yes | * Certificate Expiry * Status Code * Navigation Time * Domain Lookup Time * Connect Time * Response Time * DOM Processing Time * Avg HTTP Latency |
| Transaction | N.A. | No | No | No | No | No | * Health Status * Total Transaction Time * Transaction Completion * Screenshots |

#### Details Table Heading Rows

For tests with subtests that originate from multiple sources (i.e. multiple agents in an ASN test) the information that would otherwise be presented in the **Agents** column (see **Standard Detail Columns**) of the **Test Details** table is presented in heading rows. Each heading row represents one "from" agent and precedes the set of standard rows that represent the targets of that agent's subtests, serving to group the targets by their "from" agent.

### Standard Detail Columns

The following columns in the **Test Details** table are common to most test types:

* **Agent**: Information about the "from" agent that is the source of the subtest that is represented by this row of the table, including:

  + Type (icon): The type of agent, either global (Kentik icon) or private (see **About Synthetics Agents**).
  + Location: The country and city of the site where the agent is hosted or, for public cloud agents, the name of the cloud region where the agent is deployed.
  + Status: The agent's current status (see **Agent Status**).
  + AS: The name and number of the Autonomous System in which the agent is deployed.
* **Target**: A target is the network entity toward which a subtest is testing. If the targets of this test's subtests are not agents, then this column will be populated with the IP addresses of those targets. If the targets are agents, this column provides type, location, status, and AS information as described above for the agent column.
* **Avg Latency**: Latency on the tested path, averaged over the currently selected time slice.
* **Packet Loss**: The percent of packets lost on the tested path over the currently selected time slice.
* **Avg Jitter**: Jitter on the tested path, averaged over the currently selected time slice.
* **Details**: A button that takes you to the **Subtest Details Page** for this subtest.

### Additional Detail Columns

The following topics cover additional columns that appear in the**Test Details** table but are specific to individual test types.

#### DNS Test Columns

The following columns are specific to DNS tests (DNS Server Grid and DNS Server Monitor):

* **Resolution Time**: The average time, for DNS queries made during the current time slice, from making the query to receiving the response.
* **DNS Response Code**: The DNS RCode in the response to the DNS query, either 0 (no error) or a code describing an error. Possible DNS codes are listed in the IANA document at **https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-6**.![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(518).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)
* **DNS Results**: The DNS results returned during the currently displayed time slice, the type of which will be the **DNS Record Type** specified in the test's **DNS Settings**.

  + If no IPs are specified in the **Allowed DNS Results** field of the test's **Health Settings**, all IPs in this column are displayed in a gray lozenge.
  + If some IPs are specified in the **Allowed DNS Results** field, the allowed IPs are displayed in this column in a green lozenge, and the other IPs (not allowed) are displayed in a red lozenge.

    > **Note:** If more than one result is returned of a given kind (allowed or not allowed), the first result in the lozenge will be followed by "+" and the number of additional results. Hover over any lozenge to display the full list of results.
* **DNSSEC**: When **DNSSEC Validation** is switched On in the test's **Target and Agents** settings, this column will display the test's current validation state. Click on the indicator for one of the below states to see additional information in the **DNSSEC Results Dialog**:

  + Healthy: Each signing entity in the chain is validated as authentic.
  + Critical: At least one signing entity cannot be authenticated.

#### HTTP(S) or API Test Columns

The following columns are specific to HTTP(S) or API tests:

* **Certificate Expiry** (HTTPS only): The status of the agent’s SSL certificate. This column is displayed only when **Ignore TLS Errors** is switched Off in the test's **HTTP Settings**:

  + If the server certificate is valid, its expiry date is displayed in a lozenge whose color is based on the thresholds set with **Certificate** **Expiry** in the test's **Health Settings**).
  + If the server certificate is expired, “SSL Certificate Expired” is displayed in a red lozenge.
* **Status Code**: The HTTP Return Code (see **Host Traffic Dimensions**) returned during the currently displayed time slice.

  > **Note:** A subtest for which multiple status codes were returned during the time slice will be represented in the **Test Details** table by multiple rows.
* **Domain Lookup Time**: From when the domain lookup starts to when the domain lookup is finished.
* **Connect Time**: From when the request to open a connection is sent to the network to when the connection is opened.
* **Response Time**: From when the first byte of the response is received to when the last byte of the response is received.
* **Avg HTTP Latency**: The average time, for HTTP requests made during the currently displayed time slice, from making the request to receiving the last byte of the response. This figure is the sum the columns Domain Lookup Time, Connect Time, and Response Time.
* **Response Headers Health**: This column is linked directly to the Expected Responses section on a test’s Health tab (see **Health Settings**). If no key/value is present, the value under the column is Healthy. If the key and value fields are populated, that key/value set must be present in the response for the test to be considered Healthy.

#### Page Load Test Columns

The following columns are specific to Page Load tests:

* **Certificate Expiry** (HTTPS only): The status of the agent’s SSL certificate. This column is displayed only when **Ignore TLS Errors** is switched Off in the test's **HTTP Settings**:

  + If the server certificate is valid, its expiry date is displayed in a lozenge whose color is based on the thresholds set with **Certificate** **Expiry** in the test's **Health Settings**).
  + If the server certificate is expired, “SSL Certificate Expired” is displayed in a red lozenge.
* **Status Code**: The HTTP Return Code returned during the currently displayed time slice (for definitions, see **Status Codes**). A code is considered healthy (green lozenge) if the code is:

  + Selected in the **Valid** **HTTP** **Codes** field on a test’s Health tab
  + Under 400
* **Navigation Time**: The duration from the browser's receipt of the HTTP request to when it is ready to fetch the document.
* **Domain Lookup Time**: The duration between the start of domain lookup and the completion of domain lookup.
* **Connect Time**: The duration from when the request to open a connection is sent to the network to when the connection is opened.
* **Response Time**: The duration between receipt of the first byte of the response to receipt of the last byte.
* **DOM Processing Time**: The duration from when the browser starts to render the page to when the page is fully loaded. This corresponds to the time between domLoading and LoadEventEnd.
* **Avg HTTP Latency**: This figure is the sum of the columns Navigation Time, Domain Lookup Time, Connect Time, Response Time, and DOM Processing Time.
* **Response Headers Health**: The test status that displays in this column is set in the Expected Responses section on a test’s **Health** tab (see **Health Settings**). If no key/value is present, the value under the column is Healthy. If the key and value fields are populated, that key/value set must be present in the response for the test to be considered Healthy.

#### Transaction Test Columns

The following columns are specific to Transaction tests:

* **Health Status**: The health of the agent for the currently selected time range.
* **Total Transaction Time**: The elapsed time of the transaction (the time it took to run the Puppeteer script from start to finish).
* **Transaction Completion**: The ultimate outcome of the test:

  + Pass: The transaction completed.
  + Fail: The transaction did not complete before the test timeout was reached.
* **Screenshots**: The number of screenshots captured during the corresponding segment of the timeline.

### DNSSEC Results Dialog

This read-only dialog provides additional information about the validation performed in a DNS test when **DNSSEC Validation** is switched on in the test's **Target and Agent** settings. The dialog opens when you click a lozenge (Healthy or Critical) in the DNSSEC column of the Test Details table.  
  
The DNSSEC Results dialog displays the following information:

* **Name**: The domain (host address) being validated.
* **Qtype**: The type of DNS record to return from the query. The default is A. See DNS Record Type under **DNS Settings**.
* **Rcode**: Return code. Any result other than NoError will result in a failed test. See **IANA DNS RCODEs**.
* **Authenticated Data (flag AD)**: True or False. A false result will produce a failed test.
* **Trust Chain**: Details for each zone of the trust chain including flags, UNIX timestamps, and keys.

![DNSSEC results for kentik.com showing authentication status and DNS records.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/TCC-DNSSEC-results-dialog(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*The DNSSEC dialog.*

## Subtest Details Page

The Details pages for individual subtests are covered in the following topics.

> **Note:** There are no subtests, and thus no Subtest Details pages, for tests whose type is BGP Monitor.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(520).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*The Subtest Details page shows details about an individual subtest in a test.*

### Subtest Details UI

The **Subtest Details** page for an individual subtest includes the following UI elements:

* **Subtest Subnav controls**: The Subtest Details page includes the following UI elements on the subnav (light gray strip at the top of the page):

  + Edit Test: A button that takes you to the **Test Settings Page** for this test.
  + Share: Opens the **Share** dialog to enable you to share the current view (see **Sharing via the Share Dialog**).
  + Actions: A drop-down menu from which you can choose to export the current view as a visual report (PDF) covering the page’s visualizations and tables. A notification appears when the PDF is ready to download.
* **Subtest identifier**: A heading at the upper left of the page identifies the test by showing what it is testing from and toward. In the case of a subtest from one global agent to another, for example, the "from" could be "Chicago, IL, United States" and the "toward" could be "New York, United States."

  > **Note:** When the from or toward is an agent:
  >
  > + An icon to the left of the agent name indicates the agent's type (global or private).
  > + A field below the agent name indicates the cloud region if the agent is in a public cloud, or the AS name and number if the agent is in a data center.
  > + A green checkmark indicates that the agent status is active.
  > + A red exclamation mark indicates that the agent status is offline.
* **Time Range**: The lookback period covered by the timeline.
* **Details tabs**: A set of tabs, which vary depending on the type of test (see **Synthetics Test Types**), that provide different views of the test results for the currently selected time slice (see **Subtest Details Tabs**).

### Subtest Details Tabs

The following tabs, which vary depending on the type of test (see **Synthetics Test Types**), provide different views of the test results for the currently selected time slice:

* **Metrics**: A set of charts detailing the results from this subtest over the currently selected time range (see **Subtest Metrics Tab**).
* **Path View** (not present for Page Load subtests): A graphical representation of traceroute data from this subtest. The layout of this tab is the same as described in **Test Path View**, but only the traceroute for this individual subtest is shown.
* **Waterfall** (Page Load and Transaction tests only): Shows the load order and load duration of each element in the DOM of a tested page (see **Subtest Waterfall Tab**).
* **Results** (Transaction tests only): Shows the total completion time for the transaction as well as screenshots (if any are specified in the Puppeteer script) showing the state of the browser contents (see **Subtest Results Tab**).

### Subtest Metrics Tab

The **Metrics** tab on the Subtest Details page for a given subtest includes the following main components:

* **Health timeline**: An overall timeline, representing the currently selected time range, that is present on the Metrics tab for tests of every type (see **Metrics Tab Timeline**).
* **Metrics charts**: A set of charts, varying by test type (see **Metrics Chart Locations**), that detail the results from this subtest over the currently selected time range (see **Metrics Tab Charts**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(521).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*The Metrics tab shows charts illustrating subtest metrics.*

#### Metrics Tab Timeline

The **Metrics** tab health timeline is a visual representation of the currently selected time range that shows the overall status of the subtest for each time slice in the range (see **Status Time and Scope**). By default, the latest time slice (far right) is selected, meaning that this is the slice for which details are given in the **Metrics Tab Charts**. Hovering anywhere in the timeline changes the selection to the time slice at that point, causing a corresponding change to the information displayed in the charts and opening a popup over the timeline that gives the start time and overall status of the selected slice.

#### Metrics Tab Charts

In each of the metrics timelines, the solid line plots the metrics values measured for the time slices over the time range (see **Status Time and Scope**). The shaded areas, which show the rolling standard deviation in the metric over the range, are colored to indicate the metric's status during each time slice. Hovering on the timeline for a given slice opens a popup with values for that point in time on the main (top of page) timeline as well as on each of the charts on the metrics tab.

#### Metrics Chart Locations

As shown in the table below, the charts on the **Metrics** tab vary depending on the type of the test to which the subtest belongs.

| Test type | Traffic | Avg Latency | Packet Loss | Avg Jitter | Additional charts |
| --- | --- | --- | --- | --- | --- |
| **Network: Agent-to-Agent Tests** | | | | | |
| Agent-to-Agent | No | Yes | Yes | Yes | No |
| Network Mesh | Yes | Yes | Yes | Yes | No |
| **Network: Agent-to-Server Tests** | | | | | |
| Server IP Addresses | Yes | Yes | Yes | Yes | No |
| Server Hostname | Yes | Yes | Yes | Yes | No |
| Network Grid | Yes | Yes | Yes | Yes | No |
| **Network: Autonomous Tests** | | | | | |
| ASN | Yes | Yes | Yes | Yes | No |
| CDN | Yes | Yes | Yes | Yes | No |
| Country | Yes | Yes | Yes | Yes | No |
| Region | Yes | Yes | Yes | Yes | No |
| City | Yes | Yes | Yes | Yes | No |
| **Application: DNS Tests** | | | | | |
| DNS Server Monitor | No | No | No | No | * DNS Results * DNS Resolution |
| DNS Server Grid | No | No | No | No | * DNS Results * DNS Resolution |
| **Application: HTTP Tests** | | | | | |
| HTTP(S) or API | Yes | Yes | Yes | Yes | * HTTP Stages Breakdown * Avg HTTP Latency * HTTP Status Code |
| Page Load | No | Yes | Yes | Yes | * HTTP Stages Breakdown * Avg HTTP Latency * HTTP Status Code |
| Transaction | No | No | No | No | * Total Transaction Completion Time * Transaction Completion * Screenshots |

> **Notes:**
>
> * BGP tests have no subtests.
> * For Mesh and Agent-to-Agent tests, traffic charts are present only if the Site IP Classification setting (see **Site Field Definitions**) includes at least one valid IP address in any one of the three fields.

#### Standard Metrics Charts

The following charts in the **Metrics** tab are common to most test types:

* **Traffic**: The bitrate of actual inbound and outbound traffic over the currently specified time range, averaged within each time slice and plotted against a single timeline (inbound positive Yes axis, outbound negative):

  + Inbound: From the subtest's "toward" entity (e.g. hostname, ASN).
  + Outbound:To the subtest's "toward" entity.
* **Metrics timelines**: Three timelines that plot the following metrics over the currently selected time range, averaged within each time slice:

  + Avg Latency (ms): Latency on the subtest path.
  + Packet Loss (%): The percent of packets lost on the subtest path.
  + Avg Jitter (ms): Jitter on the subtest path.

> **Note:** A test may fail for some or all of the currently specified time range. Test failure may be caused by a timeout (the test took too long to complete) or by errors (e.g., agent or OS error). When a test fails during a given time period, no results are available to plot for that period in the metrics timelines. The way this condition is indicated in the charts depends on the type of chart and the kind of failure:
>
> * For both average latency and average jitter, the results plot will have a gap corresponding to the duration of the failure.
> * For packet loss:
>
>   + If the failure is due to error, there will be a gap in the results plot.
>   + If the failure is due to timeout, packet loss will be shown as 100%.

#### Additional Metrics Charts

The following charts are found on the **Metrics** tab for only a limited number of test types:

* **DNS Server Monitor tests**:

  + DNS Results: A set of lozenges that each display a DNS result returned during the currently displayed time slice, the type of which will be the DNS Record Type specified in the test's **DNS Settings**.
  + Resolution Time: The duration, averaged within each time slice, between making a DNS query and receiving the response.
* **HTTP(S) or API tests**:

  + HTTP Stages Breakdown: The stacked bar chart showing the amount of time, averaged within each time slice, spent in the key stages of an HTTP transaction: Domain Lookup Time, Connect Time, and Response Time, the sum of which is the time shown in the Avg HTTP Latency chart. Breaking HTTP Latency into component parts helps troubleshooting slow HTTP response or long page load times.
  + Avg HTTP Latency: The time, averaged within each time slice, from making an HTTP request to receiving the last byte of the response.
  + HTTP Status Code: The HTTP Return Codes (see **Host Traffic Dimensions**) returned over the time range.
* **Page Load tests**: In addition to the charts displayed for HTTP(S) or API tests, the Metrics tab of a Page Load test may include the following additional chart:

  + CSS Selectors*:* If CSS selectors are configured for this test (see **Configure Request Settings**) the Metrics tab will include a time series chart that shows a count of those configured selectors that were found in the loaded page. When you hover over the chart, a popup opens that shows which selectors were configured and found/not found at that point in the timeline.

### Subtest Waterfall Tab

The **Waterfall** tab shows the load order and load duration of each element in the DOM of a page tested with a Page Load test (see Page Load under Web Tests in **Synthetics Test Types**). The source of the presented information is a HAR file exported from the headless Chromium browser instance running in a ksynth App agent (see **Synthetics Agent Types**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(522).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*The Waterfall tab shows the load order and load duration of each element in the DOM.*

#### Waterfall Tab UI

The **Waterfall** tab includes the following controls and information:

* **Timeline**: Like the **Metrics Tab Timeline**, the timeline in the Waterfall tab is a visual representation of the selected time range, with each segment representing an increment within that range. The segment whose test results are currently displayed in the resources list is outlined in green. If the duration of the individual increments is greater than the testing frequency set for this page load test, then the results shown in the list will actually be an average of the tests run during that time increment.
* **Filter**: Filter the resources shown in the table to those that match the text entered in the field. The drop-down selector determines what will be matched:

  + URL: Matches text that appears in the URL of the requested resource.
  + BODY: Matches text that appears in the response body of the requested resource.
* **Type Selector**: Choose the resources that will be displayed in the table, either all types or one or more individual types (e.g., JavaScript, CSS, image).
* **Errors Only**: A checkbox that filters the resources in the list to only those with errors.
* **Resources List**: A list of the resources requested by the tested page (see **Waterfall Resources List**).

#### Waterfall Resources List

The list in the **Waterfall** tab includes the following columns that provide detailed information on the download of each resource requested by the tested page:

* **File**: The name of the requested resource.
* **Status**: The HTTP status returned in the response to the request for the resource.
* **Method**: The HTTP method of the request for the resource.
* **Domain**: The domain from which the resource was requested.
* **Type**: The MIME type of resource (e.g. HTML, CSS, JavaScript, etc.).
* **Size**: The size of the resource in bytes.
* **Time**: The duration in milliseconds between when the resource was initially requested and when it was completely downloaded.
* **Waterfall**: A visual representation of the download of each resource and of the entire page. The download for each resource is represented by a horizontal bar made up of multiple segments, each of which is colored to correspond to a different phase of the download (stalled, DNS lookup, SSL, initial connection, request sent, waiting, content downloaded). Hovering on the bar opens a popup that shows the duration of each phase as well as the total download time for the resource.

### Subtest Results Tab

The **Results** tab on the Subtest Details page for a Transaction subtest includes the following main components:

* **Health timeline**: An overall timeline, representing the currently selected time range and segmented into time slices. The timeline is the same as the timeline on a Metrics tab (see **Metrics Tab Timeline**). Hovering on a given time slice opens a popup on the timeline with the health status for that time slice as well as a popup over the completion time chart containing completion time values for that time slice.
* **Total Transaction Completion Time**: A chart that plots (as a solid line) the transaction times measured for the time slices over the time range. The shaded area around the plot shows the rolling standard deviation in completion time over the range. Both the plot and the shaded areas are colored to indicate the metric's status during each time slice.
* **Transaction Completion**: A chart on which each transaction over the indicated time range is represented as a dot over which one can hover to pop up additional information. Passing transactions are plotted on the PASS (upper) line in green, while failing transactions (the test timed out before the transaction script was completed) are plotted on the FAIL (lower) line in red.
* **Screenshots**: A set of screenshots for the time slice corresponding to the currently selected segment of the timeline. The screenshots are taken at points specified in the Transaction test script (see **About Transaction Test Scripts**).

  > **Note:** Screenshots will not be present for a given time slice if the health status during that time was Critical.

## BGP Monitor Details Page

The Details page for a BGP Monitor test is covered in the following topics.

### BGP Monitor Page Overview

A Details page for a BGP Monitor test includes the **Test Details Subnav Controls** and the **Main Test Details** that together comprise the standard UI at the top of every **Test Details Page**. Every BGP Monitor page also includes the health timeline (see **Results Tab Elements**) found at the top of the Results tab on test Details pages. The structure of the rest of the page, however, varies from the Details pages for other test types (see **BGP Monitor Panes**).

### BGP Monitor Panes

The unique content of the Details page for a BGP Monitor test is structured around a set of three panes whose content at any given time is determined by the **Chart Display Controls**. These panes contain the elements described in the topics below.

> **Notes:**
>
> * On the Details page for a BGP Monitor test, the maximum valid duration that can be set with the Time Range control (see **Main Test Details**) is one day.
> * BGP Monitor tests have no subtests and thus no subtest Details pages.

#### Identified Issues Pane

The Identified Issues pane includes a table that lists each of the events that generated an incident within the selected time range. A drop-down above the list determines which kinds of incidents are visible in the table. See **Identified Issues Controls** and **Identified Issues List**.

> **Note:** This pane is displayed only when the status for one or more time slices in the currently selected time range is Critical or Warning.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(523).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*The Identified Issues table is shown when the status of any time slice in the time range is not healthy.*

#### Upper Pane Elements

Of the remaining two panes, the upper pane contains the following main elements:

* **Reachability/Visibility**: A time series chart showing the percentage of vantage points (BGP routing data collectors; see **BGP Event Datasets**) that are able to reach the currently specified prefixes over the currently specified time range (see **Reachability Time Series**).
* **AS Path Changes**: A time series chart showing the number of AS path changes observed by the vantage points during the currently specified time range (see **AS Path Changes Time Series**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(524).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*The Show All setting displays both charts in the upper pane.*

#### AS Path Visualization Pane

The pane at bottom contains the **AS Path Visualization**, which is a diagram showing the AS Paths at the currently specified point in the time range (see **AS Path Visualization**).

> **Note:** For the elements in this pane, the "currently specified time" doesn't change when you hover over the health timeline; you must click in the timeline to change the time.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(525).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*A diagram showing the current AS Paths.*

### Chart Display Controls

The elements that are displayed at any given time within the three panes listed above are determined by the controls listed in the topics below.

> ***Notes:***
>
> * When visible, the**Identified Issues** pane appears below the health timeline and above the upper pane.
> * The upper pane and the AS Path Visualization pane can't both be hidden at the same time.
> * When either the upper pane or the AS Path Visualization pane is hidden, a checkbox to show it will be included in the control set of the other pane.

#### Identified Issues Controls

The **Identified****Issues**pane includes the following controls:

* **Filter Identified Issues**: A drop-down from which you choose which types of detected issues will be displayed in the Identified Issues list:

  + Clear: Click to see all of the **identified issues** within the specified time range (default view).
  + Unexpected Origin: Shows Origin Hijack events during the specified time range. Practically, this means another ASN (outside of the specified Upstream list, or against legitimate RPKI Origin) than yours has been observed by vantage points announcing the configured prefix(es).
  + Invalid RPKI: Shows each incident with an invalid RPKI that occurred within the specified time range.
  + Unexpected Upstream: The first hop in the AS Path detected prefixes seen from vantage points that do not match the list of allowed ASNs.
  + Low reachability: Shows each incident where the reachability (percentage of vantage points detecting this prefix) has a Warning or Critical status during the specified time range.
* **Hide Alarms**: Toggle on or off the display of the Identified Issues pane.

#### Upper Pane Controls

The upper pane appears either just below the Identified Issues pane or just below the page's health timeline if the Identified Issues pane is hidden.  
  
The following controls appear at the top of the pane:

* **Show**: A drop-down from which you choose which of the time series elements will be displayed: Show All (both visualizations), Reachability/Visibility only, or AS Path Changes only.
* **Prefix**: A drop-down from which you choose the prefix whose information is displayed in all three of the time series visualizations. The prefixes available on the drop-down are those that were specified when the test was created (see Prefixes to Monitor in **BGP Monitoring**).
* **Specific Prefix**: A drop-down, present if the Include more specific prefixes switch was on when the test was created, from which you can choose a more specific prefix to show in the time series charts. Choose Clear to revert the charts to the more general prefix selected in the Prefix drop-down.
* **AS filter**: A checkbox that appears only when an AS has been clicked in the **AS Path Visualization**. The checkbox is labeled with the ASN of the AS that was clicked, and it indicates that the page's time series charts, path diagram, and events table have all been filtered to include only information related to that AS. To remove the filter, uncheck the checkbox.
* **Hide**: Toggle on or off the display of the following visualizations:

  + Hide Alarms (present only when the Identified Issues pane is hidden): Shows the Identified Issues pane when unchecked.
  + Hide Timelines: Hides the entire upper pane when checked, including the other controls and the time-series charts themselves. Uncheck it to display the pane again.
  + Hide Paths Graph (present only when the AS Path Visualization pane is hidden): Shows the AS Path Visualization when unchecked.
* **Options** (vertical ellipses): Present only when Reachability/Visibility is displayed, it pops up a drop-down menu from which you can choose whether the visualization displays reachability by Origin or Upstream.

#### AS Path Visualization Controls

The AS Path Visualization pane includes the following checkboxes:

* **Hide ASN Name**: Eliminates all AS names from the AS Path Visualization.
* **Hide Paths Graph**: Check to hide the AS Path Visualization pane.
* **Hide Timelines** (present only when the upper pane is hidden): Uncheck to show the upper pane.

### Identified Issues List

The**Identified Issues** list appears only when only when the status for one or more time slices in the currently selected time range is Critical or Warning. The list includes the following columns:

* **Status**: The status of the test during the currently selected time slice (see **Synthetics Test Status**).
* **State**: Indicates one of the following incident states for the test:

  + Active: The test status became Critical or Warning before or during the time slice and remained so for the rest of the time slice.
  + Cleared: The test status became Critical or Warning before or during the time slice but returned to normal before the end of the time slice.
* **Type**: The types of BGP events listed, which — depending on the **Identified Issues Controls** — may include events of unexpected origin, invalid RPKI, unexpected upstream, or low reachability.
* **Prefix**: The prefix associated with the event. Click a prefix to see visualizations for that prefix in the upper and AS Path Visualization panes.
* **Details**: This column will show different details depending on the type of issue displayed (see **Identified Issues Details**).
* **Start**: The date-time during the time slice at which the incident became Active (test status became Critical or Warning).
* **End**: The date-time, if any, during the time slice at which the test status returned to normal (the incident state became Cleared). If N/A, the incident is ongoing.

#### Identified Issues Details

This column in the table can display one of the following details for the issue displayed:

* **Unexpected Origin**: The Origin AS for the prefix is shown.
* **Invalid RPKI**: The detected ASN for the prefix (which is displayed) that does not match the expected ASN.
* **Unexpected Upstream**: The unexpected ASNs from the first hop in the AS Path.
* **Low Reachability**: Shows how many vantage points from which the prefix is reachable.

### Chart Display Mode

The way that the **Reachability/Visibility** and **AS Path Changes** charts on the BGP Monitor Details page are rendered depends on the prefixes being monitored:

* **Single prefix, single AS**: The chart is rendered with a single color.
* **Single prefix, multiple ASes**: The chart is rendered with a different color assigned to each AS.

  > **Note:** This mode occurs when a prefix is advertised by more than one AS, indicating a possible hijack.
* **Multiple prefixes**: The chart is rendered with a different color assigned to each prefix.

  > **Note:** Multiple-prefix monitoring occurs when the **Include more specific prefixes** switch is on in the test's **BGP Monitoring**settings, but no specific prefix is selected with the **Specific Prefix** drop-down in the upper pane (see **Chart Display Controls**).

### Reachability Time Series

The BGP Monitor **Reachability/Visibility** time series chart shows the reachability of the target prefix over the currently specified time range. At each increment in the time range (e.g., each minute when the time range is one day) the chart shows the vantage points (public BGP routing data collectors) that are able to reach the target prefixes as a percent of all Kentik-utilized vantage points.  
  
The current direction of the chart, either Origin or Upstream, is indicated in parentheses to the right of the chart's heading. To set the direction, click the Options button (vertical ellipses) at the top right corner of the pane.  
  
The way the chart is filled in depends on the **Chart Display Mode**:

* **Single prefix, single AS**: The area below the plot line will be filled in with a single color.
* **Single prefix, multiple ASes**: The area corresponding to the percent of reachability via each AS will be filled in with a different color.
* **Multiple prefixes**: Each prefix is rendered in its own separate line with its own color.

Hovering at a given point in the chart opens a popup that shows the AS (or ASes, if there's been a hijack) that advertised the prefix and the reachability percentage at the corresponding increment in the time range.

> **Note:** If there's an upstream leak and you’ve specified ASes with the Upstream Leak Detection field in the test's **BGP Monitoring** settings, click the Options button (vertical ellipses) and select Upstream to see the leaking ASNs indicated in red.

### AS Path Changes Time Series

This time series chart shows any changes in the AS Path to a monitored prefix as observed from all vantage points. The changes are rendered as a spike extending up from the chart's timeline; the more changes at a given increment in the time range, the taller the spike. If the page is in multi-prefix mode (see **Chart Display Mode**) the spikes will be colored according to the prefixes involved, otherwise the spikes will be in a single color. Hovering over the chart opens a popup stating the time and the number of changes for each prefix.

### AS Path Visualization

This diagram shows the AS paths via which, at the currently selected increment in the time range, the monitored prefix was reachable from one or more vantage points. The hops (ASes) in each path are represented as shaded rectangles, with the monitored prefix on the right and the AS containing that prefix second from right. The rectangles are connected by lines that represent the links between the shown hops.  
  
To reduce clutter, if less than three percent of the vantage points advertise an AS in their paths to the monitored prefixes, then that AS is "pruned" from the diagram, meaning that neither the AS nor the upstream paths to it will be shown.  
  
The following information is shown in the rectangle representing each AS:

* **ASN**: The AS number of the hop.
* **AS name**: The AS name of the hop.
* **Vantage point**s (binocular icon): Indicates that the AS either includes one or more vantage points itself or is the first shown AS in a path that has been pruned upstream (thereby hiding the AS in which a vantage point actually exists).

> **Notes:**
>
> * For this diagram (unlike the time series charts) the "currently specified increment" doesn't change when you hover over the health timeline. You must click in the timeline to update the path diagram to a new point in time.
> * When you click an AS in the diagram, the page's time series charts, path diagram, and events table are all filtered to include only information related to that AS (the Identified Issues list is unaffected). A checkbox, labeled with the ASN of the clicked AS, will appear in the control set of the page's upper content pane (see AS filter in **Chart Display Controls**). To remove the filter, uncheck the checkbox.

#### AS Popup

The popup that's displayed when you hover over an AS includes the following information:

* **ASN**: The AS number of the hop.
* **AS name**: The AS name of the hop.
* **Vantage Points**: The number of vantage points in the AS.
* **Transiting Paths**: The number of advertised paths from a vantage point to a monitored prefix that include this AS.
* **Pruned Paths**: The number of paths upstream of the AS that were pruned because their ASes were advertised in less than three percent of the paths from all vantage points to the monitored prefixes.
* **Show Vantage Points (VPs)**: A link that pops up a list of the vantage points that exist within this AS or in an upstream AS in a pruned path for which this is the first shown AS (see **Show Vantage Points**).

#### Show Vantage Points

This link displays at the bottom of an **AS Popup**, and when clicked, pops up a list of the vantage points that exist within this AS, or in an upstream AS in a pruned path for which this is the first shown AS. For each such vantage point, the list shows the following information:

* **Collector**: The device (physical or virtual) with which the vantage point is peered in order to share BGP routing information. The displayed value is a concatenation of the dataset for which the device collects information and the name that the dataset provider has assigned to the device.
* **VP (peer) ASN**: The AS number of the vantage point.
* **VP (peer) IP**: The IP address of the vantage point.
* **AS Path**: The vantage point's best route to the monitored prefix.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(526).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

#### Link Popup

The popup that's displayed when you hover over a link includes the following information:

* **From**: The ASN and name of the starting AS of the link.
* **To**: The ASN and name of the ending AS of the link.
* **Paths**: The number of paths advertised by the starting AS to the ending AS.

## Test Preview Page

The **Test Preview** page is accessed by clicking the **Preview** button on the **Add/Edit Test Settings** page (see **Test Settings Page**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(527).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

*Test Preview pages are similar but not identical to Test Details pages of the same test type.*

### Test Preview Overview

Once a test has been saved on a **Test Settings Page**, many settings are locked and can't be modified. To check various protocols before saving your test, so you can see what works best, you can create a Test Preview page for the test by clicking the Preview button in the **Test Management Controls** at the top of a test's settings page.  
  
A Preview page opens in a new browser tab or window. The Test Settings page from which you created the preview will remain present in its original tab so that you can make changes to the test settings and create new previews as needed.  
  
A Test Preview page will be identical in most respects to a standard **Test Details Page** for a test of the same type, with the exception of the following key differences:

* None of your monthly test credits are used for a Test Preview.
* While a newly created test takes a few minutes to start returning data, a Test Preview page returns results within a few moments.
* A Test Preview records data only for the moment in time when it's created. Results that depend on sustained testing over time are not included.
* The UI elements on a Test Preview page may vary from those on a standard page for the same test type (see **Test Preview UI**).
* Test Previews will not send notifications even if they are configured in the test.
* Test Previews will not generate alerts and thus will not impact the **Incident List** on the Synthetics Dashboard.

> **Notes:**
>
> * Test Previews are not currently available for Autonomous tests (ASN, CDN, Country, Region, or City).
> * The Preview function for BGP Monitor tests is performed by the **BGP Route Viewer**. Like a Test Preview, the BGP Route Viewer does not consume test credits and enables you to see which protocols work best before saving a test.
> * Test Previews will only work with App Agents. Selecting one or more Network Agents for your test will produce an error when you try to preview the test.

### Test Preview UI

The UI of a Test Preview page differs in some respects from the UI of a Test Details page for the same test type. The differences that are not specific to individual tabs are covered in the topics below.

#### Unique Test Preview UI

The following elements are found on Test Preview pages but not on Test Details pages:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(528).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A06Z&se=2026-05-12T10%3A25%3A06Z&sr=c&sp=r&sig=nZYXu9%2FwMwg%2FtR4%2BX%2BgbGFsUMqRRFyNklpmWoYUdFWc%3D)

* **Run Preview Again**: A button that runs the preview again in the same browser tab.
* **Save Test**: A button that saves and starts the test, which will appear on the **Tests List** on the Test Control Center.
* **Credit usage**: An indicator (under the Save Test button) stating how many credits will be used if you save the test using the current settings.

#### Excluded Test Details UI

The following elements are found on Test Details pages but not on Test Preview pages:

* The buttons detailed in **Test Details Subnav Controls**.
* The Time Range control (see **Main Test Details**).
* The Health timeline (see **Results Tab Elements**).

### Test Preview Tabs

As with a Test Details page, the results on a Test Preview page are organized into tabs whose presence and contents depend on the type of test. The tabs included in a Test Preview overlap with, but are not identical to, those of a standard Test Details page, in part because previews don't include any test results that depend on sustained testing over time.  
  
The tabs of a Test Preview page include the following:

* **Results**: A subtab and table that provide health and performance information for the Test Preview (see **Test Preview Results Tab**).
* **Path View**: A graphical representation of traceroute data from the subtests in this test (see **Test Preview Path View Tab**).

  > **Note:** The Path View tab on a Test Preview page doesn't include the Path Hops and Geo Distance timeline (see **Path View UI**).

#### Test Preview Results Tab

The **Results** tab uses the following main elements to convey health status and performance information about the Test Preview:

* **Subtabs**: Depending on test type, the tab may include one or both of the following subtabs.

  + Mesh: A matrix providing the latency, packet loss, and jitter status of the individual tests between all of the agents in the matrix at the time of the Test Preview (see **Test Mesh Subtab**).
  + Map: A map showing the physical location of hosted agents (private or public) used for this test and the cloud region for cloud agents.
* **Identified Issues Table**: A table listing the subtests in this test whose health status was either Warning or Critical when the Test Preview was created. If all subtests were healthy, the table is not displayed. The table has the same columns as the Test Details table.
* **Test Timestamp**: An indicator giving the timestamp of the Test Preview.
* **Test Details table**: A table listing the subtests in this test (see **Tests and Subtests**) and providing performance metrics and other details about each subtest from when the Test Preview was created (see **Test Details Tab**).

#### Test Preview Path View Tab

The **Path View** tab shows a graphical representation of traceroute data from the subtests of a test and contains the following UI elements:

* **Path View Controls**: The top pane in the tab displays the current settings for the Path View (see **Path View Controls**).
* **Path View**: Shows the nodes and links in the path taken from source to destination by the traceroutes for the subtests in the test (see **Path View Diagram**):

  + Each node (hop) is represented as a colored circle whose AS is identified in the key
  + Each link between hops is represented as a line, the type and color of which indicates the characteristics of the link.
* **Path View Popups**: Informational popups that open when you hover over a source, node, link, or target (see **Path View Popups**).
* **Legend**: A key that indicates, by color, the ASes to which each of the hops shown in the Path View belongs.

#### Preview Details by Test Type

As shown in the table below, the tabs and subtabs of a Test Preview page vary depending on the test type (see **Synthetics Test Types**).

| Test Type | Results: Mesh | Results: Map | Path View |
| --- | --- | --- | --- |
| BGP Monitor | No | No | No |
| **Network: Agent-to-Agent** | | | |
| Agent-to-Agent | No | Yes | Yes |
| Network Mesh | Yes | Yes | Yes |
| **Network: Agent-to-Server** | | | |
| Server IP Address | No | Yes | Yes |
| Server Hostname | No | Yes | Yes |
| Network Grid | No | Yes | Yes |
| **Application: DNS Tests** | | | |
| DNS Server Monitor | No | Yes | No |
| DNS Server Grid | No | Yes | Yes |
| **Application: HTTP Tests** | | | |
| HTTP(S) or API | No | Yes | Yes |
| Page Load | No | Yes | Yes |
| Transaction | No | Yes | No |

> **Notes:**
>
> * Test Previews are not currently available for Autonomous tests.
> * For BGP Monitor tests, the **Preview** button on the test settings page opens the **BGP Route Viewer** rather than a Test Preview page.
> * A Test Preview never includes a **BGP** tab.
> * In a Test Preview, the **Results** tab never includes a **Sankey** subtab.

---

© 2014-25 Kentik
