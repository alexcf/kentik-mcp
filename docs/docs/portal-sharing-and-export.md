# Portal Sharing and Export

Source: https://kb.kentik.com/docs/portal-sharing-and-export

---

This article discusses data export and sharing from the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(56).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A48Z&se=2026-05-12T09%3A37%3A48Z&sr=c&sp=r&sig=5QPnPmFApJqrmAyrMpKDUnV1qn7670pps%2Fl%2BM8lU79g%3D)

*Sharing and export are accessed via buttons on the subnav.*

> **Note:** For general information on the Kentik portal, see **Portal Overview**.

## Portal View Sharing

The portal’s options for sharing views with others are covered in the following topics.

### About View Sharing

Kentik supports several different ways to share with others what you’re currently seeing in the portal. Some of these methods require the recipient to be a registered Kentik user who’s within your organization, while others enable you to share with persons outside of your organization who may not be Kentik users.

> **Note:** While the **Share** menu is implemented on most portal pages, not all sharing options described below are available on those pages (see **Sharable Views**).

#### Sharing via the Share Dialog

The **Share** dialog, which you access via the **Share** menu in the subnav (see **Portal Page Layout**), includes tabs for the following methods of sharing the current view:

* **Link**: Sharing options from the **Link** tab of the dialog (see **Sharing by Link**):

  + **Internal share**: Copy a URL of the view that you can pass along to other Kentik users within your organization (see **Internal Share Link**).
  + **Public share**: Save the view to a publicly accessible web page that can be viewed by people who aren’t Kentik users, inside or outside of your organization (see **Public Share Link**).
* **Email**: Send the current view as a report that’s attached as a PDF or CSV file (see **Sharing by Email**).
* **Subscription**: Send the view on a scheduled, recurring basis as a PDF or CSV file attached to an email (see **Sharing by Subscription**).

  > **Note:** You can also set up a subscription on the Subscriptions page or, in certain portal modules, by choosing **Subscribe** from the subnav's Actions menu.

In addition to the above tabs, the **Share** dialog also includes buttons, an **X** at top right and **Cancel** at bottom, that close the dialog without sharing and restore all fields to their values when the dialog was opened.

#### Sharing via Export

In addition to Kentik’s built-in features for sharing via link, email, or subscription, you can also download the information from some portal modules and then share it yourself. Use the Actions menu in the subnav (see **Portal Page Layout**) to download a file representing the view (see **Portal Export Options**), after which you can share that file with others using any communication method such as email or Slack.

### Sharable Views

The views that are currently sharable by one or more sharing methods are listed in the table below.

| Module | Internal Share | Public Share | Email | Subscription |
| --- | --- | --- | --- | --- |
| **Alerting** | No | No | No | Yes |
| **Capacity Planning Summary** | Yes | No | Yes | Yes |
| **Specific Capacity Plan(s)** | Yes | No | Yes | Yes |
| **Connectivity Costs - Summary** | Yes | No | Yes | Yes |
| **Connectivity Costs - Providers** | Yes | No | Yes | Yes |
| **Connectivity Costs - Connectivity Types** | Yes | No | Yes | Yes |
| **Connectivity Costs - Sites** | Yes | No | Yes | Yes |
| **Data Explorer** | Yes | Yes | Yes | No |
| **Library - Dashboards** | Yes | Yes | Yes | Yes |
| **Library - Saved View** | Yes | Yes | Yes | Yes |
| **Network Explorer** | Yes | No | Yes | No |
| **RPKI Analysis** | Yes | Yes | Yes | Yes |
| **Test Control Center - Main page, Test Details pages** | Yes | Yes (except BGP Monitor) | Yes | No |
| **Test Control Center - Subtest details** | Yes | No | Yes | No |
| **Synthetic Agents** | Yes | No | Yes | No |
| **State of the Internet** | Yes | No | Yes | No |
| **Botnet & Threat Analysis** | Yes | Yes | Yes | Yes |
| **CDN Analytics** | Yes | No | Yes | No |
| **OTT Service Tracking** | Yes | No | Yes | No |
| **Cloud Traffic Overviews (AWS, GCP, and Azure)** | Yes | Yes | Yes | Yes |
| **Settings -  (Networking) Devices** | Yes | No | Yes | Yes |
| **Settings - Users** | Yes | No | Yes | Yes |
| **Settings - Sites** | Yes | No | Yes | Yes |

> **Note**: For any module where you can share a subscription, you can also create a subscription from the **Subscriptions** page. For RPKI Analysis and the Traffic Trends & Overview modules, you must select **Dashboard** from the **Share** drop-down and then choose the desired module from the list in the **Selected** **Dashboard** drop-down.

### Sharing by Link

The **Link** tab of the **Share** dialog enables two types of link sharing.

![Sharing options for the 'Top 10 Interfaces' document with description and link provided.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PS-share-dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A48Z&se=2026-05-12T09%3A37%3A48Z&sr=c&sp=r&sig=5QPnPmFApJqrmAyrMpKDUnV1qn7670pps%2Fl%2BM8lU79g%3D)

#### Internal Share Link

An internal share is a URL that may be used by other Kentik-registered users in your own organization. To get the link:

1. In the subnav, click the **Share** button to open the Share dialog.
2. In the **Internal Share** pane at the top of the **Link** tab, click the **Copy Link** button. The URL of the current view will be copied to the clipboard.
3. Paste the link into an email, Slack, or some other form of communication to share it with a colleague within your organization.

#### Public Share Link

A public share is a publicly accessible web page that can be viewed by people inside or outside of your organization, who aren’t Kentik users (see **About Public Shares**). In the portal modules listed in **Sharable Views**, public shares are configured in the **Public Share** pane of the **Link** tab of the Share dialog.  
  
 The form in the **Public Share** pane includes the following UI elements:

* **Manage Public Shares**: A link that takes you to the portal’s settings page for **Public Shares**.
* **Title**: A field in which you enter a name for the public share. This may already be populated with the name of the page you’re sharing, but change it as needed.
* **Description**: A field in which you describe to the people with whom you'll be sharing the link why you are sharing this view with them and what they should look for, including relevant timestamps, incident details, etc.
* **Share** **Link**: A field showing the link for the page on which visitors will be able to see this public share. The last segment of the link will be pre-populated based on what you enter into the **Title** field, but you can modify this segment as desired (the segment will be validated for HTTP compliance).
* **Recipients**: A field in which you enter a comma-separated list of addresses (within or outside of your organization) to which you'd like Kentik to send an email with a link to this public share.
* **Message**: A field in which you enter the text of the email in which the link will be sent.
* **Share**: A button that sends an email message with the public share link to the specified recipients.

> **Notes:**
>
> * The above form also exists in the Link dialog that opens from the Edit button (pencil icon) in the **Public** **Shares** list (see **Link Dialog**). The Link dialog version doesn’t include the **Recipient** or **Message** fields.
> * Before creating a public share, be sure that the information in the shared visualization is appropriate for viewing by the public with no authentication.

#### Create a Public Share

To create and send a public share:

1. In the subnav, click the **Share** button to open the Share dialog.
2. In the **Public Share** pane at the bottom of the **Link** tab, specify the settings covered in **Public Share Link**.
3. Click **Share** to send an email message containing the link to people who aren’t Kentik users.

### Sharing by Email

Sharing by email means creating a report that Kentik will send as an attachment to a specified list of recipients. Email sharing is initiated from the **Email** tab of the Share dialog.

![Email sharing options for a capacity planning summary report with selected views.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/PS-share-dialog-email.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A48Z&se=2026-05-12T09%3A37%3A48Z&sr=c&sp=r&sig=5QPnPmFApJqrmAyrMpKDUnV1qn7670pps%2Fl%2BM8lU79g%3D)

#### Email Tab UI

The **Email** tab of the **Share** dialog includes the following UI elements:

* **Subject**: A field for the subject line of the email that will include the shared report.
* **Share**: A drop-down menu that determines the kind of view that you’d like to share (see **Sharable Views**).

  + The page on which you clicked the **Share** button limits the report types you can share. For example, if you clicked **Share** on a Dashboard page, you can only send a Dashboard report.
* **Selected view**: A drop-down menu listing available views of the kind specified with the Share setting. Choose the view from which you’d like to create a report to send via email.

  + Present only when there is more than one available view of the kind selected with the **Share** setting.
* **Report File Type**: A drop-down menu listing file types available for the report (depends on the kind of view specified with the Share setting).
* **Recipients**: A field in which to enter a comma-separated list of addresses (internal or external) to which to email the report.
* **Message**: A field in which to enter the body of the email.
* **Share**: A button that sends an email message with the attached report to the specified recipients.

#### Share by Email

To share by email:

1. In the subnav, click the **Share** button to open the Share dialog.
2. In the **Email** tab, specify the settings covered in **Email Tab UI**.
3. Click the **Share** button to send the report as an attachment to the specified recipients.

### Sharing by Subscription

Sharing by subscription is very similar to sharing by email, but instead of being sent only once, the view is shared on a recurring basis. A report is created from the view at a specified interval and sent to the same recipient(s) each time. Subscription sharing can be initiated:

* From the **Subscription** tab of the Share dialog
* By selecting **Subscribe** from the **Actions** menu
* By selecting **Add** **Subscription** on the **Subscriptions** page

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(107).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A48Z&se=2026-05-12T09%3A37%3A48Z&sr=c&sp=r&sig=5QPnPmFApJqrmAyrMpKDUnV1qn7670pps%2Fl%2BM8lU79g%3D)

#### Subscription Tab UI

The settings for subscriptions are specified in a form that appears — with variations as noted below — in several different contexts in the portal:

* The **Subscription** tab of the Share dialog, accessed via the subnav in numerous portal modules.
* The Subscription dialog that opens from the **Add** **Subscription** button on the **Subscriptions** page.
* The Subscribe dialog that opens from the **Page-wide Actions Menu** on the Alerting page.

The form includes the following settings and controls:

* **Manage Subscriptions** (not present in the Subscription dialog): A link that takes you to the portal’s settings page for **Subscriptions**.
* **Subscription**: A drop-down in which you can do one of the following:

  + Click **Create** **New** to show a field into which you can enter the name of the new subscription. Click **Cancel** (beside the field) to select a subscription rather than create one (this will also clear any fields on the dialog you’ve filled out).
  + Choose an existing subscription from a drop-down list, which enables you to add a recipient (yourself or someone else) to the distribution list or to modify the other properties of a report.
  > **Notes:**
  >
  > + Modifying the content of a subscription will change it not only for you but for all recipients.
  > + In the Subscription dialog, and also the Subscribe dialog for alerts and mitigations, **Subscription** is a text field rather than a drop-down as described above.
* **Share** (not present for alerts and mitigations): A drop-down menu from which you select the kind of view that you’d like to share (see **Sharable Views**).

  > **Note**:
  >
  > The available kinds vary depending on the page from which you opened the Share dialog.
* **Selected View** (not present for alerts and mitigations): A drop-down menu listing available views of the kind specified with the **Share** setting. Choose the view from which you’d like to create a report to send via email.

  > **Notes**:
  >
  > + This setting is present only when there is more than one available view of the kind selected with the **Share** setting.
  > + Certain kinds of views allow you to select more than one view in the **Selected** **View** field. Once selected, these will display as lozenges inside the field.
* **Report File Name**: A text field in which to enter the file name for the report. To include the date the report is sent, use YYYYMMDD at the beginning of the file name.
* **Report File Type**: A drop-down listing file types available for the report, which vary depending on the kind of view specified with the **Share** setting).
* **Recipients**: Three fields (To, CC, and BCC) in which you can add a comma-separated list of addresses (internal or external) to which to email the report. At least one of these fields must be populated with a minimum of one email address to share the report.
* **Schedule**: A drop-down to select the interval at which the report will be sent: Daily, Weekly, Monthly, or the Last day of the month. The interval selected determines whether the Day drop-down appears (the **Time** drop-down is always visible):

  + **Day** (on): A drop-down that appears only when Schedule is set to Weekly (select a day of the week), or Monthly (select a date from 1 to 28).
  + **Time** (at): A drop-down menu to select a time at which to send the report. It displays 30 minutes increments in 24H (in UTC).
* **Lookback** (only for a dashboard or saved view): Radio buttons to select the time range covered by the report:

  + **Days**: If selected, an accompanying text field will enable you to enter a number (up to 90). Each report will cover that number of days, counting back from the day the report is sent. The default is 7.
  + **This Month** (month to date): Cover the current calendar month to date.
  + **Last Month**: Cover the entire previous calendar month.
* **Alerting Filters** (only for an alert or mitigation): Filters that enable you to filter the contents of reports using **Alerts List Filters** (from the Alerting page) or **Mitigations List Filters** (from the Mitigations page). The filter settings in the Subscription dialog or the Subscribe dialog will default to those of the **Filters** pane on the underlying page when the dialog is opened.

#### Share by Subscription

To share by subscription:

1. In the subnav, either:

   * Click the **Share** button to open the Share dialog and select the **Subscriptions** tab, or
   * Click the **Actions** button and select **Subscribe.**
2. Specify the settings covered in **Subscription Tab UI**.
3. Click the **Share** or **Subscribe** button to send the report as an attachment to the specified recipients on a recurring basis at the specified interval.

## Portal Export Options

Many views in the various modules of the Kentik portal can be downloaded for viewing outside the portal or for sharing with others. Downloads are implemented as export options from the **Actions** menu in the subnav (see **Portal Page Layout**). Export options vary depending on the module.

### Portal Export Types

The following export types are supported (depending on the module):

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(114).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A48Z&se=2026-05-12T09%3A37%3A48Z&sr=c&sp=r&sig=5QPnPmFApJqrmAyrMpKDUnV1qn7670pps%2Fl%2BM8lU79g%3D)**Visual Reports**:

  + **Chart + Data Table**: A single PDF containing both the visualization and the results table (which functions as the chart’s legend).
  + **Chart Image**: The visualization only, exported as a PNG.
  + **Chart Vector**: The visualization only, exported as an SVG.
* **Data**:

  + **Chart Data**: Export, as CSV, the data represented in the chart.
  + **Data Table**: Export, as CSV, the data in the table.

### Exportable Views

The table below shows the portal modules from which you can export in at least one of the export types listed above.

| Module | Export? |
| --- | --- |
| **Alerting** | Y |
| **AI Advisor** | N |
| **Kentik Map** | N |
| **Library (Dashboards & Saved Views)** | Y |
| **My Kentik Portal** | N |
| **Observation Deck** | N |
| **Network Explorer** | Y |
| **Data Explorer** | Y |
| **Capacity Planning** | Y |
| **Insights** | N |
| **Raw Flow Explorer** | Y |
| **Synthetics Dashboard** | N |
| **Test Control Center** | Y |
| **Agent Management** | Y |
| **BGP Route Viewer** | N |
| **State of the Internet** | Y |
| **NMS Dashboard** | N |
| **Metrics Explorer** | Y |
| **Query Assistant** | N |
| **Devices** | Y |
| **NMS Interfaces** | Y |
| **Traffic Engineering** | N |
| **Connectivity Costs** | Y |
| **Traffic Costs** | Y |
| **Discover Peers** | N |
| **DDoS Defense** | N |
| **Mitigations** | Y |
| **Botnet & Threat Analysis** | Y |
| **RPKI Analysis** | Y |
| **CDN Analytics** | Y |
| **OTT Service Tracking** | Y |
| **Kentik Market Intelligence** | Y |
| **Cloud Dashboard** | N |
| **Kentik Kube\*** | N |
| **AWS Traffic Overview** | Y |
| **GCP Traffic Overview** | Y |
| **Azure Traffic Overview** | Y |
| **Performance Monitor** | N |
| **Cloud Pathfinder** | N |

\*If you don’t see Kentik Kube in your portal, it might not be enabled for your company’s account. Please contact your **Kentik Account Team**.

### Recent Exports Page

You can review your recently exported views by choosing **Recent Exports** from the portal navbar's **User Menu**, which takes you to the Recent Exports page.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(115).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A16%3A48Z&se=2026-05-12T09%3A37%3A48Z&sr=c&sp=r&sig=5QPnPmFApJqrmAyrMpKDUnV1qn7670pps%2Fl%2BM8lU79g%3D)
