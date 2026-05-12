# MKP Tenant Portal

Source: https://kb.kentik.com/docs/mkp-tenant-portal

---

This article discusses the MKP tenant portal, which is the My Kentik Portal (MKP) as seen by the users of a given tenant.

> **Note:** For information on the management of your organization's My Kentik Portal in the Kentik portal, see **My Kentik Portal**.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(295).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A22Z&se=2026-05-12T09%3A38%3A22Z&sr=c&sp=r&sig=BWU0m0sXgpGocHeREUNawBlZ9RGTv%2BdAvOxmxDx8BZY%3D)

*An example of the My Kentik Portal as displayed to a tenant user.*

## Tenant Portal Access

The specific elements included in your organization's MKP portal vary depending on the settings for each tenant. To see the MKP portal as it will be presented to a given tenant, you can use the View as Tenant option, or you can log in to the tenant portal as a tenant.

### **View as Tenant**

To view the MKP portal as it would appear to a tenant:![MKP-tenant-portal-access.png](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MKP-tenant-portal-access.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A22Z&se=2026-05-12T09%3A38%3A22Z&sr=c&sp=r&sig=BWU0m0sXgpGocHeREUNawBlZ9RGTv%2BdAvOxmxDx8BZY%3D)

1. Click the **Actions** button at the right of a tenant's row in the **Tenants List** (on the **Tenants** tab of the My Kentik Portal page) and select **View as Tenant** from the drop-down menu.
2. If you have an existing user profile for that tenant, you’ll be logged into the tenant portal with that profile. If you have not created a user profile, but others have, you will be logged in as the first user on the Users List based on how it was sorted when the tenant configuration was last saved. (See **Tenant User Settings** for details on how to create/modify users for a tenant).

To exit the MKP tenant portal and return to the My Kentik Portal page, click on the user menu at the right of the main navbar and choose **Exit Tenant View** from the drop-down menu.

### **Direct Tenant Portal Login**

To log in to the tenant portal directly:

1. Navigate to the URL for your MKP portal at `subdomain`.my.kentik.com. Your sub-domain is defined on the **Domain** tab of the MKP Settings page (see **MKP Domain Settings**).
2. Log in with a user profile that you previously created and activated for that tenant. See **MKP User Activation** for instructions on how to create a tenant user profile.

## Tenant Portal Page

The default landing page of the MKP tenant portal is an overview titled “Explore Your Network” and contains the following UI elements:

* **Library**: A tab that shows tenant users a Library similar (in concept, not layout) to the **Library** in the Kentik portal. The page displays a categorized set of cards that each represent a view (Dashboard or Saved View) that has been assigned to this tenant, either directly or via a package (see **Tenant View Settings**). To access a given view, click on the corresponding card.
* **Alerting pane**: If any alert policies have been assigned to this tenant, either directly (see **Tenant Policy Settings**) or via a package (see **Package Policy Settings**), then the main page of the tenant portal will include an Alerting pane at the upper right, including these indicators:

  + Active: Shows a count of alerts that are in ACTIVE state over the last 24 hours, meaning that conditions defined in the alert policy have been met, notifications have been triggered, but that the alert has not yet been cleared.
  + Cleared: Shows a count of alerts that have been cleared over the last 24 hours.
  + Ack Req: Shows a count of alerts that are in ACK REQ. state over the last 24 hours, meaning the alert was configured to require manual acknowledgement that hasn't yet been given.
* **Mitigations pane**: If any mitigations have been defined for the alerts for a tenant, then the main page of the tenant portal will include a Mitigations pane at the upper right, including these indicators:

  + Active: Shows a count of mitigations that are currently active, starting, clearing, or for which the grace period has yet to end.
  + Failed: Shows a count of mitigations that were attempted but that have either failed to start or to clear.
  + Waiting: Shows a count of mitigations for which start/stop are pending user acknowledgement (or expiration of timer) or are no longer active but require user acknowledgement.

Clicking the Alerting or Mitigations link in these panes takes you to the **Tenant Alerting Page**or **Tenant Mitigations Page**, respectively. There you can view detailed lists of alerts/mitigations and take actions on them.

> **Note**: For instructions on changing the default landing page of the MKP tenant portal to a different view of your choice, see **Tenant View Settings**. 

## Tenant Alerting Page

The link in the **Alerting** pane on the **Tenant Portal Page** takes you to the Alerting page, which provides information about alerts generated by policies assigned to a given tenant, either directly (see **Tenant Policy Settings**) or via a package (see **Package Policy Settings**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(297).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A22Z&se=2026-05-12T09%3A38%3A22Z&sr=c&sp=r&sig=BWU0m0sXgpGocHeREUNawBlZ9RGTv%2BdAvOxmxDx8BZY%3D)

*The Alerting page tells MKP tenant portal users about active or historical traffic anomalies.*

For a full description of the Alerting UI, see **Alerting**, noting the following differences for tenants vs. landlords:

* Tenants do not see the **Manage Policies** button or the **Actions** drop-down menu at the right of the subnav.
* In the **Filters** pane at the left of the page, tenants do not have the option to filter the Alerts List by Sites.

## Tenant Mitigations Page

The link in the Mitigations pane on the **Tenant Portal Page** takes you to the Mitigations page, which provides information about mitigations involving traffic related to the tenant. For a full description of the Mitigations UI see **Mitigations**, noting the following differences for tenants vs. landlords:

* Tenants are not shown the **Manage Mitigations** button or the **Actions** drop-down menu at the right of the subnav.
* Tenants are not shown the **Start a Manual Mitigation** button in the title section below the subnav.
* In the **Filters** pane at the left of the page:

  + Tenants do not have the option to filter the Mitigations List by **Method** or by **Platform**.
  + There is no **Show Tenant Mitigations** switch.

## Tenant Notifications Page

Notifications to tenants are triggered in response to network traffic conditions that meet the criteria specified in one of the policy thresholds defined in **Tenant Policy Settings**. Tenant notifications are available to a tenant when the **Enable Notifications** switch is on in the individual threshold settings for an alert policy that is assigned to that tenant (either directly or as part of a package). While this switch enables the tenant to set up notifications, the notification settings themselves are accessed from within the tenant MKP portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(299).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A22Z&se=2026-05-12T09%3A38%3A22Z&sr=c&sp=r&sig=BWU0m0sXgpGocHeREUNawBlZ9RGTv%2BdAvOxmxDx8BZY%3D)

*Notifications enable you to define channels via which your MKP portal can communicate with users.*

### Tenant Notifications Access

To access the notification settings in the MKP tenant portal:

1. Go to the MKP portal for a given tenant in the **Tenants List** using the **View as Tenant** button (see **Tenant Portal Access)**.
2. Click **Alerting** in the Alerting pane (see **Tenant Portal Page**) to go to the **Tenant Alerting Page**.
3. Click the **Manage Notifications** button in the top right of the page, which takes you to the tenant's Notification Channels page.

### Tenant Notifications UI

A tenant's Notification Channels page is used to manage the notification channels used by that tenant. The page includes the following UI elements:

* **Show/Hide Filters button**: Click to show/hide the filters pane on the left of the page.
* **Search field**: A field that shows lozenges for the currently set filters and also enables you to enter text to narrow the notification channels shown in the Channels List. The list will be filtered to those whose **Name**, **Type**, or **Status** columns contains the entered text.
* **Add Notification Channel**: A button that opens the Add Notification Channel dialog.
* **Filters pane**: A set of filters that enable you to narrow the Channels list (see **Tenant Channel Filters**).
* **Channels List**: A table listing the tenant's collection of notification channels (see **Tenant Channels List**).

#### Tenant Channel Filters

The controls in the **Filters** pane enable you to apply filters to narrow the channels listed in the Channels list. Filters are available for the following categories:

* **Status**: Checkboxes that filter by enabled/disabled status.
* **Type**: Checkboxes that filter by the type of notification channel (e.g., Custom Webhook, Email, JSON, etc.).

Once a filter is applied it is represented in the Search field as a lozenge. To remove the filter, click the **X** at the right of its lozenge or uncheck the box in the Filters pane.

### Tenant Channels List

The Channels list is a table that lists the notification channels set up in this tenant. The list includes the following columns:

* **Name**: The name assigned to the channel in the channel settings dialog (Add Notification Channel or Edit Notification Channel).
* **Type**: The type of notification channel. Supported types include Custom Webhook, Email, JSON, Microsoft Teams, Opsgenie, PagerDuty, ServiceNow, Slack, Splunk, Syslog, VictorOps, and xMatters.
* **Status**: Indicates whether the channel is currently enabled or disabled.
* **Edit** (pencil icon): A button that opens the Edit Notification Channel dialog, enabling you to change the channel settings.
* **Remove** (trash icon): A button that removes the notification channel from the tenant's collection of channels.

### Tenant Channel Settings

The settings of an individual notification channel are made in the Add Notification Channel and Edit Notification Channel dialogs. For full descriptions of the dialogs, see **Notification Dialogs UI**).

> **Note**: Tenants (or landlords in View as Tenant mode) see only the **Settings** tab in the tenant channel settings dialogs, while landlords will also see the **Preview** and **Used By** tabs.

## Tenant User Profile

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(301).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A22Z&se=2026-05-12T09%3A38%3A22Z&sr=c&sp=r&sig=BWU0m0sXgpGocHeREUNawBlZ9RGTv%2BdAvOxmxDx8BZY%3D)

*The User Profile enables users to manage their experience in the tenant portal.*

Like the main Kentik portal, the MKP tenant portal includes a User Profile that covers information and settings related to the individual user. To access the User Profile, click the user icon at the far right of the main navbar, then choose **User Profile** from the drop-down menu. On the Profile page, you'll see three tabs across the top, each of which corresponds to a set of user-specific settings that do not affect other users in this tenant.

> **Note**: Some of the user-specific settings are also available directly from the user drop-down menu, e.g., Authentication, Theme, and Visualization Color Settings.

### Tenant Profile General Settings

The **General** tab includes User Profile settings in the **User Information**, **Theme**, and **Defaults** panes, and a **Save** button to save any changes made on the tab.

#### User Information Pane

This pane displays basic information specific to the individual user (editable where indicated):

* **Role** (read-only field): Administrator or Member (see **About Users**).
* **Full name** (editable field): The user's full name as specified when the user was added or last edited.
* **Email** (editable field): The email address specified when the user was added or last edited.

#### Theme Pane

This pane displays the following setting:

* **Theme** (radio buttons): Sets the visual theme for the tenant portal for the user. The choices are Auto (OS Setting), Light, or Dark.

#### Defaults Pane

This pane displays the following setting:

* **Time zone** (radio buttons): The time zone in which times are expressed, either UTC or Local.

### Tenant Profile Authentication Settings

The **Authentication** tab contains User Profile settings that support user authentication. Details are covered in the links below:

* **Reset Password**: See **Password**
* **Two-factor Authentication**: See **Two-factor Authentication**

### Tenant Profile Visualizations Settings

The **Visualizations** tab contains settings for customizing the appearance of visualizations in the MKP tenant portal for an individual user. For details on these settings, refer to **Chart Colors**.
