# Alert Policies

Source: https://kb.kentik.com/docs/alert-policies-1

---

This article discusses the **Alert Policies** page of the Kentik portal. For more information about policy-based alerting, first read these articles:

* **Policy Alerts Overview**
* **Policy Templates**
* **Alerting**
* **Notifications**
* **Mitigations**
* **Policy Settings**

> **Note:** Alert policy configuration can be complex. Please reach out for help from our support team (see **Customer Care**).

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(277).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=jk%2FVgq%2B%2BDwZ28ff%2FUyqIzxKnbmNtoGpaGX%2FR0gr%2BelQ%3D)

*Manage your organization's alert policies from the Alert Policies page.*

## Alert Policies Page

The Alert Policies page is where you manage alert policies for Kentik’s alerting system (see **Alerts and Policies**).

To access the page:

1. Go to Settings » **Alerting** and click **Manage Policies**.
2. The **Policies List** shows your organization’s available alert policies. You can add, enable/disable, clone, delete, or edit policies from this list.

> **Note:** While admin-level users can see and use all of the controls listed throughout this article, member-level users can only:
>
> * See and filter the policies displayed in the **Policies** list.
> * Open a policy’s **Policy Details Drawer**.
> * Enable or disable policies.

### Alert Policies Page UI

The Alert Policies page contains the following UI elements:

* **Alert Policy Templates** (button): Opens the **Policy Templates** page to manage templates and add a policy from a template (see **Add Policy from Template**).
* **Policies statistics** (fields): Shows the number of policies created and enabled in your organization, along with the maximum allowable number for each.
* **Add Alert Policy** (button): Performs two actions:

  + Add Alert Policy: Opens the **Add Alert Policy Dialog**, where can create a new NMS or Classic Threshold policy from scratch.
  + Add Policy from Template: Click the down arrow and select **Add Policy from Template**to create a new policy from a template.
* **Filters** (funnel icon): Toggles the **Filters** pane between expanded and collapsed.
* **Search** (field): Narrows the list of policies based on matches with the entered text. Also displays any filters applied with the **Filters** pane.
* **Filters** (pane): A set of **Policy Filters** to narrow the list of alert policies.
* **Policies** (list): Lists your organization’s policies (see **Policies List**).
* **Policy Details** (drawer): Click a policy to open a summary of its settings (see **Policy Details Drawer**).

### Policies List

The Policies list is a table showing all available alert policies in your organization. It includes columns for policy information and actions.

* **Select All** (in header row): Click to select all policies in the list. Click again to deselect all policies.
* **Selected**: Check the boxes to select policies, causing the **Selected Policy Controls** to appear above the Policies list.
* **Status**: The policy’s status (Disabled or Enabled).
* **Type**: The policy type, one of NMS, Traffic, Cloud, or Protect (see **Policy Types**).
* **ID**: The unique policy ID assigned by Kentik.
* **Name**: The user-specified policy name. The user-specified description, if entered, is shown below the name.
* **Data Sources**: The network entities covered by the policy, as specified on the **Dataset** tab in **Policy Settings**.
* **Metrics**: The units (e.g., bits/s, packets/s) by which the alert measures incoming flow data (see **Data Funneling**). The primary metric is listed first, followed by secondary metrics (if any).
* **Dimensions**: The policy dimensions that define a key for traffic subdivision (see **About Keys**and **Dimensions Reference**).
* **Used by Tenant**: The **My Kentik Portal** tenants using the policy (“None” if not used).
* **Actions** (menu): Click the vertical dots icon to open the **Policy Action Menu**.

#### Selected Policy Controls

When one or more policies are selected, these controls appear above the **Policies** list:

* **Add/Edit Labels**: Click to open the **Labels**settings page, where you can create, edit, or remove labels for the selected policies.
* **Clear Labels**: Clears all applied labels from the selected policies.
* **Apply Labels**: Select one or more labels to apply them to all selected policies. Click outside to close the dropdown.
* **Enable/Disable**: Enable or disable all the selected policies.
* **Delete**: Opens a **Delete Policies** dialog to confirm or cancel removal of the selected policies.
* **Add Notification Channels**: Opens the **Add Notification Channels** dialog, where you assign one or more notification channels to the selected policies.
* **Policies selected**: Displays the number of selected policies.

![With policies selected, controls for applying labels and adding notification channels appear at the top.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(278).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=jk%2FVgq%2B%2BDwZ28ff%2FUyqIzxKnbmNtoGpaGX%2FR0gr%2BelQ%3D)

With policies selected, controls for applying labels and adding notification channels appear at the top.

#### Add Notification Channels

The **Add Notification Channels** dialog lets you assign one or more notification channels to the selected policies. It has the following:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AP-Add_notif_dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=jk%2FVgq%2B%2BDwZ28ff%2FUyqIzxKnbmNtoGpaGX%2FR0gr%2BelQ%3D)**Close** (button): Click the **X** in the upper right to close the dialog without saving changes.
* **Notification Channels** (field): Shows lozenges for any assigned notification channels. The lozenge to the left of each channel name indicates the channel's type.

  + Remove: Click the **X** in a channel’s lozenge to remove it.
  + Add: Click in the field to select a channel from the filterable channel list.

    > **Notes**:
    >
    > - Channels already assigned to this policy are not included in the list.
    > - Disabled channels are included but can’t be selected.
* **Add New Channel**: Opens the **Add Notification Channel** dialog (see **Notification Settings**) to create a new channel and automatically add it to the selected policies.
* **Test Notification Channels**: Sends a test notification to all assigned channels. Present only when **Notification****Channels** is populated.
* **Cancel**: Exits the dialog without changing channel assignments.
* **Continue**: Saves current channel assignments and exits.

#### Policy Action Menu

Access the Action menu for each policy via the vertical dots in either the **Policies List**,at the right of each policy's row, or the the **Policy Details Drawer**, at the right of the policy title.

The menu includes these options:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(280).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=jk%2FVgq%2B%2BDwZ28ff%2FUyqIzxKnbmNtoGpaGX%2FR0gr%2BelQ%3D)**Edit Policy**: Edit **Policy Settings** on the Edit Policy page.
* **Enable/Disable Policy**: Enables or disables a policy.
* **Clone Policy**: Starts a new policy pre-populated with settings from an existing policy (see **Clone a Policy**).
* **Delete**: Confirm deletion of a policy.
* **Debug**: Opens the **Policy Debug Dialog** with the top keys for both current traffic and baseline data.

### Policy Filters

Filter policies in the **Policies** list using the **Filters** pane on the left. By default, all policies are shown. Click the checkboxes to narrow the list by policy **Type** and/or **Status.**

The Filters pane includes:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(281).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=jk%2FVgq%2B%2BDwZ28ff%2FUyqIzxKnbmNtoGpaGX%2FR0gr%2BelQ%3D)**Close**: Click the **<** in the upper right to close the pane. Click the funnel icon to show it again.
* **Reset to default**: Clears all current filters (appears only when filters are specified).
* **Type**: Check the policy type boxes (NMS, Traffic, Cloud, or Protect) to restrict the list (see **Policy Types**).
* **Status**: Check policy status boxes (Enabled and/or Disabled) to restrict the list.
* **Policy ID**: Enter a complete policy ID to filter the list to only that policy (no partial matching).

### Policy Details Drawer

The Policy Details drawer is a read-only display of policy settings, including details not displayed in the **Policies** list. Click on a policy row to open its details drawer. Click the same row again to close the drawer.  
  
 The drawer contains the following elements:

* ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(282).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=jk%2FVgq%2B%2BDwZ28ff%2FUyqIzxKnbmNtoGpaGX%2FR0gr%2BelQ%3D)**ID number**: The unique ID assigned when the policy was created.
* **Name**: The user-specified policy name.
* **Actions**: A vertical dots icon that opens the Action menu (see **Policy Action Menu**).
* **Status**: A lozenge indicating the policy’s status (enabled: green, disabled: gray).
* **Description**: The user-specified policy description.
* **Policy Dashboard** (not present for NMS policies): The dashboard name associated with the policy, if any (see **General Policy Settings**). A lozenge labeled Preset indicates the dashboard is a Kentik-provided preset.
* **Dataset** (pane)**:** Summarizes the policy’s **Dataset** settings (see **Policy Dataset Settings**).
* **Thresholds** (pane): Summarizes the policy’s **Thresholds** settings (see **Policy Threshold Settings**). The number in brackets indicates the number of thresholds.
* **Baseline** (pane): Summarizes the policy’s **Baseline** settings (see **Policy Baseline Settings**).

### Policy Debug Dialog

The **Debug** dialog provides context to understand why a threshold in an alert policy triggered an alarm (see **About Alert Thresholds**). The included table presents information about a policy from two perspectives:

* **Current Traffic**: Shows information about the top-X keys (see **About Keys**) for the policy in current traffic. Current traffic is the set of flow records and associated traffic data in the most recently completed aggregate.

  + The duration of each aggregate is determined by the policy’s **Evaluation Frequency** setting.
  + The number of keys is defined by the **Maximum Number of Keys** (the actual number of keys may be less depending on the **Minimum Traffic Threshold** setting).
* **Baseline Data**: The table shows information related to the top-X keys in the baseline data of the same policy (see **About Historical Baselines**).

> **Notes**:
>
> * The Debug dialog is only available for enabled policies.
> * See **Policy Settings Pages** for more about threshold policy settings.

![The debug dialog provides details about the top keys in either current or baseline traffic.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AP-Debug_current.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=jk%2FVgq%2B%2BDwZ28ff%2FUyqIzxKnbmNtoGpaGX%2FR0gr%2BelQ%3D)

The debug dialog provides details about the top keys in either current or baseline traffic.

#### Policy Debug Dialog UI

The **Debug** dialog has the following UI elements:

* **Close**: Click the **X** in the upper right to close the dialog.
* **Top Keys**: Choose from the drop-down which top-X keys to show in the **Debug** list, either **Current Traffic** or **Baseline Data**.
* **Debug list**: A table showing info about top-X keys (current or baseline) for the selected policy (see **Debug Current Columns** and **Debug Baseline Columns**).

> **Note:** If insufficient traffic is available for debugging, a message replaces the Debug list:
>
> * Current traffic: “No matches found” means no traffic or no traffic for the policy to evaluate.
> * Baseline data: “No baseline data available” means no traffic or the baseline is still building.

#### Debug Current Columns

The **Current Traffic** view of the **Debug** list includes the following columns:

* **Position** (#): The ordinal position of the key in the top-X current keys for this policy.
* **Keys**: The policy’s key as defined in the **Dimensions** setting (see **Data Funneling**).
* **Entries**: The total number of flows with this key in the current traffic data aggregate.
* **Primary Metric**: The policy’s Primary Metric value, which determines the position of the key in the top-X (see **Data Funneling**).
* **Metric 2**: The policy’s first Secondary Metric value, if any (see **Data Funneling**).
* **Metric 3**: The policy’s second Secondary Metric value, if any (see **Data Funneling**).
* **First seen**: The start time of the current traffic data aggregate.

  > **Note**: The start time is in either UTC or local time depending on the **Time Zone** selected in your User Profile (see **User-specific Defaults**).

#### Debug Baseline Columns

The **Baseline Data** view of the **Debug** list includes the following columns:

* **Position** (**#**): The ordinal position of the key in the top-X baseline keys for this policy.
* **Keys**: The policy’s key as defined in the policy’s **Dimensions** setting (see **Data Funneling**).
* **Entries**: The total number of flows with this key in the baseline data.
* **Value**: A percentile value, e.g. P95, for the policy’s Primary Metric. Determines the key’s position in the top-X (see **Building the Baseline**).
* **Min Value**: The lowest one-hour rollup aggregation value for the policy’s Primary Metric over the “baseline window” (see **Building the Baseline**).
* **P50 Value**: The 50th percentile value of the policy’s Primary Metric over the “baseline window.”
* **Max Value**: The highest one-hour rollup aggregation value for the Primary Metric over the “baseline window.”
* **Chosen Time**: The timestamp of the first flow record matching the key.

## Adding a Policy

Follow these steps to add an alert policy in Kentik.

> **Note:** To add a policy, your assigned RBAC roles must have permissions equivalent to those of an Administrator or Super Administrator (see **Kentik-managed Roles**).

### Policy Creation Methods

Add an alert policy in the Kentik portal using any of the following methods:

* **From Scratch**:

  + On the **Alert Policies Page**, click **Add Alert Policy**.
  + Select either an NMS (see **NMS Policy Settings**) or Classic Threshold (see **Threshold Policy Settings**) policy and proceed to the settings page to create the new policy.
* **From Data Explorer (Classic Threshold only)**:

  + In the Data Explorer subnav, select **Create Alert Policy** from the **Actions** drop-down (see **Add a Query-based Policy**).
  + A policy settings page, pre-populated with your query settings and with policy type **Traffic** pre-selected, opens to create the new policy (see **Threshold Policy Settings**)
* **From Policy Template**:

  + On the **Alert Policies Page**,use the subnav controls to start a new policy from a pre-configured template (see **Use a Policy Template**).
* **Clone Existing Policy**:

  + In the **Policies List**,click the vertical dots icon for a policy and select **Clone Policy** (see **Clone a Policy**).
  + A policy settings page, pre-populated with settings from the cloned policy, opens to create the new policy.

### Add Alert Policy Dialog

The **Add Alert Policy** dialog offers two options: creating an **NMS** or **Classic Threshold** policy. Once you select your preferred policy, you’ll be redirected to the corresponding settings page.

The dialog, which opens from the **Add Alert Policy** button on the **Alert Policies Page**, includes:

* ![Choose an NMS or Classic Threshold policy type and finalize settings of a new alert policy.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AP-Add_policy_dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=jk%2FVgq%2B%2BDwZ28ff%2FUyqIzxKnbmNtoGpaGX%2FR0gr%2BelQ%3D)**NMS**: Click to select an NMS policy, which alerts you when devices, interfaces, or BGP neighbors are in an unhealthy state (see **NMS Policy Settings**).
* **Classic Threshold** (default): Click to select a Classic Threshold policy, which alerts you when selected metrics pass specified thresholds (see **Threshold Policy Settings**).
* **Cancel**: Click **Cancel** at lower right or **X** in the upper right to exit the dialog without saving.
* **Continue**: Click to proceed to a **Policy Settings** page to finalize details of the new policy.

### Add a Query-based Policy

A query-based alert policy is time-limited and built from contextual Data Explorer queries. You create it directly in Data Explorer using your current data view criteria (dimensions, metrics, filters).

> **Note:** Query-based policies are Traffic type (see **Policy Types**).

To create a query-based alert policy:

1. ![Menu options including 'Create Alert Policy' highlighted in the actions dropdown.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AP-add-query-based-policy.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=jk%2FVgq%2B%2BDwZ28ff%2FUyqIzxKnbmNtoGpaGX%2FR0gr%2BelQ%3D)In **Data Explorer**, create a query using the **Query** sidebar settings (data sources, dimensions, metrics, filters, etc.).
2. In the subnav, choose **Create Alert Policy** from the **Actions** drop-down to create a new policy on the Add Query-Based Policy page. The policy’s settings are populated automatically from the Data Explorer query.
3. On the **General** tab, provide a name for the policy.
4. On the **Thresholds** tab, click **Edit Conditions** and add at least one condition (see **Threshold Conditions**).
5. Check the **Policy Summary Pane** for missing fields or errors.
6. Click **Save** to create the new policy and return to Data Explorer.

> **Note**: Editing a query-based policy after it appears in the **Policies** list removes its expiration.

### Clone a Policy

Clone an alert policy to duplicate it for modification without altering the original:

1. Choose Settings » **Alert Policies** from the portal's main nav menu.
2. From the Actions menu (vertical dots) for the policy, choose **Clone Policy**. This takes you to the Clone Policy page, where the settings of a new policy will be populated with the values of the cloned policy.
3. Change the name/description of the new policy as necessary.
4. Finalize other settings and click **Save** to create the new policy.

### Use a Policy Template

Kentik provides preconfigured policies as templates to cover many common alerting needs.

There are two ways to add a policy from a template using the **Alert Polices** page:![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(288).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=jk%2FVgq%2B%2BDwZ28ff%2FUyqIzxKnbmNtoGpaGX%2FR0gr%2BelQ%3D)

* **Alert Policies page**: Click the drop-down on the **Add Alert Policy** button and select **Add Policy from Template**. Choose a template and click **Continue**.
* **Policy Templates page**: Click the Policy Templates button and click the Clone icon on a template.

Both methods open an Add Alert Policy page with default settings from the template. Tailor the settings to your needs and save the new policy (see **Policy Settings**).

> **Note:** Templates should not be used as-is, but should be customized to your network and traffic situation.

#### Add Policy from Template

The **Add Alert Policy from Template** dialog includes the following:

* ![Adding an alert policy for high BPS between internal VMs in AWS.](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/AP-Add_from_template_dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A19%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=jk%2FVgq%2B%2BDwZ28ff%2FUyqIzxKnbmNtoGpaGX%2FR0gr%2BelQ%3D)**Cancel** (buttons): Click the **X** in the upper right or **Cancel** at lower left to exit the dialog without saving.
* **Drop-down menu**: Click to choose an existing template from the list.
* **Go to Policy Templates**: Opens the Policy Templates page (see **Policy Templates**).
* **Template description**: A description of what the template is designed to detect.
* **Continue** (button): Opens the Add Alert Policy page, pre-populated with the chosen template’s settings.

  > **Note:** The **Continue** button is inactive until you select a policy template.

## Policy Settings

The pages and dialogs used to specify policy settings are discussed here.

> **Note:** For specifics about the policy settings UI, see **NMS Policy Settings**and **Threshold Policy Settings**.

### Policy Settings Pages

Alert policy settings can be accessed on the following pages:

| Page | Use to… | Locations |
| --- | --- | --- |
| **Add Alert Policy** | Configure a new policy from scratch or via template | Alert Policies page, **Add Alert Policy** button and **Add Alert Policy from Template** dialog. |
| **Edit Policy** | Edit an existing policy | Alert Policies page, **Edit Policy** in the **Policy Action Menu**. |
| **Clone Policy** | Duplicate an existing policy | Alert Policies page, **Clone Policy** in the **Policy Action Menu**. |
| **Add Query-Based Policy** | Create a policy from Data Explorerquery settings | **Data Explorer** page, **Create Alert Policy** in the **Actions** menu in the subnav. |

### Policy Settings Page UI

The policy settings pages share the same layout and the following common UI elements:

* **Help**: Click the **?** icon button in the subnav to open the Alert Policies KB article in a new tab.
* **Cancel** (button): Exits with saving changes and returns you to the Alert Policies page.
* **Create** or **Save** (button): Saves policy settings and returns you to the Alert Policies page. Active only when all settings are complete and error-free (see **Policy Summary Pane**).
* **Settings tabs** (not present in NMS policies): Tabs for making policy settings (see **Threshold Settings Tabs**).
* **Summary pane** (not present in NMS policies): Expand/collapse a summary of each tab’s settings. Also indicates completion status (see **Policy Summary Pane**).

> **Note:** The **Cancel** and **Create** or **Save** buttons are located at the upper right except for NMS policies, where they are at the bottom of the page.

## Manage Alert Policies

Follow these steps to manage alert policies in the Kentik portal.

> **Note:** See also **Add a Query-based Policy**, **Clone a Policy**, and **Use a Policy Template**.

### Add an NMS Policy

To create an NMS alert policy using the Alert Policies page:

1. Choose Settings » **Alert Policies** from the portal's main nav menu.
2. Click **Add Alert Policy** at the upper right.
3. Select **NMS** and click **Continue.**
4. Specify the settings (see **NMS Settings Page**).
5. Click **Save** to save the policy and return to the Alert Policies page.

### Add a Classic Threshold Policy

To create a Classic Threshold alert policy using the Alert Policies page:

1. Choose Settings » **Alert Policies** from the portal's main nav menu.
2. Click **Add Alert Policy** at the upper right.
3. Select **Classic Threshold** and click **Continue.**
4. On the Add Alert Policy page, specify the settings on each tab:

   1. General: See **General Policy Settings**.
   2. Dataset: See **Policy Dataset Settings**.
   3. Thresholds: See **Policy Threshold Settings**.
   4. Baseline: See **Policy Baseline Settings**.
5. Click **Save** to save the policy and return to the Alert Policies page.

### Add an NMS Threshold Policy

To create an NMS threshold alert policy using the Alert policies page:

1. Choose Settings » **Alert Policies** from the portal's main nav menu.
2. Click **Add Alert Policy** at the upper right.
3. Select **Classic Threshold** and click **Continue.**
4. On the Add Alert Policy page, specify the settings on each tab:

   1. General:

      1. Select **NMS** from the **Policy Type** selector.
      2. Adjust other settings (see **General Policy Settings**).
   2. Dataset:

      1. Specify settings covered in **Dataset NMS Settings**.
      2. Adjust other settings (see **Policy Dataset Settings**).
   3. Thresholds:

      1. Specify settings covered in **Thresholds NMS Settings**.
      2. Adjust other settings (see **Policy Threshold Settings**).
   4. Baseline: Specify the settings covered in **Policy Baseline Settings**.
5. Click **Save** to save the policy and return to the Alert Policies page.

### Edit a Policy

To edit an alert policy using the Alert Polices page:

1. Choose Settings » **Alert Policies** from the portal's main nav menu.
2. Click the vertical dots icon for a policy and select **Edit Policy**.
3. Update policy settings.
4. Click **Save** to save your changes and return to the Alert Policies page.

> **Note**: For details on available settings, see **NMS Policy Settings**and **Threshold Policy Settings**.

### Disable or Enable a Policy

To enable or disable an alert policy using the Alert Policies page:

1. Choose Settings » **Alert Policies** from the portal's main nav menu.
2. Click the vertical dots icon for a policy and select either:

   1. Disable Policy: The policy will no longer monitor its dataset, generate alerts, or trigger mitigations.
   2. Enable Policy: The policy will resume its normal functions.

### Remove a Policy

To remove an alert policy using the Alert Policies page:

1. Choose Settings » **Alert Policies** from the portal's main nav menu.
2. Click the vertical dots icon for a policy and select **Remove.**
3. Click **Remove** to confirm deletion of the policy.

### Debug a Policy

To debug an enabled alert policy using the Alert Policies page:

1. Choose Settings » **Alert Policies** from the portal's main nav menu.
2. Click the vertical dots icon for a policy and select **Debug**.

   > **Note**: You can also select **Debug** from the **Action** menu in a policy’s **Policy Details Drawer**.
3. Select **Current Traffic** or **Baseline Data** from the dropdown.

> **Note:** When insufficient traffic is available for debugging, a message will display in place of the **Debug** list (see **Policy Debug Dialog UI**).
