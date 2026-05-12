# NMS Policy Settings

Source: https://kb.kentik.com/docs/nms-policy-settings

---

This article discusses the **Add/Edit NMS Alert Policy** settings pages in the Kentik portal. For background on Kentik’s Alerting system, first read these articles:

* For the step-by-step procedures for managing alert policies, see **Manage Alert Policies**.
* For an NMS alert policies overview, see **About NMS Policies**.
* For other alerting articles, see **Policy Alerts Overview**, **Alerting**, **Notification Channels**, and **Mitigations**.

> **Note**: Kentik has deprecated **Up/Down** policies in favor of the more powerful **NMS** policies described in this article. Existing **Up/Down** policies remain available in the portal and can be managed with the legacy policy settings page.

![Add an NMS alert policy with a NMS Policy Settings page.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AP-NMS_Add_Policy_General.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A27Z&se=2026-05-12T09%3A33%3A27Z&sr=c&sp=r&sig=Ee3S2JqkXEqlsn1s8%2B4MqmfBjZ5rh6Qa65wquLyZrEo%3D)

## Add/Edit NMS Alert Policy Page

This section covers the UI elements of the Add/Edit NMS Alert Policy pages.

Access these settings pages as follows:

* **Add NMS Alert Policy Page**: Alerting » Alert Policies » **Add Alert Policy**. Select the NMS card.
* **Edit NMS Alert Policy Page**: Alerting » Alert Policies. Click **Edit Policy** from the policy’s action menu (vertical dots icon).

### Subnav

The subnav elements are as follows:

* **Cancel**: Click to exit without saving.
* **Create** (Add only): Click to create the new policy and return to the Policies page. Only active when all required settings are complete and error-free.
* **Save** (Edit only): Click to save changes and exit.
* **Remove** (Edit only): Click to remove the alert policy.

### General Tab

The **General Tab** settings:

* **Name**: The name of the policy. Maximum of 50 characters; must include at least one letter.
* **Description** (optional): A description of the policy; used to summarize what the policy looks at and indicate what it is used for.
* **Enable Policy**: Defaults to On. Click Off to disable the policy.

### Target & Filters Tab

The **Target & Filters Tab** settings:

* ![Interface for adding an NMS alert policy with target and filter options displayed.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AP-add-policy-dialog-type.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A27Z&se=2026-05-12T09%3A33%3A27Z&sr=c&sp=r&sig=Ee3S2JqkXEqlsn1s8%2B4MqmfBjZ5rh6Qa65wquLyZrEo%3D)

  *Select a policy type for a new NMS alert policy.*

  **Policy Type**: Select a type for the new policy, one of Agent, BGP Neighbor, Component, Device, Event, Interface, or Custom (see **NMS Policy Types**).
* **Measurements:** Click the field to select metrics from the dropdown (filtered by Policy Type). Displays lozenges for selected metrics. Click the **X** to remove a lozenge.
* **Devices**: Select devices for this policy to monitor:

  + **Selected devices list**: Shows lozenges for selected devices. Click **X** to remove a lozenge.
  + **Edit Devices** (button): Opens the **Data Sources**dialog.
* **Dimension Filters**: Add filters to narrow the scope of dimension values this policy will monitor.

  + **Selected filters list**: Shows cards for dimension filters as they are added.
  + **Add Dimension Filters** (button): Click to open the Filters dialog (see **Filters Settings**). Available for Agent, BGP Neighbor, Component, Device, Interface policy types only.
  + **Edit Dimensions** (button): Click to open the **Group By Dimensions** dialog. Available for Event policy type only.

### Activate & Clear Tab

The **Activate & Clear Tab** settings:

![Configuration settings for alert conditions with a major severity error highlighted.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AP-add-policy-dialog-severity(1).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A27Z&se=2026-05-12T09%3A33%3A27Z&sr=c&sp=r&sig=Ee3S2JqkXEqlsn1s8%2B4MqmfBjZ5rh6Qa65wquLyZrEo%3D)

*Configuration settings for alert conditions with a major severity error highlighted.*

* **Severity**: Choose a severity level for alerts activated by this policy. Default is Major.
* **Alert Condition Groups**: Add groups of conditions to be met before the policy triggers alerts.

  + Any/All dropdown for top-level condition group
  + Cards for conditions in the groups

    - Any/All/None dropdown for conditions
    - **Add Condition**: Opens the

      * Measurement/Metric dropdowns (filtered by Policy Type)
      * Comparison dropdown: Equals, etc.
      * Value: Enter value in the format indicated in red (e.g., float, percent as float)
      * Delete condition (X)
    - **Add Nested Condition Group**: Add up to 3 levels of nested condition groups within each top-level group.
    - **Remove Group**: Removes a nested condition group.
  + **Add Condition Group**: Add a new top-level condition group.
* **Acknowledgment Required**: Enable the switch to require a user to acknowledge this policy’s alerts (see **Alerts List****)**.
* **Manual Clear Required**: Enable the switch to require manual clearing of alerts that are no longer in an alarm state (see **Alerts List****)**.
* **Notification Channels**: View and assign notification channels to this policy (see **About Notification Channels**):

  + **Assign a Channel**: Click to drop down a filterable channel list and select a channel. Repeat as needed.
  + Assigned channels show as lozenges in the field.
  + Click the **X** in a lozenge to remove the channel.
* **Add New Channel**: A button that opens the Add Notification Channel dialog (see **Notification Dialogs UI**) where you can create a new channel and add it to the policy automatically.
* **Test Notification Channels**: Click to send test notifications to all assigned channels. Active only when notification channels are selected.
* **Activation Delay**: Enter the number of minutes to wait before starting the alert. If the alert would have cleared within this period, the alert won’t start at all.
* **Clearance Delay**: Enter the number of minutes to wait before clearing the alert. If the alert would have restarted within this period, the alert will not be cleared.

#### Data Sources Dialog

The Data Sources dialog allows you to add devices to the alert policy, including:

* **Search**: Enter terms to filter the Device Selector list.
* **Cancel**: Click the **X** at the upper right, or **Cancel** at the lower right, to exit the dialog without saving.
* **Device Selector:** Left pane listing methods of selecting devices:

  + **All Devices** (default): Include all devices.
  + **Labels**: Add devices by selecting one or more labels (see **Labels**).
  + **Sites**: Add devices by selecting one or more sites (see **Sites**).
* **View Selections**: Right pane summarizing device selections:

  + Lozenges for each selected label/site
  + Click **X** to remove the selection
  + Click **>** on a site lozenge to show assigned devices
* **Save**: Click **Save** to confirm the device selections and exit.

## NMS Policy Types

The following table describes the policy types for NMS alert policies, listing the related metrics for each:

| Policy Type | Description | Metrics |
| --- | --- | --- |
| Agent | Metrics reported by Kentik agents | * /kentik/agent/cpu    + Times   + Utilization * /kentik/agent/disk    + Free   + Utilization * /kentik/agent/memory    + Buffers   + Cached   + Free   + SReclaimable   + Total   + Utilization * /kentik/agent/status    + Agent Operational status   + Is Agent Running * /kentik/agent/status/alert    + Should trigger alert |
| BGP Neighbor | Protocol metrics for BGP | * /protocols/bgp/neighbors    + Messages In Total   + Messages In Updates   + Messages Out Total   + Messages Out Updates   + Session State   + State transitions   + Time Established * /protocols/bgp/neighbors/prefixes    + advertised   + advertised-prefixes   + Installed   + Max Prefix   + Max Prefix Clear Threshold   + Max Prefix Percent   + Max Prefix Threshold   + Received   + Received Pre-Policy   + Rejected   + Sent   + Suppressed   + Withdrawn |
| Component | Hardware component metrics | * /components    + Entity Admin Status   + Entity Operational Status   + Fantray Operational Status   + Field Replaceable Unit   + Module Operational Status   + Operational Status   + Uptime * /components/cpu/utilization    + CPU Utilization * /components/fan    + Alarm severity   + Alarm status   + Operational status   + Speed * /components/memory    + available   + Free   + Physical   + Used   + Utilization   + utilized * /components/power-supply    + Current   + Operational status   + Power   + Voltage * /components/temperature    + Temperature |
| Device | System metrics | * /system    + Availability Status   + Average RTT (µs)   + Boot Time   + online   + Status   + Uptime (Seconds) * /system/cpus/summary    + CPU Utilization |
| Event | Metrics for Syslog and SNMP Trap | Syslog (default):   * Timestamp * Severity * Message * Device Name   SNMP Trap (default):   * Timestamp * Trap Type * Device Name  **Note**: Click **Edit Dimension** to adddimensions other than the defaults. |
| Interface | Interface counters | /interfaces/counters |
| Custom | Select any available measurement | All |
