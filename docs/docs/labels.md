# Labels

Source: https://kb.kentik.com/docs/labels

---

This article discusses the management of **Labels** in the Kentik portal.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(610).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=%2B2IGqMqAZxmdgEwC6cPldNs7xOz0vwN2cG0BBUY3YwE%3D)

*Use the Labels page to manage your organization's labels, which represent collections of devices, agents, and tests.*

## About Labels

Labels are **text-based properties** that enable you to create collections of items and treat them as groups. They can be applied to various Kentik resources including:

* **Devices**
* **Synthetic Tests**
* `ksynth` **Agents**
* **Dashboards** and **Saved Views**
* **Alert Policies**
* **Credentials** and **RBAC Roles**

### **Label Use Cases**

Here are some practical examples of how you can use labels in the Kentik portal:

* **Geographic Organization**: Label devices by location (e.g., "US-East", "Europe", "APAC")
* **Functional Grouping**: Group by role (e.g., "Core-Routers", "Edge-Devices", "DMZ")
* **Environmental Separation**: Distinguish environments (e.g., "Production", "Staging", "Development")
* **Team-Based Access**: Control access based on team responsibilities
* **Service Classification**: Group devices by the services they support

#### More on Device Labels

Labels are particularly useful for dynamically controlling which network devices are included in a query used by a Kentik resource (e.g., dashboard). The group of devices included at each query runtime can be managed by simply updating the list of devices with a particular label, instead of having to revise the resource itself.

> **Notes:**
>
> * A query treats devices, agents, and tests with a given label at runtime as having that label for the entire query time range.
> * Labels can restrict access to parts of the portal (see **Using Labels with RBAC**).

### Label Appearance

Throughout the Kentik portal, item labels appear where details are displayed, filtered, or set (e.g., lists, Details panes):

* Labels are gray rounded rectangle with a user-selected color on the left.
* In a label settings fields, an **X** appears to the right of the label text. Click the **X** to remove a label.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(612).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=%2B2IGqMqAZxmdgEwC6cPldNs7xOz0vwN2cG0BBUY3YwE%3D)

## Labels Page

Labels are managed from the Labels page, accessible from Settings » **Labels** in the main nav menu.

> **Note:** Only Administrator-level users can add, modify, delete, or remove labels. Member-level users can only view existing labels.

### Labels Page UI

The Labels page includes the following UI elements:

* **Add Label controls**: Buttons and fields to add labels (see **Add Label Controls**).
* **Show/Hide Filters** (funnel icon): Toggles the **Filters** pane.
* **Search**: Displays current usage filters and allows searching labels.
* **Filters pane**: Filters the **Labels** list (see **Label Filters Pane**).
* **Labels list**: Lists your organization’s defined labels (see **Labels List**).

### Labels List

The **Labels** list displays all your organization’s existing labels. Sort the list by clicking on a column heading (ascending or descending). Each row represents a label with information and options as described below.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(613).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=%2B2IGqMqAZxmdgEwC6cPldNs7xOz0vwN2cG0BBUY3YwE%3D)

#### Label Information Columns

The following columns are displayed:

* **Label**: The color and name of the label.
* **Usage**: The numbers of devices, agents, tests, policies, saved views, or dashboards to which a label is assigned.

#### Label Admin Actions

The following actions are available only to Administrator-level users:

* **Edit Label** (pencil icon): Opens an **Edit Label Dialog** to change the label’s color or name.
* **Remove** (trash icon): Opens a dialog to confirm removal of the label.

### Add Label Controls

The **Add Label** controls at the top right enable Administrator-level users to add a new label to their organization’s collection of labels:

* **Color** (swatch): Opens a popup to select a label color (default: gray).
* **Name**: Enter a name for the new label.
* **Add Label** (active when the Name field is filled): Adds a new label with the specified name and color to the Labels list.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(614).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=%2B2IGqMqAZxmdgEwC6cPldNs7xOz0vwN2cG0BBUY3YwE%3D)

### Label Filters Pane

The left pane contains controls to filter the **Labels** list:

* **Reset To Default** (present only when one or more filters are specified): Clears all current filters.
* **Usage:** Checkboxes for the item types to which a label may be applied (agent, device, alert policy, synthetic test, dashboard, or saved view). Only selected types appear in the list.

The filter rules are:

* No selections: Includes all of your organization's existing labels.
* Selected **Usage** checkboxes: Includes labels applied to any items of the selected types. Unapplied labels are excluded.
* Search field and selected checkboxes: Includes only labels matching both criteria.

## Edit Label Dialog

The **Edit Label** dialog, accessible from the **Edit** button on each row of the **Labels List**, allows you to modify a label’s properties:

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/Lbl-Edit_dialog.png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=%2B2IGqMqAZxmdgEwC6cPldNs7xOz0vwN2cG0BBUY3YwE%3D)

* **Color selector**: A swatch displaying the current label color. Click to open a popup to change it.
* **Name**: An editable field displaying the label name.
* **Cancel**: An **X** at the top right or a button at the bottom to exit the dialog without saving changes.
* **Save**: Saves the current settings and exits the dialog.

## Manage Labels

Here, we’ll discuss the steps involved in managing labels through the Labels page.

> **Note**: For steps on applying labels to various Kentik resource types, see **Apply Labels to Resources**.

### Add a Label

To add a new label:

1. Open the Labels page (Settings » **Labels**).
2. In the **Add** **Labels** section, specify:

   1. Label color: Click the swatch to set a label color.
   2. Label name: Enter a label name.
3. Click **Add Label** to add it to the Labels list.

### Edit a Label

To edit the name or color of a label:

1. Open the Labels page (Settings » **Labels**).
2. In the **Labels List**, find the label that you'd like to modify.
3. Click **Edit** on the label’s row, then in the **Edit Label Dialog** modify the name field and/or color swatch.
4. Click **Save** to save changes.

### Remove a Label

To remove a label:

1. Open the Labels page (Settings » **Labels**).
2. In the **Labels List**, find the label that you'd like to remove.
3. Click **Remove**, then in the dialog:

   1. Click **Remove Label** to remove the label.
   2. Click **Cancel** button to exit without removing the label.

## Apply Labels to Resources

This topic covers all available locations and methods for applying and removing existing labels to Kentik resources.

> **Note**: For the steps on creating/editing/deleting labels, see **Manage Labels**.

The locations within the Kentik portal that you use to apply a label depend on type of resource:

| Resource Type | Portal Locations to Apply Labels |
| --- | --- |
| **Devices** | * **Label Controls** above the **Device List**. * Settings in the **Device Settings Dialog**. |
| **Synthetic Tests** | * **Label Controls** above the **Tests List**. * Settings on the **Test Information** tab on a **Test Settings Page**. |
| **Alert Policies** | * **Selected Policy Controls** above the **Policies List** (all policy types). * Settings on **Policy Settings Pages** (excluding NMS policy type) |
| **Credentials** | Settings in the **Credential Settings Dialogs**. |
| **Dashboards** | **Selected Item Controls**above the **Library List**. |
| **Saved Views** | * **Selected Item Controls** above the **Library List**. * Settings in the **Saved View Dialogs**. |
| `ksynth` **Agents** | * Label controls above the **Agents List**. * Settings in a **Configure Agent Dialog**. |
| **RBAC Roles** | Settings in the **Manage RBAC Role Page**. |

### Label a Device

Apply labels to devices in the following locations.

#### Label Devices on Devices Page

To apply an existing label to devices with the Devices page:

1. Open the Devices page (Settings » **Networking Devices**).
2. Select devices by checking the boxes. The **Apply Labels** field shows existing labels.
3. Click **Apply Labels** for a filterable list of labels (see **Label Controls**).
4. Highlighted labels are already applied. Click twice on a highlighted label to apply it to the new set of devices.
5. Click once to apply a non-highlighted label.
6. Click outside to close the dropdown. Labels will appear under the device name in the **Device** list.

#### Label Devices in Device Dialog

To apply an existing label to one device using a Device dialog:

1. Open the Devices page (Settings » **Networking Devices**), then do one of the following:

   1. Click **Add Device** to add a new device (see **Add a Device**).
   2. Click the **Edit** button in the Device list for an existing device.
2. In the **General** tab of the **Device** dialog, the applied labels appear in the **Labels** field.
3. Click **Labels** to filter the labels. Labels that are already applied to the device will be highlighted.
4. Select the desired labels.
5. For new devices, complete the setup and click **Add Device** to save and exit.
6. For existing devices, click **Save** to save and exit.

### Remove Device Labels

You can remove a device label from both the Devices page and a **Device** dialog.

#### Remove Labels on Devices Page

To remove a label from one or more devices:

1. Open the Devices page (Settings » **Networking Devices**).
2. Select devices by checking the boxes.
3. Click the **X** at the right of any label in the **Apply Labels** field to remove it from all checked tests.
4. Alternatively, click **Clear Labels** to eliminate all labels from those tests.

#### Remove Labels in Device Dialog

To remove a label from a single device:

1. Open the Devices page (Settings » **Networking Devices**).
2. Click **Edit** next to the device you want to modify.
3. In the **Device** dialog, General tab, click the **X** at the right of a lozenge in the **Label(s)** field to remove the label.
4. Click **Save** to save your changes and close the dialog.

### Label a Synthetics Test

A label can be applied to a synthetic test in two places:

#### Label Test in the Tests List

To apply an existing label to one or more tests:

1. Open the **Test Control Center** page from the main nav menu.
2. In the **Tests** list, check the boxes for tests to label. The **Apply Labels** field shows existing labels.
3. Click **Apply Labels** for a filterable list of labels (see **Label Controls**).
4. Highlighted labels are already applied. Click twice on a highlighted label to apply it to the new set of devices.
5. Click once to apply a non-highlighted label.
6. Click outside to close the dropdown. Labels will appear under the device name in the **Device** list.

#### Label a Test in Test Settings

To apply a label to a synthetic test on its settings page:

1. Open **Test Control Center** from the main menu and do one of the following:

   1. New test: Click **Add Test** to add a new test (see **Add a Test**).
   2. Existing test: Click the **Edit** button for the test in the **Tests** list.
2. On the **Test Settings Page**, go to the Test Information tab. For an existing test, the **Labels** field will show labels currently applied to the test.
3. Click **Labels** to open a filterable list of labels. Labels already applied to the test will be highlighted.
4. Select the labels to apply to the test. Each selected label will appear in the **Labels** field. Click outside the dropdown to close the list.

   > **Note:** Remove labels by clicking the **X** next to them.
5. Save the labels to the test and return to the Test Control Center:

   1. New test: Complete other settings then click **Create Test**.
   2. Existing test: Click **Save**.

> **Note**: Create new labels using **Add Label** (must be manually applied afterward).

### Remove Synthetics Test Labels

Labels can be removed from synthetic tests in two locations:

#### Remove Labels on Tests List

To remove a label from one or more synthetic tests on the **Tests** list:

1. Open **Test Control Center** from the main nav menu.
2. Select tests by checking the boxes next to them.
3. In the **Apply Labels** field, either:

   1. Click the **X** next to specific labels to remove them.
   2. Click **Clear Labels** to remove all labels from the selected tests.

#### Remove Labels in Test Settings

To remove a label from a synthetic test on its Edit Test page:

1. Open **Test Control Center** from the main nav menu.
2. Click **Edit** for the desired test.
3. Go to the **Test** **Information** tab. The **Labels** field will show any labels currently applied to the test.
4. Click the **X** next to any labels you want to remove.
5. Click **Save** to apply changes.

### Label an Alert Policy

Labels can be applied to alert policies in two locations:

#### Label Policy on Policies Page

To apply an existing label to one or more policies on the Policies list (**all policy types**):

1. Go to Settings » **Alert Policies**.
2. Select policies using checkboxes (**Selected Policy Controls** appear). The **Apply Labels** field will show labels already applied to the selected polices.
3. Click **Apply Labels** for a filterable list of labels (see **Label Controls**).
4. Highlighted labels are already applied. Click twice on a highlighted label to apply it to the new set of policies.
5. Click once to apply a non-highlighted label.
6. Click outside to close the dropdown. Labels will appear in the **Name** column of the **Policies** list.

   > **Note**: Labels can only be applied to NMS policies via the Policies page.

#### Label Policy in Policy Settings

To apply a label to an alert policy on its settings page (**Classic Threshold policy type only**):

1. Go to Settings » **Alert Policies** and do one of the following:

   1. New policy: Click **Add Policy** to add a new policy (see **Adding a Policy**)**.**
   2. Existing policy: Click **Edit Policy** from the **Policy Action Menu** (vertical ellipsis) for the policy.
2. Go to the **General** tab. For an existing policy, the **Labels** field will show labels currently applied to the policy.
3. Click the **Labels** field to drop down a filterable list of your organization's labels. In the list, the labels (if any) that are already applied to the selected policy will be highlighted.
4. Select the labels to apply to the policy. Each selected label will appear in the **Labels** field. Click outside the dropdown to close the list.

   > **Note:** Remove labels by clicking the **X** next to them.
5. Save the labels to the policy and return to the Policies page:

   1. New policy: Complete other settings, then click **Save**.
   2. Existing policy: Click **Save**.

> **Note**: Create new labels using **Add Label** (must be manually applied afterward).

### Remove Policy Labels

You can remove a policy’s label in two locations:

#### Remove Labels on Policies Page

To remove a label from one or more policies on the Policies page:

1. Open the Alert Policies page (Settings » **Alert Policies**).
2. Select policies using checkboxes (**Selected Policy Controls** appear). The **Apply Labels** field will show labels already applied to the selected polices.
3. In the **Apply Labels** field, either:

   1. Click the **X** next to specific labels to remove them.
   2. Click **Clear Labels** to remove all labels from the selected policies.

#### Remove Labels in Policy Settings

To remove a label from a policy on its Edit Policy page:

1. Open the Alert Policies page (Settings » **Alert Policies**).
2. Choose **Edit Policy** from the **Policy Action Menu** (vertical ellipsis) for the policy.
3. Go to the **General** tab. The **Labels** field will show labels currently applied to the policy.
4. Click the **X** next to a label to remove it.
5. Click **Save** to apply changes and return to the Policies page.

### Label a Credential

To apply an existing label to a credential:

1. Open the Credentials Vault page from the **Organization Settings** menu.
2. Find the credential you want to label and click **Edit**.
3. In the **Credential Settings Dialog**, the **Labels** field will show labels currently applied to the credential.
4. Click **Labels** for a filterable list of available labels. Labels already applied to the selected credential will be highlighted.
5. Select the labels you want to apply and click outside to close the drop-down.

   > **Note:** Remove labels by clicking the **X** next to them.
6. Click **Save** to apply changes and close the dialog.

> **Note**: Create new labels using **Add Label** (must be manually applied afterward).

### Remove Credential Labels

To remove a label from a credential:

1. Go to Organization Settings » **Credentials Vault**.
2. Find the credential you want to remove a label from and click **Edit**.
3. In the **Edit Credential** dialog, the **Labels** field will show the labels currently applied to the credential. Click the **X** next to a label to remove it.
4. Click **Save** to apply changes and close the dialog.

### Label a Dashboard

Labels can be applied to a dashboard in two locations:

#### Label Items in Library

To apply an existing label to dashboards or saved views in the Library:

1. Open the **Library** from the main nav menu.
2. Select items using checkboxes (**Selected Item Controls**appear). The **Apply Labels** field will show labels already applied to the selected items.
3. Click **Apply Labels** for a filterable list of available labels.
4. Highlighted labels are already applied. Click twice on a highlighted label to apply it to the new set of items.
5. Click once to apply a non-highlighted label.
6. Click outside to close the dropdown. Labels will appear in the **Name** column of the **Library** list.

#### Label Dashboard in Dashboard Dialog

To apply an existing label to a dashboard using a dashboard dialog (see **Dashboard Dialogs**):

1. Open **Library** from the main nav menu.
2. Open a dashboard dialog:

   1. New dashboard: Click Add New » **Dashboard.**
   2. Existing dashboard: Choose **Edit Properties** from the **Action** menu (vertical ellipsis) for the dashboard.
3. The **Labels** field will show any labels currently applied to the dashboard.
4. Click **Labels** to open a filterable list available labels. Labels already applied to this dashboard will be highlighted.
5. Select one or more labels from the filterable drop-down list. Each selected label will appear in the **Labels** field.
6. Click outside to close the dropdown.

   > **Note:** Remove labels by clicking the **X** next to them.
7. Save the applied labels to the dashboard and return to the Library:

   1. New dashboard: Complete the other settings of the dashboard and click **Add Dashboard**.
   2. Existing dashboard: Click **Save.**

> **Note**: Create new labels using **Add Label** (must be manually applied afterward).

### Remove Dashboard Labels

Labels can be removed from a dashboard in two locations:

#### Remove Labels via Library List

To remove a label from one or more items (dashboards and/or saved views) in the **Library** list:

1. Open the **Library** from the main nav menu.
2. Click the checkbox at the left of each item from which you want to remove a label. The **Apply Labels** field will show labels applied to one or more of the selected items.
3. Click the **X** next to any label you want to remove from the currently selected items
4. Alternatively, click **Clear Labels** to remove all labels from those items.

#### Remove Labels via Library Dialogs

To remove a label from a Library item using its settings dialog:

1. Open the **Library** from the main nav menu.
2. Choose **Edit Properties** from the **Action** menu (vertical ellipsis) for the item.
3. The **Labels** field shows the labels currently applied to the item.
4. Click the **X** next to any labels you want to remove.
5. Click **Save** to apply changes and return to the Library.

### Label a Saved View

A label may be applied to a saved view in two locations:

#### Label Saved View in Dialog

To apply an existing label to a single saved view using a saved view dialog:

1. Open the **Library** from the main nav menu.
2. Open a saved view dialog:

   1. New saved view: Click Add New » **Saved View**.
   2. Existing saved view: Choose **Edit Properties** from the **Action** menu (vertical ellipsis) for the saved view.
3. The **Labels** field shows the labels currently applied to the saved view.
4. Click **Labels** for a filterable list of available labels. Labels already applied to this saved view will be highlighted.
5. Select one or more labels to apply to the saved view. Each selected label will appear in the **Labels** field.
6. Click outside to close the dropdown.

   > **Note:** Remove labels by clicking the **X** next to them.
7. Save the applied labels to the saved view and return to the Library:

   1. New saved view: Complete the other settings of the saved view and click **Add Saved View**.
   2. Existing saved view: Click **Save**.

> **Note**: Create new labels using **Add Label** (must be manually applied afterward).

### Remove Saved View Labels

A label may be removed from a saved view in two locations:

* The **Library List** for removing multiple saved views simultaneously (see **Remove Labels via Library List**).
* The saved view dialog (see **Saved View Dialogs**) for removing a single saved view (see **Remove Labels via Library Dialogs**).

### Label a Synthetics Agent

Labels can be applied to Synthetics agents in two locations:

#### Label Agent in Agents List

To label an existing agent:

1. Open the Agent Management page from the portal's main nav menu.
2. On one of the **Agents List Tabs**, check the boxes for each agent you want to label. The **Apply Labels** field will show the labels already applied to the selected agents.
3. Click **Apply Labels** for a filterable dropdown of available labels (see **Label Controls**).
4. Highlighted labels are already applied. Click twice on a highlighted label to apply it to the new set of agents.
5. Click once to apply a non-highlighted label.
6. Click outside to close the dropdown. The labels will appear in the **Name** column of the **Agents** list.

![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(611).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=%2B2IGqMqAZxmdgEwC6cPldNs7xOz0vwN2cG0BBUY3YwE%3D)

#### Label a Private Agent

To label a private agent on the **Private Agents** tab or in its **Configure Agent Dialog**:

1. Open the Agent Management page from the portal's main nav menu.
2. On the **Private Agents** tab, click the agent to which you want to apply a label.
3. In the **Agent Details Sidebar**, click the **Configure** button to open the **Configure Agent** dialog. The **Labels** field will show any labels that have already been applied to the agent.
4. Click Labels for a filterable dropdown of available labels. Labels already applied to this agent will be highlighted.
5. Select one or more labels to apply to the test. Each selected label will appear in the **Labels** field.
6. Click outside to close the dropdown.

   > **Note:** Remove labels by clicking the **X** next to them.
7. Click **Save** to apply changes and return to the Agent Management page.

### Remove Synthetics Agent Labels

To remove a label from one or more synthetic agents:

1. Open the Agent Management page from the portal's main nav menu.
2. Select each agent that you want to remove a label from. The **Apply Labels** field will display all labels applied to the selected agents.
3. Click the **X** next to any label to remove it from all selected agents
4. Alternatively, click **Clear Labels** to remove all labels from those agents.

### Using Labels with RBAC

Labels can be applied to Kentik’s Role-based Access Control (RBAC) roles for some resource types. For instance, a role’s permission could grant a user access to only those Library dashboards with a specific label.

Kentik resources that support RBAC role labels:

* Devices
* Dashboards
* Saved Views
* Credentials

> **Notes:**
>
> * See **Label an RBAC Role**for instructions on adding labels to roles.
> * Refer to **About Kentik RBAC**for more about using RBAC in the Kentik portal.

#### Label an RBAC Role

To make a role's permissions contingent on labels:

1. Go to **Organization Settings** » **RBAC Management** from the navbar to open the RBAC Management page.
2. In the **Roles** list, find the role you wish to label. If needed, use the **Search** field to narrow the roles.
3. Click **Edit** on a role’s row to open the Manage RBAC Role page.
4. On the **Permissions** tab, filter the **Permissions** list to show the permission group corresponding to the role’s focus (e.g., "Dashboards").
5. Enable the desired permission. Only permissions with Full Access or Label Access displayed on the right of the permission’s row can use labels to allow/restrict access.
6. Click the permission's **Access** drop-down and select **Label Access**.
7. Click **Add Label** for a filterable dropdown of available labels, then do one of the following:

   1. Click a single label to apply it to the permission.
   2. Use the **Search** field to filter the list and click **Select All** to apply all labels.
8. The applied labels will appear next to the **Add Label** dropdown. Repeat to apply additional labels to the permission.
9. Click **Save** to apply changes to the role.
10. Click **Cancel** to return to the RBAC Management page.    
    ![](https://cdn.us.document360.io/082e25b5-afce-42d4-8f47-70bd3f1d02b7/Images/Documentation/image(615).png?sv=2022-11-02&spr=https&st=2026-05-12T09%3A17%3A19Z&se=2026-05-12T09%3A44%3A19Z&sr=c&sp=r&sig=%2B2IGqMqAZxmdgEwC6cPldNs7xOz0vwN2cG0BBUY3YwE%3D)
