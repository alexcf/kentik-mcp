# Agent Management

Source: https://kb.kentik.com/docs/agent-management

---

This article discusses the Synthetics » **Agents** page of the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(541).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A24Z&se=2026-05-12T09%3A48%3A24Z&sr=c&sp=r&sig=pKUuZFJH3UbGc6%2FuHFBI4REEpKNb9uEyhjM25L2kvjI%3D)

*The Agents page enables agent deployment and monitoring of Synthetics agent status.*

## About Agent Management

Kentik’s synthetic testing capabilities are enabled by `ksynth-agent`, a Kentik software agent that can be deployed within your infrastructure as well as elsewhere around the Internet. `ksynth-agent` can be used in either public ("global") or private deployments (see **Synthetics Agent Deployments**).  
  
The **Agents** page lets you view and manage all currently deployed `ksynth-agent` instances, both public and private. Here are some of the features:

* The location of the agents, filterable by type, is shown on the **Agents Map**.
* A filterable **Agents List** provides information about each agent and allows you to activate an agent, edit its configuration, or remove it.
* Add private agents using the **Deploy Private Agent** pane across the top.

## Agent Status

A deployed agent can be in one of the following states, which are indicated by status icons at various places in the Synthetics UI:

| Status | Description |
| --- | --- |
| Active (green checkmark circle) | The agent is activated and current. |
| Pending | The agent has been installed but not yet registered or activated (see **Activate Agent**). |
| Warning or Offline (orange triangle) | Kentik has never received test data from this agent, or the agent has been offline for between five minutes and one hour. |
| Critical or Offline (red caution hexagon) | Either Kentik has never received test data from this agent, or the agent has been offline (no results received) for 60 minutes or more. |

## Agents Page

Synthetics agents are managed in the Kentik portal on the **Agents** page, which you’ll find under Synthetics in the main navbar. The page includes the following main UI elements:

* **Export** (on subnav): A drop-down menu with two options for sharing agent details:

  + **Visual Report***:* Export a PDF of the current state of the page. Kentik will prepare the report and display a notification indicating that the report is ready. When you click the provided link, your browser will handle the download to your local machine.
* **Share** (on subnav): Share the page via link or email (see **Sharing via the Share Dialog**).
* **Deploy a Private Agent**: A button that expands into a pane in which you can activate a new private agent (see **Private Agent Deployment**).
* **Agents map**: Shows the physical location of `ksynth-agent` instances, both public and private, throughout the world (see **Agents Map**).
* **Agents list**: A tabbed table providing information about your `ksynth-agent` instances (see **Agents List**).
* **Filters pane**: A set of filters that determine what’s displayed in the Agents list (see **Agents List Filters**).
* **Agent Details pane**: A sidebar that slides out from the right edge of the page and provides details on an individual agent (see **Agent Details Sidebar**).

## Private Agent Deployment

While public/global agents are deployed by Kentik for use by all customers, private agents are deployed by each customer for their own exclusive use (see **Synthetics Agent Deployments**). Deploying a private agent involves first downloading and installing the desired `ksynth-agent` instance (see **Synthetics Agent Types**), then activating it from the **Deploy A Private Agent** pane. If the pane is collapsed, you can expand it by clicking the **Deploy a Private Agent** button at the top right of the Agents page.  
  
The expanded pane includes the following buttons:

* **View Setup Procedures**: Opens the Private Agent Setup dialog (see **Private Agent Setup**).
* **Enter Challenge Code**: Opens the Register Agent dialog (see **Enter Challenge Code**).

#### Private Agent Setup

![Instructions for setting up a private agent using Docker with command examples.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AM-private-agent-setup-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A24Z&se=2026-05-12T09%3A48%3A24Z&sr=c&sp=r&sig=pKUuZFJH3UbGc6%2FuHFBI4REEpKNb9uEyhjM25L2kvjI%3D)

The Private Agent Setup dialog provides setup procedures for deploying `ksynth-agent` via package or Docker:

1. Choose an app agent type (network or app).
2. Choose a deployment option (Debian/Ubuntu package, RPM package, or Docker)
3. Follow the on-screen instructions for installation and registration. (Same instructions are also in the KB in **Synthetics Agent**).

#### Enter Challenge Code

Depending on your `ksynth-agent` installation process, you may need to register each agent individually in the portal in order for it to be activated. To do so, you get a challenge code from the agent and enter the code on the Agents page by clicking the **Enter Challenge Code** button, which opens the Register Agent dialog. When you enter the code in the dialog and click the Register button, the agent instance corresponding to the entered code will be automatically registered and will appear in the **Private Agents** tab of the **Agents List**.

> **TIP**: For specific information about activation, see **Activate Agent**.

## Agents Map

The Agents map shows the physical location of `ksynth-agent` instances, both public agents and private, throughout the world. Hovering on an agent opens a popup with identifying information about that agent. Scrolling or double-clicking on any point in the map zooms the display into the clicked area.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(542).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A24Z&se=2026-05-12T09%3A48%3A24Z&sr=c&sp=r&sig=pKUuZFJH3UbGc6%2FuHFBI4REEpKNb9uEyhjM25L2kvjI%3D)

*The Agents Map shows the location of global and private* `ksynth-agent` *instances throughout the world.*

The Agents map is filterable using the checkboxes in the Markers pane, which enables you to show or hide agents by type. Within the global agents type, you have the option of showing or hiding sets of public cloud agents based on the cloud provider. When all cloud providers are unchecked, the remaining global agents are those that Kentik has installed in data centers in key Internet hubs worldwide (Hosted agents). As various types of agents are shown or hidden the map will zoom automatically to include all of the agents that are still shown.

## Agents List

The Agents list, a tabbed table providing information about `ksynth-agent` instances (both public and private) is covered in the following topics.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(543).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A24Z&se=2026-05-12T09%3A48%3A24Z&sr=c&sp=r&sig=pKUuZFJH3UbGc6%2FuHFBI4REEpKNb9uEyhjM25L2kvjI%3D)

*The Agents list includes tabs showing private agents and global agents, including those hosted by key cloud providers.*

### Agents List Labels

Kentik's labeling feature enables you to create a label (essentially a property whose value is text) and apply it to one or more of your synthetic tests, creating a group of tests that can be referred to (e.g., filtered for) collectively rather than individually. The controls for this feature of the Agents list are the same as for the **Label Controls** in the Test Control Center.

### Agents List Tabs

The Agents list tabs are included for each of the following types of agents:

* **Pending Agents**: The agents in your Kentik system that have not yet been registered. Present only if you have pending agents.
* **Private Agents**: Every Kentik customer can deploy as many instances of `ksynth-agent` as they care to in their own on-prem and/or cloud infrastructure (no additional license required). These private agents are for the exclusive use of the customer who deploys them (not available to other Kentik customers).
* **Global Agents**: All agents in the Kentik Global Agent Network, a worldwide network of Kentik-maintained `ksynth-agent` instances (hosted in data centers and public clouds) that enables performance testing to and from key Internet hubs worldwide.
* **App Agents**: A subset of our Kentik-maintained global agents, app agents run **Headless Chromium**to perform advanced web-layer tests, including full browser page load.
* **Broadband Agents**: A subset of our Kentik-maintained global agents, broadband agents are connected to broadband service providers to provide a true "end-user view" of network performance. Many are running in real end-user homes and as such may be less reliable than our Global and Public Cloud agents, which run in data centers or public cloud regions.
* **Public Cloud Agents**: A subset of global agents, showing only agents deployed in the infrastructure of key cloud service providers (AWS, GCP, Azure, OCI, etc.).

### Agents List Columns

The columns of the Agents list, which vary depending on the tab of the table, provide the following information for each of the above types of agents:

* **Name**: The agent’s name, as well as its type, status, public IP address, and any assigned labels (see **Agents List Name Column**).
* **Agent Type**: Displays whether the agent is an “App” or “Network” (see **Synthetics Agent Types**).

  > **Note:** The word “type” as applied to an agent can have different meanings depending on the context (see **Agent Type**).
* **Location**: The location where the agent is deployed, typically either a datacenter or a cloud region.
* **IP Address**: The public IP address assigned to the agent.
* **Private IP**: The private (local) IP address, if any, assigned to the agent in the **Configure Agent Dialog**.
* **ASN**: The name and number of the Autonomous System in which this agent is located.
* **Region**: The cloud provider region (e.g., "US West") in which the agent is located.
* **Tests**: The number of tests currently defined in your organization for this agent.
* **IP Version**: The IP version(s) tested by the agent.
* **Version**: The `ksynth-agent` version number of the agent (green if latest, orange if out of date).

The table below shows which columns appear in the tab for each `ksynth-agent` deployment type.

| Columns | Pending Agents | Private Agents | Global Agents | App Agents | Broadband Agents | Public Cloud Agents |
| --- | --- | --- | --- | --- | --- | --- |
| Name | Yes | Yes | Yes | Yes | Yes | Yes |
| Agent Type | Yes | Yes | No | No | No | No |
| Location | No | Yes | No | No | No | No |
| IP Address | Yes | No | No | No | No | No |
| Private IP | No | Yes | No | No | No | No |
| ASN | Yes | No | Yes | Yes | Yes | No |
| Region | No | No | No | No | No | Yes |
| Tests | No | Yes | Yes | Yes | Yes | Yes |
| IP Version | Yes | Yes | Yes | Yes | Yes | Yes |
| Version | Yes | Yes | No | No | No | No |

#### Agents List Name Column

Each cell in the **Name** column of the Agents list includes several pieces of information:

* **Agent type**: An indicator that shows the **Agent Type**. See the tab labels of the Agents list to see what each icon looks like.
* **Name**: For public agents, this is the Kentik-assigned name. For private agents, this is by default the name of the agent's host, but you can assign a different name with the `ksynth-agent` CLI (see **Agent CLI Arguments**).
* **Status**: The current **Agent Status** displayed as an icon.
* **Public IP Address**: The IP address for all synthetic traffic on this agent.
* **Labels**: Any labels assigned to the agent in the **Configure Agent Dialog** (private agents) or directly on the **Agents** list (all agents).

#### Agent Type

In both the Kentik UI and this KB, the term “agent type” may, depending on context, refer to either of the following:

* **Capability**: The **Agent Type** column in the **Agents** list refers to the capability of the agent, either App or Network (see **Synthetics Agent Types**).
* **Deployment**: In the context of deployment, the agent’s type corresponds to its “location” or network type, e.g. Private, Global, Broadband, or Public Cloud (see **Agents List Tabs**).

### Agents List Filters

The **Filters** pane to the right of the Agents list enables you to filter the list using the following fields and controls:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(544).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A24Z&se=2026-05-12T09%3A48%3A24Z&sr=c&sp=r&sig=pKUuZFJH3UbGc6%2FuHFBI4REEpKNb9uEyhjM25L2kvjI%3D)**Reset To Default**: Resets filters to their default setting. Appears only when filters are currently applied.
* **Search** (magnifying glass icon): Filters the Agents list to show only those agents whose name, IP address, or site name contains the entered string.
* **ASN**: A drop-down from which you can choose an ASN for which agents will be shown in the Agents list. Multiple ASNs can be selected.
* **Region**: A drop-down from which you can choose a region for which agents will be shown in the Agents list. Multiple regions can be selected.
* **Country**: A drop-down from which you can choose a country for which agents will be shown in the Agents list. Multiple countries can be selected.
* **Labels**: A drop-down list of labels from which you can choose (or remove) one or more labels.
* **IP Version**: A drop-down list of IP versions from which you can choose (or remove) one or more version.

## Agent Details Sidebar

The **Agent Details** sidebar, which appears when you click on an agent in the **Agents List**, is covered in the following topics.

### Details Sidebar UI

The **Agent Details** sidebar includes the following main UI elements:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(545).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A24Z&se=2026-05-12T09%3A48%3A24Z&sr=c&sp=r&sig=pKUuZFJH3UbGc6%2FuHFBI4REEpKNb9uEyhjM25L2kvjI%3D)**Agent ID**: The Kentik-assigned unique ID number of the agent.
* **Expand/Collapse**: A button that takes you to a full-screen view of the agent details provided in the sidebar. To return to the Agents page, use your browser’s back button.
* **Agent type**: An indicator showing if the agent is global (Kentik icon) or private (private agent icon).
* **Agent name**: The name of the agent, either globally (public agents) or within your organization (private agents).
* **Activate** (only for a private agent that isn’t yet activated): A button that opens the Activate Agent dialog, which is the same as the **Configure Agent Dialog**.
* **Configure** (activated private agents only): A button that opens the **Configure Agent Dialog**.
* **Remove** (activated private agents only): A button that opens a confirming dialog allowing you to remove the agent from your organization’s collection of deployed agents.
* **Agent tab**: Details about the agent (see **Agent Details Tab**).
* **Tests tab**: A searchable list of the agent’s currently active tests (see **Agent Tests Tab**).
* **Alerts tab**: A bar chart and table that detail the agent’s alerts for the selected timeframe (see **Agent Alerts Tab**).
* **Downtime tab**: A bar chart and table that detail the agent’s errors and downtime for the selected timeframe (see **Agent Downtime Tab**).

> ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(546).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A24Z&se=2026-05-12T09%3A48%3A24Z&sr=c&sp=r&sig=pKUuZFJH3UbGc6%2FuHFBI4REEpKNb9uEyhjM25L2kvjI%3D)**Note**:If the agent whose details you’re viewing is on the**Pending Agents** tab of the **Agents List** and has not yet been registered, you will see this blue information box above the sidebar tabs asking you to finish registration.

### Agent Details Tab

The **Agent** tab of the Agent Detailssidebar provides the following information related to a `ksynth-agent` instance:

* **Version** (private agents only): The version number of this agent.
* **IP Address**: The agent’s public IP address and its private IP (if any has been assigned to a private agent via the **Private IP** field in the **Configure Agent Dialog**).
* **Host OS**: (private agents): The type (e.g., Linux Ubuntu) and version of the server on which the agent is deployed.
* **ASN**: The name and number of the AS in which the agent is deployed.
* **Activated**: The date on which the agent was activated (see **Private Agent Deployment**) and the Kentik user who activated it.
* **Location**: The location of the agent:

  + **Private agents**: Name and location (city, country) of the site where the agent is deployed.
  + **Global agents**: Name (set by Kentik) and, if applicable, cloud region where the agent is deployed.

    > **Note**: A map showing the agent’s location displays below the agent’s information.

### Agent Tests Tab

The **Tests** tab of the Agent Details sidebar provides a list of the currently active tests involving this agent.

It includes the following UI elements:

* **Search**: A field into which you can enter text to narrow the listed tests.
* **Test**: A column heading that sorts the test type groups to either ascending or descending order.
* **Test type group**: Expandable/collapsible groups that contain tests of that type (see **Synthetics Test Types**).
* **Test name**: The name of the synthetic test to which this agent is assigned. Click on a test to go to the **Test Details Page** for that test in the Test Control Center.

### Agent Alerts Tab

The **Alerts** tab of the Agent Details sidebar provides details about any alerts that were triggered by tests run on the selected agent. You can use the **Time Range**, **Status**, and **State** controls to filter which alerts are shown in the bar chart andAgent Alerts log.

#### Alerts Tab Settings

Across the top of the **Alerts** tab is a control set that specifies what is to be shown in the chart and the table below:

* **Time Range**: Opens the **Time Range Control**, which sets the time period for which information is displayed (default: Last 1 Day).
* **Status**: A drop-down with which you can choose the severity of the alerts to show: All, Warning, or Critical.
* **State**: A drop-down with which you can choose the state of the alerts to show: All, Active, or Cleared.
* **Cancel** (x): A button that resets the **Time Range**, **Status**, and/or **State** drop-downs to their last saved settings.
* **Apply** (checkmark): A button that saves the current **Time Range**, **Status**, and **State** settings and applies them to the chart.

> **Note:** Changes made with the above controls are not visible in the chart and table until applied with the **Apply** button.

#### Alerts Tab Displays

The following elements show information about alerts that correspond to the settings made with the **Alerts Tab Settings**:

* **Chart**: A bar chart in which each segment represents one time slice within the time span specified with the **Time Range** selector. The bars shown in each slice are determined by the applied **Status** and **State** settings. The number of alerts is measured on the vertical axis and time is measured on the horizontal axis.
* **Incidents**: A count of the alerts reported during the most recent time slice (default) or the time slice last selected (hovered-over) in the chart.
* **Time** (UTC): The start time of the most recent time slice (default) or the time slice last selected in the chart.
* **Alerts log**: A table listing the alerts reported during the most recent time slice (default) or the time slice last selected in the chart (see **Agent Alerts Log**).

#### Agent Alerts Log

The **Agent Alerts** log provides information about any alerts triggered during the currently selected (e.g., last hovered over) time slice in the chart. The columns include:

* **Status**: An icon representing the severity of the alert:

  + **Orange dot**: Warning alert
  + **Red dot**: Critical alert
* **State**: An icon representing the state of the alert:

  + **Green checkmark**: The alert has been cleared.
  + **Red bell**: The alert is active.
* **Test**: The name of the test for which an alert was triggered. Click the test name to go to that test’s **Test Details Page**.
* **Test type**: Appears in parentheses to the right of the alert’s test name.
* **Duration**: The start and end date-time of the alert, if applicable.

### Agent Downtime Tab

The **Downtime** tab of the Agent Details sidebar provides details for any issues this agent experiences when Status Alerting and Notifications is enabled in the **Configure Agent Dialog** (see Downtime under **Agent Property Settings**).

#### Downtime Tab UI

The **Downtime** tab includes the following UI elements:

* **Time Range**: Opens the **Time Range Control**, which sets the time period covered by the Alerts tab (default: Last 6 hours).
* **Chart**: A bar chart that displays all errors during the selected timeframe. The number of errors is measured on the vertical axis and time is measured on the horizontal axis.
* **Selected Time** (Local): The date/time that is selected on the chart, the data for which is displayed in the **Downtime** log. The default is the most recent 5-minute increment.
* **Errors**: A summary of the errors reported during the last time frame selected (hovered over) in the chart.
* **Downtime log**: The list of errors that occurred during the timeframe selected on the chart. The columns of the list include:

  + **Error**: The type of error reported (i.e. Agent offline).
  + **Timestamp**: The start and end time (if applicable) of the error. Click the column heading to reorder the errors in ascending or descending order.

## Configure Agent Dialog

The Configure Agent dialog and the nearly identical Activate Agent dialog enable you to set certain properties of a private `ksynth-agent` instance.

> **Note:** This dialog is accessed via the **Configure** button in the **Agent Details Sidebar**.

### Configure Agent UI

The Configure Agent dialog includes the following UI elements:

* **Properties**: The fields and controls that determine the basic properties of the agent (see **Agent Property Settings**).
* **Cancel**: A button — **X** at top right or **Cancel** at bottom — that closes the dialog without saving changes. All elements will be restored to their values at the time the dialog was opened.
* **Save**: A button that saves changes to agent settings and exits the dialog.

### Agent Property Settings

The following fields and controls set the basic properties of a private `ksynth-agent` instance:

* **Name**: Your organization's name for the agent. By default, this is the name of the agent's host.

  > **Note:** This property can also be set at installation via the `ksynth-agent` CLI (see **Agent Arguments**).
* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(547).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A24Z&se=2026-05-12T09%3A48%3A24Z&sr=c&sp=r&sig=pKUuZFJH3UbGc6%2FuHFBI4REEpKNb9uEyhjM25L2kvjI%3D)**Site**: A control set for associating the agent with a site (see **About Sites**), which would typically be the location of the `ksynth-agent` instance’s host, whether that’s a data center or a VPC in a cloud provider):

  + **Site**: A filterable drop-down from which to choose a site.
  + **Create a New Site**: A button that opens additional controls (see **Create a New Site Controls**) with which you can create a new site if a site hasn't yet been defined for the location of the agent host.
* **Cloud Provider**: A drop-down from which to choose a cloud provider (e.g., AWS, GCP, Azure) if the agent is hosted in a cloud resource (otherwise leave as "None"). Choosing a provider shows the following additional controls, which are otherwise hidden:

  + **Cloud Region**: A filterable drop-down from which to choose the region within the specified cloud provider where the agent is hosted.
  + **Cloud Virtual Network**:  A field in which to enter a description of the VPC in which the agent is deployed.
* **Private IP(s)**: Fields into which you can enter an optional private (local) v4 and/or v6 IP address for the agent. A private IP allows you to establish agent-to-agent or agent mesh tests using your internal network via the **Use Agent Private IPs** setting in the **Target and Agents** tab of a test's settings page:

  + **Private IPv4**: Shown when **IP Versions Supported** is set to **v4 only** or **v4 + v6**.
  + **Private IPv6**: Shown when **IP Versions Supported** is set to **v6 only** or **v4 + v6**.
* **IP Versions Supported**: A drop-down to choose the type of IP addresses that will be supported by the agent: IPv4 only, IPv6 only, or both (default). This is particularly useful for agents that will be used for tests of type Hostname or ASN, where there could be a mix of IPv4 and IPv6 addresses.
* **Labels**: A field showing the labels, if any, currently applied to this agent. Click in the field to drop down a filterable list from which you can choose one or more labels to apply (see **Agents List Labels**). The applied labels will be added to the field. To remove a label, click on the **X** at its right.
* **Add/Edit**: A link that takes you to the settings page for **Labels**, where you can create a label to apply to this agent.
* **Status Alerting and Notifications**: A switch that shows the **Agent Alerting and Notifications** controls used to enable and configure alerting and notifications for this agent.
* **Maintenance Mode**: A switch that shows the **Maintenance Mode** controls used to enable and configure a pause in data collection and alerts for tests that use this agent as well as a pause in the agent’s own status alerts.

#### Create a New Site Controls

The following controls, shown when you click the**Create a New Site**button, are used to enter the information required for the new site:

* **Site Name**: A field for the name of the new site.
* **Street Address**: A field for the street address of the physical location of the site.
* **Use an Existing Site**: A button that will clear and hide the above fields if you decide to choose a site from the **Site** drop-down instead of creating a new site.

#### Agent Alerting and Notifications

The following controls, shown only when the**Status Alerting and Notifications** switch is On, configure alerting and notification for the agent, including the amount of downtime that will generate a notification and the channel(s) to which it will be sent:

* **Downtime**: Set the length of continuous downtime that will result in a notification (default = 5 minutes):

  + **Number of units**: A field into which you can enter the number of time units.
  + **Duration of units**: A drop-down from which you can select the time unit, either a minute (default) or a second.
* **Notification Channel** (Optional): A field showing the notification channels, if any, currently applied to this agent. Click in the field to drop down a filterable list from which you can choose one or more channels to use for notifications. The applied channels will be added to the field. To remove a channel, click on the **X** at its right.
* **Add New**: Opens the Add Notification Channel dialog, where you can configure a new notification channel in your organization (see **Notification Settings**). When you click **Add Notification Channel** in that dialog the new channel will be automatically added to the agent.

#### Maintenance Mode

When enabled, the **Maintenance Mode** switch in the **Configure Agent Dialog** pauses agent status alerts, and data collection and alerts for tests that use this agent. Once enabled, the following controls are available:

* **Maintenance Window**: A drop-down that shows the date-time range of the maintenance window. Click to pop up the **Maintenance Window Controls**.
* **Start Time** (present only when **Never Expire** is on): A field stating the date-time at which maintenance mode will start. Click in the field to edit its text or choose a start date from the calendar popup.
* **Never Expire**: A switch (default = off) that, when on, sets the maintenance mode to continue indefinitely (no ending date-time).
* **Expires in**: A lozenge that indicates when (or if) maintenance mode is set to expire.

#### Maintenance Window Controls

The following controls open by clicking the **Maintenance Window** field:

* **Start and End Date**: Input fields that show the start and end date-times of the maintenance window, which may be entered either using the calendar or directly in the datetime format YYYY-MM-DD HH:MM.
* **Calendar**: A popup in which you can click to set the start and end date-times.
* **Apply**: A button to set the maintenance window to the date-times specified in the **Start Date** and **End Date** fields. This is only visible if the date-time is modified after opening the dialog.
* **Cancel**: A button to leave the maintenance window as it was before the controls were opened. This is only visible if the date-time is modified after opening the dialog.
