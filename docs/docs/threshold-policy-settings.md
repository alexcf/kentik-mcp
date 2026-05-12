# Threshold Policy Settings

Source: https://kb.kentik.com/docs/threshold-policy-settings

---

The settings for threshold policies in Kentik's alerting system are detailed here. For more information about policy-based alerting, first read these articles:

* For an overview of threshold policies, see **About Threshold Policies**.
* For general policy-based alerting info, see **Policy Alerts Overview**.
* For policy templates info, see **Policy Templates**.
* For alerting page info, see **Alerting**.
* For alert notifications info, see **Notification Channels**.
* For alert mitigation info, see **Mitigations**.

> **Note:** Alert policy configuration can be complex. Please reach out for help by contacting Kentik support (see **Customer Care**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(296).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A09Z&se=2026-05-12T09%3A58%3A09Z&sr=c&sp=r&sig=f4yfZhAv7wbfmN8qlesVw7mWFygiRIyzKqZ89dgYPkI%3D)

*The settings page for a threshold policy.*

## Threshold Settings Tabs

The following tabs are included on all types of settings pages (add, edit, or clone; see **Policy Settings Pages**) for threshold policies:

* **General**: Defines the overall policy properties (see **General Policy Settings**).
* **Dataset**: Narrows the subset of traffic evaluated for all thresholds (see **Policy Dataset Settings**).
* **Thresholds**: Specifies up to five conditions collections that trigger alert activation (see **Alert State**and **Policy Threshold Settings**).
* **Baseline**: Configures baselines for comparing current traffic to determine deviations from the norm (see **Policy Baseline Settings**).

#### Policy Summary Pane

All Threshold Settings tabs have a Summary sidebar with collapsible cards for each settings tab.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(298).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A09Z&se=2026-05-12T09%3A58%3A09Z&sr=c&sp=r&sig=f4yfZhAv7wbfmN8qlesVw7mWFygiRIyzKqZ89dgYPkI%3D)Each card has:

* **Summary indicator**: Indicates the tab’s overall settings status:

  + Complete (green circle with checkmark): All required settings on the tab are complete and valid.
  + Incomplete(empty circle): One or more settings on the tab aren't specified.
  + Error (red octagon with X): One or more settings on the tab need to be corrected.
* **Heading**: The tab name for the settings summarized in the card.
* **Count** (Thresholds only): The number of active thresholds in the policy.
* **Expand/Collapse**: Click the arrow to toggle the expanded/collapsed card.
* **Summary**: A read-only list of the tab’s settings, with indicators for any errors that need to be corrected.

  > **Note**: Policies with incomplete settings and/or errors can't be saved.

## General Policy Settings

The General Settings tab of the policy settings pages (add, edit, clone) defines the overall policy properties. The tab includes:

* **Name**: ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(300).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A09Z&se=2026-05-12T09%3A58%3A09Z&sr=c&sp=r&sig=f4yfZhAv7wbfmN8qlesVw7mWFygiRIyzKqZ89dgYPkI%3D)The policy name. Maximum of 50 characters. Must include at least one letter.
* **Description** (optional): The policy description, summarizing what it looks at and what its used for.
* **Policy Status**: A switch to enable or disable the policy. Defaults to enabled.
* **Labels**: A selector to apply **Labels** to the policy:

  + Click to select labels from the dropdown, filter by label name as necessary.
  + Click **Add Label** to open the **New Label Dialog** to apply the labels.
  + Assigned labels show as lozenges in the field.
* **Policy Type** (DDoS or Custom policies only): A dropdown to set the policy type to DDoS or Custom. Query-based policies can only be created from Data Explorer (see **Add a Query-based Policy**).
* **Policy Expires** (Query-based policies only): Sets the length of time after which the policy will be disabled. Options include: Never, 15 days, 30 days, 60 days, or 90 days. Defaults to 90 days.

  > **Notes**:
  >
  > + If you turn off the policy status switch, the expiration switches to Never.
  > + If the **Policy Expires** time is set to Never, the policy type changes from Query-based to Custom.
* **Suppress Alerts**: A switch to turn on **Alert Suppressions** for this policy until a specified date, e.g., while establishing a baseline for new policies. When suppressions are on, the policy won't enter alarm state and won’t trigger notifications and/or mitigations.

  > **Notes:**
  >
  > + Alert suppressions are enabled by default for query-based policies.
  > + Alert suppressions can be enabled on a pattern basis on the **Alert Suppressions & Silences Page**.
* **Suppress Alerts End Date**: If **Suppress Alerts** is switched on, specify the date on which the alert suppressions will end. The default is seven days.
* **Policy Dashboard** (not present for NMS policies): Click to choose an existing **Dashboard** from the dropdown, filtering by dashboard name as necessary. This sets the destination dashboard for the **Open****Dashboard** button for this alert on the Alerting page (see **Alert-Specific Actions**).

## Policy Dataset Settings

The Dataset tab of the Policy Settings pages narrows the subset of traffic evaluated for the alert’s thresholds. The tab is divided into two panes, **Data Funneling**and **Building Your Dataset** as described below.

> **Note**: Policy metrics can't be changed while a policy is used by tenants, as indicated in the following locations:
>
> * **Policies List**: Tenant names are listed in the **Used by Tenant** column.
> * **Edit Policy page**: A banner on the **Dataset** tab shows how many tenants use the policy and has links to the **Tenant Policy Settings** for each tenant.

### Data Funneling

The **Data Funneling** pane’s settings (Dataset tab) define which network traffic will be evaluated by the policy:

* **Data Sources**: View and update the policy's data sources:

  + Data Sources list: Lists the devices and cloud resources that carried the traffic evaluated by this policy.
  + Edit Data Sources: Opens the **Data Sources Dialog** to update the data sources covered by this policy.
* **Policy Dimensions**: Lists the traffic evaluation dimensions and allows for selecting dimensions (see **Policy Dimensions**).
* **Metrics**: Lists the policy’s metrics and allows for selecting primary and secondary metrics (see **Policy Metrics**).
* **Filters**: Filter the traffic being evaluated for the alert:

  + Filters list: Lists the filters currently applied to this policy (see **About Filters**).
  + Edit Filters: Opens the Filtering Options dialog (see **Filtering Options Dialog**) to edit filters.

> **Note:** A mitigation can only be assigned to a policy if the dimensions include source or destination IP/CIDR (see **Threshold Mitigations**).

#### Policy Dimensions

The Policy Dimensions section (required, Dataset tab) shows the dimensions used for traffic evaluation and allows for selecting dimensions:

* **Dimension list**: The dimensions (see **About Dimensions**) that combine to define a key.
* **v4 CIDR & v6 CIDR**: Set custom IPv4 and IPv6 prefixes when CIDR dimensions are applied to the policy. Available for the following dimensions: Destination IP/CIDR, Source IP/CIDR, Source Next Hop IP/CIDR, and Destination Next Hop IP/CIDR.
* **Edit Dimensions**: Opens a dialog (see **Dimension Selectors**) to edit the dimensions (up to 8) of the key. The key definition determines how traffic is subdivided for evaluation (see **About Keys**).  
  **Example*:*** If the primary metric is packets/second (pps) and the group-by dimensions are “Source:AS Number” and “Destination:AS Number”, the top-X evaluation looks at all unique combinations of source/destination ASN for those with the highest traffic volume in pps.

#### Policy Metrics

The Metrics section (Dataset tab) shows the policy’s currently selected metrics and allows access for managing primary and secondary metrics:

* **Primary**: A lozenge for the metric selected as the primary unit (e.g., bits/s, packets/s) by which traffic will be evaluated to determine top-X (see **General Metrics** and **Host Traffic Metrics**).  
  **Example*:*** If the primary metric is bits/second (bps), the top-X evaluation will rank keys by bps.
* **Secondary**: A lozenge for the secondary metric (up to 2), used to specify multiple additional static comparators (see **Policy Threshold Settings**). Each added comparator is a condition for triggering an alarm. Use the **Secondary Metrics** selector to choose up to 2 secondary metrics.
* **Edit Metrics**: Opens the Metrics dialog (see **About the Metrics Dialog**) to edit the primary and secondary metrics for the policy.

  > **Note**: Each policy can have a maximum of three metrics: one primary and two secondary.

### Building Your Dataset

The Building Your Dataset pane settings specify how network traffic defined in the Data Funneling pane will be evaluated. By default, the pane is hidden. To show the pane, click **Advanced Settings** at the bottom of the tab.

The pane always includes the following controls:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(302).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A09Z&se=2026-05-12T09%3A58%3A09Z&sr=c&sp=r&sig=f4yfZhAv7wbfmN8qlesVw7mWFygiRIyzKqZ89dgYPkI%3D)**Evaluation Frequency**: The interval at which newly ingested traffic data is grouped by dimensions and the traffic data represented by the resulting keys is evaluated. Options include 60 seconds (default), 2 minutes, and 5 minutes.
* **Maximum Keys Per Evaluation**: The number of keys (unique combinations of dimension values; see **About Keys**) to evaluate for a match with the conditions specified in the alert's thresholds. Maximum valid value is 300.
* **Minimum Traffic Threshold**: The minimum value, as measured by the primary metric, that a key must have to be included in the top-X ranking and evaluated for a threshold match.

  + Specified value: Specify the minimum value (measurement units are determined by the primary metric).
  + Auto-calculate: If checked (default), the value is auto-calculated using a formula based on the settings of the comparators in the policy’s thresholds (see **Threshold Conditions**) and the **Minimum Traffic Threshold** field is greyed out.

If you’ve included more than one dimension, the pane also includes the following **Dimension Grouping** controls (see **About Dimension Grouping**):

* **Group Dimensions** (switch): Enables grouping by dimensions before the final top-X evaluation.
* **Dimensions to Group By** (shown only if Group Dimensions is enabled): Enter the number of dimensions, starting at the top of the Dimensions list, to include for grouping.
* **Maximum Keys per Group** (shown only if Group Dimensions is enabled): Enter the number of keys from each group to be included in the overall top-X for the alert. The maximum valid value is 300 or the value of the **Maximum Keys per Evaluation** setting.

#### About Dimension Grouping

Dimension grouping adds control to tracking the top-X keys for current traffic. It helps prevent keys from a high-volume area from dominating the top-X keys, allowing the alerting system to detect significant changes in other areas.

With dimension grouping disabled, the key definition (set of dimensions) specified in the **Dimensions** field (see **Data Funneling**) is used as a single unit, resulting in unique combinations of dimension values. These keys are then ranked by traffic volume to arrive at Top-X.  
  
 With dimension grouping on, traffic is evaluated in stages:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(305).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A09Z&se=2026-05-12T09%3A58%3A09Z&sr=c&sp=r&sig=f4yfZhAv7wbfmN8qlesVw7mWFygiRIyzKqZ89dgYPkI%3D)The Data Funneling section has two sets of dimensions: the grouping set and the non-grouping set.

  + The grouping set consists of the first dimension in the Policy Dimensions field and the dimension whose position in the Dimensions field corresponds to the number specified with the Dimensions to Group By control.
  + The remaining dimensions are in the non-grouping set.
  + **Example**: Dimensions are Source ASN, Full Device, Destination Country, and Destination ASN. If **Dimensions to Group By** is set to 2, the grouping set is Source ASN and Full Device.
* Traffic is initially evaluated as if the key definition is only the grouping set. Each group’s keys represent traffic with a unique combination of the grouping set dimensions (e.g., Source ASN and Full Device).
* Traffic in each group is then evaluated as if the key definition is only the non-grouping set. Each group’s keys represent traffic with a unique combination of the non-grouping dimensions (e.g., Destination Country and Destination ASN).
* The top N keys from each group are merged into a single pool, with N determined by the **Maximum Keys per Group** control.
* The keys in this pool are ranked by volume, resulting in the overall top-X keys for the alert, with X being defined by the **Maximum Keys Per Evaluation** control.

## Policy Threshold Settings

An alert policy’s threshold settings define a set of conditions that must match for the policy to be activated. Upon activation, the policy generates an alert for each key for which conditions have been matched. Each policy includes at least one threshold by default and can include up to five (Critical, Severe, Major, Warning, Minor).

### General Threshold Settings

The **Thresholds** tab includes the following general settings:

* **Diagram**: ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(306).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A09Z&se=2026-05-12T09%3A58%3A09Z&sr=c&sp=r&sig=f4yfZhAv7wbfmN8qlesVw7mWFygiRIyzKqZ89dgYPkI%3D)A visualization of the current policy configuration. Visible only if the policy is enabled.

  + Refresh: Click to refresh traffic data.
  + View in Data Explorer: Click to open the query in Data Explorer (new tab).
* **Severity selector**: Choose the threshold you are configuring: Critical, Severe, Major, Warning, or Minor.
* **Threshold Status**: Determines if the threshold is enabled (evaluating traffic data, generating alarms, etc.). Disabled thresholds can be retained for future use.
* **Threshold Description**: Enter a description of the traffic conditions monitored by the threshold.

### Threshold Conditions

The settings in the Conditions section determine what constitutes a match for the threshold (see **About Matches**). These settings are configured in the Edit Threshold Conditions dialog, accessed via the **Edit Conditions** button.

> **Notes:**
>
> * All conditions for a given threshold must be met for an alert to trigger.
> * The settings for a given threshold are independent of other thresholds.

#### Configuration settings for traffic alert thresholds in network policy management interface.Conditions General UI

The Edit Threshold Conditions dialog includes the following:

* **Cancel buttons**: Click the **X** in the upper right or the **Cancel** button at lower right to close the dialog without saving.
* **Conditions tiles**: Each tile defines a condition (see **Conditions Tile Types**).
* **Apply button**: Click **Apply** to apply conditions to the policy and return to the policy settings page.

#### Conditions Tile Types

Policy threshold conditions (Critical, Severe, etc.) are set in the Edit Threshold Conditions dialog. Each tile specifies a condition of a specific type. Some condition types might be disabled based on current policy settings (e.g., the metrics set on the policy's Dataset tab). If a tile is inactive, a brief notification will appear at the right.  
  
By default, the dialog includes tiles for these condition types:

* **Static Condition**: Sets conditions based on primary and secondary metrics from the policy's Dataset tab (see **Static Condition**). Click **Show Chart** to display a visualization of the last three days’ data.

  > **Note**: The default primary metric for a new policy is Packets/s.
* **Baseline Condition**: Sets a baseline condition, if enabled (see **Baseline Condition**).
* **Top Keys Condition**: Sets a top keys condition, if enabled (see **Top Keys Condition**)**.**
* **Interface Capacity Condition**: Sets an interface capacity condition, if enabled (see **Interface Capacity Condition**).
* **Ratio-Based Condition**: Sets a ratio-based condition, if enabled (see **Ratio-Based Condition**).

> **Note:** Enabled conditions are displayed in the Conditions section of the corresponding threshold tab (e.g., Critical, Severe).

#### Conditions Tile Common UI

Every conditions tile in the Edit Threshold Conditions dialog includes the following controls:

* **Activate**: A switch that activates/deactivates the condition. When activated, the tile expands, allowing you to specify settings.
* **Condition statement**: Indicates the applied condition and its settings. Condition statements vary by condition type.
* **Show Chart/Hide Chart** (static and ratio-based conditions only): Opens a chart showing the last three days of data for the specified metrics. The static/baseline metric is displayed in orange. If shown, **Hide Chart** hides it.
* **View in Data Explorer** (static chart only): Links to Data Explorer, populating the Query pane with this policy’s Dataset tab settings.
* **Refresh** (static chart only): Refreshes the visualization data.

#### Static Condition

A static condition compares a specified traffic metric to a fixed value in the tile. Use these to controls to specify the condition:

* **Condition value**: Enter a metric value (number) for comparison.
* **Condition metric**: Select the units for the condition value from the drop-down.

The resulting condition statement is a sentence, e.g., "If traffic is at least `100 packets/s`." The static condition is met when the current metric value meets or exceeds the specified value.

> **Notes**:
>
> * The default primary metric for a static condition is packets/s. You can change it or add secondary metrics in the Metric pane of the Dataset tab. The Edit Threshold Conditions dialog reflects the current Metrics pane settings.
> * A static condition is required to set a **Baseline Condition**.

#### Baseline Condition

The Baseline Condition tile compares the current traffic metric value to a historical value from a specified period. It’s enabled only when the Static Condition tile for the primary metric is activated and the Top Keys Condition tile is not.

It has the following settings:

* **Comparison value**: Enter a number for comparison.
* **Unit of measure**: Select the unit of measure (% or Units) for the comparison value:

  + *%*: Current metric value varies from the baseline by the percentage specified.
  + Units over baseline: Current metric value varies from the baseline by the number of units (e.g., bit/s, packets/s) specified.
* **Comparison direction**: A drop-down to choose the direction (Above or Below):

  + Above: The current value exceeds the baseline.
  + Below: The current value falls below the baseline.

The resulting condition statement is a sentence based on the enter values, e.g., "If traffic is at least `20 % above` the baseline."

#### Top Keys Condition

The Top Keys Condition tile is enabled only when the Baseline Condition is disabled. It checks for changes in the composition of the top-X keys (joining or leaving). The following controls determine how the top keys are evaluated:

* **Comparison type**: Choose the comparison type (Joins or Leaves) from the dropdown:

  + Joins: The key joins the top-X keys in current traffic.
  + Leaves: The key leaves the top-X keys in current traffic.
* **Top Keys**: The number of keys to evaluate.

  > **Note:** If left blank, it defaults to the value of the **Maximum Number of Keys** field in **Building Your Dataset** (default = 25).

The resulting condition statement is a sentence based on the entered values, e.g., "If a given key `joins` the top `25` keys in current traffic."

#### Interface Capacity Condition

The Interface Capacity tile triggers alerts based on an interface’s capacity relative to its current utilization. This requires a data source with traffic included in the policy (see **Data Funneling**). The tile is disabled unless:

* The policy's dimensions include Source/Destination Interface (see **Policy Dimensions**).
* The policy's metrics include bits/s (see **Policy Metrics**).
* No baseline condition is set in the **Baseline Condition** tile.

This condition is specified with:

* **Comparison value**: Enter a number for comparison.
* **Comparison method**: Choose the comparison method (Mbits/s or %) from the drop-down:

  + Mbits/s: Current utilization of the interface (traffic volume) in bits/s exceeds its capacity by at least the comparison value.
  + %: Current utilization of the interface, as a percent of its capacity, exceeds the comparison value.

The resulting condition statement is a sentence based on the entered values, e.g., "If utilization is at least `100 Mbits/s` over interface capacity."

#### Ratio-Based Condition

The Ratio-Based Condition tile sets a policy condition based on the ratio of any two metrics (see **Data Funneling**). It’s disabled unless the policy's dataset includes at least two metrics.  
  
The condition is defined with controls that designate a metric as "Left" and another as "Right", establishing a ratio:

* **Metrics**: Left/right metric controls, each with:

  + Metric value: Enter a number for the metric.
  + Metric: Choose a metric from the dropdown.
* **Swap Metrics** (left/right arrow icon): Click to swap the left/right metric controls.
* **Margin** (only when Bidirectional is on): Enter a fractional value to add to both parts of the ratio, e.g., a margin of 0.2 means a ratio of 1.2:1 is used in both directions.
* **Bidirectional**: A switch that controls how the two parts of the ratio are evaluated:

  + Off: The condition is met if the left/right metric ratio exceeds the ratio defined by their metric values. For example, if the left value is 10 and the right value is 1 then the condition is met when the ratio of the left metric to the right metric exceeds 10:1.
  + On: The condition is met when the metrics ratio in either direction exceeds the ratio defined by their metric values. For example, if the ratio exceeds 10:1 in either direction (left/right or right/left).

### Threshold Configuration

The Threshold Configuration settings pane (accessed via **Advanced Options** on the Thresholds tab) varies based on the Edit Threshold Conditions dialog settings:

* ![Configuration settings for traffic evaluation thresholds and key management in a network system.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AP-Threshold_config.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A09Z&se=2026-05-12T09%3A58%3A09Z&sr=c&sp=r&sig=f4yfZhAv7wbfmN8qlesVw7mWFygiRIyzKqZ89dgYPkI%3D)**If no baseline exists** (only when a **Baseline Condition** is set for the primary metric): Choose what to do when a key in the primary top-X is not present in the comparison top-X by selecting an option from the dropdown (see **No Existing Baseline**).
* **Default Value** (only when a **Baseline Condition** is set for the primary metric): Enter the value to use for comparison when **If no baseline exists** is set to “Use a Default Value” or when the Lowest/Highest Value can’t be determined from top keys (e.g., baselining hasn't yet started).
* **Top Keys** (always present): Enter numbers to override the dataset's **Maximum Keys Per Evaluation** setting (in parentheses to the right of each field, see **Building Your Dataset**):

  + Current keys: How many top keys to evaluate from current traffic.
  + Baseline keys: How many top keys to store for the baseline.

#### No Existing Baseline

The **If no baseline exists** setting is shown only when a **Baseline Condition** is set for the primary metric (see **About Historical Baselines**). It’s a drop-down that determines what to do if a key in the primary top-X is missing in the comparison top-X:

* **Use Lowest Value from Top Keys** (default): Compares the key’s current value in the primary top-X to the lowest key in the comparison set.
* **Use a Default Value**: Compare this key’s current value to the static value in the comparison set.
* **Do Not Alert**: Classify as not a match.
* **Activate an Alert**: Classify as a match.
* **Use Highest Value from Top Keys**: Compare the key’s current value in the primary top-X to the highest key in the comparison set.

### Threshold Frequency

Kentik's alerting system defines a "match" as an individual network traffic instance that matches the alert policy’s conditions. The Threshold Frequency settings determine how many matches must occur within a specified duration to trigger an alert:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(309).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A09Z&se=2026-05-12T09%3A58%3A09Z&sr=c&sp=r&sig=f4yfZhAv7wbfmN8qlesVw7mWFygiRIyzKqZ89dgYPkI%3D)

* **Times**: The number of matches required within the duration.
* **Duration value**: The number of time units in the duration.
* **Duration units**: The time unit (minutes or hours).
* **Reset period**: The number of match-free minutes after which the count of matches is reset to 0.

  > **Note:** If the Times setting is “1”, the Duration settings are irrelevant; an alert is generated immediately upon the first match.

### Threshold Actions

Threshold Actions settings determine how the alerting system responds to alerts generated by this threshold. These settings include:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(310).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A09Z&se=2026-05-12T09%3A58%3A09Z&sr=c&sp=r&sig=f4yfZhAv7wbfmN8qlesVw7mWFygiRIyzKqZ89dgYPkI%3D)**Acknowledgment Required**: Enable this switch to manually acknowledge alerts using the **Alerts List**, **Alert Details Drawer**, or **Alert Details Page**.
* **Notifications**: Choose notification channels to be notified in case of an alert for this threshold (see **Threshold Notifications**).
* **Mitigations**: Choose mitigations to be triggered in case of an alert for this threshold (see **Threshold Mitigations**).

### Threshold Notifications

Notifications can be sent to designated parties at various destinations when an alert is triggered or changes state. The **Notification Channels** field specifies the notifications sent when an alarm is triggered (see **Notification Channels****)**.

The notifications controls include:

* **Notification Channels**: Shows lozenges for the selected notification channels:

  + Add a Channel: Click in the field to filter and select a channel. Repeat as needed.
  + Remove a Channel: Click the **X** in the lozenge.
* **Add New Channel**: Opens the Add Notification Channel dialog (**Add or Edit Notification**) to configure and automatically add a new channel (see **Notification Settings**).
* **Test Notification Channels**: Sends a test notification to the selected channels.

  > **Note**: Test Notification Channels is visible only when channels are selected in the Notifications Channels field.

### Threshold Mitigations

> **Note:** The Mitigations section is only visible when the policy’s dimensions include source or destination IP/CIDR (see **Data Funneling**).

Set one or more mitigations (see **About Mitigation**) that trigger in response to an alert on this threshold.

#### Add Mitigation to Threshold

In its initial state, the Mitigations section includes:

* **Mitigation selector**: Select a mitigation from the dropdown list.
* **Add Mitigation**: Click to assign the currently selected mitigation to this threshold.

Once assigned, a configuration tile for the mitigation appears, with the following **Mitigation Settings**.

> **Note:** The mitigations assigned to a policy (automated mitigations) will escalate and de-escalate automatically as changing conditions match different thresholds in that policy.

#### Mitigation Settings

The Mitigation tile has the following controls:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(311).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A09Z&se=2026-05-12T09%3A58%3A09Z&sr=c&sp=r&sig=f4yfZhAv7wbfmN8qlesVw7mWFygiRIyzKqZ89dgYPkI%3D)**Remove**: Click the **X** at upper right to remove the mitigation from the threshold.
* **Apply Mitigation**: Specify in the dropdown when the mitigation is to be applied:

  + Immediately, when alert starts: Initiate mitigation when the threshold activates an alert.
  + After user confirmation: Initiate mitigation only after a user clicks “Approve and start the mitigation” on the Alerting page.
  + After user confirmation or timeout expires: Wait for a user to acknowledge or cancel the mitigation from the Alerting page. If the specified time period expires without user action, initiate the mitigation automatically.
  + Application Timer (only when Apply Mitigation is set to “After user confirmation or timeout expires”): The duration (in minutes) of the timer.
* **Clear Mitigation**: Choose from the dropdown when the mitigation will stop:

  + Immediately, when alert ends: Stop mitigation when the alert exits alarm state.
  + After user confirmation: Continue mitigation until canceled by a user.
  + After user confirmation or timeout expires: Continue mitigation until manually cancelled by a user or the specified time expires.
  + Application Timer (shown only when Clear Mitigation is set to “After user confirmation or timeout expires”): The duration (in minutes) of the timer.

#### Multiple Mitigations

You can add one or more additional mitigations to a threshold:

1. Use **Add Mitigation to Threshold** to add another mitigation.
2. Use **Mitigation Settings** to specify when the new mitigation will be applied and cleared.

This allows you to simultaneously trigger all mitigation methods/platforms (e.g., appliances at multiple sites) for a given set of conditions, scaling more effectively than cloning policies for each appliance.  
  
Support for multiple mitigations per threshold also enables a response for a given alarm to include a mix of mitigation types (e.g., RTBH, A10, and Radware). For example, in a multi-location DDoS response, you could:

1. De-preference or stop-announcing a BGP route on Location #1 by injecting a route whose community has been predefined as a flag for these actions.
2. Announce a broader routing table entry, less specific than /24, for Location #2 to force acceptance by Internet peers .
3. Trigger a third-party mitigation method (e.g., A10 or Radware) on Location #2 to announce more specific prefixes for internal re-direction to a scrubbing center.

#### Flowspec Mitigation Details

Mitigations with a Flowspec platform have a Traffic Matching card at the right of the standard **Mitigation Settings** tile. The card’s dimensions are either specific or inferred from policy values. Click **View Method Details** to see which Flowspec components (see **Flowspec Component Types**) are involved for a mitigation.

## Policy Baseline Settings

Baselining triggers alerts based on comparing current traffic values to a baseline. Derived from historical traffic patterns, the baseline represents "normal" values for a key. If the current value deviates from the norm and matches an alert policy threshold condition, the threshold triggers an alert.  
  
Baselining starts with a historical data set:

* Covers traffic funneled similarly to the current traffic (see **Policy Dataset Settings**).
* Covers a specified time range (baseline window).
* Includes time series data points representing traffic volume for a top-X key over a duration (1, 2, or 5 min) determined by the Dataset tab's **Evaluation Frequency** setting.

Two main stages build the baseline:

* **Building the baseline**: Smoothing spikes or drops in normal traffic by normalizing/averaging data points for each top-X key over the baseline window (options in **Building the Baseline**). The result is "buckets" each with a duration of one hour and with values for each of the top-X keys.
* **Using the baseline**: Choosing and normalizing/averaging key values in selected buckets into a final historical value for each key (options in **Using the Baseline**). The final baseline value is then used in comparisons in the conditions of each threshold.

A match between the current value and the final baseline value triggers an alert, depending on the threshold’s **Threshold Frequency**settings.

> **Note:** Historical baseline settings in a policy apply to all thresholds with baseline conditions.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(312).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A09Z&se=2026-05-12T09%3A58%3A09Z&sr=c&sp=r&sig=f4yfZhAv7wbfmN8qlesVw7mWFygiRIyzKqZ89dgYPkI%3D)

*For each one-hour bucket, the value of a given key (red bar) is derived from the time-slice values for that key.*

### Baseline Presets

Kentik provides three presets for configuring baselining. Choose a preset, and customize settings with the **Advanced Options** controls (see **Building the Baseline** and **Using the Baseline**):

* **Default**: Produces a general-purpose baseline for most alerting applications.
* **Precision**: Produces the most accurate baseline for highly detailed applications.
* **Express**: Rapidly produces a baseline that's less accurate but nonetheless useful for general applications.

> **Note**: Changing any setting reclassifies the baseline as Custom.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(313).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A20%3A09Z&se=2026-05-12T09%3A58%3A09Z&sr=c&sp=r&sig=f4yfZhAv7wbfmN8qlesVw7mWFygiRIyzKqZ89dgYPkI%3D)

*The Baseline tab includes presets for common baselining situations.*

**Common preset settings**:

* In Building the Baseline:

  + **Rollup aggregation**: Set to 98th percentile.
* In Using the Baseline:

  + **Bucket width**: Each key's value for a given hour is based only the key's value in that hour's bucket.
  + **Final aggregation**: Set to 95th percentile.
  + Use separate patterns for weekdays and weekends: Disabled by default.

**Preset-specific settings**:

| Preset Setting | Default | Precision | Express |
| --- | --- | --- | --- |
| Bucket depth: Number of top keys every hour | 25 | 300 | 25 |
| Baseline window start | 1 day ago | 1 day ago | 1 hour ago |
| Baseline window lookback | 4 weeks | 4 weeks | 1 week |
| Completeness: Days of data accumulated before use | 4 | 14 | 2 |
| Compare to: comparison data is derived from… | Current hour of the day, current day of the week | Current hour of the day, current day of the week | Every hour |

### Building the Baseline

The Building the Baseline pane (Advanced Options on the Baseline tab) lets you specify the historical data for an initial aggregation pass. Three settings answer the question: "How should we build the buckets for the baseline?":

* **Window**: Go back [`duration`] from [`end point`]. These settings determine the baseline window time range, e.g., "go back one week from one day ago":

  + **Duration**: The length of the time range. Options include 1, 2, 3, or 4 (default) weeks.
  + **End point**: Choose an end point for evaluating historical data from the dropdown. The window goes back from this point for the specified duration. Traffic data newer than the end point is excluded to prevent current spikes or anomalies from skewing baseline values. Options include 1 hour, 1 day (default), and 1 week ago.
* **Bucket depth**: Use the top [`##`] of keys in every hour of the window. Defaults to the **Maximum Keys per Evaluation** setting (see **Building Your Dataset**).
* **Rollup aggregation**: For each one-hour bucket, derive each key's value based on the [`minimum/maximumn/percentile`] of the time series datapoints for that hour. Options include Minimum, Maximum, or Percentile: 98th (default), 95th, 50th, 25th, or 5th.

> **Note:** The baseline window is "rolling," meaning older traffic data ages out and newer data is added continuously.

### Using the Baseline

The Using the Baseline pane (Advanced Options on the Baseline tab) lets you define how buckets in Building the Baseline are used for comparison values. It has four settings that answer the question: "Which buckets should we take key values from, and how should we derive a comparison value for each key?":

* **Completeness**: Ensure the baseline has at least [`##`] [`time unit`] of data to establish the "normal" value for each key. Time units are hours or days.

  + **Number field**: Enter a number representing time units.
  + **Time unit**: Select days (default) or hours.
* **Compare to**: For each key, take baseline values from the buckets for [`every/the current`] hour of [`every/the current`] day of the week. Choose from:

  + All buckets in the window: every hour of every day.
  + Buckets from the current hour of every day (default).
  + Buckets from the current hour of the current day.

    > **Note:** Options depend on the Window settings in Building the Baseline.
* **Bucket width**: For each key, derive the value for each "compare to" hour from [`just that hour/surrounding hours`].

  + If bucket width is "just that hour", the value is taken only from the bucket from that hour.
  + If bucket width is "surrounding hours", an additional line appears, allowing values from a number of hours on either side (see **Baseline Bucket Width**).
* **Final aggregation**: Derive the final comparison value for each key from the [`minimum/maximumn/percentile`] of it's values for the "compare to" hours. The drop-down specifies how the values in the chosen buckets become a single comparison value. Options include maximum, minimum, and percentile: 99th, 98th, 95th (default), 90th, 80th, 50th, 25th, 10th, or 5th.

  > **Note:** Policies watching for excess activity usually use Maximum, 98th, or 95th percentile aggregation. Policies watching for below-baseline activity typically use 25th percentile aggregation.

#### Baseline Bucket Width

The buckets from **Building the Baseline** cover one hour in the overall baseline window, with one value per key per bucket. The**Bucket width** setting in **Using the Baseline** allows you to replace single-hour values for each key with values derived from up to four hours on either side of the bucket's nominal hour. This additional aggregation step is specified when Bucket width is specified as "surrounding hours": Aggregate [`#`] hours before and after using [`minimum/maximumn/percentile`].  
  
The number of hours that can be aggregated is 1 to 4. The method of deriving the bucket value from multiple hours is minimum, maximum, or a percentile: 5th, 10th, 25th, 50th, 75th, or 90th.

### Weekend Baselining

The **Use separate patterns for weekdays and weekends** checkbox (Advanced Options on the Baseline tab) calculates baselines for weekends (UTC Saturday and Sunday) independently from weekdays (UTC Monday through Friday), accounting for variations in traffic patterns caused by day of week.

## NMS Threshold Policies

This section covers creating an NMS threshold policy in Kentik's alerting system.

> **Note:** Most NMS threshold policy settings are similar to other threshold policies and are covered in **Alert Policies**.

### General NMS Settings

The General tab defines the policy’s overall properties. Set the Policy Type to NMS on the General tab. The remaining settings are explained in **General Policy Settings**.

> **Note:** The General tab doesn’t include the Policy Dashboard setting when the policy type is NMS.

### Dataset NMS Settings

The Dataset tab narrows the traffic subset evaluated for thresholds. The Data Funneling section is unique to NMS policies (see **Building Your Dataset**for information on the remaining sections).

The Data Funneling pane’s settings parallel the corresponding settings on the Query sidebar of **Metrics Explorer**:

* **Measurement Pane**: Defines specifically what you'd like the query to look at. The pane’s policy settings are the same in Metrics Explorer, except you can only select one metric in a policy.
* **Filters Pane**: Narrows the scope of your query or narrows the results returned from the query. The pane's policy settings are the same in Metrics Explorer.

### Thresholds NMS Settings

A threshold is a collection of settings that define conditions for alert policy activation. It generates alerts for each matched key. Each policy includes at least one threshold but can have up to five (Critical, Severe, Major, Warning, Minor).  
  
Configure policy thresholds on the Thresholds tab. The tab is the same for NMS settings as other threshold policies (see **Policy Threshold Settings**), except for NMS the Conditions section determines what constitutes a match (see **About Matches**) for the threshold.  
  
 The Conditions section includes:

* **Static condition**: Indicates the selected metric on the Dataset tab and a lozenge indicating the value.
* **Baseline condition:** A lozenge indicating the current setting for the baseline condition.
* **Edit Conditions**: Opens the **Edit Threshold Conditions** dialog to specify static and baseline conditions.

> **Note:** The Static and Baseline conditions aren't shown in the Conditions section until specified in the dialog.

#### About NMS Conditions

For NMS threshold policies, all conditions are based on the selected metric on the Dataset tab. Edit the conditions in the Edit Threshold Conditions dialog. You can set one of each condition types:

* **Static**: Compare the current metric value to a fixed user-specified number.
* **Baseline** (only if a static condition exists): Compare the current metric value with a historical baseline for the same metric.

#### Edit Threshold Conditions

Access the Edit Threshold Conditions dialog via the **Edit Conditions** button in the **Conditions** section of a policy's Thresholds tab. It includes:

* **Conditions tiles**: Define each condition (see **About NMS Conditions**).
* **Cancel**: Click **Cancel** at lower right or **X** in upper right to close the dialog without saving changes.
* **Apply**: Apply the conditions to the policy and return to the policy settings page.

> **Note:** The settings in the Edit Threshold Conditions dialog for a given threshold are independent of other thresholds.

#### NMS Conditions Tiles

Each condition exists in a tile with the following controls:

* **Enable/disable**: A switch to turn On or Off the condition.

  > **Note:** The baseline condition can't be turned on unless the static condition is on.
* **Condition type**: A label that indicating if the condition is static or baseline.
* **Condition metric**: A lozenge indicating the metric evaluated for the condition
* **Condition value**: Controls to set the value that constitutes a match (see **Metrics Condition Match**).
* **Show/Hide Chart** (static condition only): Click to open a chart with data for the specified metric from the last three hours (default), last day, or last seven days. If shown, **Hide Chart** hides the chart.

#### Metrics Condition Match

The controls determining what constitutes a match vary based on condition type.

* **Static Condition**: Met when the current metric value exceeds the condition value specified with the following controls:

  + **Condition value**: Enter a number for comparison.
  + **Condition units**: The metric set in the Dataset tab (see **Dataset NMS Settings**).
* **Baseline Condition**: Met when the current metric value exceeds the baseline metric value to the extent specified with the following controls:

  + **Comparison value**: Enter a number for comparison.
  + **Comparison type**: Choose the comparison type (% or Units) from the dropdown:

    - **%**: Current metric value varies from baseline by the specified percentage.
    - **Units**: Current metric value varies from baseline by at least the specified number of units (e.g., bit/s, packets/s).
  + **Above/Below**: Choose whether the condition evaluates the chosen percentage or unit value above or below the baseline.
