# Insights

Source: https://kb.kentik.com/docs/insights

---

This article covers the Insights feature of the Kentik portal.

> **Note:** For descriptions of current insights, see **Insights Reference**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(472).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A43Z&se=2026-05-12T09%3A44%3A43Z&sr=c&sp=r&sig=oLrlU3jiTK52leyrvpThJbtR4YSdKBCMzfxMkIHUAcA%3D)

*The Insights page lists recent insights and enables insight configuration.*

## About Insights

Insights are automatic real-time notices from Kentik that tell you about interesting or anomalous traffic behaviors. An insight can be any fact that is relevant to the traffic on the page and potentially of interest to the network operator. It could be a vital alarm, such as a broken link or a bad interface card, or it could be something that seems out of the ordinary, such as unusual traffic across a backbone or a security violation.

Insights are intended to be actionable whenever possible. Building on the baselining capability of our alerting system, a potentially noteworthy current situation can be correlated with a wealth of historical data and trends.

### Accessing Insights

Insights may be accessed in a couple different ways within the Kentik portal:

* **Insights page**: The **Insights List** on the page at Core » **Insights** is a filterable table showing all recent insights. Click on a row in the table to open the **Insight Details Page** for a specific insight.
* **Insights pane**: The **Insights Pane** pops up from the **Insights** button at the right of the subnav on the **Network Explorer** page and its associated **Core Aggregate Pages** and **Core Details Pages**. The pane contains insight cards related to the traffic presented on that page. For example, on the aggregate page for devices, the insights relate to all devices in your network, while on a detail view for an individual device the insights relate to only that device.
* **Observation Deck**: The **Insights Overview** visualization may optionally be added to your **Observation Deck**. Insights for the configured period are displayed as cards within collapsible panes by insight family (see **Insights Families**).

  > **Note:**
  >
  > + Click **Configure** in the Actions menu to change the Lookback setting from the default of 1 week. Select either 1 day, 2 days, 1 week, or 2 weeks from the drop-down, then click **Save**.
  > + Click a pane to expand a family of insights, then click a card to show more information for an individual insight (see **Insights Pane**). Only one pane/card in the **Insights Overview** can be expanded at a time.

## Insights Families

Insights are categorized into "families." A given insight's family is indicated both in the **Insights** list on the **Insights Page** and in the **Details** sidebar of each **Insight Details Page**.  
  
 Current insights families include:

* **Edge**: Changes to traffic volume through the edge of your infrastructure, by amount or by ranking, primarily evaluated based on ASNs and sites.
* **Geolocation**: Changes in top-X rankings in traffic volume as evaluated for geo-related dimensions such as cities, regions, countries, etc.
* **Hybrid Cloud**: Changes in top-X rankings in traffic volume as evaluated for cloud-related dimensions such as cloud service, cloud region, subnet, etc.
* **Network Health**: Notification of possible problems on your network such as unexpectedly low or high traffic bitrates or flows per second.
* **Peering Analytics**: Identify changes to peering connections such as low IX utilization or ASN traffic shifts.
* **Protect**: Notification of possible threats to network availability.
* **Security**: Detect vulnerabilities such as traffic transmitted to embargoed countries.
* **Traffic**: Changes in top-X rankings in traffic volume as evaluated for traffic-related dimensions such as applications, ASNs, protocols, etc.
* **Traffic Comparison**: Identify differences in traffic observed by NetFlow vs. SNMP.

## Insights Pane

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(473).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A43Z&se=2026-05-12T09%3A44%3A43Z&sr=c&sp=r&sig=oLrlU3jiTK52leyrvpThJbtR4YSdKBCMzfxMkIHUAcA%3D)

The popup **Insights** pane opens from the **Insights** button on the subnav of Network Explorer and associated Core Aggregate and Core Details pages. This pane contains a vertical series of cards, one for each insight. The following UI elements are initially visible in each card:

* **Family**: The icon at the top left of the card indicates the family (category) of the insight (see **Insights Families**).
* **Description**: A summary of the specific situation leading to the generation of this insight, including links to **Core Details Pages** for devices when applicable.
* **Time**: An indicator stating how long ago the insight was originally posted.
* **Chart**: A chart showing traffic related to the insight. This inline chart enables you to get a quick feel for the situation without leaving the current page. Currently supported chart types include:

  + Stacked area and line charts when comparing time series data.
  + Horizontal bar graph when comparing two or more total values (such as total traffic this week vs. total traffic last week).
* **View Details**: A link that takes you to an **Insight Details Page** where you can see additional details about the insight.
* **View in Data Explorer**: A link that takes you to Data Explorer, where the controls will be set to show the traffic that generated the insight.
* **Feedback**: A set of buttons (thumbs up/down icons) with which you can rate the value of the insight (see **Feedback Controls**).
* **More Insights**: A button, present at the bottom of the pane when there are additional insights to display, that displays those insights.

> **Note:** Along with insights, the Insights pane will also include cards for alerts, if any, from Kentik's Alerting system.

## Insights Page

The **Insights** page provides a centralized list of your organization's insights, which you can filter according to factors such as type and impact. The page includes the following UI elements:

* **Configure Insights**: A button that takes you to the **Insights Settings Page**.
* **Show/Hide Filters**: A button (filter icon) that toggles the **Filters** pane between expanded and collapsed.
* **Group By**: Choose a property from the drop-down menu by which to group the insights in the list. You can group by:

  + None: Insights are not grouped
  + Insight Name: Insights are grouped by name, e.g., “No Flow From Device”
  + Impact: Insights are grouped by impact, e.g., “Major”
  + Family: Insights are grouped by family, e.g., “Network Health”
  + Key: Insights are grouped by key, e.g., “Device: nms-alert-router-15”
* **Search field**: A field that shows a lozenge for each filter category that is currently set in the **Filters** pane and that also enables you to enter text that will be used to filter the **Insights** list. The list will show only rows in which there’s a match between the entered text and the **Insight**, **Family**, **Key**, or **Value** columns. Click the **X** at the right of the field to clear entered text, and the **X** in a lozenge to clear the corresponding filter.
* **Filters**: A pane with controls that narrow the insights shown in the **Insights** list (see **Insights List Filters**).
* **Insights List**: A list of your organization's insights (see **Insights List**).

### Insights List

The Insights list is a table that provides information about the insights in your organization, filtered by the **Insights List Filters**. Each row in the table represents an individual insight. Click a row to open the **Insight Details Page** for that insight.  
  
 The table includes the following columns:

* **Impact**: The severity value (Critical, Severe, Major, Warning, Minor, or Notice) applied to the insight when it was created by Kentik or your organization.
* **Insight**: The name assigned to the insight by Kentik (see **Insights Reference**).
* **Family**: The family to which the insight belongs (see **Insights Families**).
* **Key**: A unique combination of values for a given set of dimensions (see **About Keys**). An insight is generated when the dimensions in that insight's key definition match specified values (absolute or relative to a baseline).
* **Value**: The dimension value that caused an insight to be generated.
* **Time**: The date-time at which the insight was generated.

### Insights List Filters

The filters in this pane determine which insights are listed in the **Insights List**:

* **Time Range**: A drop-down that shows the time range during which the insights occurred. To change the time range, click the drop-down to open the **Time Range Controls**. To close the controls, click outside of the controls.
* **Impact**: Use the checkboxes to include or exclude insights whose severity is Critical, Severe, Major, Warning, Minor, or Notice.
* **Family**: A field that displays a lozenge for each selected family (see **Insights Families**) and filters the Insights list to insights from those families. To select a family, click in the field and choose the family from the filterable drop-down list. To remove a family, click the X at the right of its lozenge.
* **Insight Name**: A field that displays a lozenge for each selected insight name and filters the Insights list to insights with that name. To select a name, click in the field and choose the name from the filterable drop-down list. To remove a name, click the X at the right of its lozenge.

#### Time Range Controls

The **Time Range** controls filter the listed insights to those that occurred within a specified time range in UTC (default is Last 24 hours):

* **From**: A date and time field that sets the start of the time range.
* **To**: A date and time field that sets the end of the time range.
* **Calendars**: Side-by-side monthly calendars that show the time range (if it spans more than one day) and enable you to change the range by clicking on a date (which changes the dates in the **From** and **To** fields). The start and end days of the time range are indicated in blue, and the intervening days are shaded in light gray. The calendars include the following controls:

  + Forward/Back: Buttons at the top left and right change the pair of displayed months.
  + Month/Year: Drop-down selectors enable you to choose the month and year.
* **Lookback list**: A list of durations displayed to the left of the calendars. Click in the list to set the duration (shown in the From and To fields) for which the time range will look back from the present. Options include Last hour, Last 8 hours, Last 24 hours, Last 7 days, Last 14 days, Last 30 days, and Last 90 days.
* **Apply**: A button that applies the specified time range and closes the time range controls.
* **Cancel**: A button that closes the time range controls without applying any time range changes.

## Insight Details Page

The Details page for an individual insight is reached via a link from one of the following locations:

* An individual card in the **Insights Pane** on the **Network Explorer** page or one of its associated Core Aggregate or Core Details pages (see **Accessing Insights**). Click the **Insights** button to open the popup, then click the **View Details** link on an individual card.
* An individual row of the **Insights List** on the Insights page. Click a row to go to the Details page for that insight.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(474).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A43Z&se=2026-05-12T09%3A44%3A43Z&sr=c&sp=r&sig=oLrlU3jiTK52leyrvpThJbtR4YSdKBCMzfxMkIHUAcA%3D)

*The Details page for an individual insight.*

### Details Subnav

The subnav of the Details page (silver bar across the top) has the following elements:

* **Breadcrumbs**: Indicates where you are in the portal. Click the **Insights** breadcrumb to go back to the **Insights List**.
* **Share**: A button that opens the Share dialog to share the information from the Details page by link or email (see **Sharing via the Share Dialog**).
* **Actions**: A drop-down menu of actions related to the Details page:

  + Export: Use this submenu to export the page as a visual report (see **Portal Export Options**).

### Details Main Display

The main display area of each Details page is divided into a set of panes, detailed below, that are intended to provide actionable details about the individual insight.

> **Note:** The types of information presented on Details pages varies depending on the family of the insight (see **Insights Families**).

#### Title Pane

The top-most pane of the page contains the following information (regardless of the family):

* **Insight name**: The name of this insight.
* **Description**: A summary of the specific situation leading to the generation of this insight.
* **Feedback**: A set of buttons (thumbs up/down icons) with which you can rate the value of the insight (see **Feedback Controls**).

#### Dimensions Pane

Depending on the insight, this pane shows the following elements arrayed across the width of the main page:

* **Instance**: The individual entity that the insight is about. For the insight Outbound Site Traffic, for example, the instance would be an individual site, whereas for the insight No Flow From Device the instance would be an individual device.
* **Dimensions**: Statistical or other informational values (depending on the insight) that illustrate the situation that generated the insight:

  + Statistics: Typically related to traffic volume, may include the average bitrate of traffic over a given time range, the percent change in traffic compared to an earlier time range, the total traffic over a given time range, the baseline of traffic over a previous time range, etc.
  + Other information: Varies by insight. for example, an insight such as Device firmware version check (Security family), will show information about the device whose firmware is being checked, such as the device name, OS, system description, and, if applicable, the deprecated and discontinued dates of the OS.

> **Notes:** The presence of this pane depends on the family of the insight.

#### Data Pane

The data pane shows traffic data related to the condition that caused the insight. Depending on the insight, this may be presented as a time series chart, a table, or both.

* **Table**: A top-X list of items (e.g., cities for the Inbound City Comparison insight) showing information such as current rank, change in rank, percent change in value, current value, and previous value.
* **Chart**: A time series chart illustrating the traffic for a set of items over a time range corresponding to the evaluation frequency of this insight.

### Details Sidebar

The right-side sidebar of the Details page provides additional details about the insight. The presence of the following fields varies depending on the family of the insight:

* **Impact**: A severity level, either Critical, Severe, Major, Warning, Minor, or Notice.
* **Starting Time**: The start of the period evaluated for the insight.
* **Ending Time**: The end of the period evaluated for the insight.
* **Family**: The family to which the insight belongs (see **Insights Families**). The name is a link that takes you to the main Insights page with the **Insights List** filtered to show only insights from the same family.
* **Frequency**: A summary of the frequency with which this insight has recently occurred. The **Show****All****Occurrences** link takes you to the main **Insights** page with the **Insights List** filtered to show all recent occurrences of the insight.
* **Dimensions**: Depending on the insight, additional information about the dimensions that are shown in the **Dimensions Pane**, including how often the same dimension value appeared in your organization's other insights in the last seven days.
* **Take Action**: Additional steps that you can take related to the insight (actions depend on family):

  + View in Data Explorer: A button that takes you to **Data Explorer**, where the controls will be set to show the traffic that generated the insight.
  + Open in Dashboard: Takes you to the corresponding dashboard (shown only when a corresponding dashboard exists).
* **Notifications**: Subscribe or unsubscribe to notifications about this insight:

  + Subscribe: A button that enables you to receive digest emails for the selected insight.
  + Customize: Opens the **Insight Notification Settings** dialog.
* **Explore More Insights**: A list of insight family names with links that take you to the main **Insights** page with the **Insights List** filtered to show only insights from the same family.

#### Insight Notification Settings

To access the **Insight Notifications Settings** dialog, click **Customize** under **Notifications** in the **Details Sidebar**. The dialog includes the following UI elements:

* **Close button**: An **X** in the upper right corner closes the dialog.
* **Subscription status**: Choose whether to receive notifications for this insight:

  + *Not subscribed*: Receive no notifications.
  + *Subscribed*: Receive notifications.
* **Subscription options**: Checkboxes that appear when **Subscribed** is selected. Check all that apply:

  + Include activity for this insight in a daily digest.
  + Include activity for this insight’s family in a daily digest.
  > **Note**: Unchecking both options will set the **Subscription Status** setting to "Not subscribed."
* **Cancel button**: Exit the dialog. Settings will be as they were before the dialog was opened.
* **Save**: Saves your notification settings.

#### Feedback Controls

The **Feedback** controls are used to provide feedback on the insight. The **Comment** field and **Submit** button are shown only when either of the rating buttons are clicked. The controls include the following UI elements:

* **Yes**: A button to give positive feedback about the insight.
* **No**: A button to give negative feedback about the insight.
* **Comment**: An optional field for entering feedback about the insight.
* **Submit**: A button to submit the feedback and hide the **Feedback** controls.

> **Note**: Once the comment field is displayed, it cannot be hidden until feedback is submitted.

## Insights Settings Page

The Insights Settings page is used to configure the display of insights. Access it from the **Configure Insights** button on the Insights page (see **Insights Page UI**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(475).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A43Z&se=2026-05-12T09%3A44%3A43Z&sr=c&sp=r&sig=oLrlU3jiTK52leyrvpThJbtR4YSdKBCMzfxMkIHUAcA%3D)

*Use the Settings page to control the display of insights.*

### Insight Settings UI

The page includes the following UI elements:

* **Family**: Choose an Insights family (e.g., Network Health) from the drop-down menu. The Insights list will show only insights from that family (choose All to show all insights).
* **Filter**: Enter text in the field to filter the **Insights** list to show only rows containing the entered text in one of the following columns: Insight Family, Insight Name.
* **Insight Settings List**: A table of your organization's insights (see **Insights Settings List**).

### Insights Settings List

The **Insights Settings** list shows all your organization's insights and enables you to set properties of individual insights. The list includes the following columns:

* **Insight Family**: The family to which the insight belongs (see **Insights Families**).
* **Insight Name**: The name assigned to the insight when it was created by Kentik or by your organization.
* **Company Enabled**: A checkbox enabling you to control visibility of an insight for all users in your organization.
* **Display in my Insight Feed**: Determines whether the insight will appear in the **Insights Pane** on **Network** **Explorer** pages.

### Manage Insights Settings

To configure the settings for a given insight:

1. Click **Insights** in the Core section of the main navbar menu.
2. On the Insights page, click the **Configure Insights** button at the upper right of the main display area.
3. On the Insights Settings page, look in the **Insights Settings List** for the insight that you'd like to configure.
4. In the insight's row:

   1. Use the **Company Enabled** checkbox to enable/disable the insight for your organization.
   2. Use the **Display in my Insight Feed** checkbox to determine whether the insight will appear in your personal insight feed.
