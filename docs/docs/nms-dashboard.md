# NMS Dashboard

Source: https://kb.kentik.com/docs/nms-dashboard

---

The NMS Dashboard of the Kentik portal is covered here.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(408).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A02Z&se=2026-05-12T09%3A34%3A02Z&sr=c&sp=r&sig=EKeZ%2B3I02iLjjwnRt5GPLGuKX7qGawAbwK%2B1%2FqDEzk4%3D)

*The NMS Dashboard provides a metrics-based overview of your network activity and state.*

> **Note:** The following articles provide more on other aspects of Kentik's Network Monitoring System:
>
> * **Network Monitoring**
> * **Kentik NMS Configuration**
> * **NMS Setup**
> * **Metrics Explorer**
> * **NMS Devices**
> * **NMS Interfaces**
> * **Query Assistant**

## About the NMS Dashboard

The Kentik portal's NMS Dashboard — reached by clicking **Network Monitoring System** in the portal's main menu — is a user-configurable high-level view of your organization's NMS implementation. The core of this page is the display area where visualizations (a.k.a. "widgets") are displayed. Each visualization focuses on a specific aspect of your NMS-reported traffic, performance, status, connectivity, or utilization. You can customize the display in terms of both content and layout to provide yourself with the information you'd most like to see, and you can save your changes so the displayed information is specific to your individual needs. All changes are unique to you as a user (distinct from the others in your organization).

## NMS Dashboard Structure

The NMS Dashboard includes the following main UI elements:

* **Title bar**: Running the full width of the page under the main navbar, the title bar contains page-wide controls (see **NMS Dashboard Title Bar**).
* **Display area**: The area beneath the title bar where the individual visualizations are displayed (see **NMS Dashboard Display Area**).
* **Add Visualizations**: A drawer that opens from the bottom of the screen and enables you to choose more visualizations to display (see **Add Visualizations Drawer**).

## NMS Dashboard Title Bar

The title bar of the NMS Dashboard runs across the top of the page and includes the following controls:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(410).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A02Z&se=2026-05-12T09%3A34%3A02Z&sr=c&sp=r&sig=EKeZ%2B3I02iLjjwnRt5GPLGuKX7qGawAbwK%2B1%2FqDEzk4%3D)**Add Visualizations**: A button that opens the **Add Visualizations Drawer** from the bottom of the screen, which enables you to choose additional visualizations to display.
* **Save Changes** (present only when changes to the display area haven't been saved): A button that saves the current state of the display area, including the collection of visualizations and their current layout.

> **Note:** Changes to the default set and layout of visualizations are user-specific; they don't affect other users in your organization.

## NMS Dashboard Display Area

The main display area of the NMS Dashboard includes a set of visualizations (a.k.a. "widgets") that show different aspects of your organization's network activity.

### Set Visualizations

The NMS Dashboard comes stocked with a standard set of visualizations in the display area, but you can change the visualizations that are shown by removing existing visualizations or adding new ones.

#### Add a Visualization

To add a visualization to the display area:

1. Click the **Add Visualizations** button
2. In the resulting**Add Visualizations Drawer**, click the tab that best describes the category of the visualization you're interested in.
3. On the resulting tab, click the card for the visualization you'd like to add. The drawer will close and the visualization will appear in the display area.

   > **Note:** To alter the location or size of the new widget, see **Change Visualization Layout**.

#### Remove a Visualization

To remove a visualization:

1. Click the **Action** button (vertical ellipses) at the upper right of the visualization.
2. Choose **Remove** from the resulting menu.

> **Note:** There's no confirmation dialog before the visualization is removed.

### Change Visualization Layout

The following visualization-level controls are available to change the appearance of visualizations in the display area (the other visualizations will rearrange to accommodate the change):

* **Reposition**: To move a visualization, click and drag from its title bar.
* **Resize**: To resize a visualization, click and drag from one of the handles that appear in the bottom corners when you hover on it.

### Visualization Action Menu

Each visualization has a vertical ellipses at its upper right that displays an Action menu. Depending on the visualization, you may see the following options:

* **Configure**: Modify how a visualization’s information is displayed (e.g. title, lookback, or group by). Once you’ve updated your preferences, click **Save** to close the configure options and click **Save Changes** at the top of the page to update the page.
* **View in Metrics Explorer**: Go to **Metrics Explorer** (in the same tab), where the **Query** sidebar will be set to the same metrics and group by dimensions used in the visualization.
* **View in Data Explorer**: Go to **Data Explorer** (in the same tab), where the **Query** sidebar will be set to show the traffic displayed in the visualization.
* **Remove**: Remove the visualization from the NMS Dashboard. Click **Save Changes** to update the page.

### Default Visualizations

Kentik users who visit the NMS Dashboard for the first time will see the default set of visualizations, which includes the following:

* **Health Map by Sites**: A map and list of your organization’s sites that details the health status of each (see **Health Map by Sites**).
* **Device Availability**: An area chart and a table listing your devices and their availability (see **Device Availability**).
* **Traffic Overview**: Visualizations, each on a separate tab, of total traffic volume and the volume of traffic matching each of the following traffic profiles (see **Simple Traffic Profile**).
* **Active Alerts**: A ring chart and table that show currently active alerts from NMS policies (see **Active Alerts**).
* **Interfaces with Highest Utilization**: A table that shows the top interfaces (and their device names) ranked by their last percentages of in and out utilization.
* **Devices with Highest CPU**: A table that shows the top 50 devices in your organization that have the highest CPU utilization.
* **Devices with Highest Memory**: A table that shows the top 10 devices in your organization that have the highest percentage of memory utilization.

> **Note:** In visualizations that include a table with a Device column, clicking a device name will take you to an **NMS Device Details Page** for that device.

#### Health Map by Sites

This visualization is made up of the following main parts:

* **Map**: Each site is represented as a dot.

  + A red dot indicates one or more alert; a green dot indicates none.
  + Hover over a dot to see a popup with details about the site’s active alerts. Clicking the site name at the top of the popup will take you to the portal's **Alerting** page, with the Alerts list filtered to show only alerts related to that site.
* **Site list**: A list of your organization’s sites and the number of currently active alerts at each.

> **Note**: For each policy whose alerts are to be covered by a Health Map by Sites visualization, the site must be specified as a Policy Dimension (see Dimensions in **Data Funneling**), and each device covered by the policy must be assigned to a site (see **Device General Settings**).

#### Device Availability

This visualization is made up of the following main parts:

* **Status**: Device’s current status (Down, Unknown, or Up).
* **Device name**: Name of the device and a link to access its NMS Device page. Clicking a device name will take you to an **NMS Device Details Page** for that device.
* **Overall**: The percentage at which the device has been available over the last 7, 14, or 30 days.
* **Segment**: If you’ve zoomed into the visualization (click and drag), this column provides the percentage at which the device has been available during that period of time.

#### Active Alerts

This visualization includes both a ring chart and a table. The ring chart shows the breakdown of the active alerts either by policy name (default) or by severity (set via Configure on the **Visualization Action Menu**).  
  
The table includes the following columns:

* **Policy**: Name of the policy that was triggered.
* **Dimensions**: Dimensions that triggered the alert.
* **Metric**: The metric/condition that triggered the alert.
* **Duration**: Length of time for which the alert has been active.

## Add Visualizations Drawer

The **Add Visualizations** drawer opens from the bottom of the screen when you click the corresponding button in the title bar. The drawer enables you to choose visualizations that you'd like to add to the display area.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(411).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A02Z&se=2026-05-12T09%3A34%3A02Z&sr=c&sp=r&sig=EKeZ%2B3I02iLjjwnRt5GPLGuKX7qGawAbwK%2B1%2FqDEzk4%3D)

*Available visualizations are categorized into tabs in the drawer.*

The visualizations that are available to add to the display area are grouped into the following tabs:

* **Alerting**: Visualizations related to alerts generated by alert policies.

  > **Note:** For each policy whose alerts are to be covered by a **Health Map by Sites** visualization:
  >
  > + Site must be specified as a Policy Dimension (see **Dimensions** in **Data Funneling**).
  > + Each device covered by the policy must be assigned to a site (see **Device General Settings**).
* **Cloud**: Visualizations related to traffic to and from your resources that are housed in public cloud providers (see **Cloud Overview**).
* **Core**: Visualizations drawn from the Core section of the Kentik portal (see **Core**), including widgets based on individual dimensions (IPs, ASNs, etc.) or broken down by data sources (devices, clouds, interfaces).
* **Edge**: Visualizations derived from Edge modules including **Connectivity Costs** and **Discover Peers**.
* **Insights**: Visualizations based on insights from Kentik's **Insights & Alerting** system.
* **NMS**: Visualizations related to Kentik’s Network Monitoring System.
* **Protect**: Information on DDoS attacks and mitigations, derived from the **DDoS Defense** module.
* **Service Provider**: Visualizations derived from Service Provider modules including **CDN Analytics**, **OTT Service Tracking**, and KMI Insights (see **Top Insights**).
* **Synthetics**: Visualizations related to synthetic testing configured in the **Test Control Center** using global and privately-deployed Kentik `ksynth` agents.
