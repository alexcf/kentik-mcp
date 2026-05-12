# Notification Channels

Source: https://kb.kentik.com/docs/notifications

---

This article covers the **Notification Channels** page in the Kentik portal.

> **Notes:**
>
> * For an introduction to policy alerting, see **Policy Alerts Overview**.
> * For more on alert policies, see **Alert Policies**.
> * For more on active/historical alerts, see **Alerting**.
> * For more on alert mitigations, see **Mitigations**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(679).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)

*Create and manage notification channels on the Notification Channels page.*

## About Notification Channels

Notification channels, configured at Settings » **Notifications**, are how Kentik triggers notifications about important events on your network, for example if traffic meets the threshold conditions of an alert policy. Each notification channel combines a specific type of notification, e.g., email, PagerDuty, Slack (see **Notification Types**), with a target, such as a set of email recipients or a destination URL.  
  
Your organization’s notification channels can be used in a variety of Kentik portal contexts:

* ![List of categories used in a document, including policies and mitigation methods.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/NC-used-by-filters.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)**Policies**: Notify when traffic meets the conditions specified in a threshold of an alert policy, resulting in the triggering of an alarm (see **Threshold Notifications**).
* **Mitigation Methods**: Notify when an automatic mitigation using a given mitigation method is triggered by conditions set in a threshold of an alert policy (see **General Method Settings**).
* **Insights**: Notify for auto-generated notices from Kentik’s Insights module. These notification channels are:

  + Shown on, but not managed through, the **Notification Channels**page.

    - Filter using the “Insights” and/or “Insight Families” checkboxes
  + Subscribed to directly from an Insight (see **Insight Notifications Settings**)
* **Tenant Policies**: Notify on alert policies for My Kentik Portal tenants, as configured in the tenant settings (see **Tenant Policy Settings**).
* **Synthetics Tests**: Notify when a test result meets the specified test criteria (see **Alerting and Notifications**).

## Notification Channels Page

Notification channels are managed from the Notification Channels page, accessible via Settings » **Notifications** from the Kentik portal’s main menu.

> **Notes:**
>
> * The Notification Channels page settings determine “how” and “to whom” notifications will be sent, but not “when” or “why”. For more information about how notifications are used, see **Alerts and Policies**.
> * Member-level users can only view the Notifications page, while Administrator users can also create and edit notification channels.

### Notification Channels Page UI

The Notification Channels page includes the following UI elements:

* **Show/Hide Filters** (funnel icon): A button to the left of the **Search** bar that toggles display of the **Filters** pane.
* **Search**: Enter text to search your list of existing notification channels.
* **Add Notification Channel** (Administrators only): A button that allows you to create a notification channel.
* **Notification Channels list**: A table listing all notification channels currently in use by your organization (see **Notifications Channels List**).
* **Filters**: A pane for filtering the list of notification channels based on multiple parameters (see **Notifications Filters**).

### Notification Channels List

The Notification Channels list is a filterable table of the notification channels in your organization. Click on the heading of any column (except **Used By**) to sort the list (ascending or descending). The table includes the following columns:

* **Name**: The user-defined name assigned to the notification.
* **Type**: The type of notification. (e.g., Email, Slack, JSON).
* **Used By**: The portal locations using the notification channel.

  + If the channel is currently in use, click the contents of the cell to open the Notification Channel Usage Detail dialog, which provides details on each entity (policy, test, etc.) that uses the channel.
* **Daily Digest**: True (checkmark) or false (blank). Shows if the notification channel is being used for a daily digest.
* **Status**: The notification channel status, either enabled (green) or disabled (red).
* **Edit** (pencil icon): Click to edit the notification channel. Only visible to Admin users.
* **Remove** (trash icon): Click to remove the notification channel from the Kentik system. Only visible to Admin users.

> **Note:** In addition to the notification channels used by alerting and synthetics, this table also shows channels subscribed to by insights, whose notifications are always via email (see **Details Sidebar**).

### The Filters pane of the Notification Channels page, showing various controls for filtering the list.Notifications Filters

The **Filters** pane enables you to narrow the Notification Channels list to only those notification channels that match all of the specified criteria, and includes the following sections:

* **Status**: Check the boxes to include enabled or disabled notification channels.
* **Insight Names**: Click to drop down a list of Insight Names to filter by, e.g. “Device Traffic Increase”.
* **Insight Families**: Click to drop down a list of Insight Families (see **Insights Families**) to filter by, e.g., “Network Health”.
* **Policies**: Click to drop down a list of your organization’s alert policies to filter by.
* **Type**: Check the boxes to filter by notification types (see **Notification Types**), e.g. “Email” or “PagerDuty”.
* **Used By**: Check the boxes to filter by where in Kentik that the notification channel is being used (e.g., “Policies” or “Mitigation Methods”).

> **Notes**:
>
> * Each selected Insight Name, Insight Family, and Policy appears as a lozenge in both the **Search** bar and in the respective **Filters** pane section.
> * To reset your selections, click **Reset To Default** at the top of the Filters pane.

## Manage Notification Channels

The following topics cover the basics of managing notification channels in Kentik.

### Add or Edit Notification Channel

Notification channels are added or edited using one of the following dialogs:

* **Add Notification Channel**: Used to create a new alert notification channel. Access by clicking the **Add Notification Channel** button on the upper right of the Notification Channels page.
* **Edit Notification Channel**: Used to modify an existing notification channel. Access by clicking the **Edit** icon (pencil) on a notification channel’s row.

To configure notification channels in these dialogs:

* Refer to the documentation for detailed information on the specific notification channel type you want to add or edit (see **Notification Types**).
* Follow any preliminary steps external to Kentik that are required to set up the notification channel.
* Set the **Common Notification Settings**.
* Set any additional settings specific to the notification channel type (refer to the individual type settings in **Notification Settings**).
* Follow any additional steps external to Kentik that are required to complete setting up the notification channel.

> **Note:** Notification channels from insights are managed separately (see **Details Sidebar**).

### Enable or Disable Notification

The option to enable or disable a notification channel is displayed during creation, with the default state being **Enabled**. The current state is displayed in the **Status** column of the **Notifications List**.

To enable/disable a notification channel:

1. On the Notification Channels page, find the row for the notification channel.
2. Click the **Edit** icon at the right of the row to open the Edit Notification Channel dialog.
3. Use the **Status** switch in the upper right to enable or disable the notification channel.

> **Note:** When you disable a notification channel, any alerts or insights that use that channel will no longer be sent.

## Notification Dialogs UI

![Settings interface for adding a notification with various options and templates.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/NT-add-notification-dialog(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)The Add Notification Channel and Edit Notification Channel dialogs share the same layout and the following common UI elements:

* **Close**: Click the **X** in the upper right corner to close the dialog. All elements will be restored to their values at the time the dialog was opened.
* **Tab selector**: Select which tab is displayed:

  + **Settings**: Contains the **Name**, **Status**, and **Type** fields (see **Common Notification Settings**) as well as other settings that vary by notification channel type.
  + **Preview** (Edit dialog only): Contains the **Preview** field and a **Test** button (see **Notifications Preview Tab**).
  + **Used By** (Edit dialog only): Contains a list of the alerts and/or insights that use this notification (see **Notifications Used By Tab**).
* **Cancel**: Click **Cancel** to exit the dialog without saving. All elements will be restored to their values at the time the dialog was opened.
* **Add Notification Channel** (Add dialog only): Click this button to save your settings for the new channel and exit the dialog.
* **Save** (Edit dialog only): Click this button to save your changes to the notification channel settings and exit the dialog.

## Notification Types

Kentik supports alert notification channels using the following notification systems/types:

* **Custom Webhook**
* **Email**
* **JSON**
* **Microsoft Teams**
* **OpsGenie**
* **PagerDuty**
* **ServiceNow**
* **Slack**
* **Splunk**
* **Syslog**
* **VictorOps**
* **xMatters**

> **Notes:**
>
> * The links above go to the settings descriptions for each notification type.
> * Insight notifications are currently supported via email channels only (see **Details Sidebar**).

## Notification Channel Settings

The settings for notification types in the Add Notification Channel and Edit Notification Channel dialogs are described here.

### Common Settings

The following settings are common to the **Settings** and **Preview** tabs for all notification channel types.

#### Settings Tab

The following common settings are present on the **Settings** tab:

* **Name**: A user-assigned name for the notification channel.
* **Status**: A toggle switch that determines whether the notification channel is enabled (available for use) or disabled.
* **Type**: The type of the notification channel. The remaining fields of the dialog vary depending on the type:

  + Add dialog: A set of buttons from which you can choose one of the types listed in **Notification Types**.
  + Edit dialog: An inactive button that serves as an indicator of the notification channel's type.

#### Preview Tab

The **Preview** tab includes the following common settings:

* **Preview**: A locked field that displays the HTML markup for the notification channel.
* **Test**: Click this button to receive a test notification via the selected notification channel type.

#### Used By Tab

The **Used By** tab contains a sortable table that shows all of the entities (e.g., alerts, tests, insights) that currently use this notification channel, including:

* **What**: The insights, insight families, mitigation methods, policies, and synthetic tests that use this channel for notifications.
* **ID**: The ID of the entity listed in the **What** column. An ID may be numeric (e.g., for policies and tests) or a string (e.g., for insights).
* **Name**: The name of the entity listed in the **What** column.
* **Link**: A link that takes you away from the Notification Channels page to the page where the entity listed in the **What** column can be managed.

  > **Note:** If the entity is a custom insight, the existing browser tab is unchanged and the link is opened in a new browser tab.
* **Remove** (trash icon): A button that opens a confirmation dialog that allows you to stop notifications via this channel from the entity listed in the **What** column.

### Custom Webhook Notification Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(682).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)Custom webhook notification channels are the most powerful and flexible way of integrating Kentik notifications with third party or custom output channels. They provide a programmatic approach to sending notifications to an API endpoint.  
  
The following fields are used to create Custom Webhook notification channels:

* **URL**: The address to which Kentik should send notifications.
* **Custom Headers**: A set of one or more text fields, each of which is used to specify one custom HTTP header to include in notifications from this channel. The following types of headers are supported:

  + Authorization
  + Any HTTP headers that start with an `x-` prefix
  + Any non `x-` prefix header that is not part of the known standard request headers (see **Common Non-Standard Headers**).
* **Add**: A button that adds another field for a custom header.
* **Custom Template**: A field in which to describe the notification channel's payload using Go-template syntax (see **Using Go Templates**). This allows you to specify the contents of the API request. For more information, see **Using Custom Webhook Templating**.
* **Uglify JSON**: A switch to enable or disable the removal of unnecessary whitespace from a JSON payload. The default is **Off**.

#### Creating Custom Webhook Templates

Kentik’s custom webhook templates use the templating syntax of the Go programming language, but we’ve added convenient functions and variables for referencing the properties and attributes of notification channels. For more information on how to use custom webhooks, see:

* **Kentik Custom Notification Templates**
* **Go Template Documentation**
* **Learn Go Template Syntax**

### Email Notification Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(683).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)In addition to the **Common Notification Settings**, use the following **Settings** tab elements for email notification channels:

* **Email Addresses**: A set of one or more text fields, each of which is used for the address of one email recipient.
* **Add**: A button that adds an email address field.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(684).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)

*An alert notification email from Kentik.*

> **Note:** A notification email generated by an alert policy contains links that enable you to view the traffic that triggered the alarm, either in the **Alert Details Page** or in the alert policy's designated "policy dashboard" (see **General Policy Settings**).

### JSON Notification Settings

> **Note:** While Kentik provides legacy support for JSON notifications, we recommend using custom webhook notifications instead.

In addition to the **Common Notification Settings**, use the following **Settings** tab field for JSON notification channels:

* **URL**: The URL to which Kentik should post JSON notifications for this channel.

#### Using JSON Notifications

JSON notification channels enable you to integrate Kentik with third-party monitoring systems so that events detected by Kentik (e.g., an anomaly detected by our Alerting system) can trigger external actions, which may include network configurations, DDoS attack mitigations, or other remedies. The JSON payload is posted to the specified Webhook URL, where it can be parsed and processed for any desired purpose. For an example of the JSON payload that will be posted, see **Sample Alert JSON**.  
  
JSON may be a good option for your notification channels if you prefer not to define the syntax (unlike custom webhooks). It provides a hard-coded format of the notification payload without customization, enabling users to rely on a stable and defined payload.

#### Secure Receipt of JSON

You can secure your receipt of JSON alert notifications from Kentik using any of the following complementary methods:

* **Filter by IP**: Filter inbound POST requests (using `iptables` on the web-service server or any firewall in the path, or in the code of the web service itself) to only those from IPs in the netblock 209.50.158.0/23.
* **Filter by HTTP header**: Filter inbound POST requests to only those with the following header: "User-Agent:KentikAlerting"
* **Use a query argument**: If you use HTTPS for the Webhook, include a query argument known only to your responding web-service.

#### Testing JSON Notifications

If problems arise when testing JSON notification channels, ensure that your web server is available to accept HTTP requests and also that it is accessible via a public URL. Kentik suggests trying the methods below to debug any issues.

* **RequestBin**: Use the free web-service `RequestBin` (**https://requestbin.com**) to collect HTTP requests in bins. You can examine these to see the contents of the request.
* **ngrok**: If you have a development web server on your local host that accepts HTTP-POST, use `ngrok` (**https://ngrok.com**) to get a unique public URL that is directly connected to your development machine.

#### Sample Insight JSON

The following example illustrates the JSON that would typically be sent as a notification for an insight (see **Insight Notification Settings**).

```
{
  "CompanyID":####,,
  "CurrentState": "n/a",
  "Description": "Insight for Through Site Comparison",
  "Dimensions": {},
  "EndTime": "2022-11-02 13:50:00 UTC",
  "InsightDataSourceType": "ksql",
  "InsightID": "k654321097",
  "InsightName": "operate.site.bpsThroughWeekOverWeek",
  "InsightPlainDescription": "Foobar Amsterdam AMS13 had 22% more through traffic (+139 Kbits/s) this week compared to last week.",
  "IsActive": true,
  "Links": {
    "InsightDetailsURL": "https://portal.kentik.com/v4/operate/insights/k611587433",
    "InsightsMainURL": "https://portal.kentik.com/v4/operate/insights",
    "InsightsSeverityURL": "https://portal.kentik.com/v4/operate/insights?severities=notice"
  },
  "Metrics": {},
  "PreviousState": "n/a",
  "StartTime": "2022-11-02 11:50:00 UTC",
  "Type": "EVENT_VIEW_TYPE_INSIGHT",
  "issue": [],
  "statistic": {}
}
```

#### Sample Alert JSON

The following example illustrates the JSON that would typically be sent in response to a change in the state of an alert (see **Threshold Notifications**).

```
{
  "ActivateSeverity": "critical",
  "AlarmBaselineDescription": "A static threshold was used (no baselining).",
  "AlarmEnd": "ongoing",
  "AlarmID": "0192d7bf-71ff-7ab8-a494-697f532902d7",
  "AlarmPolicyApplication": "core",
  "AlarmPolicyID": "57187",
  "AlarmPolicyName": "Clients_GLOBAL_TMS",
  "AlarmSeverity": "critical",
  "AlarmStart": "2024-10-29 10:08:20 UTC",
  "AlarmState": "alarm",
  "AlarmStateOld": "new",
  "AlarmThresholdID": "214982",
  "AlertBaselineSource": "0",
  "AlertDimensions": [
    "IP_dst",
    "c_ddos_ghost",
    "i_device_site_name"
  ],
  "AlertKey": [
    {
      "DimensionName": "IP_dst",
      "DimensionValue": "103.122.191.63"
    },
    {
      "DimensionName": "c_ddos_ghost",
      "DimensionValue": "FL-HK2-GRE-ipv4"
    },
    {
      "DimensionName": "i_device_site_name",
      "DimensionValue": "SV5 San Jose"
    }
  ],
  "AlertPolicyName": "Clients_GLOBAL_TMS",
  "AlertValue": {
    "Unit": "bits",
    "Value": 1761688121.2877295
  },
  "AlertValueSecond": {
    "Unit": "packets",
    "Value": 164264.93744145698
  },
  "AlertValueThird": {
    "Unit": "unique_src_ip",
    "Value": 89
  },
  "Baseline": 0,
  "CompanyID": 24677,
  "CurrentState": "alarm",
  "Description": "Alarm for Clients_GLOBAL_TMS Active",
  "Dimensions": {
    "IP_dst": "103.122.191.63",
    "c_ddos_ghost": "FL-HK2-GRE-ipv4",
    "i_device_site_name": "SV5 San Jose"
  },
  "EndTime": "ongoing",
  "EventType": "alarm",
  "IsActive": true,
  "Links": {
    "DashboardAlarmURL": "https://portal.kentik.com/v4/alerting/dashboard/767/0192d7bf-71ff-7ab8-a494-697f532902d7",
    "DetailsAlarmURL": "https://portal.kentik.com/v4/alerting/0192d7bf-71ff-7ab8-a494-697f532902d7"
  },
  "Metrics": {
    "bits": 1761688121.2877295,
    "packets": 164264.93744145698,
    "unique_src_ip": 89
  },
  "PolicyID": "57187",
  "PreviousState": "new",
  "RuleID": "01907943-df74-7d64-8537-57ff8a11894f",
  "StartTime": "2024-10-29 10:08:20 UTC",
  "ThresholdID": "214982",
  "Type": "alarm",
  "issue": [],
  "statistic": {}
}
```

#### Alert JSON Description

The following table explains the elements of the notification JSON:

| JSON element | Description | Type and values |
| --- | --- | --- |
| EventType | The type of message, either alarm or mitigation notification. | String: ALARM\_STATE\_CHANGE, MITIGATION\_STATE\_CHANGE |
| CompanyID | ID if company. | Number |
| MitigationID | ID of mitigation, if any. If no mitigation, then 0. | Number |
| AlarmID | ID of alarm. | Number |
| AlarmState | Current state of alarm. | String. See **Alert State**. |
| PolicyID | ID of the alert policy generating this notification. | Number |
| ThresholdID | ID of the specific threshold (within a policy) that was triggered to generate the alarm. | Number |
| ActivateSeverity | Severity level of the threshold generating the alarm (see Severity selector in **General Threshold Settings**). | String: Minor, Minor2, Major, Major2, Critical |
| AlarmStart | A date-time string giving the time that the alarm started. | String |
| AlarmEnd | A date-time string giving the time that the alarm ended. If not yet ended, time will be given as 0001-01-01T00:00:00Z. | String |
| LastActivate | A date-time string giving the last time that this alarm was active. | String |
| AlertPolicyName | The name of the alert policy generating this notification. | String |
| AlarmStateOld | The prior state of this alarm. | See **Alert State**. |
| AlertKey | An array of the dimensions that make up the key used to evaluate traffic for this alert policy (see **About Keys**). | Array |
| AlertKey » DimensionName | The system name of an individual dimension that is part of this policy's key. | String. Dimensions, which are based on fields in the KDE main table, are described in **Dimensions Reference**. |
| AlertKey » DimensionValue | The value of an individual dimension that is part of this policy's key. | String or number depending on dimension. |
| AlertValue | An array of current traffic metrics (primary and secondary) by which ingested flow data will be evaluated to determine top-X. | Array |
| AlertValue » Unit | The UI name of an individual current-traffic metric. | String. See **General Metrics** and **Host Traffic Metrics**. |
| AlertValue » Value | The value of an individual current-traffic metric. | Number |
| AlertBaseline | An array of baseline traffic metrics (primary and secondary) by which ingested flow data will be evaluated to determine top-X. | Array |
| AlertBaseline » Unit | The UI name of an individual baseline metric. | String. See **General Metrics** and **Host Traffic Metrics**. |
| AlertBaseline » Value | The value of an individual baseline metric. | Number |
| AlertBaselineSource | The source of the baseline information in the AlertBaseline array. | String. See **Baseline Source for Notifications**. |

#### Baseline Source for Notifications

The following table shows the source of the information returned in the `AlertBaselineSource` element of the notification JSON:

| Baseline Source Name | Description | Baseline Mode |
| --- | --- | --- |
| NO\_USE\_BASELINE | Baseline not used on this threshold at all. | Current to History |
| SKIPED\_BASELINE\_CALCULATION | Baseline fallback set to "Do Not Alarm." | Current to History |
| TRIGGER\_USED\_NO\_BASELINE | Baseline fallback set to "Alarm." | Current to History |
| CALCULATED\_USED\_FOR\_BASELINE | Baseline fallback used "Auto Calc" value | Current to History |
| DEFAULT\_USED\_FOR\_BASELINE | Baseline fallback used Static value | Current to History |
| LOWEST\_USED\_FOR\_BASELINE | Baseline fallback used the "Lowest Top-X" setting | Current to History |
| NOT\_FOUND\_EXISTS\_NO\_BASELINE | Baseline not found and no fallback option set. | Current to History |
| ACT\_CURRENT\_MISSING\_TRIGGER | Key was in the Historical Top-X but not in the Current Top-X. | History to Current |
| ACT\_CURRENT\_USED\_FOUND | Key was found in both Historical and Current Top-X, threshold exceeded. | History to Current |
| ACT\_CURRENT\_MISSING\_DEFAULT | Key was in the Historical Top-X but not in the Current Top-X. | History to Current |
| ACT\_CURRENT\_MISSING\_LOWEST | Key was in the Historical Top-X but not in the Current Top-X, lowest Top-X fallback | History to Current |

### Microsoft Teams Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(685).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)In addition to the **Common Notification Settings**, use the following **Settings** tab field for Microsoft Teams notification channels:

* **URL**: The URL to which Kentik should post Microsoft Teams notifications for this channel.

To use a Microsoft Teams notification channel in Kentik:

1. Set up an incoming webhook in Microsoft Teams as described in the Microsoft documentation topic **Create an Incoming Webhook**.
2. Note the unique webhook URL in step 5.
3. In Kentik, enter the webhook URL into the URL field described above.

### OpsGenie Notification Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(686).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)In addition to the **Common Notification Settings**, use the following **Settings** tab field for OpsGenie notification channels:

* **Token**: A unique API token used by the OpsGenie Web API.

To use an OpsGenie notification channel in Kentik:

1. Set up a corresponding “integration” in OpsGenie as described in the OpsGenie documentation topic **Create an API Integration**.
2. Once you have an OpsGenie account, you can add integrations for an existing team or create a new team of individual OpsGenie users.
3. Once you have a team, add an integration using the steps provided in OpsGenie’s **Using API Integration** topic.
4. In Kentik, copy the API key provided in OpGenie's step 5 and paste it into the **Token** field above.

When OpsGenie notification channels in Kentik are assigned to alert policies (see **Alerts and Policies**), a notification from a triggered threshold in a policy will appear in the alert list in OpsGenie. Click on an item in the list to see an Alert Detail that contains effectively the same information that's included in a JSON alert notification (see **Sample Alert JSON**).

### PagerDuty Notification Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(687).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)The following **Settings** tab field is used only for PagerDuty notification channels:

* **Token**: A unique service identifier used by the PagerDuty Events API to trigger, acknowledge, and resolve incidents.

To establish a PagerDuty notification channel in Kentik:

1. Set up a corresponding service in PagerDuty (see PagerDuty's **Configuring Services and Integrations** support page).
2. In Kentik, put the integration key created by this process into the **Token** field above.

> **Note:**The data included in a PagerDuty alert notification from Kentik is similar to a JSON alert notification (see **Sample Alert JSON**).

### ServiceNow Notification Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(688).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)In addition to the **Common Notification Settings**, use following **Settings** tab fields for ServiceNow notification channels:

* **Instance**: The URL of the ServiceNow Instance.
* **Username**: The username associated with the ServiceNow Instance.
* **Token:** The password associated with the ServiceNow Instance.

To establish a ServiceNow notification channel in Kentik:

1. Log into your ServiceNow Developer Site.
2. Create a ServiceNow “instance” as described in the ServiceNow training topic **Personal Developer Instances**.
3. In Kentik, put the instance name, username, and token generated by ServiceNow into the corresponding fields of the **Settings** tab.

When a ServiceNow notification channel created in Kentik is assigned to an alert policy, notifications triggered by that policy's thresholds will be listed on the Incidents page in ServiceNow. Click an item in the list to view the details of that incident. The data included in a ServiceNow alert notification from Kentik is similar to a JSON alert notification (see **Sample Alert JSON**).

### Slack Notification Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(689).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)In addition to the **Common Notification Settings**, use the following **Settings** tab field for Slack notification channels:

* **URL**: The URL to which Kentik should post Slack notifications.

To establish a Slack Notification channel, you need to set up a webhook in Slack (see **Sending messages using Incoming Webhooks**). Put the webhook URL generated in Slack into the **URL** field described above.

#### Slack Notification Troubleshooting

If you created a Slack notification channel from within the Kentik portal but are not receiving notifications, troubleshoot as follows:

1. Verify your allowed applications in Slack.
2. Verify your permissions in Slack.
3. In Kentik, in the **Notification Channels List** (Settings » **Notifications**), find the channel that you created to receive Slack notifications from Kentik. Use the **Remove** button (trash icon) at the right of the row to remove the channel.
4. If there are existing Slack notifications within Kentik that function as expected, compare the settings of those Slack channels for differences (for example, public vs. private settings).

Problems with Slack notification can be caused by changing the Public/Private setting of a Slack channel in your organization’s Slack system after that channel has been assigned to a Kentik notification channel (e.g., integrating a channel that’s Public, then switching that channel to Private). If you need to change the setting, first remove the notification channel from Kentik, then modify the Slack channel, and then re-create the notification channel in Kentik.

> **Note:** Removing an existing Slack-based notification channel will affect all entities (policies, tests, etc.) in your organization that use that channel. If you remove and recreate the channel, you will then need to add it back to every policy threshold to which the channel was assigned.

### Splunk Notification Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(690).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)In addition to the elements detailed in **Common Notification Settings**, the following **Settings** tab fields are present when the notification channel type is Splunk:

* **URL**: The URL to which Kentik should post Splunk notifications for this channel.
* **Token**: The HTTP Event Collector Token provided by Splunk.

To establish a Splunk notification channel in Kentik:

1. In Splunk, enable and configure the HTTP Event Collector (HEC) according to your specific type of Splunk software (see the Splunk documentation topic **Set up and use HTTP Event Collector in Splunk Web**).
2. Create an event collector token and copy its value. In Kentik, put this value into the **Token** field described above.
3. Determine the specific URL for your type of Splunk software as described in **Send data to HTTP Event Collector**. In Kentik, put this URL into the **URL** field described above.

### Syslog Notification Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(691).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)In addition to the elements detailed in **Common Notification Settings**, the following **Settings** tab fields are present when the notification channel type is System Log (syslog):

* **Host**: IP address to which notification messages will be posted.
* **Port**: Port on which syslog will listen for notifications.
* **Network Protocol**: Protocol to send data, either UDP or TCP
* **Syslog Hostname**: Name for the syslog to which messages will be posted.
* **Severity**: The severity level of the notifications sent via this channel. Options include emergency, alert, critical, etc. as defined in **RFC 5424**.
* **Facility**: The facility code of the notifications sent via this channel. Options include kernel, user, system, etc. as defined in **RFC 5424**.

> **Note:** The data included in a Syslog alert notification from Kentik is similar to a JSON alert notification (see **Sample Alert JSON**).

### VictorOps Notification Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(692).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)In addition to the elements detailed in **Common Notification Settings**, the following **Settings** tab field is present when the notification channel type is VictorOps (now known as Splunk On-Call):

* **URL**: The URL to which Kentik should post VictorOps notifications.

To establish a VictorOps notification channel in Kentik:

1. In VictorOps, enable a REST Endpoint Integration and set up a REST Endpoint Integration Routing Key, as described in the documentation topic **REST Endpoint Integration Guide - Splunk On-Call**.
2. Use the generated routing key to create a VictorOps webhook URL.
3. In Kentik, enter the webhook URL into the **URL** field described above.

### xMatters Notification Settings

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(693).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A08Z&se=2026-05-12T09%3A47%3A08Z&sr=c&sp=r&sig=wjgBE0k%2BFTH8CNII1JHZF06A6s4Ttm%2FXT8hW0QbpduI%3D)In addition to the elements detailed in **Common Notification Settings**, the following **Settings** tab field is present when the notification channel type is xMatters:

* **URL**: The URL to which Kentik should post xMatters notifications for this channel.

To establish an xMatters notification channel in Kentik:

1. Set up an incoming webhook in xMatters, as described in the xMatters documentation topic **Webhooks**.
2. In Kentik, enter the webhook URL into the **URL** field described above.
