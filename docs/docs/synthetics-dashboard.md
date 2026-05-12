# Synthetics Dashboard

Source: https://kb.kentik.com/docs/synthetics-dashboard

---

This article discusses the Synthetics Dashboard in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(492).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A29Z&se=2026-05-12T09%3A37%3A29Z&sr=c&sp=r&sig=U9B%2FiULgLU4Sy%2FboCOKj%2FFa%2B57UrxXNIU2kH9%2FJ0kE8%3D)

*The Synthetics Dashboard provides a high-level overview of synthetic testing.*

## Synthetics Dashboard UI

The Synthetics Dashboard — reached by clicking Synthetics in the portal's main menu — presents a summary of test status and agent deployment information for the Synthetics tests created in your organization. The dashboard includes the following main UI elements:

* **Test Status Summary**: A set of statistics on the health of tested links (see **Test Status Summary**).
* **Incident Log**: A list of tests reporting a status of Critical or Warning (see **Incident Log**).
* **Network Performance Meshes**: A user-defined set of mesh tests (see **Meshes**).
* **Recently Added Tests**: A list of your organization’s most recently added tests (see **Recently Added Tests**).
* **Agents by Status**: A breakdown by type of your organization’s currently deployed agents (see **Agents by Status**).
* **Credit Utilization**: A breakdown of your organization’s usage of synthetics test credits (see **Credit Utilization**).

## Test Status Summary

The **Test Status Summary** pane at the top of the page provides statistics on the health status of tested links (see **Synthetics Test Status**). The pane includes a set of Status indicators — Critical (red), Warning (orange), Healthy (green), and Failing (gray) — that provide a count of the tests whose status falls into each of those states.

> **Note:** The count is based on current status rather than status during the currently selected segment of the **Incident** **Timeline** (see **Incident Log UI**).

## Incident Log

The **Incident Log** lists the tests reporting a status of Critical or Warning (see **Synthetics Test Status**) during the currently selected time slice (segment) of the **Incident Timeline**. The time slice is set by hovering over the timeline. The segment you hover over will be outlined.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(493).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A29Z&se=2026-05-12T09%3A37%3A29Z&sr=c&sp=r&sig=U9B%2FiULgLU4Sy%2FboCOKj%2FFa%2B57UrxXNIU2kH9%2FJ0kE8%3D)

*The Incident Log shows tests whose status was abnormal during the currently selected time slice.*

### Incident Log UI

The Incident Log includes the following elements:

* **Time Range**: Indicates the time period covered by the timeline (see **Time Range Control**).
* **Filter**: A count of the number of filters that have been applied in the **Incident Log Configuration**.
* **Options** (vertical ellipsis): A button that pops up an **Options** menu in which you can click **Configure** to open the log's settings pane, where you can configure which incidents will be included in the **Incident** list.
* **Incident Timeline**: A visual representation of the currently selected time range, made up of segments representing each time slice in the range. By default the latest segment (far right) is selected. Hovering anywhere on the timeline has the following effects:

  + Changes the selection to the segment under the hover.
  + Changes the tests listed in the **Incident** list (see below) to those that occurred in the time slice represented by the segment under the hover.
  + Opens a popup over the timeline that gives the start time and overall status of the selected slice.
* **Incidents**: Indicates the number of Critical and Warning incidents during the currently selected time slice.
* **Time**: Indicates the starting date-time of the currently selected time slice.
* **Incident list**: A list of the tests whose status was either Critical or Warning during the currently selected time slice (see **Incident List**).

### Incident List

The **Incident** list is a table listing the tests with “incidents,” which means that the test’s status was either Critical or Warning during the currently selected time slice. The list includes the following columns:

* **Status**: The status of the test during the currently selected time slice (see **Synthetics Test Status**).
* **State**: Indicates one of the following incident states for the test:

  + Active: The test status became Critical or Warning before or during the time slice and remained so for the rest of the time slice.
  + Cleared: The test status became Critical or Warning before or during the time slice but returned to normal before the end of the time slice.
* **Test**: The name of the test that had an incident.
* **Start**: The date-time during the time slice at which the incident became Active (test status became Critical or Warning).
* **End**: The date-time, if any, during the time slice at which the test status returned to normal (the incident state became Cleared).

### Incident Log Configuration

Settings for the **Incident Log** are accessed by clicking **Configure** on the menu that pops up from the **Options** icon (vertical ellipsis) at the top right of the **Incident** **Log**. The log will be replaced with a configuration pane that includes the following settings:

* **Time Range**: Opens the **Time Range Control**, which sets the time period covered by the **Incident Log**.
* **Severity**:  A drop-down from which you can choose the status (All, Critical, or Warning; see **Test Status Levels**) of the incidents that will be shown in the **Incident Log**.
* **Status**: A drop-down from which you can choose the state (All, Active, or Cleared) of the incidents that will be shown in the **Incident Log**.
* **Incident filter fields**: A set of fields that can be used to apply filters that determine which incidents will appear in the **Incident Log** (see **Incident Filter Fields**).
* **Cancel**: A button that closes the configuration pane. All settings will be restored to their values at the time the pane was opened.
* **Save**: A button that saves changes to **Incident Log** settings and closes the configuration pane.

#### Incident Filter Fields

The incidents that will appear in the **Incident Log** can be narrowed with the incident filter fields. These multi-select fields all work as follows:

* Click in the field to open a drop-down with a list of items.
* Each item that you click in the drop-down will appear in a lozenge in the field.
* Each item in the field will be applied to filter the **Incident Log**.
* To remove a selected item from the field, click the **X** at the right of the item's lozenge.

The following filter fields are available for the **Incident Log**:

* **Test Name**: Choose the tests for which incidents will appear in the **Incident Log**.
* **Test Types**: Choose the test types (see **Synthetics Test Types**) for which incidents will appear in the **Incident** **Log**.
* **Labels**: Choose labels that have been applied to Synthetics tests in your organization; the **Incident Log** will be filtered to show only tests with the selected labels.
* **Private Agents:** Choose private agents deployed by your organization (see **Synthetics Agent Deployments**); the **Incident Log** will be filtered to show only tests using those agents.
* **Global Agents**: Choose global agents deployed by Kentik; the **Incident Log** will be filtered to show only tests using those agents.

### Time Range Control

The **Time Range** control combines two functions:

* An indicator that displays the current time range covered by the timeline.
* A button that opens a popup that enables you to define the time range.

The popup includes the following controls:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(494).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A29Z&se=2026-05-12T09%3A37%3A29Z&sr=c&sp=r&sig=U9B%2FiULgLU4Sy%2FboCOKj%2FFa%2B57UrxXNIU2kH9%2FJ0kE8%3D)**Time Range**: Two fields that each show a date-time, one for the start of the range and the other for the end. The fields are populated by clicking on the lookback list or on the calendars.
* **Apply**: A button that applies the time range from the values in the start and end fields and hides the popup.
* **Cancel**: A button that closes the popup and leaves the time range as it was before the popup was opened.
* **Lookback list**: A list of durations. Click in the list to set the duration (shown in the time range fields above) for which the time range will look back from the present.
* **Calendars**: Side-by-side monthly calendars that show the time range (if it spans more than one day) and enable you to change the range (shown in the time range fields above) by clicking on a date. The start and end days of the time range are indicated in blue, and the intervening days are indicated in light gray. The calendar controls include forward and back buttons to change the displayed months as well as drop-down selectors for month and year.

## Meshes

The **Meshes** pane shows up to eight matrices that each represents a mesh as defined in **Synthetics Test Types**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(495).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A29Z&se=2026-05-12T09%3A37%3A29Z&sr=c&sp=r&sig=U9B%2FiULgLU4Sy%2FboCOKj%2FFa%2B57UrxXNIU2kH9%2FJ0kE8%3D)

*Meshes indicate status to and from the included agents.*

This pane includes the following UI elements:

* **Options** (vertical ellipsis): A button at the upper right of the pane that pops up an Options menu in which you can click **Configure** to open the **Customize Meshes** pane, where you can set which of your organizations mesh tests will be included in the **Meshes** pane.
* **Meshes**: One or more matrices, each representing a mesh. As with any matrix, these present agents in a grid in which every agent is tested to and from every other agent. Inside the cell at each intersection of two agents there are three dots, which represent the current latency, packet loss, and jitter results, respectively, for subtests between those two agents in a single direction. Hovering over a cell highlights the cell for the reverse direction and opens a popup detailing the subtest results for both directions. Clicking the title of a mesh takes you to that mesh's **Test Details Page**.
* **Details** (opened when hovering on a cell): A popup providing current status and performance information from subtests between two agents in the mesh. The popup shows current latency, packet loss, and jitter statistics in both directions, and includes a **View Test Details** button that takes you to the corresponding **Subtest Details Page**.

#### Customize Meshes

The **Customize Meshes** pane enables you to choose which of your organization's meshes you want displayed in the **Meshes** pane. A maximum of eight meshes may be visible at once.  
  
The pane includes the following controls:

* **Meshes**: A list of the available meshes. Each item in the list includes the following:

  + Handle: Click and drag to set the order in which the meshes will be displayed.
  + Checkbox: Check the boxes of the meshes to display in the pane.
  + Name: The name of the mesh.
* **Cancel**: A button that closes the configuration pane. All settings will be restored to their values at the time the pane was opened.
* **Save**: A button that saves changes to settings and closes the configuration pane.

## Recently Added Tests

The **Recently Added Tests**pane lists the tests that were most recently defined in your organization. The list provides the following information for each test:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(496).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A29Z&se=2026-05-12T09%3A37%3A29Z&sr=c&sp=r&sig=U9B%2FiULgLU4Sy%2FboCOKj%2FFa%2B57UrxXNIU2kH9%2FJ0kE8%3D)**Status**: The current status of the agent (see **Synthetics Test Status**).
* **Name**: The name of the test. Click on the name to go to the **Test Details Page** for that test.
* **Type**: The type of test (see **Synthetics Test Types**).

## Agents by Status

The **Agents by Status**pane provides counts of various kinds of agents in use by your organization (see **About Synthetics Agents**). The kinds of agents that are listed are the same as those described in **Agents List Tabs**.

## Credit Utilization

The **Credit Utilization** pane provides information about your use of **Synthetics Test Credits**.

* **Monthly Allocation**: The total credits that are available to your organization this month (e.g., the credits-per-month purchased from Kentik in a SyntheticPak).
* **Actual Usage**: The credits your organization has actually used so far this month.
* **Projected Total Usage**: The credits Kentik expects your organization to use during this entire month based on the credits consumed by the tests that are currently active.
* **Projected Remaining**: The credits from your organization's allocation for this month that will go unused if your actual usage matches your projected usage.

> **Note:** To acquire additional test credits, contact Kentik Customer Success (see **Customer Care**).

---

© 2014-25 Kentik
