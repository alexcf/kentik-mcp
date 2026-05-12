# Mitigations Page

Source: https://kb.kentik.com/docs/mitigations-page

---

This article discusses the **Mitigations** page in the Kentik portal.

> **Note:** For more on mitigations in Kentik, browse the other articles in the **Mitigations** category.

![Overview of mitigation statuses, including active, failed, and waiting entries in a table.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MT-mitigations-page-main.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A22Z&se=2026-05-12T09%3A42%3A22Z&sr=c&sp=r&sig=Sa11HpJ%2BwWAODiOhviTa3a%2FpnGvCBGesfMPH70EALJE%3D)

*The Mitigations page enables management of current and past mitigations.*

## Mitigations Page UI

The **Mitigations** page provides information about mitigations currently underway in your organization. The page includes the following UI elements:

* **Manage Mitigations** (subnav button): Opens the **Manage Mitigations** page (**Settings »** **Mitigations**).
* **BGP Announcements** (subnav button): Opens the **BGP Announcements** dashboard, providing a centralized view of all BGP announcements tied to your active Kentik Protect mitigations.
* **Actions** (subnav dropdown): Choose to export or subscribe to the current view:

  + **Export**: Export the page’s content as a visual report (PDF). A notification appears when the export is ready to download.
  + **Subscribe**: Opens the **Subscription Dialog**, which enables you to subscribe to regular reports from the page, either by choosing an existing subscription (combination of email address and schedule) or specifying a new one.
* **Start Manual Mitigation**: A button that opens the **Start Manual Mitigation Dialog**.
* **Show/Hide Filters** (filter icon): A button that toggles the **Filters** pane between expanded and collapsed.
* **Group By**: Choose a property (e.g., status) from the dropdown menu to group the mitigations in the table by the value of that property. The table supports grouping by status, policy, platform, method, target, day, or week.
* **Search** (field): Narrow the mitigations shown in the Mitigations list. If text is entered the list will show only mitigations that match the text in at least one column. The field will also display any filters applied with the **Filters** pane.
* **Filters** (pane): Filters that narrow the mitigations listed in the **Mitigations List** (see **Mitigations List Filters**).
* **Action controls**: Present only when one or more checkboxes are selected in the Mitigations list, enabling you to apply an action to all selected mitigations:

  + **Action buttons**: Buttons that apply the actions detailed in **Mitigation Actions**.
  + **Selection indicator**: Indicates how many alerts are currently selected.
* **Mitigations List**: A table listing mitigations (see **Mitigations List**).

## Mitigations List

The Mitigations List on the Mitigations page is a filtered table (see **Mitigations List Filters**) providing information about mitigations triggered by your organization's alert policies. Each row in the table represents an individual mitigation. The table includes the following columns:

* **Select All** (in heading row): A checkbox for toggling the selection state of all mitigations in the list:

  + If either no checkboxes in the list itself are checked or only some are checked then clicking this checkbox will select all mitigations.
  + If all checkboxes in the list are checked, clicking this checkbox will deselect all mitigations.
* **Select** (in mitigation rows): Check the box to select this mitigation. The mitigation will be included in any action applied with the Action controls (see **Mitigation Actions**).
* **Status**: Shows the current state of the mitigation (see **Mitigation Status**).
* **Mitigation ID**:  The Kentik-generated unique ID for the mitigation.
* **Policy**: Displays the name of the alert policy that triggered the mitigation, or if the mitigation was performed manually.
* **Alert ID**: The Kentik-generated unique ID for the alert that triggered the mitigation. This will be blank for manual mitigations.
* **Platform**: The mitigation platform associated with this mitigation.
* **Method**: The mitigation method associated with this mitigation.
* **Target**: The entity to which this mitigation is applied, which may be an IP/CIDR or something else defined by, for example, a Flowspec filter, such as a protocol or port number.

  + This column also displaysany dimensions that may apply to the mitigation (automatic only) under the target.
  + If a mitigation was triggered by an alert policy (automatic mitigation), this column gives the alert key(s) that matched the condition specified in the alert policy threshold and thus triggered the alert, which in turn triggered the mitigation.
* **Date**: The date-time at which the mitigation began.
* **Min. Time Remaining**: An estimate of how much time remains for the mitigation.
* **Action** (vertical ellipsis): A popup with **Mitigation Actions** that may be available for this mitigation.

> **Note**: To see further details about an individual mitigation, click the mitigation’s row to open a **Mitigation Details Drawer** that slides out from the right of the page.

## Mitigation Status

The status shown for each row in the Mitigations list may apply to either an automated mitigation or a manual mitigation.

> **Note:** Cloudflare applies Magic Transit mitigation only when traffic volume exceeds protocol-dependent minimums (100K pps for TCP or UDP; 60K pps for ICMP or GRE). If the platform of a mitigation in the Alerting list is Cloudflare MT and the traffic volume of the alert policy threshold that triggered the mitigation is below these minimums, then the mitigation state may be indicated as active even when Cloudflare isn’t actually mitigating.

The following table lists the statuses displayed in the **Status** column (label) for automatic and manual mitigations in the Mitigations list, as well as the mitigation filter status to which each label corresponds:

| Status Filter | Status Column | Description |
| --- | --- | --- |
| **Active** | Active | The mitigation is active. |
| **Active** | Clearing | Disabling/removing a mitigation on a remote platform. This is an interim state whose destination state is Acknowledgement Required. |
| **Active** | End Grace | The mitigation has ended but the grace period has not yet expired (see Grace period in **Automated Mitigation Settings**). Only applies to automatic mitigations. |
| **Active** | Starting | Mitigation is being added or activated on the third-party mitigation platform. |
| **Failed** | Failed to clear | Unable to disable/remove a mitigation on the remote platform. |
| **Failed** | Failed to start | Mitigation was attempted but could not be added or activated on the third-party mitigation platform. |
| **Inactive** | Archived | The mitigation is no longer active and is not awaiting user acknowledgement. Inactive mitigations do not appear in the Mitigations List by default. To see them displayed, select Inactive under Status in the **Mitigations List Filters**. |
| **Inactive** | Clear | The mitigation is cleared and no longer active. Inactive mitigations do not appear in the Mitigations List by default. To see them displayed, select Inactive under Status in the **Mitigations List Filters**. |
| **Waiting** | Acknowledgement Required | The mitigation is no longer active:   * If user acknowledgement is required (see **Automated Mitigation Settings**), the mitigation will wait in this state. * If no acknowledgement is required, the mitigation will proceed to the Archived state. |
| **Waiting** | End Wait | Mitigation stop is pending: The conditions that triggered the mitigation no longer exist but one of the following is required before stopping (see Clear Mitigation in **Threshold Mitigations**):   * Expiration of timer; * User acknowledgement. |
| **Waiting** | Start Wait | Mitigation start is pending: Mitigation has been triggered but requires one of the following before starting (see Apply Mitigation in **Threshold Mitigations**):   * Expiration of timer; * User acknowledgement. |

## Mitigation Actions

Available mitigation actions vary depending on the current state of the mitigation and whether the mitigation is automatic (see **Automatic Mitigation Actions**) or manual (see **Manual Mitigation Actions**).  
  
Actions can be applied to mitigations in a couple of different ways:

* **Checkboxes**: When one or more **Select** checkboxes are checked in the Mitigations list, the list is shifted down to reveal the Action controls (see **Mitigations Page UI**), which include buttons that enable you to apply an action to all selected alerts.
* **Action** **menu**: The popup menu from the vertical ellipsis at the right of a given alert's row in the Mitigations list enables you to apply an action to that alert.

### Automatic Mitigation Actions

The table below shows the action buttons that may be available in a row representing an automatic mitigation and in the top right corner of the **Mitigation Details Drawer** for that mitigation. Both the available actions and the results of those actions vary depending on the current state of the mitigation (see **Mitigation Status**).

> **Note:** When you take manual control of an automatic mitigation, the mitigation won’t stop automatically even when the triggering alert is cleared; it will continue until manually stopped or removed.

| Status Filter | Status Description | Available Actions | Action Description |
| --- | --- | --- | --- |
| **Active** | **Active**: Mitigation is currently active. | Stop | Stop the mitigation. |
| **Active** | **Clearing**: Mitigation is in the process of being ended. | Take manual control | Take manual control of the mitigation without affecting the stop process. |
| **Active** | **End Grace**: Mitigation has ended but the grace period has not yet expired (see **Automated Mitigation Settings**). | Take manual control  Skip EndGrace, go to EndWait | Take manual control of the mitigation, which remains active.  Advance the mitigation immediately to the next state. |
| **Active** | **Starting**: Start has been requested but hasn’t yet succeeded. | Take manual control, go to Manual Starting | Take manual control and continue starting the mitigation. |
| **Failed** | **Failed to clear**: Stop has been requested but hasn’t yet succeeded. | Archive  Retry | Archive the mitigation to remove it from the Mitigations List.  Try again to stop the mitigation. |
| **Failed** | **Failed to start**: Mitigation has failed to start automatically. | Take manual control  Retry | Take manual control, at which point you can either retry or archive the mitigation.  Retry the mitigation. |
| **Inactive** | **Archived** | None | No actions are available once the mitigation is archived. |
| **Waiting** | **Ack Required**: Waiting for user acknowledgement. | Acknowledge  Take manual control | Once acknowledged, the mitigation is removed from the Mitigations List.  Take manual control and leave the mitigation stopped. |
| **Waiting** | **End Wait**: The triggering conditions no longer exist but the mitigation is waiting for user acknowledgement or expiration of timer. | Take manual control  Skip EndWait, go to Clearing | Take manual control of the mitigation, which remains active.  Advance the mitigation immediately to the next state. |
| **Waiting** | **Start Wait**: Waiting for user acknowledgement or expiration of a timer. | Approve and start the mitigation  Take manual control, go to Manual Clear | Immediately start a mitigation that has been waiting for acknowledgement or expiration of timer (see **Mitigation Settings**). Take manual control and cancel the pending mitigation. |

> **Note:** To see archived mitigations displayed, select **Inactive** under Status in the **Mitigations List Filters**.

### Manual Mitigation Actions

The table below shows the action buttons that may be available in a row representing a manual mitigation and in the top right corner of the **Mitigation Details Drawer** for that mitigation. Both the available actions and the results of those actions vary depending on the current state of the mitigation (see **Mitigation Status**).

| Status Filter | Status Description | Available Actions | Action Description |
| --- | --- | --- | --- |
| **Active** | **Active**: Mitigation is active. | Stop | Stop the mitigation, at which point you can either restart or archive it. |
| **Active** | **Clearing**: Stop has been requested but hasn’t yet succeeded. | Start | Restart the mitigation. |
| **Active** | **Starting**: Mitigation is being activated. | Stop | Stop the mitigation, at which point you can either restart or archive it. |
| **Failed** | **Failed to clear**: Stop was requested but didn’t succeed. | Archive  Retry | Remove the mitigation from the Mitigations List.   Try again to stop the mitigation. |
| **Failed** | **Failed to start**: Mitigation was attempted but could not be added or activated | Archive  Retry | Remove the mitigation from the Mitigations List.   Try again to start the mitigation, at which point you can either stop it or archive it. |
| **Inactive** | **Archived** | None | No actions are available once the mitigation is archived. |
| **Inactive** | **Cleared**: Mitigation is no longer active or waiting. | Start  Archive | Restart the mitigation.   Remove the mitigation from the Mitigations List. |

> **Note:** To see archived mitigations displayed, select **Inactive** under Status in the **Mitigations List Filters**.

## Mitigations List Filters

The mitigations displayed in the Mitigations list can be filtered using the controls in the **Filters** pane on the left. The pane includes the following filters:

* **Reset To Default** (appears only when you’ve specified one or more filters): Click to clear all current filters.
* **Time Range**: A dropdown that displays the time range within which mitigations will be included in the list. Click to set a different time with the **Time Range Selector**.
* **Status**: A set of checkboxes that allow you to select one or more status (Active, Failed, Waiting, or Inactive) to include in the list.
* **Sources**: Filter the list by whether a mitigation was run manually or automatically.
* **Mitigation ID**: Include only mitigations whose ID equals the entered numbers.
* **Alert ID**: Include only mitigations for which the ID of the triggering alarm equals the entered numbers.
* **Method**: Filter the list by selecting a method from the dropdown menu.
* **Platform**: Filter the list by selecting a platform from the dropdown menu.
* **Show Tenant Mitigations**: A switch that determines whether tenant mitigations are shown on the list.
* **Tenants** (appears only when **Show Tenant Mitigations** is toggled): Filters the list by selecting a tenant from the dropdown menu.
* **Policy**: Include only mitigations for which the name(s) of the triggering alert policy equals the one or more policies selected from the dropdown menu. Click in the Policy field and start typing in the Filter options field to see a list of policies with that text and then select one or more from the list.
* **Target Search**: Include only mitigations for which the entered text matches specific instances of the dimension that the mitigation method used to identify which traffic to mitigate.
* **Dimension Search**: Include only mitigations for which the entered text matches the dimension that the mitigation method used to identify which traffic to mitigate.
* **Exact Match**: A switch that determines whether the string entered in the Dimension Search field is matched strictly or loosely.

#### Time Range Selector

This selector is a popup that opens from the dropdown and sets the time range within which mitigations will be included in the list. The popup includes the following controls:

* **Lookback list**: Preset durations back from the current time, listed along the left side of the popup.
* **Calendars**: Side-by-side monthly calendars that enable you to click on a start date and end date.
* **Start date field**: The start of the time range, filled in from the lookback list or the calendars.
* **End date field**: The end of the time range, filled in from the lookback list or the calendars.
* **Apply**: A button that applies the time range from the values in the start and end fields and hides the popup. The applied range will be shown in the Time-range dropdown.
* **Cancel**: A button that closes the popup and leaves the time range as it was before the popup was opened.

## Mitigation Details Drawer

The details drawer for a given mitigation slides out from the right of the page when the row for that alarm is clicked in the **Mitigations List**. The drawer may show the following additional information, depending on the type of mitigation:

* **Action buttons**: Any actions available to take for the mitigation are available in the top of the drawer (see **Mitigation Actions)**.
* **Details**:

  + ![Details of mitigation policy for IP address 10.205.215.4/32 with alerts.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/MT-mitigation-details-drawer.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A22Z&se=2026-05-12T09%3A42%3A22Z&sr=c&sp=r&sig=Sa11HpJ%2BwWAODiOhviTa3a%2FpnGvCBGesfMPH70EALJE%3D)**Mitigation ID**: The Kentik-generated unique ID for the mitigation.
  + **Policy**: The alert policy that triggered the mitigation (if applicable). Click the link to open to the **Alert Policies Page** with that policy displayed.
  + **Alerts**: Click **View triggering alert →** to open the alert that triggered the mitigation. Click **View related alerts →** to open a list of other alerts (if any) that are related to the triggering alert (see **Alert Details Page**).
  + **Method**: The mitigation method associated with this mitigation.
  + **Platform**: The mitigation platform associated with this mitigation.
  + **Minutes Before Autostop**: The configured time limit (in minutes) for a manual mitigation before the system automatically stops it.
  + **Comment**: Optional user-provided notes associated with the mitigation, typically entered when a manual mitigation is initiated.
  + **Announcements**: Displays the number of related BGP announcements. Click the link to view the specific routes on the **BGP Announcements** page.
  + **Policy Evaluation Frequency**: The interval at which the Kentik alerting engine evaluates network traffic against the threshold conditions of the policy that triggered this mitigation.
  + **Inactivity Clear**: The duration of time that must pass without any traffic violating the policy thresholds before the system automatically clears the mitigation.
  + **Grace Period**: A configured buffer period that keeps the mitigation active even after triggering conditions are no longer met. This prevents the mitigation from rapidly toggling on and off (flapping) during intermittent attacks.
  + **Min. Time Remaining**: The estimated minimum time remaining before the mitigation is scheduled to clear or auto-stop.
* **Target**: The specific routing prefixes, IP addresses, and dimensions (e.g., source IP, tcp-flags) to which the mitigation is applied. This defines the exact subset of traffic matching the policy.
* **BGP Monitor Test** (button): If available for the mitigation platform, click to immediately initiate a BGP Monitor test for the targeted prefixes. You’ll be redirected to the **BGP Monitor Details** page within the Synthetics module.

  > **Note**: Clicking this button may bring up a **Confirm BGP Test Creation** dialog. If you are at your credit limit, you will not be permitted to create the test.
* **Event list**: A table listing, in chronological order, events involving the mitigation:

  + **Status**: The mitigation state at the time of the event.
  + **Event**: A description of the specific action, state transition, or system trigger that occurred at that timestamp (e.g., `user manual control`, `remote added`, or `skip wait`).
  + **Date (UTC)**: The date-time of the event.

The subnav is the gray, horizontal bar or strip located below the main navigation bar or navbar in various pages and modules. It typically contains page-wide controls and buttons specific to the page's functionality.
